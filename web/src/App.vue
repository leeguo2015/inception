<template>

    <el-container>
      <el-header>
          <Menu></Menu>
      </el-header>
      <el-container>
        <el-aside width="200px" >
          <Broadside @categoriesIDs="debouncedUpdateData" :categories="categories"/>
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
      </el-container>
    </el-container>
</template>

<script setup>
import Menu from './components/Menu.vue'
import { onMounted } from 'vue'
import { useStore } from 'vuex' // 引入useStore 方法
import { ref } from 'vue';
import { debounce } from 'lodash';
import Broadside from '@/components/broadside.vue';
import list from '@/components/blog/list.vue';
import { get} from '@/assets/api';

const categories = ref([]);
const blogTotal = ref(0);
const currentPage = ref(1);
const sharedData = ref([]);
const debouncedSharedData = ref([]);
const pageSize = 10

const getCategory = () => {
  get("category").then((res) => {
    categories.value = res.data.list;
  });
}
const getBlogList = (newPage) => {
  currentPage.value = newPage;
  const params = {
    size: pageSize,
    page: newPage
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
  sharedData.value = backendData;
};

const debouncedUpdateData = debounce((newData) => {
  handleUpdateData(newData);
  // getBlogList()
}, 1000); // 1000ms 的防抖时间

// console.log(store.state.user)  // store 实例对象
// function checkUser() {
//   const store = useStore();
//   const user = localStorage.getItem('user') || sessionStorage.getItem('user');
//   // console.log("store.state.user:", store.state.user)
//   // console.log("user", user)
//   if (user) {
//     store.commit('SET_USER', user);
//   }
// }

// onMounted(() => {
//   checkUser()
//
// })
</script>

<style scoped>
#footer{
  display: flex;
  justify-content: center; /* 水平居中 */;
}
#container{
  height: calc(100vh - 60px);
}
</style>
