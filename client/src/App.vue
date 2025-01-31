<template>
  <component :is="$route.meta.layout || 'MainLayout'" v-if="mounted">
    <router-view />
  </component>
</template>

<script lang="ts" setup>
import MainLayout from '@/layouts/main/MainLayout.vue';

import AdminLayout from '@/layouts/adminLayout/AdminLayout.vue';

const components = {
  MainLayout,
  AdminLayout,
};
const mounted: Ref<boolean> = ref(false);

onBeforeMount(async (): Promise<void> => {
  mounted.value = true;
});
</script>
