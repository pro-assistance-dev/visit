<template>
  <PString string="Выберете место проведения" margin="0 0 30px 0" font-size="16px" color="#006BB4" />
  <div>
    <GridContainer grid-gap="5px" margin="10px 0" grid-template-columns="repeat(auto-fit, minmax(280px, 1fr))">
      <PButton v-for="item in places" :key="item.id" button-class="change-button" :text="item.name" @click="selectPlace(item)" />
    </GridContainer>
  </div>
</template>

<script setup lang="ts">
import Place from '@/classes/Place';

const places: Place[] = PlacesStore.Items();

const mounted = ref(false);

onBeforeMount(async () => {
  // const filterQuery = new FilterQuery();
  // filterQuery.pagination.limit = 0;
  // filterQuery.setSortModel(placesSortsLib.byNameINN());
  await PlacesStore.GetAll();

  mounted.value = true;
});

const emit = defineEmits(['select']);

const selectPlace = async (item: Place): Promise<void> => {
  emit('select', item);
};

// const selectSearch = async (result: ISearchObject) => {
//   await Provider.store.dispatch('places/get', result.value);
//   stepper.value.actStepFunc(place.value as never);
// };
</script>

<style lang="scss" scoped>
.patient-name {
  color: #006bb4;
  font-size: 14px;
  width: 100%;
  padding: 0;
}

.medicine {
  color: #343d5c;
  font-size: 17px;
  width: 100%;
  padding: 0;
}

.hidden {
  display: none;
}

.scroll-block {
  height: 75vh;
  overflow: hidden;
  overflow-y: scroll;
  margin-left: 8px;
}

.registers-tooltip {
  &:hover {
    cursor: pointer;
  }
}

.patient-link {
  &:hover {
    cursor: pointer;
    text-decoration: underline;
  }
}

.tag-link:hover {
  color: #e4e4e4;
  cursor: pointer;
}

.flex-block {
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.icon-edit {
  width: 28px;
  height: 28px;
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

.line-item-left {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: auto;
  padding: 0;
}

.line-item-right {
  display: flex;
  justify-content: right;
  align-items: center;
  width: auto;
  padding: 0;
}

.icon-plus {
  width: 24px;
  height: 24px;
  cursor: pointer;
}

.icon-del {
  width: 10px;
  height: 10px;
  cursor: pointer;
}

:deep(.edit-icon) {
  width: 28px;
  height: 28px;
  color: #006bb4;
}

.filter-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.filter-block {
  position: relative;
  display: flex;
  z-index: 3;
  justify-content: space-between;
  align-items: flex-end;
  width: calc(100% - 20px);
  padding: 10px 10px 24px 10px;
  background: #f5f5f5;
  height: auto;
  border-bottom: 1px solid #c4c4c4;
}

.tools-block {
  display: flex;
  justify-content: right;
  align-items: center;
  margin-left: 10px;
}

.patient-count {
  margin-top: 10px;
  color: #b0a4c0;
  font-size: 14px;
}

:deep(.el-form-item) {
  margin: 8px 0 0 0;
}

:deep(.edit-button) {
  width: 40px;
  height: 40px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
  margin: 0 10px 0 0;
}

:deep(.commission-button) {
  min-width: 40px;
  max-width: 40px;
  height: 40px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
  margin: 0 10px 0 0;
}

:deep(.change-button) {
  width: 100%;
  height: 42px;
  border-radius: 5px;
  background: #ffffff;
  color: #343e5c;
  font-size: 14px;
}

:deep(.medical-commission-button) {
  width: 100%;
  height: 40px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
}

@media (max-width: 1915px) {
  .flex-block {
    width: 100%;
    display: block;
    justify-content: space-between;
    align-items: center;
  }
  .line-item {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    padding: 0;
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

  .line-item-right {
    margin: 10px 0;
    justify-content: space-between;
  }

  .item-flex:first-child {
    display: block;
    width: 100%;
    margin: 0 0px 10px 0;
  }
}
</style>
