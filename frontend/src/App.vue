<script setup>
import TodoItemList from './components/TodoItemList.vue'
import { ref } from 'vue'
import LoginModal from './components/LoginModal.vue'

const showLoginModal = ref(false);
const reloadItemList = ref(0);

function needsLogin() {
  showLoginModal.value = true;
}

function loginSuccess() {
  showLoginModal.value = false;
  reloadItemList.value++;

  console.log("login success");
}

</script>

<template>
      <LoginModal :show="showLoginModal" @closeLogin="showLoginModal = false" @loginSuccess="loginSuccess"></LoginModal>

      <TodoItemList @needsLogin="needsLogin" :key="reloadItemList"/>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
