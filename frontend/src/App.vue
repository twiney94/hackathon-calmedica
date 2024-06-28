<script setup lang="ts">
import NavBar from "@/components/NavBar.vue";
import PatientList from "@/components/PatientList.vue";
</script>

<template>
  <div class="container flex flex-col py-8 gap-8">
    <NavBar />
    <PatientList />
  </div>
</template>

<style scoped>
.status-blue {
  background-color: #417cda;
}
.status-yellow {
  background-color: #fcc858;
}
.status-orange {
  background-color: #fdba74;
}
.status-red {
  background-color: #dc2626;
}
.status-gray {
  background-color: #d1d5db;
}
</style>
