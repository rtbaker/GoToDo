<script setup>

defineProps({
  show: Boolean,
  title: String,
  description: String,
})

</script>

<template>
  <Transition name="modal">
    <div v-if="show" class="modal-mask">
      <div class="modal-container">
        <div class="modal-header">
          <h3>{{ title }}</h3>
        </div>

        <div class="modal-body">
              {{  description }}
        </div>

        <div class="modal-footer">
            <button
                type="button"
              class="modal-default-button"
              @click="$emit('close')"
            >Close</button>
        </div>
      </div>
    </div>
  </Transition>
</template>

<style>

/* Set a style for all buttons */
.modal-default-button {
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


/* Change styles for span and cancel button on extra small screens */
@media screen and (max-width: 300px) {
  span.psw {
    display: block;
    float: none;
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
  text-align: center;
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