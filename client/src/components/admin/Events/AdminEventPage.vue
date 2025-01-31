<template>
  <!-- <AdminEventPageInfo v-if="mounted" :event="event" /> -->
  <PTabs v-if="mounted" :tabs="tabs">
    <component :is="component" />
  </PTabs>
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import AdminEventPageInfo from '@/components/admin/Events/AdminEventPageInfo.vue';
const mounted = ref(false);
const event = Store.Events.Item;
const tabs = [
  {
    id: 0,
    name: 'Информация',
  },
];

const components = {
  '0': AdminEventPageInfo,
  // AdminEventPageSchedule: AdminEventPageSchedule,
  // AdminEventPageSponsors: AdminEventPageSponsors,
};
const handleSelect = async (item: ITab): Promise<void> => {
  emit('select', item);
  if (item.id) {
    choise.value = +item.id;
  }
};

const load = async () => {
  await Store.Events.Get();
  PF.AdminUI.Head.Set(PF.H.Router.Id() ? event.name : 'Добавить мероприятие', [
    Button.Warning('Добавить', remove),
    Button.Primary('К мероприятию', toEvent),
  ]);
  mounted.value = true;
  return;
};

const choise = ref(0);

const component = computed(() => {
  if (choise.value === 0) {
    return components['0'];
  }
  return 'no';
});

const toEvent = async () => {
  PF.H.Router.To(`/events/${event.slug}`);
};

const remove = async () => {
  await Store.Events.Remove(PF.H.Router.Id());
  await PF.H.Router.To('/admin/events');
};
Hooks.onBeforeMount(load);
Hooks.onBeforeRouteLeave();
</script>

<style scoped>
.field {
  width: 100%;
  height: 100%;
  background: red;
  padding: 10px;
}

.tab {
  font-size: 12px;
  text-transform: uppercase;
  color: #b0a4c0;
  border: 1px solid #999999;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  text-align: center;

  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */

  background: #f5f5f5;
  margin: -0.5px;
}

.tab:hover {
  background: #dff2f8;
}

.selected-tab {
  font-size: 12px;
  font-weight: bold;
  text-transform: uppercase;
  position: relative;
  border: 1px solid #343d5c;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  cursor: pointer;

  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */

  margin: -0.5px;
  z-index: 1;
}
</style>
