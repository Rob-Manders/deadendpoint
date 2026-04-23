<script setup lang="ts">
import { SetResponseCode, StartServer, StopServer } from '../../wailsjs/go/main/App'
import { onMounted, ref, watch } from 'vue'
import { EventsOn } from '../../wailsjs/runtime'
import { responseCodes } from "../lib/valid-response-codes";
import StartStopButton from "./StartStopButton.vue";

const port = ref(8080)
const isRunning = ref(false)
const responseCode = ref(200)

watch(responseCode, () => {
	SetResponseCode(responseCode.value)
})

onMounted(() => {
	EventsOn("server_running", running => {
		isRunning.value = running
	})
})

function toggleServer() {
	if (isRunning.value) {
		StopServer()
	} else {
		StartServer(port.value)
	}
}

</script>

<template>
	<div class="menu-bar">
		<StartStopButton @click="toggleServer" :is-running="isRunning" />

		<label for="port-input">Port:
			<input name="port-input" type="text" v-model.number="port" :disabled="isRunning"/>
		</label>


		<label for="response-code">
			Response Code:
			<select name="response-code" id="response-code" v-model.number="responseCode">
				<option v-for="c in responseCodes" :value="c.code">{{ c.code }} {{ c.name }}</option>
			</select>
		</label>
	</div>
</template>

<style scoped>
.menu-bar {
	display: flex;
	align-items: center;
	background-color: #1e1f22;
	padding: 0.5rem;
}
</style>
