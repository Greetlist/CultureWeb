<template>
  <div>
    <h1>新增文章</h1>
    <quill-editor
      class="editor"
      ref="article"
      :value="articleContent"
      :options="editorOption"
      @blur="onBlur($event)"
      @focus="onFocus($event)"
      @ready="onReady($event)"
    />
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
      editorOption: {
        modules: {
          toolbar: {
            container: container,
            handlers: {
              'image': this.imageUpload
            }
          }
        }
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
    }
    //videoUpload() {
    //}
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
  height: 40rem;
  width: 80%;
  padding-left: 10%;
  overflow: hidden;
}
</style>
