<template>
  <!-- <PageTopSections v-if="mounted" :components="components" :admin-header-params="headerParams" /> -->
</template>

<script setup lang="ts">
import Perfom from '@/classes/Perfom';
const perfom: Perfom = PerfomsStore.Item;
const mounted = ref(false);

const components = {
  AdminPerfomPageInfo: AdminPerfomPageInfo,
};

const load = async () => {
  PF.AdminUI.Head.Set(
    computed(() => (PF.H.Router.Id() ? perfom.name : 'Добавить мероприятие')),
    [Button.Warning('Удалить', remove)]
  );
  if (PF.H.Router.Id()) {
    await PerfomsStore.Get();
    mounted.value = true;
    return;
  }
  const perfom = new Perfom();
  await PerfomsStore.Create(perfom);
};

const remove = async () => {
  await PerfomsStore.Remove(PF.H.Router.Id());
  await PF.H.Router.To('/admin/perfoms');
};

Hooks.onBeforeMount(load);
Hooks.onBeforeRouteLeave();
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

.field {
  width: 100%;
  height: 100%;
  background: red;
  padding: 10px;
}

.tab {
  font-size: 12px;
  text-transform: uppercase;
  color: #b0a4c0;
  border: 1px solid #999999;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
  text-align: center;

  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */

  background: #f5f5f5;
  margin: -0.5px;
}

.tab:hover {
  background: #dff2f8;
}

.selected-tab {
  font-size: 12px;
  font-weight: bold;
  text-transform: uppercase;
  color: #343d5c;
  position: relative;
  border: 1px solid #343d5c;
  height: 40px;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  cursor: pointer;

  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */

  background: #dff2f8;
  margin: -0.5px;
  z-index: 1;
}
</style>
