<!--
 * @Author: leeguo leeguo2015@163.com
 * @Date: 2023-11-04 14:48:05
 * @LastEditors: leeguo leeguo2015@163.com
 * @LastEditTime: 2023-11-24 22:34:13
 * @FilePath: \inception\web\src\components\menu.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="hall-main">
    <div class="common-layout">
      <el-container>
        <el-aside id="broadside">
          <Broadside @categoriesIDs="debouncedUpdateData" :categories:="categories" />
        </el-aside>
        <el-main>
          <list :sharedData="debouncedSharedData" />
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
// import { useRoute } from 'vue-router';
import { debounce } from 'lodash';
import Broadside from './Broadside.vue';
import list from './blog/list.vue';
import { get} from '@/assets/api';
import { onMounted } from 'vue'



// const route = useRoute();
const categories = ref([]);
const sharedData = ref([]);
const debouncedSharedData = ref([]);
let backendData = [
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
    "content": `如果在实际运行中发现防抖未生效，可能的原因有：
确保事件触发：确认子组件Broadside确实有在适当的时候触发TagIDList事件，并且传递了预期的数据。
Lodash的导入：确保lodash库已经正确导入到项目中。如果使用的是Vue CLI或Vite等构建工具，需要确认lodash已经被安装并且在你的项目中正确导入。
作用域问题：虽然从代码上看不太可能出现这个问题，但理论上如果debounce调用不在正确的作用域内，可能导致handleUpdateData方法没有按预期工作。不过，基于您提供的代码片段，这一点看起来是正确的。
异步更新队列：Vue 3中的响应式系统有时会在下一个微任务队列中更新视图。尽管这不会直接影响防抖功能，但在调试时可能会造成一些混淆，例如立即在handleUpdateData之后检查数据或UI可能还未看到变化。
如果以上都检查无误，而防抖仍然未按预期工作，建议使用浏览器的开发者工具进行详细调试，观察事件触发和防抖函数调用的实际时机，以及数据变化情况，来定位问题所在。`,
    "datePublished": "2023-04-01"
  },
  // 更多文章...  
]

categories.value = [
    {
        "id": 2,
        "name": "BUG反馈",
        "content": "",
        "brief": ""
    },
    {
        "id": 5,
        "name": "AskCat2",
        "content": "",
        "brief": ""
    }
]
debouncedSharedData.value = backendData;

const handleUpdateData = (newData) => {
  sharedData.value = backendData;
};

const debouncedUpdateData = debounce((newData) => {
  handleUpdateData(newData);
  debouncedSharedData.value.push({
    "id": Date.now(),
    "title": "Vue 5 ",
    "author": "John Doe",
    "content": "Vue 3 带来了许多新特性...",
    "datePublished": "2023-04-01"
  },);
  console.log('debouncedSharedData.value', debouncedSharedData.value);
}, 1000); // 1000ms 的防抖时间

const getCategory = () => {
  get("category").then((res) => {
    console.log('category', res.data);
    // categoryList.value = res.data;

  });
}
onMounted(() => {
  getCategory()
})

</script>
<style>
.hall-main {
  overflow: auto;
  /* display: flex; */
}

#broadside {
  width: 180px;
}
</style>
