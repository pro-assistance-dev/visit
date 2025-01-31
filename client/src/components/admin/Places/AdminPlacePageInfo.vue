<template>
  <el-form>
    <LeftRightContainer left-width="280px">
      <template #right>
        <el-form-item label="Название" prop="surname">
          <el-input v-model="place.name" placeholder="Введите название" formatter="firstLetterUpper" @blur="update" />
        </el-form-item>
        <el-form-item label="Описание" prop="description">
          <el-input
            type="textarea"
            v-model="place.description"
            placeholder="Введите описание"
            formatter="firstLetterUpper"
            @blur="update"
          />
        </el-form-item>
      </template>
    </LeftRightContainer>
  </el-form>
</template>

<script lang="ts" setup>
import Place from '@/classes/Place';

const place: Place = PlacesStore.Item;

const update = async () => {
  PlacesStore.Update();
};
</script>

<style lang="scss" scoped>
.light-title {
  color: #a1a8bd;
  display: flex;
  align-items: center;
}

.upper {
  text-transform: uppercase;
}

.flex-center {
  display: flex;
  align-items: center;
}

.contact-icon {
  font-size: 25px;
  margin-right: 5px;
}

.el-tag {
  font-size: 14px;
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
  color: #343d5c;
  margin-top: 28px;
  margin-left: 10px;
  text-transform: uppercase;
}

.left-info {
  font-size: 12px;
  color: #343d5c;
  margin-top: 3px;
  margin-left: 10px;
}
</style>
