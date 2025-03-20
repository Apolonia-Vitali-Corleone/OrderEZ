import {createRouter, createWebHistory} from 'vue-router';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import Index from "../components/Index.vue";

const routes = [
    {path: '/login', component: Login},
    {path: '/register', component: Register},
    {path: '/index', component: Index}
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;