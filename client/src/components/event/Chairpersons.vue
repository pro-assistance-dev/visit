<template>
  <div class="chairman" v-if="chairpersons.length">
    <div class="bg"></div>
    <div class="event-layout">
      <div class="chairman-container">
        <h1>Председатель конференции</h1>
        <div class="chairman-card-row">
          <chairperson-card v-for="(chairperson, i) in chairpersons" :key="i" :chairperson="chairperson" />
        </div>
        <el-carousel
          v-touch:swipe="(direction) => $carouselSwipe(direction, 'chairpersonForTabs')"
          ref="chairpersonForTabs"
          class="chairman-carousel"
          indicator-position="outside"
          arrow="always"
          :interval="5000"
          height="580px"
          style="width: 100%"
        >
          <el-carousel-item v-for="(chairperson, i) in chairpersons" :key="i">
            <chairperson-card :chairperson="chairperson" />
          </el-carousel-item>
        </el-carousel>
        <el-carousel
          v-touch:swipe="(direction) => $carouselSwipe(direction, 'chairpersonForMobile')"
          ref="chairpersonForMobile"
          class="chairman-tablet-carousel"
          indicator-position="outside"
          arrow="always"
          :interval="4000"
          height="550px"
          style="width: 100%"
        >
          <el-carousel-item v-for="(chairperson, i) in chairpersons" :key="i">
            <chairperson-card :chairperson="chairperson" />
          </el-carousel-item>
        </el-carousel>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Chairperson from '@/classes/Chairperson';
const chairpersons: Chairperson[] = Store.Events.Item.chairpersons;
</script>

<style scoped lang="scss">
.el-carousel__item {
  display: flex;
  justify-content: space-evenly;
  align-items: center;
}
.chairman {
  position: relative;
}
.bg {
  height: 270px;
  top: 230px;
  border: none;
  background-color: #e1f1fc;
  background-size: cover;
  width: 100%;
  position: absolute;
  z-index: 0;
}
.chairman-card-row {
  display: flex;
  justify-content: space-evenly;
  width: 100%;
  height: auto;
}
h1 {
  font-size: 33px;
  font-weight: bold;
  text-align: center;
  margin: 50px 0;
}
.chairman-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
.chairman-carousel {
  display: none;
}
.chairman-tablet-carousel {
  display: none;
}
@media screen and (max-width: 1200px) {
  .chairman-card-row {
    display: none;
  }
  .chairman-carousel {
    display: unset;
  }
}
@media screen and (max-width: 1024px) {
  h1 {
    font-size: 27px;
  }
}
@media screen and (max-width: 768px) {
  .chairman-carousel {
    display: none;
  }
  .chairman-tablet-carousel {
    display: unset;
  }
  h1 {
    font-size: 24px;
  }
}
@media screen and (max-width: 480px) {
  h1 {
    font-size: 20px;
  }
}
</style>
