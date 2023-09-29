<script>
export default {
	props: ["loading"]
}
</script>

<template>
	<div v-if="loading">
		<div style="text-align: center">
			<div class="spinner"></div>
		</div>
	</div>
	<div v-if="!loading"><slot /></div>
</template>

<style></style>
