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
		/*async getProfile() {
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
		},*/
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


<!-- <template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<h6>
			<span id = "type" class = "special">
            Search
			</span>
        </h6>
		 <form action="/users/search/:username" method="get">
            <label for="username">Username:</label>
            <input type="text" id="username" name="username">
            <input type="submit" value="Submit">
        </form> 
	</div>
	<div>
		<input v-model="username" placeholder="Enter username">
		<button class="btn btn-sm-normal btn-outline-primary" @click="getProfile">Go to</button>
	</div>
	<div>
    <input type="text" v-model="username" placeholder="Enter username" pattern="^[A-Za-z][A-Za-z0-9_]{5,14}$" required>
    <button class="btn btn-sm-normal btn-outline-primary" @click="searchUsers">Search</button>
   
      <h4 v-if = user>User: {{ user }}</h4>
      <p v-if = user>Number of uploaded photos: {{ user.num_photos }}</p>
      <p v-if = user>Followers: {{ user.followers }}</p>
      <p v-if = user>Following: {{ user.following }}</p>

  </div> 
</template>  -->

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
			<input type="text" v-model="username" class="form-control" placeholder="Enter username" pattern="^[A-Za-z][A-Za-z0-9_]{5,14}$" required>
			<button class="btn btn-sm btn-outline-primary" @click="searchUsers">Search</button>
			</div>
		</div>
	</div>

	<div v-if="userExists" class="card">
		<div class="card-body">
			<h4 class="card-title" v-if="userExists">User: {{ user.username }}</h4>
				<p class="card-text" v-if="userExists">Posts: {{ user.numphotos }}</p>
				<p class="card-text" v-if="userExists">Followers: {{ user.followers }}</p>
				<p class="card-text" v-if="userExists">Following: {{ user.following }}</p>
				<button class="btn btn-sm btn-outline-primary" @click="visitProfile(user.userid)" v-if="userExists">
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