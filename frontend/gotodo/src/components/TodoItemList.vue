<script setup>
import { ref } from 'vue'
import TodoItem from './TodoItem.vue'

const todos = ref([
    { 'id': 1, 'title': "title 1", 'description': "a description for our todo number 1"},
    { 'id': 2, 'title': "title 2", 'description': "a description 2"},
    { 'id': 3, 'title': "title 3", 'description': "a description 3"},
])
</script>


<template>
    <div class="mainlist">
        <todoItems>
            <TodoItem v-for="todo in todos" :key="todo.id" :data="todo" />
        </todoItems>
    </div>
</template>

<style scoped>

.mainlist {
    margin: 10px;
    border-radius: 5px;
    width: 100%;
    height: auto;
}

todoItems {
  display: flex;
  flex-wrap: wrap;
}

</style>