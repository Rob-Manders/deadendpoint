<script lang="ts" setup>
import { onMounted, ref } from "vue"
import { EventsOn } from "../wailsjs/runtime"

import MenuBar from './components/MenuBar.vue'

const requests = ref<any[]>([])
const isRunning = ref(false)

onMounted(() => {
  EventsOn("request_received", request => {
    requests.value.unshift(request)
  })
})

</script>

<template>
  <MenuBar/>

  <h2>Requests</h2>

  <div class="request" v-for="(r, i) in requests" :key="i">
    <pre>{{ JSON.stringify(r, null, 2) }}</pre>
  </div>
</template>

<style>

</style>
