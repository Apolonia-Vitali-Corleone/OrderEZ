<template>
    <div>
        <el-table :data="goods">
            <el-table-column prop="id" label="id"/>
            <el-table-column prop="good_name" label="good_name"/>
        </el-table>
    </div>
    <div>
        <el-button type="primary" @click="dialogFormVisible = true">
            新增商品
        </el-button>

        <el-dialog v-model="dialogFormVisible" title="Shipping address" width="500">
            <el-form :model="form">
                <el-form-item label="商品名称" :label-width="formLabelWidth">
                    <el-input v-model="form.name" autocomplete="off"/>
                </el-form-item>
                <el-form-item label="商品价格" :label-width="formLabelWidth">
                    <el-input v-model="form.price" autocomplete="off"/>
                </el-form-item>
                <el-form-item label="商品库存" :label-width="formLabelWidth">
                    <el-input v-model="form.stock" autocomplete="off"/>
                </el-form-item>
            </el-form>

            <template #footer>
                <div class="dialog-footer">
                    <el-button @click="dialogFormVisible = false">取消</el-button>
                    <el-button type="primary" @click="handleAddGood">提交</el-button>
                </div>
            </template>
        </el-dialog>

        <el-button type="primary" @click="prevPage">上一页</el-button>
        <el-button type="primary" @click="nextPage">下一页</el-button>
    </div>
</template>

<script setup>
import axiosInstance from "@/utils/axiosInstance.js";
import {onMounted, reactive, ref} from "vue";
import {ElNotification} from "element-plus";

const pageSize = ref(10); // 每页显示的商品数量
const currentPage = ref(1); // 当前页码

const getAllGoods = async () => {
    try {
        const response = await axiosInstance.get('/good/goods', {
            params: {
                page: currentPage.value,
                pageSize: pageSize.value
            }
        });

        // 根据后端返回的数据结构调整
        if (Array.isArray(response.data.goods)) {
            return response.data.goods;
        } else {
            console.warn('返回的数据结构不符合预期，goods 不是数组');
            return [];
        }

    } catch (error) {
        console.error('获取商品失败', error);
        throw error;
    }
};


const goods = ref([]);

const fetchGoods = async () => {
    try {
        goods.value = await getAllGoods();
    } catch (error) {
        console.error('获取商品列表失败', error);
    }
};

onMounted(fetchGoods);

const prevPage = () => {
    if (currentPage.value > 1) {
        currentPage.value--;
        fetchGoods();
    }
};

const nextPage = () => {
    currentPage.value++;
    fetchGoods();
};

const dialogFormVisible = ref(false)
const formLabelWidth = '140px'

const form = reactive({
    name: '',
    price: '',
    stock: ''
})


const addGood = async () => {
    try {
        const response = await axiosInstance.post('/good', {
            name: form.name,
            price: form.price,
            stock: form.stock,
        });

        // 假设后端返回的数据结构包含 code 和 message 字段
        // code 为 200 表示业务操作成功
        const {code, message} = response.data;

        if (code === 200) {
            console.log('商品添加成功', response.data);
            ElNotification({
                title: 'Success',
                message: message || '商品添加成功',
                type: 'success'
            });
        } else {
            console.error('商品添加失败，业务错误', response.data);
            ElNotification({
                title: 'Error',
                message: message || '商品添加失败，请稍后重试',
                type: 'error'
            });
        }
    } catch (error) {
        console.error('商品添加失败，网络或其他错误', error);
        ElNotification({
            title: 'Error',
            message: '商品添加失败，请检查网络或稍后重试',
            type: 'error'
        });
    }
}

const handleAddGood = () => {
    dialogFormVisible.value = false;
    addGood();
};

</script>

<style scoped>

</style>