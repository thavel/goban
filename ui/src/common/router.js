import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../components/Main'
import Login from '../components/Login'
import Calendar from '../components/Calendar'

Vue.use(VueRouter)

export default new VueRouter({
  routes: [
    {
      path: '/',
      component: Main,
      children: [
        {
          path: '/',
          redirect: '/calendar'
        },
        {
          path: 'calendar',
          component: Calendar
        }
      ]
    },
    {
      path: '/login',
      component: Login
    }
  ]
})
