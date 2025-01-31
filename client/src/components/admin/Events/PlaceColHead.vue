<template>
  <div class="place-title" v-click-outside.prevent="outsideClick">
    {{ place.name }}

    <!-- <svg class="icon-plus" @click="openWindow = true"> -->
    <!--   <use xlink:href="#plus"></use> -->
    <!-- </svg> -->
    <div
      class="add-window"
      :style="{
        display: openWindow ? 'block' : 'none',
      }"
    >
      <Button @click="addSession" text="Добавить сессию" width="160px" height="32px" />
    </div>
  </div>
  <Plus />
  <Move />
</template>

<script setup lang="ts">
import Plus from '@/assets/svg/Plus.svg';
import Place from '@/classes/Place';
import Schedule from '@/classes/Schedule';

const props = defineProps({
  schedule: {
    type: Object as PropType<Schedule>,
    required: true,
  },
  place: {
    type: Object as PropType<Place>,
    required: true,
  },
});
const emits = defineEmits('addSession');
const outsideClick = () => {
  openWindow.value = false;
};
const openWindow: Ref<boolean> = ref(false);

const addSession = async () => {
  openWindow.value = false;
  emits('addSession');
};
</script>

<style lang="scss" scoped>
.hidden {
  display: none;
}
.place {
  width: 100%;
  max-width: 400px;
  background: #eefff5;
  border: 1px solid #343e5c;
  margin: 0px 5px;
  box-sizing: border-box;
  border-radius: 5px;
  // overflow: hidden;
}

.place-title {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 42px;
  background: #c9ffdf;
  font-size: 16px;
  font-weight: bold;
  color: #343e5c;
  padding: 0 10px;
  -webkit-user-select: none; /* Safari */
  -ms-user-select: none; /* IE 10 and IE 11 */
  user-select: none; /* Standard syntax */
}

.icon-plus {
  width: 28px;
  height: 28px;
  fill: #00b5a4;
  cursor: pointer;
}
.icon-plus:hover {
  transform: scale(1.2);
  transition: 0.3s;
}

.add-window {
  position: absolute;
  top: 10px;
  right: 50px;
  border: 1px solid #343e5c;
  border-radius: 5px;
  background: #ffffff;
  z-index: 4;
}
</style>
