<template>
  <router-view v-slot="{ Component }">
    <transition name="route" mode="out-in">
      <component :is="Component" class="max-h-screen"></component>
    </transition>
  </router-view>
</template>