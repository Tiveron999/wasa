import {createApp, reactive} from 'vue'
// import Vue from 'vue'
import Vuex from 'vuex'
// import store from './store'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios;
app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.use(router)
app.mount('#app')

Vue.use(Vuex);

new Vue({
    el: '#app',
    store,
    render: h => h(App)
});
