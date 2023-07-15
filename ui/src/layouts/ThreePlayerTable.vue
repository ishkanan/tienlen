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
  grid-area: 1 / 5 / 5 / 15;
}

.opponent2 {
  grid-area: 1 / 27 / 5 / 37;
}
</style>
