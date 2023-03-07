<template>
  <el-container>
    <el-header style="padding:0">
      <NavBar/>
    </el-header>
    <el-container>
      <el-aside width="10%">
        <SideBar/>
      </el-aside>
      <el-main>
        <el-scrollbar class='page-component__scroll'>
          <router-view></router-view>
        </el-scrollbar>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import NavBar from '@/components/Admin/NavBar.vue'
import SideBar from '@/components/Admin/SideBar.vue'
export default {
  name: 'AdminHome',
  components: {
    NavBar,
    SideBar
  }
}
</script>
