<template>
  <div class="wrapper">
    <div class="event-layout page-title">
      <h1>{{ event.name }}</h1>
      <button class="more-info-btn button" @click.prevent="logout">
        <span>Выйти из системы</span>
        <img :src="exitImg" />
      </button>
    </div>
    <div class="page-container">
      <div class="event-layout" style="padding: 0">
        <div class="media-container">
          <div class="media-container-stream">
            <!-- <h1 v-if="!started && event.id === '168f7119-7559-477f-ba55-47271a261363'">Трансляция начнётся 25.09.24</h1> -->
            <!-- <VideoPlayer :event="event" /> -->
          </div>
          <div class="media-container-chat">
            <Chat :chat="event.chat" :user="auth.Status.User" />
            <!-- <VideoStreamChat /> -->
          </div>
        </div>
      </div>
    </div>
    <div class="page-bottom"></div>
  </div>
</template>

<script lang="ts" setup>
import exitImg from '@/assets/img/icons/logout.png';
import Event from '@/classes/Event';

const event = Store.Events.Item;
const auth = PF.Auth;

const logout = async () => {
  auth.Status.Logout();
};

const start = new Date('25 Sep 2024 13:50');
const started = ref(false);

let i: number | undefined = undefined;

onMounted(async () => {
  i = window.setInterval(async () => {
    checkTime();
  }, 5000);
});
const checkTime = () => {
  const d = new Date();
  console.log(d, start, started.value, d > start);
  if (d > start) {
    started.value = true;
    clearInterval(i);
  }
};
</script>

<style scoped lang="scss">
.wrapper {
  background: no-repeat url('@/assets/img/header_bg2.png');
  background-size: cover;
}

// .logout-icon {
//   height: 25px;
//   width: 25px;
//   background-color: #224af2;
//   -webkit-mask-size: contain;
//   mask-size: contain;
//   -webkit-mask-position: center;
//   mask-position: center;
//   -webkit-mask-repeat: no-repeat;
//   mask-repeat: no-repeat;
//   mask-image: url('~/assets/img/icons/logout.svg');
// }
.media-container {
  display: flex;
  justify-content: space-between;
  height: 475px;
  width: 100%;
  // position: relative;

  &-stream {
    flex: 1;
    padding: 0;
  }

  &-chat {
    // flex: 1 1 auto;
  }
}

.blur {
  position: absolute;
  background: rgba(255, 255, 255, 0.9);
  padding: 10px;
  text-align: center;
  width: 100%;
  height: 100%;
  z-index: 2;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;

  &-timer {
    margin-bottom: 50px;
  }
}

h1 {
  font-size: 30px;
  font-weight: bold;
  text-align: center;
  color: white;
}

.page-title {
  display: flex;
  align-items: center;
  justify-content: space-around;

  h1 {
    padding: 20px 0;
  }

  .button {
    height: 50px;
  }
}

.page-bottom {
  display: flex;
  align-items: center;
  padding: 20px 0;
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

@media screen and (max-width: 1230px) {
  .media-container {
    flex-direction: column;
    height: auto;
  }

  .button {
    span {
      display: none;
    }
  }

  .page-bottom {
    display: none;
  }
}

@media screen and (max-width: 1024px) {
  h1 {
    font-size: 27px;
  }

  .button {
    font-size: 16px;
    padding: 5px 10px;
  }
  .callback-btn {
    padding: 5px 10px;
    border-radius: 10px;
    i {
      font-size: 15px;
      margin-right: 5px;
    }
  }
  .more-info-btn {
    i {
      font-size: 20px;
    }
    img {
      height: 20px;
    }
  }
}

@media screen and (max-width: 768px) {
  h1 {
    font-size: 24px;
  }
}

@media screen and (max-width: 480px) {
  h1 {
    font-size: 20px;
  }
}
</style>
