<script>
import axios from 'axios'

export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			users: null,
			isLoggedIn: false,
        	username: '',
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
		async checkLoginStatus() {
        try {
          // Make an HTTP request to the backend to check if the user is logged in
          const response = await axios.post('/api/session', {
            username: this.username,
          })
          const userID = response.data.userID
          if (userID) {
            this.isLoggedIn = true
          } else {
            this.isLoggedIn = false
		}
		} catch (error) {
		console.error(error)
		}
	},
	},
	mounted() {
		this.refresh()
	},
	created() {
		this.checkLoginStatus()
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
			<div class="content">
	  
	</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div>
      <template v-if="isLoggedIn">
        <!-- Show content for logged in users -->
      </template>
      <template v-else>
		<h6>
			<span id = "type" class = "special">
            Login
			</span>
        </h6>
		<base-card>
		<form @submit.prevent="checkLoginStatus">
		  <!-- <label for="username">Username</label> -->
		  <input id="username" type="text" v-model="username" placeholder="Username"/>
		  <button class="btn btn-sm btn-outline-secondary" type="submit" @click="loginButtonClicked">{{ buttonString }}Enter</button>
		  <p v-if="error">{{ error }}</p>
		  <a href="#" @click.prevent="toggleLogin">{{ textLoginString }}</a>
		</form>
	  </base-card>
	
        <!--  <form @submit.prevent="checkLoginStatus">
          <input v-model="username" type="text" placeholder="Username" />
          <button type="submit">Log In</button>
        </form> -->
      </template>
    </div>
		<!--<h6>
			<span id = "type" class = "special">
            Login
			</span>
        </h6>
		<base-card>
		<form @submit.prevent>
		  <label for="username">Username</label>
		  <input id="username" type="text" v-model="username" />
		  <button class="btn btn-sm btn-outline-secondary" @click="loginButtonClicked">{{ buttonString }}Enter</button>
		  <p v-if="error">{{ error }}</p>
		  <a href="#" @click.prevent="toggleLogin">{{ textLoginString }}</a>
		</form>
	  </base-card>-->
	</div> 
</template>


<style scoped>
button {
  width: 4rem;
  height: 2rem;
  border-radius: 8px;
  background-color: rgb(255, 255, 75);
  text-align: center;
}
label {
  display: block;
  text-align: left;
}
input {
  box-sizing: border-box;
  border-radius: 8px;
  width: 20%;
  padding: 8px 20px;
  margin-right: 5%;
}
input:focus {
  outline: 0;
  background-color: rgb(255, 175, 26);
}
</style>

