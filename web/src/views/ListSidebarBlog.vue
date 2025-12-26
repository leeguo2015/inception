<script setup>
import { onMounted, ref } from 'vue'
import { debounce } from 'lodash'
import Broadside from '@/components/Broadside.vue'
import list from '@/components/blog/list.vue'
import { api } from '@/services/api'
import { useRoute } from "vue-router"

// 响应式数据
const blogTotal = ref(0)
const currentPage = ref(1)
const blogListData = ref([])
const activeCategoryIDs = ref([])
const categories = ref([])
const loading = ref(false)
const pageSize = 10

// getCategory 函数被注释掉，因为分类获取逻辑已迁移到 Broadside 组件内部
// 如果需要从外部获取分类，可以取消注释
// const fetchCategories = async () => {
//   try {
//     const res = await api.get("category", { Type: route.params.blogType })
//     categories.value = res.data.list || []
//   } catch (error) {
//     console.error('获取分类失败:', error)
//   }
// }
// 获取博客列表
const fetchBlogList = async (page = 1) => {
  loading.value = true
  currentPage.value = page
  
  try {
    const params = {
      page,
      pageSize,
      categoryIds: activeCategoryIDs.value.length > 0 ? activeCategoryIDs.value : undefined
    }
    
    const res = await api.get("/blog", params)
    
    // 转换数据结构以匹配list组件期望的格式
    blogListData.value = (res.data.list || res.data).map(item => ({
      blog: item,
      user: item.user || { nickname: '未知作者' }
    }))
    
    blogTotal.value = res.data.total || res.data.length
  } catch (error) {
    console.error('获取博客列表失败:', error)
    blogListData.value = []
    blogTotal.value = 0
  } finally {
    loading.value = false
  }
}

// 分类变化处理（防抖）
const debouncedUpdateCategories = debounce((categoryIds) => {
  activeCategoryIDs.value = categoryIds
  fetchBlogList(1) // 重置到第一页
}, 500)

// 页面加载时初始化
onMounted(() => {
  fetchBlogList(1)
})

</script>

<template>
  <el-aside width="200px" >
    <Broadside @activeCategoriesIDs="debouncedUpdateCategories" :categories="categories"/>
  </el-aside>
  <el-container id="container">
    <el-main>
      <!-- 加载状态 -->
      <div v-if="loading" class="loading">加载中...</div>
      
      <!-- 博客列表 -->
      <div id="blog-list">
        <list :sharedData="blogListData" />
      </div>
      
      <!-- 无数据提示 -->
      <div v-if="!loading && blogListData.length === 0" class="no-data">
        暂无博客数据
      </div>
    </el-main>
    
    <!-- 分页组件 -->
    <el-footer id="footer" v-if="blogTotal > 0 && !loading">
      <el-pagination
          :page-size="pageSize"
          @update:current-page="fetchBlogList"
          :current-page="currentPage"
          layout="prev, pager, next"
          :total="blogTotal"
      />
    </el-footer>
  </el-container>

</template>

<style scoped>
#footer{
  display: flex;
  justify-content: center; /* 水平居中 */;
}
#container{
  height: calc(100vh - 60px);
}
</style>