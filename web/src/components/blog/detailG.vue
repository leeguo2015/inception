<template>
  <div class="blog-detail">

    <el-row>
      <el-col :span="24">
        <h2 id="blog-title">{{ blogInfo.blog.title}}</h2>
      </el-col>
    </el-row>
    <el-row :gutter="18" class="suppl-info">
      <el-col :span="4" v-if="blogInfo.user"> {{ blogInfo.user.nickname }}</el-col>
      <el-col :span="6" :offset="2" v-if="blogInfo.blog.created_at">{{ formatDate(blogInfo.blog.created_at) }}</el-col>
      <el-col :span="6" :offset="0">
        <el-tag
            v-for="tag in blogInfo.blog.category"
            :key="tag.id"
            size="small"
            type="info"
        >
          {{ tag.name }}
        </el-tag>
      </el-col>
    </el-row>

    <h1 id="blog-title"></h1>
<!--          <div class="meta-info">-->
<!--            <span v-if="blogInfo.user">作者：{{ blogInfo.user }}</span>-->
<!--            <span v-if="blogInfo.blog.created_at">发表于：{{ formatDate(blogInfo.blog.created_at) }}</span>-->
<!--            <span v-if="blogInfo.blog.category"> 标签：-->
<!--              <el-tag-->
<!--                v-for="tag in blogInfo.blog.category"-->
<!--                :key="tag.id"-->
<!--                size="small"-->
<!--                type="info"-->
<!--              >-->
<!--                {{ tag.name }}-->
<!--              </el-tag>-->
<!--            </span>-->
<!--          </div>-->
    <div  class="blog-content" ref="vditorContent"></div>
<!--    <VDitor ref="vditorRef" :value="blog.content" mode="preview" />-->
  </div>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue';
import Vditor from 'vditor';
import 'vditor/dist/index.css';
import {get} from "@/assets/api";
import router from "@/router";
import {useRoute} from "vue-router";
import {ElMessage} from "element-plus";

const blogInfo = ref({
  "blog": {
    "id": 0,
    "category_id": 0,
    "user_id": 0,
    "title": "加载中...",
    "sort": 0,
    "brief": "",
    "thumb": "",
    "tags": "",
    "referer": "",
    "status": 0,
    "content": "",
    "view_count": 0,
    "reply_count": 0,
    "zan_count": 0,
    "cai_count": 0,
    "created_at": "",
    "updated_at": ""
  },
  "category": {
    "id": 0,
    "name": "",
    "content": "",
    "brief": ""
  },
  "user": null
});

const vditorContent = ref(null);
const route = useRoute()

const formatDate = (dateString) => {
   const date = new Date(dateString);
   return `${date.getFullYear()}-${('0' + (date.getMonth() + 1)).slice(-2)}-${('0' + date.getDate()).slice(-2)} ${('0' + date.getHours()).slice(-2)}:${('0' + date.getMinutes()).slice(-2)}`;
 };

onMounted( () => {
  // await nextTick(); // 确保 DOM 已经更新
  try {
    const vditor = new Vditor(vditorContent.value, {
      // height: 500,
      cache: {
        enable: false
      },
      after: () => {
        // 将博客内容设置为 Vditor 的内容
        console.log(blogInfo.value);

        vditor.setValue(blogInfo.value.blog.content);
        // 切换到预览模式
        // vditor.setPreviewMode("no");

      },
      toolbarConfig: {
        hide: true, // 隐藏工具栏
        // pin:true,
      },
      // preview: {
      //   // markdown: {
      //   //   sanitize: true // 确保安全的渲染
      //   // }
      // },
    });
    get("/article/detail", {id:route.params.id}).then((res) => {
      if (res.code === 0) {
        blogInfo.value = res.data;
        console.log(res);
        vditor.setValue(blogInfo.value.blog.content);

      }else {
        ElMessage.error('获取失败');
      }

    });

    // 手动设置内容为不可编辑
    const editorElement = vditorContent.value.querySelector('.vditor-content');
    if (editorElement) {
      editorElement.contentEditable = 'false';
      editorElement.style.pointerEvents = 'none';
    }

    // 禁用用户交互
    vditorContent.value.querySelectorAll('*').forEach(element => {
      element.style.pointerEvents = 'none';
    });

  } catch (error) {
    console.error('Vditor initialization error:', error);
  }

});
</script>

<style scoped>
.blog-detail {
  margin: 0 auto;
  padding: 20px;

}
.vditor{
  border: none;
  --panel-background-color:rgba(8,0,0,0)
}
#blog-title{
  text-align: center;
  font-size: 2rem;
}

.blog-content {

}

</style>
