<template>
  <div class="carousel" :style="{ height: height, maxWidth: maxWidth }">
    <div class="carousel-track" :style="{ transform: `translateX(-${currentIndex * 100}%)`, height: height }">
      <div v-for="banner in banners" :key="banner.id" class="carousel-slide">
        <img
          :style="{ height: `calc(${height} - 20px)`, width: 'auto' }"
          :src="banner.bannerPath"
          alt="banner"
          @click="openLink(banner.bannerLink)"
        />
      </div>
    </div>

    <div v-if="showControl" class="carousel-control prev">
      <PButton skin="text" type="primary" @click="prevSlide"><ArrowLeft size="20px" /></PButton>
    </div>
    <div v-if="showControl" class="carousel-control next">
      <PButton skin="text" type="primary" @click="nextSlide"><ArrowRight size="20px" /></PButton>
    </div>

    <div v-if="showControl" class="carousel-indicators">
      <span
        v-for="(banner, index) in banners"
        :key="'indicator-' + banner.id"
        :class="{ active: index === currentIndex }"
        @click="goToSlide(index)"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import IBannerArray from '@/interfaces/IBannerArray';

const props = defineProps({
  banners: {
    type: Array as () => IBannerArray[],
    required: true,
  },
  showControl: {
    type: Boolean,
    default: true,
  },
  animationTime: {
    type: Number,
    default: 10000,
  },

  height: {
    type: String,
    default: '100px',
  },

  maxWidth: {
    type: String,
    default: '100%',
  },
});

const currentIndex = ref(0);
let intervalId: number | null = null;

const nextSlide = () => {
  currentIndex.value = (currentIndex.value + 1) % props.banners.length;
};

const startAutoSlide = () => {
  intervalId = window.setInterval(nextSlide, props.animationTime);
};

const stopAutoSlide = () => {
  if (intervalId !== null) {
    clearInterval(intervalId);
    intervalId = null;
  }
};

onMounted(() => {
  startAutoSlide();
});

onUnmounted(() => {
  stopAutoSlide();
});

const goToSlide = (index: number) => {
  currentIndex.value = index;
  stopAutoSlide();
  startAutoSlide();
};

const prevSlide = () => {
  currentIndex.value = (currentIndex.value - 1 + props.banners.length) % props.banners.length;
};

const openLink = (link: string) => {
  if (link.startsWith('http')) {
    window.open(link, '_blank');
  } else {
    console.warn('Invalid URL:', link);
  }
};
</script>

<style scoped>
.carousel {
  position: relative;
  overflow: hidden;
}

.carousel-track {
  display: flex;
  transition: transform 0.5s ease-in-out;
}

.carousel-slide {
  min-width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  height: auto;
  overflow: hidden;
}

img {
  width: auto;
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
  cursor: pointer;
}

.carousel-slide img {
  width: 100%;
  height: auto;
  cursor: pointer;
}

.carousel-control {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(0, 0, 0, 0.3);
  border: none;
  cursor: pointer;
  z-index: 10;
  width: 30px;
  height: 30px;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 30px;
}

.carousel-control:hover {
  background: rgba(0, 0, 0, 0.2);
}

.carousel-control.prev {
  left: 10px;
}

.carousel-control.next {
  right: 10px;
}

.carousel-indicators {
  position: absolute;
  bottom: 10px;
  display: flex;
  justify-content: center;
  gap: 10px;
  width: 100%;
}

.carousel-indicators span {
  display: block;
  width: 10px;
  height: 10px;
  background: gray;
  border-radius: 50%;
  cursor: pointer;
}

.carousel-indicators span.active {
  background: #f8b800;
}
</style>
