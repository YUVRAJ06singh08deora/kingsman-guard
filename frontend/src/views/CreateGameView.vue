<template>
  <div class="p-4 flex justify-center items-center">
    <div
      class="w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700"
    >
      <form class="space-y-6" @submit.prevent="createGame">
        <h5 class="text-xl font-medium text-gray-900 dark:text-white">
          Enter the Number of Rounds you want to keep in the Game
        </h5>
        <div>
          <input
            v-model="rounds"
            type="number"
            name="rounds"
            id="rounds"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            placeholder="Number of Rounds"
            required
          />
        </div>

        <button
          type="submit"
          class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        >
          Create Game
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios'
import { getAccessToken } from '@/auth/auth-service'

export default {
  name: 'JoinGameView',
  data() {
    return {
      rounds: null,
      playerID: '',
      playerName: '',
      gameCode: '',
    }
  },
  methods: {
    // Method to create a new game by calling the backend API
    async createGame() {
      if (!this.rounds || this.rounds <= 0) {
        console.error('Please enter a valid number of rounds')
        return
      }

      const backendUrl = import.meta.env.VITE_BACKEND_URL

      const accessToken = await getAccessToken()

      if (!accessToken) {
        console.error('Access token is missing')
        return
      }
      // Try to create a new game
      try {
        // Prepare the headers with Authorization
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        // Prepare the request body with the number of rounds
        const body = {
          number_of_rounds: this.rounds,
        }

        // Make a POST request to create the game with the rounds
        const response = await axios.post(
          `${backendUrl}/create_game`,
          body, // Passing the body with rounds
          { headers }, // Including the headers with Authorization
        )

        // Check if the game creation was successful and retrieve the game code
        if (response.data.success && response.data.data && response.data.data.game_code) {
          const gameCode = response.data.data.game_code
          this.gameCode = gameCode
          // Redirect to the /game route with the game_code as a query parameter
          this.$router.push({ path: `/game`, query: { game_code: gameCode } })
          this.joinGame(gameCode)
        } else {
          console.error('Game creation failed or game code is missing.')
        }
      } catch (error) {
        console.error('Error creating game:', error)
      }
    },
    
    async joinGame(gameCode: string) {
      const backendUrl = import.meta.env.VITE_BACKEND_URL

      // Ensure playerid and playerName are populated before proceeding
      if (!this.playerID) {
        console.error('Player ID is missing')
        return
      }

      const accessToken = await getAccessToken()

      if (!accessToken) {
        console.error('Access token is missing')
        return
      }

      // Try to Join the game
      try {
        // Prepare the headers with Authorization
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        const body = {
          game_code: gameCode,
          player_id: this.playerID,
          player_name: this.playerName,
        }

        // Make a POST request to create the game with the rounds
        const response = await axios.post(
          `${backendUrl}/join_game`,
          body, // Passing the body with rounds
          { headers }, // Including the headers with Authorization
        )

        // Check if the game creation was successful and retrieve the game code
        if (response.data.message) {
          console.log(response.data.message)
        } else {
          console.error('Not able to join the game')
        }
      } catch (error) {
        console.error('Error creating game:', error)
      }
    },
    async fetchUserNameAndId() {
      // Retrieve the access token from localStorage
      const accessToken = await getAccessToken()

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

        type UserAttribute = {
          Name: string
          Value: string
        }

        // Ensure UserAttributes exists in the response
        if (response.data && response.data.UserAttributes) {
          const userAttributes: UserAttribute[] = response.data.UserAttributes

          const givenName =
            userAttributes.find((attr: UserAttribute) => attr.Name === 'given_name')?.Value ||
            'Anonymous'

          const emailId =
            userAttributes.find((attr: UserAttribute) => attr.Name === 'email')?.Value ||
            'no-email@example.com'

          this.playerName = givenName
          this.playerID = emailId 
          console.log('Player Name:', givenName)
          console.log('Player ID:', emailId)
        } else {
          console.error('User attributes not found in response', response.data)
        }
      } catch (error) {
        console.error('Error fetching user details:', error)
      }
    },
  },
  
  async mounted() {
    await this.fetchUserNameAndId()  // This will run when the component is mounted
  },
}
</script>


<style scoped>
/* You can add specific styles here if needed */
</style>
