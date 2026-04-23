<script lang="ts" setup>
import { ref, onMounted } from "vue"
import { StartServer, StopServer } from "../wailsjs/go/main/App"
import { EventsOn } from "../wailsjs/runtime"

const requests = ref<any[]>([])
const isRunning = ref(false)

function startServer() {
  StartServer(8080, 200)
}

function stopServer() {
  StopServer()
}

onMounted(() => {
  EventsOn("request_received", request => {
    requests.value.unshift(request)
  })

  EventsOn("server_running", running => {
    isRunning.value = running
  })
})

</script>

<template>
  <div class="container">
    <h1>DeadEndpoint</h1>

    <div class="controls">
      <button @click="startServer" :disabled="isRunning">Start</button>
      <button @click="stopServer" :disabled="!isRunning">Stop</button>

      <span>Status:</span> <strong>{{ isRunning ? "Running" : "Stopped" }}</strong>
    </div>
  </div>

  <h2>Requests</h2>

  <div class="request" v-for="(r, i) in requests" :key="i">
    <pre>{{ JSON.stringify(r, null, 2)}}</pre>
  </div>
</template>

<style>

</style>
