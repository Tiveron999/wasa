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
  },

  methods: {
    async listFollowers() {
            try {
              const response = await this.$axios.get(`/users/${this.posterID}/followers`, {headers: {Authorization: 'Bearer ' + this.id }}); 
              this.followers = response.data;
            } catch (error) {
              console.error('Error fetching followers:', error);
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
            <h5 class="followers-modal-title">Followers</h5>
            <button type="button" class="followers-modal-close" @click="closeModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="followers-modal-body">
            <ul>
              <li v-for="(follower, index) in followers" :key="index" class="followers-list-item">
                {{ follower.username }}
              </li>
            </ul>
          </div>
          <div class="followers-modal-footer">
            <button type="button" class="followers-modal-btn followers-modal-close" @click="closeModal">Close</button>
          </div>
        </div>
      </div>
    </div>
  </template>




<style>
</style>