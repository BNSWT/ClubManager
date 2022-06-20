<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button size="small" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        :data="tableData"
        :row-class-name="tableRowClassName"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="标题" prop="title" width="550" />
        <!--        <el-table-column align="left" label="内容" prop="content" width="120" />-->
        <!--        <el-table-column align="left" label="附件路径" prop="attachment" width="120" />-->
        <!--        <el-table-column align="left" label="是否置顶" prop="stickyOnTop" width="120">-->
        <!--          <template #default="scope">{{ formatBoolean(scope.row.stickyOnTop) }}</template>-->
        <!--        </el-table-column>-->
        <el-table-column align="left" label="通知范围" prop="coverage" width="120" />
        <el-table-column align="left" label="">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateNotificationFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:">
          <el-input v-model="formData.content" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="附件路径:">
          <el-input v-model="formData.attachment" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否置顶:">
          <el-switch v-model="formData.stickyOnTop" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="通知范围:">
          <el-select v-model="formData.coverage" class="m-2" placeholder="Select" size="large">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Notification'
}
</script>

<script setup>
import {
  createNotification,
  deleteNotification,
  deleteNotificationByIds,
  updateNotification,
  findNotification,
  getNotificationList
} from '@/api/notification'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

const options = [
  {
    value: '所有成员',
    label: '所有成员',
  },
  {
    value: '活动部',
    label: '活动部',
  },
  {
    value: '学术部',
    label: '学术部',
  },
  {
    value: '宣传部',
    label: '宣传部',
  },
]

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  title: '',
  content: '',
  attachment: '',
  stickyOnTop: false,
  coverage: '所有成员',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const tableRowClassName = ({ row, rowIndex }) => {
  if (row.stickyOnTop) {
    return 'top-row'
  } else {
    return ''
  }
}

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.stickyOnTop === '') {
    searchInfo.value.stickyOnTop = null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getNotificationList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  const topData = ref([])
  let topDataSize = 0
  const otherData = ref([])
  let otherDataSize = 0

  for (let i = 0; i < tableData.value.length; i++) {
    if (tableData.value[i].stickyOnTop) {
      topData.value[topDataSize] = tableData.value[i]
      topDataSize = topDataSize + 1
    } else {
      otherData.value[otherDataSize] = tableData.value[i]
      otherDataSize = otherDataSize + 1
    }
  }

  tableData.value = topData.value.concat(otherData.value)
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteNotificationFunc(row)
  })
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteNotificationByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateNotificationFunc = async(row) => {
  const res = await findNotification({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.renoti
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteNotificationFunc = async(row) => {
  const res = await deleteNotification({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    title: '',
    content: '',
    attachment: '',
    stickyOnTop: false,
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createNotification(formData.value)
      break
    case 'update':
      res = await updateNotification(formData.value)
      break
    default:
      res = await createNotification(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}
</script>

<style>
.el-table .top-row {
  background-color: rgb(244,244,244);
}
</style>
