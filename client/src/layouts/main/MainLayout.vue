<template>
  <div class="main">
    <!-- <Header /> -->
    <div style="flex: 1">
      <slot />
    </div>
    <!-- <Footer /> -->
    <!-- <AuthModal v-if="modal.visible" @action="authAct" /> -->
  </div>
</template>

<script setup lang="ts">
// const modal: ComputedRef<AuthModal> = computed(() => Provider.store.getters['auth/modal']);
// const form: ComputedRef<AuthModal> = computed(() => Provider.store.getters['auth/form']);
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
