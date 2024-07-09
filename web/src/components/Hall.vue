<template>
  <div class="hall-main">
    <div class="common-layout">
      <el-container>
        <el-aside id="broadside">
          <Broadside @categoriesIDs="debouncedUpdateData" :categories="categories"/>
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
import { debounce } from 'lodash';
import Broadside from './Broadside.vue';
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
  get("blog").then((res) => {
    categories.value = res.data.list;
  });
}

onMounted(() => {
  getCategory()
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
  overflow: auto;
  /* display: flex; */
}

#broadside {
  width: 180px;
}
</style>
