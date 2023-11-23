<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-04 15:05:13
 * @FilePath: \inception\web\src\App.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<script setup>

import Hall from './components/Hall.vue';
import Menu from './components/Menu.vue'
</script>

<template >
  
  <div id="main">
    <div id="top-menu">
      <Menu />
    </div>
    <Hall>

    </Hall>
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

/* 
@media (max-width: 768px) {
 .top-menu {
   height: 40px; 
 }
} */

</style>
