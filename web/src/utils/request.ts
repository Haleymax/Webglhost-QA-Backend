import axios, { type AxiosResponse } from "axios";

// 创建 axios 实例
const service = axios.create({
    timeout: 10000,
});

// 请求拦截器
service.interceptors.request.use(
    (config: any) => {
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);


service.interceptors.response.use(
    (response: AxiosResponse) => {
        return response.data;
    },
    (error) => {
        if (error.response) {
            const { status, data } = error.response;
            switch (status) {
                case 401:
                    alert("未授权，请重新登录");
                    break;
                case 403:
                    alert("没有权限访问");
                    break;
                case 500:
                    alert("服务器错误");
                    break;
                default:
                    alert(data.message || "请求错误");
            }
        } else {
            alert("网络异常，请稍后重试");
        }
        return Promise.reject(error);
    }
);

export default service;