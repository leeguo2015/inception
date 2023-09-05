<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Home from './components/Home.vue'
import Menu from './components/Menu.vue'
</script>

<template >
  
  <div id="main">
    <div id="top-menu">
      <Menu />
    </div>
    <div class="master">
      <router-view></router-view>
    </div>
  </div>
</template>

<style scoped>
#main {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction:column;
  height: 100vh;
  /* 设置容器的高度为视口高度，以便看到居中效果 */
  background-color: #e9ddb6;
  background-image: url("./assets/texture.png");
  

}

#top-menu {
position: fixed;
top: 0; /* 设置menu的顶部距离为0 */
left: 0; /* 设置menu的左边距离为0 */
width: 100%; /* 设置menu的宽度为100% */
height: 50px; /* 设置menu的高度为50px */
display: flex;
justify-content: center;
align-items: center;
}

.master {
  box-sizing: border-box;
}

@media (max-width: 768px) {
 .top-menu {
   height: 40px; /* 在屏幕宽度小于768px时，设置菜单的高度为40px */
 }
}
</style>
