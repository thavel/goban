import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import router from './common/router'
import 'vuetify/dist/vuetify.min.css'

Vue.config.productionTip = false

Vue.use(Vuetify, {
  theme: {
    primary: '#1b1b1b',
    secondary: '#1f2838',
    accent: '#8896b2',
    error: '#FF5252',
    info: '#2196F3',
    success: '#4CAF50',
    warning: '#FFC107'
  },
  options: {
    customProperties: true
  }
})

Vue.prototype.$api = 'http://localhost:8000';

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
