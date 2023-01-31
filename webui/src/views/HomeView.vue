<script>

export default {
    /*components: {
        LoginView,
        HomeView,
    },*/
    data: function () {
        return {
            errormsg: null,
            loading: false,
            users: null,
            isLoggedIn: false,
        };
    },
    methods: {
        async refresh() {
            this.loading = true;
            this.errormsg = null;
            try {
                let response = await this.$axios.get("/context"); // .then(response => {console.log(response)});
                this.users = response.data;
            }
            catch (e) {
                this.errormsg = e.toString();
            }
            this.loading = false;
        },
        async checkLoginStatus() {
            try {
                const response = await this.$axios.post("/session", {
                    username: this.username,
                });
                if (response.data.userID) {
                    this.isLoggedIn = true;
                }
                else {
                    this.isLoggedIn = false;
                }
            }
            catch (error) {
                console.error(error);
                this.isLoggedIn = false;
            }
        },
    },
    mounted() {
        this.refresh();
        /*this.checkLoginStatus();*/
    },
    created() {
        /*this.checkLoginStatus();*/
    },
}
</script>

<template>
	
	<div> 
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WasaPhoto</h1>
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
			</div>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ul>
			<li> hi </li>
		</ul>
	</div>
</template>

<style>
</style>
