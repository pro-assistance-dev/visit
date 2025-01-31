<template>
  <el-form ref="form" :model="humanCopy" style="width: 100%">
    <InfoItem
      title="ФИО"
      :show-save-dialog="showSaveDialog"
      icon="edit-title"
      :with-hover="true"
      :close="closeToggle"
      width="100%"
      @keyup-enter="submit"
      @after-close="resetCopy"
    >
      <PString v-if="human.getFullName().length > 3" :string="human.getFullName()" />
      <PString v-else string="Введите данные" color="#B0A4C0" />
      <template #open-inside-content>
        <GridContainer custom-class="grid">
          <InfoItem
            title="фамилия"
            icon="edit-title"
            :with-open-window="false"
            :with-hover="false"
            border-color="#ffffff"
            base-box-margin="0 0 15px 0"
            padding="0"
            width="100%"
          >
            <el-form-item style="width: 100%" prop="surname" @change="setFilled">
              <el-input v-model="humanCopy.surname" />
              <!-- <el-input :model-value="human.surname" @input="(e) => human.setSurname(e)" @click.stop="() => undefined" /> -->
            </el-form-item>
          </InfoItem>
          <InfoItem
            title="имя"
            icon="edit-title"
            :with-open-window="false"
            :with-hover="false"
            border-color="#ffffff"
            base-box-margin="0 0 15px 0"
            padding="0"
            width="100%"
          >
            <el-form-item style="width: 100%" prop="name" @change="setFilled">
              <el-input v-model="humanCopy.name" />
              <!-- <el-input :model-value="human.name" @input="(e) => human.setName(e)" @click.stop="() => undefined" /> -->
            </el-form-item>
          </InfoItem>
          <InfoItem
            title="отчество"
            icon="edit-title"
            :with-open-window="false"
            :with-hover="false"
            border-color="#ffffff"
            base-box-margin="0 0 15px 0"
            padding="0"
            width="100%"
          >
            <el-form-item style="width: 100%" prop="patronymic" @change="setFilled">
              <el-input v-model="humanCopy.patronymic" />
              <!-- <el-input :model-value="human.patronymic" @input="(e) => human.setPatronymic(e)" @click.stop="() => undefined" /> -->
            </el-form-item>
          </InfoItem>
          <Button button-class="save-button" text="Сохранить" @click="submit" />
        </GridContainer>
      </template>
    </InfoItem>
  </el-form>
</template>

<script lang="ts">
const props = defineProps({
  human: {
    type: Object as PropType<Human>,
    required: true,
  },
});

const closeToggle: Ref<boolean> = ref(false);
const form = ref();
const humanCopy: Ref<Human> = ref(new Human(props.human));
const showSaveDialog: Ref<boolean> = ref(false);

const resetCopy = () => {
  humanCopy.value = new Human(props.human);
  showSaveDialog.value = false;
};

const updateHuman = async (): Promise<void> => {
  props.human.setFullName(humanCopy.value);
  await Provider.withHeadLoader(async () => {
    await Provider.store.dispatch('humans/update', props.human);
  });
};

const submit = async (): Promise<void> => {
  if (!validate(form)) {
    return;
  }
  closeToggle.value = !closeToggle.value;
  await updateHumanName();
};

const updateHumanName = async (): Promise<void> => {
  props.human.setEditNameMode(false);
  await updateHuman();
};

const setFilled = () => {
  showSaveDialog.value = true;
};
</script>
<style lang="scss" scoped>
.grid {
  max-width: auto;
  grid-gap: 10px;
  margin: 0;
  grid-template-columns: repeat(auto-fit, minmax(100%, 1fr));
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

:deep(.el-form-item) {
  margin: 0;
}
</style>
