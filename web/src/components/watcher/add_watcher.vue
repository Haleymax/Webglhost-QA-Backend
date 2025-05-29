<template>
    <v-sheet>
        <v-form v-model="valid" ref="formRef">
        <v-container>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="watcherName"
                :rules="watcherNameRules"
                label="name"
                placeholder="请输入Watcher名字"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="resource"
                :rules="resourceRules"
                label="resource"
                placeholder="请输入对应的资源"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="event"
                :rules="eventRules"
                label="event"
                placeholder="请输入事件名称"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="brand"
                :rules="brandRules"
                label="品牌"
                placeholder="请输入watcher对应的品牌"
                required
                full-width
                ></v-text-field>
            </v-col>
            </v-row>
            <v-row justify="center" align="center">
            <v-col cols="12">
                <v-text-field
                v-model="tag"
                :rules="tagRules"
                label="tag"
                placeholder="请输入watcher对应的tag"
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

<script setup lang="ts" name="AddWatcher">
import { ref } from 'vue'
import {splitString} from '@/common/character_processing'
import {addWatcher ,type  Watcher } from '@/api/watcher'


const watcherName = ref('')
const resource = ref('')
const event = ref('')
const brand = ref('')
const tag = ref('')

const watcherNameRules = [
    (v: string) => !!v || 'Watcher名字不能为空',
]

const resourceRules = [
    (v: string) => !!v || '资源不能为空',
]

const eventRules = [
    (v: string) => !!v || '事件名称不能为空',
]

const brandRules = [
    (v: string) => !!v || '品牌不能为空',
]

const tagRules = [
    (v: string) => !!v || 'Tag不能为空',
]


const formRef = ref()
const valid = ref(false)
const handleSubmit = async() => {
    if (valid.value) {
        const watcher:Watcher = {
          name: watcherName.value,
          resource: resource.value,
          click: event.value,
          brand: splitString(brand.value),
          tag: splitString(tag.value),
          id: null
        }
        const res = await addWatcher(watcher)
        if (res.status) {
            console.log('Watcher添加成功:', res.data)
            watcherName.value = ''
            resource.value = ''
            event.value = ''
            brand.value = ''
            tag.value = ''
        } else {
            console.error('添加Watcher失败:', res.message)
        }
        console.log('提交的Watcher数据:', watcher)
    } else {
        console.log('表单验证失败')
    }
}

</script>
