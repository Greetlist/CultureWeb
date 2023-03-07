<template>
  <el-form ref="user_form" :model="user_form" label-width="180px" size="medium">
    <el-form-item label="账户">
      <el-input v-model="user_form.account" clearable></el-input>
    </el-form-item>
    <el-form-item label="密码">
      <el-input v-model="user_form.password" type="password" autocomplete="off" clearable></el-input>
    </el-form-item>
    <el-form-item label="再次输入密码">
      <el-input v-model="user_form.double_password" type="password" autocomplete="off" clearable></el-input>
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
export default {
  name: 'AddUser',
  data: function () {
    var checkAccount = (rule, value, callback) => {
      if (value === '') {
        return callback(new Error("账户不能为空"))
      }
      callback()
    }
    var checkPassword = (rule, value, callback) => {
      if (value === '') {
        return callback(new Error("密码不能为空"))
      } else if (this.user_form.password.value !== '') {
        if (value.length < 6) {
          return callback(new Error("密码需大于6位"))
        }
      }
      callback()
    }
    var doubleCheckPassword = (rule, value, callback) => {
      if (value === '') {
        return callback(new Error("请再次输入密码"))
      } else if (value !== this.user_form.password.value) {
        return callback(new Error("两次密码不同"))
      } else {
        callback()
      }
    }
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
        is_admin: ""
      },
      rules: {
        account: [
          { validator: checkAccount, trigger: 'blur' }
        ],
        password: [
          { validator: checkPassword, trigger: 'blur' }
        ],
        double_password: [
          { validator: doubleCheckPassword, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    submit() {
      this.user_form.validate((valid) => {
        if (valid) {
          console.log("commit")
        } else {
          return false;
        }
      })
    },
    reset() {
      this.user_form.account = ""
      this.user_form.password = ""
      this.user_form.double_password = ""
      this.user_form.name = ""
      this.user_form.phone_number = ""
      this.user_form.sex = ""
      this.user_form.age = ""
      this.user_form.address = ""
      this.user_form.identify_id = ""
      this.user_form.is_active = ""
      this.user_form.is_admin = ""
    }
  }
}

</script>

<style>
.el-input {
  width: 50%
}
</style>
