import Vue from 'vue'
import VueRouter from 'vue-router'
import Main from '../components/Main'
import Login from '../components/Login'
import Dashboard from '../components/Dashboard'

Vue.use(VueRouter)

export default new VueRouter({
  routes: [
    {
      path: '/',
      component: Main,
      children: [
        {
          path: '',
          component: Dashboard
        },
      ]
    },
    {
      path: '/login',
      component: Login
    }
  ]
})
