<script>

export default {
    /*components: {
        LoginView,
        HomeView,
    },*/
    data: function () {
        return {
            errormsg: null,
            loading: false,
            isLoggedIn: false,
            images: [],
			show: false,
            selectedImage: null,
			imageURLS : {},
        };
    },
    methods: {
        async refresh() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.get("/context"); // .then(response => {console.log(response)});
                this.users = response.data;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async checkLoginStatus() {
            try {
                const response = await this.$axios.post("/session", {
                    username: this.username,
                });
                if (response.data.userID) {
                    this.isLoggedIn = true;
                }
                else {
                    this.isLoggedIn = false;
                }
            }
            catch (error) {
                console.error(error);
                this.isLoggedIn = false;
            }
        },
        async getFeed() {
		if (localStorage.getItem('isLoggedIn')) {
			this.id = localStorage.getItem('userId');
            console.log("id", this.id)
		}
		
		this.isLoading = true;
		console.log(this.isLoading)
		try {
			const response = await this.$axios.get(`/users/${this.id}/feed`, {
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
    showModal(index) {
		this.selectedImage = this.images[index];
		console.log(this.selectedImage)
		this.show = true;
		
	},
		hideModal(index) {
		this.selectedImage = this.images[index];
		this.show = false;
	},
    },
    mounted() {
        this.refresh();
        /*this.checkLoginStatus();*/
    },
    async created() {
		this.refresh()

		await this.getFeed()
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
            Home
			</span>
        </h6>
        
		<h5>
            <span id = "type" class = "normal">
                <br>Click on an image to view it in fullscreen mode, click on it again to return to the profile page.<br>
                <br>Here are your followers' uploaded images: <br>
            </span>
		</h5>

        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

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
	
				<div class="comment-counter">{{ selectedImage.numcomments }} Comments</div>	
				<div class="photo-date"> Upload date and time: {{ selectedImage.datetime }}</div>	
			</div>		

    </div>
</template>

<style>

</style>
