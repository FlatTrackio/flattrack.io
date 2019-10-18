import Vue from 'vue'
import Router from 'vue-router'
import home from '@/components/public/home'
import about from '@/components/public/about'
import contact from '@/components/public/contact'
import unknownPage from '@/components/public/unknown-page'
import privacyPolicy from '@/components/public/privacy-policy'

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
