<template>
    <v-container>
        <v-row>
            <v-col cols="12">
                <h1>Phone Management</h1>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12">
                <v-card class="pa-4">
                    <v-card-title class="text-h5">Add Phone</v-card-title>
                    <v-form>
                        <v-text-field label="Phone Name" required></v-text-field>
                        <v-text-field label="Phone Model" required></v-text-field>
                        <v-text-field label="Phone Status" required></v-text-field>
                        <v-btn color="primary" type="submit">Submit</v-btn>
                    </v-form>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<script setup lang="ts" name="AddPhone">
import { ref } from 'vue';  
</script>