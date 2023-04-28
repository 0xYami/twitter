<script setup lang="ts">
import { useUserStore } from '~/stores/user';
import { asyncFaillable } from '~/lib/utils';

definePageMeta({
  layout: false,
});

const user = useUserStore();
const username = useState('username', () => '');
const password = useState('password', () => '');
const isRegistering = useState('isRegistering', () => false);

const register = async () => {
  isRegistering.value = true;

  const tryRegister = await asyncFaillable(
    user.register({
      username: username.value,
      password: password.value,
    }),
  );

  if (tryRegister.failed) {
    isRegistering.value = false;
    return;
  }
  await navigateTo('/');
};
</script>

<template>
  <div class="h-screen flex flex-col items-center justify-center bg-gray-800">
    <div class="w-[420px] h-80 px-12 py-8 rounded-xl bg-black">
      <div v-if="isRegistering" class="h-full flex items-center justify-center">
        <div>Loading...</div>
      </div>
      <form
        v-if="!isRegistering"
        @submit.prevent="() => register()"
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
          class="absolute -bottom-2 w-full py-3 text-black font-bold rounded-3xl bg-white"
        >
          Create
        </button>
      </form>
    </div>
  </div>
</template>
