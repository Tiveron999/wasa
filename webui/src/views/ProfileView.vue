<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
		}
	},
	methods: {
		/* this.$axios.post('/session', JSON.stringify({username: this.username}), {headers: { 'Content-Type': 'application/json',  'Accept': 'application/json', 
		'Authorization': `Bearer ${localStorage.getItem('userId')}`}}) */
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			// console.log(localStorage.getItem('userId'))
			try {
				let response = await this.$axios.get("/context");
				this.users = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
		async uploadImage(event) {
			const file = event.target.files[0];
			const formData = new FormData();
			formData.append('image', file);

			const reader = new FileReader();
			reader.readAsDataURL(file);
			reader.onload = function(event) {
			const img = document.createElement("img");
			img.src = event.target.result;
			document.body.appendChild(img);
	};
  
		try {
		  const response = await this.$axios.post("/users/:user_id/photos/:photo_id", formData);
		  console.log(response.data);
		} catch (error) {
		  console.error(error);
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
					<button type="button" class="btn btn-sm btn-outline-primary" @click="refresh">
						Refresh
					</button>
					
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<h6>
			<span id = "type" class = "special">
            Profile
			</span>
        </h6>
		<ul>
			<li> profile </li>
		</ul>
	</div>
	<div>
	  <form>
		<input type="file" @change="uploadImage"/>
	  </form>
	</div>
</template>

<style>
img {
  object-fit: cover;
  width: 100%;
  height: 100%;
}

button {
  width: 5rem;
  height: 2rem;
  border-radius: 8px;
  background-color: rgb(255, 255, 75);
  text-align: center;
}

</style>
