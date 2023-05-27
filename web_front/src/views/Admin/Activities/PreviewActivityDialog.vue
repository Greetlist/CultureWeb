<template>
  <el-dialog
    title="预览文章"
    width="70%"
    :visible.sync="dialogVisible"
    :before-close="closeDialog"
  >
    <div class="ql-containter ql-snow">
      <div class="ql-editor" v-html="articleContent"></div>
    </div>
  </el-dialog>
</template>

<script>

import "quill/dist/quill.core.css"
import "quill/dist/quill.snow.css"
import "quill/dist/quill.bubble.css"

export default {
  name: "PreviewActivityDialog",
  data: function () {
    return {
      articleContent: '',
      dialogVisible: false
    }
  },
  methods: {
    closeDialog() {
      this.articleContent = ''
      this.dialogVisible = false
    }
  }
};
</script>
