/**
 * router/index.ts
 *
 * Automatic routes for `./src/pages/*.vue`
 */

// Composables
import { createRouter, createWebHistory } from 'vue-router/auto'
import Index from '@/pages/index.vue'

const routes = [
  {
    path: '/v/',
    redirect: '/v/main'
  },
  {
    path: '/v/main',
    name: 'Index',
    component: Index,
  },
  {
    path: '/v/upload',
    name: 'Upload',
    component: () => import('@/pages/device/upload.vue'),
  },
  {
    path: '/v/device',
    name: 'Device',
    component: () => import('@/pages/device/device_manage.vue'),
  },
  {
    path: '/v/updata',
    name: 'Updata',
    component: () => import('@/components/node/update_node.vue'),
  },
  {
    path: '/v/home',
    name: 'Home',
    component: () => import('@/pages/home/home.vue'),
  },
  {
    path: '/v/watcher',
    name: 'Watcher',
    component: () => import('@/pages/watcher/watcher_manage.vue'),
  },
  {
    path: '/v/phone',
    name: 'Phone',
    component: () => import('@/pages/phone/phone_manage.vue'),
  },
  {
    path: '/v/game',
    name: 'GameManage',
    component: () => import('@/pages/game/game_manage.vue'),
    
  }

]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})


router.onError((err, to) => {
  if (err?.message?.includes?.('Failed to fetch dynamically imported module')) {
    if (!localStorage.getItem('vuetify:dynamic-reload')) {
      console.log('Reloading page to fix dynamic import error')
      localStorage.setItem('vuetify:dynamic-reload', 'true')
      location.assign(to.fullPath)
    } else {
      console.error('Dynamic import error, reloading page did not fix it', err)
    }
  } else {
    console.error(err)
  }
})

router.isReady().then(() => {
  localStorage.removeItem('vuetify:dynamic-reload')
})

export default router
