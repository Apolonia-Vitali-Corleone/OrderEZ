<template>
  <div>
    <h1>注册</h1>
    <form @submit.prevent="handleRegister">
      <input v-model="username" type="text" placeholder="用户名" />
      <input v-model="password" type="password" placeholder="密码" />
      <button type="submit">注册</button>
    </form>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const username = ref('')
const password = ref('')

const handleRegister = () => {
  console.log('注册信息：', { username: username.value, password: password.value })
  // 这里可以添加实际的注册逻辑
}
</script>