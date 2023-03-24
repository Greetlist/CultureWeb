<template>
  <el-container>
    <el-header height="45px">
      <Navbar title="XXX文化馆"></Navbar>
    </el-header>
    <el-container>
      <el-aside width="220px">
        <Sidebar></Sidebar>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import { Navbar, Sidebar } from "@c/index";
export default {
  name: "AdminHome",
  components: {
    Navbar,
    Sidebar,
  },
};
</script>
<style lang="scss" scoped>
/deep/ .el-header {
  padding: 0;
}
</style>
