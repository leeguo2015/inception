<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-25 20:37:26
 * @FilePath: \inception\web\src\App.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->


<template>
  <Menu></Menu>
  <el-container style="overflow: auto;">
    <el-main>
      <!-- <RouterView /> -->
      <router-view v-slot="{ Component }">
        <transition appear mode="out-in">
          <component :is="Component" key="$route.fullPath" />
        </transition>
      </router-view>
    </el-main>
  </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue'
import { onMounted } from 'vue'
import { useStore } from 'vuex' // 引入useStore 方法
// console.log(store.state.user)  // store 实例对象
function checkUser() {
  const store = useStore();
  const user = localStorage.getItem('user') || sessionStorage.getItem('user');
  // console.log("store.state.user:", store.state.user)
  // console.log("user", user)
  if (user) {
    store.commit('SET_USER', user);
  }
}

onMounted(() => {
  checkUser()

})
</script>

<style scoped>
.route-transition-enter-active,
.route-transition-leave-active {
  transition: opacity 0.5s;
}

.route-transition-enter,
.route-transition-leave-to {
  opacity: 0;
}
</style>
