<template>
  <div :class="$style.viewport">
    <div :class="$style.player">
      <slot name="player" />
    </div>
    <div :class="$style.opponent1">
      <slot name="opponent1" />
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
  grid-template-columns: repeat(40, 2.5%);
  grid-template-rows: repeat(17, 5.88%);
}

.player {
  grid-area: 12 / 5 / 18 / 37;
}

.discard {
  grid-area: 7 / 9 / 12 / 33;
}

.opponent1 {
  grid-area: 1 / 16 / 6 / 26;
}
</style>
