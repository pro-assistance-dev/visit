<template>
  <div class="main" v-if="mounted">
    <!-- <div style="flex: 1"> -->
    <!--   <slot /> -->
    <!-- </div> -->
  </div>
</template>

<script lang="ts" setup>
const mounted: Ref<boolean> = ref(false);
onBeforeMount(async () => {
  await Store.Events.GetBySlug(PF.H.Router.Slug());
  // await Provider.store.dispatch('events/getMessages');
  mounted.value = true;
});
</script>

<style scoped>
.main {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}
</style>
