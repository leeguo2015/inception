<template>
    <div class="list">
      <div v-for="post in posts" :key="post.id" class="blog-post">
        <h2>{{ post.title }}</h2>
        <p>作者: {{ post.author }}</p>
        <p>发表日期: {{ formatDate(post.datePublished) }}</p>
        <p>{{ post.content }}</p>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, computed } from 'vue';
  
  // 假设这是从后端获取的数据
  const backendData = [
    {
        "id": 1,
        "title": "Vue 1 入门指南",
        "author": "John Doe",
        "content": "Vue 3 带来了许多新特性...",
        "datePublished": "2023-04-01"
    },
    {
        "id": 1,
        "title": "Vue 2 大师指南",
        "author": "John Doe",
        "content": "Vue 3 带来了许多新特性...",
        "datePublished": "2023-04-01"
    },
    {
        "id": 1,
        "title": "Vue 3 入坟指南",
        "author": "John Doe",
        "content": "Vue 3 带来了许多新特性...",
        "datePublished": "2023-04-01"
    },
    // 更多文章...
]
  
  // 使用ref或computed来处理数据，这里直接使用了原始数据作为示例
  const posts = ref(backendData);
  
  // 一个简单的日期格式化函数
  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return `${date.getFullYear()}-${date.getMonth()+1}-${date.getDate()}`;
  };
  </script>
  
  <style scoped>
  .list {
    display: flex;
    flex-direction: column;
    gap: 2rem;
    padding-top: 1rem;
  }
  
  .blog-post {
    border: 1px solid #ccc;
    padding: 1rem;
    border-radius: 5px;
    width: 100%;
  }
  </style>