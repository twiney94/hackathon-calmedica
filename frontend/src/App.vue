<script setup lang="ts">
import NavBar from "@/components/NavBar.vue";
import PatientList from "@/components/PatientList.vue";
import Chat from "@/components/Chat.vue";
import { ref } from "vue";

const chatIsOpened = ref(true);
const messages = ref([
  {
    id: 1,
    text: "Lorem ipsum dolor sit amet consectetur. Faucibus nibh pulvinar erat aliquam diam.",
    sender: "bot",
    metadata: "26-04-2024 14:32",
  },
  {
    id: 2,
    text: "Lorem ipsum dolor sit amet consectetur. Faucibus nibh pulvinar erat aliquam diam.",
    sender: "user",
    metadata: "26-04-2024 14:33",
  },
  {
    id: 3,
    text: "Lorem ipsum dolor sit amet consectetur. Faucibus nibh pulvinar erat aliquam diam.",
    sender: "bot",
    metadata: "26-04-2024 14:34",
  },
  {
    id: 4,
    text: "Lorem ipsum dolor sit amet consectetur. Faucibus nibh pulvinar erat aliquam diam.",
    sender: "user",
    metadata: "26-04-2024 14:35",
  },
  {
    id: 5,
    text: "Lorem ipsum dolor sit amet consectetur. Faucibus nibh pulvinar erat aliquam diam.",
    sender: "bot",
    metadata: "26-04-2024 14:36",
  },
]);
</script>

<template>
  <div class="container flex flex-col py-8 gap-8">
    <NavBar />
    <PatientList />
    <Chat
      :isOpened="chatIsOpened"
      :headInfo="'37475644738859 | Doe | John | 930208295 | 2024-06-25'"
      :messages="messages"
      @close="chatIsOpened = false"
    />
  </div>
</template>
