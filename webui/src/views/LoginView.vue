<script>
import axios from 'axios'

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			isLoggedIn: false,
        	username: '',
			userId: '',
			error: null,
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/context");
				// this.users = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		
		submit() {
		console.log(JSON.stringify({username: this.username}))
		
		let r = new RegExp('^[A-Za-z][A-Za-z0-9_]{5,14}$')
		if(r.test(this.username)) {
			console.log('ok');
		} else {
			console.log('not ok');
			alert('Invalid username. The username can only contain the characters A-Z, a-z, 0-9 and _ and has to be of length 6 - 15.')
			this.username = ''
			return
		}

		// Send the username to the Go function
		this.$axios.post('/session', JSON.stringify({username: this.username}), {headers: { 'Content-Type': 'application/json',  'Accept': 'application/json'}})
		.then(response => {
			console.log(response.status);
			// Update the local storage
			localStorage.setItem('isLoggedIn', true);
			localStorage.setItem('username', this.username);
			if (response.data.userID !== undefined) {
			const userId = response.data.userID;
			localStorage.setItem('userId', userId);
			console.log(userId)
			} else {
			console.error('userID is not defined in the response');
			}

			// Redirect to the home page
			this.$router.push('/')
		})
		.catch(error => {
			this.errormsg = error.toString();
		});
		},

		logout() {
			localStorage.setItem('isLoggedIn', false);
			localStorage.setItem('username', '');
			localStorage.setItem('userId', '');
			alert('You have successfully logged out.')
			}

	},
	mounted() {
		this.refresh()
	},
}
</script>

<template>
	<div class="login-page">
	  <div class="login-card">
		<h1 class="h2 login-title">WasaPhoto</h1>
		<div class="login-form">
		<h6>
			<span id = "type" class = "special">
			Login
			</span>
        </h6>
		<input id="username" type="text" v-model="username" placeholder="Username" />
		<button class="btn btn-sm-normal btn-outline-primary" @click="submit()">Enter</button>
		<p v-if="error">{{ error }}</p>
		
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		</div>
		<!--  <button id="logout-button" class="logout-button" @click="logout">Logout</button>  -->
	  </div>
	</div>
</template>
  


<style>

</style>

