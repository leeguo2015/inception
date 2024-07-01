<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-24 00:33:58
 * @FilePath: \inception\web\src\components\Menu.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <el-menu default-active="/" class="el-menu-demo menu-padding" mode="horizontal" text-color="#070707" :ellipsis="false"
    @select="handleSelect">

    <el-menu-item index="home" @click="gotoUrl('/')" class="menu-ico">
      <el-icon size="1em" >
        <House />
      </el-icon>
      主页
    </el-menu-item>

    <el-menu-item index="blog_list" @click="gotoUrl('blog_list')"  class="menu-ico">
      <el-icon size="1em">
        <Cherry />
      </el-icon>
      技术
    </el-menu-item>

    <el-menu-item index="article" @click="gotoUrl('article')"  class="menu-ico">
      <el-icon size="1em" >
        <HotWater />
      </el-icon>
      随笔
    </el-menu-item>
    <!-- <el-menu-item index="/blog_add" @click="gotoUrl('blog_add')">
     发布博客
        <el-icon>
          <DocumentAdd />
        </el-icon>
    </el-menu-item> -->
    <div class="flex-grow" />
    <el-menu-item index="/search" @click="gotoUrl('search')" collapse-close-icon="Search"  class="menu-ico">
      <el-icon size="1em" >
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

<script>
export default {
  data() {
    return {
    };
  },
  methods: {
    handleSelect(key, keyPath) {
    },

    gotoUrl(url) {
      this.$router.push(url);
    }

  }
}
</script>

<style>
.menu-padding {
  width: 100%;
  border-bottom: solid 0px var(--el-border-color-base) !important;
  background-color: rgba(0, 0, 0, 0) !important;

}

.flex-grow {
  flex-grow: 1;
}

.menu-ico {
  margin-top: 0.5rem;
}
</style>