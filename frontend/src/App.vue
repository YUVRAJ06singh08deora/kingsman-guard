<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppHeader from './components/AppHeader.vue'
</script>

<template>
  <div class="w-screen h-screen">
    <app-header></app-header>
    <div class="container mx-auto font-inter w-full h-[85%]">
      <RouterView />
    </div>
  </div>
</template>

<style scoped></style>
