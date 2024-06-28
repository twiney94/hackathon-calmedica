<script setup lang="ts">
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Loader2 } from 'lucide-vue-next';
import { nextTick, onMounted, PropType, watch, ref } from "vue";
import { performHttpCall } from "@/utils/http";

interface Message {
  id: number;
  text?: string;
  audio?: string;
  sender: string;
  metadata: string;
}

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
    type: Array as PropType<Message[]>,
    required: true,
  },
});
const loading = ref(false);

defineEmits(["close"]);

async function scrollToLastMessage(behavior: ScrollBehavior = "auto") {
  await nextTick();

  const chat: HTMLElement = document.querySelector(".chat") as HTMLElement;
  const lastMessage: HTMLElement = chat.querySelector(
    ".chat__message:last-child"
  ) as HTMLElement;

  lastMessage.scrollIntoView({ behavior: behavior, block: "end" });
}

watch(props.messages, async () => {
  await scrollToLastMessage("smooth");
});

onMounted(() => {
  watch(() => props.isOpened, async (newVal, oldVal) => {
    if (newVal) {
      await nextTick();
      await scrollToLastMessage();
    }
  });
});

async function addMessage() {
  const messageInput = document.querySelector(
    ".chat__input"
  ) as HTMLInputElement;
  const attachedFile = document.querySelector("#fileInput") as HTMLInputElement;

  if (
    messageInput.value.trim() === "" ||
    (attachedFile.files && attachedFile.files.length > 0)
  ) {
    return;
  }

  if (messageInput.value.trim() !== "") {
    props.messages.push({
      id: props.messages.length + 1,
      text: messageInput.value,
      sender: "user",
      metadata: formatDateTime(),
    });

    await scrollToLastMessage("smooth");
    messageInput.value = "";
  } else {
    await addAudio();
  }
}

async function addAudio() {
  const attachedFile = document.querySelector("#fileInput") as HTMLInputElement;

  if (attachedFile && attachedFile.files && attachedFile.files.length > 0) {
    const file = attachedFile.files[0];
    const reader = new FileReader();

    reader.onload = async () => {
      loading.value = true;
      props.messages.push({
        id: props.messages.length + 1,
        text: 'En cours d\'analyse par l\'IA ...',
        audio: reader.result as string,
        sender: "user",
        metadata: formatDateTime(),
      });

      await scrollToLastMessage("smooth");

      const formData = new FormData();
      formData.append("file", file);

      await performHttpCall('generate/a667dd47-1471-410a-bcd3-d6d45d442880', 'POST', formData, true)
      props.messages[props.messages.length - 1].text = 'Analyse terminée';
      loading.value = false;
    };

    reader.readAsDataURL(file);

    attachedFile.value = "";
  }
}

function formatDateTime() {
  const dateTime = new Date();
  const day = formatDateTimeValue(dateTime.getDate());
  const month = formatDateTimeValue(dateTime.getMonth() + 1);
  const year = dateTime.getFullYear();
  const hours = formatDateTimeValue(dateTime.getHours());
  const minutes = formatDateTimeValue(dateTime.getMinutes());

  function formatDateTimeValue(value: number) {
    return value < 10 ? `0${value}` : value;
  }

  return `${day}-${month}-${year} ${hours}:${minutes}`;
}
</script>

<template>
  <div
    :class="[
        'chat flex flex-col gap-2.5 shadow-2xl bg-white p-5 rounded-lg',
        { 'opened': isOpened }
    ]"
  >
    <div class="flex justify-between">
      <p class="text-xl font-semibold truncate mr-2.5">{{ headInfo }}</p>
      <button class="chat__close" @click="$emit('close')">
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
            d="M6 18 18 6M6 6l12 12"
          />
        </svg>
      </button>
    </div>
    <ul class="flex flex-col flex-1 overflow-y-auto gap-3 shadow-inner px-1.5 py-2">
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
            sender === 'user' ? 'bg-sky-200' : 'bg-gray-100',
          ]"
        >
          <audio :src="audio" controls v-if="audio" />
          <div class="chat__text-container flex items-center gap-2" v-if="text">
            <p class="chat__text text-lg" v-if="text">
              {{ text }}
            </p>
            <Loader2
                v-show="loading"
                class="w-4 h-4 mr-2 animate-spin"
            />
          </div>
        </div>
        <span
          :class="['text-base', sender === 'user' ? 'mr-2 text-right' : 'ml-2']"
        >
          {{ metadata }}
        </span>
      </li>
    </ul>
    <form class="relative flex gap-2.5" @submit.prevent="addMessage">
      <Input class="chat__input" placeholder="Tapez votre message" />
      <div class="attached-file flex justify-center items-center px-2">
        <input
          type="file"
          id="fileInput"
          class="hidden"
          accept=".mp3, .wav"
          @change="addAudio"
        />
        <label for="fileInput">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-5 cursor-pointer"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="m18.375 12.739-7.693 7.693a4.5 4.5 0 0 1-6.364-6.364l10.94-10.94A3 3 0 1 1 19.5 7.372L8.552 18.32m.009-.01-.01.01m5.699-9.941-7.81 7.81a1.5 1.5 0 0 0 2.112 2.13"
            />
          </svg>
        </label>
      </div>
      <Button class="bg-sky-500" type="submit" size="icon">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          stroke-width="1.5"
          stroke="currentColor"
          class="size-5"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M6 12 3.269 3.125A59.769 59.769 0 0 1 21.485 12 59.768 59.768 0 0 1 3.27 20.875L5.999 12Zm0 0h7.5"
          />
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
  transform: translateX(calc(100% + 14px));
  transition: transform 0.3s ease;
}

.chat.opened {
  transform: translateX(0);
}

.attached-file {
  position: absolute;
  top: 0;
  right: calc(40px + 0.5rem);
  height: 100%;
}

audio + .chat__text-container .chat__text {
  margin-left: 1rem;
}
</style>
