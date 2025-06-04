<template>
    <div class="container">
        <div style="height: 20%"></div>
        <el-text style="font-size: 36px;color:white">点餐爽后台系统</el-text>
        <div style="height: 2vh"></div>

        <el-card style="max-width: 480px">
            <el-text style="font-size: 26px">登录</el-text>
            <div style="height: 2vh"></div>
            <el-form label-width="80px">
                <el-form-item label="用户名">
                    <el-input v-model="username" placeholder="请输入用户名"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="password" type="password" placeholder="请输入密码"></el-input>
                </el-form-item>
                <el-button @click="handleLogin">登录</el-button>
            </el-form>


            <p>还没有账号？
                <router-link to="/register">去注册</router-link>
            </p>
            <!--            <el-alert v-if="errorMessage" :message="errorMessage" type="error" show-icon>{{ errorMessage }}</el-alert>-->
        </el-card>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    justify-content: center; /* 水平居中 */
    align-items: center; /* 垂直居中 */
    min-height: 97vh;
    background-color: white;
    flex-direction: column; /* 让子元素垂直排列 */
}
</style>

<script setup>
import {onMounted, ref} from 'vue';
import {useRouter} from 'vue-router';
import {ElButton, ElCard, ElForm, ElFormItem, ElInput, ElText} from 'element-plus';
import axios from "axios";

const username = ref('');
const password = ref('');
// const errorMessage = ref('');
const router = useRouter();

// const isTokenExists = () => {
//
// };

function isTokenExists() {
    let token = localStorage.getItem('token');
    if (token != null) {
        router.push('./index');
    }
}

// const handleLogin = async () => {
//     try {
//         const response = await axiosInstance.post('/user/login', {
//             username: username.value,
//             password: password.value
//         });
//
//         console.log('登录成功', response.data);
//
//         const token = response.data.token;
//         localStorage.setItem('token', token);
//
//         console.log('token:', localStorage.getItem('token'));
//
//         await router.push('./index');
//     } catch (error) {
//         console.error('登录失败', error);
//         errorMessage.value = '登录失败，请检查用户名和密码';
//     }
// };

// async function handleLogin(){
//
// }


function handleLogin() {
    // 创建一个 config 对象，包含所需的配置信息
    const config = {
        baseURL: 'http://127.0.0.1:4444', // 根据实际 API 地址修改
        timeout: 10000, // 请求超时时间
        headers: {
            'Content-Type': 'application/json' // 避免触发 CORS 预检
        }
    };

    axios.post('/user/login', {
        username: username.value,
        password: password.value
    }, config)
        .then(response => {
            console.log('登录成功', response.data);

            localStorage.setItem('token', response.data.token);

            console.log('高传晔，token已经被存储起来：token:', localStorage.getItem('token'));

            router.push('./index');
        })
        .catch(error => {
            console.error('登录失败', error);
            // errorMessage.value = '登录失败，请检查用户名和密码';
            alert("登录失败，请检查用户名和密码");
        });
}

onMounted(isTokenExists);

</script>