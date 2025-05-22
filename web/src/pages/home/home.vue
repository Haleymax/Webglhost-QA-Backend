<template>
    <v-stepper :items="['Step 1', 'Step 2', 'Step 3']">
        <template v-slot:item.1>
            <v-card title="Step One" flat>...</v-card>
        </template>

        <template v-slot:item.2>
            <v-card title="Step Two" flat>...</v-card>
        </template>

        <template v-slot:item.3>
            <v-card title="Step Three" flat>...</v-card>
        </template>
    </v-stepper>
</template>

<script setup lang="ts" name="Home">



const steps = [
    { title: 'Step 1', content: 'Step 1 Content' },
    { title: 'Step 2', content: 'Step 2 Content' },
    { title: 'Step 3', content: 'Step 3 Content' }
]

</script>