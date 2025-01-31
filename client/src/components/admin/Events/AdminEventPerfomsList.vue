<template>
  <div
    draggable="true"
    @dragstart.stop="(e) => startDrag(e, perfom)"
    @dragend.stop="(e) => endDrag(e, perfom)"
    v-for="perfom in event.perfoms"
    :key="perfom.id"
  >
    <MiniPerfomCard :perfom="perfom" />
  </div>
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import Schedule from '@/classes/Schedule';
import Perfom from '@/classes/Perfom';
const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);
const event = Store.Events.Item;

const startDrag = (e: DragEvent, perfom: Perfom) => {
  scheduler.value.startDrag(e, perfom.startTime, perfom.endTime);
};

const movePerfom = async (perfom: Perfom, schedule: Schedule, startTime: string) => {
  const newStart = startTime;
  const newEnd = Time.AddMtoHM(newStart, perfom.getDuration());
  // movePerfom(newStart, newEnd);
  // if (!props.schedule.periodIsFree(newStart, newEnd, perfom.id)) {
  //   scheduler.value.endDrag();
  //   return;
  // }
  event.removePerfom(perfom.id);
  schedule.perfoms.push(perfom);
  perfom.setPeriod(newStart, newEnd);
  // props.session.sortPerfomsByStartTime();
  // await update();

  scheduler.value.endDrag();
};
const endDrag = (e: any, perfom: Perfom) => {
  const elements = document.elementsFromPoint(e.x, e.y);
  let startTime = '';
  let scheduleId = '';
  elements.forEach((e: HTMLElement) => {
    if (e.className === 'place') {
      scheduleId = e.id;
    }
    if (e.className === 'five-min') {
      startTime = e.innerText;
    }
  });
  const schedule = event.value.schedules.find((s: Schedule) => s.id === scheduleId);
  if (!schedule || !startTime) {
    return;
  }

  movePerfom(perfom, schedule, startTime);
};
</script>
<style lang="scss">
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
  width: 100%;
  height: 42px;
  margin-bottom: 0px;
  box-sizing: border-box;
  // -webkit-user-select: none; /* Safari */
  // -ms-user-select: none; /* IE 10 and IE 11 */
  // user-select: none; /* Standard syntax */
}

.place-container {
  display: flex;
  justify-content: left;
  height: calc(100% - 75px);
  box-sizing: border-box;
  margin-left: 60px;
  z-index: 1000;
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

// .wraper {
//   display: grid;
//   grid-template-columns: 100px 1fr 1fr 1fr 1fr;
//   height: 86vh;
//   width: 100%;
// }

// .title {
//   position: sticky;
//   top: 0;
//   left: 0;
//   z-index: 1;
//   display: grid;
//   grid-template-columns: 100px 1fr 1fr 1fr 1fr;
//   height: 60px;
//   width: 100%;
//   font-weight: bold;
//   font-size: 16px;
// }

// .time-line_title {
//   background: #ffffff;
//   color: #343e5c;
//   border-right: 1px solid #959595;
//   border-bottom: 2px solid #343e5c;
//   display: flex;
//   justify-content: center;
//   align-items: center;
// }

// .column1_title {
//   background: #1979cf;
//   color: #ffffff;
//   border-right: 1px solid #959595;
//   border-bottom: 2px solid #343e5c;
//   display: flex;
//   justify-content: center;
//   align-items: center;
// }
// .column2_title {
//   background: #f3911c;
//   color: #ffffff;
//   border-right: 1px solid #959595;
//   border-bottom: 2px solid #343e5c;
//   display: flex;
//   justify-content: center;
//   align-items: center;
// }
// .column3_title {
//   background: #660b19;
//   color: #ffffff;
//   border-right: 1px solid #959595;
//   border-bottom: 2px solid #343e5c;
//   display: flex;
//   justify-content: center;
//   align-items: center;
// }
// .column4_title {
//   background: #0aa249;
//   color: #ffffff;
//   border-right: 1px solid #959595;
//   border-bottom: 2px solid #343e5c;
//   display: flex;
//   justify-content: center;
//   align-items: center;
// }

// .column1 {
//   background: #dff2f8;
//   height: 100%;
//   border-right: 1px solid #959595;
// }

// .column2 {
//   background: #fff0df;
//   height: 100%;
//   border-right: 1px solid #959595;
// }

// .column3 {
//   background: #ffd2d9;
//   height: 100%;
//   border-right: 1px solid #959595;
// }

// .column4 {
//   background: #eefff5;
//   height: 100%;
//   border-right: 1px solid #959595;
// }

// .card-event {
//   background: #ffffff;
//   border: 1px solid #959595;
//   border-radius: 5px;
//   height: calc(300px - 2px);
//   margin: 1px 2px;
//   box-sizing: border-box;
// }

// :deep(.icon-plus) {
//   width: 24px;
//   height: 24px;
//   cursor: pointer;
//   fill: #00bea5;
// }

// .plus-button {
//   width: 100%;
//   border-radius: 0;
//   color: #00bea5;
//   background: #c1efeb;
//   height: 100%;
//   margin: 0;
// }
</style>
