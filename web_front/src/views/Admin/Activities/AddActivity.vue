<template>
  <div>
    <h1>新增活动</h1>
    <el-form ref="activity_form" status-icon :rules="user_check_rules" :model="activity_form" label-width="180px" size="medium">
      <el-form-item label="标题" prop="title">
        <el-input v-model="activity_form.title" clearable></el-input>
      </el-form-item>
      <el-form-item label="简要" prop="summary">
        <el-input v-model="activity_form.summary" clearable></el-input>
      </el-form-item>
      <el-form-item label="标签" prop="label">
        <el-select
          v-model="activity_form.labels"
          clearable
          multiple
          collapse-tags
          placeholder="请选择标签"
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
      <el-form-item label="活动时间">
        <el-date-picker
          v-model="form_time_select"
          type="datetimerange"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
        >
        </el-date-picker>
      </el-form-item>

      <el-form-item label="交通信息">
        <el-input v-model="activity_form.transport" clearable type="textarea" maxlength="300"></el-input>
      </el-form-item>

      <el-form-item label="内容">
        <quill-editor
          ref="article"
          v-model="activity_form.content"
          :options="editorOption"
          @blur="onBlur($event)"
          @focus="onFocus($event)"
          @ready="onReady($event)"
          @change="onChange($event)"
        />
      </el-form-item>
      <el-form-item label-width="0" style="text-align: center;">
        <el-button type="primary" style="margin-right: 20%;" @click="onSubmit">提交</el-button>
        <el-button type="danger" @click="onReset">重置</el-button>
    </el-form>
    <el-upload
      ref="uploadInput"
      :action=actionUrl
      :before-upload=fillRequestParam
      :data=uploadData
      :on-success="uploadSuccess"
      style="display:none"
    >
      <el-button size="small" type="primary" id="imgInput" v-loading.fullscreen.lock="fullscreenLoading" element-loading-text="uploading">Upload</el-button>
    </el-upload>
  </div>
</template>

<script>

import { uploadMediaURL } from "@services/admin/index"
import { adminApi } from "@services/admin/"
import { notifyApiResult } from "@js/notify"

import { Quill } from 'vue-quill-editor'
import { container, ImageExtend } from 'quill-image-extend-module'
Quill.register('modules/ImageExtend', ImageExtend)

import "quill/dist/quill.core.css"

export default {
  name: "AddActivity",
  data: function () {
    return {
      uploadFileType: '',
      addRange: '',
      actionUrl: uploadMediaURL,
      uploadData: {},
      options: [],
      editorOption: {
        modules: {
          toolbar: {
            container: container,
            handlers: {
              'image': this.imageUpload
            }
          }
        }
      },
      form_time_select: '',
      activity_form: {
        title: '',
        summary: '',
        labels: [],
        start_time: '',
        end_time: '',
        transport: '',
        content: ''
      }
    }
  },
  created() {
    var instance = this
    instance.totalLabelList = []
    adminApi.getTotalLabel().then(function (res) {
      var request_result = res.data.request_result
      if (request_result["return_code"] !== 0) {
        instance.totalLabelList = []
      } else {
        for (let idx in res.data.labels) {
          var item = res.data.labels[idx]
          var option = {
            'value': item.label_id,
            'label': item.label_name
          }
          instance.options.push(option)
        }
      }
    })
  },
  methods: {
    onBlur(quill) {
    },
    onFocus(quill) {
    },
    onReady(quill) {
    },
    onChange(quill) {
    },
    onSubmit() {
      this.activity_form.start_time = this.form_time_select[0].toISOString()
      this.activity_form.end_time = this.form_time_select[1].toISOString()
      console.log(this.activity_form)
    },
    onReset() {
      this.resetForm()
    },
    resetForm() {
      this.form_time_select = ''
      this.activity_form.title = ''
      this.activity_form.summary = ''
      this.activity_form.labels = []
      this.activity_form.time = ''
      this.activity_form.transport = ''
      this.activity_form.content = ''
    }
  }
};
</script>
