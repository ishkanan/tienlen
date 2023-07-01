<template>
  <div class="viewport">
    <div
      v-for="(event, i) in events"
      :key="i"
      :class="{
        event,
        error: event.isError,
        warning: event.isWarning,
        success: event.isSuccess,
      }"
    >
      <template v-for="(rune, j) in event.runes">
        <p v-if="j === 0" :key="j" class="rune runeTime">
          {{ `[${event.timestamp}]:` }}
        </p>
        <p v-if="rune.isMessage" :key="j * 3 + 1" class="rune">{{ rune.message }}</p>
        <p
          v-if="rune.card"
          :key="j * 3 + 2"
          class="rune runeCard"
          :style="runeCardImageStyle(rune.card)"
        />
      </template>
    </div>
  </div>
</template>

<script setup lang="ts">
import { format } from 'date-fns'
import { orderBy } from 'lodash-es'
import { computed, inject, ref } from 'vue'
import { type Card, EventSeverity, Suit } from '../lib/models'
import { useGameStore } from '../stores/game'

interface event {
  timestamp: string
  isSuccess: boolean
  isError: boolean
  isWarning: boolean
  runes: rune[]
}

interface rune {
  isCard: boolean
  isMessage: boolean
  card?: Card
  message?: string
}

const gameStore = useGameStore()

const suitYmap: Record<Suit, number> = {
  [Suit.Diamonds]: 0,
  [Suit.Hearts]: 31,
  [Suit.Spades]: 62,
  [Suit.Clubs]: 93,
}

const events = computed(() =>
  orderBy(gameStore.events, ['timestamp'], ['desc'])
    .filter((e) => !e.toast)
    .map<event>((e) => ({
      timestamp: format(e.timestamp, 'kk:mm:ss'),
      isSuccess: e.severity === EventSeverity.Success,
      isError: e.severity === EventSeverity.Error,
      isWarning: e.severity === EventSeverity.Warning,
      runes: e.runes.map((r) => ({
        isCard: r.card !== undefined,
        isMessage: r.message !== undefined,
        card: r.card,
        message: r.message,
      })),
    })),
)

const runeCardImageStyle = (card: Card): Record<string, string> => {
  const offsetX = 21 * (card.faceValue - 1)
  const offsetY = suitYmap[card.suit]
  return {
    backgroundPosition: `top -${offsetY}px left -${offsetX}px`,
  }
}

</script>

<style scoped>
.viewport {
  width: 100%;
  height: 100%;
  max-height: 100%;
  overflow-y: scroll;

  & .event {
    margin: 4px 6px 0 6px;
    max-width: calc(100% - 12px);
    color: rgba(210, 210, 210);

    & .rune {
      display: inline-block;
      color: inherit;
    }

    & .runeCard {
      width: 21px;
      height: 32px;
      background: url(../assets/images/cards-runes.png);
      margin-left: 6px;
      vertical-align: middle;
    }

    & .runeTime {
      margin-right: 5px;
    }
  }
}

p {
  font-size: 20px;
  font-weight: 100;
  margin: 0;
  word-wrap: break-word;
}

.warning {
  color: rgba(255, 238, 0, 0.8) !important;
}

.error {
  color: rgba(220, 0, 0.8) !important;
}

.success {
  color: rgba(0, 220, 0.8) !important;
}
</style>
