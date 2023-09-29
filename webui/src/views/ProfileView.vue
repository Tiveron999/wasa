
<script>
import axios from 'axios'

import { toRaw } from 'vue';

// import UsernameModal from '../components/UsernameModal.vue';


export default {
	data: function() {
		return {
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
		}
	},
/*	components: {
		UsernameModal
	}, */
	watch: {
    // Watch for changes to the route or path
    '$route': 'onRouteChange'
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
		/*async uploadImage(event) {
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
	  },*/
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
		const response = await this.$axios.post(`/users/${this.id}/photos`, formData, 
		  {headers: { 'Content-Type': 'multipart/form-data', Authorization: 'Bearer ' + this.id}});
		  
		  /*
		  // The response data now contains an array of relative paths
			let newImage = JSON.stringify(response.data)
			this.images.push(newImage)
			console.log(newImage)
			// const imageUrl = URL.createObjectURL(newImage.image);
			// this.imageURLS[newImage.pid] = imageUrl;


			const binaryData = newImage.image

			const blob = new Blob([binaryData], { type: 'image/jpeg' }); 
			const imageUrl = URL.createObjectURL(blob);
			this.imageURLS[newImage.pid] = imageUrl;

			console.log(imageUrl) */
				
			await this.getImages()
		 /* console.log(response.data);
		  	const jsonData = response.data;
			const imageBlob = jsonData.image;

			// Create a Blob URL for the image Blob
			const imageUrl = URL.createObjectURL(imageBlob);

			// Push the URL into the array
			this.images.push(imageUrl);

			// Handle other parts of the JSON response as needed
			this.uploadedFiles = jsonData;
			*/

			// this.uploadedFiles = JSON.parse(JSON.stringify(response.data))

		} catch (error) {
		  console.error(error);
		}
	},	

		showModal(index) {
		this.selectedImage = this.images[index];
		console.log(this.selectedImage)
		this.show = true;
		
	},
		hideModal(index) {
		this.selectedImage = this.images[index];
		this.show = false;
	},
	  /*profileLink() {
      if (this.isLoggedIn) {
        return `/users/${localStorage.getItem('userId')}/profile`
      } else {
        return '/users/:user_id/profile'
      }
	},*/
	/*toggleLike(image) {
      // Update the likes property of the selected image
      image.likes = image.likes ? 0 : 1
	  if (localStorage.getItem('isLoggedIn')) {
		this.id = localStorage.getItem('userId');
		this.$axios.put(`/users/${this.id}/photos/{photo_id}/likes/${this.id}`)
      .then(response => {})
      .catch(error => {
        console.log(error);
      });
    }
	}, */

	async getImages() {
		if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}
		
		this.isLoading = true;
		console.log(this.isLoading)
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
				console.log(imageUrl);

				this.imageURLS[photo.pid] = imageUrl;
				}
				this.isLoading = false;
				console.log(this.isLoading)

		} catch (error) {
			console.log(error);
		}
	},

	async getDataURLFromImageURL(imageURL) {
	const response = await fetch(imageURL);
	let blob = await response.blob();
	if (blob.type !== 'image/jpeg') {
		blob = new Blob([blob], { type: 'image/jpeg' });
	}
	console.log(blob.type, response)
	const reader = new FileReader();
	return new Promise((resolve) => {
		reader.readAsDataURL(blob);
		reader.onload = () => {
		const dataURL = reader.result;
		console.log(dataURL)
		resolve(dataURL);
		};
	});
	},
	toggleLike(pid, posterid) {
		console.log("poster", posterid)
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
			this.getProfile()
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
				this.getProfile()
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
			try {
				const response = await this.$axios.get(`/users/${this.$route.params.user_id}/profile`, {headers: { Authorization: 'Bearer ' + this.id }});

				this.userProfile = response.data;
				console.log(this.userProfile)
			} catch (error) {
				// Handle any errors here
				console.error(error);
			}
		},

		onRouteChange(newRoute, oldRoute) {
			// You can access the new route and old route objects
			console.log('Route changed:', newRoute, oldRoute);
			if (
				oldRoute.params.user_id &&
				newRoute.params.user_id &&
				oldRoute.params.user_id !== newRoute.params.user_id &&
				oldRoute.params.profile &&
				newRoute.params.profile 
			) {
				// You are switching between user profiles
				this.getProfile()
				this.getImages()
			}
    	},




	
	},
	async created() {
		this.refresh()

		/*if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
			this.$axios.get(`/users/${this.id}/following/:following_id`, {headers: {Authorization: 'Bearer ' + this.id}})
		.then(response => {
			this.following = response.data.following;
			// console.log();
		})
		.catch(error => {
			console.log(error);
		});


		this.$axios.get(`/users/${this.id}/followers/:follower_id`, {headers: {Authorization: 'Bearer ' + this.id}})
		.then(response => {
			this.followers = response.data.followers;
		})
		.catch(error => {
			console.log(error);
		});

			}*/
		await this.getImages()
		// await this.resolveImageUrls()
		await this.getProfile()
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



		<div>

    
    <!-- 
    <button class="btn btn-primary" @click="openUsernameModal">Edit</button>


    <div class="modal" :class="{ 'show': isUsernameModalOpen }">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Change Username</h5>
            <button type="button" class="close" @click="closeUsernameModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="changeUsername">
              <div class="form-group">
                <label for="newUsername">New Username</label>
                <input
                  type="text"
                  class="form-control"
                  id="newUsername"
                  v-model="newUsername"
                  :pattern="usernamePattern"
                  title="Username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed."
                />
              </div>
              <div class="modal-footer">
                <button type="submit" class="btn btn-primary">Save</button>
                <button type="button" class="btn btn-secondary" @click="closeUsernameModal">Close</button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div> -->
 	</div>

		</div>
		
					


			<!-- fix modal appearing by clicking on a button instead of being static 
				<modal v-model="modalVisible" @modify="changeUsername">
				<div class="input-group mb-3 w-25" v-if="isSameUser()">
					<input
				type="text"
				class="form-control w-25"
				placeholder="Your new username..."
				v-model="newUsername"
				:pattern="usernamePattern"
				title="Username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed."
				/>

					<div class="input-group-append" v-if="isSameUser()">
					<button class="btn btn-outline-secondary" @click="changeUsername" >Modify</button>
					</div>
				</div>
				</modal>  --> 

				






		<div class="profile-container">
			<br>
			<div class="profile-info" v-if="userProfile">
				<h3 class="username"> {{ userProfile.username  }}</h3>
				<div class="follow-info">
				<p><strong>{{ userProfile.numphotos }}</strong> posts</p>
				<p><strong>{{ userProfile.followers  }}</strong> followers</p>
				<p><strong>{{ userProfile.following  }}</strong> following</p>
				</div>
				<button type="button" class="btn btn-sm-normal btn-outline-primary follow-button" v-if="!isSameUser()" @click="toggleFollow" 
				:class="{ following: isFollowing }">
				{{ isFollowing ? 'Following' : 'Follow' }}
				</button>
				<br>
				<button type="button" class="btn btn-sm-normal btn-outline-primary red-bg" v-if="!isSameUser()" @click="toggleBan" 
				:class="{ banned: isBanned }">
				{{ isBanned ? 'Unban' : 'Ban' }}
				</button>
			</div>
			<br> <br> <br> <br>
		</div>






		<button v-if="isSameUser()" class="btn btn-sm-normal btn-outline-primary" @click="openUsernameModal" data-bs-toggle="modal" data-bs-target="#UsernameModal">Edit</button>

		<div class="modal fade" id="UsernameModal" tabindex="-1" :class="{ 'show': isUsernameModalOpen }">
			<div class="modal-dialog">
				<div class="modal-content">
					<div class="modal-header">
						<h5 class="modal-title">Change Username</h5>
					</div>
					<div class="modal-body">
						<input
						type="text"
						class="form-control"
						placeholder="New Username"
						v-model="newUsername"
						:pattern="usernamePattern"
						title="Username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed."
						/>
					</div>
					<div class="modal-footer">
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
					<!--  <p v-if="isFirst"> Click to select </p> 
					<p v-if="isSaving"> Uploading image... </p>  -->
					<button class='btn btn-sm-normal btn-outline-primary' type="button" @click="$refs.pickImage.click()">Choose</button>
				</form>
			</div>
			<h5>
				<span id = "type" class = "normal">
					<br><br>Click on an image to view it in fullscreen mode, click on it again to return to the profile page.<br><br>
					<br>Here are the uploaded images: <br>
				</span>
			</h5>
			
			<div class="image-row">
				<!-- <div v-if="isLoading">	
					<LoadingSpinner></LoadingSpinner>
				</div> 
				<div v-else> 
					
				
				@click="showCommentModal = true"
			
			
			<CommentModal :showCommentModal="showCommentModal" @update:showCommentModal="showCommentModal = $event"></CommentModal>-->
				<img v-for="(image, index) in images" :key="image.pid" :src="imageURLS[image.pid]" :draggable="false" @click="showModal(index)"/>
				
			</div> 
		</div>
		<div id="Modal" v-if="show">
			<img :src="imageURLS[selectedImage.pid]" @click="hideModal(index)" />
			<div class="button-container">
				<button class="heart" @click="toggleLike(selectedImage.pid, selectedImage.posterid)" ></button>
				<div class="like-counter">{{ selectedImage.numlikes }} Likes</div>
				<div class="button-space"></div>
				<button class="cloud"> </button>
				
				<p>{{ showCommentModal }}</p>
				<div class="comment-counter">{{ selectedImage.numcomments }} Comments</div>	
				<div class="photo-date"> Upload date and time: {{ selectedImage.datetime }}</div>	
			</div>		
				<div class="delete-button-container">
					<button class="btn btn-sm-normal btn-outline-primary delete-photo" type="button" v-if="isSameUser()" @click="deletePhoto(selectedImage.pid)">Delete</button>
				</div>
		</div>     
		

	<!-- <div class="input-group-append">
		<button class="btn btn-sm btn-outline-primary" v-if="isSameUser()" @click="showUsernameModal()">Edit</button>
	</div>  -->
	<p v-if="error">{{ error }}</p>

	<!-- 
	<div class="modal fade" id="editUsernameModal" tabindex="-1" role="dialog" aria-labelledby="editUsernameModalLabel" aria-hidden="true" v-show="modalVisible">
		<div class="modal-dialog">
		<div class="modal-content">
			<div class="modal-header">
			<h5 class="modal-title" id="editUsernameModalLabel">Edit Username</h5>
			<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close" @click="hideUsernameModal()"></button>
			</div>
			<div class="modal-body">
			<form>
				<div class="mb-3">
				<label for="modal-input" class="form-label">New Username</label>
				<input type="text" class="form-control" id="modal-input" v-model="newUsername">
				</div>
			</form>
			</div>
			<div class="modal-footer">
			<button type="button" class="btn btn-secondary" data-bs-dismiss="modal" @click="hideUsernameModal()">Cancel</button>
			<button type="button" class="btn btn-primary" v-if="isSameUser()" @click="updateUsername()">Confirm</button>
			</div>
		</div>
		</div>
	</div>  -->

</div>

</template>


<style>

</style>

