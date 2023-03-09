<template>
  <el-form ref="user_form" status-icon :rules="user_check_rules" :model="user_form" label-width="180px" size="medium">
    <el-form-item label="账户" prop="account">
      <el-input v-model="user_form.account" clearable></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="password">
      <el-input
        v-model="user_form.password"
        type="password"
        autocomplete="off"
        clearable
      ></el-input>
    </el-form-item>
    <el-form-item label="再次输入密码" prop="double_password">
      <el-input
        v-model="user_form.double_password"
        type="password"
        autocomplete="off"
        clearable
      ></el-input>
    </el-form-item>
    <el-form-item label="昵称">
      <el-input v-model="user_form.name" clearable></el-input>
    </el-form-item>
    <el-form-item label="手机">
      <el-input v-model="user_form.phone_number" clearable></el-input>
    </el-form-item>
    <el-form-item label="性别">
      <el-select v-model="user_form.sex" placeholder="性别">
        <el-option label="男" value="male"></el-option>
        <el-option label="女" value="female"></el-option>
      </el-select>
    </el-form-item>
    <el-form-item label="年龄">
      <el-input v-model="user_form.age" clearable></el-input>
    </el-form-item>
    <el-form-item label="家庭住址">
      <el-input v-model="user_form.address" clearable></el-input>
    </el-form-item>
    <el-form-item label="身份证号码">
      <el-input v-model="user_form.identify_id" clearable></el-input>
    </el-form-item>
    <el-form-item label="是否激活">
      <el-switch v-model="user_form.is_active"></el-switch>
    </el-form-item>
    <el-form-item label="管理员权限">
      <el-switch v-model="user_form.is_admin"></el-switch>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submit">创建</el-button>
      <el-button @click="reset">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script>

import { adminApi } from "@services/admin/"

export default {
  name: "AddUser",
  data: function () {
    return {
      user_form: {
        account: "",
        password: "",
        double_password: "",
        name: "",
        phone_number: "",
        sex: "",
        age: "",
        address: "",
        identify_id: "",
        is_active: "",
        is_admin: "",
      },
      user_check_rules: {
        account: [
          { required: true, message: "请输入账户", trigger: ["change", "blur"] },
          { validator: this.checkAccount, trigger: ["change", "blur"] },
        ],
        password: [
          { required: true, message: "请输入密码", trigger: ["change", "blur"] },
          { validator: this.checkPassword, trigger: ["change", "blur"] },
        ],
        double_password: [
          { required: true, message: "请再次输入密码", trigger: ["change", "blur"] },
          { validator: this.doubleCheckPassword, trigger: ["change", "blur"] },
        ],
      },
    };
  },
  methods: {
    checkAccount(rule, value, callback) {
      if (value === "") {
        return callback(new Error("账户不能为空"));
      }
      callback();
    },
    checkPassword(rule, value, callback) {
      if (value === "") {
        return callback(new Error("密码不能为空"));
      } else if (this.user_form.password !== "") {
        if (value.length < 6) {
          return callback(new Error("密码需大于6位"));
        }
      }
      callback();
    },
    doubleCheckPassword(rule, value, callback) {
      if (value === "") {
        return callback(new Error("请再次输入密码"));
      } else if (value !== this.user_form.password) {
        return callback(new Error("两次密码不同"));
      } else {
        callback();
      }
    },
    submit() {
      this.$refs["user_form"].validate((valid) => {
        if (valid) {
          adminApi.addUser(this.user_form)
          //adminApi.getTotalUserInfo()
        } else {
          return false;
        }
      });
    },
    reset() {
      this.user_form.account = "";
      this.user_form.password = "";
      this.user_form.double_password = "";
      this.user_form.name = "";
      this.user_form.phone_number = "";
      this.user_form.sex = "";
      this.user_form.age = "";
      this.user_form.address = "";
      this.user_form.identify_id = "";
      this.user_form.is_active = "";
      this.user_form.is_admin = "";
    },
  },
};
</script>

<style>
.el-input {
  width: 50%;
}
</style>
