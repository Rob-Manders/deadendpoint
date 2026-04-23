<script setup lang="ts">
const emit = defineEmits<{ (e: 'click', event: MouseEvent): void }>()
const props = defineProps<{ isRunning: boolean }>()
</script>

<template>
	<button class="start-stop-button" @click="(event: MouseEvent) => emit('click', event)">
		<span class="indicator" :style="{ backgroundColor: props.isRunning ? '#42b983' : '#191a1c'}"/>
		{{ props.isRunning ? "Stop" : "Start" }}
	</button>
</template>

<style scoped>
.start-stop-button {
	display: flex;
	align-items: center;
	justify-content: space-between;
	cursor: pointer;
	background-color: #26282e;
	border: none;
	padding: 0.25rem 0.5rem;
	border-radius: 5px;
	font-size: 1.1rem;
	width: 5rem;
	color: #d4d4d4;
	outline: none;
}

.indicator {
	display: block;
	min-width: 0.6rem;
	min-height: 0.6rem;
	border-radius: 200vw;
	margin-right: 0.5rem;
}
</style>