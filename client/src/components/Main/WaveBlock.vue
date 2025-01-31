<template>
  <div class="wave-block">
    <waves top></waves>
    <div class="wave-container main-container">
      <h2>{{ title }}</h2>
      <div class="wave-container-content">
        <slot />
      </div>
    </div>
    <waves bottom v-if="bottomWaves"></waves>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Waves from '@/components/Main/Waves.vue';

export default defineComponent({
  name: 'WaveBlock',
  components: { Waves },
  props: {
    title: {
      type: String,
      required: true,
    },
    bottomWaves: {
      type: Boolean,
      default: true,
    },
  },
});
</script>

<style lang="scss" scoped>
.wave-block {
  position: relative;
}
.wave-container {
  padding: 48px 15px;
  &-content {
    font-size: 16px;
    margin-top: 48px;
    line-height: 1.5;
  }
}

@media screen and (max-width: 768px) {
  .wave-container {
    padding: 15px;
    &-content {
      font-size: 16px;
      margin-top: 48px;
      line-height: 1.5;
    }
  }
}
</style>