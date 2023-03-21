<template>
    <el-menu :default-active="activeIndex" class="el-menu-demo padding1" mode="horizontal" background-color="#545c64"
        text-color="#fff" :ellipsis="false" @select="handleSelect">

        <el-menu-item index="0">
            HOME <el-icon>
                <House />
            </el-icon>
        </el-menu-item>
        <el-menu-item index="2" >
            VIEW <el-icon><DataBoard /></el-icon>
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="3">SEARCH <el-icon>
                <Search />
            </el-icon></el-menu-item>
        <el-sub-menu index="4">
            <template #title>USER <el-icon><User/></el-icon></template>
            <el-menu-item index="2-1">信息更改</el-menu-item>
            <el-menu-item index="2-1">文章管理</el-menu-item>
        </el-sub-menu>
    </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
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