
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
			liked: false,
			submitted: false,
			text: 'Like',
			username: '',
			modalVisible: false,
			posts: 0,
			isUsernameModalOpen: false,
   			newUsername: '',
			usernamePattern: "/^[A-Za-z][A-Za-z0-9_]{5,14}$/",
			isFollowing: false,
		}
	},
/*	components: {
		UsernameModal
	}, */

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
			console.log(reader.result);
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

		showModal(index) {
		this.selectedImage = index;
		this.show = true;
		
	},
		hideModal(index) {
		this.selectedImage = index;
		this.show = false;
	},
	  /*profileLink() {
      if (this.isLoggedIn) {
        return `/users/${localStorage.getItem('userId')}/profile`
      } else {
        return '/users/:user_id/profile'
      }
	},*/
	toggleLike(image) {
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
	},
	getImages() {
		if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
		}

		this.$axios.get(`/users/${this.id}/photos`, {
			headers: { 'Accept': 'application/json', 'Authorization': 'Bearer ' + this.id }
		})
		.then(response => {
			console.log('hi2', response.data);

			// The response data now contains an array of relative paths
			this.images = response.data;
			console.log('images:', this.images);
		})
		.catch(error => {
			console.log(error.response);
		});
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
	toggleLike2: function()
	{
		if(this.liked) {
			this.unlikePhoto()
		} else {
			this.likePhoto()
		}
	},

	likePhoto: function() {
		this.submitted = true;

		this.$axios.put(this.$axios.put(`/users/${this.id}/photos/{photo_id}/likes/${this.id}`), function(resp) {
			this.liked = true;
			this.submitted = false;
			this.text = 'Unlike';
		});
	},

	unlikePhoto: function() {
		this.submitted = true;

		this.$axios.delete(`/users/${this.id}/photos/{photo_id}/likes/${this.id}`, function(resp) {
			this.liked = false;
			this.submitted = false;
			this.text = 'Like';
		});
	},
	showUsernameModal() {
      this.modalVisible = true;
    },
    hideUsernameModal() {
      this.modalVisible = false;
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
			await this.$axios.put(`/users/${this.id}/following/${this.$route.params.user_id}`, {headers: {Authorization: 'Bearer ' + this.id}})
			this.following += 1
		} catch (e) {
			console.error(e);
			this.errormsg = e.toString();
		}
		},
		unfollowUser() {
		// 
		},

	
	},
	mounted() {
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
	this.getImages();
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
			<div class="profile-info">
				 <!-- <h2>{{ username }}</h2> -->
				<div class="follow-info">
				<p><strong>{{ posts }}</strong> posts</p>
				<p><strong>{{ followers }}</strong> followers</p>
				<p><strong>{{ following }}</strong> following</p>
				</div>
				<button type="button" class="btn btn-sm-normal btn-outline-primary follow-button" v-if="!isSameUser()" @click="toggleFollow" 
				:class="{ following: isFollowing }">
				{{ isFollowing ? 'Following' : 'Follow' }}
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
				<form>
					<input style = "display:none" type="file" @change="uploadImage2" ref = "pickImage"/>
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
		<img v-for="(image, index) in images" :key="index" :src="`#/${image}`" :draggable="false" @click="showModal(index)"/>
			</div> 
		</div>
		<div id="Modal" v-if="show">
			<img :src="images[selectedImage]" @click="hideModal(index)" />
			<button class="heart" @click="toggleLike2" ></button>
			<div class="like-counter">{{ images.likes }} Likes</div>
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

		

</template>


<style>

</style>

