<!-- eslint-disable vue/require-v-for-key -->
<script setup lang="ts">

import Upper from './Upper.vue';
import { ref, inject } from "vue";

import { useTokenStore } from '@/stores/token';

import { userStore } from '@/stores/user';

import { useRouter } from 'vue-router'

import {APP_ID, TOKEN, SERVER_BASE} from '@/config/config';


const router = useRouter()

const uStore = userStore()

const store = useTokenStore()

const host = uStore.getHost()

interface IModel {
    id: number
    createdAt: string
    dev_name: string
    model_name: string
    send_user: string
    vibration: string
    duration: number
    token: number
}

const items = ref<IModel[]>([])

const axios: any = inject('axios')  // inject axios

const GET_ORDER = SERVER_BASE + "api/v1/dev/order"

axios
    .get(GET_ORDER + "?send_user=" + uStore.getUserName(),
    { headers: { 'Token': store.getToken() } }, 
    { withCredentials: true },
    )
    // .get(SERVER_BASE, { headers: { 'Token': store.getToken() } })
    .then((response: { data: any }) => {
        // console.log("res data", response.data)
        // if (response.data === 401) {
        //     console.log("get 401, push to login")
        //     router.push({ name: 'login' })
        // }

        items.value = response.data.data
    }).catch((err: any) => {
        // console.log("res err", err)
        // if (err.response.status === 401) {
        //     console.log("res err, get 401, push to login")
        //     router.push({ name: 'login' })
        // }
    });


function addToken() {

    
}


</script>

<template>
    <Upper />


    <div>
          <div class="bg-info mb-1 mt-1 text-start">
            <h4 class=" mt-2 text-start">Add Tokens</h4>
          </div>


        <div class="container row text-center">
            <h4 class="left-align mt-2 col-4">Token left: </h4> 
            <button type="button" class="btn btn-sm btn-warning col-4" @click="addToken">Apply for 1000 Tokens</button>

        </div>

    </div>


    <div class="container-fluid bg-info mb-1">
        <h4 class="left-align mt-2 text-start">Recently 100 times action</h4>

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
                    <td>{{ set.createdAt }}</td>
                    <td>{{ set.model_name }}</td>
                    <td>{{ set.duration }} Sec</td>
                    <td>{{ set.vibration }}</td>
                    <td>{{ set.token }}</td>
                </tr>
            </tbody>
        </table>


    </div>
</template>

<style scoped></style>
