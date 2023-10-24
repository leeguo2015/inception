<template>
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
  </template>~
  
  <script>
  export default {
    data() {
      return {
        articles: [],
        newArticle: {
          title: '',
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
  </script>

  <style scoped>  

#article-container{
  background-color: aqua;
  height: 80%;
  width: 80%;
}
</style>