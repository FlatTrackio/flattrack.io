import Vue from 'vue'
import Router from 'vue-router'
import home from '@/views/home'
import about from '@/views/about'
import contact from '@/views/contact'
import unknownPage from '@/views/unknown-page'
import privacyPolicy from '@/views/privacy-policy'

Vue.use(Router)

export default new Router({
  scrollBehavior: () => {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '/',
      name: 'home',
      component: home
    },
    {
      path: '/moreinfo',
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
      path: '/privacy-policy',
      name: 'privacy',
      component: privacyPolicy
    },
    {
      path: '/privacy',
      redirect: '/privacy-policy'
    },
    {
      path: '*',
      name: 'unknown-page',
      component: unknownPage
    }
  ]
})
