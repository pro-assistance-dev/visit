<template>
  <el-form>
    <InfoItem title="Название" :with-open-window="false" :with-hover="false" baseBoxMargin="0 0 10px 0" padding="0 5px">
      <el-form-item style="wigth: 100%" prop="surname" label="">
        <el-input v-model="perfom.name" placeholder="Введите название" formatter="firstLetterUpper" @blur="update" />
      </el-form-item>
    </InfoItem>
    <InfoItem title="Описание" :with-open-window="false" :with-hover="false" baseBoxMargin="0 0 10px 0" padding="0px 5px">
      <el-form-item style="wigth: 100%" prop="description" label="">
        <el-input v-model="perfom.description" placeholder="Введите описание" formatter="firstLetterUpper" @blur="update" />
      </el-form-item>
    </InfoItem>
    <InfoItemSelectButtons
      title="Тип выступления"
      :selected="perfom.type"
      :list="types"
      @select="setType"
      margin="10px 0 0 0"
      background="#ffffff"
    />
    <InfoItemSelectButtons
      title="Формат выступления"
      :selected="perfom.format"
      :list="formats"
      @select="setFormat"
      margin="10px 0 0 0"
      background="#ffffff"
    />
    <div class="place-item">
      <div class="bottom-buttons">
        <div class="title">Спикеры:</div>
      </div>
      <InfoItem
        title=""
        :with-open-window="false"
        :with-hover="false"
        baseBoxMargin="0 10px 20px 10px"
        padding="0px 5px"
        background="#ffffff"
      >
        <RemoteSearch :key-value="'speaker'" placeholder="Найдите спикера для добавления" max-width="100%" @select="selectSearch" />
      </InfoItem>

      <div v-for="(speaker, i) in perfom.speakers" :key="speaker.id" class="speaker-field">
        {{ speaker.human.getFullName() }}
        <div class="list-number">
          {{ i + 1 }}
        </div>
        <div class="button-del" @click="removeSpeaker(speaker.id)">Удалить</div>
      </div>
    </div>
    <Button button-class="save-button" text="Удалить выступление" @click="remove" margin="20px 0 0 0" />
  </el-form>
</template>

<script lang="ts" setup>
import Perfom from '@/classes/Perfom';
import Speaker from '@/classes/Speaker';
const speaker = SpeakersStore.Item;
const emit = defineEmits(['remove']);

const props = defineProps({
  perfom: {
    type: Object as PropType<Perfom>,
    required: true,
  },
});
const types = [
  'Доклад',
  'Диалог',
  'Дискуссия',
  'Мастер-класс',
  'Буктрейлер',
  'Поиск решения',
  'Практика',
  // 'Секция с докладами',
  // 'Диалог',
  // 'Дискуссия',
  // 'Мастер-класс',
  // 'Буктрейлер - короткий видеоролик',
  // 'Квест - поиск решения в команде',
  // 'Кейс - применение имеющихся знаний на практическом примере',
];
const setType = async (t: string) => {
  props.perfom.type = t;
  await update();
};
const setFormat = async (t: string) => {
  props.perfom.format = t;
  await update();
};
const formats = ['Очно', 'Онлайн', 'Запись'];
const update = async () => {
  await PerfomsStore.Update(props.perfom);
};

const selectSearch = async (res: ISearchElement) => {
  if (props.perfom.speakerExists(res.value)) {
    // ElMessage.warning('Спикер уже добавлен')
    return;
  }
  await SpeakersStore.Get(res.value);
  await addSpeaker();
};

const addSpeaker = async () => {
  // Provider.store.dispatch('m2m/create', props.perfom.bindSpeaker(speaker.value));
};

const removeSpeaker = async (id: string) => {
  // Provider.store.dispatch('m2m/remove', props.perfom.unbindSpeaker(id));
};
const remove = () => {
  emit('remove');
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
  margin: 2px 0 0 0;
  width: 100%;
}

:deep(.el-input__inner) {
  width: 100%;
  border: none;
  padding: 0;
  height: 36px;
}

:deep(.el-select) {
  width: 100%;
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

.speaker-field {
  position: relative;
  width: calc(100% - 42px);
  margin: 0px 10px 20px 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  padding: 12px 10px;
  background: #f9fafb;
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

.place-item {
  width: calc(100% - 22px);
  position: relative;
  padding: 0 10px;
  border: 1px solid #c3c3c3;
  border-radius: 5px;
  margin: 10px 0;
  background: #dff2f8;
}

.bottom-buttons {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-right: 15px;
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
</style>
