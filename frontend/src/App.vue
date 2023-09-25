<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavBar from './components/NavBar.vue'
import ErrorsBlock from "@/components/ErrorsBlock.vue";
</script>

<template>
  <NavBar/>

  <RouterView />

  <ErrorsBlock/>
</template>