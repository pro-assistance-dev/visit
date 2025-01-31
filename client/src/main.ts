import './assets/styles/main.scss';
import 'video.js/dist/video-js.css';
import VueVideoPlayer from '@videojs-player/vue';
import { createApp } from 'vue';

import App from './App.vue';
import router from './router';

const app = createApp(App);

app.use(router);
app.use(VueVideoPlayer);

app.directive('click-outside', {
  mounted(el, binding) {
    el.clickOutsideEvent = function (event: Event) {
      if (!(el === event.target || el.contains(event.target))) {
        binding.value(event, el);
      }
    };
    document.body.addEventListener('click', el.clickOutsideEvent);
  },
  unmounted(el) {
    document.body.removeEventListener('click', el.clickOutsideEvent);
  },
});

import setup from './setup/index';
setup(app, router);

router.isReady().then(() => {
  app.mount('#app');
});
