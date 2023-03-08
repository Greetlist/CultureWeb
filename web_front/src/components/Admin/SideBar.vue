<template>
  <div class="side-bar">
    <el-menu router class="menu-box">
      <el-submenu
        :index="' ' + item1.order"
        v-for="(item1, index) in adminMenuList"
        :key="index"
      >
        <template slot="title">
          <i :class="item1.headIcon"></i>
          <span>{{ item1.name }}</span>
        </template>
        <el-menu-item
          class="menuItem"
          @click="clickMenuItem"
          v-for="(item2, index) in item1.children"
          :key="index"
          :index="item2.path"
        >
          <span>{{ item2.name }}</span>
        </el-menu-item>
      </el-submenu>
    </el-menu>
  </div>
</template>
<script>
import { adminMenuList } from "@constant/index";

export default {
  name: "SideBar",
  data() {
    return {
      adminMenuList,
    };
  },
  methods: {
    clickTitle() {},
    clickMenuItem() {},
  },
};
</script>

<style lang="scss" scoped>
.menu-box {
  height: calc(100vh - 45px);
  text-align: left;
}
</style>
