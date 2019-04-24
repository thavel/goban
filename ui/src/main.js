import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import router from './common/router'
import 'vuetify/dist/vuetify.min.css'

Vue.config.productionTip = false

Vue.use(Vuetify, {
  theme: {
    primary: '#1976D2',
    secondary: '#424242',
    accent: '#82B1FF',
    error: '#FF5252',
    info: '#2196F3',
    success: '#4CAF50',
    warning: '#FFC107'
  },
  options: {
    customProperties: true
  }
})

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
