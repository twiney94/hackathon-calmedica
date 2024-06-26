<template>
  <nav class="bg-white shadow rounded-md">
    <div class="container mx-auto px-4">
      <div class="flex justify-between items-center py-4">
        <router-link
          to="/"
          class="text-lg font-bold text-gray-800 cursor-pointer"
        >
          <img
            :src="CalmedicaLogo"
            alt="Calmedica Logo"
            class="h-8 w-auto inline-block mr-2"
          />
        </router-link>
        <div class="flex gap-4">
          <router-link to="/login" class="text-sky-600 font-bold"
            >Afficher les patients</router-link
          >
          <router-link to="/register" class="text-gray-800"
            >Liste d'attente</router-link
          >
          <router-link to="/register" class="text-gray-800"
            >Rapport</router-link
          >
          <router-link to="/register" class="text-gray-800"
            >Tableau de bord</router-link
          >
          <router-link to="/register" class="text-gray-800"
            >Autorisations</router-link
          >
          <Popover>
            <PopoverTrigger>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke="currentColor"
                class="size-6"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M14.857 17.082a23.848 23.848 0 0 0 5.454-1.31A8.967 8.967 0 0 1 18 9.75V9A6 6 0 0 0 6 9v.75a8.967 8.967 0 0 1-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 0 1-5.714 0m5.714 0a3 3 0 1 1-5.714 0"
                />
              </svg>
            </PopoverTrigger>
            <PopoverContent>
              <div
                id="toast-message-cta"
                class="w-full max-w-xs p-4 text-gray-500 bg-white rounded-lg shadow dark:bg-gray-800 dark:text-gray-400"
                role="alert"
              >
                <div class="flex">
                  <div class="w-8 h-8 bg-orange-500 rounded-full" s />
                  <div class="ms-3 text-sm font-normal">
                    <span
                      class="mb-1 text-sm font-semibold text-gray-900 dark:text-white"
                      >01 75 45 63 58</span
                    >
                    <div class="mb-2 text-sm font-normal">
                      Cet appel nécessite une intervention (sutures
                      défaillantes).
                    </div>
                    <a
                      href="#"
                      class="inline-flex px-2.5 py-1.5 text-xs font-medium text-center text-white bg-red-500 rounded-lg hover:bg-blue-700 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-500 dark:hover:bg-blue-600 dark:focus:ring-blue-800"
                      >Intervenir</a
                    >
                  </div>
                  <button
                    type="button"
                    class="ms-auto -mx-1.5 -my-1.5 bg-white justify-center items-center flex-shrink-0 text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700"
                    data-dismiss-target="#toast-message-cta"
                    aria-label="Close"
                  >
                    <span class="sr-only">Close</span>
                    <svg
                      class="w-3 h-3"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 14 14"
                    >
                      <path
                        stroke="currentColor"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
                      />
                    </svg>
                  </button>
                </div>
              </div>
            </PopoverContent>
          </Popover>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup lang="ts">
import CalmedicaLogo from "@/assets/calmedica-blue.png";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
</script>
