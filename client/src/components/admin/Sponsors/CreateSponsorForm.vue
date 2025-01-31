<template>
  <el-form>
    <InfoItem
      title="Название"
      icon="edit-title"
      :with-open-window="false"
      :with-hover="false"
      border-color="#ffffff"
      base-box-margin="0 0 15px 0"
      padding="0"
      width="100%"
    >
      <el-form-item style="width: 100%" prop="name">
        <el-input v-model="sponsor.name" />
      </el-form-item>
    </InfoItem>
    <Button button-class="save-button" text="Создать" @click="create" />
  </el-form>
</template>

<script setup lang="ts">
import Sponsor from '@/classes/Sponsor';
const sponsor: Ref<Sponsor> = ref(Sponsor.Create());
const emit = defineEmits(['add']);

const create = async (): Promise<void> => {
  await SponsorsStore.Create(sponsor.value);
  await SponsorsStore.Get(sponsor.value.id);
  SponsorsStore.UnshiftToAll(sponsor.value);
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

.save-button {
  width: 300px;
  border-radius: 5px;
  height: 42px;
  color: #006bb4;
  background: #dff2f8;
  margin: 0 auto;
  font-size: 14px;
}
</style>
