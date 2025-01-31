<template>
  <div class="sponsor-item">
    <div class="bottom-buttons">
      <div class="title">Спонсоры</div>
      <div class="button-add" @click="openSponsorsList()">+ Добавить</div>
    </div>

    <div v-for="(sponsor, i) in event.sponsors" :key="sponsor.id" class="sponsor-item-item">
      {{ sponsor.name }}
      <div class="button-del" @click="removeSponsor(sponsor.id)">Удалить</div>
      <div class="list-number">
        {{ i + 1 }}
      </div>
    </div>
  </div>
  <PModal :show="openedList" @close="openedList = false" class="session-edit">
    <AdminSponsorsSelectList @select="addSponsor" />
  </PModal>
</template>

<script lang="ts" setup>
import Event from '@/classes/Event';

import Sponsor from '@/classes/Sponsor';
const event: Event = Store.Events.Item;

const openedList: Ref<boolean> = ref(false);

const addSponsor = async (item: Sponsor) => {
  openedList.value = false;
  if (PF.H.Classes.ExistsWithId(event.sponsors, item.id)) {
    PF.Notification.Warning('Данный спонсор уже добавлен');
    return;
  }
  // Provider.store.dispatch('m2m/create', event.value.bindSponsor(item));
};

const removeSponsor = async (id: string) => {
  // Provider.store.dispatch('m2m/remove', event.value.unbindSponsor(id));
};

const openSponsorsList = () => {
  openedList.value = true;
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

.container {
  width: 100%;
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

.el-select {
  width: 100%;
}

:deep(.el-date-editor.el-input, .el-date-editor.el-input__inner) {
  width: 100%;
}

:deep(.el-form-item) {
  display: block;
  margin-bottom: 16px;
}

:deep(.el-input__inner) {
  height: 40px;
  width: 100%;
  display: flex;
  font-family: Comfortaa, Arial, Helvetica, sans-serif;
  font-size: 15px;
  color: #343d5c;
}

:deep(.el-input__inner::sponsorholder) {
  color: #b0a4c0;
}

:deep(.el-input__icon) {
  color: #343d5c;
}

:deep(.el-form-item__label) {
  color: #b0a4c0;
  padding: 0 !important;
  text-transform: uppercase;
  margin-left: 5px;
  font-size: 14px;
  margin-bottom: 6px;
}

:deep(.el-input-number__increase) {
  border-radius: 0;
}

:deep(.el-input-number__decrease) {
  border-radius: 0;
}

:deep(.el-textarea__inner) {
  font-family: Comfortaa, Arial, Helvetica, sans-serif;
  font-size: 15px;
  color: #343d5c;
}

.button-add {
  border: none;
  background: inherit;
  color: #00b5a4;
  transition: 0.3s;
  cursor: pointer;
}

.button-add:hover {
  color: #009e8f;
}

.button-del {
  position: absolute;
  top: 13px;
  right: 36px;
  border: none;
  background: inherit;
  color: #a3a9be;
  transition: 0.3s;
  cursor: pointer;
}

.button-del:hover {
  color: #b63414;
}

.bottom-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-right: 15px;
}

.list-number {
  position: absolute;
  top: 10px;
  right: 10px;
  width: 22px;
  height: 22px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #ffffff;
  background: #1979cf;
  border-radius: 20px;
}

.title {
  font-family: 'Open Sans', sans-serif;
  font-size: 14px;
  color: #c4c4c4;
  margin: 10px;
  color: #b0a4c0;
  text-transform: uppercase;
  margin-left: 5px;
  font-size: 14px;
}

.sponsor-item {
  width: calc(100% - 22px);
  position: relative;
  padding: 0 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  margin: 0px;
  background: #dff2f8;
}

.sponsor-item-item {
  position: relative;
  width: calc(100% - 42px);
  margin: 0px 10px 20px 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  padding: 12px 10px;
  background: #f9fafb;
}

.session-edit {
  position: absolute;
  top: 20px;
  right: 20px;
  border: 1px solid #343e5c;
  border-radius: 5px;
  width: 100%;
  background: #ffffff;
  padding: 10px;
  z-index: 5;
}
</style>
