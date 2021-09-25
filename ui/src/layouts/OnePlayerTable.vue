<template>
  <div :class="$style.viewport">
    <div :class="$style.player">
      <slot name="player" />
    </div>
    <div :class="$style.discard">
      <slot name="discard" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api';

export default defineComponent({});
</script>

<style lang="postcss" module>
.viewport {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: repeat(10, 10%);
  grid-template-rows: repeat(6, 16.67%);
}

.player {
  grid-area: 5 / 2 / 7 / 10;
}

.discard {
  grid-area: 3 / 3 / 5 / 9;
}

.controls {
  grid-area: 1 / 1 / 2 / 2;
}

@media (max-width: 1100px) {
  .viewport {
    grid-template-columns: repeat(10, 10%);
    grid-template-rows: repeat(20, 5%);
  }

  .player {
    grid-area: 7 / 1 / 11 / 11;
  }

  .discard {
    grid-area: 2 / 1 / 6 / 11;
  }

  .controls {
    grid-area: 1 / 1 / 2 / 3;
  }
}
</style>
