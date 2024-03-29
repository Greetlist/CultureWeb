<template>
  <div>
    <h1>新增文章</h1>
    <el-form ref="article_form" status-icon :rules="user_check_rules" :model="article_form" label-width="180px" size="medium">
      <el-form-item label="标题" prop="title">
        <el-input v-model="article_form.title" clearable maxlength="30"></el-input>
      </el-form-item>
      <el-form-item label="简要" prop="summary">
        <el-input
          v-model="article_form.summary"
          clearable
          type="textarea"
          maxlength="300"
          show-word-limit
        >
        </el-input>
      </el-form-item>
      <el-form-item label="排序" prop="rank">
        <el-input-number v-model="article_form.rank" :min="1" :max="10" :step="1"></el-input>
      </el-form-item>
      <el-form-item label="标签" prop="label">
        <el-select
          v-model="article_form.labels"
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
      <el-form-item label="是否置顶">
        <el-switch v-model="article_form.is_top"></el-switch>
      </el-form-item>
      <el-form-item label="内容">
        <quill-editor
          ref="article"
          v-model="article_form.content"
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
      </el-form-item>
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
  name: "AddArticle",
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
      article_form: {
        title: '',
        summary: '',
        rank: '',
        label: '',
        is_top: false,
        content: '',
      }
    }
  },
  methods: {
    onBlur(quill) {
    },
    onFocus(quill) {
    },
    onReady(quill) {
    },
    onChange(event) {
      //this.$refs.article.quill.setSelection(event.text.length+1)
      //setTimeout(() => this.$refs.article.quill.setSelection(event.text.length+1), 0)
    },
    fillRequestParam(file) {
      this.uploadData.size = file.size
      this.uploadData.category = this.uploadFileType
      var fList = file.name.split('.')
      if (fList.length === 1 || fList[0] === '' && fList.length === 2) {
        this.uploadData.extension = "png"
      } else {
        this.uploadData.extension = fList.pop()
      }
    },
    imageUpload(state) {
      this.uploadFileType = 'image'
      this.addRange = this.$refs.article.quill.getSelection()
      if (state) {
        let fileInput = document.getElementById('imgInput')
        fileInput.click()
      }
    },
    uploadSuccess(res, file, fileList) {
      this.fullscreenLoading = false
      let vm = this
      let url = ''
      if (this.uploadFileType === 'image') {
        url = res.fetch_url
      }
      if (url !== '') {
        let value = url
        vm.addRange = vm.$refs.article.quill.getSelection()
        vm.$refs.article.quill.insertEmbed(vm.addRange !== null ? vm.addRange.index : 0, vm.uploadFileType, value, Quill.sources.USER)
      } else {
        console.log("inert image error")
      }
      this.$refs.uploadInput.clearFiles()
    },
    //videoUpload() {
    //}

    onSubmit() {
      adminApi.submitArticle(this.article_form).then(function (res) {
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
      this.article_form.title = ''
      this.article_form.summary = ''
      this.article_form.rank = ''
      this.article_form.labels = []
      this.article_form.is_top = false
      this.article_form.content = ''
    }
  },
  computed: {
    editor() {
      return this.$refs.article.quill
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
