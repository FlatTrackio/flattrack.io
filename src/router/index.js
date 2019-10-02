import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/public/home'
import unknownPage from '@/components/public/unknown-page'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },
    {
      path: '*',
      name: 'unknown-page',
      component: unknownPage
    },
    {
      path: '/moreinfo',
      name: 'home-more-info',
      component: home,
      redirect: '/'
    }
  ]
})
