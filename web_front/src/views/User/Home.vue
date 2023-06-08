<template>
  <div>
    <el-button v-if="type === 'mobile'" type="primary">{{ type }}</el-button>
    <el-button v-else-if="type === 'pad'" type="info">{{ type }}</el-button>
    <el-button v-else type="danger">{{ type }}</el-button>
    Width: {{ width }}, Height: {{ height }}
  </div>
</template>

<script>
export default {
  name: "HomeView",
  data() {
    return {
      type: '',
      width: '',
      height: ''
    };
  },
  mounted() {
    this.$nextTick(() => {
      window.addEventListener('resize', this.onResize);
    })
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.onResize);
  },
  methods: {
    onResize() {
      if (window.innerWidth < 500) {
        this.type = "mobile"
      } else if (window.innerWidth >= 500 && window.innerWidth <= 1200) {
        this.type = "pad"
      } else {
        this.type = "screen"
      }
      this.width = window.innerWidth
      this.height = window.innerHeight
    }
  },
  created() {
    this.onResize()
  }
};
</script>

<style lang="scss" scoped>
</style>
