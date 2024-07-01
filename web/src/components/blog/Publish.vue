<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-09-17 00:09:54
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-13 00:54:45
 * @FilePath: \inception\web\src\views\blog\publish.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
    <div class="blog-editor">
    <el-container  class="blog-editor">
      <el-aside width="200px" id="narrow-side">Aside</el-aside>
      <el-main>
        <div class="blog-editor">
    <el-form :model="blog" :rules="rules" ref="form" label-width="80px">
      <el-form-item label="标题" prop="title">
        <el-input v-model="blog.title" class="blog-title inpu-areat"  placeholder="请输入标题" ></el-input>
      </el-form-item>

      <el-form-item label="标签" prop="tags">
        <el-select v-model="blog.tags" multiple placeholder="请选择标签" filterable>
          <el-option v-for="tag in tagOptions" :key="tag" :label="tag" :value="tag" class="input-area"  ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="内容" prop="content">
        <el-input type="textarea" v-model="blog.content" ref="blog-content" class="input-text" id="blog-content" :autosize="{ minRows: 6, maxRows: 10 }"
          placeholder="请输入内容"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm">发布</el-button>
        <!-- <el-button @click="resetForm">重置</el-button> -->
      </el-form-item>
    </el-form>
  </div>
      </el-main>
    </el-container>
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

    let article_title_list = [
      {title: '标题1', content: '内容1', author : '作者1', create_time: '2023-11-13 00:54:45'},
      {title: '标题2', content: '内容2', author : '<USERNAME>', create_time: '2023-11-13 00:54:45'},
      {title: '标题3', content: '内容2', author : '<USERNAME>', create_time: '2023-11-13 00:54:45'},
      {title: '标题4', content: '内容2', author : '<USERNAME>', create_time: '2023-11-13 00:54:45'}
        
    ]

    q
    const submitForm = () => {
      // 表单提交逻辑，可自行根据实际情况编写
    }

    const resetForm = () => {
      // 表单重置逻辑，可自行根据实际情况编写
    }

    return {
      blog, tagOptions, rules, submitForm, resetForm,article_title_list
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
    padding: 0.5rem;
  }

  .blog-title {
    height: 3rem;
    font-size: 2rem;
    width: 520px;
  }
   .input-area {
   background-color: rgba(red, green, blue, 0.1)  !important;
  }
  
  #narrow-side{
    background-color: beige;
    height: 100%;
  }
  #log-content{
    background-color: rgba(red, green, blue, 0.1)  !important;
    color: red;

  }
/* ::v-deep .el-form-item .el-form-item__content .el-textarea .el-textarea__inner{
  background-color: rgba(red, green, blue, 0.95)
} */
</style>
