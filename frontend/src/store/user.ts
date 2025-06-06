import { defineStore } from 'pinia'

interface UserState {
  token: string | null
  userInfo: {
    id: number
    nickname: string
    avatarUrl: string
  } | null
}

interface LoginResponse {
  token: string
  user: UserState['userInfo']
}

export const useUserStore = defineStore('user', {
  state: (): UserState => ({
    token: null,
    userInfo: null,
  }),

  getters: {
    isLoggedIn: (state) => !!state.token,
  },

  actions: {
    setToken(token: string) {
      this.token = token
      // Store token in local storage
      uni.setStorageSync('token', token)
    },

    setUserInfo(userInfo: UserState['userInfo']) {
      this.userInfo = userInfo
    },

    async login(code: string) {
      try {
        const response = await uni.request({
          url: 'http://localhost:8080/api/v1/auth/login',
          method: 'POST',
          data: { code },
        })

        if (response.statusCode === 200) {
          const data = response.data as LoginResponse
          this.setToken(data.token)
          this.setUserInfo(data.user)
          return true
        }
        return false
      } catch (error) {
        console.error('Login failed:', error)
        return false
      }
    },

    logout() {
      this.token = null
      this.userInfo = null
      uni.removeStorageSync('token')
    },

    // Initialize store from storage
    initFromStorage() {
      const token = uni.getStorageSync('token')
      if (token) {
        this.token = token
      }
    },
  },
}) 