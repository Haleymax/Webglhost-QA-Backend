<template>
    <v-sheet>
        <v-form  ref="formRed">
            <v-container>
                <v-row justify="center" align="center">
                    <v-col cols="12">
                        <v-text-field
                            v-model="host"
                            :rules="hostRules"
                            label="主机地址"
                            placeholder="请输入主机地址"
                            required
                            full-width
                            ></v-text-field>
                    </v-col>
                </v-row>
            </v-container>
        </v-form>
    </v-sheet>
</template>

<script setup lang="ts" name="UpdataNode">
import { ref } from 'vue';


const valid = ref(true)
const formRef = ref()

const host = ref('')
const user = ref('')
const password = ref('')
const name = ref('')

const hostRules = [
    (v: string) => !!v || '主机地址不能为空'
]


</script>