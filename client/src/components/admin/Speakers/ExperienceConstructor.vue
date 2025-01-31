<template>
  <el-form>
    <el-form-item label="Заведение" prop="surname">
      <el-input
        v-model="experience.place"
        @input="experience.setPlace"
        @blur="update"
        placeholder="Введите заведение"
        formatter="firstLetterUpper"
      ></el-input>
    </el-form-item>
    <el-form-item label="Подразделение" prop="name">
      <el-input v-model="experience.division" @blur="update" placeholder="Введите подразделение" formatter="firstLetterUpper"></el-input>
    </el-form-item>
    <el-form-item label="Должность" prop="patronymic">
      <el-input v-model="experience.position" @blur="update" placeholder="Введите должность" formatter="firstLetterUpper"></el-input>
    </el-form-item>
  </el-form>
</template>

<script setup lang="ts">
import Experience from '@/classes/Experience';
const experience: Ref<Experience> = ref(Experience.Create());
const emit = defineEmits(['add']);

const props = defineProps({
  experience: {
    type: Experience as PropType<Experience>,
    required: true,
  },
});

const update = async () => {
  await ExperiencesStore.Update(props.experience);
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
