
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="缩略图URL:" prop="thumbnailUrl">
          <el-input v-model="formData.thumbnailUrl" :clearable="true"  placeholder="请输入缩略图URL" />
       </el-form-item>
        <el-form-item label="场馆名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入场馆名称" />
       </el-form-item>
        <el-form-item label="省:" prop="province">
          <el-input v-model="formData.province" :clearable="true"  placeholder="请输入省" />
       </el-form-item>
        <el-form-item label="市:" prop="city">
          <el-input v-model="formData.city" :clearable="true"  placeholder="请输入市" />
       </el-form-item>
        <el-form-item label="区/县:" prop="district">
          <el-input v-model="formData.district" :clearable="true"  placeholder="请输入区/县" />
       </el-form-item>
        <el-form-item label="详细地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入详细地址" />
       </el-form-item>
        <el-form-item label="联系电话:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入联系电话" />
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
  createVenue,
  updateVenue,
  findVenue
} from '@/api/leep/venue'

defineOptions({
    name: 'VenueForm'
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
            thumbnailUrl: '',
            name: '',
            province: '',
            city: '',
            district: '',
            address: '',
            phone: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVenue({ ID: route.query.id })
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
               res = await createVenue(formData.value)
               break
             case 'update':
               res = await updateVenue(formData.value)
               break
             default:
               res = await createVenue(formData.value)
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

<style>
</style>
