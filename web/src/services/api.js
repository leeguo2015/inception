import axios from 'axios'
import { ElMessage } from 'element-plus'

// 基础URL配置
const getBaseURL = () => {
  if (import.meta.env.VITE_API_BASE_URL) {
    return import.meta.env.VITE_API_BASE_URL + '/v1/api'
  }
  return 'http://127.0.0.1:8080/v1/api'
}

// 创建axios实例
const instance = axios.create({
  baseURL: getBaseURL(),
  timeout: 15000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
instance.interceptors.request.use(
  (config) => {
    // 可以在这里添加全局请求处理，如添加token等
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
instance.interceptors.response.use(
  (response) => {
    // 直接返回数据部分
    return response.data
  },
  (error) => {
    // 统一错误处理
    if (error.response) {
      const status = error.response.status
      console.error('请求失败，状态码：' + status)
      switch (status) {
        case 400:
          ElMessage.error(error.response.data.msg || '请求失败，请稍后重试')
          break
        case 401:
          ElMessage.error('未授权，请重新登录')
          break
        case 403:
          ElMessage.error('权限不足')
          break
        case 404:
          ElMessage.error('请求的资源不存在')
          break
        case 500:
          ElMessage.error('服务器内部错误')
          break
        default:
          ElMessage.error('请求失败，请稍后重试')
      }
    } else if (error.request) {
      ElMessage.error('网络请求失败，请检查您的网络连接')
    } else {
      ElMessage.error('请求配置错误')
    }
    return Promise.reject(error)
  }
)

// 导出具体的API方法
export const api = {
  get: (url, params = {}) => instance.get(url, { params }),
  post: (url, data = {}) => instance.post(url, data),
  put: (url, data = {}) => instance.put(url, data),
  delete: (url) => instance.delete(url),
  patch: (url, data = {}) => instance.patch(url, data)
}

// 导出axios实例，用于需要自定义配置的情况
export { instance as axiosInstance }

// 导出基础URL，用于其他需要URL的地方
export const baseURL = getBaseURL()

export default api