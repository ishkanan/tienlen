<template>
  <div slot-scope="{}" />
</template>

<script lang="ts">
import Vue from 'vue';
import { defineComponent, computed, watch } from '@vue/composition-api';
import { ToastObject, ToastPosition } from 'vue-toasted';
import { EventSeverity } from '~/lib/models';
import { game } from '~/store/game';

export default defineComponent({
  setup() {
    const events = computed(() => {
      const now = new Date();
      return game.events.filter((e) => e.timestamp >= now && e.toast);
    });

    const toastOptions = {
      position: 'top-center' as ToastPosition,
      duration: 5000,
      keepOnHover: true,
      action: {
        text: 'Close',
        onClick: (_: Event, toastObject: ToastObject) => {
          toastObject.goAway(0);
        },
      },
    };

    const toastTypeMap: Record<EventSeverity, string> = {
      [EventSeverity.Info]: 'info',
      [EventSeverity.Error]: 'error',
      [EventSeverity.Warning]: 'success',
    };

    watch(
      events,
      (val) => {
        if (val.length === 0) return;
        val.forEach((e) => {
          e.runes.forEach((r) => {
            if (!r.message) return;
            Vue.toasted.show(r.message, {
              ...toastOptions,
              type: toastTypeMap[e.severity] ?? 'info',
            });
          });
        });
      },
      { immediate: true },
    );

    return {};
  },
});
</script>

<style lang="postcss" module>
:global(.toasted) {
  font-size: 16px !important;
}
</style>
