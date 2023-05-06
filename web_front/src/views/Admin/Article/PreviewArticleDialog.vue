<template>
  <el-dialog
    title="预览文章"
    width="70%"
    :visible.sync="dialogVisible"
    :before-close="closeDialog"
  >
    <div v-html="articleContent"></div>
  </el-dialog>
</template>

<script>

export default {
  name: "PreviewArticleDialog",
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
