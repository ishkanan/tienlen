<script setup lang="ts">
import { computed, inject, ref } from 'vue'
import ConfirmDialog from '../components/ConfirmDialog.vue'
import InputDialog from '../components/InputDialog.vue'
import { useGameStore } from '../stores/game'

interface Socket {
  requestChangeName: ({ name }: { name: string }) => void
  requestResetGame: () => void
}

const socket = inject<Socket>('socket')

const gameStore = useGameStore()

const playerName = computed(() => gameStore.self?.name ?? '')

const showInput = ref(false)

const doNameInput = () => (showInput.value = true)

const handleNameInput = (confirm: boolean, name: string) => {
  if (confirm) socket?.requestChangeName({ name })
  showInput.value = false
}

const showConfirm = ref(false)

const doResetConfirm = () => (showConfirm.value = true)

const handleConfirm = (confirm: boolean) => {
  if (confirm) socket?.requestResetGame()
  showConfirm.value = false
}
</script>

<template>
  <div class="viewport">
    <button class="button is-info" @click="doNameInput">Change name</button>
    <button class="button is-danger" @click="doResetConfirm">Reset game</button>
    <InputDialog
      v-if="showInput"
      title="Enter new name"
      :default="playerName"
      @confirm="handleNameInput"
    />
    <ConfirmDialog
      v-if="showConfirm"
      title="Are you sure?"
      message="This will reset EVERYTHING and remove all disconnected players!"
      @confirm="handleConfirm"
    />
  </div>
</template>

<style scoped>
.viewport {
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: space-evenly;
}

.button {
  width: 50%;
}
</style>
