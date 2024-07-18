<template>
  <div class="hall-main">
    <div class="common-layout">
      <el-container >
        <el-aside id="broadside">
          <Broadside @categoriesIDs="debouncedUpdateData" :categories="categories"/>
        </el-aside>
        <el-main >
         <div id="blog-list">
          <list :sharedData="debouncedSharedData" />

         </div>

        </el-main>
        <el-footer>
          <!-- 底部页码组件 -->
<!--          <PageFooter />-->
<!--          <el-pagination-->
<!--              :page-size="20"-->
<!--              :pager-count="11"-->
<!--              layout="prev, pager, next"-->
<!--              :total="1000"-->
<!--          />-->
        </el-footer>

      </el-container>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { debounce } from 'lodash';
import Broadside from './broadside.vue';
import list from './blog/list.vue';
import { get} from '@/assets/api';
import { onMounted } from 'vue'

const categories = ref([]);
const sharedData = ref([]);
const debouncedSharedData = ref([]);

const getCategory = () => {
  get("category").then((res) => {
    categories.value = res.data.list;
  });
}
const getBlogList = () => {
  get("aritcle").then((res) => {
    debouncedSharedData.value = res.data.list;
  });
}

onMounted(() => {
  getCategory()
  getBlogList()
})

let backendData = []
debouncedSharedData.value = backendData;

const handleUpdateData = (newData) => {
  sharedData.value = backendData;
};

const debouncedUpdateData = debounce((newData) => {
  handleUpdateData(newData);
  // getBlogList()
}, 1000); // 1000ms 的防抖时间

</script>
<style>
.hall-main {
  //overflow: auto;
  /* display: flex; */
}

#broadside {
  width: 180px;
}
#blog-list{
  //height: calc(100vh - 60px);
  //padding-bottom: 10vh;
}
.el-main{
  padding: 0 !important;
}
.content-wrap {
  display: flex;
  //height: calc(100vh - /* header和footer的高度 */);
  height: calc(100vh - 70px);
}
.el-main {
  overflow-y: auto; /* 允许内容区域滚动 */
}
</style>
