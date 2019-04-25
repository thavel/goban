import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../components/Main'
import Login from '../components/Login'
import Absences from '../components/Absences'

Vue.use(VueRouter)

export default new VueRouter({
  routes: [
    {
      path: '/',
      component: Main,
      children: [
        {
          path: '/',
          redirect: '/absences'
        },
        {
          path: 'absences',
          component: Absences
        }
      ]
    },
    {
      path: '/login',
      component: Login
    }
  ]
})
