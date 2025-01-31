<template>
  <div v-for="n in 20 - 9">
    <div
      class="container"
      :style="{
        height: timeStep * minutes[division].lenght + 'px',
      }"
    >
      <div
        v-for="min in minutes[division]"
        class="five-min"
        :style="{
          height: timeStep - 1 + 'px',
        }"
      >
        <div class="five-min-title">
          {{ n + 8 + `:${min}` }}
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
const props = defineProps({
  division: {
    type: String,
    default: '5',
  },
  timeStep: {
    type: Number,
    default: 50,
  },
});
const height = computed(() => {});
const minutes = {
  '5': ['00', '05', '10', '15', '20', '25', '30', '35', '40', '45', '50', '55'],
  '15': ['00', '15', '30', '45'],
  '30': ['00', '30'],
  '60': ['00'],
};
</script>

<style lang="scss" scoped>
.container {
  box-sizing: border-box;
  width: 100%;
}

.container:first-child {
  margin-top: 10px;
}

.five-min {
  position: relative;
  border-top: 1px solid #c4c4c4;
  display: flex;
  justify-content: left;
  align-items: start;
  font-size: 14px;
  color: #343e5c;
}

.five-min-title {
  background: #dff2f8;
  font-size: 14px;
  color: #343e5c;
  margin-top: -8px;
  padding: 0 5px;
}
</style>
