<template>
  <v-container>
    <v-row>
      <v-col cols="12">
        <h1>添加微信小游戏</h1>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12">
        <v-card class="pa-4">
          <v-card-title class="text-h5">添加游戏</v-card-title>
          <v-form @submit.prevent="handleAddGame" ref="formRef" v-model="valid">
            <v-text-field
              label="package"
              v-model="packageName"
              :rules="packageRules"
              placeholder="请输入包名"
              required
            ></v-text-field>
            <v-select
              label="case_type"
              v-model="caseType"
              :items="caseTypeItems"
              :rules="caseTypeRules"
              placeholder="请选择轮次"
              required
            ></v-select>
            <v-text-field
              label="game_engine"
              v-model="gameEngine"
              :rules="engineRules"
              placeholder="请输入游戏引擎"
              required
            ></v-text-field>
            <v-text-field
              label="game_url"
              v-model="gameUrl"
              :rules="urlRules"
              placeholder="请输入游戏URL"
              required
            ></v-text-field>
            <v-text-field
              label="game_name"
              v-model="gameName"
              :rules="nameRules"
              placeholder="请输入游戏名称"
              required
            ></v-text-field>
            <v-text-field
              label="game_id"
              v-model="gameId"
              :rules="idRules"
              placeholder="请输入游戏ID"
              required
              type="number"
            ></v-text-field>
            <v-btn color="primary" type="submit">添加</v-btn>
          </v-form>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts" name="AddGame">
import { ref } from 'vue'
import { addGame, type Game } from '@/api/game'

const valid = ref(true)
const formRef = ref()

const packageName = ref('')
const type = ref('')
const caseType = ref('')
const gameEngine = ref('')
const gameUrl = ref('')
const gameName = ref('')
const gameId = ref('')

const caseTypeItems = ['daily', 'build_target', 'converter', 'first_round', 'second_round']

const packageRules = [(v: string) => !!v || '包名不能为空']
const typeRules = [(v: string) => !!v || '类型不能为空']
const caseTypeRules = [(v: string) => !!v || '轮次不能为空']
const engineRules = [(v: string) => !!v || '引擎不能为空']
const urlRules = [(v: string) => !!v || 'URL不能为空']
const nameRules = [(v: string) => !!v || '名称不能为空']
const idRules = [
  (v: string) => !!v || 'ID不能为空',
  (v: string) => /^\d+$/.test(v) || 'ID必须为数字'
]

const handleAddGame = async () => {
  const isValid = await formRef.value?.validate?.()
  if (!isValid) return

  const data:Game = {
    package: packageName.value,
    type: "unity 引擎小游戏",
    case_type: [caseType.value],
    game_engine: gameEngine.value,
    game_url: gameUrl.value,
    game_name: gameName.value,
    game_type: 'weixin_minigame',
    game_id: Number(gameId.value),
    status: true,
    id:"",
  }
  const res: any = await addGame(data)
  if (res.status) {
    alert('添加游戏成功: ' + res.message)
    // 清空表单
    packageName.value = ''
    type.value = ''
    caseType.value = ''
    gameEngine.value = ''
    gameUrl.value = ''
    gameName.value = ''
    gameId.value = ''
  } else {
    alert('添加游戏失败: ' + res.message)
  }
}
</script>
