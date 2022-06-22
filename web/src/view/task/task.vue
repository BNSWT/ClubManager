<template>
  <div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog(0)">新增根级任务</el-button>
      </div>
<!--      <el-table-->
<!--        :data="displayData"-->
<!--        row-key="ID"-->
<!--        style="width: 100%"-->
<!--      >-->
<!--        <el-table-column align="left" label="" prop="taskDescription" width="300" />-->
<!--        <el-table-column align="right" label="">-->
<!--          <template #default="scope">-->
<!--            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateTaskFunc(scope.row)">变更</el-button>-->
<!--            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>-->
<!--          </template>-->
<!--        </el-table-column>-->
<!--      </el-table>-->

      <el-timeline
          v-for="(activities, activitiesIndex) in displayData"
          :key = "activitiesIndex"
      >
        <el-card>{{activities.taskDescription}}
          <el-row justify="end">
            <el-button type="info" :icon="Plus" @click="openDialog(activities.ID)" circle />
            <el-button type="primary" :icon="Edit" @click="updateTaskFunc(activities)" circle />
            <el-button type="success" :icon="Check" circle />
            <el-button type="danger" :icon="deleteIcon" @click="deleteRow(activities)" circle/>
          </el-row>
        </el-card>
        <el-row></el-row>
        <el-timeline-item
            v-for="(activity, index) in activities.children"
            :key="index"
            :type="activity.taskFinished ? 'success' : '' "
            :hollow="!!activity.taskFinished"
            :timestamp="formatDate(activity.taskDeadline)"
        >
          <el-card>
            <el-row>
              <el-tag>{{assigneeList[activity.taskAssignee]}}</el-tag>
            </el-row>
            <el-row>
              {{ activity.taskDescription }}
            </el-row>
            <el-row
                v-if="activity.children"
            ></el-row>
            <el-row
              v-if="activity.children"
            >
              <el-timeline>
                <el-timeline-item
                    v-for="(child, childIndex) in activity.children"
                    :key="childIndex"
                    :type="child.taskFinished ? 'success' : '' "
                    :hollow="!!child.taskFinished"
                    :timestamp="formatDate(child.taskDeadline)"
                >
                  <el-card>
                    <el-row>
                      <el-tag>{{assigneeList[activity.taskAssignee]}}</el-tag>
                    </el-row>
                    <el-row>
                      {{child.taskDescription}}
                    </el-row>
                    <el-row justify="end">
                      <el-button type="info" :icon="Plus" @click="openDialog(child.ID)" circle />
                      <el-button type="primary" :icon="Edit" @click="updateTaskFunc(child)" circle />
                      <el-button type="success" :icon="Check" circle />
                      <el-button type="danger" :icon="deleteIcon" @click="deleteRow(child)" circle/>
                    </el-row>
                  </el-card>
                </el-timeline-item>
              </el-timeline>
            </el-row>
            <el-row justify="end">
              <el-button type="info" :icon="Plus" @click="openDialog(activity.ID)" circle />
              <el-button type="primary" :icon="Edit" @click="updateTaskFunc(activity)" circle />
              <el-button type="success" :icon="Check" circle />
              <el-button type="danger" :icon="deleteIcon" @click="deleteRow(activity)" circle/>
            </el-row>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="任务描述:">
          <el-input v-model="formData.taskDescription" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责人:">
          <el-input v-model.number="formData.taskAssignee" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="截止时间:">
          <el-date-picker v-model="formData.taskDeadline" type="date" style="width:100%" placeholder="选择日期" clearable />
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
  name: 'Task'
}
</script>

<script setup>

import { Plus, Edit, Check, Delete as deleteIcon } from '@element-plus/icons-vue'

const assigneeList = ['活动部', '学术部','宣传部']

import {
  createTask,
  deleteTask,
  deleteTaskByIds,
  updateTask,
  findTask,
  getTaskList
} from '@/api/task'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  taskDescription: '',
  taskAssignee: 0,
  taskDeadline: new Date(),
  taskReport: '',
  taskFinished: false,
  parentTask: 0,
  preTask: 0,
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
  if (searchInfo.value.taskFinished === '') {
    searchInfo.value.taskFinished = null
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

const getDisplayData = () => {
  var len = 0
  for (let i = 0; i < tableData.value.length; i++) {
    const parentID = tableData.value[i].parentTask
    if (parentID === 0) {
      displayData.value[len] = tableData.value[i]
      displayData.value[len].children = []
      len = len + 1
    } else {
      const index = displayData.value.findIndex((element) =>
        element.ID === parentID
      )
      if (index === -1) {
        const parentIndexOrigin = tableData.value.findIndex((element) => element.ID === parentID)
        const rootID = tableData.value[parentIndexOrigin].parentTask
        const rootIndex = displayData.value.findIndex((element) => element.ID === rootID)
        const childIndex = displayData.value[rootIndex].children.findIndex((element) => element.ID === parentID)
        if (displayData.value[rootIndex].children[childIndex].children === undefined) {
          displayData.value[rootIndex].children[childIndex].children = []
        }
        const childLen = displayData.value[rootIndex].children[childIndex].children.length
        displayData.value[rootIndex].children[childIndex].children[childLen] = tableData.value[i]
      } else {
        const childLen = displayData.value[index].children.length
        displayData.value[index].children[childLen] = tableData.value[i]
      }
    }
  }
  console.log(displayData.value)
}

// 查询
const getTableData = async() => {
  const table = await getTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteTaskFunc(row)
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
  const res = await deleteTaskByIds({ ids })
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
const updateTaskFunc = async(row) => {
  const res = await findTask({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.retask
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteTaskFunc = async(row) => {
  const res = await deleteTask({ ID: row.ID })
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
  formData.value.parentTask = id
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    taskDescription: '',
    taskAssignee: 0,
    taskDeadline: new Date(),
    taskReport: '',
    taskFinished: false,
    parentTask: 0,
    preTask: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  let res
  switch (type.value) {
    case 'create':
      res = await createTask(formData.value)
      break
    case 'update':
      res = await updateTask(formData.value)
      break
    default:
      res = await createTask(formData.value)
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
