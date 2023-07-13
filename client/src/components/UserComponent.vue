<template>
  <div class="flex flex-row h-full text-white p-5">
  <div class="overflow-y-auto w-full">
    <div class="py-24 sm:py-25">
        <div class="mx-auto max-w-7xl px-6 lg:px-8">
          <h2 class="text-3xl font-bold tracking-tight text-gray-300 sm:text-4xl">Users</h2>
          <p class="mt-2 text-lg leading-8 text-gray-400">All users infected</p>
        </div>
      </div>
    <TableComponent/>
  </div>
</div>
</template>
<script setup>
  import TableComponent from '@/components/TableComponent.vue'
</script>