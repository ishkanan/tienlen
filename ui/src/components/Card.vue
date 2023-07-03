<script setup lang="ts">
import { computed, ref } from 'vue'
import { type Card, Suit } from '../lib/models'

const props = defineProps<{
  card: Card,
  selectable: Boolean,
  showFace: Boolean,
}>()

const emit = defineEmits<{
  (e: 'selected', value: Boolean): void
}>()

const selected = ref(false)

const redCard = computed(() => [Suit.Hearts, Suit.Diamonds].includes(props.card.suit))

const onClick = () => {
  if (!props.selectable) return
  selected.value = !selected.value
  emit('selected', selected.value)
}
</script>

<template>
  <div
    class="tile"
    :class="{
      selectable: props.selectable,
      [`tile${card.globalRank}`]: props.showFace,
      [`tile-back-red`]: !props.showFace && redCard,
      [`tile-back-black`]: !props.showFace && !redCard,
    }" @click="onClick">
    <slot name="info" />
  </div>
</template>

<style scoped>
.tile {
  cursor: default;
  width: 80px;
  height: 120px;
  max-width: 80px;
  max-height: 120px;
}

.selectable {
  cursor: pointer;
}

.tile-back-red {
  background: url(../assets/images/cards/tile-back-red.png);
}
.tile-back-black {
  background: url(../assets/images/cards/tile-back-black.png);
}
.tile1 {
  background: url(../assets/images/cards/tile1.png);
}
.tile2 {
  background: url(../assets/images/cards/tile2.png);
}
.tile3 {
  background: url(../assets/images/cards/tile3.png);
}
.tile4 {
  background: url(../assets/images/cards/tile4.png);
}
.tile5 {
  background: url(../assets/images/cards/tile5.png);
}
.tile6 {
  background: url(../assets/images/cards/tile6.png);
}
.tile7 {
  background: url(../assets/images/cards/tile7.png);
}
.tile8 {
  background: url(../assets/images/cards/tile8.png);
}
.tile9 {
  background: url(../assets/images/cards/tile9.png);
}
.tile10 {
  background: url(../assets/images/cards/tile10.png);
}
.tile11 {
  background: url(../assets/images/cards/tile11.png);
}
.tile12 {
  background: url(../assets/images/cards/tile12.png);
}
.tile13 {
  background: url(../assets/images/cards/tile13.png);
}
.tile14 {
  background: url(../assets/images/cards/tile14.png);
}
.tile15 {
  background: url(../assets/images/cards/tile15.png);
}
.tile16 {
  background: url(../assets/images/cards/tile16.png);
}
.tile17 {
  background: url(../assets/images/cards/tile17.png);
}
.tile18 {
  background: url(../assets/images/cards/tile18.png);
}
.tile19 {
  background: url(../assets/images/cards/tile19.png);
}
.tile20 {
  background: url(../assets/images/cards/tile20.png);
}
.tile21 {
  background: url(../assets/images/cards/tile21.png);
}
.tile22 {
  background: url(../assets/images/cards/tile22.png);
}
.tile23 {
  background: url(../assets/images/cards/tile23.png);
}
.tile24 {
  background: url(../assets/images/cards/tile24.png);
}
.tile25 {
  background: url(../assets/images/cards/tile25.png);
}
.tile26 {
  background: url(../assets/images/cards/tile26.png);
}
.tile27 {
  background: url(../assets/images/cards/tile27.png);
}
.tile28 {
  background: url(../assets/images/cards/tile28.png);
}
.tile29 {
  background: url(../assets/images/cards/tile29.png);
}
.tile30 {
  background: url(../assets/images/cards/tile30.png);
}
.tile31 {
  background: url(../assets/images/cards/tile31.png);
}
.tile32 {
  background: url(../assets/images/cards/tile32.png);
}
.tile33 {
  background: url(../assets/images/cards/tile33.png);
}
.tile34 {
  background: url(../assets/images/cards/tile34.png);
}
.tile35 {
  background: url(../assets/images/cards/tile35.png);
}
.tile36 {
  background: url(../assets/images/cards/tile36.png);
}
.tile37 {
  background: url(../assets/images/cards/tile37.png);
}
.tile38 {
  background: url(../assets/images/cards/tile38.png);
}
.tile39 {
  background: url(../assets/images/cards/tile39.png);
}
.tile40 {
  background: url(../assets/images/cards/tile40.png);
}
.tile41 {
  background: url(../assets/images/cards/tile41.png);
}
.tile42 {
  background: url(../assets/images/cards/tile42.png);
}
.tile43 {
  background: url(../assets/images/cards/tile43.png);
}
.tile44 {
  background: url(../assets/images/cards/tile44.png);
}
.tile45 {
  background: url(../assets/images/cards/tile45.png);
}
.tile46 {
  background: url(../assets/images/cards/tile46.png);
}
.tile47 {
  background: url(../assets/images/cards/tile47.png);
}
.tile48 {
  background: url(../assets/images/cards/tile48.png);
}
.tile49 {
  background: url(../assets/images/cards/tile49.png);
}
.tile50 {
  background: url(../assets/images/cards/tile50.png);
}
.tile51 {
  background: url(../assets/images/cards/tile51.png);
}
.tile52 {
  background: url(../assets/images/cards/tile52.png);
}
</style>
