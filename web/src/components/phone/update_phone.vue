<template>
    <v-sheet>
        <v-form v-model="valid" ref="formRed">
            <v-container>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="serial"
                            :rules="serialRules"
                            label="序列号"
                            required
                            full-width
                            readonly
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="manufacturer"
                            :rules="manufacturerRules"
                            label="制造商"
                            required
                            full-width
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="model"
                            :rules="modelRules"
                            label="型号"
                            required
                            full-width
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="androidVersion"
                            :rules="androidVersionRules"
                            label="Android版本"
                            required
                            full-width
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="cpuabi"
                            :rules="cpuabiRules"
                            label="CPU ABI"
                            required
                            full-width
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="marketName"
                            :rules="marketNameRules"
                            label="名称"
                            required
                            full-width
                        ></v-text-field>
                    </v-col>
                </v-row>

                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="marketNameSymbol"
                            :rules="marketNameSymbolRules"
                            label="名称符号"
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

<script setup lang="ts" name="UpdatePhone">
import { updataPhone, type Phone } from '@/api/phone';
import { ref } from 'vue';

const valid = ref(true);
const props = defineProps<{
    PhoneData: Phone
}>();

const formRef = ref(null);
const id = ref(props.PhoneData.id || '');
const serial = ref(props.PhoneData.serial || '');
const manufacturer = ref(props.PhoneData.manufacturer || '');
const model = ref(props.PhoneData.model || '');
const androidVersion = ref(props.PhoneData.androidVersion || '');
const cpuabi = ref(props.PhoneData.cpuabi || '');
const marketName = ref(props.PhoneData.marketName || '');
const marketNameSymbol = ref(props.PhoneData.marketNameSymbol || '');

const serialRules = [
    (v: string) => !!v || '序列号不能为空'
];
const manufacturerRules = [ 
    (v: string) => !!v || '制造商不能为空'
];
const modelRules = [    
    (v: string) => !!v || '型号不能为空'
];  
const androidVersionRules = [
    (v: string) => !!v || 'Android版本不能为空'
];
const cpuabiRules = [
    (v: string) => !!v || 'CPU ABI不能为空'
];
const marketNameRules = [
    (v: string) => !!v || '名称不能为空'
];
const marketNameSymbolRules = [
    (v: string) => !!v || '名称符号不能为空'
];

const handleSubmit = async () => {
    if (valid.value) {
        const data: Phone = {
            id: id.value,
            serial: serial.value,
            manufacturer: manufacturer.value,
            model: model.value,
            androidVersion: androidVersion.value,
            cpuabi: cpuabi.value,
            marketName: marketName.value,
            marketNameSymbol: marketNameSymbol.value
        };
        
        const res:any = await updataPhone(data);
        if (res.status) {
            alert('更新手机信息成功: ' + res.message);
        } else {
            alert('更新手机信息失败: ' + res.message);
        }
        console.log('表单验证成功', data);
    } else {
        console.log('表单验证失败');
        alert('请检查表单输入');
    }
};

</script>