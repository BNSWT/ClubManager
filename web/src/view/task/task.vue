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
        :key="activitiesIndex"
      >
        <el-card>{{ activities.taskDescription }}
          <el-row justify="end">
            <el-button type="info" :icon="Plus" circle @click="openDialog(activities.ID)" />
            <el-button type="primary" :icon="Edit" circle @click="updateTaskFunc(activities)" />
            <el-button type="success" :icon="Check" @click="openUpload(activities.ID)" circle />
            <el-button type="danger" :icon="deleteIcon" circle @click="deleteRow(activities)" />
          </el-row>
        </el-card>
        <el-row />
        <el-timeline-item
          v-for="(activity, index) in activities.children"
          :key="index"
          :type="activity.taskFinished ? 'success' : '' "
          :hollow="!!activity.taskFinished"
          :timestamp="formatDate(activity.taskDeadline)"
        >
          <el-card>
            <el-row>
              <el-tag
                :type="activity.taskAssignee==0 ? 'success' : (activity.taskAssignee=='1' ? 'warning' : '')"
              >{{ assigneeList[activity.taskAssignee] }}</el-tag>
            </el-row>
            <el-row>
              {{ activity.taskDescription }}
            </el-row>
            <el-row
              v-if="activity.children"
            />
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
                      <el-tag :type="activity.taskAssignee==0 ? 'success' : (activity.taskAssignee=='1' ? 'warning' : '')"
                      >{{ assigneeList[activity.taskAssignee] }}</el-tag>
                    </el-row>
                    <el-row>
                      {{ child.taskDescription }}
                    </el-row>
                    <el-row justify="end">
                      <el-button type="info" :icon="Plus" circle @click="openDialog(child.ID)" />
                      <el-button type="primary" :icon="Edit" circle @click="updateTaskFunc(child)" />
                      <el-button type="success" :icon="Check" @click="openUpload(child.ID)" circle />
                      <el-button type="danger" :icon="deleteIcon" circle @click="deleteRow(child)" />
                    </el-row>
                  </el-card>
                </el-timeline-item>
              </el-timeline>
            </el-row>
            <el-row justify="end">
              <el-button type="info" :icon="Plus" circle @click="openDialog(activity.ID)" />
              <el-button type="primary" :icon="Edit" circle @click="updateTaskFunc(activity)" />
              <el-button type="success" :icon="Check" @click="openUpload(activity.ID)" circle />
              <el-button type="danger" :icon="deleteIcon" circle @click="deleteRow(activity)" />
            </el-row>
          </el-card>
        </el-timeline-item>
      </el-timeline>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="重新布置任务">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="任务描述:">
          <el-input v-model="formData.taskDescription" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="负责部门:">
          <el-select v-model="formData.taskAssignee" class="m-2" placeholder="Select" size="large">
            <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
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
    <el-dialog v-model="uploadVisible" :before-close="closeUpload" title="完成当前任务">
      <el-row>
        <form id="fromCont" method="post">
          <div class="fileUpload" @click="inputChange">
            选择文件
            <el-input  class="narrow-input" v-show="false" id="file" ref="FileInput" multiple="multiple" type="file" @change="choseFile"></el-input>
          </div>
        </form>
        <el-button :disabled="limitFileSize" type="primary" size="small" class="uploadBtn" @click="getFile">上传文件</el-button>
      </el-row>
      <el-row>
      <div class="el-upload__tip">请上传任务汇报文件（不超过5MB）</div>
      </el-row>
      <el-row>
      <div class="list">
        <transition name="list" tag="p">
          <div v-if="file" class="list-item">
            <el-icon>
              <document />
            </el-icon>
            <span>{{ file.name }}</span>
            <span class="percentage">{{ percentage }}%</span>
            <el-progress :show-text="false" :text-inside="false" :stroke-width="2" :percentage="percentage" />
          </div>
        </transition>
      </div>
      </el-row>
      <el-row>
      <el-button @click="finishTask">确认完成任务</el-button>
      </el-row>
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

const currentID = ref(0)
const file = ref(null)
const fileMd5 = ref('')
const formDataList = ref([])
const waitUpLoad = ref([])
const waitNum = ref(NaN)
const limitFileSize = ref(false)
const percentage = ref(0)
const percentageFlage = ref(true)

const assigneeList = ['活动部', '宣传部', '学术部']
const options = [
  {
    value: 0,
    label: '活动部',
  },
  {
    value: 1,
    label: '宣传部',
  },
  {
    value: 2,
    label: '学术部',
  },
]

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
import { watch, ref } from 'vue'
import SparkMD5 from 'spark-md5'
import { breakpointContinue, breakpointContinueFinish, findFile as findBKFile, removeChunk } from '@/api/breakpoint'

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

const finishTask = async() => {
  var res = await findTask({ ID: currentID.value })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.retask
  }
  formData.value.taskFinished = true
  res = updateTask(formData.value)
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '成功完成任务'
    })
    closeDialog()
    getTableData()
  }
  formData.value = {
    taskDescription: '',
    taskAssignee: 0,
    taskDeadline: new Date(),
    taskReport: '',
    taskFinished: false,
    parentTask: 0,
    preTask: 0,
  }
  uploadVisible.value=false
}

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

const uploadVisible = ref(false)

const openUpload = (id) => {
  currentID.value = id
  uploadVisible.value = true
}

const closeUpload = () => {
  uploadVisible.value = false
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

// 选中文件的函数
const choseFile = async(e) => {
  const fileR = new FileReader() // 创建一个reader用来读取文件流
  const fileInput = e.target.files[0] // 获取当前文件
  const maxSize = 5 * 1024 * 1024
  file.value = fileInput // file 丢全局方便后面用 可以改进为func传参形式
  percentage.value = 0
  if (file.value.size < maxSize) {
    fileR.readAsArrayBuffer(file.value) // 把文件读成ArrayBuffer  主要为了保持跟后端的流一致
    fileR.onload = async e => {
      // 读成arrayBuffer的回调 e 为方法自带参数 相当于 dom的e 流存在e.target.result 中
      const blob = e.target.result
      const spark = new SparkMD5.ArrayBuffer() // 创建md5制造工具 （md5用于检测文件一致性 这里不懂就打电话问我）
      spark.append(blob) // 文件流丢进工具
      fileMd5.value = spark.end() // 工具结束 产生一个a 总文件的md5
      const FileSliceCap = 1 * 1024 * 1024 // 分片字节数
      let start = 0 // 定义分片开始切的地方
      let end = 0 // 每片结束切的地方a
      let i = 0 // 第几片
      formDataList.value = [] // 分片存储的一个池子 丢全局
      while (end < file.value.size) {
        // 当结尾数字大于文件总size的时候 结束切片
        start = i * FileSliceCap // 计算每片开始位置
        end = (i + 1) * FileSliceCap // 计算每片结束位置
        var fileSlice = file.value.slice(start, end) // 开始切  file.slice 为 h5方法 对文件切片 参数为 起止字节数
        const formData = new window.FormData() // 创建FormData用于存储传给后端的信息
        formData.append('fileMd5', fileMd5.value) // 存储总文件的Md5 让后端知道自己是谁的切片
        formData.append('file', fileSlice) // 当前的切片
        formData.append('chunkNumber', i) // 当前是第几片
        formData.append('fileName', file.value.name) // 当前文件的文件名 用于后端文件切片的命名  formData.appen 为 formData对象添加参数的方法
        formDataList.value.push({ key: i, formData }) // 把当前切片信息 自己是第几片 存入我们方才准备好的池子
        i++
      }
      const params = {
        fileName: file.value.name,
        fileMd5: fileMd5.value,
        chunkTotal: formDataList.value.length
      }
      const res = await findBKFile(params)
      // 全部切完以后 发一个请求给后端 拉当前文件后台存储的切片信息 用于检测有多少上传成功的切片
      const finishList = res.data.file.ExaFileChunk // 上传成功的切片
      const IsFinish = res.data.file.IsFinish // 是否是同文件不同命 （文件md5相同 文件名不同 则默认是同一个文件但是不同文件名 此时后台数据库只需要拷贝一下数据库文件即可 不需要上传文件 即秒传功能）
      if (!IsFinish) {
        // 当是断点续传时候
        waitUpLoad.value = formDataList.value.filter(all => {
          return !(
              finishList &&
              finishList.some(fi => fi.FileChunkNumber === all.key)
          ) // 找出需要上传的切片
        })
      } else {
        waitUpLoad.value = [] // 秒传则没有需要上传的切片
        ElMessage.success('文件已秒传')
      }
      waitNum.value = waitUpLoad.value.length // 记录长度用于百分比展示
    }
  } else {
    limitFileSize.value = true
    ElMessage('请上传小于5M文件')
  }
}

const getFile = () => {
  // 确定按钮
  if (file.value === null) {
    ElMessage('请先上传文件')
    return
  }
  if (percentage.value === 100) {
    percentageFlage.value = false
  }
  formData.value.name = file.value.name
  formData.value.realPath = file.value.name
  sliceFile() // 上传切片
}

const sliceFile = () => {
  waitUpLoad.value &&
  waitUpLoad.value.forEach(item => {
    // 需要上传的切片
    item.formData.append('chunkTotal', formDataList.value.length) // 切片总数携带给后台 总有用的
    const fileR = new FileReader() // 功能同上
    const fileF = item.formData.get('file')
    fileR.readAsArrayBuffer(fileF)
    fileR.onload = e => {
      const spark = new SparkMD5.ArrayBuffer()
      spark.append(e.target.result)
      item.formData.append('chunkMd5', spark.end()) // 获取当前切片md5 后端用于验证切片完整性
      upLoadFileSlice(item)
    }
  })
}

watch(() => waitNum.value, () => { percentage.value = Math.floor(((formDataList.value.length - waitNum.value) / formDataList.value.length) * 100) })

const upLoadFileSlice = async(item) => {
  // 切片上传
  const fileRe = await breakpointContinue(item.formData)
  if (fileRe.code !== 0) {
    return
  }
  waitNum.value-- // 百分数增加
  if (waitNum.value === 0) {
    // 切片传完以后 合成文件
    const params = {
      fileName: file.value.name,
      fileMd5: fileMd5.value
    }
    const res = await breakpointContinueFinish(params)
    if (res.code === 0) {
      // 合成文件过后 删除缓存切片
      const params = {
        fileName: file.value.name,
        fileMd5: fileMd5.value,
        filePath: res.data.filePath,
      }
      ElMessage.success('上传成功')
      await removeChunk(params)
    }
  }
}

const FileInput = ref(null)
const inputChange = () => {
  FileInput.value.dispatchEvent(new MouseEvent('click'))
}

</script>

<style>
.fileUpload {
  padding: 3px 10px;
  font-size: 12px;
  height: 20px;
  line-height: 20px;
  position: relative;
  cursor: pointer;
  color: #000;
  border: 1px solid #c1c1c1;
  border-radius: 4px;
  overflow: hidden;
  display: inline-block;
}
  narrow-input{
    position: absolute;
    font-size: 100px;
    right: 0;
    top: 0;
    opacity: 0;
    cursor: pointer;
  }
  .uploadBtn{
    position: relative;
    top: -10px;
    margin-left: 15px;
  }

</style>
