<template>
  <el-dialog
    title="修改文章"
    width="85%"
    :visible.sync="dialogVisible"
  >
    <el-form ref="article_form" status-icon :rules="user_check_rules" :model="article_form" label-width="180px" size="medium" class="editor">
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
          class="editor"
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
        <el-button type="primary" style="margin-right: 20%;" @click="onSubmit">提交修改</el-button>
        <el-button type="danger" @click="closeDialog">放弃修改</el-button>
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
  </el-dialog>
</template>

<script>

import { notifyApiResult } from "@js/notify"
import { uploadMediaURL } from "@services/admin/index"
import { adminApi } from "@services/admin/"

import { Quill } from 'vue-quill-editor'
import { container, ImageExtend } from 'quill-image-extend-module'
Quill.register('modules/ImageExtend', ImageExtend)

export default {
  name: "EditArticleDialog",
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
        article_id: '',
        title: '',
        summary: '',
        rank: '',
        labels: [],
        is_top: false,
        is_enable: false,

        content: '',
        create_time: '',
        local_save_name: '',
        is_modify_content: false
      },
      dialogVisible: false
    }
  },
  methods: {
    onBlur(quill) {
    },
    onFocus(quill) {
    },
    onReady(quill) {
    },
    onChange(quill) {
      this.article_form.is_modify_content = true
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
      var instance = this
      var modify = [this.article_form]
      var req = {
        modify_list: modify
      }
      if (this.article_form.is_modify_content === false) {
        this.$notify({
            title: 'Result',
            type: 'info',
            message: '没有修改文章'
        })
        return
      }
      adminApi.batchModifyArticle(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        instance.dialogVisible = false
      })
    },
    closeDialog() {
      this.article_form.title = ''
      this.article_form.summary = ''
      this.article_form.rank = ''
      this.article_form.labels = []
      this.article_form.is_top = false
      this.article_form.content = ''
      this.dialogVisible = false

      this.article_form.content = ''
      this.article_form.create_time = ''
      this.article_form.local_save_name = ''
      this.article_form.is_modify_content = false
    }
  },
  computed: {
    editor() {
      return this.$refs.article.quill
    }
  }
};
</script>

<style lang="scss" scoped>
.editor {
  width: 80%;
  padding-left: 10%;
  overflow: hidden;
}
</style>
