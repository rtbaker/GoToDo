<script setup>
import { ref } from 'vue'
import TodoItem from './TodoItem.vue'
import { onMounted } from 'vue'
import ToDo from '../model/ToDo.js'

const emit = defineEmits(['needsLogin'])


onMounted(() => {
    emit('needsLogin')
  console.log(`the TodoItemList component is now mounted.`)
})

const todos = ref([
    ToDo.newFromJSON({ 'id': 1, 'title': "title 1", 'description': "a description for our todo number 1!", 'updatedAt': "2025-06-24T14:04:32.689170511+01:00"}),
    ToDo.newFromJSON({ 'id': 2, 'title': "title 2", 'description': "a description 2", 'updatedAt': "2025-06-24T14:04:32.689170511+01:00"}),
    ToDo.newFromJSON({ 'id': 3, 'title': "title 3", 'description': "a description 3", 'updatedAt': "2025-06-24T14:04:32.689170511+01:00"}),
])

function markTodoDone(todoId) {
    alert(`mark done: ${todoId}`)
}

function editTodo(todoId) {
    alert(`edit: ${todoId}`)
}

function deleteTodo(todoId) {
    alert(`delete: ${todoId}`)
}

</script>


<template>
    <div class="mainlist">
        <div class="todoItems">
            <TodoItem v-for="todo in todos" :todo="todo" v-bind:key="todo.id"
                @markTodoDone="markTodoDone"
                @editTodo="editTodo"
                @deleteTodo="deleteTodo"
                />
        </div>
    </div>
</template>

<style scoped>

.mainlist {
    margin: 10px;
    border-radius: 5px;
    width: 100%;
    height: auto;
}

.todoItems {
  display: flex;
  flex-wrap: wrap;
}

</style>