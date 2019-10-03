import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/public/home'
import about from '@/components/public/about'
import contact from '@/components/public/contact'
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
      path: '/moreinfo',
      name: 'home-more-info',
      component: home,
      redirect: '/'
    },
    {
      path: '/about',
      name: 'about',
      component: about
    },
    {
      path: '/contact',
      name: 'contact',
      component: contact
    },
    {
      path: '*',
      name: 'unknown-page',
      component: unknownPage
    }
  ]
})
