<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-24 00:33:58
 * @FilePath: \inception\web\src\components\Menu.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <el-menu default-active="/" :router="true" class="top-menu el-menu-demo menu-padding" mode="horizontal" text-color="#070707"
    :ellipsis="false" @select="handleSelect">
    <el-menu-item index="/" class="menu-ico">
      <el-icon size="1.5em">
        <House />
      </el-icon>
      <div class="menu-ico-font">主页</div>
    </el-menu-item>

    <el-menu-item index="/list/topic" class="menu-ico">
      <el-icon size="1.5em">
        <KnifeFork />
      </el-icon>
      <div class="menu-ico-font">技术</div>
    </el-menu-item>

    <el-menu-item index="/list/article"  class="menu-ico">
      <el-icon size="1.5em">
        <HotWater />
      </el-icon>
      <div class="menu-ico-font">随笔</div>
    </el-menu-item>
<!--    <el-menu-item index="/blog/detail/1"  class="menu-ico">-->
<!--      <el-icon size="1.5em">-->
<!--        <IceTea />-->
<!--      </el-icon>-->
<!--      <div class="menu-ico-font">详情</div>-->
<!--    </el-menu-item>-->
    <el-menu-item index="/resources"  class="menu-ico">
      <el-icon size="1.5em">
        <HotWater />
      </el-icon>
      <div class="menu-ico-font">资源</div>
    </el-menu-item>
    <!-- <el-menu-item index="/blog_add" @click="gotoUrl('blog_add')">
     发布博客
        <el-icon>
          <DocumentAdd />
        </el-icon>
    </el-menu-item> -->
    <div class="flex-grow"/>
    <div class="menu-ico">
      <el-input v-model="searchInput" style="width: 240px; margin-right: 1rem" />
    </div>

    <el-menu-item  class="menu-ico" mode="horizontal">
      <el-icon size="1.5em">
        <Search />
      </el-icon>
      搜索
    </el-menu-item>

    <!-- <el-sub-menu index="/user" v-if="this.$store.state.User">
      <template #title>

        个人信息
        <div class="icon"><el-icon>
            <User />
          </el-icon></div>
      </template> -->
    <!-- <el-menu-item index="personal"> <router-link to="/personal">个人信息</router-link></el-menu-item> -->
    <!-- <el-menu-item index="logout"><router-link to="/logout">注销</router-link></el-menu-item> -->
    <!-- </el-sub-menu> -->

    <!-- <el-menu-item index="/login" @click="gotoUrl('login')" :v-if="!this.$store.state.User"> -->

    <!-- <template #title>
 
          登录    
          <div class="icon"><el-icon>
              <User />
            </el-icon></div>
     
        </template> -->
    <!-- </el-menu-item> -->
  </el-menu>
</template>

<script setup>
import { HotWater, House, IceTea, KnifeFork } from "@element-plus/icons-vue";
import { useRoute, useRouter } from 'vue-router';
import { ref } from "vue";
const route = useRoute();
const router = useRouter();
const userId = ref(route.params.id);
const handleSelect = (key, keyPath) => {
}
const searchInput = ref("")

const gotoUrl = (keyPath) => {
  router.push(keyPath);
}


</script>

<style scoped>

.el-menu {
  width: 100%;
  background-color: rgba(0, 0, 0, 0) !important;
}
.top-menu {
  display: flex;
  align-items: center;
}
.flex-grow {
  flex-grow: 2;
}

.menu-ico-font {
  font-size: 1.2em;
  margin: 0 0.5rem;
}
</style>