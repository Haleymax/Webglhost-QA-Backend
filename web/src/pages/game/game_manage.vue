<template>
  <v-card class="pa-4" outlined>
    <v-tabs
      v-model="tab"
      bg-color="teal-darken-3"
      color="white"
      @change="onTabChange"
    >
      <v-tab value="one">tap游戏列表</v-tab>
      <v-tab value="two">rpk游戏列表</v-tab>
      <v-tab value="three">刷库</v-tab>
      <v-tab value="four">websocket</v-tab>
    </v-tabs>

    <v-card-text>
      <v-tabs-window v-model="tab">
        <v-tabs-window-item value="one">
            <WxMiniGame/>
        </v-tabs-window-item>

        <v-tabs-window-item value="two">
            <QuickMiniGame/>
        </v-tabs-window-item>

        <v-tabs-window-item value="three">
            <DatabasePage/>
        </v-tabs-window-item>

        <v-tabs-window-item value="four">
            <WebSocket/>
        </v-tabs-window-item>
      </v-tabs-window>
    </v-card-text>
  </v-card>
</template>

<script setup lang="ts" name="GameManage">
import { ref } from 'vue';
import WxMiniGame from '@/pages/game/wxmini_game.vue';
import QuickMiniGame from '@/pages/game/quick_game.vue'
import DatabasePage from '@/pages/game/database.vue';
import WebSocket from '@/pages/game/websocket.vue'

const tab = ref('one');

function onTabChange(newTab: string) {
  console.log(`当前选中的标签是: ${newTab}`);
}

</script>
