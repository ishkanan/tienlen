<script setup lang="ts">
import { computed } from 'vue'
import { sortBy } from 'lodash-es'
import CardView from '../components/Card.vue'
import { useGameStore } from '../stores/game'

const gameStore = useGameStore()

const canStart = computed(() => inLobby.value && gameStore.opponents.length > 0)

const inLobby = computed(() => gameStore.isInLobby)

const lastPlayed = computed(() => {
  if (gameStore.self?.lastPlayed) return gameStore.self
  return gameStore.opponents.find((o) => o.lastPlayed)
})

const lastPlayedCards = computed(() => sortBy(gameStore.lastPlayed || [], (c) => 52 - c.globalRank))

const needMorePlayers = computed(() => gameStore.opponents.length !== 3)

const paused = computed(() => gameStore.isPaused)

const previousWinner = computed(() => {
  if (gameStore.self?.wonLastGame) return gameStore.self
  return gameStore.opponents.find(o => o.wonLastGame)
})
</script>

<template>
  <div :class="{ viewport: true, winner: previousWinner && inLobby }">
    <h1 v-if="paused" class="title is-1 message">Game is paused!</h1>

    <template v-else-if="inLobby && previousWinner === undefined">
      <h1 v-if="needMorePlayers && !canStart" class="title is-1 message">Wait for more players...</h1>
      <h1 v-else-if="needMorePlayers && canStart" class="title is-1 message">
        Wait for more players
        <br />
        or start the game...
      </h1>
      <h1 v-else class="title is-1 message">Start the game...</h1>
    </template>

    <template v-else>
      <div class="messageAndLastPlayed">
        <h1 v-if="inLobby && previousWinner" class="title is-1 message">
          The game has finished! One more?
        </h1>
        <template v-else-if="lastPlayedCards.length > 0 && !!lastPlayed">
          <h3 class="title is-3 message">
            {{ lastPlayed.name }} played:
          </h3>
          <div class="lastPlayed">
            <div
              v-for="card in lastPlayedCards"
              class="cardWrapper"
            >
              <CardView
                :key="card.globalRank"
                :card="card"
                :selectable="false"
                :show-face="true"
              />
            </div>
          </div>
        </template>
      </div>
    </template>
  </div>
</template>

<style scoped>
.viewport {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: rgba(200, 200, 200, 0.4);
  border: black dashed 1px;
  border-radius: 10px;
}

.messageAndLastPlayed {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  flex-direction: column;
  justify-content: space-evenly;

  & h3 {
    margin: 5px 0 0;
  }
}

.lastPlayed {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

.cardWrapper {
  min-width: 60px;
}

.winner {
  flex-direction: column;
}

.message {
  background-color: transparent;
  color: black;
  text-align: center;
  text-transform: uppercase;
}
</style>
