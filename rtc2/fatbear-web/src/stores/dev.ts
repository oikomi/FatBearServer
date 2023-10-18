
import { ref } from 'vue'
import { defineStore } from 'pinia'


export const useDevStore = defineStore('dev', () => {
  const devName = ref('')

  function setDevName(u: string) {
    devName.value = u
  }

  function getDevName(): string {
    return devName.value
  }


  return { devName, setDevName, getDevName }
})
