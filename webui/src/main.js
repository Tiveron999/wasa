import {createApp, reactive} from 'vue'
import App from './App.vue'
import router from './router'
import axios from './services/axios.js';
import ErrorMsg from './components/ErrorMsg.vue'
import LoadingSpinner from './components/LoadingSpinner.vue'
import UsernameModal from './components/UsernameModal.vue'
import CommentModal from './components/CommentModal.vue';
import LikesModal from './components/LikesModal.vue';
import FollowersModal from './components/FollowersModal.vue';

import './assets/dashboard.css'
import './assets/main.css'

const app = createApp(App)
app.config.globalProperties.$axios = axios; // { ...axios}

app.component("ErrorMsg", ErrorMsg);
app.component("LoadingSpinner", LoadingSpinner);
app.component("UsernameModal", UsernameModal);
app.component("CommentModal", CommentModal);
app.component("LikesModal", LikesModal);
app.component("FollowersModal", FollowersModal);

app.use(router)
app.mount('#app')
