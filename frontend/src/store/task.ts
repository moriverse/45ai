import { defineStore } from 'pinia'

export type TaskStatus = 'PENDING_PAYMENT' | 'PROCESSING' | 'COMPLETED' | 'FAILED'

interface Task {
  id: number
  status: TaskStatus
  inputImageUrl: string
  resultImageUrl?: string
  errorMessage?: string
  templateId: number
  createdAt: string
  updatedAt: string
}

interface TaskState {
  tasks: Task[]
  currentTask: Task | null
}

export const useTaskStore = defineStore('task', {
  state: (): TaskState => ({
    tasks: [],
    currentTask: null,
  }),

  getters: {
    getTaskById: (state) => (id: number) => {
      return state.tasks.find(task => task.id === id)
    },
  },

  actions: {
    setTasks(tasks: Task[]) {
      this.tasks = tasks
    },

    setCurrentTask(task: Task | null) {
      this.currentTask = task
    },

    async fetchTasks() {
      try {
        const response = await uni.request({
          url: 'http://localhost:8080/api/v1/tasks',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`,
          },
        })

        if (response.statusCode === 200) {
          this.setTasks(response.data as Task[])
          return true
        }
        return false
      } catch (error) {
        console.error('Failed to fetch tasks:', error)
        return false
      }
    },

    async createTask(templateId: number, inputImageUrl: string) {
      try {
        const response = await uni.request({
          url: 'http://localhost:8080/api/v1/tasks',
          method: 'POST',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`,
          },
          data: {
            template_id: templateId,
            input_image_url: inputImageUrl,
          },
        })

        if (response.statusCode === 200) {
          const task = response.data as Task
          this.tasks.unshift(task)
          this.setCurrentTask(task)
          return task
        }
        return null
      } catch (error) {
        console.error('Failed to create task:', error)
        return null
      }
    },
  },
}) 