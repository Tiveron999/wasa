<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
	methods: {
		updateProfileLink() {
			console.log(localStorage.getItem('userId'))
			if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
			this.$router.push({ name: 'Profile', params: { user_id: this.id } });
			} else {
			this.$router.push("/users/:user_id/profile");
			}
		}
},
/*computed: {
    profileUrl() {
      // return the URL with the actual user id
      return `/users/${this.userId}/profile`;
    }
<<<<<<< HEAD
  },*/

  data() {
    return {
      userId: null
    };
  },
  created() {
    // get the user id from local storage
    this.userId = localStorage.getItem("userId")
	console.log(this.userId)
  },

}

</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WasaPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/session" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								Login
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								Home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/search" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
								Search
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/users/:user_id/profile" @click="updateProfileLink" class="nav-link"> 
							<!-- <router-link class="nav-link" :to="{ name: 'Profile', params: { user_id: localStorage.getItem('userId')}}"> </router-link> 
							<RouterLink :to="{ name: 'Profile', params: { user_id: localStorage.getItem('userId')}}" class="nav-link"> -->

								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#instagram"/></svg>
								Profile
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/link1" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								Menu item 1
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/link2" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								Menu item 2
							</RouterLink>
						</li>
					</ul>

					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mt-4 mb-1 text-muted text-uppercase">
						<span>Secondary menu</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<!-- <RouterLink :to="'/some/' + 'variable_here' + '/path'" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#file-text"/></svg>
								Item 1
							</RouterLink> -->
						</li>
					</ul>
				</div>
			</nav>

			<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
				<RouterView />
			</main>
		</div>
	</div>
</template>

<style>
</style>
