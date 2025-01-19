<template>
  <div class="flex justify-center items-center w-full h-full">
    <div
      class="w-[30rem] h-[15rem] border border-gray-300 bg-primary flex flex-col justify-center items-center rounded-xl"
    >
      <h1 class="text-lg font-bold pb-6">Welcome to Kingsmen Guard</h1>
      <button
        class="bg-purple-500 text-white font-semibold w-20 h-9 rounded-md hover:bg-purple-700"
        @click="login"
      >
        Login
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import { login } from '@/auth/auth-service'

export default defineComponent({
  name: 'HomeView',
  methods: {
    login() {
      login()
    },
  },
})
</script>

<style scoped></style>
