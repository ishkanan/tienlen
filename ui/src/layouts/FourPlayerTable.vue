<template>
  <div class="viewport">
    <div class="player">
      <slot name="player" />
    </div>
    <div class="opponent1">
      <slot name="opponent1" />
    </div>
    <div class="opponent2">
      <slot name="opponent2" />
    </div>
    <div class="opponent3">
      <slot name="opponent3" />
    </div>
    <div class="discard">
      <slot name="discard" />
    </div>
  </div>
</template>

<style scoped>
.viewport {
  height: 100%;
  width: 100%;
  display: grid;
  grid-template-columns: repeat(40, 2.5%);
  grid-template-rows: repeat(17, 5.88%);
}

.player {
  grid-area: 11 / 5 / 18 / 37;
}

.discard {
  grid-area: 6 / 9 / 11 / 33;
}

.opponent1 {
  grid-area: 6 / 1 / 11 / 9;
}

.opponent2 {
  grid-area: 1 / 16 / 5 / 26;
}

.opponent3 {
  grid-area: 6 / 33 / 11 / 41;
}
</style>
