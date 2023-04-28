<template>
  <div>
    <el-row type="flex">
      <el-col :span="2"><h2>文章列表</h2></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="saveAllChanges(row)">保存所有更改</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="deleteAllSelected(row)">删除所选</el-button></el-col>
    </el-row>
    <el-table
      :data="
        totalArticleList.slice((currentPage - 1) * pageSize, currentPage * pageSize)
      "
      border
      highlight-current-row
      :fit="true"
      style="width: 100%; margin-top: 15px;"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        width="55"
      >
      </el-table-column>
      <el-table-column
        align="center"
        v-for="item in tableColumns"
        :key="item.col"
        :prop="item.col"
        :label="item.name"
        :sortable="item.sort"
      >
        <template slot-scope="{ row }">
          <el-switch
            v-if="item.col === 'is_delete' || item.col === 'is_enable' || item.col === 'is_top'"
            v-model="row[item.col]"
            active-color="#13ce66"
          />
          <el-input
            v-else-if="item.col === 'title' || item.col === 'summary'"
            v-model="row[item.col]"/>
          <template v-else-if="item.col === 'article_link'">
            <el-button size="small" type="primary" icon="el-icon-edit" circle @click="openEditDialog(row)"></el-button>
            <el-button size="small" type="success" icon="el-icon-view" circle @click="openViewDialog(row)"></el-button>
          </template>
          <template v-else-if="item.col === 'operation'">
            <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteArticle(row)"></el-button>
          </template>
          <el-input-number 
            v-else-if="item.col === 'rank'"
            v-model="row[item.col]"
            size="small"
            :min="1"
            :max="10"
            :step="1"></el-input-number>
          <el-input
            v-else
            disabled
            v-model="row[item.col]"
          />
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[10, 20, 50]"
      :page-size="pageSize"
      layout="total,sizes,prev,pager,next,jumper"
      :total="totalArticleList.length"
      background
      style="text-align: center;"
    >
    </el-pagination>
  </div>
</template>

<script>
import { adminApi } from "@services/admin/"

export default {
  name: "ArticalList",
  data: function() {
    return {
      totalArticleList: [],
      currentPage: 1,
      pageSize: 10,
      tableColumns: [
        {"col": "title", "name": "标题", "sort": false},
        {"col": "summary", "name": "摘要", "sort": false},
        {"col": "article_link", "name": "内容", "sort": false},
        {"col": "visit_number", "name": "访问数量", "sort": true},
        {"col": "state", "name": "状态", "sort": true},
        {"col": "is_enable", "name": "是否激活", "sort": false},
        {"col": "is_top", "name": "是否置顶", "sort": false},
        {"col": "rank", "name": "排序权重", "sort": true},
        {"col": "create_time", "name": "创建时间", "sort": true},
        {"col": "update_time", "name": "更新时间", "sort": true},
        {"col": "operation", "name": "操作", "sort": false},
      ],
      selected_rows: []
    }
  },
  methods: {
    handleSizeChange(val) {
      this.pageSize = val
    },
    handleCurrentChange(val) {
      this.currentPage = val
    },
    handleSelectionChange(val) {
      this.selected_rows = val
    },
    deleteArticle(row) {
      console.log(row)
    },
    saveAllChanges() {
      console.log(row)
    },
    deleteAllSelected() {
      console.log(row)
    }
  },
  created() {
    var instance = this
    adminApi.getTotalArticle().then(function (res) {
      var request_result = res.data.request_result
      if (request_result["return_code"] !== 0) {
        instance.$notify({
            title: 'Result',
            type: 'error',
            message: '查询失败'
        })
        instance.totalArticleList = []
      } else {
        instance.$notify({
            title: 'Result',
            type: 'success',
            message: '查询成功'
        })
        for (let idx in res.data.article_list) {
          var item = res.data.article_list[idx]
          item["article_link"] = "/get_article" + "?article_id=" + item["local_save_name"]
          instance.totalArticleList.push(item)
        }
      }
    })
  }
};
</script>
