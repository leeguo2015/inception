<template>
    <div class="blog-detail">
      <h1>{{ blog.title }}</h1>
      <div class="meta-info">
        <span>作者：{{ blog.author }}</span>
        <span>发表于：{{ formatDate(blog.created_at) }}</span>
        <span>标签：
          <el-tag
            v-for="tag in blog.tags"
            :key="tag"
            size="small"
            type="info"
            @close="removeTag(tag)"
          >
            {{ tag }}
          </el-tag>
        </span>
      </div>
      <div v-html="blog.content"></div>
      <div class="stats">
        <span>阅读量：{{ blog.views }}</span>
        <span>点赞：{{ blog.likes }}</span>
      </div>
      <h2>评论</h2>
      <div class="comments">
        <div v-for="comment in blog.comments" :key="comment.id" class="comment">
          <p><strong>{{ comment.user }}：</strong> {{ comment.content }}</p>
          <span>发布于：{{ formatDate(comment.posted_at) }}</span>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  import { ElTag } from 'element-plus';
  
  const blog = ref({
    /* 您提供的JSON数据 */
    
    "id": 12345,
    "title": "如何有效管理个人时间",
    "created_at": "2023-04-01T10:30:00Z",
    "tags": ["时间管理", "效率提升", "生活黑客"],
    "author": "时光导师",
    "content": "在这个快节奏的时代，有效管理个人时间变得尤为重要。首先，设定明确的目标，无论是短期还是长期的，都能帮助你保持焦点。其次，合理规划每一天，使用时间块技术分配任务，确保工作与休息的平衡。此外，学会说“不”，避免无关紧要的事务侵占你的时间。最后，利用技术工具，如日历应用和待办事项列表，来辅助你的日常安排。",
    "views": 1235,
    "likes": 567,
    "comments": [
      {
        "id": 987,
        "user": "小李子",
        "content": "非常实用的建议！特别是时间块技术，我试了之后效率确实提高了。",
        "posted_at": "2023-04-02T15:45:00Z"
      },
      {
        "id": 988,
        "user": "时间旅者",
        "content": "同意楼上，我也开始尝试设定每日三件事，感觉生活更有条理了。",
        "posted_at": "2023-04-03T08:00:00Z"
      }
    ]
  
  });
  
  // 格式化日期函数
  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return `${date.getFullYear()}-${('0' + (date.getMonth() + 1)).slice(-2)}-${('0' + date.getDate()).slice(-2)} ${('0' + date.getHours()).slice(-2)}:${('0' + date.getMinutes()).slice(-2)}`;
  };
  
  // 示例函数，实际上在生产环境中您可能需要处理逻辑来删除标签
  const removeTag = (tag) => {
    console.log(`Tag "${tag}" would be removed.`);
  };
  </script>
  
  <style scoped>
  /* 在此处添加自定义样式 */
  .blog-detail {
    padding: 2rem;
    color: black;
    background-color: aqua;
  }
  .meta-info {
    margin-bottom: 1rem;
  }
  .stats span {
    margin-right: 1rem;
  }
  .comment {
    margin-bottom: 1rem;
  }
  </style>