import { NavigationGuardNext, RouteLocationNormalized } from 'vue-router';

import EventPage from '@/components/event/EventPage.vue';
// import { isAuthorized } from '@/router/index';

export default [
  {
    path: '/events/:slug',
    name: 'EventPage',
    component: EventPage,
    // beforeEnter(to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext): void {
    //     isAuthorized(next);
    // },
  },
];
