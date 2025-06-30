<script setup>
import { ref } from 'vue'
import eyeUrl from '@/assets/eye.svg';
import eyeCrossedUrl from '@/assets/eye-crossed.svg';

const loginAPI=import.meta.env.VITE_API_URL + "/api/1.0/login"
const errorMessage = ref("")
const email = ref("")
const password = ref("")
const passwordFieldType = ref("password")
const passwordEyeImage = ref(eyeUrl)

const emit = defineEmits(['closeLogin', 'loginSuccess'])

defineProps({
  show: Boolean
})

function doLogin() {
    if (email.value.length == 0) {
        errorMessage.value = "Please enter an email address."
        return;
    }
    
    if (!emailIsValid(email.value)) {
        errorMessage.value = "Sorry this does not look like a valid email address."
        return;
    }

    if (password.value.length == 0) {
        errorMessage.value = "Please enter a password."
        return;
    }

    postLoginAPI();
}

// Actually call the API, done with await as there is no point returning
// to the user until an answer
async function postLoginAPI() {
  try {
    const response = await fetch(loginAPI, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email: email.value, password: password.value }),
      credentials: "include",
    });

    if (!response.ok) {
      console.log(`Response status: ${response.status}`);

      const json = await response.json();
      console.log(json);
      errorMessage.value = json.message
      return;
    }

    // all good then close
    emit('loginSuccess');
  } catch (error) {
    errorMessage.value = error.message;
  }
}

// Basic check
function emailIsValid (email) {
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)
}

// Password field hide/show text
function switchPassword() {
    passwordFieldType.value = (passwordFieldType.value === "password" ? "text" : "password")
    passwordEyeImage.value = (passwordEyeImage.value === eyeUrl ? eyeCrossedUrl : eyeUrl)
}

</script>

<template>
  <Transition name="modal">
    <div v-if="show" class="modal-mask">
      <div class="modal-container">
        <div class="modal-header">
          <h3>Login</h3>
        </div>

        <div class="modal-body">
          <label for="uname">Email</label>
            <input v-model="email" type="text" placeholder="Enter Username" name="uname" required>
        
            <div class="password-label">
                <label for="psw">Password</label>
                <img class="password-eye" :src="passwordEyeImage" @click="switchPassword"></img>
            </div>
            <input v-model="password" :type="passwordFieldType" placeholder="Enter Password" name="psw" required>

            <button
             type="button"
                class="modal-default-button"
                @click="doLogin"
            >OK</button>

            <div class="error">
              {{  errorMessage }}
            </div>
        </div>

        <div class="modal-footer">
            <button
                type="button"
              class="modal-default-button cancel-btn"
              @click="$emit('close')"
            >Cancel</button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style>
.password-label {
    display: flex;
    align-items: center;
}

.password-eye {
    width: 20px;
    height: 20px;
    margin-left: 5px;
}
/* stuff nicked from w3schools Login Form example (https://www.w3schools.com/howto/howto_css_login_form.asp) */

/* Full-width inputs */
input[type=text], input[type=password] {
  width: 100%;
  padding: 12px 20px;
  margin: 8px 0;
  display: inline-block;
  border: 1px solid #ccc;
  box-sizing: border-box;
}

/* Set a style for all buttons */
button {
  background-color: #04AA6D;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  cursor: pointer;
  width: 100%;
  display: block;
}

label {
    font-weight: bold;
}

/* Add a hover effect for buttons */
button:hover {
  opacity: 0.8;
}

/* Extra style for the cancel button (red) */
.cancel-btn {
  width: auto;
  padding: 10px 18px;
  background-color: #f44336;
}

/* Change styles for span and cancel button on extra small screens */
@media screen and (max-width: 300px) {
  span.psw {
    display: block;
    float: none;
  }
  .cancelbtn {
    width: 100%;
  }
}

.error {
    min-height: 25px;
}
/* stuff from vue tutorial on modals: */

.modal-footer {
    border-top: 1px solid darkgrey;
    background-color:#f1f1f1;
    padding: 10px;
}

.modal-mask {
  position: fixed;
  z-index: 9998;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  transition: opacity 0.3s ease;
}

.modal-container {
  width: 40%;
  margin: auto;
  padding: 0px 0px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
}

@media screen and (max-width: 400px) {
  .modal-container {
    width: 95%;
  }
}

.modal-header h3 {
  text-align: center;
  margin-top: 10px;
  color: #42b983;
}

.modal-body {
  margin-top: 10px;
  padding: 16px;
  height: 100%;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editing
 * these styles.
 */

.modal-enter-from {
  opacity: 0;
}

.modal-leave-to {
  opacity: 0;
}

.modal-enter-from .modal-container,
.modal-leave-to .modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>