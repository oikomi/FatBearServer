
import { ref } from 'vue'
import { defineStore } from 'pinia'


export const useTokenStore = defineStore('token', () => {
  const token = ref('')
  function setToken(t: string) {
    token.value = t
  }

  function getToken(): string {
    return token.value
  }


  return { token, getToken, setToken }
})
