<script lang="ts">
import { computed, watch } from 'vue'
import {
  type Toasted,
  type ToastObject,
  type ToastPosition,
} from 'vue-toasted'
import { EventSeverity } from '../lib/models'
import { useGameStore } from '../stores/game'
import { app } from '../main'

interface DynamicToasted {
  toasted: Toasted
}

// injected by vue-toasted at init
// see https://github.com/shakee93/vue-toasted/blob/master/src/index.js
const toaster = ((app as unknown) as DynamicToasted).toasted

const gameStore = useGameStore()

const events = computed(() => {
  const now = new Date()
  return gameStore.events.filter(e => e.timestamp >= now && e.toast)
})

const toastOptions = {
  position: 'top-center' as ToastPosition,
  duration: 5000,
  keepOnHover: true,
  action: {
    text: 'Close',
    onClick: (_: Event, toastObject: ToastObject) => {
      toastObject.goAway(0)
    },
  },
}

const toastTypeMap: Record<EventSeverity, string> = {
  [EventSeverity.Info]: 'info',
  [EventSeverity.Error]: 'error',
  [EventSeverity.Warning]: 'success',
  [EventSeverity.Success]: 'info',
}

watch(
  events,
  (val) => {
    if (val.length === 0) return
    val.forEach((e) => {
      e.runes.forEach((r) => {
        if (!r.message) return
        toaster.show(r.message, {
          ...toastOptions,
          type: toastTypeMap[e.severity] ?? 'info',
        })
      })
    })
  },
  { immediate: true },
)
</script>

<template>
  <div slot-scope="{}" />
</template>

<style scoped>
:global(.toasted) {
  font-size: 16px !important;
}
</style>
