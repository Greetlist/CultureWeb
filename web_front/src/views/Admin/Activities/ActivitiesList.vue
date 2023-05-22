<template>
  <div>
    <el-row type="flex">
      <el-col :span="2"><h2>活动列表</h2></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="saveAllChanges(row)">保存所有更改</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="deleteAllSelected()">删除所选</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px">
        <el-select
          v-model="currentSelectLabel"
          collapse-tags
          size="medium"
          placeholder="筛选标签"
          clearable
          @change="changeActivityList"
          @clear="resetActivityList"
        >
          <el-option
            v-for="item in totalLabelList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-col>
      <el-col :span="2" :offset="12" style="padding-top: 15px">
        <el-input v-model="searchKeyWord" placeholder="输入关键字"/>
      </el-col>
      <el-col :span="2" style="padding-top: 15px">
        <el-button type="success" @click="searchActivity">搜索</el-button>
      </el-col>
    </el-row>
    <el-table
      :data="
        showActivityList.slice((currentPage - 1) * pageSize, currentPage * pageSize)
      "
      border
      highlight-current-row
      :fit="true"
      style="width: 100%; margin-top: 15px;"
      :row-key="(row) => {return row.activity_id}"
      @selection-change="handleSelectionChange"
      ref="artileTable"
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
        :width="item.width"
      >
        <template slot-scope="scope">
          <el-switch
            v-if="item.col === 'is_enable'"
            v-model="scope.row[item.col]"
            active-color="#13ce66"
            @change=hasEditRow(scope.row)
          />
          <el-input
            v-else-if="item.col === 'title' || item.col === 'summary'"
            v-model="scope.row[item.col]"
            @change=hasEditRow(scope.row)
          />
          <template v-else-if="item.col === 'activity_link'">
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
            <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteSingleActivity(scope.row)"></el-button>
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
      :total="showActivityList.length"
      background
      style="text-align: center;"
    >
    </el-pagination>

    <el-dialog
      title="批量修改确认"
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
          v-for="item in dialogShowColumns"
          :key="item.col"
          :prop="item.col"
          :label="item.name"
          :sortable="item.sort"
        >
          <template slot-scope="scope">
            <el-switch
              v-if="item.col === 'is_enable'"
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
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="batchModifyActivity()">确认修改</el-button></el-col>
        <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="closeModifyDialog()">取消修改</el-button></el-col>
      </el-row>
    </el-dialog>

    <el-dialog
      title="批量删除确认"
      :visible.sync="batchDeleteVisible"
      width="70%"
      :before-close="closeDeleteDialog"
    >
      <el-table
        :data="batchSelectedRows"
        border
        highlight-current-row
        :fit="true"
        style="width: 100%; margin-top: 15px;"
      >
        <el-table-column
          align="center"
          v-for="item in dialogShowColumns"
          :key="item.col"
          :prop="item.col"
          :label="item.name"
          :sortable="item.sort"
        >
          <template slot-scope="scope">
            <el-switch
              v-if="item.col === 'is_enable'"
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
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="doBatchDelete(false)">确认删除</el-button></el-col>
        <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="closeDeleteDialog()">取消删除</el-button></el-col>
      </el-row>
    </el-dialog>

    <el-dialog
      title="单条删除确认"
      :visible.sync="singleDeleteVisible"
      width="20%"
    >
      <el-row type="flex">
        <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="doBatchDelete(true)">确认删除</el-button></el-col>
        <el-col :span="2" :offset="15" style="padding-top: 15px"><el-button type="success" @click="singleDeleteVisible=false">取消删除</el-button></el-col>
      </el-row>
    </el-dialog>

    <EditActivityDialog ref="editDialog"></EditActivityDialog>
    <PreviewActivityDialog ref="previewDialog"></PreviewActivityDialog>
</template>
<script>
import { adminApi } from "@services/admin/"
import { notifyApiResult } from "@js/notify"
import EditActivityDialog from "@views/Admin/Activity/EditActivityDialog.vue"
import PreviewActivityDialog from "@views/Admin/Activity/PreviewActivityDialog.vue"

export default {
  name: "ActivityList",
  data: function() {
    return {
      totalActivityList: [],
      showActivityList: [],
      currentPage: 1,
      pageSize: 10,
      tableColumns: [
        {"col": "title", "name": "标题", "sort": false},
        {"col": "summary", "name": "摘要", "sort": false},
        {"col": "activity_link", "name": "内容", "sort": false, width: 100},
        {"col": "visit_number", "name": "访问数量", "sort": true, width: 100},
        {"col": "labels", "name": "标签", "sort": false, width: 200},
        {"col": "is_enable", "name": "是否激活", "sort": false, width: 75},
        {"col": "create_time", "name": "创建时间", "sort": true, width: 100},
        {"col": "update_time", "name": "更新时间", "sort": true, width: 100},
        {"col": "operation", "name": "操作", "sort": false, width: 75},
      ],
      dialogShowColumns: [
        {"col": "title", "name": "标题", "sort": false},
        {"col": "summary", "name": "摘要", "sort": false},
        {"col": "labels", "name": "标签", "sort": false},
        {"col": "is_enable", "name": "是否激活", "sort": false},
      ],
      batchSelectedRows: [],
      totalLabelList: [],
      totalLabelMap: '',
      modifiedMap: '',
      modifiedRowsList: [],
      modifyVisible: false,
      batchDeleteVisible: false,
      singleDeleteVisible: false,
      singleDeleteRow: '',

      currentSelectLabel: '',
      searchKeyWord: ''
    }
  },
  components: {
    EditActivityDialog,
    PreviewActivityDialog
  },
  methods: {
    handleSizeChange(val) {
      this.pageSize = val
    },

    handleCurrentChange(val) {
      this.currentPage = val
    },

    handleSelectionChange(val) {
      console.log(val)
      this.batchSelectedRows = val
    },

    hasEditRow(row) {
      row.is_modify_content = false
      row.content = ""
      row.create_time = row.create_time
      row.local_save_name = row.local_save_name
      this.modifiedMap[row.activity_id] = row
    },

    convertHtml(text) {
      return `<span> ${text} </span>`
    },

    deleteSingleActivity(row) {
      this.singleDeleteRow = row
      this.singleDeleteVisible = true
    },

    saveAllChanges() {
      for (let activity_id in this.modifiedMap) {
        this.modifiedRowsList.push(this.modifiedMap[activity_id])
      }
      this.modifyVisible = true
    },

    deleteAllSelected() {
      this.batchDeleteVisible = true
    },

    doBatchDelete(isSingle) {
      var deleteActivityIDList = []
      if (isSingle) {
        deleteActivityIDList.push(this.singleDeleteRow.activity_id)
      } else {
        for (let idx in this.batchSelectedRows) {
          deleteActivityIDList.push(this.batchSelectedRows[idx].activity_id)
        }
      }
      if (deleteActivityIDList.length === 0) {
        this.$notify({
            title: 'Result',
            type: 'info',
            message: '没有数据'
        })
        this.batchDeleteVisible = false
        return
      }

      var req = {
        'delete_list': deleteActivityIDList
      }

      var instance = this
      adminApi.batchDeleteActivity(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.queryAllActivity()
        }
      })
      this.singleDeleteVisible = false
      this.batchDeleteVisible = false
      this.batchSelectedRows = []
      console.log(this.$refs)
      this.$refs.artileTable.clearSelection()
    },

    closeModifyDialog() {
      this.modifyVisible = false
      this.modifiedRowsList = []
    },

    closeDeleteDialog() {
      this.batchDeleteVisible = false
    },

    queryAllActivity() {
      var instance = this
      instance.totalActivityList = []
      instance.showActivityList = []
      adminApi.getTotalActivity().then(function (res) {
        var request_result = res.data.request_result
        console.log(request_result)
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] !== 0) {
          instance.totalActivityList = []
          instance.showActivityList = []
        } else {
          for (let idx in res.data.activity_list) {
            var item = res.data.activity_list[idx]
            item["activity_link"] = "/get_activity" + "?activity_id=" + item["local_save_name"]
            var currentLabelList = []
            for (let idx in item['labels']) {
              currentLabelList.push(item['labels'][idx].label_id)
            }
            item['labels'] = currentLabelList
            instance.totalActivityList.push(item)
            instance.showActivityList.push(item)
          }
        }
      })
    },

    batchModifyActivity() {
      if (this.modifiedRowsList.length === 0) {
        this.$notify({
            title: 'Result',
            type: 'info',
            message: '没有数据'
        })
        this.modifyVisible = false
        return
      }

      var instance = this
      var req = {
        modify_list: instance.modifiedRowsList
      }
      adminApi.batchModifyActivity(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.modifiedRowsList = []
          instance.modifiedMap = new Map()
          instance.queryAllActivity()
        }
      })
      this.modifyVisible = false
    },

    queryAllLabel() {
      var instance = this
      instance.totalLabelList = []
      instance.totalLabelMap = new Map()
      adminApi.getTotalLabel().then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
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
    },

    changeActivityList(val) {
      console.log(val)
      this.showActivityList = []
      for (let idx in this.totalActivityList) {
        if (this.totalActivityList[idx]['labels'].indexOf(val) !== -1) {
          this.showActivityList.push(this.totalActivityList[idx])
        }
      }
    },

    resetActivityList() {
      this.showActivityList = this.totalActivityList
    },

    searchActivity() {
      if (this.searchKeyWord === '') {
        this.queryAllActivity()
        return
      }
      var req = {
        'key_word': this.searchKeyWord
      }
      var instance = this
      adminApi.searchActivity(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.totalActivityList = []
          instance.showActivityList = []
          for (let idx in res.data.activity_list) {
            var item = res.data.activity_list[idx]
            item["activity_link"] = "/get_activity" + "?activity_id=" + item["local_save_name"]
            var currentLabelList = []
            for (let idx in item['labels']) {
              currentLabelList.push(item['labels'][idx].label_id)
            }
            item['labels'] = currentLabelList
            instance.totalActivityList.push(item)
            instance.showActivityList.push(item)
          }
        }
      })
    },

    getContent(row, target) {
      var req = {
        'create_time': row.create_time,
        'local_save_name': row.local_save_name,
        'activity_id': row.activity_id,
        'is_admin': true,
      }
      var instance = this
      adminApi.getActivityContent(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (target === "edit") {
          instance.$refs.editDialog.activity_form.content = res.data.activity_content
          instance.$refs.editDialog.dialogVisible = true
        } else {
          instance.$refs.previewDialog.activityContent = res.data.activity_content
          instance.$refs.previewDialog.dialogVisible = true
        }
      })
    },

    openEditDialog(row) {
      this.$refs.editDialog.options = this.totalLabelList

      this.$refs.editDialog.activity_form.activity_id = row.activity_id
      this.$refs.editDialog.activity_form.title = row.title
      this.$refs.editDialog.activity_form.summary = row.summary
      this.$refs.editDialog.activity_form.labels = row.labels

      this.$refs.editDialog.activity_form.create_time = row.create_time
      this.$refs.editDialog.activity_form.local_save_name = row.local_save_name
      this.getContent(row, "edit")
    },
    openViewDialog(row) {
      this.getContent(row, "preview")
    }
  },

  created() {
    this.modifiedMap = new Map()
    this.queryAllActivity()
    this.queryAllLabel()
  }
};
</script>
