<template>
  <div>
    <h1>文章列表</h1>
  </div>
</template>

<script>
import { adminApi } from "@services/admin/"

export default {
  name: "ArticalList",
  data: function() {
    return {
      totalArticleList: []
    }
  },
  methods: {
  },
  created() {
    var res = adminApi.getTotalArticle()
    request_result = res.data.request_result
    this.totalArticleList = res.data.article_list
  }
};
</script>
