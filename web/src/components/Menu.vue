<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-16 23:43:26
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-24 00:33:58
 * @FilePath: \inception\web\src\components\Menu.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <el-menu default-active="/" class="el-menu-demo padding1" mode="horizontal" background-color="#617172"
    text-color="#fff" :ellipsis="false" @select="handleSelect">

    <el-menu-item index="home" @click="gotoUrl('/')">
      主页<el-icon>
          <House />
        </el-icon>
      
    </el-menu-item>
    <el-menu-item index="blog_list" @click="gotoUrl('blog_list')">

        博客 <el-icon>
          <Document />
        </el-icon>
    </el-menu-item>
    <el-menu-item index="/blog_add" @click="gotoUrl('blog_add')">
     发布博客
        <el-icon>
          <DocumentAdd />
        </el-icon>


    </el-menu-item>
    <div class="flex-grow" />
    <el-menu-item  index="/search"  @click="gotoUrl('search')"  collapse-close-icon="Search">
      搜索
        <el-icon>
          <Search />
        </el-icon>

    </el-menu-item>
    <el-sub-menu index="/user" v-if="this.$store.state.User">
      <template #title>

        个人信息
        <div class="icon"><el-icon>
            <User />
          </el-icon></div>
      </template>
      <!-- <el-menu-item index="personal"> <router-link to="/personal">个人信息</router-link></el-menu-item> -->
      <!-- <el-menu-item index="logout"><router-link to="/logout">注销</router-link></el-menu-item> -->
    </el-sub-menu>

    <el-menu-item index="/login" @click="gotoUrl('login')" :v-if="!this.$store.state.User">

        <template #title>
 
          登录    
          <div class="icon"><el-icon>
              <User />
            </el-icon></div>
     
        </template>
    </el-menu-item>
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
.icon {
  margin: 0 4px;
}

.padding1 {
  padding: 0 1rem;
  width: 100%;
  border-bottom: solid 0px var(--el-border-color-base) !important;

}

.flex-grow {
  flex-grow: 1;
}</style>