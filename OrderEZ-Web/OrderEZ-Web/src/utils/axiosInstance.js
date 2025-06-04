import axios from 'axios';

// 创建 Axios 实例
const axiosInstance = axios.create({
    baseURL: 'http://127.0.0.1:4444', // 根据实际 API 地址修改
    timeout: 10000, // 请求超时时间
    headers: {
        'Content-Type': 'application/json' // 避免触发 CORS 预检
    }
});

// 请求拦截器
axiosInstance.interceptors.request.use(
    (config) => {
        // 从本地存储中获取 token
        const token = localStorage.getItem('token');
        if (token) {
            // 将 token 添加到请求头中
            config.headers.Authorization = `${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

export default axiosInstance;