<template>
  <div class="highlight-list-container">
    <div class="highlight-actions">
      <el-button type="primary" @click="openHighlightDialog">添加高光片段</el-button>
    </div>
    
    <el-table :data="highlights" style="width: 100%">
      <el-table-column prop="highlightStartTime" label="开始时间" width="180">
        <template #default="scope">
          {{ formatTime(scope.row.highlightStartTime) }}
        </template>
      </el-table-column>
      <el-table-column prop="highlightEndTime" label="结束时间" width="180">
        <template #default="scope">
          {{ formatTime(scope.row.highlightEndTime) }}
        </template>
      </el-table-column>
      <el-table-column prop="highlightType" label="类型" min-width="200"/>
      <el-table-column label="操作" width="240" fixed="right">
        <template #default="scope">
          <el-button type="primary" link @click="playHighlight(scope.row)">播放</el-button>
          <el-button type="primary" link @click="editHighlight(scope.row)">编辑</el-button>
          <el-button type="danger" link @click="deleteHighlight(scope.row)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑高光片段' : '添加高光片段'"
      width="500px"
    >
      <el-form :model="highlightForm" :rules="rules" ref="formRef" label-width="120px">
        <el-form-item label="开始时间(秒)" prop="highlightStartTime">
          <el-input-number v-model="highlightForm.highlightStartTime" :min="0"/>
        </el-form-item>
        <el-form-item label="结束时间(秒)" prop="highlightEndTime">
          <el-input-number v-model="highlightForm.highlightEndTime" :min="0"/>
        </el-form-item>
        <el-form-item label="类型" prop="highlightType">
          <el-input v-model="highlightForm.highlightType"/>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="saveHighlight">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {
  createVideoHighlight,
  getVideoHighlightList,
  updateVideoHighlight,
  deleteVideoHighlight
} from '@/api/leep/videoHighlight'
import { ElMessage, ElMessageBox } from 'element-plus'

const props = defineProps({
  videoId: {
    type: String,
    required: true
  },
  userId: {
    type: String,
    required: true
  },
  videoRef: {
    type: Object,
    required: true
  }
})

const emit = defineEmits(['highlight-play'])

const highlights = ref([])
const dialogVisible = ref(false)
const isEdit = ref(false)
const highlightForm = ref({
  videoId: '',
  userId: '',
  highlightStartTime: 0,
  highlightEndTime: 0,
  highlightType: ''
})

const formatTime = (seconds) => {
  const mins = Math.floor(seconds / 60)
  const secs = seconds % 60
  return `${mins}:${secs.toString().padStart(2, '0')}`
}

// 修改加载高光列表的方法
const loadHighlights = async () => {
  try {
    const response = await getVideoHighlightList({
      page: 1,
      pageSize: 50,
      videoId: props.videoId,
      userId: props.userId
    })
    if (response.code === 0) {
      highlights.value = response.data.list
    }
  } catch (error) {
    ElMessage.error('加载高光列表失败')
  }
}

const openHighlightDialog = () => {
  isEdit.value = false
  highlightForm.value = {
    videoId: props.videoId,
    userId: props.userId,
    highlightStartTime: 0, // 默认从0开始
    highlightEndTime: 0, // 用户手动输入结束时间
    highlightType: ''
  }
  dialogVisible.value = true
}

const editHighlight = (row) => {
  isEdit.value = true
  highlightForm.value = { ...row }
  dialogVisible.value = true
}

// 添加表单验证
const validateTimes = (rule, value, callback) => {
  if (highlightForm.value.highlightEndTime <= highlightForm.value.highlightStartTime) {
    callback(new Error('结束时间必须大于开始时间'))
  } else {
    callback()
  }
}

const rules = {
  highlightStartTime: [
    { required: true, message: '请输入开始时间', trigger: 'blur' },
    { type: 'number', min: 0, message: '开始时间不能小于0', trigger: 'blur' }
  ],
  highlightEndTime: [
    { required: true, message: '请输入结束时间', trigger: 'blur' },
    { type: 'number', min: 0, message: '结束时间不能小于0', trigger: 'blur' },
    { validator: validateTimes, trigger: 'blur' }
  ],
  highlightType: [
    { required: true, message: '请输入高光类型', trigger: 'blur' }
  ]
}

const formRef = ref(null)

// 修改保存高光的方法，确保包含 videoId 和 userId
const saveHighlight = async () => {
  try {
    await formRef.value.validate()
    
    // 确保表单数据包含视频ID和用户ID
    highlightForm.value.videoId = props.videoId
    highlightForm.value.userId = props.userId
    
    if (isEdit.value) {
      await updateVideoHighlight(highlightForm.value)
    } else {
      await createVideoHighlight(highlightForm.value)
    }
    ElMessage.success(isEdit.value ? '更新成功' : '创建成功')
    dialogVisible.value = false
    loadHighlights()
  } catch (error) {
    if (error !== false) { // validate() 返回 false 时不显示错误
      ElMessage.error('操作失败')
    }
  }
}

// 修改删除高光的方法，确保包含 videoId 和 userId
const deleteHighlight = async (row) => {
  try {
    await ElMessageBox.confirm('确定要删除这个高光片段吗？')
    await deleteVideoHighlight({ 
      id: row.id,
      videoId: props.videoId,
      userId: props.userId
    })
    ElMessage.success('删除成功')
    loadHighlights()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('删除失败')
    }
  }
}

const playHighlight = (row) => {
  emit('highlight-play', row)
}

onMounted(() => {
  loadHighlights()
})
</script>

<style scoped>
.highlight-list-container {
  margin-top: 20px;
}

.highlight-actions {
  margin-bottom: 15px;
}
</style>