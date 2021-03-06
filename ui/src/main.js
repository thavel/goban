import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import VCalendar from 'v-calendar'
import router from './common/router'
import auth from './common/auth'
import 'vuetify/dist/vuetify.min.css'

Vue.config.productionTip = false

Vue.use(Vuetify, {
  theme: {
    primary: '#1a3300',
    secondary: '#284d00',
    tertiary: '#356600',
    accent: '#4f9900',
    error: '#FF5252',
    info: '#2196F3',
    success: '#4CAF50',
    warning: '#FFC107'
  },
  options: {
    customProperties: true
  }
})

Vue.use(VCalendar, {componentPrefix: 'vc'});

Vue.prototype.$auth = auth;
Vue.prototype.$api = window.location.href.split("/ui/")[0];
//Vue.prototype.$api = 'http://localhost:8000';

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')
