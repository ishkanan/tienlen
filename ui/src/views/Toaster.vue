<template>
  <div slot-scope="{}"/>
</template>

<script lang="ts">
import Vue from 'vue';
import { ToastObject, ToastPosition } from 'vue-toasted';
import { GameEvent, GameEventKind } from '~/lib/models';
import { game } from '~/store/game';

const toastOptions = {
  position: 'bottom-right' as ToastPosition,
  duration: window.innerWidth < 1100 ? 3000 : 5000,
  keepOnHover: true,
  action: {
    text: 'Close',
    onClick: (_: Event, toastObject: ToastObject) => {
      toastObject.goAway(0);
    },
  },
};

export default Vue.extend({
  computed: {
    events(): GameEvent[] {
      return game.events;
    },
  },

  watch: {
    events: {
      immediate: true,
      handler(value: GameEvent[]) {
        if (value.length === 0) return;
        value.forEach(event => {
          Vue.toasted.show(
            event.message,
            {
              ...toastOptions,
              type: event.kind === GameEventKind.Error ? 'error' : 'info',
            },
          );
        });
        game.events = [];
      },
    },
  },
});
</script>

<style lang="postcss" module>
:global(.toasted) {
  font-size: 16px !important;
}

@media (max-width: 1100px) {
  :global(.toasted) {
    font-size: 13px !important;
  }
}
</style>