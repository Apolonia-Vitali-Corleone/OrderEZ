<template>
    <div class="register_container">
        <h1>注册</h1>
        <form @submit.prevent="handleRegister">
            <input v-model="username" type="text" placeholder="用户名"/>
            <input v-model="password" type="password" placeholder="密码"/>
            <button type="submit">注册</button>
        </form>

        <!-- 登录按钮，跳转到登录页 -->
        <p>已经有账号？
            <router-link to="/login">去登录</router-link>
        </p>
        <div v-if="errorMessage">{{ errorMessage }}</div>

    </div>
</template>

<script setup>
import {ref} from 'vue'

const username = ref('')
const password = ref('')
const errorMessage = ref('');

const handleRegister = () => {
    console.log('注册信息：', {username: username.value, password: password.value})
    // 这里可以添加实际的注册逻辑
}
</script>

<style scoped>
.register_container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    min-height: 97vh;
    background-image: url('../asserts/lyf.jpg');
}

</style>

