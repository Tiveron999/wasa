<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			users: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/context");
				this.users = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async getProfile() {
      try {
        const res = await axios.get('/users', {params: {username: this.username}});
        console.log(res.data); // Log the response data
      } catch (err) {
        console.log(err);
      }
	}
	},
	mounted() {
		this.refresh()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<form action="/users/search/:username" method="get">
            <label for="username">Username:</label><br>
            <input type="text" id="username" name="username"><br>
            <input type="submit" value="Submit">
        </form> 
	</div>
	<div>
	  <input v-model="username" placeholder="Enter username">
	  <button @click="getProfile">Get Profile</button>
	</div>
</template>

<style>
</style>