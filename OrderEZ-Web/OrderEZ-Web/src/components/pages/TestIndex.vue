<template>
    <el-container class="container">
        <!-- 头部区域 -->
        <el-header class="header">欢迎使用本系统</el-header>

        <el-container class="body-container">
            <!-- 侧边栏 -->
            <el-aside class="aside">
                <el-menu :default-active="currentPage"
                         class="el-menu-vertical-demo"
                         @select="handleMenuItemClick">
                    <el-menu-item index="UserManagement">
                        <el-icon>
                            <icon-menu></icon-menu>
                        </el-icon>
                        <span>用户管理</span>
                    </el-menu-item>
                    <el-menu-item index="GoodManagement">
                        <el-icon>
                            <icon-menu/>
                        </el-icon>
                        <span>商品管理</span>
                    </el-menu-item>

                    <el-menu-item index="CartManagement">
                        <el-icon>
                            <icon-menu/>
                        </el-icon>
                        <span>购物车</span>
                    </el-menu-item>
                </el-menu>
                <div class="spacer" style="height: 50vh"></div>
                <el-button @click="handlerLogout">退出登录</el-button>
            </el-aside>

            <el-container class="content">
                <!-- 主内容区 -->
                <el-main class="main">
                    <!--                    &lt;!&ndash; 错误提示 &ndash;&gt;-->
                    <!--                    <el-alert-->
                    <!--                            v-if="errorMessage"-->
                    <!--                            :title="errorMessage"-->
                    <!--                            type="error"-->
                    <!--                            show-icon-->
                    <!--                            style="margin-bottom: 20px"-->
                    <!--                    />-->
                    <router-view></router-view>
                </el-main>

                <!-- 底部 -->
                <el-footer class="footer">
                    Powered by <span style="color: red;">BOIHUIKE</span>
                </el-footer>
            </el-container>
        </el-container>
    </el-container>
</template>

<script setup>
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import axios from "axios";

const router = useRouter();
const errorMessage = ref('');
const currentPage = ref(localStorage.getItem('currentPage') || 'UserManagement');

// const handleMenuItemClick = (index) => {
//
// };

function handleMenuItemClick(index) {
    currentPage.value = index;
    localStorage.setItem('currentPage', index);

    switch (index) {
        case 'UserManagement':
            router.push({name: 'UserManagement'});
            break;
        case 'GoodManagement':
            router.push({name: 'GoodManagement'});
            break;
        case 'CartManagement':
            router.push({name: 'CartManagement'});
            break;
        default:
            break;
    }
}

// const handlerLogout = async () => {
//     const token = localStorage.getItem('token');
//     if (!token) {
//         errorMessage.value = '未找到令牌，请重新登录';
//         return;
//     }
//
//     const loadingInstance = ElLoading.service({
//         lock: true,
//         text: '正在退出登录，请稍候...',
//         background: 'rgba(0, 0, 0, 0.7)',
//     });
//
//     try {
//         await logout(token);
//         localStorage.removeItem('token');
//         delete axiosInstance.defaults.headers.common['Authorization'];
//         loadingInstance.close();
//         window.location.href = '/login';
//     } catch (error) {
//         loadingInstance.close();
//         errorMessage.value = '登出失败，请检查网络';
//         console.error('登出失败', error);
//     }
// };

// const logout = async (token) => {
//     axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token}`;
//     await axiosInstance.post('/user/logout');
// };


function handlerLogout() {
    const token = localStorage.getItem('token');
    if (!token) {
        errorMessage.value = '未找到令牌，请重新登录';
        window.location.href = '/login';
    }

    const config = {
        baseURL: 'http://127.0.0.1:4444', // 根据实际 API 地址修改
        timeout: 10000, // 请求超时时间
        headers: {
            'Content-Type': 'application/json', // 避免触发 CORS 预检
            'Authorization': `${token}`
        }
    };

    axios.post("/user/logout", null, config)
        .then(response => {
            localStorage.removeItem('token');
            alert("登出成功");
            window.location.href = '/login';
        })
        .catch(error => {
            let errorMsg = '登出失败，请检查网络';
            if (error.response) {
                // 服务器返回了错误状态码
                errorMsg = `登出失败，服务器返回状态码: ${error.response.status}`;
            } else if (error.request) {
                // 请求发送了，但没有收到响应
                errorMsg = '登出失败，没有收到服务器响应';
            }
            localStorage.removeItem('token');
            alert(errorMsg);
            window.location.href = '/login';
        });
}

</script>

<style scoped>
.container {
    height: 97vh;
}

.header {
    background-color: white;
    color: black;
    line-height: 60px;
    font-size: 30px;
    flex-shrink: 0;
}

.footer {
    background-color: white;
    color: black;
    text-align: center;
    line-height: 50px;
    font-size: 18px;
    flex-shrink: 0;
}

.aside {
    background-color: white;
    width: 200px;
    color: black;
    text-align: center;
}

.content {
    display: flex;
    flex-direction: column;
    color: white;
}

.main {
    background-color: white;
    flex: 1;
}
</style>
