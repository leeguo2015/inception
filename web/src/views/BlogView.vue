<template>
  <el-card>
    <div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:12px">
      <div>
        <h2>{{ post?.title }}</h2>
        <div>
          <el-tag v-for="t in post?.tags || []" :key="t" style="margin-right:6px">{{ t }}</el-tag>
        </div>
      </div>
      <div>
        <el-button v-if="post" @click="$router.push({ name: 'BlogEdit', params: { id: post.id } })">编辑</el-button>
        <el-button @click="$router.push({ name: 'BlogList' })">返回</el-button>
      </div>
    </div>

    <div v-if="post" v-html="post.content" class="blog-content"></div>
    <div v-else>加载中或未找到文章</div>
  </el-card>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useBlogStore } from '@/stores/blog'

const route = useRoute()
const id = route.params.id
const blog = useBlogStore()
const post = ref(null)

onMounted(async () => {
  post.value = await blog.fetchById(id)
})
</script>

<style scoped>
.blog-content img { max-width:100%; }
</style>