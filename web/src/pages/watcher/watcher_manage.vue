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
                        label="查找监控"
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
                    添加监控
                </v-btn>

                <v-dialog
                    v-model="dialog"
                    max-width="600px"
                >
                    <v-card>
                        <v-card-title>添加监控</v-card-title>
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

const headers = [
    { text: 'ID', value: 'id' },
    { text: '名称', value: 'name' },
    { text: '类型', value: 'type' },
    { text: '操作', value: 'action', sortable: false }
]

import { ref } from 'vue'

const watchers = ref([
    { id: 1, name: '监控1', type: '类型A' },
    { id: 2, name: '监控2', type: '类型B' },
    { id: 3, name: '监控3', type: '类型C' }
])

const search = ref('')
const dialog = ref(false)
const del = ref(false)
const update = ref(false)
</script>