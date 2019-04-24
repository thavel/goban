<template>
<v-app id="app">
  <v-container fluid v-if="loading">
    <v-layout align-center justify-center row fill-height>
      <v-flex xs12>
        <v-layout align-center justify-center column fill-height>
          <div id="loader"><div></div><div></div><div></div><div></div></div>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
  <router-view v-if="!loading"/>
</v-app>
</template>

<script>
import { Auth } from '@/common/auth';
import axios from 'axios';

export default {
  name: 'app',
  data: () => ({
    loading: true
  }),
  mounted() {
    this.autoLogin();
  },

  methods: {
    autoLogin: async function() {
      let token = Auth.retrive();
      if (token == null) {
        this.failure();
        return;
      }
      try {
        let res = await axios.post(this.$api + '/auth/token', {token: token});
        Auth.set(res.data.token);
        this.success();
      } catch(e) {
        this.failure();
      }
    },
    success: function() {
      this.loading = false;
      this.$router.replace({path: '/'});
    },
    failure: function() {
      this.loading = false;
      this.$router.replace({path: '/login'});
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
#loader {
  display: inline-block;
  position: relative;
  width: 130px;
  height: 30px;
  margin: auto;
}
#loader div {
  position: absolute;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  background: var(--v-primary-base);
  animation-timing-function: cubic-bezier(0, 1, 1, 0);
}
#loader div:nth-child(1) {
  left: 12px;
  animation: loader-step1 0.6s infinite;
}
#loader div:nth-child(2) {
  left: 12px;
  animation: loader-step2 0.6s infinite;
}
#loader div:nth-child(3) {
  left: 52px;
  animation: loader-step2 0.6s infinite;
}
#loader div:nth-child(4) {
  left: 90px;
  animation: loader-step3 0.6s infinite;
}
@keyframes loader-step1 {
  0% {
    transform: scale(0);
  }
  100% {
    transform: scale(1);
  }
}
@keyframes loader-step3 {
  0% {
    transform: scale(1);
  }
  100% {
    transform: scale(0);
  }
}
@keyframes loader-step2 {
  0% {
    transform: translate(0, 0);
  }
  100% {
    transform: translate(38px, 0);
  }
}
</style>
