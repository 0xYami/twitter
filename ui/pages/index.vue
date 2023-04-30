<script setup lang="ts">
import { useMutation, useQuery } from '@tanstack/vue-query';
import axios from 'axios';
import {
  CalendarIcon,
  FaceSmileIcon,
  GifIcon,
  ListBulletIcon,
  PhotoIcon,
} from '@heroicons/vue/24/outline';
import { useUserStore } from '~/stores/user';

type Tweet = {
  text: string;
  createdAt: string;
  user: {
    name: string;
    handle: string;
  };
};

const userStore = useUserStore();
const content = useState<null | string>('content', () => null);

const tweetsQuery = useQuery({
  queryKey: ['tweets'],
  queryFn: async () => {
    const { data } = await axios.get<Tweet[]>('http://localhost:4000/tweets/latest');
    return data;
  },
});

const createTweetQuery = useMutation({
  mutationKey: ['createTweet'],
  mutationFn: async () => {
    const body = { text: content.value };
    return axios.post('http://localhost:4000/tweets', body, {
      withCredentials: true,
    });
  },
});
</script>

<template>
  <main class="w-1/2 border-x border-x-neutral-800">
    <div class="p-4 text-xl font-bold border-b border-b-neutral-800">Home</div>
    <div class="h-28 flex px-4 py-2 space-x-4 border-b border-b-neutral-800">
      <img
        :src="'https://avatar.vercel.sh/' + userStore.username + 'foo.svg'"
        alt="Avatar"
        class="w-11 h-11 rounded-full"
      />
      <form
        @submit.prevent="() => createTweetQuery.mutate()"
        class="w-full flex flex-col justify-between"
      >
        <input
          type="text"
          v-model="content"
          placeholder="What's happening?"
          class="w-full h-12 text-xl bg-black focus:outline-none"
        />
        <div class="flex items-center justify-between">
          <div class="flex space-x-4">
            <PhotoIcon class="w-5 h-5 text-blue-400" />
            <GifIcon class="w-5 h-5 text-blue-400" />
            <ListBulletIcon class="w-5 h-5 text-blue-400" />
            <FaceSmileIcon class="w-5 h-5 text-blue-400" />
            <CalendarIcon class="w-5 h-5 text-blue-400" />
          </div>
          <button
            type="submit"
            :disabled="!content"
            class="inline-flex px-4 py-1.5 rounded-3xl bg-blue-500 disabled:opacity-50"
          >
            <svg
              v-if="createTweetQuery.isLoading.value"
              class="animate-spin h-5 w-5 text-white"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
            >
              <circle
                class="opacity-25"
                cx="12"
                cy="12"
                r="10"
                stroke="currentColor"
                stroke-width="4"
              ></circle>
              <path
                class="opacity-75"
                fill="currentColor"
                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
              ></path>
            </svg>
            <span v-else>Tweet</span>
          </button>
        </div>
      </form>
    </div>
    <div v-if="tweetsQuery.isLoading.value">Loading...</div>
    <ul v-else>
      <li v-for="tweet in tweetsQuery.data.value" class="border-b border-neutral-800">
        <Tweet :user="tweet.user" :text="tweet.text" :createdAt="tweet.createdAt" />
      </li>
    </ul>
  </main>
</template>
