<template>
  <el-dialog
    title="预览活动"
    width="70%"
    :visible.sync="dialogVisible"
    :before-close="closeDialog"
  >
    <div class="ql-containter ql-snow">
      <div class="ql-editor" v-html="activityContent"></div>
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
      activityContent: '',
      dialogVisible: false
    }
  },
  methods: {
    closeDialog() {
      this.activityContent = ''
      this.dialogVisible = false
    }
  }
};
</script>
