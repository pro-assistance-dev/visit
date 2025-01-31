<template>
  <div style="min-height: 100vh">
    <div class="bg">
      <div class="admin-main-container">
        <PLeftSlider :open="true" background="var(--left-slider-background)" width-opened="230px" width-closed="20px">
          <AdminSideMenu :border="false" />
        </PLeftSlider>
        <div class="admin-container">
          <AdminHeader style="position: sticky; z-index: 2" :shadow="false" />
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup></script>

<style lang="scss" scoped>
.admin-main-container {
  display: flex;
  justify-content: left;
  height: 100vh;
  box-sizing: border-box;
}

.admin-container {
  width: 100%;
  box-sizing: border-box;
  height: 100%;
  overflow: hidden;
  overflow-y: auto;
}

.bg {
  height: 100%;
  box-sizing: border-box;
  width: 100%;
  background: var(--background-color);
}
</style>
