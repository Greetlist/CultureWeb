<template>
  <div>
    <el-row type="flex">
      <el-col :span="2"><h2>标签列表</h2></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="primary" @click="addNewLabel()">新增标签</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="saveAllChanges()">保存所有更改</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="deleteVisible=true">删除所选</el-button></el-col>
    </el-row>
    <el-table
      :data="
        totalLabelList.slice((currentPage - 1) * pageSize, currentPage * pageSize)
      "
      border
      highlight-current-row
      :fit="true"
      style="width: 100%; margin-top: 15px;"
      :row-key="(row) => {return row.label_id}"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        :reserve-selection="true"
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
          <el-input
            v-if="item.col === 'label_name'"
            v-model="scope.row[item.col]"
            @change=hasEditRow(scope.row)
          />
          <template v-else-if="item.col === 'operation'">
            <template v-if="scope.row['is_new'] === false">
              <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteExistLabel(scope.row)"></el-button>
            </template>
            <template v-else>
              <el-button size="small" type="success" icon="el-icon-circle-check" @click="addSingleLabel(scope.$index, scope.row)"></el-button>
              <el-button size="small" type="danger" icon="el-icon-circle-close" @click="deleteNewLabel(scope.$index)"></el-button>
            </template>
          </template>
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
      :total="totalLabelList.length"
      background
      style="text-align: center;"
    >
    </el-pagination>

    <el-dialog
      title="确认修改"
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
          <template slot-scope="{ row }">
            <div v-html="convertHtml(row[item.col])"></div>
          </template>
        </el-table-column>
      </el-table>
      <el-row type="flex">
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="batchModifyLabel()">确认</el-button></el-col>
        <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="closeModifyDialog()">取消</el-button></el-col>
      </el-row>
    </el-dialog>

    <el-dialog
      title="确认删除"
      :visible.sync="deleteVisible"
      width="70%"
      :before-close="closeDeleteDialog"
    >
      <el-table
        :data="selectedRows"
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
          <template slot-scope="{ row }">
            <div v-html="convertHtml(row[item.col])"></div>
          </template>
        </el-table-column>
      </el-table>
      <el-row type="flex">
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="doDelete(false)">确认</el-button></el-col>
        <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="closeDeleteDialog()">取消</el-button></el-col>
      </el-row>
    </el-dialog>

    <el-dialog
      title="单条删除确认"
      :visible.sync="singleDeleteVisible"
      width="20%"
    >
      <el-row type="flex">
        <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="doDelete(true)">确认删除</el-button></el-col>
        <el-col :span="2" :offset="15" style="padding-top: 15px"><el-button type="success" @click="singleDeleteVisible=false">取消删除</el-button></el-col>
      </el-row>
    </el-dialog>

  </div>
</template>

<script>
import { adminApi } from "@services/admin/"
import { notifyApiResult } from "@js/notify"

export default {
  name: "LabelList",
  data: function() {
    return {
      totalLabelList: [],
      currentPage: 1,
      pageSize: 10,
      tableColumns: [
        {"col": "label_id", "name": "标签ID", "sort": false},
        {"col": "label_name", "name": "标签名称", "sort": false},
        {"col": "article_number", "name": "关联文章数量", "sort": true},
        {"col": "operation", "name": "操作", "sort": false},
      ],
      modifyShowColumns: [
        {"col": "label_id", "name": "标签ID", "sort": false},
        {"col": "label_name", "name": "标签名称", "sort": false},
        {"col": "article_number", "name": "关联文章数量", "sort": true},
      ],
      selectedRows: [],
      modifiedMap: '',
      modifiedRowsList: [],
      modifyVisible: false,
      deleteVisible: false,
      singleDeleteVisible: false,
      singleDeleteRow: ''
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
      this.modifiedMap[row.label_id] = row
    },

    deleteExistLabel(row) {
      this.singleDeleteRow = row
      this.singleDeleteVisible = true
    },
    doDelete(isSingle) {
      var labels = []
      if (isSingle) {
        labels.push(this.singleDeleteRow.label_id)
      } else {
        for (let idx in this.selectedRows) {
          labels.push(this.selectedRows[idx].label_id)
        }
      }
      if (labels.length === 0) {
        this.$notify({
            title: 'Result',
            type: 'info',
            message: '没有数据'
        })
        this.singleDeleteVisible = false
        this.deleteVisible = false
        return
      }

      var req = {
        'labels': labels
      }
      var instance = this
      adminApi.deleteLabel(req).then(function (res) {
        notify.ApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.queryAllLabel()
        }
      })
      this.singleDeleteVisible = false
      this.deleteVisible = false
    },
    deleteNewLabel(index) {
      this.totalLabelList.splice(index, 1)
    },

    addSingleLabel(index, row) {
      var req = {
        'label_name': row.label_name
      }
      var instance = this
      adminApi.addSingleLabel(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] !== 0) {
          instance.totalLabelList.splice(index, 1)
        } else {
          instance.totalLabelList[index]["is_new"] = false
        }
      })
    },
    convertHtml(text) {
      return `<span> ${text} </span>`
    },

    addNewLabel() {
      var newLine = {
        "label_id": "-",
        "label_name": "",
        "article_num": 0,
        "is_new": true
      }
      this.totalLabelList.unshift(newLine)
    },
    saveAllChanges() {
      for (let label_id in this.modifiedMap) {
        this.modifiedRowsList.push(this.modifiedMap[label_id])
      }
      this.modifyVisible = true
    },
    closeModifyDialog() {
      this.modifyVisible = false
      this.modifiedRowsList = []
    },
    closeDeleteDialog() {
      this.deleteVisible = false
    },
    queryAllLabel() {
      var instance = this
      instance.totalLabelList = []
      adminApi.getTotalLabel().then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] !== 0) {
          instance.totalLabelList = []
        } else {
          console.log(res.data)
          for (let idx in res.data.labels) {
            var item = res.data.labels[idx]
            item["is_new"] = false
            instance.totalLabelList.push(item)
          }
        }
      })
    },
    batchModifyLabel() {
      var instance = this
      var req = {
        modify_list: instance.modifiedRowsList
      }
      console.log(req)
      adminApi.batchModifyLabel(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        instance.queryAllLabel()
      })
    }
  },
  created() {
    this.modifiedMap = new Map()
    this.queryAllLabel()
  }
};
</script>
