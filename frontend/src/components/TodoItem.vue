<script setup>
import { computed } from 'vue' 
import ToDo from '../model/ToDo.js'

//const props = defineProps({
defineProps({
    todo: {
        type: ToDo,
        required: true
    },
})

const itemStatus = computed(() => {
  return 'normal'
})

</script>

<template>
    <div class="item" :class="itemStatus">
        <div class="content">
            <h2>{{  todo.title }}</h2>
            <p>{{ todo.description }}</p>
        </div>

        <div class="footer">
            <div class="info">
               {{  todo.displayDate() }}
            </div>
            <div class="actionlist">
                <img @click="$emit('markTodoDone', todo.id)" alt="Add Todo" src="../assets/done.svg" class="actionlogo" />
                <img @click="$emit('editTodo', todo.id)" src="../assets/edit.svg" class="actionlogo" />
                <img @click="$emit('deleteTodo', todo.id)" src="../assets/delete.svg" class="actionlogo" />
            </div>
        </div>
    </div>
</template>

<style scoped>
.item {
    border-radius: 8px;
    border-color: black;
    border-style: solid;
    border-width: 1px;
    width: 100%;
    margin: 5px;
    padding: 5px;
    display: flex;
    flex-direction: column;
    background-color: white;
}

.normal {
    background-color: white;
}

.warning {
    background-color: orange;
}

.overdue {
    background-color: red;
}

.footer {
    margin-top: auto;
}

.content {
    margin-bottom: 15px;
}

.actionlist {
    float: right;
    height: auto;
    line-height: 1;
}

.info {
    display: inline-block;
    font-size: 80%;
}
.actionlogo {
    height: 25px;
    width: 25px;
}

img {
    margin-left: 5px;
    border-radius: 12.5px;
}

@media (min-width: 600px) {
    .item {
        width: 250px;
    }
}

@media (hover: hover) {
  img:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}
</style>