<script>
import { mapActions } from "vuex";
import BaseCard from "../components/UI/BaseCard.vue";
export default {
  components: { BaseCard },
  data() {
	return {
	  loginSelected: true,
	  username: "",
	  password: "",
	  error: ""
	};
  },
  methods: {
	...mapActions({ login: "auth/login", signup: "auth/signup" }),
	loginButtonClicked() {
	  if (this.loginSelected) {
		this.login({ username: this.username, password: this.password })
		  .then(() => {
			this.$router.push({ name: "Posts" });
		  })
		  .catch(error => {
			this.error = error;
		  });
	  } else {
		this.signup({ username: this.username, password: this.password })
		  .then(() => {
			this.$router.push({ name: "Posts" });
		  })
		  .catch(error => {
			this.error = error;
		  });
	  }
	},
	toggleLogin() {
	  this.loginSelected = !this.loginSelected;
	}
  },
  computed: {
	buttonString() {
	  if (this.loginSelected) {
		return "LOGIN";
	  } else {
		return "SIGNUP";
	  }
	},
	textLoginString() {
	  if (this.loginSelected) {
		return "Signup instead";
	  } else {
		return "Login instead";
	  }
	}
  }
};
</script>

