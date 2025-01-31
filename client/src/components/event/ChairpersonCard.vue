<template>
  <div class="chairperson-card">
    <div class="img-container">
      <img :src="chairperson.photo.getImageUrl()" :alt="chairperson.photo.originalName" />
    </div>
    <h4>{{ chairperson.name }}</h4>
    <div class="regalis">{{ chairperson.regalias }}</div>
  </div>
</template>

<script lang="ts">
import Chairperson from '@/classes/Chairperson';

defineProps({
  chairperson: {
    type: Object as PropType<Chairperson>,
    required: true,
  },
});
</script>

<style lang="scss" scoped>
.chairperson-card {
  box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;
  border-radius: 20px;
  width: 340px;
  height: inherit;
  padding: 30px;
  font-size: 18px;
  word-break: break-word;
  background-color: white;
  z-index: 1;
  box-sizing: border-box;
}
.regalis {
  line-height: 1.5;
}
h4 {
  font-size: 24px;
  font-weight: bold;
  text-align: start;
  margin: 0 0 8px 0;
}
.img-container {
  width: 100%;
  text-align: center;
  margin-bottom: 20px;
}
img {
  border-radius: 50%;
  height: 220px;
}
@media screen and (max-width: 768px) {
  h4 {
    font-size: 20px;
  }
  .chairperson-card {
    font-size: 16px;
  }
  img {
    height: 180px;
  }
  .chairperson-card {
    width: 280px;
  }
}
@media screen and (max-width: 480px) {
  img {
    height: 130px;
  }
  h4 {
    font-size: 18px;
  }
  .chairperson-card {
    font-size: 14px;
  }
  .chairperson-card {
    width: 200px;
  }
}
</style>

