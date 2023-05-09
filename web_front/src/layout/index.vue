<template>
  <el-container>
    <el-header height="45px">
      <Navbar></Navbar>
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
  name: "LayoutMain",
  components: {
    Navbar,
    Sidebar,
  },
};
</script>
<style lang="scss" scoped>
/deep/ {
  .el-header {
    padding: 0;
  }
  .el-aside {
    a {
      text-decoration: none;
    }
  }
  .el-menu {
    border-right: none;
  }
  .el-scrollbar {
    height: calc(100vh - 45px);
    border-right: solid 1px #e6e6e6;
  }
  .el-scrollbar__bar.is-vertical > div {
    display: none;
  }
  .el-scrollbar__wrap {
    overflow-x: hidden;
  }
  .el-main {
    height: calc(100vh - 45px);
  }
}
</style>
