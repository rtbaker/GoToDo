<script setup>
import { computed } from 'vue' 
import ToDo from '../model/ToDo.js'

const props = defineProps({
    todo: {
        type: ToDo,
        required: true
    },
})

const itemStatus = computed(() => {
    if (props.todo.completed === true) {
        return 'done-bg';
    }

    let age = (Date.now() - props.todo.createdAt) / 1000 / 60 / 60 / 24;   // age in days
    
    if (age <= 7) {
        return 'less-than-week-bg';
    }

    if (age <= 14) {
        return 'one-weeks-old-bg';
    }

    return 'two-weeks-old-bg';
})

</script>

<template>
    <div class="item" :class="itemStatus">
        <div class="header">
            <h2>{{  todo.title }}</h2> Priority: {{ todo.priority }}
        </div>

        <div class="content description">
            <pre>{{ todo.description }}</pre>
        </div>

        <div class="footer">
            <div class="info">
               {{  todo.displayCreatedDate() }}
            </div>
            <div class="actionlist">
                <img @click="$emit('markTodoDone', todo)" alt="Add Todo" src="../assets/done.svg" class="actionlogo doneactionlogo" :class="todo.completed ? 'done-completed' : 'done-notcompleted'" />
                <img @click="$emit('editTodo', todo)" src="../assets/edit.svg" class="actionlogo editactionlogo" />
                <img @click="$emit('deleteTodo', todo)" src="../assets/delete.svg" class="actionlogo deleteactionlogo" />
            </div>
        </div>
    </div>
</template>

<style scoped>
.item {
    border-radius: 4px;
    border-color: black;
    border-style: solid;
    border-width: 1px;
    width: 100%;
    margin: 5px;
    padding: 5px;
    display: flex;
    flex-direction: column;
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

.description {
    background-color: rgba(220,220,220,1);
    padding: 5px;
}

/* https://stackoverflow.com/questions/248011/how-do-i-wrap-text-in-a-pre-tag */
pre {
    white-space: pre-wrap;       /* Since CSS 2.1 */
    white-space: -moz-pre-wrap;  /* Mozilla, since 1999 */
    white-space: -pre-wrap;      /* Opera 4-6 */
    white-space: -o-pre-wrap;    /* Opera 7 */
    word-wrap: break-word;       /* Internet Explorer 5.5+ */
}

.actionlist {https://stackoverflow.com/questions/248011/how-do-i-wrap-text-in-a-pre-tag
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
  img.doneactionlogo:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }

  img.deleteactionlogo:hover {
      background-color: hsla(0, 100%, 37%, 0.7);
    }

    img.editactionlogo:hover {
            background-color: hsla(260, 100%, 55%, 0.2);
        }
}

img.done-completed {
    background-color: hsla(160, 100%, 37%, 0.2);
}
</style>