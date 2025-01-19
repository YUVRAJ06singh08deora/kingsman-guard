<template>
  <div class="p-4 flex justify-center items-center">
    <div
      class="w-full max-w-sm bg-white border border-gray-200 rounded-lg shadow dark:bg-slate-300 dark:border-gray-700"
    >
      <div class="flex flex-col items-center pb-10">
        <iframe
          src="https://lottie.host/embed/984a7979-ea3d-4590-85ca-f8f729341f2f/whx4znowUR.lottie"
          class="w-full h-48"
        ></iframe>

        <!-- Displaying the logged-in user's name dynamically -->
        <h5 class="mb-1 text-xl font-medium text-gray-1500 dark:text-black">{{ userName }}</h5>

        <div class="flex mt-4 md:mt-6">
          <!-- Button to create a new game -->
          <button
            @click="redirectToCreateGame"
            class="inline-flex items-center px-4 py-2 text-sm font-medium text-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          >
            Create Game
          </button>
          <a
            @click="redirectToJoinGame"
            class="py-2 px-4 ms-2 text-sm font-medium text-gray-900 focus:outline-none bg-white rounded-lg border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700"
            >Join a Game</a
          >
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios'

export default {
  name: 'GameDashView',
  data() {
    return {
      userName: '', // Will hold the user's name
    }
  },
  mounted() {
    this.fetchUserName()
  },
  methods: {
    // Method to fetch the user's name from AWS Cognito
    async fetchUserName() {
      // Retrieve the access token from localStorage
      const accessToken = localStorage.getItem('access_token')

      if (!accessToken) {
        console.error('Access token is missing')
        return
      }

      // Prepare the request headers and body
      const headers = {
        'Content-Type': 'application/x-amz-json-1.1',
        'X-Amz-Target': 'AWSCognitoIdentityProviderService.GetUser',
        Authorization: `Bearer ${accessToken}`,
      }

      const body = {
        AccessToken: accessToken,
      }

      try {
        const response = await axios.post('https://cognito-idp.us-east-1.amazonaws.com/', body, {
          headers,
        })

        // Define the type for the UserAttributes
        type UserAttribute = {
          Name: string
          Value: string
        }

        // Find the "given_name" attribute in the response and set it
        const userAttributes: UserAttribute[] = response.data.UserAttributes
        const givenName = userAttributes.find(
          (attr: UserAttribute) => attr.Name === 'given_name',
        )?.Value

        if (givenName) {
          this.userName = givenName
        } else {
          console.error('Given name not found in response')
        }
      } catch (error) {
        console.error('Error fetching user details:', error)
      }
    },

   // Method to redirect the user to the Create Game page
   redirectToCreateGame() {
      // Redirect to /creategame page
      this.$router.push({ path: '/creategame' })
    },

    // Method to redirect the user to the Join Game page
    redirectToJoinGame() {
      // Redirect to /creategame page
      this.$router.push({ path: '/joingame' })
    },
    
  },
}
</script>

<style scoped>
/* You can add specific styles here if needed */
</style>
