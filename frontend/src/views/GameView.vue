<template>
  <div class="p-4 flex justify-between space-x-6">
    <!-- Left Section (Game Details, Leaderboard, etc.) -->
    <div class="w-1/3">
      <div class="bg-gray-800 p-6 rounded-lg shadow-lg text-white">
        <h2 class="text-xl font-bold mb-4">Game Code : {{ gameCode }}</h2>
        <!-- Display dynamic game code -->
        <h2 class="text-xl font-bold mb-4">Round Details</h2>
        <div class="mb-4">
          <span class="block text-sm opacity-75">Round Number:</span>
          <span class="block text-lg font-semibold">{{ roundNumber }}</span>
        </div>

        <h2 class="text-xl font-bold mb-4">Leaderboard</h2>
        <ul class="space-y-3">
          <!-- Loop through the leaderboard array and display player info -->
          <li v-for="(player, index) in leaderboard" :key="index" class="flex justify-between">
            <span class="text-sm">{{ player.player_name }}</span>
            <span class="font-semibold">{{ player.total_score }} Points</span>
          </li>
        </ul>
      </div>

      <!-- Container for buttons with margin-top to create space from leaderboard -->
      <div class="mt-6 flex flex-col items-center w-full">
        <!-- Start Game and End Game buttons centered and each takes half of the container width -->
        <div class="flex w-full space-x-2 mb-4">
          <button
            @click="startGame"
            type="button"
            class="focus:outline-none text-white bg-green-700 hover:bg-green-800 focus:ring-4 focus:ring-green-300 font-medium rounded-lg text-sm px-5 py-2.5 w-1/2"
          >
            Start Game
          </button>
          <button
            @click="endGame"
            type="button"
            class="focus:outline-none text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 w-1/2"
          >
            End Game
          </button>
        </div>

        <!-- Play button centered below Start/End Game buttons -->
        <button
          @click="playGame"
          type="button"
          class="focus:outline-none text-white bg-yellow-400 hover:bg-yellow-500 focus:ring-4 focus:ring-yellow-300 font-medium rounded-lg text-sm px-5 py-2.5 w-full max-w-xs"
        >
          Next Round
        </button>
      </div>
    </div>

    <!-- Right Section (Player Cards) -->
    <div class="w-2/3 bg-slate-50 rounded-lg flex flex-wrap justify-center space-x-6">
      <!-- Player 1 Card -->
      <div
        class="flex-shrink-0 m-6 relative overflow-hidden bg-red-500 rounded-lg max-w-xs shadow-lg group"
        @click="handleCardClick(players[0])"
        :class="{
          'cursor-pointer': players[0]?.role !== 'Sepoy',
          'cursor-not-allowed': players[0]?.role === 'Sepoy',
        }"
      >
        <svg
          class="absolute bottom-0 left-0 mb-8 scale-150 group-hover:scale-[1.65] transition-transform"
          viewBox="0 0 375 283"
          fill="none"
          style="opacity: 0.1"
        >
          <rect
            x="159.52"
            y="175"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 159.52 175)"
            fill="white"
          ></rect>
          <rect
            y="107.48"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 0 107.48)"
            fill="white"
          ></rect>
        </svg>
        <div
          class="relative pt-10 px-10 flex items-center justify-center group-hover:scale-110 transition-transform"
        >
          <div
            class="block absolute w-48 h-48 bottom-0 left-0 -mb-24 ml-3"
            style="
              background: radial-gradient(black, transparent 60%);
              transform: rotate3d(0, 0, 1, 20deg) scale3d(1, 0.6, 1);
              opacity: 0.2;
            "
          ></div>
          <iframe
            src="https://lottie.host/embed/984a7979-ea3d-4590-85ca-f8f729341f2f/whx4znowUR.lottie"
            class="w-full h-full"
          ></iframe>
        </div>
        <div class="relative text-white px-6 pb-6 mt-6">
          <span class="block opacity-75 -mb-1">{{ players[0]?.player_name || 'Player 1' }}</span>
          <div class="flex justify-between">
            <span class="block font-semibold text-xl">
              <span v-if="players[0]?.role === 'Sepoy' || players[0]?.player_id === userID || roleVisibility">
                {{ players[0]?.role || 'Role' }}
              </span>
              <span v-else>{{ '' }}</span>
            </span>
            <span
              class="block bg-white rounded-full text-orange-500 text-xs font-bold px-3 py-2 leading-none flex items-center"
            >
              {{ players[0]?.current_round_score || '0' }}*
            </span>
          </div>
        </div>
      </div>

      <!-- Player 2 Card -->
      <div
        class="flex-shrink-0 m-6 relative overflow-hidden bg-teal-500 rounded-lg max-w-xs shadow-lg group"
        @click="handleCardClick(players[1])"
        :class="{
          'cursor-pointer': players[1]?.role !== 'Sepoy',
          'cursor-not-allowed': players[1]?.role === 'Sepoy',
        }"
      >
        <svg
          class="absolute bottom-0 left-0 mb-8 scale-150 group-hover:scale-[1.65] transition-transform"
          viewBox="0 0 375 283"
          fill="none"
          style="opacity: 0.1"
        >
          <rect
            x="159.52"
            y="175"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 159.52 175)"
            fill="white"
          ></rect>
          <rect
            y="107.48"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 0 107.48)"
            fill="white"
          ></rect>
        </svg>
        <div
          class="relative pt-10 px-10 flex items-center justify-center group-hover:scale-110 transition-transform"
        >
          <div
            class="block absolute w-48 h-48 bottom-0 left-0 -mb-24 ml-3"
            style="
              background: radial-gradient(black, transparent 60%);
              transform: rotate3d(0, 0, 1, 20deg) scale3d(1, 0.6, 1);
              opacity: 0.2;
            "
          ></div>
          <iframe
            src="https://lottie.host/embed/984a7979-ea3d-4590-85ca-f8f729341f2f/whx4znowUR.lottie"
            class="w-full h-full"
          ></iframe>
        </div>
        <div class="relative text-white px-6 pb-6 mt-6">
          <span class="block opacity-75 -mb-1">{{ players[1]?.player_name || 'Player 2' }}</span>
          <div class="flex justify-between">
            <span class="block font-semibold text-xl">
              <span v-if="players[1]?.role === 'Sepoy' || players[1]?.player_id === userID || roleVisibility">
                {{ players[1]?.role || 'Role' }}
              </span>
              <span v-else>{{ '' }}</span>
            </span>
            <span
              class="block bg-white rounded-full text-teal-500 text-xs font-bold px-3 py-2 leading-none flex items-center"
            >
              {{ players[1]?.current_round_score || '0' }}*
            </span>
          </div>
        </div>
      </div>

      <!-- Player 3 Card -->
      <div
        class="flex-shrink-0 m-6 relative overflow-hidden bg-green-500 rounded-lg max-w-xs shadow-lg group"
        @click="handleCardClick(players[2])"
        :class="{
          'cursor-pointer': players[2]?.role !== 'Sepoy',
          'cursor-not-allowed': players[2]?.role === 'Sepoy',
        }"
      >
        <svg
          class="absolute bottom-0 left-0 mb-8 scale-150 group-hover:scale-[1.65] transition-transform"
          viewBox="0 0 375 283"
          fill="none"
          style="opacity: 0.1"
        >
          <rect
            x="159.52"
            y="175"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 159.52 175)"
            fill="white"
          ></rect>
          <rect
            y="107.48"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 0 107.48)"
            fill="white"
          ></rect>
        </svg>
        <div
          class="relative pt-10 px-10 flex items-center justify-center group-hover:scale-110 transition-transform"
        >
          <div
            class="block absolute w-48 h-48 bottom-0 left-0 -mb-24 ml-3"
            style="
              background: radial-gradient(black, transparent 60%);
              transform: rotate3d(0, 0, 1, 20deg) scale3d(1, 0.6, 1);
              opacity: 0.2;
            "
          ></div>
          <iframe
            src="https://lottie.host/embed/984a7979-ea3d-4590-85ca-f8f729341f2f/whx4znowUR.lottie"
            class="w-full h-full"
          ></iframe>
        </div>
        <div class="relative text-white px-6 pb-6 mt-6">
          <span class="block opacity-75 -mb-1">{{ players[2]?.player_name || 'Player 3' }}</span>
          <div class="flex justify-between">
            <span class="block font-semibold text-xl">
              <span v-if="players[2]?.role === 'Sepoy' || players[2]?.player_id === userID || roleVisibility">
                {{ players[2]?.role || 'Role' }}
              </span>
              <span v-else>{{ '' }}</span>
            </span>
            <span
              class="block bg-white rounded-full text-purple-500 text-xs font-bold px-3 py-2 leading-none flex items-center"
            >
              {{ players[2]?.current_round_score || '0' }}*
            </span>
          </div>
        </div>
      </div>

      <!-- Player 4 Card -->
      <div
        class="flex-shrink-0 m-6 relative overflow-hidden bg-orange-500 rounded-lg max-w-xs shadow-lg group"
        @click="handleCardClick(players[3])"
        :class="{
          'cursor-pointer': players[3]?.role !== 'Sepoy',
          'cursor-not-allowed': players[3]?.role === 'Sepoy',
        }"
      >
        <svg
          class="absolute bottom-0 left-0 mb-8 scale-150 group-hover:scale-[1.65] transition-transform"
          viewBox="0 0 375 283"
          fill="none"
          style="opacity: 0.1"
        >
          <rect
            x="159.52"
            y="175"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 159.52 175)"
            fill="white"
          ></rect>
          <rect
            y="107.48"
            width="152"
            height="152"
            rx="8"
            transform="rotate(-45 0 107.48)"
            fill="white"
          ></rect>
        </svg>
        <div
          class="relative pt-10 px-10 flex items-center justify-center group-hover:scale-110 transition-transform"
        >
          <div
            class="block absolute w-48 h-48 bottom-0 left-0 -mb-24 ml-3"
            style="
              background: radial-gradient(black, transparent 60%);
              transform: rotate3d(0, 0, 1, 20deg) scale3d(1, 0.6, 1);
              opacity: 0.2;
            "
          ></div>
          <iframe
            src="https://lottie.host/embed/984a7979-ea3d-4590-85ca-f8f729341f2f/whx4znowUR.lottie"
            class="w-full h-full"
          ></iframe>
        </div>
        <div class="relative text-white px-6 pb-6 mt-6">
          <span class="block opacity-75 -mb-1">{{ players[3]?.player_name || 'Player 4' }}</span>
          <div class="flex justify-between">
            <span class="block font-semibold text-xl">
              <span v-if="players[3]?.role === 'Sepoy' || players[3]?.player_id === userID || roleVisibility">
                {{ players[3]?.role || 'Role' }}
              </span>
              <span v-else>{{ '' }}</span>
            </span>
            <span
              class="block bg-white rounded-full text-purple-500 text-xs font-bold px-3 py-2 leading-none flex items-center"
            >
              {{ players[3]?.current_round_score || '0' }}*
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import axios from 'axios'
import { getAccessToken } from '@/auth/auth-service'

// Define the structure of a player object
interface Player {
  player_id: string
  player_name: string
  total_score: number
  round_number: number
  role: string
  current_round_score: number
}

interface ApiResponse {
  success: boolean
  message: string
  data: Player[]
}

export default {
  name: 'GameView',
  data() {
    return {
      userID: '',
      userName: '',
      gameCode: '', // Store the game code
      leaderboard: [] as Player[], // Array to store leaderboard data with explicit type
      players: [] as Player[], // Array to store player data
      roundNumber: 0, // Round number
      roleVisibility : false
    }
  },
  methods: {
    // Fetch the game code from URL query params
    async fetch_game_code() {
      const gameCodeQuery = this.$route.query.game_code
      if (gameCodeQuery && typeof gameCodeQuery === 'string') {
        this.gameCode = gameCodeQuery
      } else {
        this.gameCode = 'No game code'
      }
    },

    // Fetch leaderboard data from API with game_code in query params
    async fetch_leaderboard() {
      const backendUrl = import.meta.env.VITE_BACKEND_URL
      const accessToken = await getAccessToken()

      try {
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        console.log('Fetching leaderboard for game code:', this.gameCode)

        const response = await axios.get<ApiResponse>(
          `${backendUrl}/get_leaderboard?game_code=${this.gameCode}`,
          { headers },
        )

        console.log('Leaderboard fetched:', response)

        if (response.data.success) {
          
          // Fill leaderboard with actual players and default players if not present
          const leaderboardData = response.data.data

          // Make sure we always have 4 players
          const players = [...leaderboardData]

          // Fill missing players with default values
          while (players.length < 4) {
            players.push({
              player_id: '',
              player_name: `Player ${players.length + 1}`,
              total_score: 0,
              round_number: 0,
              role: 'Unavailable',
              current_round_score: 0,
            })
          }

          // Sort players alphabetically by player_name
          players.sort((a, b) => a.player_name.localeCompare(b.player_name))

          // Update the leaderboard array
          this.leaderboard = players

          // Assuming round number is consistent across players, get it from the first player
          this.roundNumber = leaderboardData[0]?.round_number || 0

          // Set players array from leaderboard data
          this.players = this.leaderboard
        } else {
          console.error('Failed to fetch leaderboard:', response.data.message)
        }
      } catch (error) {
        console.error('Error fetching leaderboard data:', error)
      }
    },

    // Start the game
    async startGame() {
      const backendUrl = import.meta.env.VITE_BACKEND_URL
      const accessToken = await getAccessToken()

      try {
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        const response = await axios.post(
          `${backendUrl}/start_game`,
          { game_code: this.gameCode },
          { headers },
        )

        if (response.data.success) {
          console.log('Game started successfully:', response.data.message)
        } else {
          console.error('Failed to start game:', response.data.message)
        }
      } catch (error) {
        console.error('Error starting the game:', error)
      }
    },

    // End the game
    async endGame() {
      const backendUrl = import.meta.env.VITE_BACKEND_URL
      const accessToken = await getAccessToken()

      try {
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        const response = await axios.post(
          `${backendUrl}/end_game`,
          { game_code: this.gameCode },
          { headers },
        )

        if (response.data.success) {
          console.log('Game ended successfully:', response.data.message)
          this.$router.push({ name: 'leaderboard', query: { game_code: this.gameCode } })
        } else {
          console.error('Failed to end game:', response.data.message)
        }
      } catch (error) {
        console.error('Error ending the game:', error)
      }
    },

    // Function for Play Game button
    async playGame() {
      const backendUrl = import.meta.env.VITE_BACKEND_URL
      const accessToken = await getAccessToken()
      this.roundNumber += 1 // Increment round number

      const payload = {
        player_1_id: this.players[0]?.player_id,
        player_2_id: this.players[1]?.player_id,
        player_3_id: this.players[2]?.player_id,
        player_4_id: this.players[3]?.player_id,
        game_code: this.gameCode,
        round_number: this.roundNumber,
      }

      try {
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        const response = await axios.post(`${backendUrl}/assign_roles`, payload, { headers })

        if (response.data.success) {
          this.roleVisibility = false
          console.log('Round Processed successfully:', response.data.message)
          this.fetch_leaderboard()
        } else {
          console.error('Failed to process the round:', response.data.message)
        }
      } catch (error:any) {
        if(error.response.data.message === 'roundsOver') {
          this.$router.push({ name: 'leaderboard', query: { game_code: this.gameCode } })
          console.error('All Rounds Finished', error.response.data.message)
        } else {
          console.error('Error playing the game:', error)
        }
      }
    },

    async processSepoyGuess(selectedPlayerId: string) {
      // Find the Sepoy's player_id
      const sepoyPlayer = this.players.find((player) => player.role === 'Sepoy')
      if (!sepoyPlayer) {
        console.error('No Sepoy found')
        return
      }

      const sepoyId = sepoyPlayer.player_id

      // Prepare the payload for the API call
      const payload = {
        sepoy_id: sepoyId,
        selected_player_id: selectedPlayerId,
        round_id: this.roundNumber,
        game_code: this.gameCode,
        guess: 'Thief',
      }

      // Call the API to process the guess
      const backendUrl = import.meta.env.VITE_BACKEND_URL
      const accessToken = await getAccessToken()

      try {
        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        const response = await axios.post(`${backendUrl}/sepoy_guess`, payload, { headers })

        if (response.data.success) {
          console.log('Sepoy guess processed successfully:', response.data.message)
          this.fetch_leaderboard()
        } else {
          console.error('Failed to process Sepoy guess:', response.data.message)
        }
      } catch (error) {
        console.error('Error processing Sepoy guess:', error)
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

          this.userName = givenName
          this.userID = emailId
          console.log('Player Name:', givenName)
          console.log('Player ID:', emailId)
        } else {
          console.error('User attributes not found in response', response.data)
        }
      } catch (error) {
        console.error('Error fetching user details:', error)
      }
    },
    async handleCardClick(player: Player) {
      // Find the logged-in user's role from the players list
      const loggedInUser = this.players.find((p) => p.player_id === this.userID)

      if (!loggedInUser) {
        console.log('Logged-in user not found!')
        return
      }
      // Only process when the clicked player has the 'Sepoy' role
      if (loggedInUser.role === 'Sepoy') {
        // Process the Sepoy's guess action
        this.processSepoyGuess(player.player_id)
        this.roleVisibility = true
      } else {
        // Optionally show a message or log
        alert('Only the Sepoy player can click on this card!')
      }
    },
  },
  async mounted() {
    await this.fetch_game_code() // Fetch game code from URL
    await this.fetch_leaderboard() // Fetch leaderboard data after game code is available
    await this.fetchUserNameAndId() // Fetch user name and ID
  },
}
</script>

<style scoped></style>
