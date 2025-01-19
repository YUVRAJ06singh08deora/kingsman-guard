<template>
  <div class="p-4 flex justify-center items-center">
    <div
      class="w-full max-w-sm p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8 dark:bg-gray-800 dark:border-gray-700"
    >
      <form class="space-y-6" @submit.prevent="joinGame">
        <h5 class="text-xl font-medium text-gray-900 dark:text-white">Enter the Game Code</h5>
        <div>
          <label for="gamecode" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Game Code</label>
          <input
            v-model="gameCode" 
            type="text"
            name="gamecode"
            id="gamecode"
            class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
            placeholder="Enter the Game code received from the host"
            required
          />
        </div>

        <button
          type="submit"
          class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
        >
          Join Game
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios'
import { getAccessToken } from '@/auth/auth-service';

export default {
  name: 'JoinGameView',
  data() {
    return {
      gameCode: '',
      playerID: '', 
      playerName: '',
    }
  },
  methods: {
    async joinGame() {
      const backendUrl = import.meta.env.VITE_BACKEND_URL;

      // Ensure playerId and gameCode are populated before proceeding
      if (!this.gameCode) {
        console.error('Game code is missing');
        return;
      }

      if (!this.playerID) {
        console.error('Player ID is missing');
        return;
      }

      const accessToken = await getAccessToken();

      if (!accessToken) {
        console.error('Access token is missing');
        return;
      }

      try {
        // Prepare the headers with Authorization
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        };

        const body = {
          game_code: this.gameCode, // Use this.gameCode from the data
          player_id: this.playerID,
          player_name: this.playerName,
        };

        // Make a POST request to join the game
        const response = await axios.post(
          `${backendUrl}/join_game`,
          body, 
          { headers }
        );

        // Check if the game join was successful
        if (response.data.message) {
          console.log(response.data.message);

          // Redirect to the game page with the gameCode as a query parameter
          this.$router.push({ name: 'game', query: { game_code: this.gameCode } });
        } else {
          console.error('Not able to join the game');
        }
      } catch (error) {
        console.error('Error joining game:', error);
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
    await this.fetchUserNameAndId();
  },
}
</script>

<style scoped></style>
