import AdminSponsorPage from '@/components/admin/Sponsors/AdminSponsorPage.vue';
import AdminSponsorsList from '@/components/admin/Sponsors/AdminSponsorsList.vue';

export default [
  {
    path: '/admin/sponsors',
    name: 'AdminSponsorsList',
    component: AdminSponsorsList,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/sponsors/new',
    name: 'AdminSponsorPageNew',
    component: AdminSponsorPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/sponsors/:id',
    name: 'AdminSponsorPageEdit',
    component: AdminSponsorPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
];
