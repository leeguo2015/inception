<template>
  <div class="icp-record">
    <p>网站备案号：<a :href="recordLink" target="_blank">{{ recordNumber }}</a></p>
  </div>
</template>

<script>
export default {
  name: 'IcpRecord',
  props: {
    recordNumber: {
      type: String,
      required: true,
    },
    recordLink: {
      type: String,
      default: 'http://beian.miit.gov.cn/',
    },
  },
};
</script>

<style scoped>
.icp-record {
  position: fixed;
  bottom: 20px; /* 距离底部 20px */
  right: 20px; /* 距离右侧 20px */
  background-color: rgba(255, 255, 255, 0.8); /* 半透明背景 */
  padding: 10px 15px;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  font-size: 14px;
  color: #666;
  z-index: 1000; /* 确保它在最前面 */
}

.icp-record a {
  color: #333;
  text-decoration: none;
}

.icp-record a:hover {
  text-decoration: underline;
}
</style>
