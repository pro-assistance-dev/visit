<template>
  <div class="header-container">
    <div class="conf">
      <div>
        <div class="conf-type">
          <img src="../src/assets/img/ellipse.png" />
          <span style="margin-left: 5px"> {{ event.type }} </span>
        </div>
        <div class="conf-name">
          {{ event.name }}
        </div>
      </div>
      <div class="conf-description">
        {{ event.description }}
      </div>
      <div class="btn-group">
        <!-- <button class="more-info-btn button" @click="$scroll('#program')"> -->
        <!--   <span>Узнать подробнее</span><el-icon><Right /></el-icon> -->
        <!-- </button> -->
      </div>
    </div>
    <div class="auth-form">
      <AuthForm v-if="!PF.Auth.Status.IsAuth" @submit="submitAuthForm" />
      <!-- <UserInfoForm v-else-if="PF.Auth.isAuth && !user.isFullInfo()" /> -->
    </div>
  </div>
</template>

<script setup lang="ts">
// import Event from '@/classes/Event';
//
const event = Store.Events.Item;

const submitAuthForm = async () => {
  if (PF.Auth.Form.IsLogin()) {
    await login();
  }
  if (PF.Auth.Form.IsRegister()) {
    await register();
  }
};
const login = async () => {
  await PF.S.Auth.Login();
};

const register = async () => {
  await PF.S.Auth.Register();
};
// const auth = computed(() => Provider.store.getters['auth/auth']);
// const user = computed(() => auth.value.user?.get());
</script>

<style scoped>
.header-container {
  display: flex;
  justify-content: space-between;
}

.conf {
  text-align: start;
  line-height: 1.5;
  color: white;
  font-weight: bold;
  margin-bottom: 10px;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  flex: 55%;
  margin-right: 10px;
}
.conf-name,
.conf-description {
  margin: 20px 0;
}
.conf-name {
  font-size: 33px;
}
.auth-form {
  flex: 45%;
  margin-left: 10px;
}

.conf-type {
  display: flex;
  align-items: center;
}

.conf-type,
.conf-description {
  font-size: 24px;
}

.conf-description {
  font-weight: unset;
}

.auth-modal-button {
  display: none;
}
.btn-group {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
  width: auto;
}

@media screen and (max-width: 1200px) {
  .conf,
  .auth-form {
    flex: 50%;
  }
}

@media screen and (max-width: 1024px) {
  .conf-type,
  .conf-description {
    font-size: 18px;
  }
  .conf-type {
    img {
      height: 10px;
    }
  }
  .conf-name {
    font-size: 27px;
  }
  /* .auth-form { */
  /*   display: none; */
  /* } */
  .auth-modal-button {
    display: flex;
  }
  .header-container {
    display: block;
  }
}
@media screen and (max-width: 768px) {
  .conf-name {
    font-size: 24px;
  }
}
@media screen and (max-width: 480px) {
  .conf {
    text-align: center;
  }
  .btn-group {
    align-items: center;
  }
  .conf-name {
    font-size: 20px;
  }
  .conf-type,
  .conf-description {
    font-size: 16px;
  }
}

.more-info-btn {
  border-radius: 50px;
  padding: 10px 20px;
  margin: 5px;
  i,
  img {
    font-size: 25px;
    font-weight: bold;
    margin-left: 10px;
    color: #224af2;
  }
  img {
    height: 25px;
  }
}
</style>
