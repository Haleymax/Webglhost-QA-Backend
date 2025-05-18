<template>
    <v-sheet>
        <v-form v-model="valid" ref="formRef">
        <v-container>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="host"
                :rules="hostRules"
                label="主机地址"
                placeholder="请输入主机地址"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="port"
                :rules="portRules"
                label="端口"
                placeholder="请输入端口"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="user"
                :rules="userRules"
                label="用户名"
                placeholder="请输入用户名"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="password"
                :rules="passwordRules"
                label="密码"
                placeholder="请输入密码"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center" class="mt-4">
            <v-col cols="12">
                <v-btn color="primary" block @click="handleSubmit">提交</v-btn>
            </v-col>
            </v-row>
        </v-container>
        </v-form>
    </v-sheet>
</template>

<script setup lang="ts" name="AddNode">
import { ref } from 'vue'
import  {addNode, type Node} from '@/api/device'
import type { baseResponse } from '@/api/response_data'

const host = ref('')
const port = ref()
const user = ref('')
const password = ref('')

const hostRules = [
    (v: string) => !!v || '主机地址不能为空',
]
const portRules = [
    (v: number) => !!v || '端口不能为空',
    (v: number) => /^\d+$/.test(v) || '端口必须为数字',
]
const userRules = [
    (v: string) => !!v || '用户名不能为空',
]
const passwordRules = [
    (v: string) => !!v || '密码不能为空',
]
const formRef = ref()
const valid = ref(false)
const handleSubmit = async() => {
    if (valid.value) {
        const node: Node = {
            host: host.value,
            port: port.value,
            user: user.value,
            password: password.value,
        }

        const res: baseResponse = (await addNode(node)).data
        if (res.status) {
            alert('添加节点成功: ' + res.message)
        } else {
            alert('添加节点失败: ' + res.message)
        }
        

        console.log('添加节点', { host: host.value, port: port.value, user: user.value, password: password.value })
    } else {
        console.log('表单验证失败')
    }
}

</script>