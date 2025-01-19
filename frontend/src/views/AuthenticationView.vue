<template>
  <LoadingSpinner />
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'
import { userManager } from '@/auth/oidc-config'
import { useRouter } from 'vue-router'
import LoadingSpinner from '@/components/LoadingSpinner.vue'

export default {
  name: 'AuthCallback',
  components: {
    LoadingSpinner,
  },
  setup() {
    const router = useRouter()
    const loading = ref(true)

    const handleCallback = async () => {
      try {
        const loggedInUser = await userManager.signinRedirectCallback()

        if (loggedInUser.access_token && loggedInUser.id_token && loggedInUser.refresh_token) {
          localStorage.setItem('access_token', loggedInUser.access_token)
          localStorage.setItem('id_token', loggedInUser.id_token)
          localStorage.setItem('refresh_token', loggedInUser.refresh_token)

          router.push('/gamedashboard')
        } else {
          console.error('Missing token(s) after login callback')
          router.push('/')
        }
      } catch (error) {
        console.error('Error handling login callback:', error)
        localStorage.removeItem('access_token')
        localStorage.removeItem('id_token')
        localStorage.removeItem('refresh_token')
        router.push('/')
      } finally {
        loading.value = false
      }
    }
    onMounted(() => {
      handleCallback()
    })

    return {
      loading,
    }
  },
}
</script>
