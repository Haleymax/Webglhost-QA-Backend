<template>
    <v-card
        title="微信小游戏列表管理"
        flat
    >
        <template v-slot:text>
            <v-row align="center" class="mb-4" no-gutters>
                <v-col cols="3" class="pr-2">
                    <v-text-field
                        v-model="searchName"
                        label="游戏名称"
                        prepend-inner-icon="mdi-magnify"
                        variant="outlined"
                        hide-details
                        dense
                    ></v-text-field>
                </v-col>
                <v-col cols="2" class="pr-2">
                    <v-select
                        v-model="case_type"
                        :items="['daily', 'build_target', 'converter', 'first_round', 'second_round']"
                        label="游戏轮次"
                        variant="outlined"
                        hide-details
                        dense
                    ></v-select>
                </v-col>
                <v-col cols="5" class="d-flex justify-end align-center">
                    <v-btn
                        color="primary"
                        class="mx-1"
                        @click="handleSearch"
                        variant="elevated"
                    >
                        查找游戏
                    </v-btn>
                    <v-btn
                        color="success"
                        class="mx-1"
                        @click="addgame = true"
                        variant="elevated"
                    >
                        添加游戏
                    </v-btn>
                    <v-dialog
                        v-model="addgame"
                        max-width="600px"
                    >
                        <v-card>
                            <v-card-title>添加微信小游戏</v-card-title>
                            <AddGame />
                        </v-card>
                        </v-dialog>
                    <v-btn
                        color="info"
                        class="mx-1"
                        @click="refreshGames"
                        variant="elevated"
                    >
                        刷新
                    </v-btn>
                </v-col>
            </v-row>
        </template>
        <v-data-table
            :headers="headers"
            :items="games"
        >
            <template v-slot:item.action="{ item }">
                <v-btn
                    icon
                    @click="() => console.log(item.name)"
                >
                    <v-icon>mdi-delete</v-icon>
                </v-btn>
                <v-btn
                    icon
                    @click="onEditGame(item)"
                >
                    <v-icon>mdi-pencil</v-icon>
                </v-btn>
            </template>
        </v-data-table>
        <v-dialog
            v-model="updategame"
            max-width="600px">
            <v-card>
                <v-card-title>编辑微信小游戏</v-card-title>
                <UpdateGame v-if="editGame":gameData="editGame" />
            </v-card>
        </v-dialog>
    </v-card>
</template>

<script setup lang="ts" name="WxMiniGame">
import { ref, onMounted } from 'vue';
import { findAllWxGame, searchGame, type SearchGame, type Game } from '@/api/game';
import AddGame from '@/components/game/add_game.vue';
import UpdateGame from '@/components/game/update_game.vue';

const searchName = ref('');
const case_type = ref('daily');
const games = ref([]);
const loading = ref(false);
const addgame = ref(false);
const updategame = ref(false);
const editGame = ref<Game | null>(null);

const headers = ref([
    { text: 'package', value: 'package'},
    { text: 'type', value: 'type'},
    { text: 'case_type', value: 'case_type', sortable: true},
    { text: 'game_engine', value: 'game_engin', sortable: true },
    { text: 'game_url', value: 'game_url'},
    { text: 'game_name', value: 'game_name'},
    { text: 'game_id', value: 'game_id', sortable: true },
    { text: 'action', value: 'action', sortable: false }
]);

onMounted(async() => {
    await refreshGames();
});

const refreshGames = async () => {
    if (loading.value) return;
    loading.value = true;
    try {
        console.log('刷新游戏列表');
        searchName.value = '';
        case_type.value = 'daily'; // 恢复默认值
        const res:any = await findAllWxGame();
        if (res.status) {
            games.value = res.games;
        } else {
            console.error('刷新游戏列表失败:', res.message);
        }
    } finally {
        loading.value = false;
    }
};

const handleSearch = async() => {
    if (loading.value) return;
    loading.value = true;
    try {
        console.log('搜索游戏:', searchName.value, '类型:', case_type.value);
        const SearchDate:SearchGame = {
            game_name: searchName.value,
            case_type: case_type.value,
            game_type: "weixin_minigame"
        };
        const res:any = await searchGame(SearchDate);
        if (res.status) {
            games.value = res.games;
            if (games.value === null) {
                console.warn('没有找到匹配的游戏');
                games.value = [];
            }
            console.log('搜索结果:', games.value);
        } else {
            console.error('搜索游戏失败:', res.message);
        }
    } finally {
        loading.value = false;
    }
};

const onEditGame = (item: Game) => {
    editGame.value = item;
    updategame.value = true;
};
</script>
