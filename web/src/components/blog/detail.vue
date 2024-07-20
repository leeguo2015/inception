<!--<template>-->
<!--    <div class="blog-detail">-->
<!--      <h1>{{ blog.title }}</h1>-->
<!--      <div class="meta-info">-->
<!--        <span>作者：{{ blog.author }}</span>-->
<!--        <span>发表于：{{ formatDate(blog.created_at) }}</span>-->
<!--        <span>标签：-->
<!--          <el-tag-->
<!--            v-for="tag in blog.tags"-->
<!--            :key="tag"-->
<!--            size="small"-->
<!--            type="info"-->
<!--            @close="removeTag(tag)"-->
<!--          >-->
<!--            {{ tag }}-->
<!--          </el-tag>-->
<!--        </span>-->
<!--      </div>-->
<!--      <div v-html="blog.content"></div>-->
<!--      <div class="stats">-->
<!--        <span>阅读量：{{ blog.views }}</span>-->
<!--        <span>点赞：{{ blog.likes }}</span>-->
<!--      </div>-->
<!--      <h2>评论</h2>-->
<!--      <div class="comments">-->
<!--        <div v-for="comment in blog.comments" :key="comment.id" class="comment">-->
<!--          <p><strong>{{ comment.user }}：</strong> {{ comment.content }}</p>-->
<!--          <span>发布于：{{ formatDate(comment.posted_at) }}</span>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
<!--  </template>-->

<!--  <script setup>-->
<!--  import { ref } from 'vue';-->
<!--  import { ElTag } from 'element-plus';-->
<!--  -->
<!--  const blog = ref({-->
<!--    /* 您提供的JSON数据 */-->

<!--    "id": 12345,-->
<!--    "title": "如何有效管理个人时间",-->
<!--    "created_at": "2023-04-01T10:30:00Z",-->
<!--    "tags": ["时间管理", "效率提升", "生活黑客"],-->
<!--    "author": "时光导师",-->
<!--    "content": "在这个快节奏的时代，有效管理个人时间变得尤为重要。首先，设定明确的目标，无论是短期还是长期的，都能帮助你保持焦点。其次，合理规划每一天，使用时间块技术分配任务，确保工作与休息的平衡。此外，学会说“不”，避免无关紧要的事务侵占你的时间。最后，利用技术工具，如日历应用和待办事项列表，来辅助你的日常安排。",-->
<!--    "views": 1235,-->
<!--    "likes": 567,-->
<!--    "comments": [-->
<!--      {-->
<!--        "id": 987,-->
<!--        "user": "小李子",-->
<!--        "content": "非常实用的建议！特别是时间块技术，我试了之后效率确实提高了。",-->
<!--        "posted_at": "2023-04-02T15:45:00Z"-->
<!--      },-->
<!--      {-->
<!--        "id": 988,-->
<!--        "user": "时间旅者",-->
<!--        "content": "同意楼上，我也开始尝试设定每日三件事，感觉生活更有条理了。",-->
<!--        "posted_at": "2023-04-03T08:00:00Z"-->
<!--      }-->
<!--    ]-->

<!--  });-->
<!--  -->
<!--  // 格式化日期函数-->
<!--  const formatDate = (dateString) => {-->
<!--    const date = new Date(dateString);-->
<!--    return `${date.getFullYear()}-${('0' + (date.getMonth() + 1)).slice(-2)}-${('0' + date.getDate()).slice(-2)} ${('0' + date.getHours()).slice(-2)}:${('0' + date.getMinutes()).slice(-2)}`;-->
<!--  };-->
<!--  -->
<!--  // 示例函数，实际上在生产环境中您可能需要处理逻辑来删除标签-->
<!--  const removeTag = (tag) => {-->
<!--    console.log(`Tag "${tag}" would be removed.`);-->
<!--  };-->
<!--  </script>-->
<!--  -->
<!--  <style scoped>-->
<!--  /* 在此处添加自定义样式 */-->
<!--  .blog-detail {-->
<!--    padding: 2rem;-->
<!--    /* background-color: aqua; */-->
<!--  }-->
<!--  .meta-info {-->
<!--    margin-bottom: 1rem;-->
<!--  }-->
<!--  .stats span {-->
<!--    margin-right: 1rem;-->
<!--  }-->
<!--  .comment {-->
<!--    margin-bottom: 1rem;-->
<!--  }-->
<!--  </style>-->

<template>
  <div class="blog-detail">
    <h1>{{ blog.title }}</h1>
    <div class="blog-info">
      <p>发布时间：{{ blog.created_at }}</p>
      <p>浏览次数：{{ blog.view_count }}</p>
    </div>
    <div id="content" class="blog-content"></div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import Vditor from 'vditor';
import 'vditor/dist/index.css';

// JSON 数据
const blog = {
  "id": 46943,
  "category_id": 6,
  "user_id": 1,
  "title": "golang中使用装饰器技巧",
  "sort": 0,
  "brief": "",
  "thumb": "",
  "tags": "null",
  "referer": "",
  "status": 0,
  "content": "**很多场景需求， 是需要在函数调用时候， 使用装饰器的，方便在函数使用前和使用后，做一些业务操作。 比如实现一个函数计时。**\r\n\r\n1. python 中实现很简单。\r\n\r\n```python\r\nimport time\r\n\r\ndef timeit(func):\r\n    def wrapper(*args, **kwargs):\r\n        start_time = time.time()  # 记录函数开始执行的时间\r\n        result = func(*args, **kwargs)  # 执行函数\r\n        end_time = time.time()  # 记录函数结束执行的时间\r\n        print(f\"{func.__name__} took {end_time - start_time:.4f} seconds\")\r\n        return result\r\n    return wrapper\r\n\r\n# 使用装饰器\r\n@timeit\r\ndef example_function(n):\r\n    # 一个示例函数，计算从 1 到 n 的和\r\n    return sum(range(1, n + 1))\r\n\r\n# 调用示例函数\r\nresult = example_function(10000)\r\nprint(f\"Result: {result}\")\r\n```\r\n\r\n2.golang 中， 实现方案也不少，比如常见的为装饰器模式，但是由于参数比较固定（可用范型，或interface,不过还是恶心）导致不同类型的func都需要额外工作负担。\r\n\r\n```go\r\npackage main\r\n\r\nimport (\r\n\t\"fmt\"\r\n\t\"time\"\r\n)\r\n\r\n// exampleFunction 计算从0到n（包括n）的所有整数的和，并返回结果。\r\n// 参数：\r\n//\r\n//\tn int - 要计算的整数的上限（包括该整数）。\r\n//\r\n// 返回值：\r\n//\r\n//\tint - 从0到n（包括n）的所有整数的和。\r\nfunc exampleFunction(n int) int {\r\n\tresult := 0\r\n\tfor i := 0; i <= n; i++ {\r\n\t\tresult += i\r\n\t}\r\n\treturn result\r\n}\r\n\r\ntype example func(int) int\r\n\r\n// wrapTime 是一个包装函数，用于给传入的 example 类型的函数 efun 加上时间统计的功能\r\n// 参数 efun: 待包装的函数，类型为 example\r\n// 返回值: 返回包装后的函数，类型为 example\r\nfunc wrapTime(efun example) example {\r\n\tfn := func(n int) int {\r\n\t\tstart := time.Now()\r\n\t\tdefer func() {\r\n\t\t\tfmt.Println(time.Since(start))\r\n\t\t}()\r\n\t\treturn efun(n)\r\n\t}\r\n\treturn fn\r\n}\r\n\r\nfunc main() {\r\n\twrappedFunc := wrapTime(exampleFunction)\r\n\tfmt.Println(wrappedFunc(10000))\r\n}\r\n```\r\n\r\n3. 第三种奇巧淫技（有点丑，但很好用）， 虽然大家都知道defer， 但是只能作为return函数前使用。在进入函数的时候并不会执行， 如果写在defer外边， 又会污染函数。这个方法的原因是defer作为一个栈，在压栈的时候是会执行部分操作。在此操作时候进行响应逻辑即可。\r\n   ```go\r\n   package main\r\n\r\n   import (\r\n   \t\"fmt\"\r\n   \t\"time\"\r\n   )\r\n\r\n   func exampleFunction(n int) int {\r\n   \tdefer func() func() {\r\n   \t\tstart := time.Now()\r\n   \t\treturn func() {\r\n   \t\t\tfmt.Println(\"执行时间:\", time.Since(start))\r\n   \t\t}\r\n   \t}()()\r\n   \tresult := 0\r\n   \tfor i := 0; i <= n; i++ {\r\n   \t\tresult += i\r\n   \t}\r\n   \treturn result\r\n   }\r\n\r\n   func main() {\r\n   \tfmt.Println(exampleFunction(10000))\r\n   }\r\n\r\n   ```\r\n\r\n诸君可以自己试试。\r\n",
  "view_count": 93,
  "reply_count": 1,
  "zan_count": 0,
  "cai_count": 0,
  "created_at": "2024-06-25 08:46:21",
  "updated_at": "2024-07-20 05:03:42"
};

onMounted(() => {
  const vditor = new Vditor('content', {
    mode: 'preview',
    height: '100%',
    cache: { enable: false },
    preview: {
      markdown: {},
    },
    after: () => {
      vditor.setValue(blog.content);
    }
  });
});
</script>

<style scoped>
.blog-detail {
  padding: 20px;
}

.blog-info {
  margin-bottom: 20px;
  color: #666;
}

.blog-content {
  background-color: #fff;
  border: 1px solid #ddd;
  padding: 20px;
  border-radius: 4px;
}
</style>
