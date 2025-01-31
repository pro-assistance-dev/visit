<template>
  <component :is="$route.meta.layout || 'MainLayout'" v-if="mounted">
    <router-view />
  </component>
</template>

<script lang="ts" setup>
import MainLayout from '@/layouts/main/MainLayout.vue';

const components = {
  MainLayout,
};
const mounted: Ref<boolean> = ref(false);

onBeforeMount(async (): Promise<void> => {
  mounted.value = true;
});
</script>
