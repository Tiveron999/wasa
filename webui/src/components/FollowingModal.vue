<script>

export default {
    data() {
        return {
            following: [],
        };
    },

  props: {
    showFollowingModal: Boolean,
    userID: Number,
    followingsnum: Number,
  },

  watch: {
    followingsnum (newvalue) {
      // Call the method when the route changes
      this.listFollowing()
    }
  },

  methods: {
    async listFollowing() {
        
            try {
              const response = await this.$axios.get(`/users/${this.userID}/following`, {headers: {Authorization: 'Bearer ' + this.id }}); 
              this.following = response.data;
            } catch (error) {
              console.error('Error fetching following:', error);
            }
          },


    closeModal() {
        this.$emit('update:showFollowingModal', false);
    },
    
  },

  created() {
    this.listFollowing();
  },
};

</script>


<template>
    <div class="following-modal" v-if="showFollowingModal">
      <div class="following-modal-dialog">
        <div class="following-modal-content">
          <div class="following-modal-header">
            <h5 class="following-list-modal-title">Following</h5>
            <button type="button" class="following-list-modal-close" @click="closeModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="following-list-modal-body">
            <ul>
              <li v-for="(follower, index) in following" :key="index" class="following-list-item">
                {{ follower }}
              </li>
            </ul>
          </div>
          <div class="following-list-modal-footer">
            <button type="button" class="following-list-modal-btn following.list-modal-close" @click="closeModal">Close</button>
          </div>
        </div>
      </div>
    </div>
  </template>




<style>
</style>