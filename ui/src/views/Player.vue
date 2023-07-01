<script setup lang="ts">
import { computed, inject, ref, watch } from 'vue'
import CardView from '../components/Card.vue'
import Hand from '../components/Hand.vue'
import { type Card, Suit, EventSeverity } from '../lib/models'
import { type Socket } from '../lib/socket'
import { ordinalise, startFlashTitle } from '../lib/utils'
import { useGameStore } from '../stores/game'

const socket: Socket | undefined = inject('socket')

const autoPassed = ref(false)
const autoPassing = ref(false)
const selectedRanks = ref<number[]>([])

const gameStore = useGameStore()

const player = computed(() => gameStore.self)
const winPlace = computed(() => {
  const placed = gameStore.winPlaces.findIndex(p => p.position === gameStore.self?.position)
  return placed === -1 ? 0 : placed + 1
})
const ordinalisedWinPlace = computed(() => ordinalise(winPlace.value))
const waiting = computed(
  () => gameStore.isInProgress && (!player.value?.isTurn ?? false) && winPlace.value === 0,
)

const canStart = computed(() => !gameStore.isInProgress && gameStore.opponents.length > 0)
const canPass = computed(
  () =>
    gameStore.isInProgress &&
    !gameStore.isPaused &&
    !gameStore.firstRound &&
    !gameStore.newRound &&
    (player.value?.isTurn ?? false) &&
    !autoPassing.value &&
    !autoPassed.value,
)
const canPlay = computed(
  () =>
    gameStore.isInProgress &&
    !gameStore.isPaused &&
    (player.value?.isTurn ?? false) &&
    !autoPassing.value &&
    !autoPassed.value,
)

const cardsSelected = computed(() => selectedRanks.value.length > 0)
const showHand = computed(() => gameStore.isInProgress && winPlace.value === 0)
const hand = computed(() => gameStore.selfHand)
const unfaced = computed(() => {
  return {
    suit: Suit.Spades,
    faceValue: 2,
    globalRank: 1,
    suitRank: 1,
  }
})
const paused = computed(() => gameStore.isPaused)
const passed = computed(() => player.value?.isPassed ?? false)
const newRound = computed(() => gameStore.newRound)

const doStart = () => socket?.requestStartGame()
const doPass = () => socket?.requestTurnPass()
const doPlay = () => {
  const cards = selectedRanks.value.reduce<Card[]>((memo, rank) => {
    const card = gameStore.selfHand.find(c => c.globalRank === rank)
    if (!card) return memo
    memo.push(card)
    return memo
  }, [])
  socket?.requestTurnPlay({ cards })
}
const onSelected = (ranks: number[]) => (selectedRanks.value = ranks)

watch(
  canPlay,
  (val) => {
    // we auto-skip if we have less cards than other players' last played
    // (since it is impossible to beat with less cards)
    if (
      canPlay.value &&
      gameStore.lastPlayed.length > hand.value.length &&
      (!player.value?.lastPlayed ?? false)
    ) {
      autoPassing.value = true
      window.setTimeout(() => {
        doPass()
        gameStore.pushEvent({
          severity: EventSeverity.Warning,
          runes: [
            { message: 'You were auto-passed as you have less cards than the current hand.' },
          ],
        })
        autoPassing.value = false
        autoPassed.value = true
      }, 2000)
      return
    }

    const name = player.value?.name || ''
    val && startFlashTitle(`Tiến lên || ${name}`, '★ ★ IT IS YOUR TURN ★ ★')
  },
  { immediate: true },
)

watch(
  newRound,
  (val) => {
    if (val) autoPassed.value = false
  },
  { immediate: true },
)

watch(
  paused,
  (val) => {
    const name = player.value?.name || ''
    !val &&
      player.value?.isTurn &&
      startFlashTitle(`Tiến lên || ${name}`, '★ ★ IT IS YOUR TURN ★ ★')
  },
  { immediate: true },
)
</script>

<template>
  <div v-if="player" class="viewport">
    <div class="controls">
      <button v-if="canStart" @click="doStart">Start game</button>
      <button v-if="canPlay" :disabled="!cardsSelected" @click="doPlay">Play cards</button>
      <button v-if="canPass" class="danger" @click="doPass">Pass turn</button>
      <h2 v-if="autoPassing">Your turn will be automatically passed...</h2>
      <h2 v-if="passed">You have passed! Sit tight.</h2>
      <h2 v-else-if="waiting">Waiting for turn...</h2>
      <h2 v-if="winPlace > 0 && !canStart">All done bucko! Have a break.</h2>
    </div>

    <div class="hand">
      <div v-if="winPlace > 0" class="placed">
        <h2 class="note">{{ ordinalisedWinPlace }}</h2>
      </div>

      <Hand v-else-if="showHand" :cards="hand" @selected="onSelected" />

      <CardView
        v-else
        class="unfaced"
        :card="unfaced"
        :selectable="false"
        :show-face="false"
      />
    </div>

    <div class="nameBar">
      <h3 :class="{ isTurn: player.isTurn }">{{ player.name }}</h3>
    </div>
  </div>
</template>

<style scoped>
.viewport {
  width: 100%;
  max-height: 100%;
  margin: 15px 0 0 0;
}

.controls {
  width: 100%;
  color: white;
  height: 40px;
  display: flex;
  flex-direction: row;
  justify-content: center;

  & button:nth-child(2) {
    margin-left: 80px;
  }

  & h1 {
    color: white;
  }
}

.hand {
  width: 100%;
  margin-top: 15px;
  display: flex;
  flex-direction: row;
  justify-content: center;

  & .placed {
    width: 100px;
    height: 120px;
    background: url(../assets/images/trophy.png);
    background-size: cover;
    background-repeat: no-repeat;
    margin-top: 30px;
    text-align: center;

    & .note {
      background-color: black;
      border-radius: 5px;
      color: #f2f2f2;
      padding: 0px;
      width: 60%;
      margin: 40px auto auto auto;
    }
  }

  & .unfaced {
    margin-top: 30px;
  }
}

.nameBar {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  color: #f2f2f2;

  & h3 {
    margin: 10px 0 10px 0;
    padding: 2px 6px 2px 6px;
  }

  & .isTurn {
    background-color: #f2f2f2;
    border-radius: 5px;
    color: black;
    border: 1px solid black;
    padding: 1px 5px 1px 5px;
  }
}
</style>
