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
 <div class="modal fade my-modal" :id="modal_id" tabindex="-1" aria-hidden="true">
  <div class="modal-dialog modal-dialog-centered modal-dialog-scrollable">
    <div class="modal-content bg-primary text-white">
      <div class="modal-header">
        <h1 class="modal-title fs-5 text-center" :id="modal_id">Comments</h1>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
      </div>

      

      <div class="modal-footer d-flex justify-content-center w-100">
        <div class="row w-100">
          <div class="col-10">
            <div class="mb-3 me-auto">
              <textarea class="form-control bg-dark text-white" id="exampleFormControlTextarea1" 
                placeholder="Add a comment..." rows="1" maxLength="30" v-model="commentValue"></textarea>
            </div>
          </div>

          <div class="col-2 d-flex align-items-center">
            <button type="button" class="btn btn-light" 
              @click.prevent="addComment" 
              :disabled="commentValue.length < 1 || commentValue.length > 30">
              Send
            </button>
          </div>
        </div>
      </div>

    </div>
  </div>
</div>



</template>




<style>
</style>