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

<body class="login-page">
  


	<div id="login-container">
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" id="refresh-button" class="btn btn-sm btn-outline-primary" style="float: right;" @click="refresh">
						Refresh
					</button>
				</div>
					
					
				<div class="btn-group me-2">
					
				<button id="logout-button" @click="logout">Logout</button>
				</div>  
			</div>
		</div>
			
			
	  
	</div>
		

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		
		<div>
     <!-- <template v-if="isLoggedIn">
         Show content for logged in users 
      </template>
      <template v-else>-->
		<h6>
			<span id = "type" class = "special">
            Login
			</span>
        </h6>
		<!-- <base-card> 
		<form @submit.prevent="submit"> -->
		  <!-- <label for="username">username</label> -->
		  <input id="username" type="text" v-model="username" placeholder="Username"/>
		  <button class="btn btn-sm-normal btn-outline-primary" @click="submit()">Enter</button>  <!--@click="loginButtonClicked" {{ buttonString }} type="submit" -->
		  <p v-if="error">{{ error }}</p>
		 <!--  <a href="#" @click.prevent="toggleLogin">{{ textLoginString }}</a> 
		</form> -->
	 <!-- </base-card> -->
	

    </div>
		<!--<h6>
			<span id = "type" class = "special">
            Login
			</span>
        </h6>
		<base-card>
		<form @submit.prevent>
		  <label for="username">username</label>
		  <input id="username" type="text" v-model="username" />
		  <button class="btn btn-sm btn-outline-secondary" @click="loginButtonClicked">{{ buttonString }}Enter</button>
		  <p v-if="error">{{ error }}</p>
		  <a href="#" @click.prevent="toggleLogin">{{ textLoginString }}</a>
		</form>
	  </base-card>-->

	</div>
	
</body>	
</template>


<style>

</style>

