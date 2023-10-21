<script setup lang="ts">

// https://github.com/WeiLin-Liao/vue-min-chat/blob/master/src/components/message-pabel.vue

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
    room_name: string
    send_user: string
    msg: string
    createdAt: string
}

const items = ref<IModel[]>([])

let inMsg = ref('')


const axios: any = inject('axios')  // inject axios

// dev
const ROOM_MSG_LIST_URL = SERVER_BASE + "api/v1/room/msg"

function getMsg() {

axios
  .get(ROOM_MSG_LIST_URL + "?name=host", { headers: { 'Token': store.getToken() } }, { withCredentials: true })
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
}

  function sendMsg() {
    axios
    .post(ROOM_MSG_LIST_URL,
      {
        name: "host",
        send_user: uStore.getUserName(),
        msg: inMsg.value
      },

      { headers: { 'Token': store.getToken() } }, { withCredentials: true }

    )
    // .get(SERVER_BASE, { headers: { 'Token': store.getToken() } })
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

    inMsg.value = ''

  }

  getMsg()
  setInterval(getMsg, 500);

</script>

<template>

      <div class="container right-side">
          <!-- <div class="current-chatter">Chat Window</div> -->
          <!-- <button
            id="log-out"
            type="button"
            class="close"
            data-dismiss="alert"
            aria-label="Close"
          > -->
            <!-- <span aria-hidden="true">&times;</span>
          </button> -->
        </div>
        <div class="chat-window container-fluid">
          <div class="system-message ">
            Conversation starts ! Say Hi to others.
          </div>
          <div class="system-message">11:05PM, 7/14/2023</div>
          <div  v-for="(item, index) in items"  :key='item.id + index' class="current-chatter-wrapper">
            <div class="chatter-avatar">
              <img src="/src/assets/SVG/SVG/touxiang.svg" alt="..." />
            </div>
            <div class="chatter-message container">{{item.send_user}} says:  {{item.msg}}</div>
          </div>
        </div>
        <!-- <div class="row mt-2">
          <div
            id="user-input-value"
            class="col-10 user-text"
            contenteditable="true"></div>
          <div class="btn-send">
            <button
              type="button"
              id="btn-send-message"
              class="col-1 btn btn-outline-success">
              Send
            </button>
          </div>
        </div> -->

        <div class="container row mt-2">
        <!-- <div class="col-auto">
            <label for="staticEmail2" class="visually-hidden">Email</label>
            <input type="text" readonly class="form-control-plaintext" id="staticEmail2" value="email@example.com">
        </div> -->
        <div class="col-6">
            <label for="inputMsg" class="visually-hidden secai">Password</label>
            <input type="text" class="form-control" id="inputMsg" placeholder="input msg" v-model="inMsg">
        </div>
        <div class="col-6">
            <button type="submit" class="btn secai fontcss mb-3" @click="sendMsg">Send msg</button>
        </div>

      </div>

</template>

<style scoped>
.app {
  /*display: flex;*/
  display: none;
  width: 100%;
  height: 80%;
  -webkit-box-shadow: -1px -1px 32px -4px rgba(157, 60, 60, 0.64);
  -moz-box-shadow: -1px -1px 32px -4px rgba(166, 80, 80, 0.64);
  box-shadow: -1px -1px 32px -4px rgba(163, 115, 115, 0.64);
  padding: 0;
  border: 1px solid #6b2e43;
  border-radius: 5px;
  position: absolute;
  left: 15%;
  top: 10%;
  /*hide scrollbar in IE*/
  -ms-overflow-style:none;
  /*hide scrollbar in firwfox*/
  overflow-y:-moz-scrollbars-none;
}

.fontcss {

color: azure;
}

.secai {
	background-image: linear-gradient(-45deg, #C24362, #6450A4);
}

.left-side {
  width: 30%;
  background-color: #2b2e36;
  border-right: 1px solid grey;
}
.left-side .user-wrapper {
  width: 100%;
  height: 15%;
  display: flex;
}
.user-wrapper .user-avatar {
  width: 40%;
  height: 100%;
}
.user-avatar img {
  width: 100%;
  height: 100%;
}
.user-wrapper .user-name {
  width: 70%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: white;
}
.other-users-wrapper {
  height: 85%;
  width: 100%;
}
.other-users-wrapper .list-group {
  height: 100%;
  display: flex;
  width: 100%;
}
.other-user-wrapper {
  display: flex;
  height: auto;
  background-color: #2b2e36;
  color: white;
  border: 1px solid black;
  border-radius: 15px;
  -webkit-box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  -moz-box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  font-style: italic;
  transition: all 0.1s;
}
.other-user-wrapper:hover {
  border-radius: 20px;
  border: 3px solid #17c426;
  font-style: normal;
  font-weight: 800;
}
.current-select{
  background-color:rgb(153, 142, 142);
  border: 2px solid black; 
  font-style: normal;
  font-weight: 600;
}
.other-user-wrapper:hover .other-user-avatar {
  width: 45px;
  height: 45px;
}
.other-user-avatar {
  width: 35px;
  height: 35px;
  line-height: 35px;
}
.other-user-avatar img {
  width: 100%;
  height: 100%;
}
.other-user-name {
  width: 80%;
  height: auto;
  display: flex;
  align-items: center;
  justify-content: center;
}

.right-side {
  background-color: rgb(139, 56, 131);
}

.header button:focus{
    outline: 0;
}
.header button:hover{
  color: white;
}
.right-side .header {
  width: 100%;
  height: 15%;
  background-color: #623068;
  color: white;
  font-style: italic;
  font-size: 30px;
  font-weight: 600;
  display: flex;
  align-items: flex-start;
  border-bottom:1px solid grey;
}
.header .current-chatter {
  width: 90%;
  height: 60%;
  display: flex;
  align-items: center;
  justify-content: center;
}
.right-side .chat-window {
  width: 100%;
  height: 60%;
  background-color: #da77ba;
  padding-top: 2%;
  overflow-y: scroll;
}
.chat-window .system-message{
  width: 100%;
  margin-bottom: 2%;
  height: auto;
  text-align: center;
  color: rgb(222, 214, 214);
  font-size: 70%;
}
.chat-window::-webkit-scrollbar {
  width: 10px;
}

.chat-window::-webkit-scrollbar-thumb {
  border-radius: 5px;
  background: rgba(100, 100, 100, 0.7);
}

.chat-window .current-chatter-wrapper {
  width: 70%;
  height: auto;
  margin-bottom: 2%;
  display: flex;
  float: left;
}
.current-chatter-wrapper .chatter-avatar {
  width: 42px;
  height: 42px;

  margin-right: 10px;
}
.chatter-avatar img {
  width: 80%;
  height: 80%;
  float: right;
}
.current-chatter-wrapper .chatter-message {
  width: auto;
  max-width: 80%;
  margin-left: 0%;
  min-height: 42px;
  line-height: 35px;
  height: auto;
  background-color: rgb(211, 170, 193);
  border: 1px solid rgb(172, 171, 171);
  -webkit-box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  -moz-box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
}


.chat-window .me-wrapper {
  width: 70%;
  height: auto;
  margin-bottom: 2%;
  display: flex;
  float: right;
  position: relative;
  left: 1%;
  justify-content: flex-end;
}

.me-wrapper .me-message {
  -webkit-box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  -moz-box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  box-shadow: -1px -1px 32px -4px rgba(0, 0, 0, 0.64);
  width: auto;
  max-width: 80%;
  min-height: 42px;
  height: auto;
  line-height: 35px;
  background-color: gold;
  margin-right: 10px;
  border: 2px solid black;
}
.me-wrapper .me-avatar {
  width: 42px;
  height: 42px;
}
.me-avatar img {
  width: 100%;
  height: 100%;
}
.right-side .user-function {
  padding-top: 2px;
  padding-bottom: 1px;
  width: 100%;
  height: 5%;
  background-color: #2b2e36;
  border-top: 1px solid grey;
  display: flex;
  border-bottom: 1px solid grey;
}

.user-function .emoji {
  height: 100%;
  width: auto;
  line-height: 100%;
  margin-left: 2%;
}

.emoji img {
  height: 23px;
  width: 23px;
  opacity: 0.7;
  filter:alpha(opacity=70);       /* IE */
  -moz-opacity:0.7;              /* old Mozilla */
  -khtml-opacity:0.7;
}
.emoji img:hover{
  opacity: 1;
  filter:alpha(opacity=100);       /* IE */
  -moz-opacity:1;              /* old Mozilla */
  -khtml-opacity:1;
}
.user-function .local-file {
  height: 100%;
  width: auto;
  line-height: 100%;
  margin-left: 3%;
  
}
.local-file img {
  height: 22px;
  width: 22px;
  opacity: 0.7;
  filter:alpha(opacity=70);       /* IE */
  -moz-opacity:0.7;              /* old Mozilla */
  -khtml-opacity:0.7;
  position: relative;
  
}
.local-file img:hover {
  opacity: 1;
  filter:alpha(opacity=100);       /* IE */
  -moz-opacity:1;              /* old Mozilla */
  -khtml-opacity:1;
}
.right-side .user-input {
  width: 100%;
  height: 20%;
  background-color: #3e414a;
}
.user-input .user-text {
  padding: 1% 1% 1% 1%;
  height: 65%;
  color: white;
  outline: none;
  overflow-y: scroll;
}
.user-input .user-text::-webkit-scrollbar {
  width: 10px;
}

.user-input .user-text::-webkit-scrollbar-thumb {
  border-radius: 5px;
  background: rgba(100, 100, 100, 0.7);
}

.user-input .btn {
  height: 2%;
  float: right;
  margin-right: 2%;
  padding: 0.225rem 0.225rem;
}

</style>
