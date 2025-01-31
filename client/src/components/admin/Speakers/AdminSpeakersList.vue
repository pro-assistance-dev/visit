<template>
  <AdminListWrapper v-if="mounted" pagination show-header>
    <AdminSpeakersListFilters />
    <div class="scroll-block">
      <div class="patient-count">Количество выступающих: {{ count }}</div>
      <div v-for="speaker in speakers" :key="speaker.id">
        <CollapseItem :is-collaps="false" padding="0 8px">
          <template #inside-title>
            <div class="flex-block" @click.prspeaker="() => undefined">
              <div class="item-flex">
                <div class="line-item-left">
                  <Button
                    v-if="showButton"
                    button-class="edit-button"
                    color="#006bb4"
                    icon="edit"
                    icon-class="edit-icon"
                    @click="edit(speaker.id)"
                  />
                  <FioToggleForm :human="speaker.human" />
                </div>
              </div>
              <div class="item-flex">
                <GridContainer custom-class="grid">
                  <div class="line-item-left">
                    <AdminSpeakerRegalias :speaker="speaker" @update="update(speaker)" />
                  </div>
                </GridContainer>
              </div>
            </div>
          </template>
        </CollapseItem>
      </div>
    </div>
  </AdminListWrapper>
  <PModal :show="showAddModal" title="Добавить выступающего" @close="showAddModal = false">
    <CreateSpeakerForm @add="showAddModal = false" />
  </PModal>
</template>

<script setup lang="ts">
import Speaker from '@/classes/Speaker';
import SpeakersSortsLib from '@/libs/sorts/SpeakersSortsLib';

const speakers: Speaker[] = SpeakersStore.Items();
const count: Ref<number> = SpeakersStore.Count();
const showAddModal = ref(false);
const showButton = ref(false);
const mounted = ref(false);

const load = async () => {
  await SpeakersStore.FTSP();
  PF.AdminUI.Head.Set('Выступающие', [Button.Primary('Добавить', () => (showAddModal.value = true))]);
  mounted.value = true;
};

onMounted(() => {
  showButton.value = true;
});

Hooks.onBeforeMount(load);

const edit = async (id?: string) => {
  if (!id) {
    return;
  }
  await PF.H.Router.To(`/admin/speakers/${id}`);
};

const remove = async (id: string) => {
  await SpeakersStore.Remove(id);
};

const update = async (speaker: Speaker) => {
  await SpeakersStore.Update(speaker);
};
</script>

<style lang="scss" scoped>
.button {
  width: auto;
  height: 34px;
  border-radius: 5px;
  color: #006bb4;
  font-size: 12px;

  &-filter {
    background: #ffffff;
  }

  &-download {
    background: #dff2f8;
  }
}

:deep(.button-register) {
  width: auto;
  height: 34px;
  border-radius: 5px;
  color: #006bb4;
  background: #ffffff;
  font-size: 12px;
}

:deep(.name-item) {
  margin: 0;
  width: auto;
  border-color: #ffffff;
  padding: 0;
}

.grid {
  grid-gap: 10px;
  margin: 0;
}

.plus-button {
  width: 100%;
  height: 34px;
  border-radius: 5px;
  color: #00bea5;
  background: #c1efeb;
}

.save-picker-button {
  width: 100%;
  height: 34px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
}

.gender-button {
  width: 42px;
  border-radius: 5px;
  height: 42px;
  color: #006bb4;
  background: #dff2f8;
  margin: 2px 10px 0 0;
  font-size: 18px;
}

.save-button {
  width: 100%;
  border-radius: 5px;
  height: 42px;
  color: #006bb4;
  background: #dff2f8;
  margin: 2px 10px 0 0;
  font-size: 14px;
}

:deep(.edit-button) {
  min-width: 40px;
  max-width: 40px;
  height: 40px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
  margin: 0;
  margin-right: 10px;
}

:deep(.files-buttons) {
  width: auto;
  height: 34px;
  border-radius: 5px;
  color: #006bb4;
  background: #dff2f8;
  font-size: 12px;

  &:hover {
    background: #dff2f8;
  }
}

.edv {
  font-size: 14px;
  padding: 0;
  margin: 0 5px 0 0;

  &-active {
    color: #b0a4c0;
  }
}

.patient-name {
  color: #006bb4;
  font-size: 17px;
  min-width: 150px;
  width: 100%;
  padding: 0;
}

.hidden {
  display: none;
}

.scroll-block {
  height: calc(100% - 120px);
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

:deep(.edit-icon) {
  width: 28px;
  height: 28px;
  color: #006bb4;
}

.item-flex {
  display: flex;
  width: 100%;
  justify-content: space-between;
  align-items: center;
  margin: 0 10px 0 0;
}

.item-flex:last-child {
  margin: 0;
}

.line-item-left {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  // margin-right: 10px;
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

:deep(.icon-plus) {
  fill: #00b5a4;
  width: 24px;
  height: 24px;
  cursor: pointer;
}

.icon-filter {
  width: 24px;
  height: 24px;
  cursor: pointer;
  stroke: #006bb4;
  fill: none;
}

.icon-del {
  width: 10px;
  height: 10px;
  cursor: pointer;
}

.patient-count {
  margin-top: 10px;
  color: #b0a4c0;
  font-size: 14px;
}

@media (max-width: 1600px) {
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
</style>
