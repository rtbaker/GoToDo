<script setup>
import TodoItemList from './components/TodoItemList.vue'
import { ref } from 'vue'
import LoginModal from './components/LoginModal.vue'

const showLoginModal = ref(false);

function needsLogin() {
  showLoginModal.value = true;
}

</script>

<template>
      <LoginModal :show="showLoginModal" @close="showLoginModal = false"></LoginModal>
      
      <TodoItemList @needsLogin="needsLogin"/>

      <div>
        <a href=""><img alt="Add Todo" class="logo" src="./assets/add.svg" width="35" height="35" /></a>
      </div>
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
