export default [
  {
    path: '/',
    name: 'home',
    component: () => import('@/frontend/views/home.vue')
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('@/frontend/views/about.vue')
  },
  {
    path: '/contact',
    name: 'contact',
    component: () => import('@/frontend/views/contact.vue')
  },
  {
    path: '/privacy-policy',
    name: 'privacy',
    component: () => import('@/frontend/views/privacy-policy.vue')
  },
  {
    path: '/privacy',
    redirect: '/privacy-policy'
  },
  {
    path: '*',
    name: 'unknown-page',
    component: () => import('@/frontend/views/unknown-page.vue')
  }
]
