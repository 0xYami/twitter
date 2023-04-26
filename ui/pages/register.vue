<script setup lang="ts">
import { useUserStore } from '~/stores/user';

const user = useUserStore();
const username = useState('username', () => '');
const password = useState('password', () => '');
</script>

<template>
  <div class="h-screen flex flex-col items-center justify-center bg-gray-800">
    <div class="w-[420px] h-80 px-12 pt-8 pb-2 rounded-xl bg-black">
      <form
        @submit.prevent="() => user.register({ username, password })"
        class="relative h-full flex flex-col items-start space-y-4"
      >
        <div class="text-3xl font-bold mb-2">Create your account</div>
        <input
          type="text"
          v-model.trim="username"
          placeholder="Username"
          autofocus
          required
          class="w-full px-2 py-4 border-[0.3px] border-neutral-700 outline-none outline-offset-0 focus:outline-blue-700 rounded bg-black"
        />
        <input
          type="password"
          v-model.trim="password"
          placeholder="Password"
          required
          pattern=".{8,}"
          class="w-full px-2 py-4 border-[0.3px] border-neutral-700 outline-none outline-offset-0 focus:outline-blue-700 rounded bg-black"
        />
        <button
          type="submit"
          class="absolute bottom-3 w-full py-3 text-black font-bold rounded-3xl bg-white"
        >
          Create
        </button>
      </form>
    </div>
  </div>
</template>
