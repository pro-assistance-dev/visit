<template>
  <!-- <ContextWindow :x="x" :y="y" @close="contextMenuOpened = false" v-if="contextMenuOpened"> -->
  <!--   <Button @click="remove" text="Удалить" width="calc(100% - 10px)" height="32px" /> -->
  <!-- </ContextWindow> -->
  <div v-if="perfom.show" class="empty-space" :style="externalHeight" @contextmenu.stop />

  <div
    v-show="perfom.show"
    class="speaker"
    :style="cardStyle"
    draggable="true"
    @contextmenu.stop.prevent="openContextMenu"
    @dblclick.stop="edit"
    @click.stop
    @dragstart.stop="startDrag"
    @dragend.stop="endDrag"
  >
    <template v-if="showInfo">
      <div class="speaker-title">
        <svg class="icon-move">
          <use xlink:href="#move"></use>
        </svg>
        <div class="title-hidden">
          <div class="title-header">
            <PString :string="Time.GetPeriod(perfom.startTime, perfom.endTime)" width="auto" font-size="16px" />
            <PString :string="'&nbsp' + perfom.type + ', ' + perfom.format" width="auto" font-size="16px" />
          </div>
          <div class="title-footer">
            <div v-for="speaker in perfom.speakers" class="flex-item">
              {{ speaker.human.getInitialsName() }}
            </div>
          </div>
        </div>
      </div>
      <svg class="icon-speaker-edit" @click="edit">
        <use xlink:href="#speaker_edit"></use>
      </svg>
    </template>
    <div class="perfom-name">
      {{ perfom.name }}
    </div>
    <div class="bottom-drag" draggable="true" @dragstart.stop.self="startDragEnd" @dragend.stop.self="endDrag">
      <div class="field"></div>
      <div class="color-line"></div>
    </div>
  </div>
  <!-- <div class="bottom-drag" draggable="true" @click.stop @dragstart.stop="bottomDragStart" @dragend.stop="dragEndTime"></div> -->
  <SpeakerEdit />
  <Move />
</template>

<script setup lang="ts">
import Move from '@/assets/svg/Move.svg';
import SpeakerEdit from '@/assets/svg/SpeakerEdit.svg';
import Perfom from '@/classes/Perfom';
import Schedule from '@/classes/Schedule';
import ScheduleDragger from '@/classes/ScheduleDragger';
const contextMenuOpened = ref(false);
const x = ref(0);
const y = ref(0);

const openContextMenu = (e: MouseEvent) => {
  x.value = e.clientX; // get the mouse position relative to the element
  y.value = e.clientY;
  contextMenuOpened.value = true;
};

const remove = () => {
  contextMenuOpened.value = false;
  emits('remove');
};

const props = defineProps({
  perfom: {
    type: Object as PropType<Perfom>,
    required: true,
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
let movedPerfoms: Perfom[] = [];

watch(
  () => [dragger.value.time, dragger.value.schedule],
  () => {
    if (props.perfom.id !== dragger.value.dragItem.id) {
      return;
    }
    moveToSchedule();
    const newMovedPerfoms = dragger.value.schedule.movePerfom(props.perfom, scheduler.value, dragger.value);
    if (newMovedPerfoms.length > 0) {
      movedPerfoms = newMovedPerfoms;
    }
  }
);

const event = Store.Events.Item;
const emits = defineEmits(['edit', 'startDrag', 'remove']);

const startDrag = (e: DragEvent) => {
  dragger.value.start(e, props.perfom, props.schedule);
};

const startDragEnd = (e: DragEvent) => {
  dragger.value.stretchDown(e, props.perfom);
};

const moveToSchedule = () => {
  if (dragger.value.schedule.id === props.schedule.id) {
    return;
  }
  event.copyPerfomToSchedule(props.schedule, props.perfom, dragger.value.schedule.id);
  scheduler.value.sortByStartTime(dragger.value.schedule.perfoms);
};

const endDrag = async () => {
  props.perfom.show = true;
  event.removeDuplicatesPerfoms(props.perfom.id, dragger.value.schedule.id);
  dragger.value.end();
  //
  scheduler.value.sortByStartTime(props.schedule.perfoms);
  //
  // for (let index = 0; index < movedPerfoms.length; index++) {
  //   await Provider.store.dispatch('perfoms/update', movedPerfoms[index]);
  // }
  movedPerfoms = [];
};

const edit = async () => {
  emits('edit');
};

const update = async () => {
  await PerfomsStore.Update(props.perfom);
};

const cardStyle = computed(() => {
  const h = (scheduler.value.timelinePxHeight * props.perfom.getDuration()) / scheduler.value.timelineStep - 4;
  return {
    height: computed(() => h + 'px').value,
    backgroundColor: props.perfom.color,
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

.perfom-name {
  display: flex;
  justify-content: left;
  align-items: center;
  font-size: 18px;
  overflow: hidden;
  color: #f5f6f8;
  padding: 5px;
  font-style: oblique;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.6);
}

.title-header {
  display: flex;
  justify-content: left;
  align-items: center;
  font-size: 16px;
  width: 100%;
  transition: 0.2s;
  height: 24px;
  color: #f5f6f8;
}

.title-time {
  font-size: 20px;
  padding: 0 10px 0 0px;
}

.title-footer {
  display: flex;
  justify-content: left;
  align-items: center;
  font-size: 14px;
  height: 18px;
  overflow: hidden;
  font-style: italic;
  color: #d9d9d9;
}

.flex-item {
  margin-right: 5px;
}

.field {
  height: calc(100% - 2px);
  width: 100%;
  background: inherit;
}

.color-line {
  margin-top: -11px;
  height: 0px;
  width: 100%;
  border-bottom: 3px solid #ffffff;
}

.hidden {
  display: none;
}

.speaker {
  position: relative;
  box-sizing: border-box;
  margin: 0 2px 4px 2px;
  background: #e3e3e3;
  border: 1px solid #343e5c;
  border-radius: 3px;
  width: calc(100% - 4px);
  // transition: 0.2s;
  overflow: hidden;
  &:hover {
    cursor: pointer;
    box-shadow: 2px 2px 5px #343e5c;
  }
}

.speaker-title {
  width: 100%;
  padding-right: 5px;
  padding-left: 3px;
  box-sizing: border-box;
  color: #ffffff;
  height: 44px;
  display: flex;
  justify-content: left;
  align-items: center;
  border-bottom: 1px solid rgba(245, 246, 248, 0.5);
}

.empty-space {
  height: 0px;
}

.icon-speaker-edit {
  position: absolute;
  top: 2px;
  right: 3px;
  width: 24px;
  height: 24px;
  stroke: #ffffff;
  cursor: pointer;
  transition: 0.3s;
}

.icon-speaker-edit:hover {
  stroke: #ffffff;
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

.speaker-title:hover > .icon-move {
  visibility: visible;
}

.speaker-title:hover > .title-hidden {
  opacity: 0.2;
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
  top: 8px;
  left: 46%;
  visibility: hidden;
  width: 28px;
  height: 28px;
  fill: #ffffff;
  stroke: #343e5c;
  animation-name: ripple;
  animation-duration: 1s;
  animation-iteration-count: infinite;
  padding-right: 8px;
}
</style>
