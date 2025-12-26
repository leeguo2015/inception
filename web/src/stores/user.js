import { defineStore } from 'pinia'

export const useUserStore = defineStore('user', {
  state: () => ({
    user: {},
    token: localStorage.getItem('token') || ''
  }),
  getters: {
    isLogged: (state) => !!state.token
  },
  actions: {
    setUser(userInfo) {
      this.user = userInfo
    },
    setToken(token) {
      this.token = token
      localStorage.setItem('token', token)
    },
    logout() {
      this.user = {}
      this.token = ''
      localStorage.removeItem('token')
    }
  }
})