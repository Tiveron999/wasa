<script>
export default {
	props: ["loading"]
}
</script>

<template>
	<div v-if="loading">
		<div style="text-align: center">
			<div class="loading-circle">
				<div class="progress-bar"></div>
				<div class="progress-bar yellow"></div>
				<div class="progress-bar green"></div>
				<div class="progress-bar blue"></div>
			</div>
		</div>
	</div>
	<div v-if="!loading"><slot /></div>
</template>

<style></style>
