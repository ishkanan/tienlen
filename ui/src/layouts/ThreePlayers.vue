<template>
  <div :class="$style.viewport">
    <div :class="$style.player">
      <slot name="player"/>
    </div>
    <div :class="$style.opponent1">
      <slot name="opponent1"/>
    </div>
    <div :class="$style.opponent2">
      <slot name="opponent2"/>
    </div>
    <div :class="$style.discard">
      <slot name="discard"/>
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
  grid-area: 1 / 2 / 3 / 4;
}

.opponent2 {
  grid-area: 1 / 8 / 3 / 10;
}

.discard {
  grid-area: 3 / 3 / 5 / 9;
}

@media (max-width: 1100px) {
  .viewport {
    grid-template-columns: repeat(10, 10%);
    grid-template-rows: repeat(20, 5%);
  }

  .opponent1 {
    grid-area: 1 / 3 / 3 / 9;
    margin-top: 10px;
  }

  .opponent2 {
    grid-area: 3 / 3 / 5 / 9;
  }

  .discard {
    grid-area: 5 / 1 / 9 / 11;
  }

  .player {
    grid-area: 10 / 1 / 14 / 11;
  }
}
</style>