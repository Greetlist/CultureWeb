<template>
  <div class="login-view">
    <h1>XXX文化馆</h1>
    <el-form ref="loginForm" :model="userInfo" class="login-form">
      <el-form-item
        prop="name"
        autocomplete="off"
        :rules="{
          required: true,
          message: '请输入用户名',
          trigger: 'blur',
        }"
      >
        <el-input v-model="userInfo.name" placeholder="用户名或邮箱"></el-input>
      </el-form-item>
      <el-form-item
        prop="passwd"
        :rules="{
          required: true,
          message: '请输入密码',
          trigger: 'blur',
        }"
      >
        <el-input
          type="password"
          v-model="userInfo.passwd"
          placeholder="密码"
          show-password
        ></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="onLogin('loginForm')" class="login-button"
      >登录</el-button
    >
  </div>
</template>

<script>
import { userApi } from "@services/index";

export default {
  name: "LoginView",
  data() {
    return {
      userInfo: {
        name: "",
        passwd: "",
      },
    };
  },
  methods: {
    onLogin(formName) {
      this.$refs[formName].validate(async (valid) => {
        let loading = this.$loading({
          lock: true,
          text: "登录中",
          duration: 60000,
        });
        if (valid) {
          const { userInfo } = this;
          const { status, data } = await userApi.login({
            account: userInfo?.name ?? "",
            passwd: userInfo?.passwd ?? "",
          });
          loading.close();
          // TODO修改以下判断条件
          if (status != 200 || data?.login_succ) {
            this.$message.error("登录失败！");
            return;
          } else {
            this.$message({
              message: "登录成功！",
              type: "success",
              duration: 1000,
            });
            setTimeout(() => {
              this.$router.replace("/admin");
            }, 1000);
          }
        } else {
          return;
        }
      });
    },
  },
};
</script>
<style lang="scss" scoped>
.login-view {
  width: 40%;
  margin: 180px auto 0;
  min-width: 100px;
  .login-button {
    width: 100%;
  }
  /deep/ {
    .el-button--primary {
      border: none;
      color: #fff;
    }
  }
}
</style>
