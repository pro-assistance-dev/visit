<template>
  <div
    class="admin-list-wrapper"
    :style="{
      height: `calc(100% - ${outTopMargin})`,
    }"
  >
    <div class="admin-list-wrapper-content">
      <slot />
    </div>
    <div v-if="store" class="pagination">
      <PPagination v-if="pagination" :store="store" />
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps({
  pagination: {
    type: Boolean,
    default: false,
  },
  store: {
    type: Object as PropType<IStorePaginator>,
    default: undefined,
  },
  outTopMargin: {
    type: String,
    required: false,
    default: '70px',
  },
});
</script>

<style lang="scss" scoped>
.admin-list-wrapper {
  position: relative;
  display: flex;
  flex-direction: column;
  overflow: hidden;
  padding: 0;
  box-sizing: border-box;
  font-family: var(--base-font);
  font-size: var(--font-size-default);

  &-content {
    overflow: hidden;
    overflow-y: auto;
    box-sizing: border-box;
    height: 100%;
    margin: 0 10px;
  }
}

.pagination {
  width: 100%;
  position: absolute;
  left: 0px;
  bottom: 0px;
  display: flex;
  justify-content: center;
}
</style>
