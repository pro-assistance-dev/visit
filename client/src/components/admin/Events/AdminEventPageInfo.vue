<template>
  <LeftRightContainer left-width="280px">
    <template #left>
      <UploaderImage
        :file-info="event.image"
        :height="280"
        :default-ratio="1"
        :emit-crop="true"
        @ratio="(e) => (element.ratio = e)"
        @crop="savePhoto"
      />
      <div class="left-title">
        Статус:
        <PString string="Ожидается" color="#006BB4" margin="0 10px" />
      </div>
    </template>
    <template #right>
      <div class="container">
        <PInput v-model="event.name" placeholder="Введите название" formatter="firstLetterUpper" @blur="update" />
        <!-- <el-select v-model="event.type" placeholder="Выбирите тип мероприятия" @change="update" -->
        <!--   ><el-option v-for="item in eventTypes" :key="item.value" :label="item.label" :value="item.value">{{ -->
        <!--     item.label -->
        <!--   }}</el-option></el-select -->
        <!-- > -->
        <PTextArea type="textarea" v-model="event.description" placeholder="Введите адрес" formatter="firstLetterUpper" @blur="update" />
        <DateInputRange v-model:start="startDate" v-model:end="endDate" @setStart="setStart" @setEnd="setEnd" />
      </div>
      <div class="place-item" :style="{ marginTop: '30px' }">
        <div class="bottom-buttons">
          <div class="title">Залы</div>
          <div class="button-add" @click="openPlacesList(eventDay)">+ Добавить</div>
        </div>

        <div v-for="(place, i) in event.places" :key="place.id" class="place-item-item">
          {{ place.name }}
          <div class="button-del" @click="removePlace(place.id)">Удалить</div>
          <div class="list-number">
            {{ i + 1 }}
          </div>
        </div>
      </div>
    </template>
  </LeftRightContainer>
  <PModal :show="openedPlacesList" @close="openedPlacesList = false" class="session-edit">
    <AdminPlacesSelectList @select="addPlace" />
  </PModal>
  <Alarm />
</template>

<script lang="ts" setup>
import Alarm from '@/assets/svg/Alarm.svg';
import { EventNewDatesResult } from '@/classes/Event';
import AdminPlacesSelectList from './AdminPlacesSelectList.vue';

import Place from '@/classes/Place';

const eventTypes = [
  {
    value: 'Конференция',
    label: 'Конференция',
  },
  {
    value: 'Научно-практическая конференция',
    label: 'Научно-практическая конференция',
  },
  {
    value: 'Научно-практическая конференция',
    label: 'Научно-практическая конференция',
  },
  {
    value: 'Учебно-научная конференция',
    label: 'Учебно-научная конференция',
  },
  {
    value: 'Научно-методическая конференция',
    label: 'Научно-методическая конференция',
  },
  {
    value: 'Научно-методический семинар',
    label: 'Научно-методический семинар',
  },
  {
    value: 'Съезд',
    label: 'Съезд',
  },
  {
    value: 'Конгресс',
    label: 'Конгресс',
  },
];

const event = Store.Events.Item;
const openedPlacesList: Ref<boolean> = ref(false);
const openedDatesSelect: Ref<boolean> = ref(false);

const update = async () => {
  Provider.withHeadLoader(async () => {
    await Provider.store.dispatch('events/update');
  });
};

const savePhoto = async () => {
  await FileInfosStore.Create(event.image);
  event.imageId = event.image.id;
  await update();
};

const addPlace = async (item: Place) => {
  openedPlacesList.value = false;
  if (event.placeExists(item.id as string)) {
    PF.Notification.Warning('Данное место проведения уже добавлено');
    return;
  }
  // Provider.store.dispatch('m2m/create', event.bindPlace(item));
  addSchedule(item.id as string);
  for (let j = 0; j < event.eventDays.length; j++) {
    await addSchedule(item.id, event.eventDays[j].id);
  }
};

const startDate = ref(event.getStartDate());
const endDate = ref(event.getEndDate());

const setStart = async (date: Date) => {
  startDate.value = date;
  if (startDate.value > endDate.value) {
    endDate.value = startDate.value;
  }
  const comp = event.setDates(startDate.value, endDate.value);
  await updateEventComponents(comp);
};

const setEnd = async (date: Date) => {
  endDate.value = date;
  if (startDate.value > endDate.value) {
    startDate.value = endDate.value;
  }
  const comp = event.setDates(startDate.value, endDate.value);
  await updateEventComponents(comp);
};

const updateEventComponents = async (eventComponents: EventNewDatesResult): Promise<void> => {
  for (let index = 0; index < eventComponents.newSchedules.length; index++) {
    await SchedulesStore.Create(eventComponents.newSchedules[index]);
  }
  for (let index = 0; index < eventComponents.daysForDelete.length; index++) {
    await EventDaysStore.Remove(eventComponents.daysForDelete[index].id);
  }
  for (let index = 0; index < eventComponents.newEventDays.length; index++) {
    await EventDaysStore.Create(eventComponents.daysForDelete[index]);
  }
  for (let index = 0; index < eventComponents.schedulesForDelete.length; index++) {
    await EventDaysStore.Remove(eventComponents.schedulesForDelete[index].id);
  }
  for (let index = 0; index < eventComponents.perfomsForUpdate.length; index++) {
    await PerfomsStore.Update(eventComponents.perfomsForUpdate[index]);
  }
};

const addSchedule = async (placeId?: string, eventDayId?: string) => {
  const item = event.addSchedule(placeId, eventDayId);
  await SchedulesStore.Create(item);
};

const removePlace = async (id: string) => {
  const { perfomsForUpdate, schedulesForDelete } = event.unbindPerfomsFromPlace(id);
  for (let index = 0; index < schedulesForDelete.length; index++) {
    await SchedulesStore.Remove(schedulesForDelete[index].id);
  }
  for (let index = 0; index < perfomsForUpdate.length; index++) {
    await PerfomsStore.Update(perfomsForUpdate[index]);
  }
  // Provider.store.dispatch('m2m/remove', event.unbindPlace(id));
};

const openPlacesList = () => {
  openedPlacesList.value = true;
};
</script>

<style lang="scss" scoped>
.light-title {
  color: #a1a8bd;
  display: flex;
  align-items: center;
}

.upper {
  text-transform: uppercase;
}

.flex-center {
  display: flex;
  align-items: center;
}

.contact-icon {
  font-size: 25px;
  margin-right: 5px;
}

.el-tag {
  font-size: 14px;
}

.container {
  box-sizing: border-box;
  width: 100%;
  height: calc(100% - 100px);
}

.register-tag {
  &:hover {
    cursor: pointer;
    border-width: 2px;
  }
}

:deep(.el-upload--picture-card) {
  width: 150px;
  font-size: 50px;
}

:deep(.el-upload--picture-card i) {
  font-size: 50px;
  color: #00b5a4;
  padding: 0 114px;
}

.left-title {
  font-size: 14px;
  font-weight: bold;
  color: #343d5c;
  margin-top: 28px;
  margin-left: 10px;
  text-transform: uppercase;
  display: flex;
  justify-content: left;
  align-items: center;
}

.left-info {
  font-size: 12px;
  color: #343d5c;
  margin-top: 3px;
  margin-left: 10px;
}

.info-title {
  width: 100%;
  box-sizing: border-box;
  padding: 10px;
  border: 1px solid #b83616;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-radius: 5px;
  margin-top: 40px;
}

.el-select {
  width: 100%;
}

:deep(.el-date-editor.el-input, .el-date-editor.el-input__inner) {
  width: 100%;
}

:deep(.el-form-item) {
  display: block;
  margin-bottom: 16px;
}

:deep(.el-input__inner) {
  height: 40px;
  width: 100%;
  display: flex;
  font-family: Comfortaa, Arial, Helvetica, sans-serif;
  font-size: 15px;
  color: #343d5c;
}

:deep(.el-input__inner::placeholder) {
  color: #b0a4c0;
}

:deep(.el-input__icon) {
  color: #343d5c;
}

:deep(.el-form-item__label) {
  color: #b0a4c0;
  padding: 0 !important;
  text-transform: uppercase;
  margin-left: 5px;
  font-size: 14px;
  margin-bottom: 6px;
}

:deep(.el-input-number__increase) {
  border-radius: 0;
}

:deep(.el-input-number__decrease) {
  border-radius: 0;
}

:deep(.el-textarea__inner) {
  font-family: Comfortaa, Arial, Helvetica, sans-serif;
  font-size: 15px;
  color: #343d5c;
}

.button-add {
  border: none;
  background: inherit;
  color: #00b5a4;
  transition: 0.3s;
  cursor: pointer;
}

.button-add:hover {
  color: #009e8f;
}

.button-del {
  position: absolute;
  top: 13px;
  right: 36px;
  border: none;
  background: inherit;
  color: #a3a9be;
  transition: 0.3s;
  cursor: pointer;
}

.button-del:hover {
  color: #b63414;
}

.bottom-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-right: 15px;
}

.list-number {
  position: absolute;
  top: 10px;
  right: 10px;
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  background: #1979cf;
  border-radius: 20px;
}

.title {
  font-family: 'Open Sans', sans-serif;
  font-size: 14px;
  color: #c4c4c4;
  margin: 10px;
  color: #b0a4c0;
  text-transform: uppercase;
  margin-left: 5px;
  font-size: 14px;
}

.place-item {
  width: calc(100% - 22px);
  position: relative;
  padding: 0 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  margin: 0px;
  background: #dff2f8;
}

.place-item-item {
  position: relative;
  width: calc(100% - 42px);
  margin: 0px 10px 20px 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  padding: 12px 10px;
  background: #f9fafb;
}

.session-edit {
  position: absolute;
  top: 20px;
  right: 20px;
  border: 1px solid #343e5c;
  border-radius: 5px;
  width: 100%;
  background: #ffffff;
  padding: 10px;
  z-index: 5;
}

.icon-alarm {
  width: 40px;
  height: 40px;
  fill: #b83616;
  margin: 0 10px 0 0;
}
</style>
