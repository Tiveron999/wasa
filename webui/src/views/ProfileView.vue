
<script>
import axios from 'axios'

import { toRaw } from 'vue';



export default {
	data: function() {
		return {
			youAreBanned: false,
			errormsg: null,
			loading: false,
			isLoggedIn: false,
			userId: '',
			error: null,
			images: [],
			show: false,
    		selectedImage: null,
			followers: 0,
      		following: 0,
			appearance: "default",
			isLiked: false,
			likes: 0,
			submitted: false,
			text: 'Like',
			username: '',
			posts: 0,
			isUsernameModalOpen: false,
   			newUsername: '',
			usernamePattern: "/^[A-Za-z][A-Za-z0-9_]{5,14}$/",
			isFollowing: false,
			isBanned: false,
			uploadedFilesa: null, 
			imageURLS : {},
			userProfile: null,
			showCommentModal: false,
			isLoading: false,
			showLikesModal: false,
			showFollowersModal: false,
			showFollowingModal: false,
			followers_list: [],
			following_list: [],
		}
	},

	watch: {
    // Watch for changes to the route or path
    '$route': 'onRouteChange'
  	},

	methods: {
		async refresh() {
			this.youAreBanned = false,
			this.isBanned = false,
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
		
	  async uploadImage2(event) {
		const formData = new FormData();
		const files = event.target.files;
		for (let i = 0; i < files.length; i++) {
			formData.append('image', files[i]);
			const reader = new FileReader();
			reader.readAsDataURL(files[i]);
			reader.onload = () => {
			this.images.push(reader.result);
			// console.log(reader.result);
			};
	  }
	  try {
		if (localStorage.getItem('isLoggedIn')) {
		this.id = localStorage.getItem('userId');
	}
		  const response = await this.$axios.post(`/users/${this.id}/photos`, formData, 
		  {headers: { 'Content-Type': 'multipart/form-data', Authorization: 'Bearer ' + this.id}});
		  // console.log(response.data.photoId);
		  // console.log(toRaw(this.images));
		  
		} catch (error) {
		  console.error(error);
		}
	},

	async uploadImage3(event) {
		let imagefiles = event.target.files;
		const formData = new FormData();
		const imageName = 'image'; 

		if (imagefiles.length != 1) return 

		formData.append(imageName, imagefiles[0], imagefiles[0].name);

		if (localStorage.getItem('isLoggedIn')) {
		this.id = localStorage.getItem('userId');
		}

		try {
			await this.$axios.post(`/users/${this.id}/photos`, formData, 
		 	 {headers: { 'Content-Type': 'multipart/form-data', Authorization: 'Bearer ' + this.id}});
		  
			await this.getImages()
			await this.getProfile()
		
		} catch (error) {
		  console.error(error);
		}
	},	

		showModal(index) {
		this.selectedImage = this.images[index];
		this.show = true;
		
	},
		hideModal(index) {
		this.selectedImage = this.images[index];
		this.show = false;
	},

	async getImages() {
		if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
		
		this.isLoading = true;
	
		
		try {
			const response = await this.$axios.get(`/users/${this.$route.params.user_id}/photos`, {
				headers: { 'Accept': 'application/json', 'Authorization': 'Bearer ' + this.id }
			})

			console.log("over here", response.data)

			 // Assuming response.data is an array of photo objects
			 this.images = response.data;

			 if (this.images === null){
				this.images = []
			 }
			 

			for (let i = 0; i < this.images.length; i++) {
				const photo = this.images[i];

				// Create a Blob from the Base64 data
				const binaryImageData = atob(photo.image);
				const dataArray = new Uint8Array(binaryImageData.length);

				for (let i = 0; i < binaryImageData.length; i++) {
				dataArray[i] = binaryImageData.charCodeAt(i);
				}
				const blob = new Blob([dataArray], { type: 'image/jpeg' });

				// Create a data URL from the Blob
				const imageUrl = URL.createObjectURL(blob);

				// imageUrl now contains the data URL for the image


				this.imageURLS[photo.pid] = imageUrl;
				}
				this.isLoading = false;
		

		} catch (error) {
			console.log(error);
		}
	},

	toggleLike(pid, posterid) {
		if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
	
		if (this.isLiked) {
			this.unlikePhoto(pid, posterid);
		} else {
			this.likePhoto(pid, posterid);
		}
		// Toggle the liked state
		this.isLiked = !this.isLiked;
	},

	async likePhoto(pid, posterid) {
		try {
			const response = await this.$axios.put(`/users/${posterid}/photos/${pid}/likes/${this.id}`, null, {headers: {Authorization: 'Bearer ' + this.id}})

			if (response.status === 200) {
			const data = response.data;
			const likesCount = data.likes; // Access the "likes" property from the response

			// Now "likesCount" contains the updated number of likes
			console.log(`Number of likes: ${likesCount}`);

			this.selectedImage.numlikes = likesCount;
			}
			
		} catch (e) {
			console.error("error: ", e);
			this.errormsg = e.toString();
		}
		},
		async unlikePhoto(pid, posterid) {
			try {
				const response = await this.$axios.delete(`/users/${posterid}/photos/${pid}/likes/${this.id}`, {headers: { Authorization: 'Bearer ' + this.id }});

			if (response.status === 200) {
			const data = response.data;
			const likesCount = data.likes; // Access the "likes" property from the response

			// Now "likesCount" contains the updated number of likes
			console.log(`Number of likes: ${likesCount}`);

			this.selectedImage.numlikes = likesCount;
			}

			} catch (e) {
				console.error("error: ", e);
				this.errormsg = e.toString();
			}
		},
	
	openUsernameModal() {
		console.log('here1: ' + this.isUsernameModalOpen);
		this.isUsernameModalOpen = true;
		console.log('here2: ' + this.isUsernameModalOpen);
	},
	closeUsernameModal() {
		console.log('here1: ' + this.isUsernameModalOpen);
		this.isUsernameModalOpen = false;
		console.log('here2: ' + this.isUsernameModalOpen);
	},

	changeUsername() {
		if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
		// Username is not valid
		const regex = /^[A-Za-z][A-Za-z0-9_]{5,14}$/;
		if (!regex.test(this.newUsername)) {
			alert('The username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed.');
			return;
		}
		// Make PUT request to update username
		this.$axios.put(`/users/${this.id}`, { username: this.newUsername }, {headers: {Authorization: 'Bearer ' + this.id}})
		.then(() => {
			this.username = ""
			this.isUsernameModalOpen = false;
			alert('You have successfully changed your username.')
			this.getProfile()
		})
		.catch(error => {

			if (error.response.status === 500) {                    // change error code
				// This username already exists error
				alert('This username already exists.');
			} else {
				// Other errors
				console.error(error);
			}
		});
	},

	isSameUser() {
			const userId = this.$route.params.user_id;
			const storedUserId = localStorage.getItem('userId');
			return userId === storedUserId;
		},

		toggleFollow() {
			if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
		// If the user is following, unfollow on click
		if (this.isFollowing) {
			this.unfollowUser();
		} else {
			this.followUser(); // Follow the user on click
		}

		// Toggle the following state
		this.isFollowing = !this.isFollowing;
		},
		async followUser() {
		try {
			console.log(this.id, this.$route.params.user_id)
			const result = await this.$axios.put(`/users/${this.id}/following/${this.$route.params.user_id}`, null, {headers: {Authorization: 'Bearer ' + this.id}})
			console.log(result)
			this.followers += 1
			console.log("followers", this.followers)
			await this.getProfile()
		} catch (e) {
			console.error("error: ", e);
			this.errormsg = e.toString();
		}
		},
		async unfollowUser() {
			try {
				const result = await this.$axios.delete(`/users/${this.id}/following/${this.$route.params.user_id}`, {headers: { Authorization: 'Bearer ' + this.id }});
				console.log(result);
				this.followers -= 1;
				await this.getProfile()
			} catch (e) {
				console.error("error: ", e);
				this.errormsg = e.toString();
			}
		},

		async deletePhoto(photoId) {
			try {
				const response = await this.$axios.delete(`/users/${this.id}/photos/${photoId}`, {
				headers: {
					'Content-Type': 'application/json', "Authorization": 'Bearer ' + this.id 
				},
				});

				if (response.status === 204) {
				// Photo deleted successfully
				console.log(`Photo with ID ${photoId} deleted successfully.`);
				alert('Photo deleted successfully.');
				this.show = false;
				this.getImages()
				} else if (response.status === 404) {
				// The photo does not exist
				console.error(`Photo with ID ${photoId} does not exist.`);
				alert('This photo does not exist');
				} else {
				console.error(`Error deleting photo with ID ${photoId}. Status: ${response.status}`);
				}
			} catch (error) {
				console.error(`Error deleting photo with ID ${photoId}: ${error.message}`);
			}
		},
		
		async getProfile() {
			if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
			try {
				const response = await this.$axios.get(`/users/${this.$route.params.user_id}/profile`, {headers: { Authorization: 'Bearer ' + this.id }});

				this.userProfile = response.data;
				console.log(this.userProfile)
			} catch (error) {
				// Handle any errors here
				console.error(error);
			}
		},
		async getYourBans() {
			if (localStorage.getItem('isLoggedIn')) {
			this.id = Number(localStorage.getItem('userId'));
		}
			// save time
			if(this.isSameUser()){
				return 
			}
			try {
				const response = await this.$axios.get(`/users/${this.id}/bans`, {headers: { Authorization: 'Bearer ' + this.id }});

				let bans = response.data;
	
				for (let i = 0; i < bans.length; i++) {
	
					if (bans[i].userid===this.id && bans[i].bannedid === Number(this.$route.params.user_id)){
				
						this.isBanned=true
						break
					}
				};
				
	 			
			} catch (error) {
				// Handle any errors here
				console.error(error);
			}
		},

		async getSearchBans() {
			if (localStorage.getItem('isLoggedIn')) {
			this.id = Number(localStorage.getItem('userId'));
		}
			// save time
			if(this.isSameUser()){
				return 
			}

			try {
				const response = await this.$axios.get(`/users/${this.$route.params.user_id}/bans`, {headers: { Authorization: 'Bearer ' + this.id }});

				let bans = response.data;

				for (let i = 0; i < bans.length; i++) {

					if (bans[i].userid === Number(this.$route.params.user_id) && bans[i].bannedid === this.id) {

						this.youAreBanned = true;
						break
						}

				};
				
			} catch (error) {
				// Handle any errors here
				console.error(error);
			}
		},

		toggleBan() {
			if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
		// If the user is banned, unban on click
		if (this.isBanned) {
			this.unbanUser();
		} else {
			this.banUser(); // Ban the user on click
		}

		// Toggle the ban state
		this.isBanned = !this.isBanned;
		},
		async banUser() {
		try {

			await this.$axios.put(`/users/${this.id}/banned/${this.$route.params.user_id}`, null, {headers: {Authorization: 'Bearer ' + this.id}})
		} catch (e) {
			console.error("error: ", e);
			this.errormsg = e.toString();
		}
		},
		async unbanUser() {
			try {
				await this.$axios.delete(`/users/${this.id}/banned/${this.$route.params.user_id}`, {headers: { Authorization: 'Bearer ' + this.id }});
			} catch (e) {
				console.error("error: ", e);
				this.errormsg = e.toString();
			}
		},

		onRouteChange(newRoute, oldRoute) {
			// Access the new route and old route objects
			console.log('Route changed:', newRoute, oldRoute);
			if (
				oldRoute.params.user_id &&
				newRoute.params.user_id &&
				oldRoute.params.user_id !== newRoute.params.user_id
			) {
				console.log('switched');
				// You are switching between user profiles
				this.refresh()
				this.getYourBans()
				this.getSearchBans()
				this.getProfile()
				this.getImages()
			}
    	},
		updateComments(newComments) {
			console.log(`New number of comments: ${newComments}`);
			this.selectedImage.numcomments = newComments
		},

		updateIsFollowing(flag) {
			console.log("check pt 2:", flag)
			this.isFollowing = flag
		},
	
	},
	async mounted() {
		this.refresh()
		
		console.log("are you?", this.youAreBanned)

		await this.getYourBans()
		await this.getSearchBans()
		await this.getProfile()
		await this.getImages()
	},
	computed: {
			/*usernamePattern() {
		return "^[A-Za-z][A-Za-z0-9_]{5,14}$";
	}, */
		isInvalid() {
			return !this.newUsername.match(this.usernamePattern);
		}
	},


}

</script>




<template>

<div>
	<div>
		<div>
			<div
				class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
				<h1 class="h2">WasaPhoto</h1>
				<div class="btn-toolbar mb-2 mb-md-0">
					<div class="btn-group me-2">				
						
					</div>
				</div>

			</div>
			<h6>
				<span id = "type" class = "special">
				Profile
				</span>
			</h6>
		

			<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

	

			</div>
			


		<div v-if=!this.youAreBanned>


			<div class="profile-container">
				<br>
				<div class="profile-info" v-if="userProfile">
					<h3 class="username"> {{ userProfile.username  }}</h3>
					<div class="follow-info">
					<p><strong>{{ userProfile.numphotos }}</strong> posts</p>
					
					<div class="follower-container">
						<div @click="showFollowersModal = true" class="follower-button">
							<p><strong>{{ userProfile.followers }}</strong> followers</p>
						</div>
						<FollowersModal
							:showFollowersModal="showFollowersModal"
							:userID="userProfile.userid" :sameUser="isSameUser()"
							@update:showFollowersModal="showFollowersModal = $event" 
							@following-check="updateIsFollowing"
							/>
					</div>
	
					<div class="following-container">
						<div @click="showFollowingModal = true" class="following-button">
							<p><strong>{{ userProfile.following }}</strong> following</p>
						</div>
						<FollowingModal
							:showFollowingModal="showFollowingModal"
							:userID="userProfile.userid" 
							@update:showFollowingModal="showFollowingModal = $event"
							/>
					</div>
					</div>
					<button type="button" class="btn btn-sm-normal btn-outline-primary follow-button bump-up-buttons" v-if="!isSameUser()" @click="toggleFollow" 
					:class="{ following: isFollowing }">
					{{ isFollowing ? 'Following' : 'Follow' }}
					</button>
					<br>
					<button type="button" class="btn btn-sm-normal btn-outline-primary red-bg bump-up-buttons" v-if="!isSameUser()" @click="toggleBan" 
					:class="{ banned: isBanned }">
					{{ isBanned ? 'Unban' : 'Ban' }}
					</button>
				</div>
				<br> <br> <br> <br>
			</div>




			<button v-if="isSameUser()" class="btn btn-sm-normal btn-outline-primary" @click="openUsernameModal">Edit</button>
			<div class="username-modal fade" id="UsernameModal" tabindex="-1" v-if="isUsernameModalOpen" :class="{ 'show': isUsernameModalOpen }">
				<div class="username-modal-dialog">
					<div class="username-modal-content">
						<div class="username-modal-header">
							<h5 class="username-modal-title">Change Username</h5>
						</div>
						<div class="username-modal-body">
							<input
							type="text"
							class="form-control"
							placeholder="New Username"
							v-model="newUsername"
							:pattern="usernamePattern"
							title="Username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed."
							/>
						</div>
						<div class="username-modal-footer">
							<button type="button" class="btn btn-sm-normal btn-outline-secondary green-bg"  @click="changeUsername">Save</button>
							<button type="button" class="btn btn-sm-normal btn-outline-primary red-bg" @click="closeUsernameModal">Close</button>

						</div>
					</div>
				</div>
			</div>



			<br> <br><br>
			<div id="image-container"> 
				<div v-if="isSameUser()"> 
					<h5>
						<span id = "type" class = "normal">
						<br>Use the button below to add images to your profile.<br>
						</span>
					</h5>
					<form enctype="multipart/form-data">
						<input style = "display:none" type="file"  @change="uploadImage3" ref = "pickImage" accept="image/jpeg"/>
						<!-- <p v-if="isSaving"> Uploading image... </p>  -->
						<button class='btn btn-sm-normal btn-outline-primary' type="button" @click="$refs.pickImage.click()">Choose</button>
					</form>
				</div>
				<h5>
					<span id = "type" class = "normal">
						<br><br>Click on an image to view it in fullscreen mode, click on it again to return to the profile page.<br><br>
						<br>Here are the uploaded images: <br>
					</span>
				</h5>
				

				<LoadingSpinner v-if="isLoading" />
					<div v-else> 

						
				<div class="image-row">
					
					<img v-for="(image, index) in images" :key="image.pid" :src="imageURLS[image.pid]" :draggable="false" @click="showModal(index)"/>
					
				</div> 
			</div>
			<div id="Modal" v-if="show">
				<img :src="imageURLS[selectedImage.pid]" @click="hideModal(index)" />
				<div class="button-container">
					<button class="heart" @click="toggleLike(selectedImage.pid, selectedImage.posterid)" ></button>
					<button class="like-counter-button" @click="showLikesModal = true"> {{ selectedImage.numlikes }} Likes</button>
    				<LikesModal :showLikesModal="showLikesModal" 
					:posterID="selectedImage.posterid" :pid="selectedImage.pid" 
					@update:showLikesModal="showLikesModal = $event" />
					<div class="button-space"></div>
					<button class="cloud" @click="showCommentModal = true"> </button>
					<CommentModal :showCommentModal="showCommentModal" 
					:posterID="selectedImage.posterid" :pid="selectedImage.pid" 
					@update:showCommentModal="showCommentModal = $event"
					@update-comments="updateComments"></CommentModal>
					<div class="comment-counter">{{ selectedImage.numcomments }} Comments</div>	
					<div class="photo-date"> Upload date and time: {{ selectedImage.datetime }}</div>	
				</div>		
					<div class="delete-button-container">
						<button class="btn btn-sm-normal btn-outline-primary delete-photo" type="button" v-if="isSameUser()" @click="deletePhoto(selectedImage.pid)">Delete</button>
					</div>
			</div>     
		</div>

		<p v-if="error">{{ error }}</p>

		</div>
		<div v-else>
			<span id = "type" class="normal" style="white-space: nowrap;">
			<br>You are banned by this user.<br>
			</span>
		</div>


	</div>

</div>

</template>


<style>

</style>

