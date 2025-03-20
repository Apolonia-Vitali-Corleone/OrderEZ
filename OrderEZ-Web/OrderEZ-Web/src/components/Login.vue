<template>
  <div>
    <h1>登录</h1>
    <form @submit.prevent="handleLogin">
      <input v-model="username" type="text" placeholder="用户名"/>
      <input v-model="password" type="password" placeholder="密码"/>
      <button type="submit">登录</button>
    </form>
    <div v-if="errorMessage">{{ errorMessage }}</div>
  </div>
</template>

<script setup>
import {ref} from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router';

const username = ref('')
const password = ref('')
const errorMessage = ref('')
const router = useRouter()

const handleLogin = async () => {
  console.log('登录信息：', {username: username.value, password: password.value})
  try {
    const response = await axios.post('http://127.0.0.1:4444/user/login', {
      username: username.value,
      password: password.value,
    });
    console.log('登录成功', response.data)
    await router.push('/index')
  } catch (error) {
    console.log('登录失败!', error)
    errorMessage.value = '登录失败！请检查用户名和密码'
  }
}
</script>