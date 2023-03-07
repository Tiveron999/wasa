<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			username: "",
			
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
				const res = await this.$axios.get('/users', {params: {username: this.username}});
				console.log(res.data); // Log the response data
			} catch (err) {
				console.log(err);
			}
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
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="refresh">
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
		<button @click="getProfile">Search</button>
	</div>
</template>

<style scoped>
button {
  width: 5rem;
  height: 2rem;
  border-radius: 8px;
  background-color: rgb(255, 255, 75);
  text-align: center;
}
label {
  display: block;
  text-align: center;
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