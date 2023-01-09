import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView, 
		/*beforeEnter: (to, from, next) => {
			if (!isLoggedIn()) {
				next('/session')
			  } else {
				next()
			  }
		  },*/},
		  {path: '/profile', component: ProfileView, 
		  beforeEnter: (to, from, next) => {
			if (!isLoggedIn()) {
				next('/session')
			  } else {
				next()
			  }
			},},
		{path: '/link1', component: HomeView, beforeEnter: (to, from, next) => {
			if (!isLoggedIn()) {
				next('/session')
			  } else {
				next()
			  }
		  },},
		{path: '/link2', component: HomeView, 
		/*beforeEnter: (to, from, next) => {
			if (isLoggedIn()) {
			  next();
			} else {
			  next('/session');
			}
		  },*/
		},
		{path: '/some/:id/link', component: HomeView, beforeEnter: (to, from, next) => {
			if (!isLoggedIn()) {
				next('/session')
			  } else {
				next()
			  }
		  },},
		{path: '/session', component: LoginView},
		{path: '/users/:id/profile', component: HomeView, beforeEnter: (to, from, next) => {
			if (!isLoggedIn()) {
				next('/session')
			  } else {
				next()
			  }
		  },},
		
	]
})

export default router
