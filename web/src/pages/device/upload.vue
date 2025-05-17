<template>
    <v-sheet class="mx-auto pa-5" max-width="400" rounded="lg" elevation="2">
        <v-form v-model="valid" ref="formRef">
            <v-container>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-select v-model="host" 
                        :items="nodes" 
                        item-title="name" 
                        item-value="host" 
                        :rules="HostRules"
                        label="选择节点" 
                        placeholder="请选择节点" 
                        required 
                        full-width></v-select>
                    </v-col>
                </v-row>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-file-input v-model="file" 
                        :rules="FileRules" 
                        label="上传文件" 
                        placeholder="请选择文件" 
                        required
                        full-width></v-file-input>
                    </v-col>
                </v-row>
                <v-row justify="center" align="center" class="mt-4">
                    <v-col cols="12">
                        <v-btn color="primary" block @click="handleSubmit">
                            提交
                        </v-btn>
                    </v-col>
                </v-row>
            </v-container>
        </v-form>
    </v-sheet>
</template>


<script setup lang="ts" name="Upload">
import { ref } from 'vue'



const host = ref('')
const file = ref<File | null>(null)
const formRef = ref()
const valid = ref(false)

const nodes = [
    { name: '节点1', host: '10.86.97.1' },
    { name: '节点2', host: '10.86.97.2' },
    { name: '节点3', host: '10.86.97.3' },
    { name: '节点4', host: '10.86.97.4' },
    { name: '节点5', host: '10.86.97.5' },
    { name: '节点6', host: '10.86.97.6' },
]

const HostRules = [
    (v: string) => !!v || '主机地址不能为空',
]

const FileRules = [
    (v: File | null) => !!v || '文件不能为空',
]

const handleSubmit = () => {
    console.log('表单数据:', {
        host: host.value,
        file: file.value,
    })
    alert('表单数据已提交')
}

</script>