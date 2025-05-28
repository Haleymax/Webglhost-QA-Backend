<template>
    <v-sheet>
        <v-form v-model="valid" ref="formRed">
            <v-container>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="user"
                            :rules="userRules"
                            label="用户名"
                            placeholder="请输入新的用户名"
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
                            placeholder="请输入新的密码"
                            required
                            full-width
                            ></v-text-field>
                    </v-col>
                </v-row>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="name"
                            :rules="nameRules"
                            label="节点名称"
                            placeholder="请输入新的节点名称"
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

<script setup lang="ts" name="UpdataNode">
import { ref } from 'vue';
import { type updata, updataNode} from '@/api/device'
import type { baseResponse } from '@/api/response_data'


const valid = ref(true)
const formRef = ref()

const props = defineProps({
    parentData: String
})

const host:string = String(props.parentData)
const user = ref('')
const password = ref('')
const name = ref('')

const hostRules = [
    (v: string) => !!v || '主机地址不能为空'
]

const userRules = [
    (v: string) => !!v || '用户名不能为空'
]

const passwordRules = [
    (v: string) => !!v || '密码不能为空'
]

const nameRules = [
    (v: string) => !!v || '节点名称不能为空'
]

const handleSubmit = async() => {
    
    if (valid.value) {
        const node:updata = {
            host: host,
            user: user.value,
            password: password.value,
            name: name.value
        }
        console.log('表单验证成功', node)
        const res:baseResponse = (await updataNode(node)).data
        if (res.status) {
            alert('更新节点成功: ' + res.message)
        } else {
            alert('更新节点失败: ' + res.message)
        }
    } else {
        console.log('表单验证失败')
    }
}
</script>