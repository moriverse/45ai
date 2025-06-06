<template>
  <view class="container">
    <view class="header">
      <text class="title">Generate Portrait</text>
    </view>

    <view class="upload-section">
      <view
        v-if="!selectedImage"
        class="upload-box"
        @tap="chooseImage"
      >
        <text class="upload-text">Tap to upload photo</text>
      </view>
      <image
        v-else
        :src="selectedImage"
        mode="aspectFit"
        class="preview-image"
      />
    </view>

    <view class="action-section">
      <button
        class="generate-btn"
        :disabled="!selectedImage || isGenerating"
        @tap="startGeneration"
      >
        {{ isGenerating ? 'Generating...' : 'Generate Portrait' }}
      </button>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useTaskStore } from '@/store/task'
import { useUserStore } from '@/store/user'

const taskStore = useTaskStore()
const userStore = useUserStore()

const selectedImage = ref<string>('')
const isGenerating = ref(false)
const templateId = ref<number>(0)

// Get template ID from route params
onLoad((options) => {
  if (options.templateId) {
    templateId.value = parseInt(options.templateId)
  }
})

const chooseImage = async () => {
  try {
    const res = await uni.chooseImage({
      count: 1,
      sizeType: ['compressed'],
      sourceType: ['album', 'camera'],
    })

    if (res.tempFilePaths && res.tempFilePaths.length > 0) {
      selectedImage.value = res.tempFilePaths[0]
    }
  } catch (error) {
    console.error('Failed to choose image:', error)
    uni.showToast({
      title: 'Failed to select image',
      icon: 'none',
    })
  }
}

const startGeneration = async () => {
  if (!selectedImage.value || !templateId.value) return

  isGenerating.value = true

  try {
    // First upload the image to get a URL
    const uploadRes = await uni.uploadFile({
      url: 'http://localhost:8080/api/v1/upload',
      filePath: selectedImage.value,
      name: 'file',
      header: {
        'Authorization': `Bearer ${userStore.token}`,
      },
    })

    if (uploadRes.statusCode === 200) {
      const { url } = JSON.parse(uploadRes.data)
      
      // Create generation task
      const task = await taskStore.createTask(templateId.value, url)
      
      if (task) {
        // Navigate to task detail page
        uni.navigateTo({
          url: `/pages/task/detail?id=${task.id}`,
        })
      } else {
        throw new Error('Failed to create task')
      }
    } else {
      throw new Error('Failed to upload image')
    }
  } catch (error) {
    console.error('Generation failed:', error)
    uni.showToast({
      title: 'Generation failed',
      icon: 'none',
    })
  } finally {
    isGenerating.value = false
  }
}
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

.upload-section {
  margin: 30rpx 0;
}

.upload-box {
  width: 100%;
  height: 400rpx;
  background: #f5f5f5;
  border: 2rpx dashed #d9d9d9;
  border-radius: 12rpx;
  display: flex;
  align-items: center;
  justify-content: center;
}

.upload-text {
  color: #999;
  font-size: 28rpx;
}

.preview-image {
  width: 100%;
  height: 400rpx;
  border-radius: 12rpx;
}

.action-section {
  margin-top: 40rpx;
}

.generate-btn {
  width: 100%;
  height: 88rpx;
  background: #1890ff;
  color: #fff;
  border-radius: 44rpx;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 32rpx;
}

.generate-btn[disabled] {
  background: #d9d9d9;
  color: #fff;
}
</style> 