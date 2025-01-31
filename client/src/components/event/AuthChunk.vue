<template>
  <div class="auth-card">
    <div class="auth-card-header">
      <template v-if="!showRefreshForm">
        <h3 v-if="!showLoginForm && !showRestoreForm">Зарегистрируйтесь <span v-if="isEvent">для участия</span></h3>
        <h3 v-else-if="showLoginForm">Авторизируйтесь <span v-if="isEvent">для участия</span></h3>
        <h3 v-else-if="showRestoreForm">Введите email для восстановления пароля</h3>
      </template>
      <h3 v-if="showRefreshForm">Введите новый пароль</h3>
    </div>
    <div>
      <template v-if="!showRefreshForm">
        <LoginChunk v-if="showLoginForm" @toggle="toggleForm" @restore="showRestore" />
        <RegisterChunk v-else-if="!showLoginForm && !showRestoreForm" @toggle="toggleForm" @restore="showRestore" />
        <RestoreChunk v-else-if="showRestoreForm" @login="openLoginForm" @register="openRegisterForm" />
      </template>
      <RefreshChunk v-if="showRefreshForm" />
    </div>
  </div>
</template>

<script lang="ts" setup>
const props = defineProps({
  isEvent: {
    type: Boolean,
    default: false,
  },
});

const email = ref('');
const showLoginForm: Ref<boolean> = ref(true);
const showRestoreForm: Ref<boolean> = ref(false);
const showRefreshForm: Ref<boolean> = ref(false);
const toggleForm = () => {
  showLoginForm.value = !showLoginForm.value;
};
const loginStatus: ComputedRef<'login' | 'register' | 'forgotPassword' | 'passwordChange'> = computed(
  () => Provider.store.getters['auth/loginStatus']
);

const openLoginForm = () => {
  showLoginForm.value = true;
  showRestoreForm.value = false;
};

const openRegisterForm = () => {
  showLoginForm.value = false;
  showRestoreForm.value = false;
};

onBeforeMount(() => {
  if (loginStatus.value === 'passwordChange') {
    showRefreshForm.value = true;
    showLoginForm.value = false;
    showRestoreForm.value = false;
  } else {
    Provider.store.commit('auth/setLogin');
  }
});

const showRestore = async () => {
  showLoginForm.value = false;
  showRestoreForm.value = true;
};
</script>

<style scoped lang="scss">
h3 {
  font-size: 24px;
  font-weight: bold;
  word-break: break-word;
  margin: 0 0 8px 0;
}
.auth-card {
  background: white;
  padding: 30px 50px;
  text-align: center;
  margin: auto;
  border-radius: 20px;
  max-width: 700px;

  &-header {
    display: flex;
    justify-content: space-around;
    margin: 10px 0 10px;
  }
}
.auth-card-header {
  margin-bottom: 15px;
}

.auth-login-link {
  color: blue;
  margin: 10px 0 5px 0;

  &:hover {
    cursor: pointer;
    text-decoration: underline;
  }
}

@media screen and (max-width: 1024px) {
  h3 {
    font-size: 20px;
  }
  h2 {
    font-size: 18px;
  }
  .auth-card {
    width: auto;
  }
}
@media screen and (max-width: 480px) {
  h3 {
    font-size: 18px;
  }
  .auth-card {
    padding: 0;
  }
}
</style>
