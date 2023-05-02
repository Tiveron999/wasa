<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
      		user: null,
			
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
				if (localStorage.getItem('isLoggedIn')) {
					this.id = localStorage.getItem('userId');
				}
			
				const res = await this.$axios.get('/users', {headers: {Authorization: 'Bearer ' + this.id}}); // {params: {username: this.username}}
				console.log(res.data);
				this.$router.push('/users/:user_id/profile')

			} catch (err) {
				console.log(err);
			}
		},
		searchUsers() {
			if (localStorage.getItem('isLoggedIn')) {
					this.id = localStorage.getItem('userId');
				}
			this.$axios.get(`/users?username=${this.username}`, {headers: {Authorization: 'Bearer ' + this.id}})
				.then(response => {
				this.user = response.data;
				})
				.catch(error => {
				console.error(error);
				});
    	},

      }
	,
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
					<button type="button" id="refresh-button" class="btn btn-sm btn-outline-primary" @click="refresh">
						Refresh
					</button>
					
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<h6>
			<span id = "type" class = "special">
            Search
			</span>
        </h6>
		<!-- <form action="/users/search/:username" method="get">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username">
            <input type="submit" value="Submit">
        </form> -->
	</div>
	<div>
		<input v-model="username" placeholder="Enter username">
		<button class="btn btn-sm-normal btn-outline-primary" @click="getProfile">Go to</button>
	</div>
	<div>
    <input type="text" v-model="username" placeholder="Enter username" pattern="^[A-Za-z][A-Za-z0-9_]{5,14}$" required>
    <button class="btn btn-sm-normal btn-outline-primary" @click="searchUsers">Search</button>
    <div v-if="user">
      <h1>{{ user.username }}</h1>
      <p>Number of uploaded photos: {{ user.num_photos }}</p>
      <p>Followers: {{ user.followers }}</p>
      <p>Following: {{ user.following }}</p>
    </div>
  </div>
</template>


<style>

</style>