import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import EventView from '@/views/EventView.vue'
import GroupView from "@/views/GroupView.vue";
import AlertSettingsView from "@/views/AlertSettingsView.vue";
import AlertsView from "@/views/AlertsView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/events/:uuid',
      name: 'event',
      component: EventView
    },
    {
      path: '/groups/:uuid',
      name: 'group',
      component: GroupView
    },
    {
      path: '/settings/alerts',
      name: 'alerts_settings',
      component: AlertSettingsView
    },
    {
      path: '/alerts',
      name: 'alerts',
      component: AlertsView
    }
  ]
})

export default router
