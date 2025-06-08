<template>
    <v-stepper v-model="activeStep" :items="stepTitles">
        <template v-slot:item.1>
            <v-card :title="steps[0].title" flat>
                {{ steps[0].content }}
            </v-card>
        </template>

        <template v-slot:item.2>
            <v-card :title="steps[1].title" flat>
                {{ steps[1].content }}
            </v-card>
        </template>

        <template v-slot:item.3>
            <v-card :title="steps[2].title" flat>
                {{ steps[2].content }}
            </v-card>
        </template>

        <template v-slot:item.4>
            <v-card :title="steps[3].title" flat>
                {{ steps[3].content }}
            </v-card>
        </template>
    </v-stepper>
</template>

<script setup lang="ts" name="Home">
import { ref, computed } from 'vue'

const steps = [
    { title: 'apk上传', content: '选择对应的远程节点，以及选择对应的apk包，将对应的apk包上传至指定的远程节点上' },
    { title: '手机管理', content: '实现对手机信息的管理，可以添加远程主机上的手机序列号进行入库，也可以实现对现有数据的修改' },
    { title: '节点管理', content: '可以实现对远程节点信息的管理，以方便需要用到远程操作的时候选择节点' },
    { title: 'watcher管理', content: '可以实现对watcher数据的管理，以及添加新的watcher和将数据库中的watcher刷入redis缓存中' }
]

const stepTitles = computed(() => steps.map(step => step.title))
const activeStep = ref(1)
</script>