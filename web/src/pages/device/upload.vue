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
import { h, onMounted, ref } from 'vue'
import { type nodesResponse} from '@/api/response_data'
import { getNodes} from '@/api/device'
import { uploadFile } from '@/api/device'



const host = ref('')
const file = ref<File | null>(null)
const formRef = ref()
const valid = ref(true)

const nodes = ref([])

const HostRules = [
    (v: string) => !!v || '主机地址不能为空',
]

const FileRules = [
    (v: File | null) => !!v || '文件不能为空',
]

const handleSubmit = async () => {
    const isValid = await formRef.value?.validate?.()
    if (!isValid) return
    if (!file.value) {
        alert('请选择文件')
        return
    }
    try {
        const formData = new FormData()
        formData.append('host', host.value)
        formData.append('file', file.value)
        const res: nodesResponse = await uploadFile(formData)
        if (res.status) {
            alert('上传文件成功')
        } else {
            alert('上传文件失败: ' + res.message)
        }
    } catch (e) {
        alert('上传异常')
    } finally {
    }
}

onMounted(async() => {
    const res:nodesResponse = await getNodes()
    console.log('获取节点列表', res)
    if (res.status) {
        nodes.value = res.nodes
    } else {
        alert('获取节点列表失败: ' + res.message)
    }
})

</script>