<script>

export default {
    data() {
        return {
            followers: [],
        };
    },

  props: {
    showFollowersModal: Boolean,
    userID: Number,
    sameUser: Boolean,
  },

  methods: {
    async listFollowers() {
        
            try {
              const response = await this.$axios.get(`/users/${this.userID}/followers`, {headers: {Authorization: 'Bearer ' + this.id }}); 
              this.followers = response.data;
              this.checkIsFollowing(this.followers)
            } catch (error) {
              console.error('Error fetching followers:', error);
            }
          },

    checkIsFollowing(followers) {
        if (!this.sameUser) {
            // check if the user is already following so to update the button
            const isFollowing = followers.some(follower => follower.id === this.id);
            console.log(isFollowing)
            this.$emit('following-check', isFollowing);
        }
    },

    closeModal() {
        this.$emit('update:showFollowersModal', false);
    },
    
  },

  created() {
    this.listFollowers();
  },
};

</script>


<template>
    <div class="followers-modal" v-if="showFollowersModal">
      <div class="followers-modal-dialog">
        <div class="followers-modal-content">
          <div class="followers-modal-header">
            <h5 class="followers-list-modal-title">Followers</h5>
            <button type="button" class="followers-list-modal-close" @click="closeModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="followers-list-modal-body">
            <ul>
              <li v-for="(follower, index) in followers" :key="index" class="followers-list-item">
                {{ follower }}
              </li>
            </ul>
          </div>
          <div class="followers-list-modal-footer">
            <button type="button" class="followers-list-modal-btn followers.list-modal-close" @click="closeModal">Close</button>
          </div>
        </div>
      </div>
    </div>
  </template>




<style>
</style>