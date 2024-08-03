/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-11-23 23:51:03
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-12-23 10:55:53
 * @FilePath: \inception\web\src\assets\api.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// api.js

import axios from 'axios';
import { ElMessage } from 'element-plus';
const  baseURL =  'http://124.223.114.178:8080'
const  baseURLAPI = baseURL + '/v1/api'
// 创建一个 axios 实例，配置一些全局的请求设置
const instance = axios.create({
  baseURL: baseURLAPI, // 设置接口的基础URL
  // baseURL: 'https://124.223.114.178:8081/v1/api', // 设置接口的基础URL
  timeout: 15000, // 设置请求超时时间
  // 在这里可以设置其他的请求配置，比如请求头等
});

// 添加响应拦截器
instance.interceptors.response.use(
  response => {
    // 这里可以进行其他的响应数据处理
    // 比如判断响应状态码，处理错误等
    // 只返回数据的部分
    return response.data;
  },
  error => {
    ElMessage.error('请求服务器失败');
    // 这里处理请求错误
    return Promise.reject(error);
  }
);

// 封装一个全局的 GET 请求函数
export function get(url, params) {
  return instance.get(url, { params });
}

// 封装一个全局的 POST 请求函数
export function post(url, data) {
  return instance.post(url, data);
}
// 可以继续封装其他类型的请求函数，比如 PUT、DELETE 等

export  default baseURL

