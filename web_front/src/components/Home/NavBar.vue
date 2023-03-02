<template>
  <div class="nav-bar">
    <div class="nav-bar-box">
      <span>{{ title }}</span>
    </div>
    <div class="empty-placeholder"></div>
  </div>
</template>
<script>
export default {
  name: "NavBar",
  props: {
    title: {
      type: String,
      default: "默认标题",
    },
  },
};
</script>
<style lang="scss" scoped>
.nav-bar {
  line-height: 45px;
  padding: 0 15px;
  text-align: left;
  color: #fff;
  background-color: #606266;
  .nav-bar-box {
    position: fixed;
    width: 100%;
    height: 45px;
    top: 0;
  }
  .empty-placeholder {
    height: 45px;
  }
}
</style>
