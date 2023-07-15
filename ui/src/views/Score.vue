<script setup lang="ts">
import { computed } from 'vue'
import { orderBy } from 'lodash-es'
import { useGameStore } from '../stores/game'

interface ScoreLine {
  playerName: string
  score: number
  delta: number
}

const gameStore = useGameStore()

const deltas = computed(() => {
  if (gameStore.isInProgress) return {}

  return gameStore.winPlaces.reduce<Record<number, number>>((memo, player, i) => {
    memo[player.position] = gameStore.opponents.length - i
    return memo
  }, {})
})

const scores = computed<ScoreLine[]>(() => {
  const selfScore = {
    playerName: gameStore.self?.name || 'YOU',
    score: gameStore.self?.score || 0,
    delta: (gameStore.self && deltas.value[gameStore.self.position]) || 0,
  }

  const allScores = gameStore.opponents
    .map(opponent => ({
      playerName: opponent.name,
      score: opponent.score,
      delta: deltas.value[opponent.position] || 0,
    }))
    .concat(selfScore)

  return orderBy(allScores, ['score', 'playerName'], ['desc', 'asc'])
})
</script>

<template>
  <div class="viewport">
    <div v-for="scoreLine in scores" :key="scoreLine.playerName" class="scoreLine">
      <h3>{{ scoreLine.playerName }}</h3>
      <h3 v-if="scoreLine.delta > 0">{{ scoreLine.score }} ( + {{ scoreLine.delta }} )</h3>
      <h3 v-else>{{ scoreLine.score }}</h3>
    </div>
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

  & .scoreLine {
    width: 90%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  & h3 {
    margin: 0;
  }
}
</style>
