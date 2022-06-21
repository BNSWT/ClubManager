<template>
  <div>
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
        @row-click="openDetail"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="" prop="title" width="550" />
        <!--        <el-table-column align="left" label="内容" prop="content" width="120" />-->
        <!--        <el-table-column align="left" label="附件路径" prop="attachment" width="120" />-->
        <!--        <el-table-column align="left" label="是否置顶" prop="stickyOnTop" width="120">-->
        <!--          <template #default="scope">{{ formatBoolean(scope.row.stickyOnTop) }}</template>-->
        <!--        </el-table-column>-->
        <el-table-column align="right" label="" prop="coverage" width="120">
          <template #default="scope">
            <el-tag
              :type="scope.row.coverage=='所有成员' ? 'danger' : (scope.row.coverage=='活动部' ? 'success' : (scope.row.coverage=='宣传部' ? 'warning' : ''))"
            >{{ scope.row.coverage }}</el-tag>
          </template>
        </el-table-column>
        <el-table-column align="right" label="">
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="新建通知">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="内容:">
          <el-input
            v-model="formData.content"
            clearable
            type="textarea"
            placeholder="请输入"
            :rows="10"
          />
        </el-form-item>
        <el-form-item label="">
          <form id="fromCont" method="post">
            <div class="fileUpload" @click="inputChange">
              选择附件
              <input class="narrow-input" v-show="false" id="file" ref="FileInput" multiple="multiple" type="file" @change="choseFile">
            </div>
          </form>
          <el-button :disabled="limitFileSize" type="primary" size="small" class="uploadBtn" @click="getFile">上传附件</el-button>
        </el-form-item>
        <el-form-item label="">
          <div class="el-upload__tip">请上传不超过5MB的文件</div>
        </el-form-item>
        <el-form-item label="">
          <div class="list">
            <transition name="list" tag="p">
              <div v-if="file">
                <el-icon>
                  <document />
                </el-icon>
                <span>{{ file.name }}</span>
                <span class="percentage" align="right">{{ percentage }}%</span>
                <el-progress :show-text="false" :text-inside="false" :stroke-width="2" :percentage="percentage" />
              </div>
            </transition>
          </div>
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

    <el-dialog v-model="detailVisible" :before-close="closeDetail" center>
      <template #header="">
        <div>
          <el-row>
            <el-col :span="20" justify="begin"><h4>{{ detailData.title }}</h4></el-col>
            <el-col :span="2" justify="end">
              <el-tag
                :type="detailData.coverage=='所有成员' ? 'danger' : (detailData.coverage=='活动部' ? 'success' : (sdetailData.coverage=='宣传部' ? 'warning' : ''))"
              >{{ detailData.coverage }}
              </el-tag>
            </el-col>
            <el-col v-if="detailData.stickyOnTop" :span="2" offset="20" justify="end">
              <el-tag type="info"><el-icon><CollectionTag /></el-icon>置顶</el-tag>
            </el-col>
            <el-col v-if="!detailData.stickyOnTop" :span="2" offset="20" justify="end">
              <el-tag type="info"><el-icon><Paperclip /></el-icon>普通</el-tag>
            </el-col>
          </el-row>
        </div>
      </template>
      <el-row><span>{{ detailData.content }}</span></el-row>
      <el-row height="2px">
        <div />
      </el-row>
      <el-row>
        <a :href="detailData.attachmentPath"  style="color:blue;">下载附件{{attachmentName}}</a>
      </el-row>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" type="primary" @click="closeDetail">确 定</el-button>
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
import SparkMD5 from 'spark-md5'
import {
  findFile,
  breakpointContinueFinish,
  removeChunk,
  breakpointContinue
} from '@/api/breakpoint'

const attachmentName = ref('')
const file = ref(null)
const fileMd5 = ref('')
const formDataList = ref([])
const waitUpLoad = ref([])
const waitNum = ref(NaN)
const limitFileSize = ref(false)
const percentage = ref(0)
const percentageFlage = ref(true)

import {
  createNotification,
  deleteNotification,
  deleteNotificationByIds,
  updateNotification,
  findNotification,
  getNotificationList
} from '@/api/notification'

// 全量引入格式化工具 请按需保留
import { Download as downloadIcon } from  '@element-plus/icons-vue'
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, watch } from 'vue'

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
  attachmentPath: '',
  stickyOnTop: false,
  coverage: '所有成员',
})

const detailData = ref({
  title: '',
  content: '',
  attachmentPath: '',
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

const detailVisible = ref(false)

const openDetail = (row, column, event) => {
  detailVisible.value = true
  detailData.value = {
    title: row.title,
    content: row.content,
    attachmentPath: row.attachmentPath,
    stickyOnTop: row.stickyOnTop,
    coverage: '所有成员',
  }
  attachmentName.value = row.attachmentPath.split('/')[row.attachmentPath.split('/').length - 1]
}

const closeDetail = () => {
  detailVisible.value = false
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
    attachmentPath: '',
    stickyOnTop: false,
  }
}
// 弹窗确定
const enterDialog = async() => {
  console.log("form:")
  console.log(formData.value)
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
      const res = await findFile(params)
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
  formData.value.attachmentPath = 'https://github.com/BNSWT/ClubManager/raw/main/server/fileDir/' + file.value.name
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

  .fileName{
  display: inline-block;
  vertical-align: top;
  margin: 6px 15px 0 15px;
}
    .uploadBtn{
  position: relative;
  top: -10px;
  margin-left: 15px;
}
    .tips{
  margin-top: 30px;
  font-size: 14px;
  font-weight: 400;
  color: #606266;
}
    .el-divider{
  margin: 0 0 30px 0;
}

.list{
  margin-top:15px;
}
.list-item {
  display: block;
  margin-right: 10px;
  color: #606266;
  line-height: 25px;
  margin-bottom: 5px;
  width: 40%;
  .percentage{
    float: right;
  }
}
.list-enter-active, .list-leave-active {
  transition: all 1s;
}
.list-enter, .list-leave-to
  /* .list-leave-active for below version 2.1.8 */ {
  opacity: 0;
  transform: translateY(-30px);
}

.el-link {
  margin-right: 8px;
}

.line {
  width: 100%;
}

.line div {
  width: 100%;
  height: 0;
  border-top: 1px solid var(--el-border-color);
}

.el-table .top-row {
  background-color: rgb(244,244,244);
}
</style>
