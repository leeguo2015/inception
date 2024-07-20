<template>
  <div class="broadside-main">
    <div v-for="tag in categories" :key="tag.id">
      <div class="broad-tag">
        <el-tag size="small" 
          :type="activeCategoriesIDs.includes(tag.id) ? 'success' : 'info'"
          @click="addActive(tag.id)">
          {{ tag.name }}
        </el-tag>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { defineEmits } from 'vue';
const emits = defineEmits(['activeCategoriesIDs']);
const props = defineProps({
  categories: {
    type: Array,
    required: true
  }
});

let activeCategoriesIDs = ref([]);
const addActive = (id) => {
  const index = activeCategoriesIDs.value.indexOf(id);
  if (index > -1) {
    activeCategoriesIDs.value.splice(index, 1);
  } else {
    activeCategoriesIDs.value.push(id);
  }
  emits('activeCategoriesIDs', activeCategoriesIDs.value);
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
