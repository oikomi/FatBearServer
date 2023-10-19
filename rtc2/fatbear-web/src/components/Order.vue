<!-- eslint-disable vue/require-v-for-key -->
<script setup lang="ts">

import Upper from './Upper.vue';
import { ref, inject } from "vue";

import { useTokenStore } from '@/stores/token';

import { userStore } from '@/stores/user';

import { useRouter } from 'vue-router'

const router = useRouter()

const uStore = userStore()

const store = useTokenStore()

const host = uStore.getHost()

interface IModel {
    id: number
    dev_name: string
    model_name: string
    send_user: string
    vibration: string
    duration: number
    token: number
}

const items = ref<IModel[]>([])

const axios: any = inject('axios')  // inject axios

// const SERVER_BASE = "http://127.0.0.1:8080/"
const SERVER_BASE = "https://120.55.60.98/"


const GET_ORDER = SERVER_BASE + "api/v1/dev/order"

axios
    .get(GET_ORDER, { send_user: uStore.getUserName() },{ headers: { 'Token': store.getToken() } }, { withCredentials: true },
    )
    // .get(SERVER_BASE, { headers: { 'Token': store.getToken() } })
    .then((response: { data: any }) => {
        console.log("res data", response.data)
        // if (response.data === 401) {
        //     console.log("get 401, push to login")
        //     router.push({ name: 'login' })
        // }

        items.value = response.data.data
    }).catch((err: any) => {
        console.log("res err", err)
        // if (err.response.status === 401) {
        //     console.log("res err, get 401, push to login")
        //     router.push({ name: 'login' })
        // }

    });


</script>

<template>
    <Upper />

    <div class="container bg-info mb-1">
        <h4 class="left-align mt-2">Recently 100 times action</h4>

        <table class="table mb-2">
            <thead>
                <tr>
                    <th scope="col">#</th>
                    <th scope="col">Time</th>
                    <th scope="col">Model</th>
                    <th scope="col">Duration</th>
                    <th scope="col">Vibration</th>
                    <th scope="col">Tokens</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(set, index) in items" :key='set.id'>
                    <th scope="row">{{ index }}</th>
                    <td>{{ set.duration }} Sec</td>
                    <td>{{ set.vibration }}</td>
                    <td>{{ set.token }}</td>
                </tr>
            </tbody>
        </table>


    </div>
</template>

<style scoped></style>
