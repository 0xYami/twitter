<script setup lang="ts">
import {
  ArrowPathRoundedSquareIcon,
  ArrowUpTrayIcon,
  ChatBubbleOvalLeftIcon,
  EllipsisHorizontalIcon,
  HeartIcon,
} from '@heroicons/vue/24/outline';
import { getElapsedHoursFromDate } from '~/lib/utils';

defineProps({
  user: {
    type: Object,
    required: true,
  },
  text: {
    type: String,
    required: true,
  },
  createdAt: {
    type: String,
    required: true,
  },
});
</script>

<template>
  <div class="h-auto min-h-28 flex px-4 py-3 space-x-3">
    <img :src="user.avatarURL" alt="Avatar" class="w-12 h-12 rounded-full" />
    <div class="w-full flex flex-col justify-between">
      <div class="relative w-full flex items-center space-x-1">
        <div class="font-bold">{{ user.name }}</div>
        <div class="text-neutral-500">@{{ user.handle }}</div>
        <div class="w-1 h-1 rounded-full bg-neutral-500" />
        <div class="text-neutral-500">{{ getElapsedHoursFromDate(new Date(createdAt)) }}h</div>
        <button type="button" class="absolute right-0 group hover:bg-blue-800/30 p-1 rounded-full">
          <EllipsisHorizontalIcon class="w-5 h-5 text-neutral-500 group-hover:text-blue-500" />
        </button>
      </div>
      <p class="break-all">{{ text }}</p>
      <ul class="flex items-center mt-3 space-x-6">
        <li>
          <button type="button" class="flex items-center space-x-3">
            <ChatBubbleOvalLeftIcon class="w-5 h-5 text-neutral-500" />
            <span class="text-sm text-neutral-500">0</span>
          </button>
        </li>
        <li>
          <button type="button" class="flex items-center space-x-3">
            <ArrowPathRoundedSquareIcon class="w-5 h-5 text-neutral-500" />
            <span class="text-sm text-neutral-500">0</span>
          </button>
        </li>
        <li>
          <button type="button" class="flex items-center space-x-3">
            <HeartIcon class="w-5 h-5 text-neutral-500" />
            <span class="text-sm text-neutral-500">0</span>
          </button>
        </li>
        <li>
          <button type="button" class="flex items-center space-x-3">
            <ArrowUpTrayIcon class="w-4 h-4 text-neutral-500" />
          </button>
        </li>
      </ul>
    </div>
  </div>
</template>
