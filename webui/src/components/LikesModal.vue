<script>

export default {
    data() {
        return {
            likes: [],
        };
    },

  props: {
    showLikesModal: Boolean,
    posterID: Number,
    pid: Number,

  },

  methods: {
    async listLikes() {
            if (localStorage.getItem('isLoggedIn')) {
                this.id = localStorage.getItem('userId');
              }
            try {
              const response = await this.$axios.get(`/users/${this.posterID}/photos/${this.pid}/likes`, {headers: {Authorization: 'Bearer ' + this.id }}); 
              this.likes = response.data;
              console.log("likes", this.likes)
            } catch (error) {
              console.error('Error fetching likes:', error);
            }
          },


    closeModal() {
        this.$emit('update:showLikesModal', false);
    },
    
  },

  created() {
    this.listLikes();
  },
};

</script>



<template>
    <div class="likes-list-modal" v-if="showLikesModal">
      <div class="likes-list-modal-dialog">
        <div class="likes-list-modal-content">
          <div class="likes-list-modal-header">
            <h5 class="likes-list-modal-title">Likes</h5>
            <button type="button" class="likes-list-modal-close" @click="closeModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="likes-list-modal-body">
            <ul>
              <li v-for="(liker, index) in likes" :key="index" class="likes-list-item">
                {{ liker }}
              </li>
            </ul>
          </div>
          <div class="likes-list-modal-footer">
            <button type="button" class="likes-list-modal-btn likes-list-modal-close" @click="closeModal">Close</button>
          </div>
        </div>
      </div>
    </div>

</template>




<style>
</style>