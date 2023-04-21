import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'




localStorage.clear();
// Initialize the variables
localStorage.getItem('isLoggedIn') === 'true' || false,
localStorage.getItem('username') || '',
localStorage.getItem('userId') || ' '



const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: HomeView, beforeEnter: (to, from, next) => {
			if (localStorage.getItem('isLoggedIn') === 'true') {
			  next()
			} else {
			  next('/session')
			}
		  }
		},
		  {path: '/users/:user_id/profile', component: ProfileView,  name: 'Profile',
		  beforeEnter: (to, from, next) => {
			if (localStorage.getItem('isLoggedIn') === 'true') {
				/*const userId = localStorage.getItem('userId');
				if (userId) {
				  to.params.user_id = userId;
				}*/
				next()
			  } else {
				next('/session')
			  }
			}},
		{path: '/session', component: LoginView},
		{ path: '/users', component: SearchView, beforeEnter: (to, from, next) => {
		if (localStorage.getItem('isLoggedIn') === 'true') {
			next()
		} else {
			next('/session')
		}
		}
	},
		
	]
})

export default router
