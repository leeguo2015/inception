<template>
    <el-container>
      <el-header>
          <Menu ></Menu>
      </el-header>
      <el-container>
        <router-view :key="$route.fullPath"></router-view>
      </el-container>
      <IcpRecord  v-if="isHomePage"  recordNumber="渝ICP备2022014533号-1" />

    </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue'
import IcpRecord from './components/IcpRecord.vue';
import { onMounted } from 'vue'
import { useRoute } from 'vue-router';
import { computed } from 'vue';
// import { useStore } from 'vuex' // 引入useStore 方法
// import store from "@/assets/global";

// console.log(store.state.user)  // store 实例对象
// function checkUser() {
//   const store = useStore();
//   const user = localStorage.getItem('user') || sessionStorage.getItem('user');
//   if (user) {
//     store.commit('SET_USER', user);
//   }
// }
const route = useRoute();

const isHomePage = computed(() => {
  return route.path === '/'; // 判断是否为主页
});


onMounted(() => {
  // checkUser()

})
</script>

<style scoped>

</style>
