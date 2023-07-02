<script setup lang="ts">
import { computed, watch } from 'vue'
import {
  toast,
  type ToastOptions,
  type ToastPosition,
  type ToastType,
} from 'vue3-toastify'
import { EventSeverity } from '../lib/models'
import { useGameStore } from '../stores/game'

const gameStore = useGameStore()

const events = computed(() => {
  const now = new Date()
  return gameStore.events.filter(e => e.timestamp >= now && e.toast)
})

const toastOptions: ToastOptions = {
  position: 'top-center' as ToastPosition,
  autoClose: 5000,
  pauseOnHover: true,
  closeOnClick: true,
}

const toastTypeMap: Record<EventSeverity, ToastType> = {
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
        toast(r.message, {
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
