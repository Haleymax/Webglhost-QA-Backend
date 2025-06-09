<template>
    <v-card
        title="微信小游戏列表管理"
        flat
    >
        <template v-slot:text>
            <v-row align="center">
                <v-col cols="3">
                    <v-text-field
                        v-model="search"
                        label="游戏名称"
                        prepend-inner-icon="mdi-magnify"
                        variant="outlined"
                        hide-details
                    ></v-text-field>
                </v-col>
                <v-col cols="3">
                    <v-select
                        v-model="gameType"
                        :items="['全部', '单机游戏', '网络游戏']"
                        label="case_type"
                        variant="outlined"
                        hide-details
                    ></v-select>
                </v-col>
                <v-col cols="3">
                    <v-select
                        v-model="gameRound"
                        :items="['第一轮', '第二轮', '第三轮']"
                        label="游戏轮次"
                        variant="outlined"
                        hide-details
                    ></v-select>
                </v-col>
            </v-row>
        </template>
    </v-card>
</template>

<script setup lang="ts" name="WxMiniGame">
import { ref, computed } from 'vue';

const search = ref('');

const gameRound = ref('第一轮');
const gameType = ref('全部');

// 示例游戏数据
const games = ref([
    { name: '游戏1', type: '单机游戏' },
    { name: '游戏2', type: '网络游戏' },
    { name: '游戏3', type: '单机游戏' },
]);

// 过滤后的游戏列表
const filteredGames = computed(() => {
    return games.value.filter(game => {
        const matchesSearch = game.name.includes(search.value);
        const matchesType = gameType.value === '全部' || game.type === gameType.value;
        return matchesSearch && matchesType;
    });
});
</script>