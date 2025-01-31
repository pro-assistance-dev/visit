<template>
  <div
    :key="head"
    class="admin-header-bottom"
    :style="{
      boxShadow: shadow ? '5px 0 5px 3px rgba(0,0,0,0.3)' : 'none',
    }"
  >
    <PFlex justify-content="space-between" width="100%" margin="0 20px 0 40px">
      <PFlex>
        <PButton v-if="PF.H.Router.Id() || PF.H.Router.Slug()" type="primary" width="50px" margin="0 20px 0 0" @click="PF.H.Router.Back()">
          <IconArrowLeft size="20px" />
        </PButton>
        <PString :string="head.title" font-size="26px" font-weight="bold" />
      </PFlex>
      <div class="button-group">
        <div v-for="item in head.buttons" :key="item.text" class="flex-item">
          <PButton
            v-if="item.action && item.condition"
            :key="item.condition"
            :type="item.type"
            :text="item.text"
            padding="0 20px"
            @click="action(item.action)"
          />
        </div>
      </div>
    </PFlex>
  </div>
</template>

<script lang="ts" setup>
defineProps({
  shadow: { type: Boolean, default: true },
});
const buttonClicked: Ref<boolean> = ref(false);
const head = PF.AdminUI.Head;
const action = (f: CallableFunction) => {
  buttonClicked.value = true;
  f();
  buttonClicked.value = false;
};
</script>

<style lang="scss" scoped>
.admin-header-bottom {
  position: relative;
  z-index: 999;
  box-sizing: border-box;
  background: var(--admin-header-background-color);
  display: flex;
  align-items: center;
  padding: 10px;
  margin-bottom: 10px;
  color: var(--admin-header-font-color);
  height: var(--admin-header-height, 60px);
}

h4 {
  font-weight: normal;
  font-size: 18px;
}
.flex-between {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  margin: 0 20px 0 40px;
}
.button-group {
  display: flex;
  justify-content: right;
  align-items: center;
}

.flex-item {
  margin: 0 0 0 10px;
}

.flex-line {
  box-sizing: border-box;
  display: flex;
  justify-content: left;
  align-items: center;
}
</style>
