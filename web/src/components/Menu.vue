<template>
  <el-menu :default-active="activeIndex" class="el-menu-demo padding1" mode="horizontal" background-color="#617172"
    text-color="#fff" :ellipsis="false" @select="handleSelect">

    <el-menu-item index="home">
      <router-link to="/home">主页<el-icon>
          <House />
        </el-icon>
      </router-link>
    </el-menu-item>
    <el-menu-item index="blog_list">
      <router-link to="/blog_list"> 博客 <el-icon>
          <Document />
        </el-icon></router-link>
    </el-menu-item>
    <el-menu-item router="true" index="/blog_add">
      <router-link to="/blog_add">发布博客
        <el-icon>
          <DocumentAdd />
        </el-icon>
      </router-link>

    </el-menu-item>
    <div class="flex-grow" />
    <el-menu-item router="true" index="/search" collapse-close-icon="Search">
      <router-link to="/search">搜索
        <el-icon>
          <Search />
        </el-icon>
      </router-link>
    </el-menu-item>
    <el-sub-menu index="/login">
      <template #title>

        个人信息
        <!-- <div class="icon"><el-icon><User /></el-icon></div> -->
      </template>
      <el-menu-item index="personal"> <router-link to="/personal">个人信息</router-link></el-menu-item>
      <el-menu-item index="logout"><router-link to="/logout">注销</router-link></el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script>
export default {
  data() {
    return {
      activeIndex: '1',
      activeIndex2: '1'
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
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
}
</style>