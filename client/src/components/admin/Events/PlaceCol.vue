<template>
  <div :id="schedule.id" class="place" :style="placeStyle">
    <PlaceColHead :schedule="schedule" :place="place" @addSession="addSession" />
    <div class="field">
      <div v-if="opts.perfomsShowed" class="perfom-field">
        <PerfomCard
          v-for="(perfom, i) in schedule.perfoms"
          :key="perfom.id"
          :perfom="perfom"
          perfom-field
          :prev-perfom="schedule.perfoms[i - 1]"
          :schedule="schedule"
          :external-height="emptyHeight(perfom, schedule.findPrevPerfom(perfom))"
          @edit="editPerfom(perfom)"
          @remove="removePerfom(perfom)"
        />
      </div>
      <div class="session-field">
        <SessionCard
          v-for="(session, i) in schedule.sessions"
          :key="session.id"
          :fix-head="showOnlySessions"
          :session="session"
          :prev-session="schedule.sessions[i - 1]"
          :schedule="schedule"
          :external-height="emptyHeight(session, schedule.findPrevSession(session))"
          @edit="editSession(session)"
        />
      </div>
    </div>
    <PModal v-if="openPerfomConstructor" :show="openPerfomConstructor" @close="openPerfomConstructor = false">
      <PerfomConstructor :perfom="selectedPerfom" @remove="removePerfom" @create="addPerfom" />
    </PModal>
    <PModal v-if="showEditSessionModal" :show="showEditSessionModal" title="Редактировать сессию" @close="showEditSessionModal = false">
      <SessionConstructor :start-time="selectedSession" :session="selectedSession" @close="showEditSessionModal = false" />
      <Button button-class="del-button" text="Удалить" @click="removeSession" />
    </PModal>
  </div>
</template>

<script setup lang="ts">
import Perfom from '@/classes/Perfom';
import Place from '@/classes/Place';
import Schedule from '@/classes/Schedule';
import ScheduleDragger from '@/classes/ScheduleDragger';
import ScheduleOptions from '@/classes/ScheduleOptions';
import Session from '@/classes/Session';

const props = defineProps({
  schedule: {
    type: Object as PropType<Schedule>,
    required: true,
  },
  place: {
    type: Object as PropType<Place>,
    required: true,
  },
  showOnlySessions: {
    type: Boolean as PropType<boolean>,
    required: true,
  },
});
const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);

const event = Store.Events.Item;
const dragger: ComputedRef<ScheduleDragger> = computed(() => Provider.store.getters['schedules/dragger']);
const opts: ComputedRef<ScheduleOptions> = computed(() => Provider.store.getters['schedules/options']);
const showEditSessionModal = ref(false);

// const removeSession = async (sessionId?: string) => {
//   props.schedule.removeSession(sessionId);
//   Provider.store.dispatch('sessions/remove', sessionId);
// };

const selectedPerfom: Ref<Perfom | undefined> = ref(undefined);
const selectedSession: Ref<Session | undefined> = ref(undefined);
const openPerfomConstructor: Ref<boolean> = ref(false);

const editPerfom = (perfom: Perfom) => {
  selectedPerfom.value = perfom;
  openPerfomConstructor.value = true;
};
const editSession = (session: Session) => {
  selectedSession.value = session;
  showEditSessionModal.value = true;
};

const placeStyle = computed(() => {
  return {
    height: scheduler.value.getScheduleHeight() + 45 + 'px',
  };
});

const addSession = async () => {
  const item = props.schedule.addSession();
  await SessionsStore.Create(item);
};
const openWindow: Ref<boolean> = ref(false);

const addPerfom = async () => {
  selectedPerfom.value = props.schedule.addPerfom();
  await PerfomsStore.Create(selectedPerfom.value);
  openPerfomConstructor.value = true;
  openWindow.value = false;
};

const removePerfom = async (perfom?: Perfom) => {
  const p = perfom ?? selectedPerfom.value;
  if (!p) {
    return;
  }
  props.schedule.removePerfom(p.id as string);
  p.scheduleId = undefined;
  p.eventId = event.id;
  event.perfoms.push(p);
  openPerfomConstructor.value = false;
  await PerfomsStore.Remove(p.id);
};

const removeSession = async () => {
  if (!selectedSession.value) {
    return;
  }
  props.schedule.removeSession(selectedSession.value.id as string);
  selectedSession.value.scheduleId = undefined;
  // selectedSession.value.eventId = event.value.id;
  event.sessions.push(selectedSession.value);
  showEditSessionModal.value = false;
  await SessionsStore.Remove(selectedSession.value.id);
};

const emptyHeight = (perfom: Perfom | Session, prevPerfom?: Perfom) => {
  const startMinutes = Time.HMtoM(perfom.startTime);
  const prevMinutes = prevPerfom && prevPerfom.show ? Time.HMtoM(prevPerfom.endTime) : Time.HMtoM(scheduler.value.startHM);
  const hh = ((startMinutes - prevMinutes) / scheduler.value.timelineStep) * scheduler.value.timelinePxHeight;

  return {
    height: computed(() => hh + 'px').value,
  };
};
</script>

<style lang="scss" scoped>
.cover {
  background: grey;
  position: absolute;
  width: 400px;
  height: 400px;
  z-index: 1000;
}

.i-c {
  display: flex;
  align-content: justify-content;
  justify-content: space-between;
  width: 100%;
  min-width: 400px;
}

.hidden {
  display: none;
}

.border {
  width: 10px;
  opacity: 0.2;
  background: black;
  z-index: 1000;
}

.border-vert {
  height: 10px;
  opacity: 0.2;
  background: black;
  z-index: 1000;
}

.border-bottom {
  height: 10px;
  opacity: 0.2;
  background: black;
  z-index: 1000;
  margin-bottom: -20px;
}

.del-button {
  width: 100%;
  border-radius: 5px;
  height: 42px;
  color: #ff4c68;
  background: #ecc7c7;
  margin: 10px 10px 0 0;
  font-size: 14px;
}

.place {
  width: 100%;
  max-width: 400px;
  background: #eefff5;
  border: 1px solid #343e5c;
  margin: 0px 5px;
  box-sizing: border-box;
  border-radius: 5px;
  overflow: hidden;
}

.field {
  position: relative;
  width: 100%;
}

.session-field {
  position: absolute;
  box-sizing: border-box;
  overflow: hidden;
  width: 100%;
  top: 0;
  left: 0;
  z-index: 0;
}

.perfom-field {
  position: absolute;
  box-sizing: border-box;
  overflow: hidden;
  width: 80%;
  top: 0;
  left: 0;
  z-index: 1;
}

.place-title {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 42px;
  background: #c9ffdf;
  font-size: 16px;
  font-weight: bold;
  color: #343e5c;
  padding: 0 10px;
}

.icon-plus {
  width: 28px;
  height: 28px;
  fill: #00b5a4;
  cursor: pointer;
}

.icon-plus:hover {
  transform: scale(1.2);
  transition: 0.3s;
}

.add-window {
  position: absolute;
  top: 10px;
  right: 50px;
  border: 1px solid #343e5c;
  border-radius: 5px;
  background: #ffffff;
  z-index: 4;
}
</style>
