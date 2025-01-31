import { createRouter, createWebHistory, NavigationGuardNext, RouteLocationNormalized, RouteRecordRaw } from 'vue-router';

import MainPage from '@/components/Main/MainPage.vue';
import AuthRoutes from '@/router/AuthRoutes';
import EventsRoutes from '@/router/EventsRoutes';
import indexAdminRoutes from '@/router/indexAdminRoutes';

//
// export const isAuthorized = (next: NavigationGuardNext): void => {
//   const user = localStorage.getItem('user');
//   if (user) {
//     // store.commit('auth/setIsAuth', true);
//   }
//   next();
// };
//
// export const authGuard = (): void => {
//   if (!TokenService.isAuth()) {
//     router.push('/');
//   }
// };
//
// export const adminGuard = (): void => {
//   console.log('UserService.isAdmin()', !UserService.isAdmin());
//   if (!UserService.isAdmin()) {
//     router.push('/');
//   }
// };

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/main',
  },
  {
    path: '/main',
    name: 'MainPage',
    component: MainPage,
    meta: { main: true, title: 'Главная' },
  },
  ...EventsRoutes,
  ...AuthRoutes,
  ...indexAdminRoutes,
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
