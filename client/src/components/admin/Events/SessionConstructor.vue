<template>
  <el-form class="block">
    <InfoItem title="Название" :with-open-window="false" :with-hover="false" baseBoxMargin="0 0 10px 0" padding="0 5px">
      <el-form-item style="width: 100%" prop="name">
        <el-input v-model="session.name" @blur="update" />
      </el-form-item>
    </InfoItem>
    <InfoItem title="Выбор" :with-open-window="false" :with-hover="false" baseBoxMargin="0 0 10px 0" padding="0 0 0 5px">
      <RemoteSearch :key-value="'speaker'" placeholder="Выберите председателя" max-width="100%" @select="addChair" />
    </InfoItem>
    <div>Председатели:</div>
    <div v-for="(chair, i) in session.chairs" :key="chair.id" class="speaker-field">
      {{ chair.human.getFullName() }}
      <div class="list-number">
        {{ i + 1 }}
      </div>
      <div class="button-del" @click="removeChair(chair.id)">Удалить</div>
    </div>
  </el-form>
</template>

<script setup lang="ts">
import Session from '@/classes/Session';
import Speaker from '@/classes/Speaker';

const props = defineProps({
  session: {
    type: Object as PropType<Session>,
    required: true,
  },
});

const speaker: Speaker = SpeakersStore.Item;
const emits = defineEmits(['close']);

const update = async () => {
  await SessionsStore.Update(props.session);
};
const close = async () => {
  emits('close');
};

const addChair = async (res: ISearchElement) => {
  await SpeakersStore.Get(res.value);
  // await Provider.store.dispatch('m2m/create', props.session.bindChair(speaker.value));
};

const removeChair = async (id?: string) => {
  // await Provider.store.dispatch('m2m/remove', props.session.unbindChair(id as string));
};
</script>

<style lang="scss" scoped>
.speaker-field {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  margin: 10px 0px;
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
</style>
