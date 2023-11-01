
<script setup lang="ts">
// import Upper from './Upper.vue';

import Order from '@/components/Order.vue'
import Chat from './Chat.vue';


import { ref, inject, onMounted } from "vue";
import { useRouter } from 'vue-router'

import { useTokenStore } from '@/stores/token';

import { userStore } from '@/stores/user';

import { useDevStore } from '@/stores/dev';
import { APP_ID, TOKEN, SERVER_BASE } from '@/config/config';


const uStore = userStore()

const store = useTokenStore()

const devStore = useDevStore()

const axios: any = inject('axios')  // inject axios

axios.defaults.withCredentials = true

const SEND_TIP_URL = SERVER_BASE + "api/v1/devorder/order"
const DEV_SET_URL = SERVER_BASE + "api/v1/dev/set"

const GET_TOKEN_URL = SERVER_BASE + "api/v1/user/getToken"


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
  .get(DEV_SET_URL + "?model_name=" + "host",

    { headers: { 'Token': store.getToken() } },
    { withCredentials: true },

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

  let userToken = ref(0);

  function getUserToken() {

    axios
  .get(GET_TOKEN_URL + "?name=" + uStore.getUserName(),
    { headers: { 'Token': store.getToken() } },
    { withCredentials: true },
  )
  // .get(SERVER_BASE + "health", { headers: { 'Token': store.getToken() } })
  .then((response: { data: any }) => {
    console.log("res data", response.data)
    if (response.data === 401) {
      console.log("get 401, push to login")
    }
    userToken.value = response.data.data
  }).catch((err: any) => {
    console.log("res err", err)
    if (err.response.status === 401) {
      console.log("res err, get 401, push to login")
      router.push({ name: 'login' })
    }
  });
  }

  onMounted(() => {
  console.log(`the component is now mounted.`)
  getUserToken()

})

  // setInterval(getUserToken, 1000);


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


const tselect = ref('')
const tc = ref(0)
let lastToken: number = 0

function setting() {
  router.push({ name: 'order' })
}

function sendTip() {

  console.log("tselect : ", tselect.value)


  if (tselect.value === '20') {
    lastToken = 20
  } else if (tselect.value === '50') {
    lastToken = 50
  } else if (tselect.value === '100') {
    lastToken = 100
  }

  if (tc.value !== 0) {
    lastToken = tc.value
  }

  console.log("lastToken is ", lastToken)

  console.log("send tip , token is  ", store.getToken())
  axios
    .post(SEND_TIP_URL,
      {
        'dev_name': devStore.getDevName(),
        'model_name': "host",
        'send_user': uStore.getUserName(),
        'vibration': "Medium",
        'duration': 2,
        'token': lastToken,
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
      alert("send success")
    }).catch((err: any) => {
      console.log("res err", err)
      if (err.response.status === 401) {
        console.log("res err, get 401, push to login")
        router.push({ name: 'login' })
      }
      alert("send failed")

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
// const APP_ID = "ac1f26c994ea4f978a11ce1251424920";
// const TOKEN = "007eJxTYPh4TveowQ+tuAPJ8Y9epC/ldvLSOrj6dSq/5LMkR3u90CoFhsRkwzQjs2RLS5PURJM0S3OLREPD5FRDI1NDEyMTSyODY7cNUhsCGRnyXP8xMEIhiM/OkJZYkpSaWMTAAACkgSAs";

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

  console.log("client join the live channle ok")

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

  <body>

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


  <div class="container-fluid row mt-1 mb-2">
    <div class="col-6 ">
      <div class="row input-group">
        <div class="col-2">
          <!-- <div class="input-group mb-3"> -->
          <!-- <span class="input-group-text" id="basic-addon1">room</span>
            <input type="text" class="form-control" placeholder="room" aria-label="room" aria-describedby="basic-addon1"
              v-model="roomName"> -->
          <!-- </div> -->

        </div>
        <div class="col-4">
          <button class="btn btn-primary w-45  me-1" @click="joinLive">Join</button>
          <!-- <button class="btn btn-primary w-45 " @click="leaveChannel">Leave</button> -->
        </div>
      </div>
      <div class="container-fluid">
        <!-- <div class="col-8"> -->
        <div class="container-fluid row col-2 xiaotubiaored rounded">
          <h6 class=" fontcss mt-1 ">Live</h6>
        </div>


        <video class="col-12 object-fit: fill" v-show="isJoined" id="remote-video">remote</video>

        <!-- <div v-if="isJoined" class="waiting">

          </div> -->

        <!-- <video class="col-8" v-show="isVideoOn" id="camera-video">local</video> -->
        <!-- <video class="col-8" v-show="isVideoOn" id="camera-video"></video>
          <video class="col-8" v-show="isVideoSubed" id="remote-video"></video>
          <div v-if="isJoined && !isVideoSubed" class="waiting"> -->
        <!-- You can shared channel {{ channel }} to others..... -->
        <!-- </div> -->
      </div>
    </div>
    <!-- </div> -->

    <div class="col-6 setting">

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

          <h4 class="left-align mt-1 fontcss">Tip Menu</h4>

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

          <div class="dandivfushang">

          <div class="mt-2 ">
            <h5 class="fontcss text-center">User ID: {{ uStore.getUserName() }}, Token Left: {{userToken}}</h5>
          </div>


          <div class="container mt-2  text-center">
            <div class="form-check btn-group text-center" role="group" aria-label="Basic radio toggle button group">
              <input type="radio" class="form-check-input btn-check  " name="btnradio" id="btnradio1" autocomplete="off"
                checked value="20" v-model="tselect">
              <label class="btn btn-outline-primary me-4" for="btnradio1">20</label>

              <input type="radio" class="form-check-input btn-check " name="btnradio" id="btnradio2" autocomplete="off"
                value="50" v-model="tselect">
              <label class="btn btn-outline-primary me-4" for="btnradio2">50</label>

              <input type="radio" class="form-check-input btn-check " name="btnradio" id="btnradio3" autocomplete="off"
                value="100" v-model="tselect">
              <label class="btn btn-outline-primary me-4" for="btnradio3">100</label>
            </div>

          </div>

          <div class="text-center mt-2">

            <div class="container row ">
              <div class="input-group  col-8 mb-3" width="20" height="20">
                <span class="input-group-text" id="basic-addon1">Custom Amount</span>
                <input type="text" class="form-control" placeholder="" aria-label="room" aria-describedby="basic-addon1"
                  v-model="tc">
              </div>
            </div>

            <div class="row text-center mt-2">

              <button class="offset-4 col-2 btn  d-inline-flex align-items-center secai" type="button" @click="sendTip">
                Send
                <svg class="bi ms-1" width="20" height="20">
                  <use xlink:href="#arrow-right-short" />
                </svg>
              </button>


              <button type="button" class="col-2 btn  ms-2 secai" data-bs-toggle="modal" data-bs-target="#exampleModal">
                Setting
              </button>
            </div>


            </div>

            <!-- <button type="button" class="btn btn-sm btn-outline-secondary" @click="setting">setting</button> -->

            <div>

              <!-- Button trigger modal -->

              <!-- Modal -->
              <div class="modal fade modal-xl" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel"
                aria-hidden="true">
                <div class="modal-dialog">
                  <div class="modal-content">
                    <div class="modal-header">
                      <h1 class="modal-title fs-5 " id="exampleModalLabel">Setting</h1>
                      <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body ">
                      
                      <Order />

                    </div>
                    <div class="modal-footer">
                      <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
                      <!-- <button type="button" class="btn btn-primary">Save changes</button> -->
                    </div>
                  </div>
                </div>
              </div>

            </div>

          </div>



        </div>
      </div>


      <div class="tab-pane fade" id="nav-chat" role="tabpanel" aria-labelledby="nav-chat-tab" tabindex="0">

        <Chat />

      </div>

    </div>

  </div>

</body>

</template>


<style scoped>

html,
body {
	width : 100%;
  height : 100%;
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
  background-color:rgba(0,0,0,0);

}

.xiaotubiao {
  background-image: linear-gradient(-45deg, #1c181c, #8675b5);
  /* background-color:rgba(20, 19, 19, 0); */

}

.dandivfushang {
  background-image: linear-gradient(-45deg, #c95b7c, #4d4b52);
}


.xiaotubiaored {
  background-image: linear-gradient(-45deg, #e80f45, #8d8898);
  /* background-color:rgba(20, 19, 19, 0); */

}



.fontcss {

  color: azure;
}

</style>
