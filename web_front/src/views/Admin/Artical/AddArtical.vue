<template>
  <div>
    <h1>新增文章</h1>
    <el-form ref="article_form" status-icon :rules="user_check_rules" :model="article_form" label-width="180px" size="medium" class="editor">
      <el-form-item label="标题" prop="title">
        <el-input v-model="article_form.title" clearable></el-input>
      </el-form-item>
      <el-form-item label="简要" prop="summary">
        <el-input v-model="article_form.summary" clearable></el-input>
      </el-form-item>
      <el-form-item label="排序" prop="rank">
        <el-input-number v-model="article_form.rank" :min="1" :max="10" :step="1"></el-input>
      </el-form-item>
      <el-form-item label="标签" prop="label">
        <el-select v-model="article_form.label" clearable>
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

import { Quill } from 'vue-quill-editor'
import { container, ImageExtend } from 'quill-image-extend-module'
Quill.register('modules/ImageExtend', ImageExtend)

export default {
  name: "AddArtical",
  data: function () {
    return {
      articleContent: '',
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
      adminApi.submitArticle(this.article_form)
    },
    onReset() {
      this.article_form.title = ''
      this.article_form.summary = ''
      this.article_form.rank = ''
      this.article_form.label = ''
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
    this.options = [
      {
        value: '文化建设',
        label: '文化建设'
      },
      {
        value: '旅游宣传',
        label: '旅游宣传'
      }, {
        value: '美食推广',
        label: '美食推广'
      }, {
        value: '活动预览',
        label: '活动预览'
      }, {
        value: '前瞻报告',
        label: '前瞻报告'
      }
    ]
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
