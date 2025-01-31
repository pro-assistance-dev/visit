import AdminEventPage from '@/components/admin/Events/AdminEventPage.vue';
import AdminEventsList from '@/components/admin/Events/AdminEventsList.vue';

export default [
  {
    path: '/admin/events',
    name: 'AdminEventsList',
    component: AdminEventsList,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/events/new',
    name: 'AdminEventPageNew',
    component: AdminEventPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/events/:id',
    name: 'AdminEventPageEdit',
    component: AdminEventPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
];
