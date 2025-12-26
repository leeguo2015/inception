<template>
  <div class="editor-container">
    <el-card class="editor-card">
      <el-form label-position="top" class="editor-form">
        <el-form-item label="标题">
          <el-input v-model="title" placeholder="填写标题" />
        </el-form-item>

        <el-row :gutter="16">
          <el-col :xs="24" :sm="12">
            <el-form-item label="标签">
              <div class="tag-input-wrapper">
                <el-input 
                  v-model="newTag" 
                  placeholder="输入标签后按回车添加"
                  @keyup.enter="addTag"
                  size="small"
                >
                  <template #append>
                    <el-button @click="addTag" size="small">添加</el-button>
                  </template>
                </el-input>
                <div class="tags-display">
                  <el-tag
                    v-for="tag in tags"
                    :key="tag"
                    closable
                    size="small"
                    @close="removeTag(tag)"
                    style="margin-right: 8px; margin-top: 8px;"
                  >
                    {{ tag }}
                  </el-tag>
                </div>
              </div>
            </el-form-item>
          </el-col>
          <el-col :xs="24" :sm="12" class="meta-col">
            <div class="meta-info">
              <div v-if="isEdit" class="meta-line">编辑中： {{ id }}</div>
            </div>
          </el-col>
        </el-row>

        <el-form-item label="内容">
          <div class="editor-wrapper">
            <QuillEditor v-model="content" class="quill-editor" />
          </div>
        </el-form-item>

        <div class="actions">
          <el-button @click="cancel">取消</el-button>
          <el-button v-if="isEdit" type="danger" @click="remove">删除</el-button>
          <el-button type="primary" @click="save">保存</el-button>
        </div>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { QuillEditor } from '@vueup/vue-quill'
import '@vueup/vue-quill/dist/vue-quill.snow.css'
import { useBlogStore } from '@/stores/blog'

const route = useRoute()
const router = useRouter()
const blog = useBlogStore()

const id = route.params.id || null
const isEdit = !!id

const title = ref('')
const content = ref('')
const newTag = ref('')
const tags = ref([])

onMounted(async () => {
  if (isEdit) {
    const p = await blog.fetchById(id)
    if (p) {
      title.value = p.title || ''
      content.value = p.content || ''
      tags.value = p.tags || []
    }
  }
})

async function save() {
  if (isEdit) {
    await blog.update(id, { title: title.value, content: content.value, tags: tags.value })
    router.push({ name: 'BlogView', params: { id } })
  } else {
    const newPost = await blog.add({ title: title.value, content: content.value, tags: tags.value })
    const newId = newPost && (newPost.id || newPost.blog_id || newPost.data?.id)
    router.push({ name: 'BlogView', params: { id: newId || '' } })
  }
}

function cancel() {
  router.back()
}

async function remove() {
  await blog.remove(id)
  router.push({ name: 'BlogList' })
}

// 标签操作方法
function addTag() {
  const tag = newTag.value.trim()
  if (tag && !tags.value.includes(tag)) {
    tags.value.push(tag)
    newTag.value = ''
  }
}

function removeTag(tag) {
  const index = tags.value.indexOf(tag)
  if (index > -1) {
    tags.value.splice(index, 1)
  }
}
</script>

<style scoped>
.editor-container {
  width: 100%;
  display: flex;
  justify-content: center;
  padding: 16px;
}
.editor-card {
  width: 100%;
  max-width: 980px;
  box-shadow: 0 6px 18px rgba(0,0,0,0.06);
  border-radius: 8px;
}
.editor-form {
  padding: 6px 4px;
}
.editor-wrapper {
  min-height: 65vh;
  min-width: 100%;
  border: 1px solid var(--el-border-color, #e6e6e6);
  border-radius: 6px;
  overflow: hidden;
  background: #fff;
}
.quill-editor {
  height: 65vh;
  min-height: 500px;
  display: flex;
  flex-direction: column;
}
/* 调整 quill 内部容器以撑满高度 */
.quill-editor .ql-container {
  flex: 1;
  min-height: 0; /* 允许容器收缩 */
}
:deep(.ql-editor) {
  min-height: 65vh;
  border: none !important;
  outline: none !important;
  box-shadow: none !important;
}
/* toolbar 和编辑区圆角/边框处理 */
.quill-editor .ql-toolbar {
  border-bottom: 1px solid var(--el-border-color, #e6e6e6);
  background: #fafafa;
}
.actions {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 12px;
}
.meta-col {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}
.meta-info .meta-line {
  color: #909399;
  font-size: 13px;
}

/* 标签输入样式 */
.tag-input-wrapper {
  width: 256px;

}

.tags-display {
  margin-top: 8px;
  min-height: 40px;
  /* border: 1px solid #e6e6e6; */
  border-radius: 4px;
  padding: 8px;
  /* background: #f9f9f9; */
}
@media (max-width: 576px) {
  .quill-editor { height: 320px; }
  .editor-card { padding: 8px; }
}
</style>