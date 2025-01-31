<template>
  <div class="title-user-info">Расскажите о себе, чтобы присоединиться к вебинару:</div>
  <el-form :model="user" :label-position="'left'" :rules="rules" ref="form">
    <el-form-item prop="lastName">
      <el-input placeholder="Фамилия" v-model.trim="user.lastName" />
    </el-form-item>
    <el-form-item prop="firstName">
      <el-input placeholder="Имя" v-model.trim="user.firstName" />
    </el-form-item>
    <el-form-item>
      <el-input placeholder="Отчество" v-model.trim="user.patronName" />
    </el-form-item>
    <el-form-item prop="phone">
      <el-input placeholder="Контактный телефон" v-model.trim="user.phone" />
    </el-form-item>
    <el-form-item prop="company">
      <el-input placeholder="Организация" v-model="user.company" />
    </el-form-item>
    <el-form-item prop="post">
      <el-input placeholder="Должность" v-model="user.post" />
    </el-form-item>
    <div class="btn-group">
      <button class="button" @click.prevent="submit">Сохранить</button>
      <button class="button" @click.prevent="logout">Выйти</button>
    </div>
  </el-form>
</template>

<script setup lang="ts">
const auth = computed(() => Provider.store.getters['auth/auth']);
const user = ref(auth.value.user?.get());
const form = ref();

const privacyRule = (_: unknown, value: string, callback: MyCallbackWithOptParam) => {};
const rules = ref({
  firstName: [{ required: true, trigger: 'blur', message: 'Необходимо ввести имя' }],
  lastName: [{ required: true, trigger: 'blur', message: 'Необходимо ввести фамилию' }],
  // phone: [{ validator: phoneRule, trigger: 'blur' }],
  company: [{ required: true, trigger: 'blur', message: 'Необходимо ввести компанию' }],
  post: [{ required: true, trigger: 'blur', message: 'Необходимо указать занимаемую должность' }],
  privacy: [{ validator: privacyRule, trigger: 'change' }],
});

const submit = async (): Promise<void> => {
  if (!(await validate(form, true))) {
    return;
  }
  await Provider.store.dispatch('users/update', user.value);
  auth.value.user.set(user.value);
};
const logout = () => {
  auth.value.logout();
};
</script>

<style scoped lang="scss">
.form-row {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  .el-form-item {
    margin: 15px;
  }
}

@media screen and (max-width: 900px) {
  .form-row {
    flex-direction: column;
    .el-form-item {
      width: 100%;
    }
  }
}
.auth-btn {
  color: white;
  background-color: #224af2;
}
.btn-group {
  width: 100%;
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
}

.title-user-info {
  color: #ffffff;
  display: flex;
  justify-content: left;
  align-items: center;
  height: 40px;
  font-size: 18px;
  margin-bottom: 10px;
}
</style>
