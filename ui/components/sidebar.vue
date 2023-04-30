<script setup lang="ts">
type FollowSuggestion = {
  username: string;
  handle: string;
  avatarURL: string;
};

const followSuggestions: FollowSuggestion[] = [
  {
    username: 'Alice',
    handle: 'alice',
    avatarURL: 'https://avatar.vercel.sh/alice.svg',
  },
  {
    username: 'John',
    handle: 'john',
    avatarURL: 'https://avatar.vercel.sh/john.svg',
  },
  {
    username: 'Bob',
    handle: 'bob',
    avatarURL: 'https://avatar.vercel.sh/bob.svg',
  },
];
</script>

<template>
  <div class="grow flex flex-col p-4 space-y-4">
    <input type="text" placeholder="Search Twitter" class="w-full p-3 rounded-3xl bg-neutral-900" />
    <div class="py-4 bg-bunker-500 rounded-xl">
      <div class="px-4 text-xl font-bold mb-2">Who to follow</div>
      <div
        v-for="suggestion in followSuggestions"
        class="flex items-center justify-between px-4 py-2"
      >
        <div class="flex items-center space-x-3">
          <img :src="suggestion.avatarURL" alt="Avatar" class="w-11 h-11 rounded-full" />
          <div class="flex flex-col">
            <div>{{ suggestion.username }}</div>
            <div class="text-neutral-400">@{{ suggestion.handle }}</div>
          </div>
        </div>
        <button
          type="button"
          class="bg-white text-sm text-black font-bold px-4 py-1.5 rounded-full"
        >
          Follow
        </button>
      </div>
    </div>
  </div>
</template>
