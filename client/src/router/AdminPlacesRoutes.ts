import AdminPlacePage from '@/components/admin/Places/AdminPlacePage.vue';
import AdminPlacesList from '@/components/admin/Places/AdminPlacesList.vue';

export default [
  {
    path: '/admin/places',
    name: 'AdminPlacesList',
    component: AdminPlacesList,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/places/new',
    name: 'AdminPlacePageNew',
    component: AdminPlacePage,
  },
];
