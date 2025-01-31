<template>
  <div
    v-if="opened"
    v-click-outside.prevent="outsideClick"
    class="add-window"
    :style="{
      left: left,
      top: top,
    }"
  >
    <Button text="Добавить сессию" width="calc(100% - 10px)" height="32px" @click="addSession" />
    <Button text="Добавить выступление" width="calc(100% - 10px)" height="32px" @click="addPerfom" />
  </div>
</template>

<script setup lang="ts">
import Schedule from '@/classes/Schedule';

const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);
const opened = ref(false);
const x = ref(0);
const y = ref(0);

const left = computed(() => x.value + 'px');
const top = computed(() => y.value + 'px');

const event: Event = Store.Events.Item;
const outsideClick = () => {
  if (opened.value) {
    opened.value = false;
  }
};
const open = (e: any) => {
  x.value = e.clientX;
  y.value = e.clientY;

  opened.value = true;
  selectObjects(e);
};

onBeforeMount(() => {
  window.addEventListener('dblclick', open);
});

let startTime = '';
let scheduleId = '';
let schedule: Schedule | undefined;

const selectObjects = (e: any) => {
  const elements = document.elementsFromPoint(e.x, e.y);
  elements.forEach((e: HTMLElement) => {
    if (e.className === 'place') {
      scheduleId = e.id;
    }
    if (e.className === 'five-min') {
      startTime = e.innerText;
    }
  });
};

const addPerfom = async () => {
  schedule = event.schedules.find((s: Schedule) => s.id === scheduleId);
  if (!schedule) {
    return;
  }
  const item = schedule.addPerfom();
  item.startTime = startTime;

  const nextPerfom = schedule.findNextPerfom(item);
  if (nextPerfom) {
    item.endTime = nextPerfom.startTime;
  } else {
    item.endTime = Time.AddMtoHM(item.startTime, 20);
  }
  scheduler.value.sortByStartTime(schedule.perfoms);
  await PerfomsStore.Create(item);
  opened.value = false;
};

const addSession = async () => {
  schedule = event.schedules.find((s: Schedule) => s.id === scheduleId);
  if (!schedule) {
    return;
  }
  const item = schedule.addSession();
  item.startTime = startTime;
  item.endTime = Time.AddMtoHM(item.startTime, 20);
  scheduler.value.sortByStartTime(schedule.sessions);
  await SessionsStore.Create(item);

  opened.value = false;
};
</script>
<style lang="scss">
.add-window {
  position: fixed;
  border: 1px solid #343e5c;
  border-radius: 5px;
  background: #ffffff;
  z-index: 10;
  width: 300px;
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}
</style>
