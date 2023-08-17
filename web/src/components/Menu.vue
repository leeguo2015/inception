<template>
    <el-menu :default-active="activeIndex" class="el-menu-demo padding1" mode="horizontal" background-color="#545c64"
        text-color="#fff" :ellipsis="false" @select="handleSelect">

        <el-menu-item index="0">
            <router-link to="/">主页</router-link> <el-icon>
                <House />
                
            </el-icon> 
        </el-menu-item>
        <el-menu-item index="2" >
            文章 <el-icon><DataBoard /></el-icon>
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="3">搜索 <el-icon>
                <Search />
            </el-icon></el-menu-item>
        <el-sub-menu index="4">
            <template #title>个人信息 <el-icon><User/></el-icon></template>
            <el-menu-item index="2-1">个人信息</el-menu-item>
            <el-menu-item index="2-1">注销</el-menu-item>
        </el-sub-menu>
    </el-menu>
</template>

<script>

export default {
    data(){
        return {
            msg :"User"
        }
    }
    
}

</script>

<style>
.padding1 {
    padding: 0 1rem;
    ;
}

.flex-grow {
    flex-grow: 1;
}</style>