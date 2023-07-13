<template>
    <div class="bg-gray-700 flex flex-col md:flex-row h-screen p-5 md:space-x-5">
        <div class=" mb-5 md:w-60">
            <div class="bg-gray-800 p-5 rounded-lg border-x-2 border-hacker">
            <img src="@/assets/logo.png" alt="logo" class="w-20 mx-auto mb-12">
            <button v-on:click="switchComponent(item.current, item.name)" v-for="item in navigation" :key="item.name" :class="[activeNavigation === item.name ? 'text-hacker border-gray-200' : 'text-gray-300 hover:border-hacker hover:text-gray-300', 'group flex items-center rounded-md px-2 py-2 text-base font-medium border border-gray-800 mt-2 w-40 transition duration-200']">
                      <component :is="item.icon" class="mr-4 h-6 w-6 flex-shrink-0"/>
                      {{ item.name }}
            </button>
        </div>
        </div>
        <div class="bg-gray-800 w-full rounded-lg border-x-2 border-hacker">
            <slot>
                <component :is="currentComponent"/>
            </slot>
        </div>
    </div>
</template>
<script setup>
import { ref } from 'vue';

import {
    MapIcon,
    UsersIcon,
    NewspaperIcon
  } from '@heroicons/vue/24/outline'

  import UserComponent from '@/components/UserComponent.vue'
  import MapsComponent from '@/components/MapsComponent.vue'
  import NewsComponent from '@/components/NewsComponent.vue'

const navigation = [
    { name: 'Users', icon: UsersIcon, current: UserComponent },
    { name: 'News', icon: NewspaperIcon, current: NewsComponent },
    { name: 'Maps', icon: MapIcon, current: MapsComponent },
  ]

const currentComponent = ref(UserComponent)
const activeNavigation = ref('Users')

function switchComponent(selectedComponent, selectedNavigation) {
    currentComponent.value = selectedComponent;
    activeNavigation.value = selectedNavigation;
}
</script>