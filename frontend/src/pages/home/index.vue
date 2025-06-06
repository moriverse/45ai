<template>
  <view class="container">
    <view class="header">
      <text class="title">AI Portrait Generator</text>
    </view>

    <view class="template-grid">
      <view
        v-for="template in templates"
        :key="template.id"
        class="template-card"
        @tap="selectTemplate(template)"
      >
        <image
          :src="template.coverImageUrl"
          mode="aspectFill"
          class="template-image"
        />
        <view class="template-info">
          <text class="template-name">{{ template.name }}</text>
          <text class="template-price">¥{{ template.price }}</text>
        </view>
      </view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface Template {
  id: number
  name: string
  description: string
  coverImageUrl: string
  price: number
  stylePrompt: string
}

const templates = ref<Template[]>([])

const fetchTemplates = async () => {
  try {
    const response = await uni.request({
      url: 'http://localhost:8080/api/v1/templates',
      method: 'GET',
    })

    if (response.statusCode === 200) {
      templates.value = response.data as Template[]
    }
  } catch (error) {
    console.error('Failed to fetch templates:', error)
    uni.showToast({
      title: 'Failed to load templates',
      icon: 'none',
    })
  }
}

const selectTemplate = (template: Template) => {
  uni.navigateTo({
    url: `/pages/generate/index?templateId=${template.id}`,
  })
}

onMounted(() => {
  fetchTemplates()
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

.template-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 20rpx;
}

.template-card {
  background: #fff;
  border-radius: 12rpx;
  overflow: hidden;
  box-shadow: 0 2rpx 10rpx rgba(0, 0, 0, 0.1);
}

.template-image {
  width: 100%;
  height: 400rpx;
}

.template-info {
  padding: 20rpx;
}

.template-name {
  font-size: 28rpx;
  font-weight: bold;
  margin-bottom: 10rpx;
}

.template-price {
  font-size: 24rpx;
  color: #ff4d4f;
}
</style> 