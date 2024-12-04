import { createRouter, createWebHistory } from 'vue-router'

import search from '../components/Main/search_key.vue'
import state from '../components/Main/state.vue'

const routerHistory = createWebHistory()
// createWebHashHistory hash 路由
// createWebHistory history 路由
// createMemoryHistory 带缓存 history 路由
const router = createRouter({
  history: routerHistory,
  routes: [
    {
      path:'/search',
      name:'search',
      component:search
    },
    {
      path:'/state',
      name:'state',
      component:state
    } ,
    {
      path: '/',
      redirect:'state'
    }
  ]
})
 
export default router