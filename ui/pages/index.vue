<script setup lang="ts">
import {
  CalendarIcon,
  FaceSmileIcon,
  GifIcon,
  ListBulletIcon,
  PhotoIcon,
} from '@heroicons/vue/24/outline';
</script>

<template>
  <main class="w-1/2 border-x border-x-neutral-800">
    <div class="p-4 text-xl font-bold border-b border-b-neutral-800">Home</div>
    <div class="h-28 flex px-4 py-2 space-x-4 border-b border-b-neutral-800">
      <img src="https://avatar.vercel.sh/foo.svg" alt="Avatar" class="w-11 h-11 rounded-full" />
      <div class="w-full flex flex-col justify-between">
        <input type="text" class="w-full h-12 text-xl bg-black" placeholder="What's happening?" />
        <div class="flex items-center justify-between">
          <div class="flex space-x-4">
            <PhotoIcon class="w-5 h-5 text-blue-400" />
            <GifIcon class="w-5 h-5 text-blue-400" />
            <ListBulletIcon class="w-5 h-5 text-blue-400" />
            <FaceSmileIcon class="w-5 h-5 text-blue-400" />
            <CalendarIcon class="w-5 h-5 text-blue-400" />
          </div>
          <button
            type="button"
            disabled
            class="px-4 py-1.5 rounded-3xl bg-blue-500 disabled:opacity-50"
          >
            Tweet
          </button>
        </div>
      </div>
    </div>
  </main>
</template>
