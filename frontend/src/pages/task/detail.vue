<template>
  <view class="container">
    <view class="header">
      <text class="title">Generation Progress</text>
    </view>

    <view class="task-info">
      <view class="status-section">
        <text class="status-label">Status:</text>
        <text :class="['status-value', task?.status.toLowerCase()]">
          {{ getStatusText(task?.status) }}
        </text>
      </view>

      <view class="images-section">
        <view class="image-container">
          <text class="image-label">Original Photo</text>
          <image
            :src="task?.inputImageUrl"
            mode="aspectFit"
            class="task-image"
          />
        </view>

        <view class="image-container">
          <text class="image-label">Generated Portrait</text>
          <image
            v-if="task?.resultImageUrl"
            :src="task?.resultImageUrl"
            mode="aspectFit"
            class="task-image"
          />
          <view
            v-else
            class="loading-placeholder"
          >
            <text>Processing...</text>
          </view>
        </view>
      </view>

      <view
        v-if="task?.errorMessage"
        class="error-section"
      >
        <text class="error-message">{{ task.errorMessage }}</text>
      </view>
    </view>

    <view class="action-section">
      <button
        v-if="task?.status === 'COMPLETED'"
        class="save-btn"
        @tap="saveImage"
      >
        Save to Album
      </button>
      <button
        v-if="task?.status === 'FAILED'"
        class="retry-btn"
        @tap="retryTask"
      >
        Try Again
      </button>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useTaskStore } from '@/store/task'
import type { TaskStatus } from '@/store/task'

const taskStore = useTaskStore()
const taskId = ref<number>(0)
const task = ref<ReturnType<typeof taskStore.getTaskById> | null>(null)
let pollingInterval: number | null = null

// Get task ID from route params
onLoad((options) => {
  if (options.id) {
    taskId.value = parseInt(options.id)
    task.value = taskStore.getTaskById(taskId.value)
  }
})

const getStatusText = (status?: TaskStatus) => {
  switch (status) {
    case 'PENDING_PAYMENT':
      return 'Waiting for Payment'
    case 'PROCESSING':
      return 'Processing'
    case 'COMPLETED':
      return 'Completed'
    case 'FAILED':
      return 'Failed'
    default:
      return 'Unknown'
  }
}

const startPolling = () => {
  pollingInterval = setInterval(async () => {
    await taskStore.fetchTasks()
    task.value = taskStore.getTaskById(taskId.value)
    
    // Stop polling if task is completed or failed
    if (task.value?.status === 'COMPLETED' || task.value?.status === 'FAILED') {
      stopPolling()
    }
  }, 5000) // Poll every 5 seconds
}

const stopPolling = () => {
  if (pollingInterval) {
    clearInterval(pollingInterval)
    pollingInterval = null
  }
}

const saveImage = async () => {
  if (!task.value?.resultImageUrl) return

  try {
    await uni.saveImageToPhotosAlbum({
      filePath: task.value.resultImageUrl,
    })
    uni.showToast({
      title: 'Saved to album',
      icon: 'success',
    })
  } catch (error) {
    console.error('Failed to save image:', error)
    uni.showToast({
      title: 'Failed to save image',
      icon: 'none',
    })
  }
}

const retryTask = () => {
  uni.navigateBack()
}

onMounted(() => {
  if (task.value?.status === 'PROCESSING') {
    startPolling()
  }
})

onUnmounted(() => {
  stopPolling()
})
</script>

<style>
.container {
  padding: 20rpx;
}

.header {
  padding: 20rpx 0;
  margin-bottom: 30rpx;
}

.title {
  font-size: 36rpx;
  font-weight: bold;
}

.task-info {
  background: #fff;
  border-radius: 12rpx;
  padding: 20rpx;
  margin-bottom: 30rpx;
}

.status-section {
  display: flex;
  align-items: center;
  margin-bottom: 20rpx;
}

.status-label {
  font-size: 28rpx;
  color: #666;
  margin-right: 10rpx;
}

.status-value {
  font-size: 28rpx;
  font-weight: bold;
}

.status-value.pending_payment {
  color: #faad14;
}

.status-value.processing {
  color: #1890ff;
}

.status-value.completed {
  color: #52c41a;
}

.status-value.failed {
  color: #ff4d4f;
}

.images-section {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20rpx;
  margin: 20rpx 0;
}

.image-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.image-label {
  font-size: 24rpx;
  color: #666;
  margin-bottom: 10rpx;
}

.task-image {
  width: 100%;
  height: 300rpx;
  border-radius: 8rpx;
}

.loading-placeholder {
  width: 100%;
  height: 300rpx;
  background: #f5f5f5;
  border-radius: 8rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #999;
  font-size: 24rpx;
}

.error-section {
  margin-top: 20rpx;
  padding: 20rpx;
  background: #fff2f0;
  border: 1rpx solid #ffccc7;
  border-radius: 8rpx;
}

.error-message {
  color: #ff4d4f;
  font-size: 24rpx;
}

.action-section {
  margin-top: 40rpx;
}

.save-btn,
.retry-btn {
  width: 100%;
  height: 88rpx;
  border-radius: 44rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.save-btn {
  background: #1890ff;
  color: #fff;
}

.retry-btn {
  background: #ff4d4f;
  color: #fff;
}
</style> 