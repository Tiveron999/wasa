<script>
export default {
    data() {
        return {
            newComment: '',
            comments: [],
        };
    },
    props: {
        showCommentModal: Boolean,
        

  },
  mounted() {
    console.log('CommentModal mounted');
  },
  updated() {
    console.log('CommentModal updated');
  },
  
    methods: {
        addComment() {
            // Logic to add a new comment
            const comment = {
                author: 'Your Name',
                timestamp: new Date().toLocaleString(),
                text: this.newComment,
            };
            this.comments.push(comment);
            this.newComment = '';
        },
        closeModal() {
      this.$emit('update:showCommentModal', false);
        },
    },
};
</script>



<template>
 <!-- Comment Modal -->
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
          <!-- Add your comment input and display logic here -->
          <div class="comment-form-group">
            <label for="commentInput">Add a Comment</label>
            <textarea class="comment-form-control" id="commentInput" rows="3" v-model="newComment"></textarea>
          </div>
          <button class="comment-btn comment-btn-primary" @click="addComment">Submit Comment</button>
          <!-- Display existing comments here -->
          <div v-for="(comment, index) in comments" :key="index" class="comment">
            <p class="comment-author">{{ comment.author }} - {{ comment.timestamp }}</p>
            <p class="comment-text">{{ comment.text }}</p>
          </div>
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