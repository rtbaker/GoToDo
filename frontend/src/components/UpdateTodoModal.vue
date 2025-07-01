<script setup>
import { ref, watch, computed } from 'vue'
import ToDo from '../model/ToDo.js'

const createAPI=import.meta.env.VITE_API_URL + "/api/1.0/todos"
const errorMessage = ref("")


const emit = defineEmits(['closeUpdate', 'updateNeedsLogin', 'todoUpdated'])

const props = defineProps({
  showUpdate: Boolean,
  todo: {
        type: ToDo,
        required: true
  },
})

const title = ref(null);
const description = ref(null);
const priority = ref(null);

function doUpdate() {
    if (title.value.length == 0) {
        errorMessage.value = "Please enter a title."
        return;
    }

    if (description.value.length == 0) {
        errorMessage.value = "Please enter a description."
        return;
    }

    postUpdateAPI();
}

// Actually call the API, done with await as there is no point returning
// to the user until an answer
async function postUpdateAPI() {
  try {
    const response = await fetch(createAPI + '/' + props.todo.id, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title: title.value, description: description.value, priority: Number(priority.value) }),
      credentials: "include",
    });

    if (!response.ok) {
      console.log(`Response status: ${response.status}`);

      if (response.status === 401) {
        emit("updateNeedsLogin");
      }

      const json = await response.json();
      console.log(json);
      errorMessage.value = json.message

      return;
    }

    const json = await response.json();
          
    let newTodo = ToDo.newFromJSON(json);
    emit("todoUpdated", newTodo)

    // all good then close
    emit('closeUpdate');
  } catch (error) {
    errorMessage.value = error.message;
  }
}

// Reset values
watch(
  () => props.showUpdate,
  () => {
    title.value = props.todo.title;
    description.value = props.todo.description;
    priority.value = props.todo.priority;
  }
);

</script>

<template>
  <Transition name="update-modal">
    <div v-if="showUpdate" class="update-modal-mask">
      <div class="update-modal-container">
        <div class="update-modal-header">
          <h3>Update Todo</h3>
        </div>

        <div class="update-modal-body">
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

        <div class="update-modal-footer">
          <button
                       type="button"
                          class="update-std-button"
                          @click="doUpdate"
                      >OK</button>
            <button
                type="button"
              class="update-std-button update-cancel-btn"
              @click="$emit('closeUpdate')"
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
.update-std-button {
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
.update-cancel-btn {
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

  .update-std-button {
    width: 100%;
  }
}

.error {
    min-height: 25px;
}
/* stuff from vue tutorial on modals: */

.update-modal-footer {
    border-top: 1px solid darkgrey;
    background-color:#f1f1f1;
    padding: 10px;
}

.update-modal-mask {
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

.update-modal-container {
  width: 40%;
  margin: auto;
  padding: 0px 0px;
  background-color: #fff;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.33);
  transition: all 0.3s ease;
}

@media screen and (max-width: 400px) {
  .update-modal-container {
    width: 95%;
  }
}

.update-modal-header h3 {
  text-align: center;
  margin-top: 10px;
  color: #42b983;
}

.update-modal-body {
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

.update-modal-enter-from {
  opacity: 0;
}

.update-modal-leave-to {
  opacity: 0;
}

.update-modal-enter-from .update-modal-container,
.update-modal-leave-to .update-modal-container {
  -webkit-transform: scale(1.1);
  transform: scale(1.1);
}
</style>