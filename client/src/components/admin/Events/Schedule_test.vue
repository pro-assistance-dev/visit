<template>
  <AdminListWrapper pagination show-header>
    <AdminEventsListFilters />
    <div class="global-wraper">
      <div class="title">
        <div class="time-line_title">
          <Button button-class="plus-button" icon="plus" icon-class="icon-plus" />
        </div>
        <div class="column1_title">Зал "Театр"</div>
        <div class="column2_title">Зал "Москва"</div>
        <div class="column3_title">Зал “Лена”+“Яна”</div>
        <div class="column4_title">Зал “Нева”+“Дон”</div>
      </div>
      <div class="wraper">
        <div class="time-line">
          <Hour :hour="9" />
          <Hour :hour="10" />
          <Hour :hour="11" />
          <Hour :hour="12" />
          <Hour :hour="13" />
          <Hour :hour="14" />
          <Hour :hour="15" />
          <Hour :hour="16" />
          <Hour :hour="17" />
          <Hour :hour="18" />
        </div>
        <div class="column1">
          <div class="card-event"></div>
        </div>
        <div class="column2"></div>
        <div class="column3"></div>
        <div class="column4"></div>
      </div>
    </div>
  </AdminListWrapper>
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import EventsSortsLib from '@/libs/sorts/EventsSortsLib';

const load = async () => {
  FTSP.Get().setS(EventsSortsLib.byName());
  await Store.Events.FTSP({ ftsp: FTSP.Get() });
};

Hooks.onBeforeMount(load, {
  adminHeader: {
    title: 'Мероприятия',
    buttons: [{ text: 'Добавить', type: 'primary', action: () => PF.H.Router.To('/admin/events/new') }],
  },
  pagination: { storeModule: 'events', action: 'ftsp' },
});

const edit = async (id?: string) => {
  if (!id) {
    return;
  }
  await PF.H.Router.To(`/admin/events/${id}`);
};

const remove = async (id: string) => {
  await Store.Events.Remove(id);
};
const update = async (event: Event) => {
  await Store.Events.Update(event);
};
</script>

<style lang="scss" scoped>
.global-wraper {
  position: relative;
  overflow-y: scroll;
}
.wraper {
  display: grid;
  grid-template-columns: 100px 1fr 1fr 1fr 1fr;
  height: 65vh;
  width: 100%;
}

.title {
  position: sticky;
  top: 0;
  left: 0;
  z-index: 1;
  display: grid;
  grid-template-columns: 100px 1fr 1fr 1fr 1fr;
  height: 60px;
  width: 100%;
  font-weight: bold;
  font-size: 16px;
}

.time-line_title {
  background: #ffffff;
  color: #343e5c;
  border-right: 1px solid #959595;
  border-bottom: 2px solid #343e5c;
  display: flex;
  justify-content: center;
  align-items: center;
}

.column1_title {
  background: #1979cf;
  color: #ffffff;
  border-right: 1px solid #959595;
  border-bottom: 2px solid #343e5c;
  display: flex;
  justify-content: center;
  align-items: center;
}
.column2_title {
  background: #f3911c;
  color: #ffffff;
  border-right: 1px solid #959595;
  border-bottom: 2px solid #343e5c;
  display: flex;
  justify-content: center;
  align-items: center;
}
.column3_title {
  background: #660b19;
  color: #ffffff;
  border-right: 1px solid #959595;
  border-bottom: 2px solid #343e5c;
  display: flex;
  justify-content: center;
  align-items: center;
}
.column4_title {
  background: #0aa249;
  color: #ffffff;
  border-right: 1px solid #959595;
  border-bottom: 2px solid #343e5c;
  display: flex;
  justify-content: center;
  align-items: center;
}
.time-line {
  background: #ffffff;
  height: 100%;
}

.column1 {
  background: #dff2f8;
  height: 100%;
  border-right: 1px solid #959595;
}

.column2 {
  background: #fff0df;
  height: 100%;
  border-right: 1px solid #959595;
}

.column3 {
  background: #ffd2d9;
  height: 100%;
  border-right: 1px solid #959595;
}

.column4 {
  background: #eefff5;
  height: 100%;
  border-right: 1px solid #959595;
}

.card-event {
  background: #ffffff;
  border: 1px solid #959595;
  border-radius: 5px;
  height: calc(300px - 2px);
  margin: 1px 2px;
  box-sizing: border-box;
}

:deep(.icon-plus) {
  width: 24px;
  height: 24px;
  cursor: pointer;
  fill: #00bea5;
}

.plus-button {
  width: 100%;
  border-radius: 0;
  color: #00bea5;
  background: #c1efeb;
  height: 100%;
  margin: 0;
}
</style>
