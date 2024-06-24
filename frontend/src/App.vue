<script setup lang="ts">
import { Button } from "@/components/ui/button";
import CalmedicaLogo from "@/assets/calmedica-flat.png";
</script>

<template>
  <div class="container">
    <img :src="CalmedicaLogo" alt="Calmedica Logo" />
    <Button class="bg-calmedica">Click me</Button>
  </div>
</template>
