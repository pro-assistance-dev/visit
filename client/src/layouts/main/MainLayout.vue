<template>
  <div class="main">
    <div style="flex: 1">
      <slot />
    </div>
  </div>
</template>

<script setup lang="ts">
const fullscreen: Ref<boolean> = ref(window.matchMedia('(max-width: 1200px)').matches);
const handleResize = () => {
  fullscreen.value = window.matchMedia('(max-width: 1200px)').matches;
};

onMounted(() => {
  window.addEventListener('resize', handleResize);
});

onUnmounted(() => window.removeEventListener('resize', handleResize));
</script>

<style scoped>
.main {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>
