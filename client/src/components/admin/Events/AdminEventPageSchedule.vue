<template>
  <AdminScheduleTable />
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import Perfom from '@/classes/Perfom';
import ISearchElement from '@/interfaces/ISearchElement';

const event: Event = Store.Events.Item;
const perfom: Perfom = PerfomsStore.Item;
const openedPerfomConstructor: Ref<boolean> = ref(false);

const selectSearch = async (res: ISearchElement) => {
  await PerfomsStore.Get(res.value);
  addPerfom(perfom);
};
const addPerfom = async (item: Perfom) => {
  openedPerfomConstructor.value = false;
  // Provider.store.dispatch('m2m/create', event.bindPerfom(item));
};

const removePerfom = async (id: string) => {
  // Provider.store.dispatch('m2m/remove', event.unbindPerfom(id));
};

const selectedPerfom = ref(new Perfom());

const copyPerfom = (perfom: Perfom) => {
  selectedPerfom.value = perfom;
};
const bindPerformToSchedule = async () => {
  // const m2m = selectedEventDay.value.getScheduleByPlaceId(selectedP.value).bindPerfom(selectedPerfom.value, selectedT.value);
  // Provider.store.dispatch('m2m/create', m2m);
};
</script>
