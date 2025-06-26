<script setup>
import { computed } from 'vue' 


defineProps({
  data: {
    type: Object,
    required: true,
  },
})

const itemStatus = computed(() => {
  return 'normal'
})
</script>

<template>
    <item :class="itemStatus">
        <content>
            <h2>{{  data.title }}</h2>
            <p>{{ data.description }}</p>
        </content>

        <footer>
            <info>
                date here
            </info>
            <actionlist>
                <a href="http://www.ic.ac.uk"><img alt="Add Todo" src="../assets/done.svg" class="actionlogo" /></a>
                <a href=""><img alt="Add Todo" src="../assets/delete.svg" class="actionlogo" /></a>
            </actionlist>
        </footer>
    </item>
</template>

<style scoped>
item {
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

footer {
    margin-top: auto;
}

content {
    margin-bottom: 15px;
}

actionlist {
    float: right;
    height: auto;
    line-height: 1;
}

.actionlogo {
    height: 25px;
    width: 25px;
}

a {
    margin-left: 5px;
    display:inline-block;
    border-radius: 5px;
    height: 25px;
}

@media (min-width: 600px) {
    item {
        width: 250px;
    }
}

@media (hover: hover) {
  a:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}
</style>