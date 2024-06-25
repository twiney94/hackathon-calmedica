<script setup lang="ts">
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import {nextTick, onMounted, watch} from "vue";

const props = defineProps({
  isOpened: {
    type: Boolean,
    default: false,
  },
  headInfo: {
    type: String,
    required: true,
  },
  messages: {
    type: Array,
    required: true,
  },
})

defineEmits(['close'])

async function scrollToLastMessage() {
  await nextTick();

  const chat: HTMLElement = document.querySelector('.chat') as HTMLElement;
  const lastMessage: HTMLElement = chat.querySelector('.chat__message:last-child') as HTMLElement;

  lastMessage.scrollIntoView({ behavior: 'smooth', block: 'end' })
}

watch(props.messages, async () => {
  await scrollToLastMessage();
})

onMounted(async () => {
  await scrollToLastMessage();
})

async function addMessage() {
  const messageInput = document.querySelector('.chat__input') as HTMLInputElement;
  const attachedFile = document.querySelector('#fileInput') as HTMLInputElement;

  if(
      messageInput.value.trim() === '' ||
      attachedFile.files && attachedFile.files.length > 0
  ) {
    return;
  }

  if(messageInput.value.trim() !== '') {
    const dateTime = new Date();
    const formatDateTime = `${dateTime.getDate()}-${dateTime.getMonth() + 1}-${dateTime.getFullYear()} ${dateTime.getHours()}:${dateTime.getMinutes()}`;

    props.messages.push({
      id: props.messages.length + 1,
      text: messageInput.value,
      sender: 'user',
      metadata: `Orange - Reçu - ${formatDateTime}`
    })

    await scrollToLastMessage();
    messageInput.value = '';
  } else {
    addAudio();
  }
}

function addAudio() {
  const attachedFile = document.querySelector('#fileInput') as HTMLInputElement;
  const dateTime = new Date();
  const formatDateTime = `${dateTime.getDate()}-${dateTime.getMonth() + 1}-${dateTime.getFullYear()} ${dateTime.getHours()}:${dateTime.getMinutes()}`;
  const file = attachedFile.files[0];
  const reader = new FileReader();

  reader.onload = () => {
    props.messages.push({
      id: props.messages.length + 1,
      audio: reader.result as string,
      sender: 'user',
      metadata: `Orange - Reçu - ${formatDateTime}`
    })

    scrollToLastMessage();
  }

  reader.readAsDataURL(file);

  attachedFile.value = '';
}
</script>

<template>
  <div
      class="chat flex flex-col gap-2.5 shadow-2xl bg-white p-5 rounded-lg"
      v-show="isOpened"
  >
    <div class="flex justify-between">
      <p class="text-xl font-semibold truncate mr-2.5">{{ headInfo }}</p>
      <button class="chat__close" @click="$emit('close')">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-6">
          <path stroke-linecap="round" stroke-linejoin="round" d="M6 18 18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
    <ul class="flex flex-col flex-1 overflow-y-auto gap-3">
      <li
          v-for="{ id, text, audio, sender, metadata } in messages"
          :key="id"
          :class="[
              'chat__message flex flex-col w-fit max-w-96',
              sender === 'user' ? 'ml-auto' : '',
          ]"
      >
        <div
            :class="[
                'rounded-lg shadow-md p-2.5',
                sender === 'user' ? 'bg-sky-500 text-white' : 'bg-sky-100',
            ]"
        >
          <p
              class="text-lg"
              v-if="text"
          >
            {{ text }}
          </p>
          <audio
              :src="audio"
              controls
              v-if="audio"
          />
        </div>
        <span
            :class="[
                'text-base',
                sender === 'user' ? 'mr-2 text-right' : 'ml-2',
            ]"
        >
          {{ metadata }}
        </span>
      </li>
    </ul>
    <form class="relative flex gap-2.5" @submit.prevent="addMessage">
      <Input
          class="chat__input"
          placeholder="Tapez votre message"
      />
      <div class="attached-file flex justify-center items-center px-2">
        <input type="file" id="fileInput" class="hidden" accept=".mp3" @change="addAudio"/>
        <label for="fileInput">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
               class="size-5 cursor-pointer"
          >
            <path stroke-linecap="round" stroke-linejoin="round" d="m18.375 12.739-7.693 7.693a4.5 4.5 0 0 1-6.364-6.364l10.94-10.94A3 3 0 1 1 19.5 7.372L8.552 18.32m.009-.01-.01.01m5.699-9.941-7.81 7.81a1.5 1.5 0 0 0 2.112 2.13" />
          </svg>
        </label>
      </div>
      <Button class="bg-sky-500" type="submit" size="icon">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor"
             class="size-5"
        >
          <path stroke-linecap="round" stroke-linejoin="round" d="M6 12 3.269 3.125A59.769 59.769 0 0 1 21.485 12 59.768 59.768 0 0 1 3.27 20.875L5.999 12Zm0 0h7.5" />
        </svg>
      </Button>
    </form>
  </div>
</template>

<style scoped>
.chat {
  position: fixed;
  top: 70px;
  right: 14px;
  bottom: 50px;
  width: 672px;
}

.attached-file {
  position: absolute;
  top: 0;
  right: calc(40px + 0.5rem);
  height: 100%;
}
</style>