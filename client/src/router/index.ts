import { createRouter, createWebHistory, NavigationGuardNext, RouteLocationNormalized, RouteRecordRaw } from 'vue-router';

import MainPage from '@/components/main/MainPage.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'MainPage',
    component: MainPage,
    meta: { main: true, title: 'Главная' },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
