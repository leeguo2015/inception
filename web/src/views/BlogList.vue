<template>
  <el-card>
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:12px">
      <h3>文章列表</h3>
      <el-button type="primary" @click="$router.push({ name: 'BlogNew' })">写文章</el-button>
    </div>

    <el-table :data="posts" style="width:100%">
      <el-table-column prop="title" label="标题"></el-table-column>
      <el-table-column label="标签">
        <template #default="{ row }">
          <el-tag v-for="t in row.tags" :key="t" style="margin-right:6px">{{ t }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="{ row }">
          <el-button size="mini" @click="$router.push({ name: 'BlogView', params: { id: row.id } })">查看</el-button>
          <el-button size="mini" @click="$router.push({ name: 'BlogEdit', params: { id: row.id } })">编辑</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script setup>
import { onMounted, computed } from 'vue'
import { useBlogStore } from '@/stores/blog'

const blog = useBlogStore()
onMounted(() => blog.fetchList())
const posts = computed(() => blog.posts)
</script>

<style scoped></style>