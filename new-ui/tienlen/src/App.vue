<script setup lang="ts">
import { computed } from 'vue'
import Game from './views/Game.vue'
import Intro from './views/Intro.vue'
import Toaster from './views/Toaster.vue'
import { useGameStore, ConnectionState } from './stores/game'

const gameStore = useGameStore()

const connected = computed(() => gameStore.connState === ConnectionState.Connected)
</script>

<template>
  <div class="app">
    <Intro v-if="!connected" />
    <Game v-else />
    <Toaster />
  </div>
</template>

<style scoped>
.app {
  height: 100vh;
  width: 100vw;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
