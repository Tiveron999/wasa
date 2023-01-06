import {createRouter, createWebHashHistory} from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'


const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{path: '/', component: HomeView, 
		beforeEnter: (to, from, next) => {
			if (this.$isLoggedIn()) {
			  next();
			} else {
			  next('/login');
			}
		  },},
		{path: '/link1', component: HomeView, beforeEnter: (to, from, next) => {
			if (this.$isLoggedIn()) {
			  next();
			} else {
			  next('/login');
			}
		  },},
		{path: '/link2', component: HomeView, 
		/*beforeEnter: (to, from, next) => {
			if (isLoggedIn()) {
			  next();
			} else {
			  next('/login');
			}
		  },*/
		},
		{path: '/some/:id/link', component: HomeView, beforeEnter: (to, from, next) => {
			if (this.$isLoggedIn()) {
			  next();
			} else {
			  next('/login');
			}
		  },},
		{path: '/login', component: LoginView, 
		beforeEnter: (to, from, next) => {
			if (!this.$isLoggedIn()) {
			  next();
			} else {
			  next('/');
			}
		  },
		},
	]
})

export default router
