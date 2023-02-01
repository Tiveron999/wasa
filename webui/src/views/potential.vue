<script>
  import axios from 'axios';
  
  export default {
	methods: {
	  async uploadImage(event) {
		const file = event.target.files[0];
		const formData = new FormData();
		formData.append('image', file);
  
		try {
		  const response = await this.$axios.post("/users/:user_id/photos/:photo_id", formData);
		  console.log(response.data);
		} catch (error) {
		  console.error(error);
		}
	  }
	}
  };
  </script>


<template>
	<div>
	  <form>
		<input type="file" @change="uploadImage"/>
	  </form>
	</div>
  </template>
  
  
  