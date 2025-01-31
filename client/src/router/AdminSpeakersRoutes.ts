import AdminSpeakerPage from '@/components/admin/Speakers/AdminSpeakerPage.vue';
import AdminSpeakersList from '@/components/admin/Speakers/AdminSpeakersList.vue';

export default [
  {
    path: '/admin/speakers',
    name: 'AdminSpeakersList',
    component: AdminSpeakersList,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/speakers/new',
    name: 'AdminSpeakerPageNew',
    component: AdminSpeakerPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
  {
    path: '/admin/speakers/:id',
    name: 'AdminSpeakerPageEdit',
    component: AdminSpeakerPage,
    meta: {
      layout: 'AdminLayout',
    },
  },
];
