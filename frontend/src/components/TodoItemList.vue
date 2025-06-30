<script setup>
import { ref } from 'vue'
import TodoItem from './TodoItem.vue'
import { onMounted } from 'vue'
import ToDo from '../model/ToDo.js'

const emit = defineEmits(['needsLogin'])

// urls
const getTodosAPI=import.meta.env.VITE_API_URL + "/api/1.0/todos"

onMounted(() => {
  getTodos();
  console.log(`the TodoItemList component is now mounted.`)
})

const todos = ref(new Map())

function markTodoDone(todoId) {
    alert(`mark done: ${todoId}`)
}

function editTodo(todoId) {
    alert(`edit: ${todoId}`)
}

function deleteTodo(todoId) {
    alert(`delete: ${todoId}`)
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
    console.log(items);
    return items.sort((a,b) => b.priority - a.priority);
}

function loadItems(itemList) {
    for (const item of itemList) {
        let todo = ToDo.newFromJSON(item);
        todos.value.set(todo.id, todo);
    }
}
</script>


<template>
    <div class="mainlist">
        <div class="todoItems">
            <TodoItem v-for="todo in sortedItems()" :todo="todo" v-bind:key="todo.id"
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