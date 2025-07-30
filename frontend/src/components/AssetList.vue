<template>
  <div>
    <h2>Asset List</h2>
    <ul>
      <li v-for="asset in assets" :key="asset.id">
        {{ asset.name }} ({{ asset.category }}) - {{ asset.status }}
      </li>
    </ul>
  </div>
</template>

<script setup>
import axios from 'axios'
import { ref, onMounted } from 'vue'

const assets = ref([])

const loadAssets = async () => {
  const res = await axios.get('http://localhost:8080/api/assets')
  assets.value = res.data
}

onMounted(() => loadAssets())
</script>
