import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  scrollBehavior: () => {
    return { x: 0, y: 0 }
  },
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/home')
    },
    {
      path: '/about',
      name: 'about',
      component: () => import('@/views/about')
    },
    {
      path: '/contact',
      name: 'contact',
      component: () => import('@/views/contact')
    },
    {
      path: '/privacy-policy',
      name: 'privacy',
      component: () => import('@/views/privacy-policy')
    },
    {
      path: '/privacy',
      redirect: '/privacy-policy'
    },
    {
      path: '*',
      name: 'unknown-page',
      component: () => import('@/views/unknown-page')
    }
  ]
})
