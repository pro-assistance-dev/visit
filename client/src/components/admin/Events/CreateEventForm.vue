<template>
    <div class="container">
        <PInput label="Название"  v-model="event.name" placeholder="Введите название" formatter="firstLetterUpper" />
        <DateInputRange label="Даты проведения" v-model:start="startDate" v-model:end="endDate" @setStart="setStart" @setEnd="setEnd" />
    </div>
    <Button button-class="save-button" text="Создать" @click="create" />
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import EventDay from '@/classes/EventDay';
const emit = defineEmits(['add']);

const create = async (): Promise<void> => {
  if (event.name === '') {
    PF.Message.Error('Введите название')
    return
  }
  await Store.Events.Create();
  for (let index = 0; index < event.eventDays.length; index++) {
    await EventDaysStore.Create(event.eventDays[index]);
  }
  Store.Events.UnshiftToAll(event);
  emit('add');
};
const event: Event = Store.Events.Item;
const startDate = ref(event.getStartDate());
const endDate = ref(event.getEndDate());

onBeforeMount(() => {
  event.init()
})


const setStart = async (date: Date) => {
  startDate.value = date;
  if (startDate.value > endDate.value) {
    endDate.value = startDate.value;
  }
  event.setDates(startDate.value, endDate.value);
  setPeriod()
};

const setEnd = async (date: Date) => {
  endDate.value = date;
  if (startDate.value > endDate.value) {
    startDate.value = endDate.value;
  }
  setPeriod()
};
const setPeriod =() => {
  event.eventDays = []
  const newDates = Dates.GetDatesFromPeriod(startDate.value, endDate.value);
  newDates.forEach((d:Date) => {
    const newEventDay = EventDay.Create(event.id as string, d);
    event.eventDays.push(newEventDay)
  });
}
onBeforeUnmount(() => {
Store.Events.Set()
})
</script>

<style  scoped>
.container {
  width: 100%;
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
  margin-top: 28px;
  margin-left: 10px;
  text-transform: uppercase;
}

.left-info {
  font-size: 12px;
  margin-top: 3px;
  margin-left: 10px;
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
  }


  :deep(.el-form-item__label) {
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
    color: #B63414;
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
    margin: 0px ;
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
</style
