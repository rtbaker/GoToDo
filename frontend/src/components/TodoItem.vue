<script setup>
import { computed } from 'vue' 

const props = defineProps({
  data: {
    type: Object,
    required: true,
  },
})

const itemStatus = computed(() => {
  return 'normal'
})

function markDone() {
    alert(`marked done: ${props.data.id}`)
}

function editTodo() {
    alert(`edit todo: ${props.data.id}`)
}

function deleteTodo() {
    alert(`delete todo: ${props.data.id}`)
}

</script>

<template>
    <div class="item" :class="itemStatus">
        <div class="content">
            <h2>{{  data.title }}</h2>
            <p>{{ data.description }}</p>
        </div>

        <div class="footer">
            <div class="info">
                date here
            </div>
            <div class="actionlist">
                <img @click="markDone" alt="Add Todo" src="../assets/done.svg" class="actionlogo" />
                <img @click="editTodo" alt="Edit Todo" src="../assets/edit.svg" class="actionlogo" />
                <img @click="deleteTodo" alt="Delete Todo" src="../assets/delete.svg" class="actionlogo" />
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