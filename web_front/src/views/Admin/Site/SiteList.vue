<template>
  <div>
    <el-row type="flex">
      <el-col :span="2"><h2>场地列表</h2></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="primary" @click="addNewSite()">新增场地</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="success" @click="saveAllChanges()">保存所有更改</el-button></el-col>
      <el-col :span="2" style="padding-top: 15px"><el-button type="danger" @click="deleteVisible=true">删除所选</el-button></el-col>
    </el-row>
    <el-table
      :data="
        totalSiteList.slice((currentPage - 1) * pageSize, currentPage * pageSize)
      "
      border
      highlight-current-row
      :fit="true"
      style="width: 100%; margin-top: 15px;"
      :row-key="(row) => {return row.site_id}"
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
          <template v-if="scope.row['is_new'] === false">
            <div v-if="item.col === 'site_id'" v-html="convertHtml(scope.row[item.col])"></div>
            <el-button v-else-if="item.col === 'operation'" size="small" type="danger" icon="el-icon-delete" @click="deleteExistSite(scope.row)"></el-button>
            <el-input
              v-else
              v-model="scope.row[item.col]"
              @change=hasEditRow(scope.row)
            />
          </template>
          <template v-else>
            <div v-if="item.col === 'site_id'" v-html="convertHtml(scope.row[item.col])"></div>
            <template v-else-if="item.col === 'operation'">
              <el-button size="small" type="success" icon="el-icon-circle-check" @click="addSingleSite(scope.$index, scope.row)"></el-button>
              <el-button size="small" type="danger" icon="el-icon-circle-close" @click="deleteNewSite(scope.$index)"></el-button>
            </template>
            <el-input
              v-else
              v-model="scope.row[item.col]"
              @change=hasEditRow(scope.row)
            />
          </template>
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
      :total="totalSiteList.length"
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
        <el-col :span="2" :offset="20" style="padding-top: 15px"><el-button type="danger" @click="batchModifySite()">确认</el-button></el-col>
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
          :site="item.name"
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
  name: "SiteList",
  data: function() {
    return {
      totalSiteList: [],
      currentPage: 1,
      pageSize: 10,
      tableColumns: [
        {"col": "site_id", "name": "场地ID", "sort": true},
        {"col": "site_name", "name": "场馆名称", "sort": false},
        {"col": "location", "name": "场馆地址", "sort": false},
        {"col": "phone_number", "name": "联系电话", "sort": false},
        {"col": "contact_name", "name": "联系人", "sort": false},
        {"col": "operation", "name": "操作", "sort": false},
      ],
      modifyShowColumns: [
        {"col": "site_id", "name": "场地ID", "sort": true},
        {"col": "site_name", "name": "场馆名称", "sort": false},
        {"col": "location", "name": "场馆地址", "sort": false},
        {"col": "phone_number", "name": "联系电话", "sort": false},
        {"col": "contact_name", "name": "联系人", "sort": false},
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
      this.modifiedMap[row.site_id] = row
    },

    deleteExistSite(row) {
      this.singleDeleteRow = row
      this.singleDeleteVisible = true
    },
    doDelete(isSingle) {
      var sites = []
      if (isSingle) {
        sites.push(this.singleDeleteRow.site_id)
      } else {
        for (let idx in this.selectedRows) {
          sites.push(this.selectedRows[idx].site_id)
        }
      }
      if (sites.length === 0) {
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
        'sites': sites
      }
      var instance = this
      adminApi.deleteSite(req).then(function (res) {
        notify.ApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] === 0) {
          instance.queryAllSite()
        }
      })
      this.singleDeleteVisible = false
      this.deleteVisible = false
    },
    deleteNewSite(index) {
      this.totalSiteList.splice(index, 1)
    },

    addSingleSite(index, row) {
      var req = {
        'site_name': row.site_name,
        'location': row.location,
        'phone_number': row.phone_number,
        'contact_name': row.contact_name
      }
      var instance = this
      adminApi.addSingleSite(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] !== 0) {
          instance.totalSiteList.splice(index, 1)
        } else {
          instance.totalSiteList[index]["is_new"] = false
        }
      })
    },
    convertHtml(text) {
      return `<span> ${text} </span>`
    },

    addNewSite() {
      var newLine = {
        "site_id": "-",
        "site_name": "",
        "location": "",
        "phone_number": "",
        "contact_name": "",
        "is_new": true
      }
      this.totalSiteList.unshift(newLine)
    },
    saveAllChanges() {
      for (let site_id in this.modifiedMap) {
        this.modifiedRowsList.push(this.modifiedMap[site_id])
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
    queryAllSite() {
      var instance = this
      instance.totalSiteList = []
      adminApi.getTotalSite().then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        if (request_result["return_code"] !== 0) {
          instance.totalSiteList = []
        } else {
          console.log(res.data)
          for (let idx in res.data.sites) {
            var item = res.data.sites[idx]
            item["is_new"] = false
            instance.totalSiteList.push(item)
          }
        }
      })
    },
    batchModifySite() {
      var instance = this
      var req = {
        modify_list: instance.modifiedRowsList
      }
      console.log(req)
      adminApi.batchModifySite(req).then(function (res) {
        var request_result = res.data.request_result
        notifyApiResult(instance, request_result["return_code"], request_result["error_msg"])
        instance.queryAllSite()
      })
    }
  },
  created() {
    this.modifiedMap = new Map()
    this.queryAllSite()
  }
};
</script>
