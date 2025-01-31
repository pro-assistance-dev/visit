<template>
  <div class="empty-space" :style="externalHeight" />
  <div :id="session.id" class="session" :style="cardStyle" @dblclick.stop="edit">
    <div v-if="showInfo" class="session-title" draggable="true" @dragend.self="endDrag" @dragstart.self="startDrag">
      <svg class="icon-move">
        <use xlink:href="#move"></use>
      </svg>
      {{ session.getMacroInfo() }}
      <svg class="icon-session-edit" @click="edit">
        <use xlink:href="#session_edit"></use>
      </svg>
      <div>
        <div>
          {{ session.name }}
        </div>
        <div v-for="chair in session.chairs">
          {{ chair.human.getFullName() }}
        </div>
      </div>
    </div>
    <div class="bottom-drag" draggable="true" @dragstart.self="startDragEnd" @dragend.self="endDrag">
      <div class="field"></div>
      <div class="color-line"></div>
    </div>
  </div>
  <Plus />
  <Move />
  <SessionSettings />
</template>

<script setup lang="ts">
import Move from '@/assets/svg/Move.svg';
import Plus from '@/assets/svg/Plus.svg';
import SessionSettings from '@/assets/svg/SessionSettings.svg';
import Schedule from '@/classes/Schedule';
import Session from '@/classes/Session';
const props = defineProps({
  session: {
    type: Object as PropType<Session>,
    required: true,
  },
  prevSession: {
    type: Object as PropType<Session | undefined>,
  },
  schedule: {
    type: Object as PropType<Schedule>,
    required: true,
  },
  externalHeight: {
    type: Number as PropType<number>,
    default: 0,
  },
});

const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);
const dragger: ComputedRef<ScheduleDragger> = computed(() => Provider.store.getters['schedules/dragger']);
let moved: Session[] = [];
const hoverInfo = ref(false);

watch(
  () => [dragger.value.time, dragger.value.schedule],
  () => {
    if (props.session.id !== dragger.value.dragItem.id) {
      return;
    }
    // moveToSchedule();
    const newMovedPerfoms = dragger.value.schedule.moveSession(props.session, scheduler.value, dragger.value);
    if (newMovedPerfoms.length > 0) {
      moved = newMovedPerfoms;
    }
  }
);

// const moveToSchedule = () => {
//   if (dragger.value.schedule.id === props.schedule.id) {
//     return;
//   }
//   event.value.copyPerfomToSchedule(props.schedule, props.session, dragger.value.schedule.id);
//   scheduler.value.sortByStartTime(dragger.value.schedule.perfoms);
// };
const startDrag = (e: DragEvent) => {
  dragger.value.start(e, props.session, props.schedule);
};

const startDragEnd = (e: DragEvent) => {
  dragger.value.stretchDown(e, props.session);
};

const endDrag = async () => {
  props.session.show = true;
  // event.value.removeDuplicatesPerfoms(props.perfom.id, dragger.value.schedule.id);
  dragger.value.end();
  //
  scheduler.value.sortByStartTime(props.schedule.sessions);
  //
  for (let index = 0; index < moved.length; index++) {
    await SessionsStore.Update(moved[index]);
  }
  moved = [];
};

const dragEndTime = async () => {
  if (Time.Gt(props.session.startTime, scheduler.value.activeTime)) {
    scheduler.value.endDrag();
    return;
  }
  if (!scheduler.value.timeIsFree(props.session.startTime, scheduler.value.activeTime, props.session.id, props.schedule.sessions)) {
    scheduler.value.endDrag();
    return;
  }
  props.session.endTime = scheduler.value.activeTime;
  await update();
  scheduler.value.endDrag();
};

const emits = defineEmits(['remove', 'changeTime', 'edit']);
// const showEditModal = ref(false);

const edit = async () => {
  emits('edit');
};

const update = async () => {
  await SessionsStore.Update(props.session);
};

const cardStyle = computed(() => {
  const h = (scheduler.value.timelinePxHeight * props.session.getDuration()) / scheduler.value.timelineStep - 4;
  return {
    height: computed(() => h + 'px').value,
  };
});
const showInfo = computed(() => {
  return scheduler.value.timelineStep === '5' || scheduler.value.timelineStep === '15';
});
</script>

<style lang="scss" scoped>
.bottom-drag {
  position: absolute;
  z-index: 1000;
  bottom: -10px;
  right: 0px;
  height: 20px;
  width: 60px;
  &:hover {
    cursor: row-resize;
  }
}

.field {
  height: calc(100% - 2px);
  width: 100%;
  opacity: 0;
}

.color-line {
  margin-top: -11px;
  height: 0px;
  width: 100%;
  border-bottom: 3px solid #006bb4;
}

.hidden {
  display: none;
}

.session {
  position: relative;
  margin: 0 1px;
  background: #ffffff;
  border: 1px solid #343e5c;
  border-top-right-radius: 3px;
  border-bottom-right-radius: 3px;
  z-index: 4;
  width: calc(100% - 4px);
  &:hover {
    cursor: pointer;
    // border: 2px solid white;
    box-shadow: 2px 2px 5px #000; /* Параметры тени */
    border: 1px solid #5178ec;
  }
}

.session-title {
  position: absolute;
  width: 100%;
  box-sizing: border-box;
  color: #006bb4;
  cursor: pointer;
  height: 80px;
}

.session-title:hover {
  z-index: 20;
  background: #dff2f8;
  cursor: pointer;
  border-bottom: 1px solid #379fff;
}

.session-body {
  width: 100%;
  // overflow: hidden;
}

.empty-space {
  height: 0px;
}

.icon-session-edit {
  position: absolute;
  top: 49px;
  right: 8px;
  width: 24px;
  height: 24px;
  fill: #006bb4;
  cursor: pointer;
  transition: 0.3s;
}

.icon-session-edit:hover {
  fill: #379fff;
  transform: scale(1.2);
}

.icon-plus {
  position: absolute;
  top: 7px;
  right: 40px;
  width: 28px;
  height: 28px;
  fill: #00b5a4;
  cursor: pointer;
}
.icon-plus:hover {
  transform: scale(1.2);
  transition: 0.3s;
}

.session-title:hover > .icon-move {
  visibility: visible;
}

@keyframes ripple {
  0% {
    transform: scale(1, 1);
  }
  50% {
    transform: scale(1.15, 1.15);
  }
}

.icon-move {
  position: absolute;
  top: 9px;
  right: 0px;
  visibility: hidden;
  width: 24px;
  height: 24px;
  fill: #dff2f8;
  stroke: #747474;
  animation-name: ripple;
  animation-duration: 1s;
  animation-iteration-count: infinite;
  padding-right: 8px;
}
</style>
