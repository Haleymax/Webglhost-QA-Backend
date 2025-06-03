<template>
    <v-card
        title="手机信息管理"
        flat
    >
        <template v-slot:text>
            <v-row align="center">
                <v-col cols="8">
                    <v-text-field
                        v-model="search"
                        label="查找手机信息"
                        prepend-inner-icon="mdi-magnify"
                        variant="outlined"
                        hide-details
                        single-line
                    ></v-text-field>
                </v-col>

                <v-col cols="2">
                    <v-btn
                        color="primary"
                        @click="addphone = true"
                    >
                        添加手机信息
                    </v-btn>

                    <v-dialog
                        v-model="addphone"
                        max-width="600px"
                    >
                        <v-card>
                            <v-card-title>添加手机信息</v-card-title>
                            <AddPhone />
                        </v-card>
                    </v-dialog>
                </v-col>
            </v-row>
        </template>

        <v-data-table
            :headers="headers"
            :items="phoneList"
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
    </v-card>
</template>

<script setup lang="ts" name="PhoneManage">
import { ref } from 'vue';
import AddPhone from '@/components/phone/add_phone.vue';

const search = ref('');
const addphone = ref(false);


const headers = [
    { text: '手机ID', value: 'id' },
    { text: '手机名称', value: 'name' },
    { text: '手机状态', value: 'status' },
    { text: '操作', value: 'action', sortable: false }
];

const phoneList = ref([
    { id: 1, name: '手机A', status: '在线' },
    { id: 2, name: '手机B', status: '离线' },
    { id: 3, name: '手机C', status: '在线' }
]);


</script>