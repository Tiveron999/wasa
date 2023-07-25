<script>
export default {	

    props: {
    isUsernameModalOpen: {
      type: Boolean,
      required: true,
    },
    newUsername: {
      type: String,
      required: true,
    },
    usernamePattern: {
      type: RegExp,
      required: true,
    },
    }
}
</script>


<template>
	<div class="modal fade" id="UsernameModal" tabindex="-1" :class="{ 'show': isUsernameModalOpen }">
		<div class="modal-dialog">
			<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title">Change Username</h5>
				</div>
				<div class="modal-body">
					<input
					type="text"
					class="form-control"
					placeholder="New Username"
					v-model="newUsername"
					:pattern="usernamePattern"
					title="Username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed."
					/>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-primary"  @click="changeUsername">Save</button>
					<button type="button" class="btn btn-secondary" @click="closeUsernameModal">Close</button>
				</div>
			</div>
		</div>
	</div>
</template>


<style>

</style>
