


import { ref } from 'vue'
import { defineStore } from 'pinia'


export const userStore = defineStore('user', () => {
    const userName = ref('')
    const host = ref('')

    function setUserName(u : string) {
        userName.value = u
      }
  
      function getUserName() : string {
          return userName.value
      }
  

    function setHost(h : string) {
      host.value = h
    }

    function getHost() : string {
        return host.value
    }
  
    return { userName, host, setUserName, getUserName, getHost, setHost }
  })
  