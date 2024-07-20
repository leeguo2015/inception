<template>
    <el-container>
      <el-header>
          <Menu ></Menu>
      </el-header>
      <el-container>
        <router-view :key="$route.fullPath"></router-view>
      </el-container>
    </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue'
import { onMounted } from 'vue'
import { useStore } from 'vuex' // 引入useStore 方法
import store from "@/assets/global";

console.log(store.state.user)  // store 实例对象
function checkUser() {
  const store = useStore();
  const user = localStorage.getItem('user') || sessionStorage.getItem('user');
  if (user) {
    store.commit('SET_USER', user);
  }
}

onMounted(() => {
  // checkUser()

})
</script>

<style scoped>

</style>
