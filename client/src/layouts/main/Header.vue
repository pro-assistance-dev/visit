<template>
  <div class="header">
    <AdaptiveContainer :menu-width="'300px'" :mobile-width="'768px'">
      <template #menu>
        <div v-for="item in menu" :key="item.name" class="menu">
          <a :href="`#${item.scrollTo}`">{{ item.name }}</a>
        </div>
        <div class="menu">
          <button class="main-btn" @click="openLoginModal">Войти</button>
        </div>
      </template>

      <template #icon>
        <svg class="icon-menu">
          <use xlink:href="#icon-menu"></use>
        </svg>
      </template>
    </AdaptiveContainer>
    <div class="header-bottom main-container">
      <div class="header-bottom-left">
        <Transition name="slide-fade-bottom">
          <div class="phone">
            <div class="phone_logo">
              <Transition name="spin-fade">
                <img :src="logoPng" />
              </Transition>
            </div>
          </div>
        </Transition>
      </div>
      <div class="header-bottom-right">
        <h1>
          Центр содействия <br />
          развитию здравоохранения
        </h1>
        <h4>
          Кроссфункциональня команда, <br />
          нацеленная на результат
        </h4>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import logoPng from '@/assets/img/promo/logo.png';

const menu = [
  { name: 'О нас', scrollTo: 'about' },
  { name: 'Услуги', scrollTo: 'projects' },
  { name: 'Библиотека', scrollTo: 'lib' },
  // { name: 'Материалы', scrollTo: 'materials' },
  { name: 'Контакты', scrollTo: 'contacts' },
];
// const auth: ComputedRef<Auth> = computed(() => Provider.store.getters['auth/auth']);
// const authModal: ComputedRef<AuthModal> = computed(() => Provider.store.getters['auth/modal']);

// const openLoginModal = () => {
//   if (auth.value.isAuth) {
//     return;
//   }
//   authModal.value.open();
// };
//
// const logout = async () => {
//   auth.value.logout();
// };
</script>

<style scoped>
.icon-menu {
  width: 40px;
  height: 40px;
  stroke: #3d1819;
  margin-top: 10px;
  cursor: pointer;
}

.icon-menu:hover {
  stroke: #702527;
}

.sq {
  height: 40px;
  width: 40px;
  background: #0056b3;
}

.flex {
  display: flex;
  width: 100%;
  justify-content: space-between;
}

.menu {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100px;
  padding: 10px;
  font-size: 14px;
  text-transform: uppercase;
  cursor: pointer;

  a {
    color: white;
    padding: 14px;
    font-weight: 500;
    text-transform: uppercase;
    text-decoration: none;
  }
}
.menu {
  a:hover {
    color: #3d1819;
  }
}

.menu:last-child {
  width: auto;
}

.menu-container {
  max-width: 1200px;
  width: 100%;
  margin: 0 auto;
  display: flex;
  justify-content: center;
  align-items: center;
}

.burger {
  display: none;
  width: 40px;
  height: 40px;
  background: #ffffff;
}

.header {
  position: relative;
  width: 100%;
  background: #f57e20;
  background: linear-gradient(135deg, #f57e20 0%, #f1562b 100%);
  height: 760px;
  overflow: hidden;
  color: white;
}

.header-top {
  max-width: 100%;
  margin: 0 auto;
  padding: 48px 15px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

ul {
  list-style-type: none;
  display: flex;
  align-items: center;
  li {
    a {
      color: white;
      padding: 12px;
      font-weight: 500;
      text-transform: uppercase;
      text-decoration: none;
    }
    a:hover {
      color: #0056b3;
      text-decoration: underline;
      transition:
        color 0.15s ease-in-out,
        background-color 0.15s ease-in-out,
        border-color 0.15s ease-in-out,
        box-shadow 0.15s ease-in-out;
      transition-duration: 0.15s, 0.15s, 0.15s, 0.15s;
      transition-timing-function: ease-in-out, ease-in-out, ease-in-out, ease-in-out;
      transition-delay: 0s, 0s, 0s, 0s;
      transition-property: color, background-color, border-color, box-shadow;
    }
  }
}

.header-bottom {
  margin: 0 auto;
  position: relative;
  display: flex;
  // margin: 0 24px;
  &-left,
  &-right {
    width: 100%;
  }
  &-right {
    padding: 80px 0;
    text-align: center;
    h1,
    h4 {
      font-weight: 500;
      line-height: 1.2;
    }
    h1 {
      font-size: 40px;
    }
    h4 {
      font-size: 24px;
    }
  }
  .phone {
    background: {
      image: url('@/assets/img/promo/phone.png');
      repeat: no-repeat;
    }
    width: 580px;
    height: 1050px;
    right: 0;
    left: 0;
    position: absolute;
    display: flex;
    justify-content: center;

    .phone_logo {
      img {
        width: 100%;
      }
      margin-top: 50%;
      width: 40%;
    }
  }
}
.slide-fade-bottom {
  &-enter-active {
    transition: all 0.5s ease-out;
  }
  &-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
  }
  &-enter-from,
  &-leave-to {
    transform: translateY(30px);
    opacity: 0;
  }
}
.spin-fade {
  &-enter-active {
    transition: all 0.5s ease-out;
  }
  &-leave-active {
    transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
  }
  &-enter-from,
  &-leave-to {
    transform: rotate(360deg);
    opacity: 0;
  }
}

.drop-down {
  display: flex;
  flex-direction: column;
}

// .header-top.main-container {
//   display: flex;
//   justify-content: center;
//   align-items: center;
//   position: fixed;
//   width: 100%;
//   top: 0px;
//   left: 50%;
//   transform: translateX(-50%);
//   z-index: 100;
//   background: #f57e20;
//   padding: 0px;
//   box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.7);
// }

.main-btn {
  font-size: 14px;
}

html,
body {
  margin: 0;
  padding: 0;
  overflow-x: hidden;
  height: 100%;
  width: 100%;
  background: white;
  font-family: 'Montserrat';
  scroll-behavior: smooth;
}

.main-container {
  max-width: 1170px;
  margin: 0 auto;
  h2 {
    font-size: 40px;
    font-weight: normal;
    text-align: center;
    position: relative;
  }
  h2:after {
    content: '';
    height: 3px;
    background: red;
    width: 87px;
    position: absolute;
    bottom: -10px;
    left: 0;
    right: 0;
    margin: auto;
  }
  p {
    line-height: 1.5;
    font-size: 16px;
  }
}

.main-btn {
  padding: 5px 15px;
  color: white;
  font-size: 16px;
  line-height: 1.5;
  border-radius: 30px;
  background-color: #3d1819;
  text-transform: uppercase;
  cursor: pointer;
  border: none;
}

button {
  font-family: 'Montserrat';
}

.button {
  padding: 5px 15px;
  border: 1px solid #dcdfe6;
  background-color: white;
  color: #263238;
  border-radius: 15px;
  font-size: 18px;
  margin: 5px;
  display: flex;
  align-items: center;
  white-space: nowrap;
  &:hover {
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    // filter: brightness(120%);
    cursor: pointer;
  }
}

.event-layout {
  max-width: 1200px;
  width: 100%;
  padding: 0 15px;
  margin: 0 auto;
  box-sizing: border-box;
}

@media screen and (max-width: 1024px) {
  .header-bottom-left {
    display: none;
  }
}

@media screen and (max-width: 768px) {
  .logo {
    display: none;
  }

  .ul-menu {
    padding-left: 0px;
  }

  .header-top.main-container {
    padding: 0;
  }

  .menu {
    width: 200px;
    justify-content: left;
  }

  .main-btn {
    font-size: 14px;
  }

  .menu-container {
    display: flex;
    justify-content: left;
    align-items: center;
  }

  .burger {
    display: block;
  }
}
</style>
