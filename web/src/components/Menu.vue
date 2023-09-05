<template>
    <el-menu :default-active="activeIndex" class="el-menu-demo padding1" mode="horizontal" background-color="#617172"
        text-color="#fff" :ellipsis="false" @select="handleSelect">

        <el-menu-item index="0">
            <router-link to="/">主页</router-link> <el-icon>
                <House />

            </el-icon>
        </el-menu-item>
        <el-menu-item index="2">
            文章 <el-icon>
                <DataBoard />
            </el-icon>
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="3">
            搜索 
            <el-icon>
                <Search />
            </el-icon>
        </el-menu-item>
        <el-sub-menu index="4">
            <template #title><router-link to="/login">个人信息 </router-link><el-icon>
                    <User />
                </el-icon></template>
            <el-menu-item index="2-1">个人信息</el-menu-item>
            <el-menu-item index="2-1">注销</el-menu-item>
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
.padding1 {
    padding: 0 1rem;
    width: 100%;
    border-bottom: solid 0px var(--el-border-color-base)!important;

}

.flex-grow {
    flex-grow: 1;
}


</style>