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
                <v-col cols="2">
                    <v-btn
                        width="20px"
                        color="primary"
                        @click="refreshPhoneList"
                    >
                        刷新
                    </v-btn>
                </v-col>
            </v-row>
        </template>
        <v-data-table
            :headers="headers"
            :items="phoneList"
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
                <v-card-title>编辑手机信息</v-card-title>
                <UpdatePhone v-if="editPhone" :PhoneData="editPhone"/>
            </v-card>
        </v-dialog>
    </v-card>
</template>

<script setup lang="ts" name="PhoneManage">
import { onMounted, ref } from 'vue';
import AddPhone from '@/components/phone/add_phone.vue';
import { deletePhone, getAllPhoneInfo, type Phone } from '@/api/phone';
import { processPhoneInfo } from '@/utils/processing_phone_info';
import UpdatePhone from '@/components/phone/update_phone.vue';

const search = ref('');
const addphone = ref(false);
const editPhone = ref<Phone | null>(null);
const update = ref(false);


const headers = [
    { text: '序列号', value: 'serial' },
    { text: '生产厂商', value: 'manufacturer' },
    { text: '型号', value: 'model' },
    { text: '安卓版本', value: 'androidVersion'},
    { text: 'CPU架构', value: 'cpuabi' },
    { text: '市场名称', value: 'marketName' },
    { text: '市场名称符号', value: 'marketNameSymbol' },
    { text: '操作', value: 'action', sortable: false }
];

const phoneList = ref([]);

onMounted(async () => {
    const res:any = await getAllPhoneInfo();
    let data = [];
    if (res.status) {
        data = res.phones;
        data = processPhoneInfo(data);
        // 如果 data 是 Map 数组，转成对象数组
        if (data.length && data[0] instanceof Map) {
            data = data.map(m => Object.fromEntries(m));
        }
        phoneList.value = data;
    }
});

const openEditDialog = (item: Phone) => {
    editPhone.value = item;
    update.value = true;
};

const handleDelete = async (item: Phone) => {
    if (confirm(`确定删除手机信息: ${item.serial} 吗？`)) {
        const res:any = await deletePhone(item.serial);
        if (res.status) {
            alert('删除成功: ' + res.message);
            refreshPhoneList();
        } else {
            alert('删除失败: ' + res.message);
        }
    }
};

const refreshPhoneList = async () => {
    const res:any = await getAllPhoneInfo();
    let data = [];
    if (res.status) {
        data = res.phones;
        data = processPhoneInfo(data);
        // 如果 data 是 Map 数组，转成对象数组
        if (data.length && data[0] instanceof Map) {
            data = data.map(m => Object.fromEntries(m));
        }
        phoneList.value = data;
    }
};
</script>
