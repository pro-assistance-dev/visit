<template>
  <AdminListWrapper v-if="mounted" pagination show-header :store="PlacesStore">
    <AdminPlacesListFilters />
    <div class="scroll-block">
      <!-- <div class="patient-count">Количество пациентов: {{ count }}</div> -->
      <div v-for="place in places" :key="place.id">
        <CollapseItem :is-collaps="false" padding="0 8px">
          <template #inside-title>
            <div class="flex-block" @click.prplace="() => undefined">
              <div class="item-flex">
                <div class="line-item-left">
                  <!-- <Button button-class="edit-button" color="#006bb4" icon="edit" icon-class="edit-icon" @click="edit(place.id)" /> -->
                  <SingleNameToggleForm :with-name="place" @submit="update(place)" />
                  <Button button-class="edit-button" color="#006bb4" icon="del" icon-class="edit-icon" @click="remove(place.id)" />
                </div>
              </div>
            </div>
          </template>
        </CollapseItem>
      </div>
    </div>
  </AdminListWrapper>
  <PModal :show="showAddModal" title="Добавить локацию" @close="showAddModal = false">
    <CreatePlaceForm @add="showAddModal = false" />
  </PModal>
</template>

<script setup lang="ts">
import Place from '@/classes/Place';

const showAddModal = ref(false);

const mounted = ref(false);
const places: Place[] = PlacesStore.Items();

const load = async () => {
  await PlacesStore.FTSP();
  mounted.value = true;
  PF.AdminUI.Head.Set('Локации', [Button.Primary('Добавить', () => (showAddModal.value = true))]);
};

Hooks.onBeforeMount(load);

const edit = async (id?: string) => {
  if (!id) {
    return;
  }
  await PF.H.Router.To(`/admin/places/${id}`);
};

const remove = async (id: string) => {
  await PlacesStore.Remove(id);
};
const update = async (place: Place) => {
  await PlacesStore.Update(place);
};
</script>

<style lang="scss" scoped>
:deep(.edit-icon) {
  width: 28px;
  height: 28px;
  color: #006bb4;
}
.scroll-block {
  height: calc(100% - 120px);
  overflow: hidden;
  overflow-y: scroll;
  margin-left: 8px;
}

.flex-block {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.item-flex {
  display: flex;
  width: 100%;
  justify-content: space-between;
  align-items: center;
  margin: 0 10px 0 0;
}

.item-flex:last-child {
  margin: 10px 0;
}
:deep(.edit-button) {
  min-width: 40px;
  max-width: 40px;
  height: 40px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
  margin: 0px;
  margin-right: 10px;
}
.line-item-left {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  margin-right: 10px;
  padding: 0;
}

.line-item-right {
  display: flex;
  justify-content: right;
  align-items: center;
  width: 100%;
  min-width: 210px;
  padding: 0;
}

@media (max-width: 1915px) {
  .flex-block {
    width: 100%;
    display: block;
    justify-content: space-between;
    align-items: center;
  }
  .item-flex {
    display: flex;
    width: 100%;
    justify-content: space-between;
    align-items: center;
    margin-top: 10px;
  }

  .item-flex:last-child {
    margin-top: 10px;
    margin-bottom: 10px;
  }
}

@media (max-width: 560px) {
  .line-item-left {
    margin: 10px 0;
  }

  .item-flex:first-child {
    display: block;
    width: 100%;
    margin: 0 0 10px 0;
  }
}
</style>
