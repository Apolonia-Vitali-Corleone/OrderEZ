<template>
    <div>
        <el-table :data="cartItems">
            <el-table-column prop="cart_id" label="购物车商品ID"/>
            <el-table-column prop="count" label="购物车商品数量"/>
        </el-table>
    </div>
</template>

<script setup>
import {onMounted, ref} from "vue";
import axios from "axios";

const cartItems = ref([]);


function getCartItems() {
    const token = localStorage.getItem('token');
    if (!token) {
        alert('未找到令牌，请重新登录');
        window.location.href = '/login';
        return;
    }

    axios.get('http://127.0.0.1:4444/cart', {
        headers: {
            'Content-Type': 'application/json', // 避免触发 CORS 预检
            'Authorization': `${token}`
        }
    })
        .then(response => {
            cartItems.value = response.data;
        })
        .catch(error => {
            console.error('错误信息:', error);
            alert('获取购物车商品数据失败，请检查网络或重试。错误信息：' + error);
        });
}

onMounted(getCartItems);

</script>

<style scoped>

</style>