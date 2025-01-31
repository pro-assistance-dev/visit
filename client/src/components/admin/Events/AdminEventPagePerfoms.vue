<template>
  <h3>Доклады</h3>
  <Button text="Добавить доклад" @click="openedPerfomConstructor = true" />
  <RemoteSearch :key-value="'perfom'" placeholder="Найдите доклад для добавления" max-width="100%" @select="selectSearch" />
  <VueDraggable
    ref="el"
    v-model="event.perfoms"
    :sort="false"
    :group="{ name: 'perfoms', put: false }"
    :clone="copyPerfom"
    @end="bindPerformToSchedule"
  >
    <div v-for="perfom in event.perfoms" :key="perfom.id">
      <div style="display: flex">
        {{ perfom.name }}
        <Button @click="removePerfom(perfom.id)" text="x" />
      </div>
    </div>
  </VueDraggable>
  <PModal :show="openedPerfomConstructor" @close="openedPerfomConstructor = false">
    <PerfomConstructor @create="addPerfom" />
  </PModal>
</template>

<script setup lang="ts">
import Event from '@/classes/Event';
import Perfom from '@/classes/Perfom';
import { VueDraggable } from 'vue-draggable-plus';
import ISearchElement from '@/interfaces/ISearchElement';
const event = Store.Events.Item;
const perfom = PerfomsStore.Item;
const openedPerfomConstructor: Ref<boolean> = ref(false);

import RemoteSearch from '@/services/components/RemoteSearch.vue';
const selectSearch = async (res: ISearchElement) => {
  await PerfomsStore.Get(res.value);
  addPerfom(perfom);
};

const addPerfom = async (item: Perfom) => {
  openedPerfomConstructor.value = false;
  // Provider.store.dispatch('m2m/create', event.value.bindPerfom(item));
};

const removePerfom = async (id: string) => {
  // Provider.store.dispatch('m2m/remove', event.value.unbindPerfom(id));
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
