<template>
  <PString :string="title" font-size="20px" margin="20px auto" />
  <GridContainer
    grid-template-columns="repeat(auto-fit, minmax(30%, 1fr))"
    background="#ffffff"
    max-width="100%"
    grid-gap="10px"
    margin="0 auto"
  >
    <PInput v-model="user.firstName" margin="0" label="Имя" placeholder="Введите Ваше имя" />
    <PInput v-model="user.phone" margin="0" label="Номер телефона" placeholder="Введите ваш номер телефона" />
  </GridContainer>

  <button class="auth-btn button" @click.prevent="submit">Отправить</button>

  <!-- <PButton skin="base" type="primary" text="Отправить" margin="40px auto 0 auto" justify-content="center" @click="submit" /> -->

  <!-- <div class="form-container">
    <el-form :model="user" :label-position="'left'" :rules="rules" ref="form">
      <h3>Обратный звонок</h3>
      <el-form-item prop="firstName">
        <el-input placeholder="Имя" v-model.trim="user.firstName" />
      </el-form-item>
      <el-form-item prop="phone">
        <el-input placeholder="Контактный телефон" v-model.trim="user.phone" />
      </el-form-item>
      <div class="btn-group">
        <button class="auth-btn button" @click.prevent="submit">Отправить</button>
      </div>
    </el-form>
  </div> -->
</template>

<script lang="ts" setup>
import User from '@/classes/User';

defineProps({
  title: {
    type: String,
    default: 'Обратный звонок',
  },
});

const emits = defineEmits(['send']);

const form = ref();
const user = ref(new User());
const phoneRule = (_: unknown, value: string, callback: MyCallbackWithOptParam) => {
  const phoneRegExp = /^(\+7|7|8)?[\s\-]?\(?[489][0-9]{2}\)?[\s\-]?[0-9]{3}[\s\-]?[0-9]{2}[\s\-]?[0-9]{2}$/;
  if (!value) {
    callback(new Error('Необходимо указать контактный номер телефона'));
  } else if (!phoneRegExp.test(value)) {
    callback(new Error('Пожалуйста, введите корректный номер телефона'));
  } else {
    callback();
  }
};
const rules = ref({
  firstName: [{ required: true, trigger: 'blur', message: 'Необходимо ввести имя' }],
  phone: [{ validator: phoneRule, trigger: 'blur' }],
});

const submit = async () => {
  if (!(await validate(form, true))) {
    return;
  }
  emits('send', user.value.firstName, user.value.phone);
  user.value = new User();
};
</script>

<style lang="scss" scoped>
.form-container {
  max-width: 700px;
  margin: 0 auto;
}
h3 {
  font-size: 24px;
  font-weight: bold;
  text-align: center;
  margin-bottom: 20px;
}
@media screen and (max-width: 1024px) {
  h3 {
    font-size: 20px;
  }
}
@media screen and (max-width: 480px) {
  h3 {
    font-size: 18px;
  }
}
</style>
