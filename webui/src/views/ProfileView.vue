
<script>
import axios from 'axios'

import { toRaw } from 'vue';




/*window.onload = function() {
    let likeCounter = document.querySelector('.like-counter');
    let heartButton = document.querySelector('.heart');

    heartButton.addEventListener('click', function() {
        let currentLikes = parseInt(likeCounter.textContent);
        likeCounter.textContent = currentLikes + 1 + ' likes';
    });
}*/



/*const track = document.getElementById("image-track");

const handleOnDown = e => track.dataset.mouseDownAt = e.clientX;

const handleOnUp = () => {
  track.dataset.mouseDownAt = "0";  
  track.dataset.prevPercentage = track.dataset.percentage;
}

const handleOnMove = e => {
  if(track.dataset.mouseDownAt === "0") return;
  
  const mouseDelta = parseFloat(track.dataset.mouseDownAt) - e.clientX,
        maxDelta = window.innerWidth / 2;
  
  const percentage = (mouseDelta / maxDelta) * -100,
        nextPercentageUnconstrained = parseFloat(track.dataset.prevPercentage) + percentage,
        nextPercentage = Math.max(Math.min(nextPercentageUnconstrained, 0), -100);
  
  track.dataset.percentage = nextPercentage;
  
  track.animate({
    transform: `translate(${nextPercentage}%, -50%)`
  }, { duration: 1200, fill: "forwards" });
  
  for(const image of track.getElementsByClassName("image")) {
    image.animate({
      objectPosition: `${100 + nextPercentage}% center`
    }, { duration: 1200, fill: "forwards" });
  }
}

window.onmousedown = e => handleOnDown(e);

window.ontouchstart = e => handleOnDown(e.touches[0]);

window.onmouseup = e => handleOnUp(e);

window.ontouchend = e => handleOnUp(e.touches[0]);

window.onmousemove = e => handleOnMove(e);

window.ontouchmove = e => handleOnMove(e.touches[0]);
*/


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
			usernamePattern: "/^[A-Za-z][A-Za-z0-9_]{5,14}$/"
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
	this.$axios.get(`/users/${this.id}/photos`, {headers: {'Accept': 'application/json', Authorization: 'Bearer ' + this.id}})
	.then(async response => {
		console.log(response.data);
		// Update images data with retrieved image URLs
		const dataURLs = await Promise.all(response.data.map(url => this.getDataURLFromImageURL(url)));
		console.log(this.images)
		console.log(dataURLs)
		this.images = dataURLs;
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
    this.isUsernameModalOpen = true;
  },
  closeUsernameModal() {
    this.isUsernameModalOpen = false;
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
		if (error.response.status === 500) {
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
					
					<!-- <button type="button" class="btn btn-sm btn-outline-primary" v-if="isSameUser()" @click="openUsernameModal">
						Edit
					</button> -->


				<div v-if="isSameUser()" class="modal" :class="{ 'show': isUsernameModalOpen }">
				<div class="modal-dialog">
					<div class="modal-content">
					<div class="modal-header">
						<h5 class="modal-title">Change Username</h5>
						<button type="button" class="close" @click="closeUsernameModal">
						<span aria-hidden="true">&times;</span>
						</button>
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
						<button type="button" class="btn btn-primary" :disabled="isInvalid" @click="changeUsername">Save</button>
						<button type="button" class="btn btn-secondary" @click="closeUsernameModal">Close</button>
					</div>
					</div>
				</div>
				</div>


			<!-- fix modal appearing by clicking on a button instead of being static -->
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
				</modal>   

				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<h6>
			<span id = "type" class = "special">
            Profile
			</span>
        </h6>
	
	</div>
	<!-- <div id="image-track" data-mouse-down-at="0" data-prev-percentage="0">   -->
 <!--  <img class="image" src="https://images.unsplash.com/photo-1610194352361-4c81a6a8967e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1674&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1618202133208-2907bebba9e1?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1495805442109-bf1cf975750b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1548021682-1720ed403a5b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1496753480864-3e588e0269b3?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2134&q=80" draggable="false" />
   
-->
	<div>
		<br>
		<p>Posts: {{ posts }}</p>
		<p>Followers: {{ followers }}</p>
		<p>Following: {{ following }}</p>
		<br>
	</div>
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
		<img v-for="(image, index) in images" :key="index" :src="image" :draggable="false" @click="showModal(index)"/>
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

