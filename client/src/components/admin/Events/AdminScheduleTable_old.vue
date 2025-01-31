<template>
  <div>
    <Button @click="addEventDay" text="Добавить день" />
  </div>
  <div v-for="eventDay in event.eventDays" :key="eventDay.id">
    <div id="head" style="display: flex">
      <div>Время</div>
      <div v-for="place in eventDay.places" :key="place.name">
        {{ place.name }}
        <Button @click="removePlace(eventDay, place.id)" text="x" />
      </div>
      <Button @click="openPlacesList(eventDay)" text="Добавить место проведения" />
    </div>
    <div id="body" style="display: flex">
      <div id="КолонкаВремени">
        <div v-for="t in getTimeArray()" :key="t">
          <div style="display: flex">
            <div>
              {{ t }}
            </div>
            <div style="display: flex">
              <div v-for="place in eventDay.places" :key="place.id">
                <div class="v-resize-border" style="margin: 5px">
                  <VueDraggable v-model="eventDay.places" :group="{ name: 'perfoms', pull: true, put: () => putF(eventDay, t, place) }">
                    <div :class="getSlotClasses(eventDay.getScheduleByPlaceId(place.id), place.id, t)">[___________]</div>
                  </VueDraggable>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <PModal :show="openedPlacesList" @close="openedPlacesList = false">
    <AdminPlacesSelectList @select="addPlace" />
  </PModal>
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import Place from '@/classes/Place';
import EventDay from '@/classes/EventDay';
import { VueDraggable } from 'vue-draggable-plus';
import Schedule from '@/classes/Schedule';
const event = Store.Events.Item;
const openedPlacesList: Ref<boolean> = ref(false);
const selectedEventDay: Ref<EventDay> = ref(new EventDay());

const openPlacesList = (eventDay: EventDay) => {
  selectedEventDay.value = eventDay;
  openedPlacesList.value = true;
};

const getSlotClasses = (schedule: Schedule, placeId: string, t: string) => {
  return {
    'filled-slot': schedule.getPerfomByStartTime(t),
  };
};
const addPlace = async (item: Place) => {
  openedPlacesList.value = false;
  if (selectedEventDay.value.placeExists(item.id as string)) {
    ElMessage.warning('Данное место проведения уже добавлено');
    return;
  }
  // Provider.store.dispatch('m2m/create', selectedEventDay.value.bindPlace(item));
  addSchedule(item.id as string);
};

const addSchedule = async (placeId: string) => {
  const item = selectedEventDay.value.addSchedule(placeId);

  SchedulesStore.Create(item);
};

const removePlace = async (eventDay: EventDay, id: string) => {
  // Provider.store.dispatch('m2m/remove', eventDay.unbindPlace(id));
};

const addEventDay = async () => {
  await EventDaysStore.Create(event.addEventDay());
};

const selectedTime = ref('');
const selectedPlaceId = ref('');

const putF = async (eventDay: EventDay, timeStart: string, place: Place) => {
  selectedTime.value = timeStart as string;
  selectedPlaceId.value = place.id as string;
  selectedEventDay.value = eventDay;
};
</script>
<style lang="scss">
.filled-slot {
  background: red;
  resize: vertical;
  cursor: pointer;
}
.v-resize-border {
  border-top: 9px black;
  border-bottom: 9px black;
  cursor: ns-resize;
  margin: 10px;
}
</style>
