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
                    @click="() => { console.log(item) }"
                >
                    <v-icon>mdi-delete</v-icon>
                </v-btn>
                <v-btn
                    icon
                    @click="() => { console.log(item) }"
                >
                    <v-icon>mdi-pencil</v-icon>
                </v-btn>
            </template> 
        </v-data-table>
    </v-card>
</template>

<script setup lang="ts" name="DeviceManage">
import AddNode from '@/components/add_node.vue'
import { h, ref } from 'vue'

const search = ref('')

const dialog = ref(false)

const headers = [
    { text: '节点名称', value: 'name' },
    { text: '节点地址', value: 'host' },
    { text: '操作', value: 'action', sortable: false },
]

const nodes = [
    { name: '节点1', host: '10.86.97.1' },
    { name: '节点2', host: '10.86.97.2' },
    { name: '节点3', host: '10.86.97.3' },
    { name: '节点4', host: '10.86.97.4' },
    { name: '节点5', host: '10.86.97.5' },
    { name: '节点6', host: '10.86.97.6' },
]


</script>