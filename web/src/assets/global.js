/*
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-11-22 23:57:44
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-12-23 10:49:57
 * @FilePath: \inception\web\src\assets\global.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// store/index.js
import { createStore } from 'vuex'
//
const store = createStore({
  state() {
    return {
      user : {},
      token:{},
    }
  },

  mutations: {
    SET_USER(state, user_info) {
      console.log(state, user_info)
      state.user = user_info;
    },
    SET_TOKEN(state, token) {
      state.token = token;
      localStorage.setItem('token', token);
    }
  },
  actions: {

  },
  getters: {

  }
})
// const store ={}
export default store
