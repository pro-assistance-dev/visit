<template>
  <el-form>
    <InfoItem title="фамилия" icon="edit-title" :with-open-window="false" margin="0" padding="0" width="100%">
      <el-form-item style="width: 100%" prop="name">
        <el-input v-model="place.name" />
      </el-form-item>
    </InfoItem>
    <Button button-class="save-button" text="Создать" @click="create" />
  </el-form>
</template>

<script setup lang="ts">
import Place from '@/classes/Place';
const place: Ref<Place> = ref(Place.Create());
const emit = defineEmits(['add']);
const create = async (): Promise<void> => {
  await PlacesStore.Create(place.value);
  await PlacesStore.Get(place.value.id);
  PlacesStore.UnshiftToAll(place.value);
  emit('add');
};
</script>
<style lang="scss" scoped>
div {
  font-size: 14px;
}
h4 {
  margin-bottom: 0;
}
.header-center {
  text-align: center;
}
.filter {
  border-radius: 20px;
  width: 70%;
  margin-left: auto;
  margin-right: auto;
  margin-bottom: 40px;
}

.collapseHeader {
  padding-left: 10px;
  line-height: 15px;
}

.el-collapse-item {
  background-color: white;
  margin-bottom: 20px;
  padding: 4px;
  border-radius: 10px;
  width: 100%;
  box-sizing: border-box;
}
:deep(.el-collapse-item__wrap) {
  border-bottom: none;
}

:deep(.el-collapse-item__content) {
  padding: 20px;
}
:deep(.el-collapse-item__content),
:deep(.el-collapse-item__header) {
  font-size: 14px;
  padding-left: 20px;
}

:deep(.el-collapse-item__header) {
  height: 50px;
  font-size: 16px;
  font-weight: bold;
}
:deep(.el-collapse-item__header),
:deep(.el-collapse-item__wrap) {
  border: none;
}

.collapse-content-container {
  margin-left: 10px;
}

:deep(.el-form-item) {
  height: 38px;
  margin: 1px 0 0 10px;
  align-items: center;
}

:deep(.el-form-item__content) {
  align-items: center;
}

:deep(.el-input__inner) {
  height: 38px;
  margin: 0;
  padding: 0;
}

:deep(.el-input) {
  min-height: 38px;
}

:deep(.el-input__wrapper) {
  box-shadow: none;
  height: 36px;
}

:deep(.el-input__wrapper:hover) {
  box-shadow: none;
}

:deep(.el-input__wrapper.is-focus) {
  box-shadow: none;
}

.save-button {
  width: 300px;
  border-radius: 5px;
  height: 42px;
  color: #006bb4;
  background: #dff2f8;
  margin: 20px auto 0 auto;
  font-size: 14px;
}
</style>
