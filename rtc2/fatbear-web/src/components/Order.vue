<!-- eslint-disable vue/require-v-for-key -->
<script setup lang="ts">

import { ref, inject } from "vue";

import { useTokenStore } from '@/stores/token';

import { userStore } from '@/stores/user';

import { useRouter } from 'vue-router'

import { APP_ID, TOKEN, SERVER_BASE, MSG_INTERVAL } from '@/config/config';
import { dataType } from "element-plus/es/components/table-v2/src/common.mjs";


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

const GET_ORDER = SERVER_BASE + "api/v1/devorder/order"

const ADD_TOKEN = SERVER_BASE + "api/v1/user/addToken"
const GET_TOKEN = SERVER_BASE + "api/v1/user/getToken"


function getOrderF() {
    axios
        .get(GET_ORDER + "?send_user=" + "host",
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

            console.log(items.value)
            console.log("get order success")


            // alert("get token success")

        }).catch((err: any) => {
            // console.log("res err", err)
            // if (err.response.status === 401) {
            //     console.log("res err, get 401, push to login")
            //     router.push({ name: 'login' })
            // }
            alert("get order failed: " + err)
        });

}

let leftToken = ref(0)


function getToken() {

    axios
        .get(GET_TOKEN + "?name=view",
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
            console.log("get token success")
            console.log(response.data)
            leftToken.value = response.data.data
        }).catch((err: any) => {
            // console.log("res err", err)
            // if (err.response.status === 401) {
            //     console.log("res err, get 401, push to login")
            //     router.push({ name: 'login' })
            // }

            alert("get token failed: " + err)
        });

}


// setInterval(getOrderF, 1000);


function addToken() {

    axios
        .post(ADD_TOKEN,
            {
                'name': uStore.getUserName(),
                'token': 1000,
            },
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

            alert("add token success")

        }).catch((err: any) => {
            // console.log("res err", err)
            // if (err.response.status === 401) {
            //     console.log("res err, get 401, push to login")
            //     router.push({ name: 'login' })
            // }

            alert("add token failed: " + err)

        });

}


</script>

<template>
    <div class="secai text-start">
        <div class="row mb-1 mt-2 text-start">
            <img class="col-2 " src="/src/assets/SVG/SVG/jiaqiang.svg" alt="" width="50" height="39">
            <h4 class=" col-10 text-start fontcss">Add Tokens</h4>
        </div>


        <div class="container row text-start mb-2">
            <h4 class="left-align mt-2 col-4 fontcss">Token left: {{leftToken}}</h4>
            <button class="col-2 btn btn-primary mt-1 mb-2 secai" @click="getToken">Refresh</button>
            <button type="button" class="btn offset-1 col-4 xiaotubiaored fontcss" @click="getToken">Apply for 1000
                Tokens</button>
        </div>

    </div>


    <div class="container-fluid secai text-start mb-1">
        <div class="container-fluid row">
            <h4 class="left-align col-4 mt-2 text-start secai fontcss">Recently 100 times action</h4>
            <button class="offset-1 col-2 btn btn-primary mt-1 mb-2 secai" @click="getOrderF">Refresh</button>
        </div>

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

<style scoped>
html,
body {
    width: 100%;
    height: 100%;
    /* background-color:#C24362; */
    background-image: linear-gradient(-45deg, #C24362, #6450A4);
    /* background-image: url(/src/assets/SVG/chunbeijing/meinv.svg); */
    /* background-repeat: no-repeat; */
}


.secai {
    background-image: linear-gradient(-45deg, #C24362, #6450A4);
}

.maincss {
    background-image: url(/src/assets/SVG/chunbeijing/meinv.svg);

}

.setting {
    background-image: linear-gradient(-45deg, #ba4068, #8474b1);
}

.dandiv {
    background-image: linear-gradient(-45deg, #c95b7c, #7252c4);
}

.tablecss {
    /* background-image: linear-gradient(-45deg, #753c4d, #594b7e); */
    background-color: rgba(0, 0, 0, 0);

}

.xiaotubiao {
    background-image: linear-gradient(-45deg, #1c181c, #8675b5);
    /* background-color:rgba(20, 19, 19, 0); */

}

.xiaotubiaored {
    background-image: linear-gradient(-45deg, #e80f45, #8d8898);
    /* background-color:rgba(20, 19, 19, 0); */

}



.fontcss {

    color: azure;
}
</style>
