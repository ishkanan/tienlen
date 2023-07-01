<script setup lang="ts">
import { computed, watch } from 'vue'
import InGameLayout from '../layouts/InGame.vue'
import OnePlayerLayout from '../layouts/OnePlayerTable.vue'
import TwoPlayerLayout from '../layouts/TwoPlayerTable.vue'
import ThreePlayerLayout from '../layouts/ThreePlayerTable.vue'
import FourPlayerLayout from '../layouts/FourPlayerTable.vue'
import { type Player } from '../lib/models'
import { setTitle } from '../lib/utils'
import { useGameStore } from '../stores/game'
import ActivityView from '../views/Activity.vue'
import ControlsView from '../views/Controls.vue'
import DiscardView from '../views/Discard.vue'
import OpponentView from '../views/Opponent.vue'
import PlayerView from '../views/Player.vue'
import ScoreView from '../views/Score.vue'

const positionMap: Record<number, number[]> = {
  1: [2, 3, 4],
  2: [3, 4, 1],
  3: [4, 1, 2],
  4: [1, 2, 3],
}

const gameStore = useGameStore()

const player = computed(() => gameStore.self)

const opponents = computed<Player[]>(() => {
  if (!player.value) return []
  const positions = positionMap[player.value.position]
  if (!positions) return []
  return positions.reduce<Player[]>((memo, pos) => {
    const player = gameStore.opponents.find(p => p.position === pos)
    if (!player) return memo
    memo.push(player)
    return memo
  }, [])
})

const canStart = computed(() => !gameStore.isInProgress)
watch(
  canStart,
  (val) => {
    val && setTitle('Tiến lên || in lobby ...')
  },
  { immediate: true },
)

const paused = computed(() => gameStore.isPaused)
watch(
  paused,
  (val) => {
    val && setTitle('Tiến lên || game paused ...')
  },
  { immediate: true },
)
</script>

<template>
  <InGameLayout>
    <template #gameTable>
      <OnePlayerLayout v-if="player && opponents.length === 0">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
      </OnePlayerLayout>
      <TwoPlayerLayout v-if="player && opponents.length === 1">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
        <template #opponent1>
          <OpponentView :position="opponents[0].position" />
        </template>
      </TwoPlayerLayout>
      <ThreePlayerLayout v-if="player && opponents.length === 2">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
        <template #opponent1>
          <OpponentView :position="opponents[0].position" />
        </template>
        <template #opponent2>
          <OpponentView :position="opponents[1].position" />
        </template>
      </ThreePlayerLayout>
      <FourPlayerLayout v-if="player && opponents.length === 3">
        <template #discard><DiscardView /></template>
        <template #player><PlayerView /></template>
        <template #opponent1>
          <div class="nudge">
            <OpponentView :position="opponents[0].position" />
          </div>
        </template>
        <template #opponent2>
          <OpponentView :position="opponents[1].position" />
        </template>
        <template #opponent3>
          <div class="nudge">
            <OpponentView :position="opponents[2].position" />
          </div>
        </template>
      </FourPlayerLayout>
    </template>
    <template #scoreArea>
      <ScoreView />
    </template>
    <template #controlsArea>
      <ControlsView />
    </template>
    <template #activityArea>
      <ActivityView />
    </template>
  </InGameLayout>
</template>

<style scoped>
.nudge {
  margin-top: 6%;
}
</style>
