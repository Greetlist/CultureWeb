<template>
  <div class="nav-bar">
    <el-row>
      <el-col :span="18">
        <div>{{ webBasicInfo.web_name }}</div>
      </el-col>
      <el-col :span="6" class="user-info">
        <div>
          <span class="welcome-msg">欢迎您，管理员</span>
          <span class="logout-btn" @click="signout">退出</span>
        </div>
      </el-col>
    </el-row>
  </div>
</template>
<script>
import { mapState } from "vuex"
export default {
  name: "NavBar",
  methods: {
    signout() {},
  },
  computed: mapState([
    'webBasicInfo'
  ])
};
</script>
<style lang="scss" scoped>
.nav-bar {
  line-height: 45px;
  padding: 0 15px;
  text-align: left;
  color: #fff;
  background-color: #606266;
  .user-info {
    text-align: right;
    .welcome-msg {
      margin: 0 10px;
    }
  }
}
</style>
