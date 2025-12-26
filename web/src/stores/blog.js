import { defineStore } from 'pinia'
import { api } from '@/services/api'

export const useBlogStore = defineStore('blog', {
  state: () => ({
    posts: [],
    current: null,
  }),
  actions: {
    async fetchList() {
      const res = await api.get('/blog')
      this.posts = res.data || []
      return this.posts
    },
    async fetchById(id) {
      const res = await api.get(`/blog/${id}`)
      this.current = res.data || null
      return this.current
    },
    async add({ title, content, tags = [] }) {
      const res = await api.post('/blog/', { title, content, tags })
      return res.data
    },
    async update(id, { title, content, tags = [] }) {
      const res = await api.put(`/blog/${id}`, { title, content, tags })
      return res.data
    },
    async remove(id) {
      const res = await api.delete(`/blog/${id}`)
      return res.data
    }
  }
})