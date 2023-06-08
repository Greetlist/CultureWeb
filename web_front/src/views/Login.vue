<template>
  <div class="login-view">
    <h2 style="text-align: center">{{ webBasicInfo.web_name }}</h2>
    <el-form ref="loginForm" :model="userInfo" class="login-form">
      <el-form-item
        prop="account"
        autocomplete="off"
        :rules="{
          required: true,
          message: '请输入用户名',
          trigger: 'blur',
        }"
      >
        <el-input v-model="userInfo.account" placeholder="用户名或邮箱"></el-input>
      </el-form-item>
      <el-form-item
        prop="password"
        :rules="{
          required: true,
          message: '请输入密码',
          trigger: 'blur',
        }"
      >
        <el-input
          type="password"
          v-model="userInfo.password"
          placeholder="密码"
          show-password
        ></el-input>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="onLogin('loginForm')" class="login-button">
      登录
    </el-button>
    <el-button type="primary" @click="toHome()" class="login-button">
      Home
    </el-button>
  </div>
</template>

<script>
import { userNormalApi } from "@services/user/normal/";
import { mapState } from "vuex"

export default {
  name: "LoginView",
  data() {
    return {
      userInfo: {
        account: "",
        password: "",
      }
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
          const { status, data } = await userNormalApi.login({
            account: userInfo?.account ?? "",
            password: userInfo?.password ?? "",
          });
          loading.close();
          // TODO修改以下判断条件
          console.log(data)
          var res = data.request_result
          if (res.return_code != 0) {
            this.$message.error(res.error_msg);
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
    toHome() {
      this.$router.replace("/home");
    }
  },
  created() {
    var instance = this
    userNormalApi.getWebBasicInfo().then(function (res) {
      instance.$store.commit('setWebBasicInfo', res.data.web_basic_info)
    })
    userNormalApi.checkLogin().then(function (res) {
      var data = res.data
      if (data.request_result.return_code == 0 && data.is_login === true) {
        instance.$router.replace("/admin")
      }
    })
  },
  computed: mapState([
    'webBasicInfo'
  ])
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
}
</style>
