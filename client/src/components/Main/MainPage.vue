<template>
  <WaveBlock id="about" title="О нас">
    <p>
      Организация создана в целях предоставления услуг и поддержки социально ориентированных инициатив в сфере развития здравоохранения,
      содействия медицинскому профессиональному сообществу в решении вопросов, касающихся медицинской науки, просвещения, создания условий
      для реализации проектов в сфере здравоохранения и поддержки пациентов
    </p>
  </WaveBlock>
  <div id="projects" class="colored">
    <div class="main-container">
      <h2>Наши проекты и услуги</h2>
      <p style="max-width: 100%; margin: 0 20px; padding-top: 20px">
        Специалисты, которые привлекаются к участию в проектах, задействованы в самых различных профессиональных сферах: от бизнес-аналитики
        и системы здравоохранения до телевидения и журналистики
      </p>
      <ul class="services">
        <li v-for="item in services" :key="item.name" class="services-item">
          <div>{{ item.name }}</div>
        </li>
      </ul>
    </div>
  </div>
  <WaveBlock v-if="eventsLoaded && events.length > 0" id="lib" title="Библиотека">
    <div class="image-container">
      <CarouselImages v-if="eventsLoaded" :events="events" />
    </div>
  </WaveBlock>
  <div class="image-container"></div>
  <WaveBlock id="contacts" title="Контакты" :bottom-waves="false">
    <div class="contacts-content">
      <div class="contacts-content-left">
        <div class="contacts-content-left-info">
          <div>
            <img :src="phoneImg" alt="..." />
          </div>
          <h4>+7 (916) 078 33 70</h4>
        </div>

        <div class="contacts-content-left-info">
          <div>
            <img :src="emailImg" alt="..." />
          </div>
          <h4>pro-assistance@mail.ru</h4>
        </div>

        <div class="contacts-content-left-links">
          <img src="@/assets/img/icons/facebook.png" alt="" />
          <img src="@/assets/img/icons/twitter.png" alt="" />
          <img src="@/assets/img/icons/instagram.png" alt="" />
        </div>
      </div>
      <div class="contacts-content-right">
        <input type="text" placeholder="Имя Фамилия" />
        <input type="text" placeholder="Email" />
        <input type="text" placeholder="Заголовок" />
        <input type="text" placeholder="Сообщение" />
        <div>
          <button class="main-btn">Отправить</button>
        </div>
      </div>
    </div>
  </WaveBlock>
</template>

<script lang="ts" setup>
import Event from '@/classes/Event';
const events: Event[] = Store.Events.Items;

const services = [
  { icon: 'video.png', name: 'Фото и видео поддержка' },
  { icon: 'edit-3.png', name: 'Дизайн, оформление и иллюстрации' },
  { icon: 'search.png', name: 'Сервисы по сбору и анализу данных' },
  { icon: 'monitor.png', name: 'Поддержка и разработка приложений' },
  { icon: 'map.png', name: 'Проведение и организация мероприятий' },
  { icon: 'printer.png', name: 'Брендирование продуктов и мероприятий' },
];

const eventsLoaded: Ref<boolean> = ref(false);
onBeforeMount(async () => {
  await Store.Events.GetAll();
  eventsLoaded.value = true;
});
</script>

<style lang="scss" scoped>
.colored {
  padding: 200px 0;
  background: linear-gradient(135deg, #f25a2a 18%, #f67f1e 41%, rgba(255, 222, 0, 0.51) 100%);
}

.image-container {
  padding: 60px 0;
}
.ellipse-icon {
  margin-right: 20px;
  height: 86px;
  background: #fff;
  border-radius: 100%;
  text-align: center;
  display: flex;
  flex-direction: column;
  justify-content: center;
  padding: 0 22px;
  img {
    width: 42px;
    vertical-align: middle;
    border-style: none;
    overflow-clip-margin: content-box;
    overflow: clip;
    box-sizing: border-box;
  }
}
.services {
  margin-top: 50px;
  column-count: 2;
  &-item {
    width: 60%;
    display: flex;
    align-items: center;
    margin: 0 auto;
    padding-top: 50px;
    font-size: 16px;
    font-weight: bold;
  }
}
.contacts {
  &-content {
    display: flex;
    margin: 48px 0;
  }
  &-content-left {
    width: 60%;

    &-info {
      display: flex;
      align-items: center;
      img {
        margin-right: 24px;
      }
    }

    &-links {
      margin: 16px 0 0 48px;
      img {
        margin: 10px;
      }
    }
  }
  &-content-right {
    width: 40%;
    display: flex;
    flex-direction: column;
    input {
      font-size: 16px;
      font-weight: 400;
      line-height: 1.5;
      color: #495057;
      padding: 6px 12px;
      border: none;
      background: transparent;
      border-bottom: 1px solid black;
      margin-bottom: 15px;
    }
    input:focus {
      outline: none !important;
      outline: 0;
      outline-color: initial;
      outline-style: initial;
      outline-width: 0px;
      box-shadow: 0 0 0 0.2rem rgb(0 123 255 / 25%);
    }
  }
}

@media screen and (max-width: 768px) {
  .services {
    margin-top: 20px;
    column-count: 1;
    padding: 0 20px;
    &-item {
      min-width: 280px;
      max-width: 500px;
      display: flex;
      align-items: center;
      margin: 0 auto;
      padding-top: 20px;
      font-size: 16px;
      font-weight: bold;
    }
  }

  .ul {
    display: block;
  }

  .contacts {
    &-content {
      display: block;
      margin: 20px 0;
    }
    &-content-left {
      min-width: 280px;
      max-width: 280px;
      margin: 0 auto;
    }
    &-content-right {
      min-width: 280px;
      max-width: 600px;
      margin: 0 auto;
    }
  }

  .image-container {
    padding: 60px 0;
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
</style>
