<template>
  <div>
    <el-row type="flex">
      <el-col :span="2"><h2>预约列表</h2></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="saveAllChanges(row)">保存所有更改</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="deleteAllSelected()">删除所选</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px">
        <el-select
          v-model="currentSelectLabel"
          collapse-tags
          size="medium"
          placeholder="筛选场馆"
          clearable
          @change="changeReservationList"
          @clear="resetReservationList"
        >
          <el-option
            v-for="item in totalSiteList"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </el-col>
    </el-row>
    <el-table
      :data="
        showReservationList.slice((currentPage - 1) * pageSize, currentPage * pageSize)
      "
      border
      highlight-current-row
      :fit="true"
      style="width: 100%; margin-top: 15px;"
      :row-key="(row) => {return row.reservation_id}"
      @selection-change="handleSelectionChange"
      ref="reservationTable"
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
          <template v-else-if="item.col === 'site'">
            <el-select
              v-model="scope.row[item.col]"
              multiple
              collapse-tags
              size="small"
              @change="hasEditRow(scope.row)"
            >
              <el-option
                v-for="item in totalSiteList"
                :key="item.value"
                :label="item.label"
                :value="item.value"
              />
            </el-select>
          </template>
          <template v-else-if="item.col === 'operation'">
            <el-button size="small" type="danger" icon="el-icon-delete" @click="deleteSingleReservation(scope.row)"></el-button>
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
      :total="showReservationList.length"
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
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="batchModifyReservation()">确认修改</el-button></el-col>
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

    <EditReservationDialog ref="editDialog"></EditReservationDialog>
</template>
<script>
import { adminApi } from "@services/admin/"
import { notifyApiResult } from "@js/notify"

export default {
  name: "ReservationList",
  data: function() {
    return {
      totalReservationList: [],
      showReservationList: [],
      currentPage: 1,
      pageSize: 10,
      tableColumns: [
        {"col": "reservation_id", "name": "ID", "sort": true},
        {"col": "site", "name": "场馆", "sort": false},
        {"col": "start_time", "name": "开始时间", "sort": false, width: 200},
        {"col": "end_time", "name": "结束时间", "sort": false, width: 200},
        {"col": "operation", "name": "操作", "sort": false, width: 75},
      ],
      dialogShowColumns: [
        {"col": "reservation_id", "name": "标题", "sort": false},
        {"col": "site", "name": "摘要", "sort": false},
        {"col": "start_time", "name": "开始时间", "sort": false},
        {"col": "end_time", "name": "结束时间", "sort": false},
      ],
      batchSelectedRows: [],
      totalSiteList: [],
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
    EditReservationDialog,
    PreviewReservationDialog
  },
  methods: {
    handleSizeChange(val) {
      this.pageSize = val
    },

    handleCurrentChange(val) {
      this.currentPage = val
    },

    handleSelectionChange(val) {
      this.batchSelectedRows = val
    },

    hasEditRow(row) {
      row.is_modify_content = false
      row.content = ""
      row.create_time = row.create_time
      row.local_save_name = row.local_save_name
      this.modifiedMap[row.reservation_id] = row
    },

    convertHtml(text) {
      return `<span> ${text} </span>`
    },

    deleteSingleReservation(row) {
      this.singleDeleteRow = row
      this.singleDeleteVisible = true
    },

    saveAllChanges() {
      for (let reservation_id in this.modifiedMap) {
        this.modifiedRowsList.push(this.modifiedMap[reservation_id])
      }
      this.modifyVisible = true
    },

    deleteAllSelected() {
      this.batchDeleteVisible = true
    },

    doBatchDelete(isSingle) {
      var deleteReservationIDList = []
      if (isSingle) {
        deleteReservationIDList.push(this.singleDeleteRow.reservation_id)
      } else {
        for (let idx in this.batchSelectedRows) {
          deleteReservationIDList.push(this.batchSelectedRows[idx].reservation_id)
        }
      }
      if (deleteReservationIDList.length === 0) {
        this.$notify({
            title: 'Result',
            type: 'info',
            message: '没有数据'
        })
        this.batchDeleteVisible = false
        return
      }

      var req = {
        'delete_list': deleteReservationIDList
      }

      var instance = this
      adminApi.batchDeleteReservation(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.queryAllReservation()
        }
      })
      this.singleDeleteVisible = false
      this.batchDeleteVisible = false
      this.batchSelectedRows = []
      this.$refs.reservationTable.clearSelection()
    },

    closeModifyDialog() {
      this.modifyVisible = false
      this.modifiedRowsList = []
    },

    closeDeleteDialog() {
      this.batchDeleteVisible = false
    },

    queryAllReservation() {
      var instance = this
      instance.totalReservationList = []
      instance.showReservationList = []
      adminApi.getTotalReservation().then(function (res) {
        var request_result = res.data.request_result
        console.log(request_result)
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] !== 0) {
          instance.totalReservationList = []
          instance.showReservationList = []
        } else {
          for (let idx in res.data.reservation_list) {
            var item = res.data.reservation_list[idx]
            item["reservation_link"] = "/get_reservation" + "?reservation_id=" + item["local_save_name"]
            var currentLabelList = []
            for (let idx in item['labels']) {
              currentLabelList.push(item['labels'][idx].label_id)
            }
            item['labels'] = currentLabelList
            instance.totalReservationList.push(item)
            instance.showReservationList.push(item)
          }
        }
      })
    },

    batchModifyReservation() {
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
      adminApi.batchModifyReservation(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.modifiedRowsList = []
          instance.modifiedMap = new Map()
          instance.queryAllReservation()
        }
      })
      this.modifyVisible = false
    },

    queryAllSite() {
      var instance = this
      instance.totalSiteList = []
      adminApi.getTotalSite().then(function (res) {
        var request_result = res.data.request_result
        if (request_result["return_code"] !== 0) {
          instance.totalSiteList = []
        } else {
          for (let idx in res.data.sites) {
            var item = res.data.sites[idx]
            var option = {
              'value': item.site_id,
              'label': item.site_name
            }
            instance.options.push(option)
          }
        }
      })
    },

    changeReservationList(val) {
      this.showReservationList = []
      for (let idx in this.totalReservationList) {
        if (this.totalReservationList[idx]['site'] === val) {
          this.showReservationList.push(this.totalReservationList[idx])
        }
      }
    },

    resetReservationList() {
      this.showReservationList = this.totalReservationList
    },

    getContent(row, target) {
      var req = {
        'create_time': row.create_time,
        'local_save_name': row.local_save_name,
        'reservation_id': row.reservation_id,
        'is_admin': true,
      }
      var instance = this
      adminApi.getReservationContent(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (target === "edit") {
          instance.$refs.editDialog.reservation_form.content = res.data.reservation_content
          instance.$refs.editDialog.dialogVisible = true
        } else {
          instance.$refs.previewDialog.reservationContent = res.data.reservation_content
          instance.$refs.previewDialog.dialogVisible = true
        }
      })
    }
  },

  created() {
    this.modifiedMap = new Map()
    this.queryAllReservation()
    this.queryAllSite()
  }
};
</script>
