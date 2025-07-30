<template>
  <div>
    <h2>Add New Asset</h2>
    <input v-model="name" placeholder="Name" />
    <input v-model="category" placeholder="Category" />
    <input v-model="status" placeholder="Status (Available/Assigned)" />
    <button @click="addAsset">Add</button>
  </div>
</template>

<script setup>
import axios from 'axios'
import { ref } from 'vue'

const name = ref('')
const category = ref('')
const status = ref('')

const addAsset = async () => {
  await axios.post('http://localhost:8080/api/assets', {
    name: name.value,
    category: category.value,
    status: status.value
  })
  name.value = category.value = status.value = ''
}
</script>
