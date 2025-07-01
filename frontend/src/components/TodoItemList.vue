<script setup>
import { ref } from 'vue'
import TodoItem from './TodoItem.vue'
import { onMounted } from 'vue'
import ToDo from '../model/ToDo.js'
import CreateTodoModal from './CreateTodoModal.vue'
import UpdateTodoModal from './UpdateTodoModal.vue'
import MessageModal from './MessageModal.vue'
import * as TodoAPI from '../api/todo.js'

const emit = defineEmits(['needsLogin'])
const showCreateModal = ref(false);
const showUpdateModal = ref(false);
const showMessageModal = ref(false)
const messageTitle = ref("title")
const messageDescr = ref("description")
const currentTodo = ref(new ToDo(1,"", "", "", "", false)) // Dummy current todo to start (used by update modal)

// urls
const getTodosAPI=import.meta.env.VITE_API_URL + "/api/1.0/todos"

onMounted(() => {
  getTodos();
  // console.log(`the TodoItemList component is now mounted.`)
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
    currentTodo.value = todo;
    showUpdateModal.value = true;
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

function todoUpdated(todo) {
    console.log("Todo updated");
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
    console.log("CreateTodoModal needs login");
    showCreateModal.value = false;
    emit('needsLogin')
}

function updateNeedsLogin() {
    console.log("UpdateTodoModal needs login");
    showUpdateModal.value = false;
    emit('needsLogin')
}

</script>


<template>
    <MessageModal @close="showMessageModal = false" :show="showMessageModal" :title="messageTitle" :description="messageDescr"></MessageModal>
    <CreateTodoModal @createNeedsLogin="createNeedsLogin" :showCreate="showCreateModal" @closeCreate="showCreateModal = false" @todoCreated="todoCreated"></CreateTodoModal>
    <UpdateTodoModal :todo="currentTodo" @updateNeedsLogin="updateNeedsLogin" :showUpdate="showUpdateModal" @closeUpdate="showUpdateModal = false" @todoUpdated="todoUpdated"></UpdateTodoModal>

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
        <br/>
        <span class="indicator less-than-week less-than-week-bg">Less than a week old</span>
        <span class="indicator one-weeks-old one-weeks-old-bg">More than one week old</span>
        <span class="indicator two-weeks-old two-weeks-old-bg">More than two weeks old</span>
        <span class="indicator done-indicator done-bg">DONE</span>
    </div>
</template>

<style scoped>

.indicator {
    border-width: 1px;
    border-style: solid;
    border-right-style: hidden;
    border-color: black;
    margin-top: 2px;
    padding: 2px;
}
.less-than-week {
    background-color: white;
}

.done-indicator {
    border-right-style: solid;
}
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