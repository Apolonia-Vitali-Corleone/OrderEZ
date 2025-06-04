<template>
    <div>
        <el-table :data="users">
            <el-table-column prop="id" label="id"/>
            <el-table-column prop="username" label="username"/>
        </el-table>
    </div>
    <div>
        <el-button type="primary" @click="prevPage">上一页</el-button>
        <el-button type="primary" @click="nextPage">下一页</el-button>
    </div>
</template>

<script setup>
import axiosInstance from "@/utils/axiosInstance.js";
import {onMounted, ref} from "vue";

const pageSize = ref(10); // 每页显示的商品数量
const currentPage = ref(1); // 当前页码

const getAllUsers = async () => {
    try {
        const response = await axiosInstance.get('/user/users');

        // 根据后端返回的数据结构调整
        if (Array.isArray(response.data.users)) {
            return response.data.users;
        } else {
            console.warn('返回的数据结构不符合预期，users 不是数组');
            return [];
        }

    } catch (error) {
        console.error('获取用户失败', error);
        throw error;
    }
};


const users = ref([]);

const fetchUsers = async () => {
    try {
        users.value = await getAllUsers();
    } catch (error) {
        console.error('获取用户列表失败', error);
    }
};

onMounted(fetchUsers);


const prevPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--;
        fetchUsers();
    }
};

const nextPage = () => {
    currentPage.value++;
    fetchUsers();
};
</script>


<style scoped>

</style>