<template>
    <v-container>
        <v-row>
            <v-col cols="12">
                <h1>Phone Management</h1>
            </v-col>
        </v-row>
        <v-row>
          <v-col cols="6">
                <v-card class="pa-4">
                    <v-card-title class="text-h5">查询手机信息</v-card-title>
                    <v-form @submit.prevent="handFindPhoneInfo">
                        <v-select v-model="host"
                                  :items="nodes"
                                  item-title="name"
                                  item-value="host"
                                  label="选择节点"
                                  placeholder="请选择节点"
                                  required
                                  full-width></v-select>
                        <v-select v-model="querySerial"
                                  :items="PhoneList"
                                  item-title="serial"
                                  item-value="serial"
                                  label="选择手机序列号"
                                  placeholder="请选择手机序列号"
                                  required
                                  full-width></v-select>
                        <v-btn color="primary" type="submit">查询手机信息</v-btn>
                    </v-form>
                </v-card>
            </v-col>
            <v-col cols="6">
                <v-card class="pa-4">
                    <v-card-title class="text-h5">添加手机</v-card-title>
                    <v-form @submit.prevent="handAddPhone">
                        <v-text-field
                          label="serial"
                          v-model="addSerial"
                          :rules="serialRules"
                          placeholder="请输入手机序列号"
                          required
                          ></v-text-field>
                        <v-text-field
                          label="manufacturer"
                          v-model="manufacturer"
                          placeholder="请输入手机制造商"
                          :rules="manufacturerRules"
                          required
                          ></v-text-field>
                        <v-text-field
                          label="model"
                          v-model="model"
                          placeholder="请输入手机型号"
                          :rules="modelRules"
                          required
                        ></v-text-field>
                        <v-text-field
                          label="androidVersion"
                          v-model="androidVersion"
                          placeholder="请输入Android版本"
                          :rules="androidVersionRules"
                          required
                          ></v-text-field>
                        <v-text-field
                          label="cpuabi"
                          v-model="cpuabi"
                          placeholder="请输入CPU ABI"
                          :rules="cpuabiRules"
                          required
                          ></v-text-field>
                        <v-text-field
                          label="marketName"
                          v-model="marketName"
                          placeholder="请输入名称"
                          :rules="marketNameRules"
                          required
                          ></v-text-field>
                        <v-text-field
                          label="marketNameSymbol"
                          v-model="marketNameSymbol"
                          :rules="marketNameSymbolRules"
                          placeholder="请输入名称符号"
                          required></v-text-field>
                        <v-btn color="primary" type="submit">添加</v-btn>
                    </v-form>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script setup lang="ts" name="AddPhone">
import { onMounted, ref, watch} from 'vue';
import { type nodesResponse} from '@/api/response_data'
import { getNodes, getPhoneInfo, getPhoneList} from '@/api/device'
import { type GetPhoneInfo } from '@/api/device'
import { da } from 'vuetify/locale';

const PhoneList = ref([])
const nodes = ref([])
const host = ref(null);
const querySerial = ref(null);
const addSerial = ref('');
const manufacturer = ref('');
const model = ref('');
const androidVersion = ref('');
const cpuabi = ref('');
const marketName = ref('');
const marketNameSymbol = ref('');

const serialRules = [
    (v: string) => !!v || '序列号不能为空',
    (v: string) => v.length <= 50 || '序列号不能超过50个字符'
];

const manufacturerRules = [
    (v: string) => !!v || '制造商不能为空',
    (v: string) => v.length <= 50 || '制造商不能超过50个字符'
];

const modelRules = [
    (v: string) => !!v || '型号不能为空',
    (v: string) => v.length <= 50 || '型号不能超过50个字符'
];

const androidVersionRules = [
    (v: string) => !!v || 'Android版本不能为空',
    (v: string) => v.length <= 20 || 'Android版本不能超过20个字符'
];

const cpuabiRules = [
    (v: string) => !!v || 'CPU ABI不能为空',
    (v: string) => v.length <= 20 || 'CPU ABI不能超过20个字符'
];

const marketNameRules = [
    (v: string) => !!v || '名称不能为空',
    (v: string) => v.length <= 50 || '名称不能超过50个字符'
];

const marketNameSymbolRules = [
    (v: string) => !!v || '名称符号不能为空',
    (v: string) => v.length <= 10 || '名称符号不能超过10个字符'
];


watch(host, async (newVal, oldVal) => {
    if (newVal !== oldVal) {
        const res = await getPhoneList(newVal);
        if (res.status) {
            if(res.serials.length == 0) {
                PhoneList.value = [];
                alert('该节点没有手机信息');

            }else {
                PhoneList.value = res.serials
            }
        } else {
            alert('获取节点列表失败: ' + res.message)
        }
        console.log('host 发生变化:', newVal);
    }
});

onMounted(async() => {
    const res:nodesResponse = await getNodes()
    console.log('获取节点列表', res)
    if (res.status) {
        nodes.value = res.nodes
    } else {
        alert('获取节点列表失败: ' + res.message)
    }
})


const handFindPhoneInfo = async () => {
    const data:GetPhoneInfo = {
        host: host.value,
        serial: querySerial.value
    };
    const res = await getPhoneInfo(data);
    ;
    if (res.status) {
        addSerial.value = res.info.serial;
        manufacturer.value = res.info.manufacturer;
        model.value = res.info.model;
        androidVersion.value = res.info.androidVersion;
        cpuabi.value = res.info.cpuabi;
        marketName.value = res.info.marketName;
        marketNameSymbol.value = res.info.marketNameSymbol;
    } else {
        alert('查询手机信息失败: ' + res.message);
    }
}

const handAddPhone = async () => {
    // 实现添加手机的逻辑
    // 校验字段并调用添加接口
}
</script>
