<template>
    <v-card
        title="节点管理"
        flat
    >
        <template v-slot:text>
            <v-row align="center">
                <v-col cols = "10">
                    <v-text-field
                        v-model="search"
                        label="查找节点"
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
                    添加节点
                </v-btn>

                <v-dialog
                    v-model="dialog"
                    max-width="600px"
                >
                    <v-card>
                        <v-card-title>添加节点</v-card-title>
                        <AddNode />
                    </v-card>
                </v-dialog>

                </v-col>
            </v-row>
        </template>
        <v-data-table
            :headers="headers"
            :items="nodes"
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
                <v-dialog
                    v-model="update"
                    max-width="600px"
                >
                    <v-card>
                        <v-card-title>更新节点</v-card-title>
                        <UpdataNode :parentData="item.host"/>
                    </v-card>
                </v-dialog>
            </template>
        </v-data-table>
    </v-card>
</template>

<script setup lang="ts" name="DeviceManage">
import AddNode from '@/components/node/add_node.vue'
import UpdataNode from '@/components/node/update_node.vue'
import { h, onMounted, ref } from 'vue'
import { type nodesResponse} from '@/api/response_data'
import { getNodes } from '@/api/device'

const search = ref('')

const dialog = ref(false)
const nodes = ref()

const del = ref(false)
const update = ref(false)

const headers = [
    { text: '节点名称', value: 'name' },
    { text: '节点地址', value: 'host' },
    { text: '操作', value: 'action', sortable: false },
]

onMounted(async() => {
    const res:nodesResponse = await getNodes()
    console.log('获取节点列表', res)
    if (res.status) {
        nodes.value = res.nodes
        console.log('获取节点列表成功', res.nodes)
    } else {
        alert('获取节点列表失败: ' + res.message)
    }
})


</script>
