<template>
  <div>
    <h1>新增预约</h1>
    <el-form ref="reservation_form" status-icon :rules="user_check_rules" :model="reservation_form" label-width="180px" size="medium">
      <el-form-item label="用途" prop="summary">
        <el-input
          v-model="reservation_form.summary"
          clearable
          type="textarea"
          maxlength="300"
          show-word-limit
        >
        </el-input>
      </el-form-item>
      <el-form-item label="场馆" prop="label">
        <el-select
          v-model="reservation_form.site"
          clearable
          collapse-tags
          placeholder="请选择场馆"
          size="medium"
          style="width: 15rem;"
        >
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="活动占用时间">
        <el-date-picker
          v-model="form_time_select"
          type="datetimerange"
          start-placeholder="开始时间"
          end-placeholder="结束时间"
        >
        </el-date-picker>
      </el-form-item>
      <el-form-item label-width="0" style="text-align: center;">
        <el-button type="primary" style="margin-right: 20%;" @click="onSubmit">提交</el-button>
        <el-button type="danger" @click="onReset">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>

import { uploadMediaURL } from "@services/admin/index"
import { adminApi } from "@services/admin/"
import { notifyApiResult } from "@js/notify"

export default {
  name: "AddReservation",
  data: function () {
    return {
      options: [],
      reservation_form: {
        usage: '',
        site: '',
        start_time: '',
        end_time: ''
      }
    }
  },
  methods: {
    onSubmit() {
      adminApi.submitReservation(this.reservation_form).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.resetForm()
        }
      })
    },
    onReset() {
      this.resetForm()
    },
    resetForm() {
      this.reservation_form.usage = ''
      this.reservation_form.site = ''
      this.reservation_form.start_time = ''
      this.reservation_form.end_time = ''
    }
  },
  created() {
    var instance = this
    instance.totalSiteList = []
    adminApi.getTotalSite().then(function (res) {
      var request_result = res.data.request_result
      if (request_result["return_code"] !== 0) {
        instance.totalSiteList = []
      } else {
        for (let idx in res.data.sites) {
          var item = res.data.sites[idx]
          var option = {
            'value': item.site_id,
            'label': item.site_name
          }
          instance.options.push(option)
        }
      }
    })
  }
};
</script>

<style lang="scss" scoped>
.ql-container{
  overflow: auto;
  max-height: 33rem;
}

.ql-container ::-webkit-scrollbar{
  width: 10px;
  height: 10px;
}
.ql-container ::-webkit-scrollbar-thumb{
  background: #666666;
  border-radius: 5px;
}
.ql-container ::-webkit-scrollbar-track{
  background: #ccc;
  border-radius: 5px;
}
</style>
