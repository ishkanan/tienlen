<template>
  <div :class="$style.viewport">
    <div :class="$style.player">
      <slot name="player"/>
    </div>
    <div :class="$style.opponent1">
      <slot name="opponent1"/>
    </div>
    <div :class="$style.discard">
      <slot name="discard"/>
    </div>
    <div :class="$style.controls">
      <slot name="controls"/>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({});
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

.opponent1 {
  grid-area: 1 / 5 / 3 / 7;
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
    grid-area: 8 / 1 / 12 / 11;
  }

  .opponent1 {
    grid-area: 1 / 3 / 3 / 9;
    margin-top: 10px;
  }

  .discard {
    grid-area: 3 / 1 / 7 / 11;
  }

  .controls {
    grid-area: 1 / 1 / 2 / 3;
  }
}
</style>