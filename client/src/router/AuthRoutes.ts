import RefreshPasswordPage from '@/components/RefreshPasswordPage.vue';

export default [
  {
    path: '/restore/password/:userId/:uniqueId',
    name: 'RefreshPasswordPage',
    meta: { title: 'Сброс пароля' },
    component: RefreshPasswordPage,
  },
];
