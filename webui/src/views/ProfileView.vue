

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" id="refresh-button" class="btn btn-sm btn-outline-primary" @click="refresh">
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
		<p>Followers: {{ followers }}</p>
		<p>Following: {{ following }}</p>
		<br>
	</div>
		<div id="image-container"> 
			<h5>
				<span id = "type" class = "normal">
				<br>Use the button below to add images to your profile.<br>
				</span>
			</h5>
			<form>
				<input type="file" @change="uploadImage2"/>
			</form>
			<h5>
				<span id = "type" class = "normal">
					<br><br>Click on an image to view it in fullscreen mode, click on it again to return to the profile page.<br><br>
					<br>Here are the uploaded images: <br>
				</span>
			</h5>
	
			 <div class="image-row">
	
		<img v-for="(image, index) in images" :key="index" :src="image" :draggable="false" v-if="images.length > 0" @click="showModal(index)"/>
			</div> 
		</div>
		<div id="Modal" v-if="show">
			<img :src="images[selectedImage]" @click="hideModal(index)" />
			<button class="heart"></button>
			<div class="like-counter">0 Likes</div>
		</div>     
	
</template>


<style>

</style>


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
			// images: [],
			show: false,
    		selectedImage: null,
			followers: 0,
      		following: 0,
			appearance: "default"
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
			console.log(this.images[i]);
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
		this.$axios.put(`/users/${this.id}/photos/{photo_id}/${this.id}`)
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
        this.images = dataURLs.map(dataURL => ({ src: dataURL }));
    })
    .catch(error => {
        console.log(error.response);
    });
	},
	async getDataURLFromImageURL(imageURL) {
		const response = await fetch(imageURL);
		console.log(response)
		const blob = await response.blob();
		const reader = new FileReader();
		return new Promise((resolve) => {
			reader.readAsDataURL(blob);
			reader.onload = () => {
				const dataURL = reader.result;
				resolve(dataURL);
				console.log(dataURL)
			};
		});
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
	/*created () {
  },*/
}


</script>