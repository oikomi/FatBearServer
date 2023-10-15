
<script setup lang="ts">
import Upper from './Upper.vue';

import { ref, inject } from "vue";
import { useRouter } from 'vue-router'

const axios: any = inject('axios')  // inject axios

const SERVER_BASE = "http://120.55.60.98:8080/health"

const SERVER_BASE_ONLINE = "https://120.55.60.98/health"


const router = useRouter()

axios
  .get(SERVER_BASE_ONLINE, { withCredentials: true })
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
const APP_ID = "ac1f26c994ea4f978a11ce1251424920";
const TOKEN = "007eJxTYNA5vt2ZzSI2Or9B51p5yP7s0tmPKhT+P0pYdfmqpfZK52cKDInJhmlGZsmWliapiSZpluYWiYaGyamGRqaGJkYmlkYG7j/VUxsCGRk41K8xMjJAIIjPzpCWWJKUmljEwAAAp2kgEA==";


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

    client.on("user-published", onUserPublish);
  }

  await client.join(APP_ID, channel.value, TOKEN, 123);

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

async function startLive() {
  publishAudio()
  publishVideo()
}

</script>


<template>
  <Upper />

  <div class="container row mt-1 mb-2">
    <div class="col-8 bg-info-subtle">
      <div class="row input-group">
        <div class="col-4">
          <div class="input-group mb-3">
            <span class="input-group-text" id="basic-addon1">user</span>
            <input type="text" class="form-control" placeholder="Username" aria-label="Username"
              aria-describedby="basic-addon1">
          </div>

        </div>
        <div class="col-8">
          <button class="btn btn-primary w-45  me-1" @click="startLive">Join</button>
          <button class="btn btn-primary w-45 " @click="leaveChannel">Leave</button>
        </div>
      </div>
      <div class="row">
        <div class="col-8">
          <video class="col-8" v-show="isVideoOn" id="camera-video"></video>
          <video class="col-8" v-show="isVideoSubed" id="remote-video"></video>
          <div v-if="isJoined && !isVideoSubed" class="waiting">
            <!-- You can shared channel {{ channel }} to others..... -->
          </div>
        </div>


        <div class="col-4">
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">Card title</h5>
            </div>
          </div>

          <div class="card">
            <div class="card-body">
              <h5 class="card-title">Card title</h5>
            </div>
          </div>

        </div>

      </div>

    </div>


    <div class="col-4 bg-primary-subtle">

      <nav>
        <div class="nav nav-tabs row" id="nav-tab" role="tablist">
          <button class="nav-link active col" id="nav-setting-tab" data-bs-toggle="tab" data-bs-target="#nav-home"
            type="button" role="tab" aria-controls="nav-home" aria-selected="true">Setting</button>
          <button class="nav-link col" id="nav-chat-tab" data-bs-toggle="tab" data-bs-target="#nav-chat" type="button"
            role="tab" aria-controls="nav-chat" aria-selected="false">Chat</button>
        </div>
      </nav>
      <div class="tab-content" id="nav-tabContent">
        <div class="tab-pane fade show active" id="nav-home" role="tabpanel" aria-labelledby="nav-setting-tab"
          tabindex="0">
          <div class="bg-info mb-1">
            <h4 class="left-align">Connect Toys</h4>
          </div>

          <div class="container">
            <h5 class="mb-2 text-center">VibCrafter APP</h5>
            <div class="input-group mb-3">
              <span class="input-group-text" id="basic-addon1">User ID</span>
              <input type="text" class="form-control" placeholder="user id" aria-label="Username"
                aria-describedby="basic-addon1">
            </div>

            <div class="input-group mb-3">
              <span class="input-group-text" id="basic-addon2">Password</span>
              <input type="text" class="form-control" placeholder="password" aria-label="Recipient's username"
                aria-describedby="basic-addon2">
            </div>

            <div class="col-auto">
              <button type="submit" class="btn btn-primary mb-2 text-center">Login</button>
            </div>

          </div>


        </div>
      </div>
      <div class="tab-pane fade" id="nav-chat" role="tabpanel" aria-labelledby="nav-chat-tab" tabindex="0">
        chat

        <h2 class="left-align">Easemob Chat Examples</h2>

      </div>

    </div>

  </div>
</template>