<template>
  <div class="main" v-if="mounted">
    <div class="program">
      <EventHeader />
      <EventStreamChat v-if="PF.Auth.Status.IsAuth" />
      <div class="header" v-else>
        <div class="event-layout">
          <EventRegistration />
        </div>
      </div>
      <div class="event-layout">
        <div class="program-table">
          <EventSchedules />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
const mounted: Ref<boolean> = ref(false);
onBeforeMount(async () => {
  await Store.Events.GetBySlug(PF.H.Router.Slug());
  mounted.value = true;
});
</script>
<style scoped>
.header {
  width: 100%;
  background: no-repeat url('@/assets/img/header_bg2.png');
  background-size: cover;
  display: flex;
  align-items: center;
  height: auto;
  padding: 50px 0;
}
.program {
  background: no-repeat url('@/assets/img/event_layout_BG.png');
  background-size: cover;

  &-table {
    text-align: center;
  }
}

@media screen and (min-width: 1024px) {
  .program-list {
    display: none;
  }
}
</style>
