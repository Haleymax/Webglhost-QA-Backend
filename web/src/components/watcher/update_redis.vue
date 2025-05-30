<template>
    <v-sheet>
        <v-form v-model="valid" ref="formRef">
        <v-container>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-select
                    label="Env环境"
                    v-model="env"
                    :items="env_items"
                ></v-select>
            </v-col>
            </v-row>

            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-select
                    label="runtime环境"
                    v-model="runtime"
                    :items="runtime_items"
                ></v-select>
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

<script setup lang="ts" name="UpdateRedis">
import { ref } from 'vue'
import { updateRedis } from '@/api/watcher'

const env = ref('')
const runtime = ref('')

const env_items = ['dev', 'test', 'prod']

const runtime_items = ['core_sdk', 'xiaomi', 'runtime']


const formRef = ref()
const valid = ref(false)
const handleSubmit = async() => {
    if (valid.value) {
        const data:any = {
            env: env.value,
            runtime: runtime.value
        }
        const res = await updateRedis(data)
        if (res.status) {
            console.log('更新redis缓存成功:', res.data)
            alert('更新redis缓存成功: ' + res.message)
        } else {
            console.error('更新redis缓存失败:', res.message)
            alert('更新redis缓存失败: ' + res.message)
        }
        console.log('提交的数据:', watcher)
        alert('更新redis缓存成功: ' + res.message)
    } else {
        console.log('表单验证失败')
        alert('更新redis缓存成功: ' + res.message)
    }
}

</script>
