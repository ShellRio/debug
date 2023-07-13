<template>
    <div class="relative text-hacker w-2/4 mx-auto">
      <div class="absolute inset-0 flex items-center" aria-hidden="true">
        <div class="w-full border-t border-hacker" />
      </div>
      <div class="relative flex justify-center">
        <span class="px-2">
          <PlusIcon class="h-5 w-5 bg-gray-800" aria-hidden="true" />
        </span>
      </div>
    </div>
  </template>
  
  <script setup>
  import { PlusIcon } from '@heroicons/vue/20/solid'
  </script>