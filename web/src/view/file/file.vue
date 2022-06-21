<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog(0)">新增根级文件/目录</el-button>
      </div>
      <el-table
        :data="displayData"
        :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
        row-key="ID"
        style="width: 100%"
      >
        <el-table-column align="left" label="" prop="name" width="300" />
        <el-table-column align="right" label="">
          <template #default="scope">
            <el-button
              icon="plus"
              size="small"
              type="primary"
              link
              @click="openDialog(scope.row.ID)"
            >新增子文件/目录</el-button>
            <el-button
              icon="edit"
              size="small"
              type="primary"
              link
            >下载文件/目录</el-button>
            <el-button
              icon="delete"
              size="small"
              type="primary"
              link
              @click="deleteRow(scope.row)"
            >删除文件/目录</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="名称:">
          <el-input v-model="formData.name" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="真实路径:">
          <el-input v-model="formData.realPath" clearable placeholder="请输入" />
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
  name: 'File'
}
</script>

<script setup>
import {
  createFile,
  deleteFile,
  deleteFileByIds,
  updateFile,
  findFile,
  getFileList
} from '@/api/file'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  parentID: 0,
  name: '',
  realPath: '',
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const displayData = ref([])
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
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

const getDisplayData = () => {
  var len = 0
  for (let i = 0; i < tableData.value.length; i++) {
    const parentID = tableData.value[i].parentID
    if (parentID === 0) {
      displayData.value[len] = tableData.value[i]
      displayData.value[len].children = []
      len = len + 1
    } else {
      const index = displayData.value.findIndex((element) =>
        element.ID === parentID
      )
      const childLen = displayData.value[index].children.length
      displayData.value[index].children[childLen] = tableData.value[i]
    }
  }
  console.log(displayData.value)
}

// 查询
const getTableData = async() => {
  const table = await getFileList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
    getDisplayData()
  }
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
    deleteFileFunc(row)
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
  const res = await deleteFileByIds({ ids })
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
const updateFileFunc = async(row) => {
  const res = await findFile({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.refile
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteFileFunc = async(row) => {
  const res = await deleteFile({ ID: row.ID })
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
const openDialog = (id) => {
  type.value = 'create'
  dialogFormVisible.value = true
  formData.value.parentID = id
  console.log('Opening dialog')
  console.log(formData.value)
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    parentID: 0,
    name: '',
    realPath: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
  console.log('Submiting form')
  console.log(formData.value)
  let res
  switch (type.value) {
    case 'create':
      res = await createFile(formData.value)
      break
    case 'update':
      res = await updateFile(formData.value)
      break
    default:
      res = await createFile(formData.value)
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
</style>
