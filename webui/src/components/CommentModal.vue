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
<div class="modal" v-if="showCommentModal">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title">Comments</h5>
        <button type="button" class="close" @click="closeModal">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <div class="form-group">
          <label for="commentInput">Add a Comment</label>
          <textarea class="form-control" id="commentInput" rows="3" v-model="newComment"></textarea>
        </div>
        <button class="btn btn-primary" @click="addComment">Submit Comment</button>
        <div v-for="(comment, index) in comments" :key="index">
          <p>{{ comment.author }} - {{ comment.timestamp }}</p>
          <p>{{ comment.text }}</p>
        </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-secondary" @click="closeModal">Close</button>
      </div>
    </div>
  </div>
</div>


</template>




<style>
</style>