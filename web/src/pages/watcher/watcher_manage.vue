<template>
    <v-card
        title="watcher管理"
        flat
    >
        <template v-slot:text>
            <v-row align="center">
                <v-col cols="10">
                    <v-text-field
                        v-model="search"
                        label="查找watcher"
                        prepend-inner-icon="mdi-magnify"
                        variant="outlined"
                        hide-details
                        single-line
                    ></v-text-field>
                </v-col>
                <v-col cols="2">
                    <v-btn
                        color="primary"
                        @click="dialog = true"
                    >
                    添加watcher
                </v-btn>

                <v-dialog
                    v-model="dialog"
                    max-width="600px"
                >
                    <v-card>
                        <v-card-title>添加watcher</v-card-title>
                        <!-- AddWatcher component should be defined elsewhere -->
                        <AddWatcher />
                    </v-card>
                </v-dialog>

                </v-col>
            </v-row>
        </template>
        <v-data-table
            :headers="headers"
            :items="watchers"
            :search="search"
        >
            <template v-slot:item.action="{ item }">
                <v-btn
                    icon
                    @click="del = true"
                >
                    <v-icon>mdi-delete</v-icon>
                </v-btn>

                <v-btn
                    icon
                    @click="update = true"
                >
                    <v-icon>mdi-pencil</v-icon>
                </v-btn>
            </template>
        </v-data-table>
        <v-dialog
            v-model="update"
            max-width="600px">
            </v-dialog>
    </v-card>
</template>

<script setup lang="ts" name="WatcherManage">

import { getWatchers } from '@/api/watcher'
import { onMounted, ref } from 'vue'
import { type Watcher } from '@/api/watcher'

const headers = [
    { text: '名称', value: 'name' },
    { text: '类型', value: 'type' },
    { text: '资源', value: 'resource' },
    { text: '点击', value: 'click' },
    { text: '品牌', value: 'brand' },
    { text: '标签', value: 'tag' },
    { text: '操作', value: 'action', sortable: false }
]

const watchers = ref<Watcher[]>([])

const search = ref('')
const dialog = ref(false)
const del = ref(false)
const update = ref(false)


onMounted(async () => {
    const res:any = await getWatchers()
    if (res.status) {
        watchers.value = res.watcher
    } else {
        console.error('获取监控列表失败:', res.message)
    }
})

</script>
