<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-17 00:09:54
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-13 00:54:45
 * @FilePath: \inception\web\src\views\blog\publish.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->

<!-- <template>
  <div id="article-container">
    <el-container>
      <el-header>
          <el-button @click="createArticle">发表文章</el-button>
        </el-header>
        <el-main>
          <el-article v-for="article in articles" :key="article.id">
            <el-article-header>
              <el-article-title>{{ article.title }}</el-article-title>
              <el-article-meta>发布于 {{ article.createdAt }}</el-article-meta>
            </el-article-header>
            <el-article-content>{{ article.content }}</el-article-content>
          </el-article>
        </el-main>
        <el-footer>
          <el-button @click="createArticle">发表文章</el-button>
        </el-footer>
    </el-container>
  </div> 
</template> -->
<!-- 
 <script>
export default {
  data() {
    return {
      articles: [],
      newArticle: {
        title: 'aaaaaaaaa',
        content: ''
      }
    }
  },
  methods: {
    createArticle() {
      if (!this.newArticle.title || !this.newArticle.content) {
        this.$message.error('请填写标题和内容！');
        return;
      }
      this.articles.push({ ...this.newArticle, id: Date.now() });
      this.$message.success('文章发表成功！');
      this.newArticle = { title: '', content: '' };
    }
  }
}
</script> -->
<template>
  <div class="blog-editor">
    <el-form :model="blog" :rules="rules" ref="form" label-width="80px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="blog.title" class="blog-title"  placeholder="请输入标题"></el-input>
      </el-form-item>

      <el-form-item label="标签" prop="tags">
        <el-select v-model="blog.tags" multiple placeholder="请选择标签" filterable>
          <el-option v-for="tag in tagOptions" :key="tag" :label="tag" :value="tag"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="内容" prop="content">
        <el-input type="textarea" v-model="blog.content" ref="blog-content" id="blog-content" :autosize="{ minRows: 6, maxRows: 10 }"
          placeholder="请输入内容"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm">发布</el-button>
        <el-button @click="resetForm">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { reactive } from 'vue'
import { ElForm, ElFormItem, ElInput, ElSelect, ElOption, ElButton } from 'element-plus'
import Quill from 'quill'
// import 'quill/dist/quill.snow.css'
// import { ref, watch } from 'vue'
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'

import { quillEditor } from 'vue-quill-editor'


export default {
  components: { ElForm, ElFormItem, ElInput, ElSelect, ElOption, ElButton, quillEditor },

  setup() {
    // const quillContainer = ref("blog-content")
    // const editorContent = ref('blog-content')

    // watch(editorContent, value => {
    //   // 监听editorContent的变化，并将其同步到Quill编辑器中
    //   if (quillEditor && quillEditor.root.innerHTML !== value) {
    //     quillEditor.root.innerHTML = value
    //   }
    // })

    // let quillEditor = null

    // // 在组件挂载后初始化Quill编辑器
    // const initQuillEditor = () => {
    //   quillEditor = new Quill(quillContainer.value, {
    //     theme: 'snow', // 这里选择了snow主题，你可以根据需要选择合适的主题
    //     placeholder: '请输入内容',
    //     // 其他Quill编辑器的配置项可以按需添加
    //   })

    //   // 监听Quill编辑器内容变化，并将其同步到editorContent中
    //   quillEditor.on('text-change', () => {
    //     const html = quillEditor.root.innerHTML
    //     if (editorContent.value !== html) {
    //       editorContent.value = html
    //     }
    //   })
    // }

    const blog = reactive({
      title: '',
      tags: [],
      content: ''
    })



    const tagOptions = ['标签1', '标签2', '标签3'] // 标签选项

    const rules = {
      title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
      tags: [{ required: true, message: '请选择标签', trigger: 'blur' }],
      content: [{ required: true, message: '请输入内容', trigger: 'blur' }]
    }

    const submitForm = () => {
      // 表单提交逻辑，可自行根据实际情况编写
    }

    const resetForm = () => {
      // 表单重置逻辑，可自行根据实际情况编写
    }

    return {
      blog, tagOptions, rules, submitForm, resetForm
      // quillContainer,
      // editorContent,
      // initQuillEditor
    }
  }
}
</script>
<style scoped>  /* #article-container {
    background-color: aqua;
    height: 100%;
    width: 100%;
  } */

  .blog-editor {
    /* background-color: aqua; */
    /* max-width: 800px; */
    margin: 0 auto;
    height: 100%;
    width: 100%;
    padding: 1rem;
  }

  .blog-title {
    height: 3rem;
    font-size: 2rem;
    width: 520px;
  }
</style> 
