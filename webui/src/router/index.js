import {createRouter, createWebHashHistory} from 'vue-router'
import Vuex from 'vuex'
// import store from './store'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '../views/ProfileView.vue'
import SearchView from '../views/SearchView.vue'


const store = new Vuex.Store({
	state: {
	  isLoggedIn: false,
	  username: ''
	},
	actions: {
		login({ commit }, userId) {
			commit('setIsLoggedIn', true)
			commit('setUserId', userId)
		}
	},	
	mutations: {
	  setIsLoggedIn (state, isLoggedIn) {
		state.isLoggedIn = isLoggedIn
	  },
	  setUsername (state, username) {
		state.username = username
	  },
	  setUserId(state, userId) {
        state.userId = userId
    }
	}
  })

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		{ path: '/', component: HomeView, beforeEnter: (to, from, next) => {
			if (store.state.isLoggedIn) {
			  next()
			} else {
			  next('/session')
			}
		  }
		},
		  {path: '/profile', component: ProfileView, 
		  beforeEnter: (to, from, next) => {
			if (!isLoggedIn()) {
				next('/session')
			  } else {
				next()
			  }
			},},
		{path: '/link1', component: HomeView},
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
			if (store.state.isLoggedIn) {
				next()
			} else {
				next('/')
			}
		  },},
		{path: '/session', component: LoginView},
		{path: '/users/:id/profile', component: HomeView, beforeEnter: (to, from, next) => {
			if (store.state.isLoggedIn) {
				next()
			} else {
				next('/profile')
			}
		  },},
		{ path: '/search', component: SearchView, beforeEnter: (to, from, next) => {
		if (store.state.isLoggedIn) {
			next()
		} else {
			next('/session')
		}
		}
	},
		
	]
})

export default router
