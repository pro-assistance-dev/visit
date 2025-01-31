<template>
  <div class="day">
    <div class="day-date">{{ $dateTimeFormatter.format(eventDay.date) }}</div>
    <Hour :dragenter="drag" v-for="h in scheduler.getHoursArr()" :hour="h" />
  </div>
</template>

<script setup lang="ts">
import EventDay from '@/classes/EventDay';

defineProps({
  eventDay: {
    type: Object as PropType<EventDay>,
    required: true,
  },
});
const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);
const drag = () => {
  console.log('dragHor');
};
</script>

<style lang="scss" scoped>
.day {
  position: relative;
  margin-bottom: 20px;
}

.day-date {
  width: 100%;
  height: 50px;
  background: #006bb4;
  color: #ffffff;
  font-size: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 51px;
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
}
</style>
