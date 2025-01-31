<template>
  <AddScheduleItemModal></AddScheduleItemModal>
  <!-- <RightSliderContainer title-close="Показать список докладов" title-open="Скрыть список докладов"> -->
  <!--   <template #header> -->
  <!--     <svg class="icon-plus2"> -->
  <!--       <use xlink:href="#plus"></use> -->
  <!--     </svg> -->
  <!-- </template> -->
  <!-- <AdminEventPerfomsList /> -->
  <!-- </RightSliderContainer> -->

  <div class="global-wraper" @drag="moveCursor" @dblclick="addObj" @dragend="end">
    <div class="control-panel">
      <div class="time-menu">
        Временной шаг расписания, мин:
        <div class="interval-selection">
          <div v-for="step in Object.values(TimelineSteps)" :key="step" class="interval" @click="scheduler.setTimelineStep(step)">
            {{ step }}
          </div>
        </div>
      </div>
      <div class="toggle">
        <svg v-if="opts.perfomsShowed" class="activated-icon non-active" @click="opts.perfomsShowed = false">
          <use xlink:href="#non-active" />
        </svg>
        <svg v-else class="activated-icon active" @click="opts.perfomsShowed = true">
          <use xlink:href="#active" />
        </svg>
        Скрыть выступления
      </div>
    </div>

    <div class="time-line">
      <TimeDay v-for="eventDay in event.eventDays" :key="eventDay.id" :event-day="eventDay" />
    </div>
    <template v-for="eventDay in event.eventDays" :key="eventDay.id">
      <div class="place-container" :style="placeContainerStyle">
        <PlaceCol v-for="place in event.places" :key="place.id" :place="place" :schedule="event.getSchedule(eventDay.id, place.id)" />
      </div>
    </template>
  </div>
  <Plus />
  <Toggle />
</template>
<script setup lang="ts">
import Plus from '@/assets/svg/Plus.svg';
import Toggle from '@/assets/svg/Toggle.svg';
import Event from '@/classes/Event';
import ScheduleDragger from '@/classes/ScheduleDragger';

// const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);
// const dragger: ComputedRef<ScheduleDragger> = computed(() => Provider.store.getters['schedules/dragger']);
// const opts: ComputedRef<ScheduleDragger> = computed(() => Provider.store.getters['schedules/options']);
//
// const event: Event = Store.Events.Item;
//
// const placeContainerStyle = computed(() => {
//   return {
//     height: scheduler.value.getScheduleHeight() + 63 + 'px',
//   };
// });
//
// const moveCursor = (e: any) => {
//   if (!dragger.value.active()) {
//     return;
//   }
//   document.elementsFromPoint(e.x, e.y).forEach((e: HTMLElement) => {
//     if (e.className === 'five-min') {
//       dragger.value.time = e.innerText;
//     }
//     if (e.className === 'place') {
//       dragger.value.setSchedule(event.getScheduleById(e.id));
//     }
//   });
// };
</script>
<style lang="scss">
.toggle {
  display: flex;
  justify-content: left;
  align-items: center;
}

.control-panel {
  display: flex;
  justify-content: left;
  align-items: center;
  height: 40px;
}

.pos {
  position: relative;
}

.icon-plus2 {
  position: absolute;
  top: 15px;
  right: 15px;
  width: 28px;
  height: 28px;
  fill: #00b5a4;
  cursor: pointer;
}
.icon-plus:hover {
  transform: scale(1.2);
  transition: 0.3s;
}

.global-wraper {
  position: relative;
  overflow: auto;
  height: calc(100% - 30px);
  box-sizing: border-box;
}

.time-menu {
  display: flex;
  justify-content: left;
  align-items: center;
  height: 42px;
  margin: 0 10px;
  box-sizing: border-box;
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
  width: auto;
}

.place-container {
  display: flex;
  justify-content: left;
  height: calc(100% - 75px);
  box-sizing: border-box;
  margin-left: 60px;
  margin-top: 58px;
}

.time-line {
  position: absolute;
  height: 100%;
  width: 100%;
}

.interval-selection {
  display: flex;
  height: 42px;
  justify-content: space-between;
  align-items: center;
  width: 160px;
  padding: 0 10px;
}

.interval {
  border: 1px solid #343e5c;
  border-radius: 5px;
  height: 32px;
  width: 32px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: #ffffff;
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
}

.activated-icon {
  width: 40px;
  height: 40px;
  cursor: pointer;
  transition: 0.3s;
  margin-right: 10px;
}

.active {
  fill: #1979cf;
}

.non-active {
  fill: #c4c4c4;
}
</style>
