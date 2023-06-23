<template>
     <body id="body-pd">
     
      <div class="l-navbar" id="nav-bar">
        <nav class="nav">
            <div class="nav_list">
                <li class="nav-item">
                    <RouterLink to="/session" class="nav_link">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
                    Login
                    </RouterLink>
                </li>
                <li class="nav-item">
                    <RouterLink to="/" class="nav_link">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
                    Home
                    </RouterLink>
                </li>
                <li class="nav-item">
                    <RouterLink to="/users" class="nav_link">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#search"/></svg>
                    Search
                    </RouterLink>
                </li>
                <li class="nav-item">
                    <RouterLink to="/users/:user_id/profile" @click="updateProfileLink" class="nav_link">
                    <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#instagram"/></svg>
                    Profile
                    </RouterLink>
                </li>       
            </div>

              
          
          <a href="#" class="nav_link">
            <svg class="feather"><use href="/feather-sprite-v4.29.0.svg#log-out"/></svg>
            <span class="nav_name">Log out</span>
          </a>
      
        </nav>
      </div>
     
    </body>
  </template>
  
  <style>
  
  </style>
  
  <script>
  document.addEventListener("DOMContentLoaded", function(event) {   
  
   const linkColor = document.querySelectorAll('.nav_link')
   
   function colorLink(){
   if(linkColor){
   linkColor.forEach(l=> l.classList.remove('active'))
   this.classList.add('active')
   }
   }
   linkColor.forEach(l=> l.addEventListener('click', colorLink))
   });


  export default {
    data: function () {
        return {};
    },
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
  };
  </script>
  