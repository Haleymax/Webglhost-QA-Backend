<template>
    <v-sheet>
        <v-form v-model="valid" ref="formRed">
            <v-container>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="name"
                            :rules="nameRules"
                            label="name"
                            placeholder="请输入watcher名字"
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
                            placeholder="请输入对应的触发事件"
                            required
                            full-width
                            ></v-text-field>
                    </v-col>
                </v-row>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-combobox
                            v-model="brand"
                            :rules="brandRules"
                            label="brand"
                            placeholder="请输入watcher对应的手机品牌（可输入多个，用回车分隔）"
                            multiple
                            chips
                            clearable
                            required
                            full-width
                            ></v-combobox>
                    </v-col>
                </v-row>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-combobox
                            v-model="tag"
                            :rules="tagRules"
                            label="tag"
                            placeholder="请输入watcher对应的tag（可输入多个，用回车分隔）"
                            multiple
                            chips
                            clearable
                            required
                            full-width
                            ></v-combobox>
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

<script setup lang="ts" name="UpdateWatcher">
import { ref } from 'vue';
import type { Watcher } from '@/api/watcher';
import type { baseResponse } from '@/api/response_data';
import { updateWatcher } from '@/api/watcher';

const valid = ref(true)

const props = defineProps<{
    parentData: Watcher
}>()

const formRed = ref(null)

const id = ref(props.parentData.id || '')
const name = ref(props.parentData.name || '')
const resource = ref(props.parentData.resource || '')
const event = ref(props.parentData.click || '')
const brand = ref(props.parentData.brand || [])
const tag = ref(props.parentData.tag || [])

const nameRules = [
    (v: string) => !!v || '名称不能为空'
]

const resourceRules = [
    (v: string) => !!v || '资源不能为空'
]

const eventRules = [
    (v: string) => !!v || '事件不能为空'
]

const brandRules = [
    (v: string[]) => !!v || '品牌不能为空'
]

const tagRules = [
    (v: string[]) => !!v || '标签不能为空'
]

const handleSubmit = async() => {
    if (valid.value) {
        const watcher: Watcher = {
            id: id.value,
            name: name.value,
            resource: resource.value,
            click: event.value,
            brand: brand.value,
            tag: tag.value
        }
        const res: baseResponse = await updateWatcher(watcher);
        if (res.status) {
            alert('更新成功: ' + res.message);
        } else {
            alert('更新失败: ' + res.message);
        }
    } else {
        console.error('表单验证失败');
    }
}
</script>
