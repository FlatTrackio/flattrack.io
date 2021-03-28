export default [
  {
    path: '/',
    name: 'home',
    component: () => import('@//views/home.vue')
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('@//views/about.vue')
  },
  {
    path: '/contact',
    name: 'contact',
    component: () => import('@//views/contact.vue')
  },
  {
    path: '/privacy-policy',
    name: 'privacy',
    component: () => import('@//views/privacy-policy.vue')
  },
  {
    path: '/privacy',
    redirect: '/privacy-policy'
  },
  {
    path: '*',
    name: 'unknown-page',
    component: () => import('@//views/unknown-page.vue')
  }
]
