<template>
  <div>
    <div class="gva-form-box">
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
        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateNotification,
  findNotification
} from '@/api/notification'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
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

const route = useRoute()
const router = useRouter()
const type = ref('')
const formData = ref({
  title: '',
  content: '',
  attachment: '',
  stickyOnTop: false,
  coverage: '所有成员',
})

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findNotification({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.renoti
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
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
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
