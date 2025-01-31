import AdminPerfomPage from '@/components/admin/Perfoms/AdminPerfomPage.vue';
import AdminPerfomsList from '@/components/admin/Perfoms/AdminPerfomsList.vue';

export default [
  {
    path: '/admin/perfoms',
    name: 'AdminPerfomsList',
    component: AdminPerfomsList,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/perfoms/new',
    name: 'AdminPerfomPageNew',
    component: AdminPerfomPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/perfoms/:id',
    name: 'AdminPerfomPageEdit',
    component: AdminPerfomPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
];
