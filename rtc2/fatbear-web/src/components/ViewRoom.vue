
<script setup lang="ts">
import Upper from './Upper.vue';

import { ref, inject } from "vue";
import { useRouter } from 'vue-router'

import { useTokenStore } from '@/stores/token';

import { userStore } from '@/stores/user';

import { useDevStore } from '@/stores/dev';


const uStore = userStore()

const store = useTokenStore()

const devStore = useDevStore()

const axios: any = inject('axios')  // inject axios

axios.defaults.withCredentials = true

// const SERVER_BASE = "http://127.0.0.1:8080/"
const SERVER_BASE = "https://120.55.60.98/"

const SEND_TIP_URL = SERVER_BASE + "api/v1/dev/order"
const DEV_SET_URL = SERVER_BASE + "api/v1/dev/set"


const router = useRouter()

interface IModel {
  id: number
  model_name: string
  vibration: number
  duration: number
  token: number
}

const items = ref<IModel[]>([])

axios
  .get(DEV_SET_URL, { headers: { 'Token': store.getToken() } }, { withCredentials: true }
    ,
    { "model_name": "host" }
  )
  // .get(SERVER_BASE + "health", { headers: { 'Token': store.getToken() } })
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


// axios
//   .get(SERVER_BASE + "health", { headers: { 'Token': store.getToken() } }, { withCredentials: true })
//   // .get(SERVER_BASE + "health", { headers: { 'Token': store.getToken() } })
//   .then((response: { data: any }) => {
//     console.log("res data", response.data)
//     if (response.data === 401) {
//       console.log("get 401, push to login")
//       router.push({ name: 'login' })
//     }
//   }).catch((err: any) => {
//     console.log("res err", err)
//     if (err.response.status === 401) {
//       console.log("res err, get 401, push to login")
//       router.push({ name: 'login' })
//     }
// });

function sendTip() {

  console.log("send tip , token is  ", store.getToken())
  axios
    .post(SEND_TIP_URL, 
    {
        'dev_name': devStore.getDevName(),
        'model_name': "host",
        'send_user': uStore.getUserName(),
        'vibration': "Medium",
        'duration': 10,
        'token': 100
      },
      { headers: { 'Token': store.getToken() } }, { withCredentials: true }
    )
    // .get(SERVER_BASE + "health", { headers: { 'Token': store.getToken() } })
    .then((response: { data: any }) => {
      console.log("res data", response.data)
      if (response.data === 401) {
        console.log("get 401, push to login")
        router.push({ name: 'login' })
      }
    }).catch((err: any) => {
      console.log("res err", err)
      if (err.response.status === 401) {
        console.log("res err, get 401, push to login")
        router.push({ name: 'login' })
      }
    });



}

import type {
  IAgoraRTCClient,
  IAgoraRTCRemoteUser,
  ICameraVideoTrack,
  IMicrophoneAudioTrack,
} from "agora-rtc-sdk-ng";

import AgoraRTC from "agora-rtc-sdk-ng";

let track: ICameraVideoTrack;
let audioTrack: IMicrophoneAudioTrack;
let client: IAgoraRTCClient;

let isJoined = ref(false);
let isAudioPubed = ref(false);
let isVideoPubed = ref(false);
let isAudioSubed = ref(false);
let isVideoSubed = ref(false);

let isVideoOn = ref(false);
let isAudioOn = ref(false);

let roomName = ref('');

async function turnOnCamera() {
  isVideoOn.value = !isVideoOn.value;

  if (track) {
    return track.setEnabled(!track.enabled);
  }
  track = await AgoraRTC.createCameraVideoTrack();
  track.play("camera-video");
}

async function turnOnMicrophone() {
  isAudioOn.value = !isAudioOn.value;

  if (audioTrack) {
    return audioTrack.setEnabled(!audioTrack.enabled);
  }

  audioTrack = await AgoraRTC.createMicrophoneAudioTrack();
  audioTrack.play();
}

const channel = ref("fatbear");
// you can apply appid follow the guide https://www.agora.io/en/blog/how-to-get-started-with-agora/
const APP_ID = "ac1f26c994ea4f978a11ce1251424920";
const TOKEN = "007eJxTYIjz9PMVdWl9PydrlkXdhspzhRel7q/XZSpyTry5vu2e1TsFhsRkwzQjs2RLS5PURJM0S3OLREPD5FRDI1NDEyMTSyODI0Z6qQ2BjAwsZmysjAwQCOKzM6QlliSlJhYxMAAAZEkejw==";


async function joinChannel() {
  if (!channel.value) {
    channel.value = "fatbear";
  }

  if (isJoined.value) {
    await leaveChannel();
  }

  if (!client) {
    client = AgoraRTC.createClient({
      mode: "rtc",
      codec: "vp8",
    });

    // client.setClientRole("audience", {
    //   level: 2
    // });

    client.on("user-published", onUserPublish);

    // await client.join(APP_ID, channel.value, TOKEN, "1234567890");

    // isJoined.value = true;

  }

  await client.join(APP_ID, channel.value, TOKEN, 1234567890);

  isJoined.value = true;
}

async function leaveChannel() {
  isJoined.value = false;
  isAudioPubed.value = false;
  isVideoPubed.value = false;

  // add
  isVideoOn.value = false

  client && (await client.leave());
}

async function onUserPublish(
  user: IAgoraRTCRemoteUser,
  mediaType: "video" | "audio"
) {
  if (mediaType === "video") {
    const remoteTrack = await client.subscribe(user, mediaType);
    remoteTrack.play("remote-video");
    isVideoSubed.value = true;
  }
  if (mediaType === "audio") {
    const remoteTrack = await client.subscribe(user, mediaType);
    remoteTrack.play();
    isAudioSubed.value = true;
  }
}

async function joinLive() {
  if (!isVideoOn.value) {
    await turnOnCamera();
  }

  if (!isJoined.value) {
    await joinChannel();
  }
  // await client.publish(track);
  // isVideoPubed.value = true;
}


async function publishVideo() {
  if (!isVideoOn.value) {
    await turnOnCamera();
  }

  if (!isJoined.value) {
    await joinChannel();
  }
  await client.publish(track);
  isVideoPubed.value = true;
}

async function publishAudio() {
  if (!isAudioOn.value) {
    await turnOnMicrophone();
  }
  if (!isJoined.value) {
    await joinChannel();
  }

  await client.publish(audioTrack);
  isAudioPubed.value = true;
}

</script>


<template>
  <Upper />

  <svg xmlns="http://www.w3.org/2000/svg" class="d-none">
    <symbol id="arrow-right-short" viewBox="0 0 16 16">
      <path fill-rule="evenodd"
        d="M4 8a.5.5 0 0 1 .5-.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5A.5.5 0 0 1 4 8z" />
    </symbol>
    <symbol id="x-lg" viewBox="0 0 16 16">
      <path fill-rule="evenodd"
        d="M13.854 2.146a.5.5 0 0 1 0 .708l-11 11a.5.5 0 0 1-.708-.708l11-11a.5.5 0 0 1 .708 0Z" />
      <path fill-rule="evenodd" d="M2.146 2.146a.5.5 0 0 0 0 .708l11 11a.5.5 0 0 0 .708-.708l-11-11a.5.5 0 0 0-.708 0Z" />
    </symbol>
  </svg>


  <div class="container row mt-1 mb-2">
    <div class="col-8 bg-info-subtle">
      <div class="row input-group">
        <div class="col-2">
          <!-- <div class="input-group mb-3"> -->
          <!-- <span class="input-group-text" id="basic-addon1">room</span>
            <input type="text" class="form-control" placeholder="room" aria-label="room" aria-describedby="basic-addon1"
              v-model="roomName"> -->
          <!-- </div> -->

        </div>
        <div class="col-8">
          <button class="btn btn-primary w-45  me-1" @click="joinLive">Join</button>
          <!-- <button class="btn btn-primary w-45 " @click="leaveChannel">Leave</button> -->
        </div>
      </div>
      <div class="row">
        <div class="col-8">

          <video class="col-8" v-show="isJoined" id="remote-video">remote</video>

          <div v-if="isJoined" class="waiting">

          </div>

          <!-- <video class="col-8" v-show="isVideoOn" id="camera-video">local</video> -->
          <!-- <video class="col-8" v-show="isVideoOn" id="camera-video"></video>
          <video class="col-8" v-show="isVideoSubed" id="remote-video"></video>
          <div v-if="isJoined && !isVideoSubed" class="waiting"> -->
          <!-- You can shared channel {{ channel }} to others..... -->
        </div>
      </div>
    </div>
    <!-- </div> -->



    <div class="col-3 bg-primary-subtle">

      <nav>
        <div class="nav nav-tabs row" id="nav-tab" role="tablist">
          <button class="nav-link active col" id="nav-setting-tab" data-bs-toggle="tab" data-bs-target="#nav-home"
            type="button" role="tab" aria-controls="nav-home" aria-selected="true">Send</button>
          <button class="nav-link col" id="nav-chat-tab" data-bs-toggle="tab" data-bs-target="#nav-chat" type="button"
            role="tab" aria-controls="nav-chat" aria-selected="false">Chat</button>
        </div>
      </nav>
      <div class="tab-content" id="nav-tabContent">
        <div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-setting-tab"
          tabindex="0">

          <h4 class="left-align mt-1">Tip Menu</h4>

          <div>
            <table class="table mb-2">
              <thead>
                <tr>
                  <th scope="col">#</th>
                  <th scope="col">Activity</th>
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

          <div>
            User ID: {{ uStore.getUserName() }} Token Left
          </div>


          <div class="container mt-2 text-center">
            <div class=" btn-group text-center" role="group" aria-label="Basic radio toggle button group">
              <input type="radio" class=" btn-check btn-rounded" name="btnradio" id="btnradio1" autocomplete="off"
                checked>
              <label class="btn btn-outline-primary" for="btnradio1">20</label>

              <input type="radio" class=" btn-check btn-rounded" name="btnradio" id="btnradio2" autocomplete="off">
              <label class="btn btn-outline-primary" for="btnradio2">50</label>

              <input type="radio" class=" btn-check btn-rounded" name="btnradio" id="btnradio3" autocomplete="off">
              <label class="btn btn-outline-primary" for="btnradio3">100</label>
            </div>

          </div>

          <div class="text-center mt-2">

            <div class="input-group mb-3">
              <span class="input-group-text " id="basic-addon1">Custom Amount</span>
              <input type="text" class="form-control" placeholder="" aria-label="room" aria-describedby="basic-addon1"
                v-model="roomName">
            </div>

            <div class="text-center mt-2">

              <button class="btn btn-primary d-inline-flex align-items-center" type="button" @click="sendTip">
                Send
                <svg class="bi ms-1" width="20" height="20">
                  <use xlink:href="#arrow-right-short" />
                </svg>
              </button>
            </div>


          </div>


        </div>
      </div>


      <div class="tab-pane fade" id="nav-chat" role="tabpanel" aria-labelledby="nav-chat-tab" tabindex="0">
        chat

        <h2 class="left-align">Easemob Chat Examples</h2>

      </div>

    </div>

</div></template>