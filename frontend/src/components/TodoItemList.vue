<script setup>
import { ref } from 'vue'
import TodoItem from './TodoItem.vue'
import { onMounted } from 'vue'
import ToDo from '../model/ToDo.js'
import CreateTodoModal from './CreateTodoModal.vue'
import MessageModal from './MessageModal.vue'
import * as TodoAPI from '../api/todo.js'

const emit = defineEmits(['needsLogin'])
const showCreateModal = ref(false);
const showMessageModal = ref(false)
const messageTitle = ref("title")
const messageDescr = ref("description")

// urls
const getTodosAPI=import.meta.env.VITE_API_URL + "/api/1.0/todos"

onMounted(() => {
  getTodos();
  console.log(`the TodoItemList component is now mounted.`)
})

const todos = ref(new Map())

function markTodoDone(todo) {
    TodoAPI.markTodoDone(
            todo,
            (data) => { todos.value.set(todo.id, ToDo.newFromJSON(data))}, 
            function (message) {
                messageTitle.value = "Error marking as done";
                messageDescr.value = message;
                showMessageModal.value = true;
            },
            () => emit('needsLogin') 
        );
}

function editTodo(todo) {
    alert(`edit: ${todo.id}`)
}

function deleteTodo(todo) {
    TodoAPI.deleteTodo(
        todo,
        () => todos.value.delete(todo.id),
        function (message) {
            messageTitle.value = "Error deleting";
            messageDescr.value = message;
            showMessageModal.value = true;
        },
        () => emit('needsLogin') 
    );
}

function todoCreated(todo) {
    console.log("New todo added");
    todos.value.set(todo.id, todo);
}

async function getTodos() {
    try {
        const response = await fetch(getTodosAPI, {
          method: "GET",
          credentials: "include",
        });
    
        if (!response.ok) {
          console.log(`Response status: ${response.status}`);

          if (response.status === 401) {
            emit("needsLogin");
          }

          const json = await response.json();
          console.log(json);

          return;
        }
    
        // all good then close
        const json = await response.json();
        loadItems(json);
        //console.log(json);
      } catch (error) {
        console.log(error.message);
      }
}

function sortedItems() {
    let items = Array.from(todos.value.values());
    items.sort((a,b) => b.priority - a.priority);
    return items.sort((a,b) => a.completed === b.completed ? 0 : a.completed ? 1 : -1)
}

function loadItems(itemList) {
    for (const item of itemList) {
        let todo = ToDo.newFromJSON(item);
        todos.value.set(todo.id, todo);
    }
}

function createNeedsLogin() {
    console.log("TodoItemList needs login");
    showCreateModal.value = false;
    emit('needsLogin')
}

</script>


<template>
    <MessageModal @close="showMessageModal = false" :show="showMessageModal" :title="messageTitle" :description="messageDescr"></MessageModal>
    <CreateTodoModal @createNeedsLogin="createNeedsLogin" :showCreate="showCreateModal" @closeCreate="showCreateModal = false" @todoCreated="todoCreated"></CreateTodoModal>

    <div class="mainlist">
        <div class="todoItems">
            <TodoItem v-for="todo in sortedItems()" :todo="todo" v-bind:key="todo.id"
                @markTodoDone="markTodoDone"
                @editTodo="editTodo"
                @deleteTodo="deleteTodo"
                />
        </div>
    </div>

    <div>
        <img @click="showCreateModal = true" alt="Add Todo" class="logo" src="../assets/add.svg" width="35" height="35" />
        <br/>Order: higher priority first, completed tasks last.
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