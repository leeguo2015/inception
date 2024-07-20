<script setup>
import { onMounted } from 'vue'
import { ref } from 'vue';
import { debounce } from 'lodash';
import Broadside from '@/components/Broadside.vue';
import list from '@/components/blog/list.vue';
import { get} from '@/assets/api';
import { useRouter, useRoute } from "vue-router";

const categories = ref([]);
const blogTotal = ref(0);
const currentPage = ref(1);
const sharedData = ref([]);
const debouncedSharedData = ref([]);
const CategoryIDs = ref([]);
const pageSize = 10
const route = useRoute();
const router = useRouter();
const getCategory = () => {
  const params = {
    Type :route.params.blogType,
  };
  get("category", params).then((res) => {
    categories.value = res.data.list;
  });
}
const getBlogList = (newPage) => {
  currentPage.value = newPage;
  const params = {
    size: pageSize,
    page: newPage,
    Type :route.params.blogType,
    CategoryIDs:  `[${CategoryIDs.value.map(id => encodeURIComponent(id)).join(',')}]`
  };
  get("article",params).then((res) => {
    debouncedSharedData.value = res.data.list;
    blogTotal.value = res.data.total;
  });
}

onMounted(() => {
  getCategory()
  getBlogList()
})

let backendData = []
debouncedSharedData.value = backendData;

const handleUpdateData = (newData) => {
  CategoryIDs.value = newData;
  console.log(newData);
};

const debouncedUpdateData = debounce((newData) => {
  // console.log(newData);
  // handleUpdateData(newData);
  CategoryIDs.value = newData;
  getBlogList()
}, 500); // 1000ms 的防抖时间

</script>

<template>
  <el-aside width="200px" >
    <Broadside @activeCategoriesIDs="debouncedUpdateData" :categories="categories"/>
  </el-aside>
  <el-container id="container">

    <el-main>
      <div id="blog-list">
        <list :sharedData="debouncedSharedData" />
      </div>
    </el-main>
    <el-footer id="footer" v-if="blogTotal>0">
      <el-pagination
          :page-size="10"
          @update:current-page="getBlogList"
          :current-page="currentPage"
          layout="prev, pager, next"
          :total="blogTotal"
      />
    </el-footer>
  </el-container>

</template>

<style scoped>
#footer{
  display: flex;
  justify-content: center; /* 水平居中 */;
}
#container{
  height: calc(100vh - 60px);
}
</style>