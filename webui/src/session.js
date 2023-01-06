import Vue from 'vue';
import jwt from 'jsonwebtoken';

const isLoggedIn = () => {
  // Check for the presence of a JWT in local storage
  const token = localStorage.getItem('token');
  if (token) {
    // JWT is present, verify it using the secret key
    try {
      jwt.verify(token, secretKey);
      return true;
    } catch (error) {
      // JWT is invalid, return false
      return false;
    }
  } else {
    // JWT is not present, user is not logged in
    return false;
  }
};

Vue.prototype.$isLoggedIn = isLoggedIn;