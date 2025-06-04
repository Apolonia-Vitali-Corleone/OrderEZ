import {createApp} from 'vue';
import {createRouter, createWebHistory} from 'vue-router';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import App from '@/App.vue';
import Index from "@/components/pages/TestIndex.vue";
import Login from '@/components/pages/Login.vue';
import Register from "@/components/pages/Register.vue";
import GoodManagement from "@/components/pages/management/GoodManagement.vue";
import UserManagement from "@/components/pages/management/UserManagement.vue";
import CartManagement from "@/components/pages/management/CartManagement.vue";

// 定义路由
const routes = [
    { path: '/', redirect: '/index' },
    { path: '/login', name: 'Login', component: Login },
    { path: '/register', name: 'Register', component: Register },
    {
        path: '/index',
        name: 'Index',
        component: Index,
        meta: { requiresAuth: true },
        children: [
            {
                path: '/user_management',
                name: 'UserManagement',
                component: UserManagement
            },
            {
                path: '/cart_management',
                name: 'CartManagement',
                component: CartManagement
            },
            {
                path: '/good_management',
                name: 'GoodManagement',
                component: GoodManagement
            }
        ]
    }
];

// 创建 Vue Router 实例
const router = createRouter({
    history: createWebHistory(),
    routes
});

// 全局路由守卫
router.beforeEach((to, from, next) => {
    // 获取token
    const token = localStorage.getItem('token');

    if (!token && to.meta.requiresAuth) {
        // ❌ 未登录且访问需要权限的页面 -> 跳转到 Login
        next({name: 'Login'});
    } else if (token && (to.name === 'Login' || to.name === 'Register')) {
        // ✅ 已登录但访问 Login/Register -> 跳转到主页
        next({name: 'Index'});
    } else {
        next();
    }
});

// 创建 Vue 应用
const app = createApp(App);

// 使用ElementPlus
app.use(ElementPlus)

// 挂载 Vue Router
app.use(router);
app.mount('#app');