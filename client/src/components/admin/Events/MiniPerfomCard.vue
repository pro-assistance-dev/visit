<template>
  <div class="speaker">
    <div class="speaker-title">
      {{ perfom.name }}
    </div>
    <svg class="icon-speaker-edit">
      <use xlink:href="#speaker_edit"></use>
    </svg>
    <!-- <div class="info">info</div> -->
  </div>

  <SpeakerEdit />
</template>

<script setup lang="ts">
import Perfom from '@/classes/Perfom';
import SpeakerEdit from '@/assets/svg/SpeakerEdit.svg';

defineProps({
  perfom: {
    type: Object as PropType<Perfom>,
    required: true,
  },
});

// const scheduler: ComputedRef<Scheduler> = computed(() => Provider.store.getters['schedules/scheduler']);

const emits = defineEmits(['edit']);

const edit = async () => {
  emits('edit');
};
</script>

<style lang="scss" scoped>
.hidden {
  display: none;
}

.speaker {
  position: relative;
  box-sizing: border-box;
  margin-top: 10px;
  background: #e3e3e3;
  border: 1px solid #343e5c;
  border-radius: 3px;
  z-index: 5;
  width: calc(100% - 4px);
  overflow: hidden;
}

.speaker-title {
  width: 100%;
  margin-right: 45px;
  padding-left: 10px;
  box-sizing: border-box;
  color: #343e5c;
  min-height: 45px;
  border-bottom: 1px solid #343e5c;
  display: flex;
  justify-content: left;
  align-items: center;
}

.icon-speaker-edit {
  position: absolute;
  top: 5px;
  right: 5px;
  width: 32px;
  height: 32px;
  stroke: #343e5c;
  cursor: pointer;
  transition: 0.3s;
}

.icon-speaker-edit:hover {
  stroke: #343e5c;
  transform: scale(1.2);
}

// .speaker-edit {
//   position: absolute;
//   top: 20px;
//   right: 20px;
//   border: 1px solid #343e5c;
//   border-radius: 5px;
//   width: 240px;
//   background: #ffffff;
// }

.info {
  display: flex;
  justify-content: left;
  align-items: center;
  padding: 0 10px;
  color: #343e5c;
  background: #ffffff;
  height: 100%;
}
</style>
