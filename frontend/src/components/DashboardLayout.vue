<template>
  <div class="ml-5 py-5 mx-auto flex">
    <app-sidebar></app-sidebar>

    <div class="p-5 ml-5 border border-gray-300 bg-primary rounded-lg w-[75%] h-auto">
      <slot></slot>
    </div>
  </div>
</template>

<script lang="ts">
import AppSidebar from './AppSidebar.vue'

export default {
  components: {
    AppSidebar,
  },
}
</script>
