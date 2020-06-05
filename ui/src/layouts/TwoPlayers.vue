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
</style>