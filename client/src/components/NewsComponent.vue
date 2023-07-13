<template>
    <div class="flex flex-row h-full text-white p-5">
  <div class="overflow-y-auto w-full">
    <div class="my-12">
        <div class="mx-auto max-w-7xl px-6 lg:px-8">
          <h2 class="text-3xl font-bold tracking-tight text-gray-200 sm:text-4xl">Last news</h2>
          <p class="mt-2 text-lg leading-8 text-gray-400">Discover latest news about tech.</p>
        </div>
        <div class="mt-10 space-y-10 space-x-5">
          <div class="flex flex-col item-center">
          <div class="mt-2 flex rounded-md shadow-sm w-96 mx-auto">
            <div class="relative flex flex-grow items-stretch focus-within:z-10">
              <input v-model="newSearch" class="block w-full pl-2 rounded-none rounded-l-md border-0 py-1.5 text-gray-900 ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 sm:text-sm sm:leading-6" placeholder="Search..." />
            </div>
            <button @click="searchNews()" type="button" class="relative -ml-px inline-flex items-center gap-x-1.5 rounded-r-md px-3 py-2 text-sm font-semibold text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-500 bg-gray-200">
              <DocumentMagnifyingGlassIcon class="-ml-0.5 h-5 w-5 text-gray-400" aria-hidden="true" />
            </button>
          </div>
        </div>
          <button @click="switchCategory(item.name)" v-for="item in navCategories" :key="item.index" class="border font-semibold py-1 px-3 rounded-full">{{item.name}}</button>
        </div>
      </div>
      <div class="flex flex-col items-center">
        <div class="flex justify-between w-1/2 -mb-5">
          <button @click="switchLanguages(item.ref)" class="bg-gray-700 py-2 px-5 rounded-t-xl" v-for="item in languages" :key="item.index">{{item.name}}</button>
        </div>
      <article v-for="article in news" :key="article.index">
        <a :href="article.url" target="_blank">
        <div class="relative isolate flex flex-col gap-8 lg:flex-row my-5 bg-gray-700 p-5 rounded-xl hover:bg-gray-600 hover:cursor-pointer hover:scale-105 transition duration-500">
            <div class="relative sm:aspect-[2/1] lg:aspect-square lg:w-64 lg:shrink-0">
              <img :src="article.urlToImage" alt="" class="absolute inset-0 h-full w-full rounded-2xl bg-gray-50 object-cover" />
              <div class="absolute inset-0 rounded-2xl ring-1 ring-inset ring-gray-900/10" />
            </div>
            <div>
              <div class="flex items-center gap-x-4 text-xs">
                <time :datetime="article.publishedAt" class="text-gray-500">{{ article.publishedAt }}</time>
                <a :href="article.url" class="relative z-10 rounded-full py-1.5 px-3 font-medium text-gray-600 bg-gray-200">{{ article.source.name }}</a>
              </div>
              <div class="group relative max-w-xl">
                <h3 class="mt-3 text-lg font-semibold leading-6 text-hacker text-left ">
                    <span class="absolute inset-0" />
                    {{ article.title }}
                </h3>
                <p class="mt-5 text-sm leading-6 text-gray-300 text-justify">{{ article.description }}</p>
                <p class="mt-5 text-sm leading-6 text-gray-300 text-justify">{{ article.content }}</p>
              </div>
              <div class="mt-6 flex border-tpt-6">
                <div class="relative flex items-center gap-x-4">
                  <div class="text-sm leading-6">
                    <p class="font-semibold text-gray-300">
                        <span class="absolute inset-0" />
                        {{ article.author }}
                    </p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </a>
            <DividerComponent />
          </article>
        </div>
  </div>
</div>

</template>
<script setup>
import {ref, onBeforeMount, shallowRef} from 'vue'
// import News from '@/data/news-test.json'
import DividerComponent from './DividerComponent.vue';
import { DocumentMagnifyingGlassIcon } from '@heroicons/vue/20/solid'
import {getAllNews} from '@/services/news.js'

let news = shallowRef({})

let selectedCategory = 'hacker';
let selectedLanguage = 'fr';

let newSearch = ref('')


onBeforeMount(async () => {
	await updateNews();
});

const navCategories = [
  {
    name: 'hacker',
  },
  {
    name: 'hacking',
  },
  {
    name: 'cyber security',
  },
  {
    name: 'tech',
  },
  {
    name: 'malware',
  },
];
const languages = [
  {
    name: 'Français',
    ref: 'fr',
  },
  {
    name: 'English',
    ref: 'en',
  },
];

async function switchCategory(itemCategory) {
  selectedCategory = itemCategory;
  await updateNews();
}
async function switchLanguages(language) {
  selectedLanguage = language;
  await updateNews();
}

async function searchNews() {
  selectedCategory = newSearch.value;
  await updateNews();
}

async function updateNews() {
 const response = await getAllNews(selectedCategory, selectedLanguage);
  news.value = response.data.articles;
}
</script>