<template>
    <div class="modal" tabindex="-1" role="dialog">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title">Change Username</h5>
            <button type="button" class="close" @click="closeUsernameModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <input
              type="text"
              class="form-control"
              placeholder="New Username"
              v-model="newUsername"
              :pattern="usernamePattern"
              title="Username must start with a letter and be between 6 and 15 characters long. Only alphanumeric characters and underscores are allowed."
            />
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-primary" @click="changeUsername">Save</button>
            <button type="button" class="btn btn-secondary" @click="closeUsernameModal">Close</button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        newUsername: "",
        usernamePattern: /^[A-Za-z][A-Za-z0-9_]{5,14}$/,
      };
    },
    methods: {
        closeUsernameModal() {
        // Emit a custom event to notify the parent component to close the modal
        this.$emit("close-modal");
      },
      changeUsername() {
        // Emit a custom event to notify the parent component to change the username
        this.$emit("change-username", this.newUsername);
        // Reset the newUsername field after emitting the event
        this.newUsername = "";
      },
    },
  };
  </script>
  
  <style>
  /* Add any custom styles for the modal here */
  </style>
  