<template>
  <el-form>
    <LeftRightContainer left-width="280px">
      <template #right>
        <el-form-item label="Название" prop="surname">
          <el-input v-model="sponsor.name" placeholder="Введите название" formatter="firstLetterUpper" @blur="update" />
        </el-form-item>
        <el-form-item label="Описание" prop="description">
          <el-input
            type="textarea"
            v-model="sponsor.description"
            placeholder="Введите описание"
            formatter="firstLetterUpper"
            @blur="update"
          />
        </el-form-item>
        <div class="tools-buttons">
          <button class="admin-add" @click.prevent="addBanner()">+ Добавить</button>
        </div>
        <AdminGallery @add-image="upload" :default-ratio="4 / 3" :file-list="sponsor.banners" @remove="removeBanner" />
      </template>
    </LeftRightContainer>
  </el-form>
</template>

<script lang="ts" setup>
import Sponsor from '@/classes/Sponsor';
import Banner from '@/classes/Banner';

const sponsor: Sponsor = SponsorsStore.Item;

const update = async () => {
  await SponsorsStore.Update(sponsor);
};
const addBanner = async () => {
  const item = sponsor.addBanner();
  await BannersStore.Create(item);
};
const removeBanner = async (item: Banner) => {
  await BannersStore.Remove(item.id);
};
const upload = async (item: Banner) => {
  FileInfosStore.Create(item.fileInfo);
  item.fileInfoId = item.fileInfo.id;
  await BannersStore.Update(item);
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
