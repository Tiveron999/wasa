<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
      		user: null,
			userExists: false,
			searched: false,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/context");
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		
		searchUsers() {
			this.user = {}; // (re-)initialize the user object

			if (!this.username) {
				this.user = null;
				return;
			}

			if (localStorage.getItem('isLoggedIn')) {
				this.id = localStorage.getItem('userId');
			}

			try {
				this.$axios.get(`/users?username=${this.username}`, {headers: { Authorization: 'Bearer ' + this.id },})
				.then((response) => {
				this.user = response.data;
				console.log(this.user)
				if (this.user.userid) {
					this.userExists = true;
					} else {
					this.user = null
					this.userExists = false;
					}
				this.searched = true;

				})
				.catch((error) => {
					if (error.response) {
						console.error('Status:', error.response.status);
						console.error('Data:', error.response.data);
					}
					this.user = null;
				});
			} catch (e) {
				this.errormsg = e.toString();
			}
		},

		visitProfile(searchedID) {
			this.$router.push({ name: 'Profile', params: { user_id: searchedID } });
		}


      }
	,
	mounted() {
		this.refresh()
		
	}
}
</script>

<template>
	<div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
			<div class="btn-group me-2"></div>
			</div>
		</div>

	<h6>
		<span id = "type" class = "special">
		Search
		</span>
    </h6>
	
	<br>
	<div class="card mb-3">
		<div class="card-body">
			<h6 class="card-title">Search User</h6>
			<div class="input-group">
				<div class="input-container">
					<input type="text" v-model="username" class="form-control" placeholder="Enter username" pattern="^[A-Za-z][A-Za-z0-9_]{5,14}$" required>
				</div>
				<div class="button-container">
					<button class="btn btn-sm btn-outline-primary bump-down-buttons" @click="searchUsers">Search</button>
				</div>

			</div>
		</div>
	</div>

	<div v-if="userExists" class="card">
		<div class="card-body">
			<h4 class="card-title">User: {{ user.username }}</h4>
				<p class="card-text">Posts: {{ user.numphotos }}</p>
				<p class="card-text">Followers: {{ user.followers }}</p>
				<p class="card-text">Following: {{ user.following }}</p>
				<button class="btn btn-sm btn-outline-primary" @click="visitProfile(user.userid)">
				Go to
				</button>
		</div>
	</div>

	<p v-else-if="searched && !userExists" class="no-result-text text-muted d-flex justify-content-center">User does not exist.</p>

	<!--<p v-else class="no-result-text d-flex justify-content-center">No users found.</p> -->



  
	<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

	</div>


</template>


<style>

</style>