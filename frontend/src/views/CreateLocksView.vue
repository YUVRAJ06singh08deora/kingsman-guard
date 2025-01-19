<template>
  <dashboard-layout>
    <h1 class="text-xl font-bold text-gray-700">Add New Lock</h1>
    <form @submit="submitHandler">
      <div class="mt-5">
        <label for="lock_type" class="block text-base font-semibold text-gray-700">Lock Type</label>
        <div class="mt-2 grid grid-cols-1">
          <select
            id="lock_type"
            name="lock_type"
            v-model="formData.lock_type"
            autocomplete="off"
            class="col-start-1 row-start-1 w-full appearance-none rounded-md bg-white py-1.5 pl-3 pr-8 text-base text-gray-700 outline outline-1 -outline-offset-1 outline-gray-300 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-600"
          >
            <option value="" disabled selected>Please Select one Option</option>
            <option value="Window">Window</option>
            <option value="Door">Door</option>
            <option value="Terrace">Terrace</option>
          </select>
        </div>
      </div>

      <div class="mt-5">
        <p class="text-base font-semibold text-gray-700">Lock Status:</p>
        <div class="mt-2 flex gap-10">
          <div class="flex items-center gap-x-2">
            <input
              id="locked"
              type="radio"
              :value="true"
              v-model="formData.lock_status"
              class="relative size-4 appearance-none rounded-full border border-gray-300 bg-white checked:border-indigo-600 checked:bg-indigo-600"
            />
            <label for="locked" class="block text-base font-medium text-gray-700">Active</label>
          </div>
          <div class="flex items-center gap-x-2">
            <input
              id="unlocked"
              type="radio"
              :value="false"
              v-model="formData.lock_status"
              class="relative size-4 appearance-none rounded-full border border-gray-300 bg-white checked:border-indigo-600 checked:bg-indigo-600"
            />
            <label for="unlocked" class="block text-base font-medium text-gray-700">Inactive</label>
          </div>
        </div>
      </div>

      <div class="mt-5">
        <label for="temp_guest_code" class="block text-base font-semibold text-gray-700"
          >Temporary Guest Code:</label
        >
        <div class="mt-2">
          <input
            type="number"
            id="temp_guest_code"
            v-model.number="formData.temp_guest_code"
            class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-700 outline outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400"
            autocomplete="off"
            placeholder="123456"
          />
        </div>
      </div>

      <div class="mt-5">
        <p class="text-base font-semibold text-gray-700">Is Locked:</p>
        <div class="mt-2 flex gap-10">
          <div class="flex items-center gap-x-2">
            <input
              id="is_locked_yes"
              type="radio"
              :value="true"
              v-model="formData.is_locked"
              class="relative size-4 appearance-none rounded-full border border-gray-300 bg-white checked:border-indigo-600 checked:bg-indigo-600"
            />
            <label for="is_locked_yes" class="block text-base font-medium text-gray-700">Yes</label>
          </div>
          <div class="flex items-center gap-x-2">
            <input
              id="is_locked_no"
              type="radio"
              :value="false"
              v-model="formData.is_locked"
              class="relative size-4 appearance-none rounded-full border border-gray-300 bg-white checked:border-indigo-600 checked:bg-indigo-600"
            />
            <label for="is_locked_no" class="block text-base font-medium text-gray-700">No</label>
          </div>
        </div>
      </div>

      <div class="mt-5">
        <label for="door_password" class="block text-base font-medium text-gray-700"
          >Lock Password:</label
        >
        <div class="mt-2">
          <input
            type="password"
            id="door_password"
            v-model="formData.door_password"
            class="block w-full rounded-md bg-white px-3 py-1.5 text-base text-gray-700 outline outline-1 -outline-offset-1 outline-gray-300 placeholder:text-gray-400"
            autocomplete="new-password"
            placeholder="Password"
          />
        </div>
      </div>

      <button
        type="submit"
        class="mt-6 rounded-md bg-indigo-600 py-2 px-4 text-white hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
      >
        Submit
      </button>
    </form>
  </dashboard-layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import axios from 'axios'
import DashboardLayout from '@/components/DashboardLayout.vue'
import { getAccessToken } from '@/auth/auth-service'

interface FormData {
  lock_type: string
  lock_status: boolean
  temp_guest_code: number | null
  is_locked: boolean
  door_password: string
}

export default defineComponent({
  name: 'CreateLocks',
  components: {
    DashboardLayout,
  },
  data(): { formData: FormData } {
    return {
      formData: {
        lock_type: '',
        lock_status: false,
        temp_guest_code: null,
        is_locked: false,
        door_password: '',
      },
    }
  },
  methods: {
    async submitHandler(e: Event) {
      e.preventDefault()

      if (
        !this.formData.lock_type ||
        this.formData.temp_guest_code === null ||
        !this.formData.door_password
      ) {
        alert('Please fill all required fields.')
        return
      }

      try {
        const backendUrl = import.meta.env.VITE_BACKEND_URL

        const accessToken = await getAccessToken()

        if (!accessToken) {
          alert('Access token is missing. Please log in again.')
          return
        }

        const headers = {
          Authorization: `Bearer ${accessToken}`,
          'Content-Type': 'application/json',
        }

        const response = await axios.post(
          `${backendUrl}/create_lock`,
          JSON.stringify(this.formData),
          { headers },
        )

        alert(`${response.data.message}`)
      } catch (error: any) {
        console.error('Error creating lock:', error)
        alert(`Failed to create lock: ${error.response?.data?.message || error.message}`)
      }
    },
  },
})
</script>
