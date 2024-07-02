<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-11-04 14:48:05
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-24 22:34:13
 * @FilePath: \inception\web\src\components\menu.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<!--<script setup>-->

<!--</script>-->
<template>
<div class="HallContent" >  
    <broadside v-if="!$route.path.includes('detail')"/>
<!--     <RouterView />-->
    <!-- <div class="sub-content">
      <router-view name="subContent"></router-view>
    </div> -->
</div>
</template>
<script setup>
import { useRoute } from 'vue-router';
import { RouterLink, RouterView } from 'vue-router'
import broadside  from './Broadside.vue'

const route = useRoute();
</script>
<style> 
.HallContent {
  box-sizing: border-box;
  padding: 0.5em;

}
</style>

