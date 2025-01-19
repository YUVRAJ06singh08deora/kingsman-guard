import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

import AuthenticationView from '@/views/AuthenticationView.vue'
import GameView from '@/views/GameView.vue'
import GameDashView from '@/views/GameDashView.vue'
import JoinGameView from '@/views/JoinGameView.vue'
import CreateGameView from '@/views/CreateGameView.vue'
import LeaderBoardView from '@/views/LeaderBoardView.vue'


function isAuthenticated() {
  return localStorage.getItem('access_token') !== null
}

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/authenticate',
      name: 'authenticate',
      component: AuthenticationView,
    },
    {
      path: '/game',
      name: 'game',
      component: GameView,
    },
    {
      path: '/gamedashboard',
      name: 'gamedashboard',
      component: GameDashView,
    },
    {
      path: '/joingame',
      name: 'joingame',
      component: JoinGameView,
    },
    {
      path: '/creategame',
      name: 'creategame',
      component: CreateGameView,
    },
    {
      path: '/leaderboard',
      name: 'leaderboard',
      component: LeaderBoardView,
    },
    
  ],
})

// Navigation guard for All routes
router.beforeEach((to, from, next) => {
  if (to.name !== 'authenticate' && !isAuthenticated()) {
    if (to.path !== '/') {
      next('/')
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
