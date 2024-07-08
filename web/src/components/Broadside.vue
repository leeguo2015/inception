<template>
  <div class="broadside-main">
    <div v-for="tag in categories" :key="tag.id">
      <div class="broad-tag">
        <el-tag size="small" 
          :type="activeList.includes(tag.id) ? 'success' : 'info'"
          @click="addActive(tag.id)">
          {{ tag.name }}
        </el-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
  categories: {
    type: Array,
    required: true
  }
});

let activeList = ref([]);

const addActive = (id) => {
  const index = activeList.value.indexOf(id);
  if (index > -1) {
    activeList.value.splice(index, 1);
  } else {
    activeList.value.push(id);
  }
};
</script>

<style>
.broadside-main {
  margin: 40px 20px;
  padding: 1em;
  min-height: 30vh;
}

.broad-tag {
  margin-top: 5px;
}
</style>
