import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'



// Initialize the variables
localStorage.getItem('isLoggedIn') === 'true' || false,
localStorage.getItem('username') || ''


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
		  {path: '/profile', component: ProfileView, 
		  beforeEnter: (to, from, next) => {
			if (localStorage.getItem('isLoggedIn') === 'true') {
				next()
			  } else {
				next('/session')
			  }
			}},
		{path: '/link1', component: HomeView},
		{path: '/link2', component: HomeView},
		{path: '/some/:id/link', component: HomeView, beforeEnter: (to, from, next) => {
			if ((localStorage.getItem('isLoggedIn') === 'true')) {
				next()
			} else {
				next('/session')
			}
		  }},
		{path: '/session', component: LoginView},
		{path: '/users/:id/profile', component: HomeView, beforeEnter: (to, from, next) => {
			if ((localStorage.getItem('isLoggedIn') === 'true')) {
				next()
			} else {
				next('/session')
			}
		  }},
		{ path: '/search', component: SearchView, beforeEnter: (to, from, next) => {
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
