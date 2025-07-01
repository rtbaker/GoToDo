<script setup>
import { ref, watch } from 'vue'
import ToDo from '../model/ToDo.js'

const createAPI=import.meta.env.VITE_API_URL + "/api/1.0/todos"
const errorMessage = ref("")
const title = ref("")
const description = ref("")
const priority = ref(1)

const emit = defineEmits(['closeCreate', 'createNeedsLogin', 'todoCreated'])

const props = defineProps({
  showCreate: Boolean
})

function doCreate() {
    if (title.value.length == 0) {
        errorMessage.value = "Please enter a title."
        return;
    }

    if (description.value.length == 0) {
        errorMessage.value = "Please enter a description."
        return;
    }

    postCreateAPI();
}

// Actually call the API, done with await as there is no point returning
// to the user until an answer
async function postCreateAPI() {
  try {
    const response = await fetch(createAPI, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title: title.value, description: description.value, priority: Number(priority.value) }),
      credentials: "include",
    });

    if (!response.ok) {
      console.log(`Response status: ${response.status}`);

      if (response.status === 401) {
        emit("createNeedsLogin");
      }

      const json = await response.json();
      console.log(json);
      errorMessage.value = json.message

      return;
    }

    const json = await response.json();
          
    let newTodo = ToDo.newFromJSON(json);
    emit("todoCreated", newTodo)

    // all good then close
    emit('closeCreate');
  } catch (error) {
    errorMessage.value = error.message;
  }

}

// Reset values
watch(
  () => props.showCreate,
  () => {
    title.value = "";
    description.value = "";
    priority.value = 1;
  }
);

</script>

<template>
  <Transition name="create-modal">
    <div v-if="showCreate" class="create-modal-mask">
      <div class="create-modal-container">
        <div class="create-modal-header">
          <h3>Create Todo</h3>
        </div>

        <div class="create-modal-body">
          <label for="title">Title</label>
          <input v-model="title" type="text" placeholder="Enter Title" name="title" required>
          
          <label for="description">Description</label>
          <textarea v-model="description" rows="10" placeholder="Enter Description" name="description" required></textarea>
           
          <label for="rangeSlider" class="sliderValue">Priority - {{ priority }}</label>
           <input type="range" class="form-range" v-model="priority" min="0" max="10" step="1"/>
        
          <div class="error">
            {{  errorMessage }}
          </div>
        </div> 

        <div class="create-modal-footer">
          <button
                       type="button"
                          class="create-std-button"
                          @click="doCreate"
                      >OK</button>
            <button
                type="button"
              class="create-std-button create-cancel-btn"
              @click="$emit('closeCreate')"
            >Cancel</button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style>
textarea {
  width: 100%;
}

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

input[type=range] {
  width: 50%;
  display: block;
}

/* Set a style for all buttons */
.create-std-button {
  background-color: #04AA6D;
  color: white;
  padding: 14px 20px;
  margin: 8px 0;
  border: none;
  cursor: pointer;
}

label {
    font-weight: bold;
}

/* Add a hover effect for buttons */
button:hover {
  opacity: 0.8;
}

/* Extra style for the cancel button (red) */
.create-cancel-btn {
  padding: 14px 20px;
  background-color: #f44336;
  margin-left: 10px;
}

/* Change styles for span and cancel button on extra small screens */
@media screen and (max-width: 300px) {
  span.psw {
    display: block;
    float: none;
  }

  .create-std-button {
    width: 100%;
  }
}

.error {
    min-height: 25px;
}
/* stuff from vue tutorial on modals: */

.create-modal-footer {
    border-top: 1px solid darkgrey;
    background-color:#f1f1f1;
    padding: 10px;
}

.create-modal-mask {
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

.create-modal-container {
  width: 40%;
  margin: auto;
  padding: 0px 0px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
}

@media screen and (max-width: 400px) {
  .create-modal-container {
    width: 95%;
  }
}

.create-modal-header h3 {
  text-align: center;
  margin-top: 10px;
  color: #42b983;
}

.create-modal-body {
  margin-top: 10px;
  padding: 16px;
  height: 100%;
}

/*
 * The following styles are auto-applied to elements with
 * transition="modal" when their visibility is toggled
 * by Vue.js.
 *
 * You can easily play with the modal transition by editingcreate-
 * these styles.
 */

.create-modal-enter-from {
  opacity: 0;
}

.create-modal-leave-to {
  opacity: 0;
}

.create-modal-enter-from .create-modal-container,
.create-modal-leave-to .create-modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>