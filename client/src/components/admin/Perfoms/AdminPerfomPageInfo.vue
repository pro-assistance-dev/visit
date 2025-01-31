<template>
  <LeftRightContainer left-width="280px">
    <template #right>
      <PerfomConstructor />
    </template>
  </LeftRightContainer>
</template>

<script lang="ts" setup>
const update = async () => {
  await PerfomsStore.Update();
};
</script>

<style lang="scss" scoped>
.light-title {
  color: #a1a8bd;
  display: flex;
  align-items: center;
}

.upper {
  text-transform: uppercase;
}

.flex-center {
  display: flex;
  align-items: center;
}

.contact-icon {
  font-size: 25px;
  margin-right: 5px;
}

.el-tag {
  font-size: 14px;
}

.register-tag {
  &:hover {
    cursor: pointer;
    border-width: 2px;
  }
}

:deep(.el-upload--picture-card) {
  width: 150px;
  font-size: 50px;
}

:deep(.el-upload--picture-card i) {
  font-size: 50px;
  color: #00b5a4;
  padding: 0 114px;
}

.left-title {
  font-size: 14px;
  font-weight: bold;
  color: #343d5c;
  margin-top: 28px;
  margin-left: 10px;
  text-transform: uppercase;
}

.left-info {
  font-size: 12px;
  color: #343d5c;
  margin-top: 3px;
  margin-left: 10px;
}
</style>
