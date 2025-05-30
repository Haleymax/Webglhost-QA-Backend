<template>
    <v-card
        title="watcher管理"
        flat
    >
        <template v-slot:text>
            <v-row align="center">
                <v-col cols="6">
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
                        <AddWatcher/>
                    </v-card>
                </v-dialog>

                </v-col>

                <v-col cols="2">
                    <v-btn
                        color="primary"
                        @click="update_redis = true"
                    >
                    更新redis缓存
                </v-btn>

                <v-dialog
                    v-model="update_redis"
                    max-width="600px"
                >
                    <v-card>
                        <v-card-title>更新redis缓存</v-card-title>
                        <UpdateRedis/>
                    </v-card>
                </v-dialog>

                </v-col>
                <v-col cols="2">
                    <v-btn
                        color="primary"
                        @click=handleUpdate
                    >
                    刷新
                </v-btn>
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
                    @click="handleDelete(item)"
                >
                    <v-icon>mdi-delete</v-icon>
                </v-btn>

                <v-btn
                    icon
                    @click="openEditDialog(item)"
                >
                    <v-icon>mdi-pencil</v-icon>
                </v-btn>
            </template>
        </v-data-table>
        <v-dialog
            v-model="update"
            max-width="600px"
        >
            <v-card>
                <v-card-title>更新watcher</v-card-title>
                <UpdateWatcher v-if="editWatcher" :parentData="editWatcher"/>
            </v-card>
        </v-dialog>
    </v-card>
</template>

<script setup lang="ts" name="WatcherManage">

import { getWatchers } from '@/api/watcher'
import { onMounted, ref } from 'vue'
import { type Watcher , deleteWatcher} from '@/api/watcher'
import AddWatcher from '@/components/watcher/add_watcher.vue'
import UpdateWatcher from '@/components/watcher/update_watcher.vue'
import UpdateRedis from '@/components/watcher/update_redis.vue'

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
const update_redis = ref(false)
const update = ref(false)
const editWatcher = ref<Watcher | null>(null)

onMounted(async () => {
    const res:any = await getWatchers()
    if (res.status) {
        watchers.value = res.watcher
    } else {
        console.error('获取监控列表失败:', res.message)
    }
})

const handleDelete = async (watcher: Watcher) => {
    console.log('删除watcher ID:', watcher.id)
    const res = await deleteWatcher(watcher)
    if (res.status) {
        console.log('删除成功:', res.message)
        alert('删除成功: ' + res.message)
        handleUpdate(watcher)
    } else {
        console.error('删除失败:', res.message)
    }
}

const handleUpdate = async (watcher?: Watcher) => {
    const res = await getWatchers()
    if (res.status) {
        watchers.value = res.watcher
    } else {
        console.error('获取监控列表失败:', res.message)
    }
}

const openEditDialog = (watcher: Watcher) => {
    editWatcher.value = watcher
    update.value = true
}

</script>
