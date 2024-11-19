<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="视频ID，UUID:" prop="id">
          <el-input v-model="formData.id" :clearable="true"  placeholder="请输入视频ID，UUID" />
       </el-form-item>
        <el-form-item label="上传用户的ID:" prop="userId">
          <el-input v-model="formData.userId" :clearable="true"  placeholder="请输入上传用户的ID" />
       </el-form-item>
        <el-form-item label="视频路径:" prop="videoUrl">
          <el-input v-model="formData.videoUrl" :clearable="true"  placeholder="请输入视频路径" />
       </el-form-item>
        <el-form-item label="视频类型：local, aliyun, upyun:" prop="videoType">
          <el-input v-model="formData.videoType" :clearable="true"  placeholder="请输入视频类型：local, aliyun, upyun" />
       </el-form-item>
        <el-form-item label="创建时间:" prop="createdAt">
          <el-date-picker v-model="formData.createdAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="更新时间:" prop="updatedAt">
          <el-date-picker v-model="formData.updatedAt" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script setup>
import {
  createVideos,
  updateVideos,
  findVideos
} from '@/api/leep/videos'

defineOptions({
    name: 'VideosForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            id: '',
            userId: '',
            videoUrl: '',
            videoType: '',
            createdAt: new Date(),
            updatedAt: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVideos({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createVideos(formData.value)
               break
             case 'update':
               res = await updateVideos(formData.value)
               break
             default:
               res = await createVideos(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>