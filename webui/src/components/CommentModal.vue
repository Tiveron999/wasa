<script>


export default {
    data() {
        return {
            newComment: '',
            comments: [],
            newCommentsNumber: 0,
        };
    },
    props: {
        showCommentModal: Boolean,
        posterID: Number,
        pid: Number,


  },
    methods: {
        async submitComment() {
          if (localStorage.getItem('isLoggedIn')) {
            this.id = localStorage.getItem('userId');
          }
          try {
            const commentText = this.newComment;
            console.log(commentText)
            
            const response = await this.$axios.post(`/users/${this.posterID}/photos/${this.pid}/comments`, JSON.stringify({comment: commentText}),
              {headers: {'Content-Type': 'application/json', Authorization: 'Bearer ' + this.id }});

            
              this.newCommentsNumber = response.data.comments
              this.updateCommentsNumber(this.newCommentsNumber)
    
            // Clear the input field after submission
            this.newComment = '';

            this.listComments()

            } catch (e) {
            console.error("error: ", e);
            this.errormsg = e.toString();
          }

          },

          async listComments() {
            if (localStorage.getItem('isLoggedIn')) {
                this.id = localStorage.getItem('userId');
              }
            try {
              const response = await this.$axios.get(`/users/${this.posterID}/photos/${this.pid}/comments`, {headers: {Authorization: 'Bearer ' + this.id }}); 
              this.comments = response.data;
            } catch (error) {
              console.error('Error fetching comments:', error);
            }
          },

          isCommenter(commenterID) {
            const storedUserId = localStorage.getItem('userId');
            return commenterID === storedUserId;
        },

        deleteComment(cid) {
          // Assuming you have a way to identify comments by index
          const commentToDelete = this.comments[index];
          
          // Make an API request or perform any action to delete the comment
          // For example, you can use Axios for the request
          // Replace 'yourApiEndpoint' with your actual API endpoint
          axios.delete('yourApiEndpoint/' + commentToDelete.CID)
            .then(() => {
              // Remove the comment from the comments array
              this.comments.splice(index, 1);
            })
            .catch((error) => {
              console.error('Error deleting comment:', error);
            });
        },


        closeModal() {
          this.$emit('update:showCommentModal', false);
        },
        updateCommentsNumber(newComments) {
          this.$emit('update-comments', newComments);
        },
    },

    created() {
    this.listComments();
  },
};
</script>



<template>

 <div class="comment-modal" v-if="showCommentModal">
    <div class="comment-modal-dialog">
      <div class="comment-modal-content">
        <div class="comment-modal-header">
          <h5 class="comment-modal-title">Comments</h5>
          <button type="button" class="comment-modal-close" @click="closeModal">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="comment-modal-body">
          <!-- comments -->
          <div v-for="(comment, index) in comments" :key="index" class="comment">
            <p class="comment-author">{{ comment.commenterid }} - {{ comment.datetime }}</p>
            <p class="comment-text">{{ comment.text }}</p>
            <i class="fas fa-trash-alt" v-if="isCommenter(comment.commenterid)" @click="deleteComment(comment.cid)"></i>
          </div>
          <div class="comment-form-group">
            <label for="commentInput">Add a Comment</label>
            <textarea class="comment-form-control" id="commentInput" rows="3" v-model="newComment"></textarea>
          </div>
          <button class="comment-btn comment-btn-primary" @click="submitComment">Post</button>
        </div>
        <div class="comment-modal-footer">
          <button type="button" class="comment-btn comment-btn-secondary" @click="closeModal">Close</button>
        </div>
      </div>
    </div>
  </div>


</template>




<style>
</style>