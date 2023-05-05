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
        <template slot-scope="scope">
          <el-switch
            v-if="item.col === 'is_enable' || item.col === 'is_top'"
            v-model="scope.row[item.col]"
            active-color="#13ce66"
            @change=hasEditRow(scope.row)
          />
          <el-input
            v-else-if="item.col === 'title' || item.col === 'summary'"
            v-model="scope.row[item.col]"
            @change=hasEditRow(scope.row)
          />
          <template v-else-if="item.col === 'article_link'">
            <el-button size="small" type="primary" icon="el-icon-edit" circle @click="openEditDialog(scope.row)"></el-button>
            <el-button size="small" type="success" icon="el-icon-view" circle @click="openViewDialog(scope.row)"></el-button>
          </template>
          <template v-else-if="item.col === 'labels'">
            <el-select
              v-model="scope.row[item.col]"
              multiple
              collapse-tags
              size="small"
              @change="hasEditRow(scope.row)"
            >
              <el-option
                v-for="item in totalLabelList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>
          <template v-else-if="item.col === 'operation'">
            <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteArticle(scope.row)"></el-button>
          </template>
          <el-input-number 
            v-else-if="item.col === 'rank'"
            v-model="scope.row[item.col]"
            size="small"
            :min="1"
            :max="10"
            :step="1"
            @change=hasEditRow(scope.row)
          >
          </el-input-number>
          <div v-else v-html="convertHtml(scope.row[item.col])"></div>
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

    <el-dialog
      title="二次确认"
      :visible.sync="modifyVisible"
      width="70%"
      :before-close="closeModifyDialog"
    >
      <el-table
        :data="modifiedRowsList"
        border
        highlight-current-row
        :fit="true"
        style="width: 100%; margin-top: 15px;"
      >
        <el-table-column
          align="center"
          v-for="item in modifyShowColumns"
          :key="item.col"
          :prop="item.col"
          :label="item.name"
          :sortable="item.sort"
        >
          <template slot-scope="scope">
            <el-switch
              v-if="item.col === 'is_enable' || item.col === 'is_top'"
              v-model="scope.row[item.col]"
              active-color="#13ce66"
              disabled
            />
            <el-input
              v-else-if="item.col === 'title' || item.col === 'summary'"
              v-model="scope.row[item.col]"
              disabled
            />
            <template v-else-if="item.col === 'labels'">
              <el-tag
                v-for="label in scope.row.labels"
                type="primary"
                size="small"
                effet="dark"
              >
              {{ totalLabelMap[label].label_name }}
              </el-tag>
            </template>
            <div v-else v-html="convertHtml(scope.row[item.col])"></div>
          </template>
        </el-table-column>
      </el-table>
      <el-row type="flex">
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="batchModifyArticle()">确认</el-button></el-col>
        <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="closeModifyDialog()">取消</el-button></el-col>
      </el-row>
    </el-dialog>
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
        {"col": "labels", "name": "标签", "sort": false},
        {"col": "is_enable", "name": "是否激活", "sort": false},
        {"col": "is_top", "name": "是否置顶", "sort": false},
        {"col": "rank", "name": "排序权重", "sort": true},
        {"col": "create_time", "name": "创建时间", "sort": true},
        {"col": "update_time", "name": "更新时间", "sort": true},
        {"col": "operation", "name": "操作", "sort": false},
      ],
      modifyShowColumns: [
        {"col": "title", "name": "标题", "sort": false},
        {"col": "summary", "name": "摘要", "sort": false},
        {"col": "labels", "name": "标签", "sort": false},
        {"col": "is_enable", "name": "是否激活", "sort": false},
        {"col": "is_top", "name": "是否置顶", "sort": false},
        {"col": "rank", "name": "排序权重", "sort": false},
      ],
      selectedRows: [],
      totalLabelList: [],
      totalLabelMap: '',
      modifiedMap: '',
      modifiedRowsList: [],
      modifyVisible: false,
      deleteVisible: false
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
      this.selectedRows = val
    },
    hasEditRow(row) {
      console.log(row)
      this.modifiedMap[row.article_id] = row
    },
    deleteArticle(row) {
      console.log(row)
    },
    convertHtml(text) {
      return `<span> ${text} </span>`
    },
    saveAllChanges() {
      for (let article_id in this.modifiedMap) {
        this.modifiedRowsList.push(this.modifiedMap[article_id])
      }
      this.modifyVisible = true
    },
    deleteAllSelected() {
      console.log(this.selectedRows)
    },
    closeModifyDialog() {
      this.modifyVisible = false
      this.modifiedRowsList = []
    },
    displayApiResult(returnCode) {
      if (returnCode !== 0) {
        this.$notify({
            title: 'Result',
            type: 'error',
            message: '调用失败'
        })
        this.totalArticleList = []
      } else {
        this.$notify({
            title: 'Result',
            type: 'success',
            message: '调用成功'
        })
      }
    },
    queryAllArticle() {
      var instance = this
      instance.totalArticleList = []
      adminApi.getTotalArticle().then(function (res) {
        var request_result = res.data.request_result
        instance.displayApiResult(request_result["return_code"])
        if (request_result["return_code"] !== 0) {
          instance.totalArticleList = []
        } else {
          for (let idx in res.data.article_list) {
            var item = res.data.article_list[idx]
            item["article_link"] = "/get_article" + "?article_id=" + item["local_save_name"]
            var currentLabelList = []
            for (let idx in item['labels']) {
              currentLabelList.push(item['labels'][idx].label_id)
            }
            item['labels'] = currentLabelList
            instance.totalArticleList.push(item)
          }
        }
      })
    },
    batchModifyArticle() {
      var instance = this
      var req = {
        modify_list: instance.modifiedRowsList
      }
      console.log(req)
      adminApi.batchModifyArticle(req).then(function (res) {
        var request_result = res.data.request_result
        instance.displayApiResult(request_result["return_code"])
        instance.queryAllArticle()
      })
    },
    queryAllLabel() {
      var instance = this
      instance.totalLabelList = []
      instance.totalLabelMap = new Map()
      adminApi.getTotalLabel().then(function (res) {
        var request_result = res.data.request_result
        instance.displayApiResult(request_result["return_code"])
        if (request_result["return_code"] === 0) {
          for (let idx in res.data.labels) {
            var item = res.data.labels[idx]
            var option = {
              'value': item.label_id,
              'label': item.label_name
            }
            instance.totalLabelList.push(option)
            instance.totalLabelMap[item.label_id] = item
          }
        }
      })
    }
  },
  created() {
    this.modifiedMap = new Map()
    this.queryAllArticle()
    this.queryAllLabel()
  }
};
</script>
