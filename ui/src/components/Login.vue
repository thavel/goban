<template>
<v-container fluid fill-height class="login">
  <v-layout align-center justify-center>
    <v-flex xs12 sm8 md3 class="text-xs-center">
      <img src="../assets/logo.png" class="login-logo"/>
      <v-card class="elevation-12 round" max-width="500px">
        <v-toolbar color="secondary" dark>
          <v-toolbar-title>Authentication</v-toolbar-title>
        </v-toolbar>
        <form ref="loginForm" @submit.prevent="login" lazy-validation>
          <v-card-text>
            <v-text-field
              prepend-icon="person"
              v-model="input.user"
              name="user"
              label="user"
              type="text"
              :rules="rules">
            </v-text-field>
            <v-text-field
              prepend-icon="lock"
              v-model="input.password"
              name="password"
              label="Password"
              type="password"
              :rules="rules">
            </v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="secondary" type="submit" round>Login</v-btn>
          </v-card-actions>
        </form>
      </v-card>
    </v-flex>
  </v-layout>
  <v-snackbar v-model="error" color="error" bottom>
    Invalid credentials
    <v-btn flat @click="error = false">ok</v-btn>
  </v-snackbar>
</v-container>
</template>

<script>
import { Auth } from '../common/auth';
import axios from 'axios';

export default {
  data: () => ({
    errorTimeout: 3000,
    error: false,
    input: {
      user: null,
      password: null
    },
    rules: [
      v => !!v || 'Field required'
    ]
  }),
  mounted() {
    Auth.clean();
  },

  methods: {
    login: async function() {
      try {
        var res = await axios.post(this.$api + '/auth/token', this.input);
        Auth.set(res.data.token);
        this.$router.replace({path: '/'});
      } catch(e) {
        this.error = true;
      }
    }
  }
}
</script>

<style>
.login {
  background: linear-gradient(var(--v-secondary-base), var(--v-accent-base));
}
.login-logo {
  margin-top: -75px;
  max-width: 250px;
  margin-bottom: 15px;
}
.login .v-card.round {
  border-radius: 18px;
}
</style>
