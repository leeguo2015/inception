<script setup>
import { RouterLink, RouterView } from 'vue-router'
import Home from './components/Home.vue'

</script>

<template >
  <div id="main">
    <Home class="child" />
    <router-link to="/">Go to Home</router-link>
    <router-link to="/about">Go to About</router-link>
    <router-link to="/login">Login</router-link>
  </div>

  <router-view></router-view>
</template>

<style scoped>

#main {
  /* display: flex; */
  height: 100%; 
  width: 100%;
  background-color: gray;

}


</style>
