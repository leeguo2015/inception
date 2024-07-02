<template>
  <div class="broadside-main">
    <div v-for="tag in data" :key="tag.id">
      <!-- :type="activeList.value.includes(tag.id)? 'info' : 'sucess'"  -->
      <!-- :type="tag? 'info' : 'sucess'"  -->

      <div class="broad-tag">
        <el-tag size="small" 
        @click="addActive(tag.id)" >
          {{ tag.content }}

        </el-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref,computed } from 'vue'
let data = ref([
  {
    "id": 1,
    "content": "编程",
    "created_at": "2023-04-05T10:00:00Z",
    "is_hot": true,
    "description": "涵盖各种编程语言和技术的讨论"
  },
  {
    "id": 2,
    "content": "人工智能",
    "created_at": "2023-06-9T15:30:00Z",
    "is_hot": false,
    "description": "探讨AI最新进展和应用"
  },
  {
    "id": 3,
    "content": "数据分析",
    "created_at": "2023-03-18T08:45:00Z",
    "is_hot": true,
    "description": "关于数据挖掘、统计分析的技巧和案例"
  },
  {
    "id": 4,
    "content": "Web开发",
    "created_at": "2023-01-20T12:00:00Z",
    "is_hot": false,
    "description": "前端与后端技术分享，包括框架、库的使用"
  },
  {
    "id": 5,
    "content": "软件工程",
    "created_at": "2023-02-24T17:15:00Z",
    "is_hot": true,
    "description": "软件开发的最佳实践和项目管理"
  }
])
let activeList = ref([])

const addActive = (id) => {
  let index = activeList.value.indexOf(id);
  if (index > -1) {
    activeList.value.splice(index, 1)
  } else {
    activeList.value.push(id)
  }
}

</script>

<style>
.broadside-main {
  margin: 40px 20px;
  padding: 1em;
  /* width: 100%; */
  min-height: 30vh;
  /* height: 100vh; */
  /* background-color: rgb(35, 37, 36, 0.7); */
}

.broad-tag {
  margin-top: 5px;
}
.active{
  background-color: red;
}
</style>
