<template>
  <el-menu
    :default-active="1"
    class="el-menu-demo"
    mode="horizontal"
    @select="handleSelect"
    background-color="#545c64"
    text-color="#fff"
    active-text-color="#ffd04b">
    <el-menu-item index="1">用户首页</el-menu-item>
  </el-menu>
</template>


<script>
export default {
  name: 'NavBar'
}
</script>
