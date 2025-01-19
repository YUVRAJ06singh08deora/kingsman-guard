<script lang="ts">
import { logout } from '@/auth/auth-service'

export default {
  name: 'AppHeader',
  methods: {
    async handleLogout() {
      localStorage.removeItem('access_token')
      localStorage.removeItem('id_token')
      localStorage.removeItem('refresh_token')
      await logout()
    },
  },
}
</script>

<template>
  <header class="w-full font-inter bg-primary border border-gray-300 py-4 relative">
    <div class="container mx-auto w-full">
      <div class="absolute mx-4 pt-0.5">
        <!-- Logout button -->
        <button @click="handleLogout" class="flex justify-center items-center text-gray-500">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-5 mr-1"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0 3-3m0 0-3-3m3 3H9"
            />
          </svg>

          Logout
        </button>
      </div>
      <div class="flex justify-center items-center">
        
        <h1 class="text-lg md:text-xl font-semibold mx-2">Kingsman Guard</h1>
      </div>
    </div>
  </header>
</template>

<style scoped></style>
