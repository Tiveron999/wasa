<script>
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
			profileLink: '',
			images: [],
			show: false,
    		selectedImage: null
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
			};
	  }
	  try {
		  const response = await this.$axios.post("/users/:user_id/photos/:photo_id", formData);
		  console.log(response.data);
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
	mounted() {
		this.refresh()
	},
	/*created () {
    // Get the current login status from local storage
    this.isLoggedIn = localStorage.getItem('isLoggedIn') || false
  },*/
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
	<!-- <div id="image-track" data-mouse-down-at="0" data-prev-percentage="0">   -->
 <!--  <img class="image" src="https://images.unsplash.com/photo-1610194352361-4c81a6a8967e?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1674&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1618202133208-2907bebba9e1?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1495805442109-bf1cf975750b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1548021682-1720ed403a5b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80" draggable="false" />
  <img class="image" src="https://images.unsplash.com/photo-1496753480864-3e588e0269b3?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=2134&q=80" draggable="false" />
   
-->
   	
		<div id="image-container"> 
			<form>
				<input type="file" @change="uploadImage2"/>
			</form>
			<br>Here are the uploaded images: <br>
	
			 <div class="image-row">
		<img v-for="(image, index) in images" :key="index" :src="image" :draggable="false" @click="showModal(index)"/>
			</div> 
		</div>
		<div id="Modal" v-if="show">
			<img :src="images[selectedImage]" @click="hideModal(index)" />
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

img {
	object-fit: cover;
	width: 33%;
	height: 300px;
	float: left;
	margin: 0.1% 0.1%;
}

.image-row {
	display: flex;
	flex-wrap: wrap;
	float: left;
	width: 100%;
	overflow: auto;
}

#image-container {
  object-fit: cover;
  width: 100%;
  height: 100%;
  overflow: auto;
  margin: 0 auto;
  display: inline-block; 
  user-select: none;
}

#Modal {
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
  z-index: 1;
  position: fixed;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  overflow: auto;
}

#Modal img {
	object-fit: contain;
	margin: auto;
	display: flex;
	align-items: center;
	justify-content: center;
	width: 50%;
	height: auto;
	max-width: 80%;
  	max-height: 80%;
}

/*
#image-track {
  display: flex;
  gap: 4vmin;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(0%, -50%); 
  user-select: none; 
}

#image-track > .image {
  width: 40vmin;
  height: 56vmin;
  object-fit: cover;
  object-position: 100% center;
}
*/

</style>
