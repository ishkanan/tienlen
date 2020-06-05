<template>
  <div slot-scope="{}"/>
</template>

<script lang="ts">
import Vue from 'vue';
import { ToastObject, ToastPosition } from 'vue-toasted';
import { GameEvent, GameEventKind } from '~/lib/models';
import { game } from '~/store/game';

const toastOptions = {
  position: 'top-right' as ToastPosition,
  duration: 10000,
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
  font-size: 18px !important;
}
</style>