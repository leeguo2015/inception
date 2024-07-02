<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-24 00:33:58
 * @FilePath: \inception\web\src\components\Menu.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <el-menu default-active="/" class="top-menu el-menu-demo menu-padding" mode="horizontal" text-color="#070707"
    :ellipsis="false" @select="handleSelect">

    <el-menu-item index="home" @click="gotoUrl('/list')" class="menu-ico">
      <el-icon size="1.5em">
        <House />
      </el-icon>
      <div class="menu-ico-font">主页</div>
    </el-menu-item>

    <el-menu-item index="blog_list" @click="gotoUrl('/list/blog')" class="menu-ico">
      <el-icon size="1.5em">
        <KnifeFork />
      </el-icon>
      <div class="menu-ico-font">技术</div>
    </el-menu-item>

    <el-menu-item index="article" @click="gotoUrl('/list/article')" class="menu-ico">
      <el-icon size="1.5em">
        <HotWater />
      </el-icon>
      <div class="menu-ico-font">随笔</div>
    </el-menu-item>
    <el-menu-item index="detail" @click="gotoUrl('/blog/detail/1')" class="menu-ico">
      <el-icon size="1.5em">
        <IceTea />
      </el-icon>
      <div class="menu-ico-font">详情</div>
    </el-menu-item>
    <!-- <el-menu-item index="/blog_add" @click="gotoUrl('blog_add')">
     发布博客
        <el-icon>
          <DocumentAdd />
        </el-icon>
    </el-menu-item> -->
    <div class="flex-grow"></div>
    <el-menu-item index="/search" @click="gotoUrl('search')" collapse-close-icon="Search" class="menu-ico">
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

<script>
// import { KnifeFork } from '@element-plus/icons-vue/dist/types';

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
.el-menu {
  width: 100%;
  background-color: rgba(0,0,0,0) !important;   
}

.flex-grow {
  flex-grow: 1;
}

.menu-ico {
  font-size: 2em;
  margin-top: 0.5rem;
}

.menu-ico-font {
  font-size: 1.2em;
  margin: 0 0.5rem;
}
</style>