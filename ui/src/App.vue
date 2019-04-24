<template>
<v-app id="app">
  <div id="loader"><div></div><div></div><div></div><div></div></div>
  <v-content>
    <router-view/>
  </v-content>
</v-app>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  padding-top: 60px;
}

#loader {
  display: inline-block;
  position: relative;
  width: 64px;
  height: 64px;
  margin: auto;
}
#loader div {
  position: absolute;
  width: 14px;
  height: 14px;
  border-radius: 50%;
  background: var(--v-primary-base);
  animation-timing-function: cubic-bezier(0, 1, 1, 0);
}
#loader div:nth-child(1) {
  left: 6px;
  animation: loader-step1 0.6s infinite;
}
#loader div:nth-child(2) {
  left: 6px;
  animation: loader-step2 0.6s infinite;
}
#loader div:nth-child(3) {
  left: 26px;
  animation: loader-step2 0.6s infinite;
}
#loader div:nth-child(4) {
  left: 45px;
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
    transform: translate(19px, 0);
  }
}
</style>
