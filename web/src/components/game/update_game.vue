<template>
  <v-sheet>
    <v-form v-model="valid" ref="formRef">
      <v-container>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="id"
              label="ID"
              readonly
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="packageName"
              :rules="packageRules"
              label="包名"
              required
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="type"
              :rules="typeRules"
              label="类型"
              required
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-select
              v-model="caseType"
              :items="caseTypeItems"
              :rules="caseTypeRules"
              label="轮次"
              required
              full-width
            ></v-select>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="gameEngine"
              :rules="engineRules"
              label="游戏引擎"
              required
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="gameUrl"
              :rules="urlRules"
              label="游戏URL"
              required
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="gameName"
              :rules="nameRules"
              label="游戏名称"
              required
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-text-field
              v-model="gameId"
              :rules="idRules"
              label="游戏ID"
              required
              type="number"
              full-width
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row justify="center" align="center">
          <v-col cols="12">
            <v-select
              v-model="status"
              :items="statusItems"
              :rules="statusRules"
              label="状态"
              required
              full-width
            ></v-select>
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

<script setup lang="ts" name="UpdateGame">
import { ref } from 'vue'
import { type Game } from '@/api/game'
import { updateGame } from '@/api/game'

const valid = ref(true)
const formRef = ref(null)
const props = defineProps<{
  gameData: Game
}>()

const id = ref(props.gameData.id || '')
const packageName = ref(props.gameData.package || '')
const type = ref(props.gameData.type || '')
const caseType = ref(props.gameData.case_type || '')
const gameEngine = ref(props.gameData.game_engine || '')
const gameUrl = ref(props.gameData.game_url || '')
const gameName = ref(props.gameData.game_name || '')
const gameId = ref(props.gameData.game_id || '')
const status = ref(props.gameData.status ? '启用' : '禁用')

const caseTypeItems = ['daily', 'build_target', 'converter', 'first_round', 'second_round']
const statusItems = ['启用', '禁用']

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
const statusRules = [(v: string) => !!v || '状态不能为空']

const handleSubmit = async () => {
  if (valid.value) {
    const data: Game = {
      id: id.value,
      package: packageName.value,
      type: type.value,
      case_type: caseType.value,
      game_engine: gameEngine.value,
      game_url: gameUrl.value,
      game_name: gameName.value,
      game_type: 'weixin_minigame',
      game_id: Number(gameId.value),
      status: status.value === '启用'
    }
    const res: any = await updateGame(data)
    if (res.status) {
      alert('更新游戏信息成功: ' + res.message)
    } else {
      alert('更新游戏信息失败: ' + res.message)
    }
  }
}
</script>
