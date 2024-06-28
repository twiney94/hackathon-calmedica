<script setup lang="ts">
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Loader2 } from "lucide-vue-next";
import { nextTick, onMounted, PropType, watch, ref, computed } from "vue";
import { performHttpCall } from "@/utils/http";
import type { Patient, Message } from "@/types/patient";

const messages = ref<Message[]>([]);

const model = defineModel({
  prop: "isOpened",
  event: "update:isOpened",
});

const props = defineProps({
  patient: {
    type: Object as PropType<Patient>,
    required: true,
  },
  isOpened: {
    type: Boolean,
    default: false,
  },
});

const emit = defineEmits(["close", "newStatus", "notify"]);

async function scrollToLastMessage(behavior: ScrollBehavior = "auto") {
  if (!messages.value.length) return;
  await nextTick();

  const chat: HTMLElement = document.querySelector(".chat") as HTMLElement;
  const lastMessage: HTMLElement = chat.querySelector(
    ".chat__message:last-child"
  ) as HTMLElement;

  lastMessage.scrollIntoView({ behavior: behavior, block: "start" });
}

watch(messages.value, async () => {
  await scrollToLastMessage("smooth");
});

onMounted(() => {
  watch(
    () => props.isOpened,
    async (newVal, oldVal) => {
      if (newVal) {
        await nextTick();
        await scrollToLastMessage();
      }
    }
  );
});

const headerText = computed(() => {
  return props.patient ? `Suivi du patient ${props.patient.phone}` : "Chat";
});

async function addAudio() {
  const attachedFile = document.querySelector("#fileInput") as HTMLInputElement;

  if (attachedFile && attachedFile.files && attachedFile.files.length > 0) {
    const chatInput = document.querySelector(
      ".chat__input"
    ) as HTMLInputElement;
    const file = attachedFile.files[0];
    const reader = new FileReader();
    attachedFile.disabled = true;

    chatInput.disabled = true;

    reader.onload = async () => {
      messages.value.push({
        content: "Analyse en cours...",
        patient_id: props.patient.uuid,
        created_at: formatDateTime(),
        loading: true,
        audio: reader.result as string,
      });

      await scrollToLastMessage("smooth");

      const formData = new FormData();
      formData.append("file", file);

      const response = await performHttpCall(
        `generate/${props.patient.uuid}`,
        "POST",
        formData,
        true
      );

      const newMessage = response["DESC"];

      messages.value[messages.value.length - 1].content = "Analyse terminée";
      messages.value[messages.value.length - 1].loading = false;
      attachedFile.disabled = false;
      chatInput.disabled = false;

      await postMessage(newMessage);
      emit("newStatus", {
        status: response["STATUS"].toLowerCase(),
        keywords: response["KEYWORDS"],
        patient: props.patient,
      });
      await scrollToLastMessage("smooth");
      await addNotification(newMessage, props.patient.uuid);
    };

    reader.readAsDataURL(file);

    attachedFile.value = "";
  }
}

async function fetchData() {
  const response = await performHttpCall<any>(
    `messages/${props.patient.uuid}`,
    "GET"
  );
  response.data.forEach((message: any) => {
    message.created_at = formatDateTime(message.created_at);
  });
  messages.value = response.data;
}

async function postMessage(message: string) {
  const body = {
    content: message,
    patient_Id: props.patient.uuid,
  };
  const response = await performHttpCall<any>(`messages`, "POST", body);
  const newMessage = response.data;
  newMessage.created_at = formatDateTime(response.data.created_at);
  messages.value.push(response.data);
}

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
    postMessage(messageInput.value);
    await scrollToLastMessage("smooth");
    messageInput.value = "";
  } else {
    await addAudio();
  }
}

watch(
  () => props.patient,
  async () => {
    await fetchData();
    await scrollToLastMessage();
  }
);

const pillColor = computed(() => {
  return props.patient ? `bg-${props.patient.status}-500` : "bg-sky-500";
});

function formatDateTime(backendDateTime?: string) {
  let dateTime;
  if (!backendDateTime) {
    dateTime = new Date();
  } else {
    dateTime = new Date(backendDateTime);
  }
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

async function addNotification(message: string, patient_id: string) {
  const response = await performHttpCall(
    `patients/${patient_id}/notifications`,
    "POST",
    {
      message: message,
    }
  );

  if (response.status === 200) {
    emit("notify");
  }
}
</script>

<template>
  <div
    :class="[
      'chat flex flex-col gap-2.5 shadow-2xl bg-white p-5 rounded-lg',
      { opened: isOpened },
    ]"
    v-if="modelValue"
  >
    <div class="flex justify-between">
      <span
        class="text-xl font-semibold truncate mr-2.5 flex items-center gap-2"
      >
        <div class="w-4 h-4 rounded-full" :class="pillColor"></div>
        {{ headerText }}
      </span>
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
    <ul
      class="flex flex-col flex-1 overflow-y-auto gap-3 shadow-inner px-1.5 py-2"
    >
      <li
        v-for="{
          id,
          content,
          patient_id,
          created_at,
          audio,
          loading,
        } in messages"
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
          <div
            class="chat__text-container flex items-center gap-2"
            v-if="content"
          >
            <p class="chat__text text-lg" v-if="content">
              {{ content }}
            </p>
            <Loader2 v-show="loading" class="w-4 h-4 mr-2 animate-spin" />
          </div>
        </div>
        <span
          :class="['text-base', sender === 'user' ? 'mr-2 text-right' : 'ml-2']"
        >
          {{ created_at }}
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
          accept=".mp3, .wav, .m4a"
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
