<template>
  <div :dragenter="drag" class="container" :style="containerStyle">
    <div
      v-for="m in scheduler.timelines"
      class="five-min"
      :style="lineStyle"
      :class="{ 'selected-time': scheduler.activeTime === Time.ConcatHM(hour, m) }"
    >
      <div class="five-min-title">
        {{ Time.ConcatHM(hour, m) }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  hour: {
    type: Number,
    default: 0,
  },
});

const drag = () => {
  console.log('dragHor');
};
const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);

const containerStyle = computed(() => {
  return {
    height: scheduler.value.getHourHeight() + 'px',
  };
});

const lineStyle = computed(() => {
  return {
    height: scheduler.value.timelinePxHeight - 1 + 'px',
  };
});
</script>

<style lang="scss" scoped>
.container {
  box-sizing: border-box;
  width: 100%;
  z-index: 0;
}

.container:nth-child(2) {
  margin-top: 10px;
}

.selected-time {
  // color: red;
  // background: red;
  // border-color: red;
  box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.7);
}

.five-min {
  position: relative;
  border-top: 1px solid #c4c4c4;
  display: flex;
  justify-content: left;
  align-items: start;
  font-size: 14px;
  color: #343e5c;
  pointer-events: auto;
}

.five-min-title {
  background: #dff2f8;
  font-size: 14px;
  color: #343e5c;
  margin-top: -8px;
  padding: 0 5px;
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
}
</style>
