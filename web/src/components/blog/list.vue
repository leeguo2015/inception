<template>
  <div class="list">
    <div v-for="post in sharedData" :key="post.blog.id" class="blog-post"
      @click="router.push(`/blog/detail/` + post.blog.id)">
      <el-row>
        <el-col :span="24">
          <h2>{{ post.blog.title }}</h2>
        </el-col>
      </el-row>
      <el-row :gutter="18" class="suppl-info">
        <el-col :span="4"> {{ post.user.nickname }}</el-col>
        <el-col :span="6" :offset="0">{{ formatDate(post.blog.created_at) }}</el-col>
      </el-row>
      <p>{{ formatText(post.blog.content) }}</p>
    </div>
  </div>
</template>

<script setup>
import { slice } from 'lodash';
import { ref, computed } from 'vue';
import { useRouter, useRoute } from "vue-router";

const route = useRoute();
const router = useRouter();
const backendData = []
const posts = ref(backendData);
const props = defineProps({
  sharedData: Array
})
// 一个简单的日期格式化函数
const formatDate = (timeStr) => {
  var date = new Date(timeStr.replace(/-/g, '/')); // 注意：这里用"/"替换"-"是为了兼容不同的浏览器
  date.setHours(date.getHours() + 8);
  function formatDate(date) {
    var year = date.getFullYear();
    var month = ("0" + (date.getMonth() + 1)).slice(-2); // getMonth返回的是0-11，所以需要+1
    var day = ("0" + date.getDate()).slice(-2);
    var hours = ("0" + date.getHours()).slice(-2);
    var minutes = ("0" + date.getMinutes()).slice(-2);
    var seconds = ("0" + date.getSeconds()).slice(-2);
    return year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;
  }

  return formatDate(date);
};

// const shouldShowText = computed(() => props.originalText.length <= 300);

const formatText = (text) => {
  return text.length > 300 ? text.slice(0, 300) + '...' : text;
};
</script>

<style scoped>
.list {
  gap: 2rem;
  height: 100%;
  /* //display: flex;
  flex-direction: column;
  //display: auto; */
  /* overflow: auto; */
}

.blog-post {
  border: 1px solid #ccc;
  padding: 1rem;
  margin: 1rem 0;
  border-radius: 5px;
  width: 100%;
}

.suppl-info {
  color: grey;
  font-size: 0.8rem;
  font-style: italic;
}
</style>