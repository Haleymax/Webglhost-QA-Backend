<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer">
      <v-list-item title="菜单" subtitle=""></v-list-item>
      <v-divider></v-divider>
      <v-list>
        <v-list-item to="/v/home">
          <v-list-item-icon>
            <v-icon>mdi-home</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>主页</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item to="/v/upload">
          <v-list-item-icon>
            <v-icon>mdi-upload</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>apk上传</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item to="/v/phone">
          <v-list-item-icon>
            <v-icon>mdi-database</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>手机管理</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item to="/v/device">
          <v-list-item-icon>
            <v-icon>mdi-cloud-arrow-up-outline</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>节点管理</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
          <v-list-item to="/v/watcher">
            <v-list-item-icon>
              <v-icon>mdi-eye</v-icon>
            </v-list-item-icon>
            <v-list-item-content>
              <v-list-item-title>watcher管理</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-divider></v-divider>
        <v-list-item to="/v/game">
          <v-list-item-icon>
            <v-icon>mdi-gamepad-variant</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>游戏列表管理</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      <v-divider></v-divider>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-app-bar-title>宿主QA小工具</v-app-bar-title>
    </v-app-bar>

    <v-main>
      <RouterView />
    </v-main>
  </v-app>
</template>

<script setup lang="ts" name="Index">
import { ref } from 'vue'

const drawer = ref(false)
</script>
