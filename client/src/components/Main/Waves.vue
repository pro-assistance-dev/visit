<template>
  <div class="wave-line" :class="{ top, bottom }">
    <div class="wave-line-1" :style="`background-position-x: ${wave1pos * 1}px;`"></div>
    <div class="wave-line-2" :style="`background-position-x: ${wave2pos * 0.7}px;`"></div>
    <div class="wave-line-3" :style="`background-position-x: ${wave3pos * 0.2}px;`"></div>
  </div>
</template>

<script lang="ts">
export default defineComponent({
  name: 'Waves',
  props: {
    top: {
      type: Boolean,
      default: false,
    },
    bottom: {
      type: Boolean,
      default: false,
    },
  },
  setup() {
    let scrollOffset = 0;
    let previousOffset = 0;
    let rememberedOffset: number | null = null;
    const wave1pos: ComputedRef<number> = computed(() => scrollOffset + 0);
    const wave2pos: ComputedRef<number> = computed(() => scrollOffset + 190);
    const wave3pos: ComputedRef<number> = computed(() => scrollOffset + 980);

    const handleScroll = () => {
      if (scrollOffset > previousOffset && rememberedOffset != null) {
        rememberedOffset = null;
      }
      previousOffset = scrollOffset;
      scrollOffset = window.scrollY;
    };

    onMounted(() => {
      window.addEventListener('scroll', handleScroll);
    });

    onUnmounted(() => {
      window.removeEventListener('scroll', handleScroll);
    });

    return {
      wave1pos,
      wave2pos,
      wave3pos,
    };
  },
});
</script>

<style lang="scss" scoped>
.wave-line.top {
  top: -120px;
}
.wave-line.bottom {
  bottom: -128px;
  transform: rotate(180deg) scaleX(-1);
}
.wave-line {
  position: absolute;
  left: 0;
  z-index: 20;
  width: 100%;
}
.wave-line-1,
.wave-line-2,
.wave-line-3 {
  background-image: url(@/assets/img/promo/wave.png);
  position: absolute;
  height: 130px;
  min-width: 100%;
}
.wave-line-2 {
  opacity: 0.2;
  top: -25px;
}
.wave-line-3 {
  opacity: 0.1;
  top: -50px;
}
</style>

