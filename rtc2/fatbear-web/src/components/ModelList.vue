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
    room_name: string
    creator: string
    room_status: number
}

const items = ref<IModel[]>([])

const axios: any = inject('axios')  // inject axios

// const SERVER_BASE = "http://127.0.0.1:8080/"
const SERVER_BASE = "https://120.55.60.98/"


// dev
const ROOM_LIST_URL = SERVER_BASE + "api/v1/room/list"

axios
  .get(ROOM_LIST_URL, { headers: { 'Token': store.getToken() } }, { withCredentials: true })
  // .get(SERVER_BASE, { headers: { 'Token': store.getToken() } })
  .then((response: { data: any }) => {
    console.log("res data", response.data)
    if (response.data === 401) {
      console.log("get 401, push to login")
      router.push({ name: 'login' })
    }

    items.value = response.data.data
  }).catch((err: any) => {
    console.log("res err", err)
    if (err.response.status === 401) {
      console.log("res err, get 401, push to login")
      router.push({ name: 'login' })
    }

  });

async function join() {

  router.push({ name: 'view' })

}


</script>

<template>
  <Upper />
  
<div album py-5 bg-body-tertiary >
  <div class="container row">
    <!-- <div v-for="(room, index) in items"  :key='room.id'> -->
        <div v-for="(room, index) in items"  :key='room.id' class="col-2 card shadow-sm">
            <!-- <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" role="img" aria-label="Placeholder: Thumbnail" preserveAspectRatio="xMidYMid slice" focusable="false"><title>Placeholder</title><rect width="100%" height="100%" fill="#55595c"/>
                <text x="50%" y="50%" fill="#eceeef" dy=".3em">
                Thumbnail</text>
            </svg> -->
            <img src="../assets/zhubo.jpeg" class="img-fluid">
            <div class="card-body">
              <p class="card-text">{{index}}, {{ room.room_name }}, {{ room.creator }}</p>
              <div class="d-flex justify-content-between align-items-center">
                <div class="btn-group">
                  <button type="button" class="btn btn-sm btn-outline-secondary" @click="join">Join</button>
                  <!-- <button type="button" class="btn btn-sm btn-outline-secondary">Edit</button> -->
                </div>
                <!-- <small class="text-body-secondary">9 mins</small> -->
              </div>
            </div>
          </div>
    <!-- </div> -->
  </div>

</div>

</template>

<style scoped>

.bd-placeholder-img {
        font-size: 1.125rem;
        text-anchor: middle;
        -webkit-user-select: none;
        -moz-user-select: none;
        user-select: none;
      }


</style>
