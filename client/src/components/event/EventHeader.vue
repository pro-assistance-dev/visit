<template>
  <div class="event-layout">
    <div class="header-container">
      <img class="logo-img" :src="ProsLogo" />
      <CarouselBanners :banners="bannerArray" :showControl="false" height="100px" max-width="900px" />
      <div class="header-phone">
        <button class="callback-btn button" @click="callbackHandler">
          <span> Обратный звонок</span>
        </button>
      </div>
    </div>
    <div class="header-container-mobile">
      <PFlex width="100%" justify-content="space-between" margin="20px 0">
        <img class="logo-img" :src="ProsLogo" />
        <div class="header-phone">
          <button class="callback-btn button" @click="callbackHandler">
            <span> Обратный звонок</span>
          </button>
        </div>
      </PFlex>
      <CarouselBanners :banners="bannerArray" :showControl="false" height="100px" max-width="900px" />
    </div>
    <PModal :show="callbackFormShow" :closable="true" @close="callbackHandler">
      <CallBackForm @send="sendDataToBot" />
    </PModal>

    <!-- <PModal :show="showAddModal" title="Добавить спонсора" @close="showAddModal = false">
      <CreateSponsorForm @add="showAddModal = false" />
    </PModal> -->
  </div>
</template>

<script lang="ts" setup>
import ProsLogo from '@/assets/img/logo2.png';

import Sponsor from '@/classes/Sponsor';
import IBannerArray from '@/interfaces/IBannerArray';

const event = Store.Events.Item;

const fullscreen: Ref<boolean> = ref(window.matchMedia('(max-width: 768px)').matches);

const bannerArray: IBannerArray[] = [
  { id: '1', bannerPath: new URL('@/assets/img/logo4.png', import.meta.url).href, bannerLink: '' },
  { id: '2', bannerPath: new URL('@/assets/img/logo8.png', import.meta.url).href, bannerLink: '' },
  { id: '3', bannerPath: new URL('@/assets/img/logo10.png', import.meta.url).href, bannerLink: '' },
  { id: '4', bannerPath: new URL('@/assets/img/logo4.png', import.meta.url).href, bannerLink: '' },
  { id: '5', bannerPath: new URL('@/assets/img/logo8.png', import.meta.url).href, bannerLink: '' },
  { id: '6', bannerPath: new URL('@/assets/img/logo10.png', import.meta.url).href, bannerLink: '' },
];

const callbackFormShow: Ref<boolean> = ref(false);

const sendDataToBot = async (name: string, phone: string) => {
  const message = `Заявка на обратный звонок%0AИмя: ${name}%0AТелефон: ${phone}`;
  callbackFormShow.value = false;
  try {
    // await Provider.store.dispatch('telegram/sendCallBackMessage', message);
  } catch {
    // ElNotification.error({
    //   title: 'Обратный звонок',
    //   message: 'Сервис временно недоступен. Пожалуйста, воспользуйтесь указанным номером телефона',
    // });
    return;
  }
  // ElNotification.success({
  //   title: 'Обратный звонок',
  //   message: 'Спасибо за заявку, с вами свяжутся в ближайшее время',
  // });
  callbackFormShow.value = false;
};
const toggleCallbackDialog = () => {
  callbackFormShow.value = !callbackFormShow.value;
};
const callbackHandler = async () => {
  toggleCallbackDialog();
  // if (!isAuth.value) {
  //   toggleCallbackDialog();
  // } else {
  //   await sendDataToBot(user.value.firstName, user.value.phone);
  // }
};

const handleResize = () => {
  fullscreen.value = window.matchMedia('(max-width: 768px)').matches;
};
onMounted(async () => {
  window.addEventListener('resize', handleResize);
});
onUnmounted(() => window.removeEventListener('resize', handleResize));
</script>

<style lang="scss" scoped>
.header-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 33px auto;
  max-width: 1170px;
}

.logo-img {
  height: 74px;
  padding: 0 10px 0 0;
}

.header-phone {
  font-size: 24px;
  font-weight: bold;
  display: flex;
  align-items: center;
  width: 200px;

  span {
    white-space: nowrap;
    margin-right: 10px;
  }
}

.button {
  padding: 5px 15px;
  border: 1px solid #dcdfe6;
  background-color: #ffffff;
  color: #263238;
  border-radius: 15px;
  font-size: 18px;
  margin: 5px;
  display: flex;
  align-items: center;
  white-space: nowrap;
}

.button:hover {
  box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}

.callback-btn {
  font-weight: bold;
}

.header-container-mobile {
  display: none;
  margin: 33px auto;
  max-width: 1170px;
}

@media screen and (max-width: 768px) {
  .header-container-mobile {
    display: block;
  }

  .header-container {
    display: none;
  }
}
</style>
