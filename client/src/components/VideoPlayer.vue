<template>
  <div class="video-container">
    <VideoPlayer
      @mounted="playerMounted"
      :src="src"
      controls
      muted
      fluid
      ref="videoPlayer"
      class="video-js vjs-matrix vjs-big-play-centered vjs-layout-huge"
    >
    </VideoPlayer>
  </div>
</template>

<script lang="ts">
import { VideoPlayer, VideoPlayerEvents } from '@videojs-player/vue';
import 'video.js/dist/video-js.css';
// import p from 'videojs-contrib-quality-levels';
import Event from '@/classes/Event';

export default defineComponent({
  components: {
    VideoPlayer,
  },
  props: {
    event: {
      type: Object as PropType<Event>,
      required: true,
    },
  },
  setup(props) {
    const src = `${import.meta.env.VITE_APP_HLS_SERVER}:${import.meta.env.VITE_APP_HLS_PORT}/hls/${props.event.id}.m3u8`;
    // const src =
    //   'https://streaming.disk.yandex.net/hls/U2FsdGVkX1-_ouVJIf5q620-bCLuW_FZTzFGod0iUFyG-OwsctUL77QCi71k905a7nbFo_QWyU4NIWUCkXBhKeDr-5NLWzWpsfPvi9Kk8DdgUQQGKTPUnjMbyNmp9thflcutcgj5xOrIjVKrdWvJNG8rlspo9X56HfzywLKHlFGIcaL6KdcQI32V5b0_4A3ZTu_yOVY-wYET2CdaxQohZq1AJvUK6pV5Nh7UTPoZbqtYiiIz9nbD_POQJ65slzvQXAIkZ1aPnl2s8Nf6tFP3lgG2iklwKVM-4BYBO_ANA0QWhSHexuwrFRGOn3XXKIdos5dWLMcFYoIQUeaqbif5lk1RxqGb_KCvDg0Fhdtj61A/60b9dae134770/cd4486722e69d6346a846ac904e1823888266ce918ae83c413eaffb932066787/480p/playlist.m3u8?from=disk&vsid=4e891be983988dcea9e32517edb730eb0efd4e2e85dexWEBx1149x1701535322&source_index=0&session_data=1&preview=1&t=1701535324834';
    // const plugins = { 'videojs-contrib-quality-levels': p };
    const playerMounted = (e: any) => {
      e.player.play();
    };
    return {
      playerMounted,
      src,
    };
  },
});
</script>

<style>
.video-container {
  height: 100%;
  overflow: hidden;
  display: flex;
  align-items: center;
}

.vjs-matrix.video-js {
  color: #f57e20;
  /* border-radius: 20px; */
  border-style: none;
  overflow: hidden;
  background-color: #3d1819;
  /*put other stuff you need here*/
}

/* Change the border of the big play button. */
.vjs-matrix .vjs-big-play-button {
  border-color: #f57e20;
  /*put other stuff you need here*/
}

/* Change the color of various "bars". */
.vjs-matrix .vjs-volume-level,
.vjs-matrix .vjs-play-progress,
.vjs-matrix .vjs-slider-bar {
  background: #f57e20;
  /*put other stuff you need here*/
}
video[poster] {
  object-fit: cover;
}
.vjs-poster {
  background-size: cover;
  background-position: inherit;
}
</style>
