<script setup lang="ts">
import { format } from 'date-fns'
import { orderBy } from 'lodash-es'
import { computed } from 'vue'
import GameEvent, { type event } from '../components/GameEvent.vue'
import { EventSeverity } from '../lib/models'
import { useGameStore } from '../stores/game'

const gameStore = useGameStore()

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
</script>

<template>
  <div class="viewport">
    <div v-for="(event, i) in events" :key="i">
      <GameEvent :event="event" />
    </div>
  </div>
</template>

<style scoped>
.viewport {
  width: 100%;
  height: 100%;
  max-height: 100%;
  overflow-y: scroll;
}
</style>
