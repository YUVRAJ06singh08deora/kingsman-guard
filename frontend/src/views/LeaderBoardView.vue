<template>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg mt-10">
      <!-- Table with LeaderBoard caption -->
      <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <caption
          class="p-5 text-lg font-semibold text-center rtl:text-center text-gray-900 bg-white dark:text-white dark:bg-gray-800"
        >
          LeaderBoard
        </caption>
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" class="px-6 py-3">Player Name</th>
            <th scope="col" class="px-6 py-3">Position</th>
            <th scope="col" class="px-6 py-3">Score</th> <!-- New Score column -->
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="(player, index) in leaderboard"
            :key="player.player_id"
            class="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
          >
            <th
              scope="row"
              class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
            >
              {{ player.player_name }}
            </th>
            <td class="px-6 py-4">{{ index + 1 }}</td> <!-- Position (rank) -->
            <td class="px-6 py-4">{{ player.total_score }}</td> <!-- Display player score -->
          </tr>
        </tbody>
      </table>
  
      <!-- Centered "Home" button -->
      <div class="flex justify-center mt-6">
        <button 
          type="button" 
          @click="goToHome"
          class="text-white bg-gradient-to-r from-teal-400 via-teal-500 to-teal-600 hover:bg-gradient-to-br focus:ring-4 focus:outline-none focus:ring-teal-300 dark:focus:ring-teal-800 font-medium rounded-lg text-sm px-5 py-2.5 text-center">
          Home
        </button>
      </div>
    </div>
  </template>
  
  <script lang="ts">
  import axios from 'axios'
  import { getAccessToken } from '@/auth/auth-service'
  
  // Define a Player type interface
  interface Player {
    player_id: string;
    player_name: string;
    total_score: number;
    round_number: number;
    role: string;
    current_round_score: number;
  }
  
  export default {
    name: 'LeaderBoardView',
    data() {
      return {
        gameCode: '', // Replace with your dynamic game code or prop
        leaderboard: [] as Player[], // List to hold leaderboard players, typed as Player[]
        roundNumber: 0, // To track round number, if needed
      }
    },
    async mounted() {
      // Fetch leaderboard data when the component is mounted
      await this.fetch_game_code()
      await this.fetchLeaderboard()
    },
    methods: {
      async fetch_game_code() {
        const gameCodeQuery = this.$route.query.game_code
        if (gameCodeQuery && typeof gameCodeQuery === 'string') {
          this.gameCode = gameCodeQuery
        } else {
          this.gameCode = 'No game code'
        }
      },
      // Fetch leaderboard data from API with game_code in query params
      async fetchLeaderboard() {
        const backendUrl = import.meta.env.VITE_BACKEND_URL
        const accessToken = await getAccessToken()
  
        try {
          const headers = {
            Authorization: `Bearer ${accessToken}`,
            'Content-Type': 'application/json',
          }
  
          console.log('Fetching leaderboard for game code:', this.gameCode)
  
          const response = await axios.get(
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
  
            // Sort players by total_score in descending order
            players.sort((a, b) => b.total_score - a.total_score)
  
            // Update the leaderboard array
            this.leaderboard = players
  
            // Assuming round number is consistent across players, get it from the first player
            this.roundNumber = leaderboardData[0]?.round_number || 0
          } else {
            console.error('Failed to fetch leaderboard:', response.data.message)
          }
        } catch (error) {
          console.error('Error fetching leaderboard data:', error)
        }
      },
      // Method to handle "Home" button click and navigate to /gamedashboard route
      goToHome() {
        this.$router.push('/gamedashboard') // Redirect to the /gamedashboard route
      }
    },
  }
  </script>
  
  <style scoped>
  /* Add any necessary custom styles here */
  </style>
  