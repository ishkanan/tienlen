<script setup lang="ts">
import { v4 as uuidv4 } from 'uuid'
import { type Card } from '../lib/models'

export interface event {
  timestamp: string
  isSuccess: boolean
  isError: boolean
  isWarning: boolean
  runes: rune[]
}

export interface rune {
  isCard: boolean
  isMessage: boolean
  card?: Card
  message?: string
}

defineProps<{
  event: event,
}>()

const uuid = () => uuidv4()
</script>

<template>
  <div
    :class="{
      event,
      error: event.isError,
      warning: event.isWarning,
      success: event.isSuccess,
    }"
  >
    <template v-for="(rune, index) in event.runes">
      <p v-if="index === 0" :key="uuid()" class="rune runeTime">
        {{ `[${event.timestamp}]:` }}
      </p>
      <p v-if="rune.isMessage" :key="uuid()" class="rune">{{ rune.message }}</p>
      <p
        v-if="rune.card"
        :key="uuid()"
        :class="{ [`tile${rune.card.globalRank}`]: true }"
        class="rune runeCard"
      />
    </template>
  </div>
</template>

<style scoped>
.event {
  margin: 4px 6px 0 6px;
  max-width: calc(100% - 12px);
  color: rgba(210, 210, 210);
}

p {
  font-size: 16px;
  font-weight: 90;
  margin: 0;
  word-wrap: break-word;
}

.warning {
  color: rgba(255, 238, 0, 0.8) !important;
}

.error {
  color: rgba(220, 0, 0.8) !important;
}

.success {
  color: rgba(0, 220, 0.8) !important;
}

.rune {
  display: inline-block;
  color: inherit;
}

.runeTime {
  margin-right: 5px;
}

.runeCard {
  width: 21px;
  height: 32px;
  margin-left: 6px;
  vertical-align: middle;
}

.tile1 {
  background: url(../assets/images/card-runes/tile1.png);
}
.tile2 {
  background: url(../assets/images/card-runes/tile2.png);
}
.tile3 {
  background: url(../assets/images/card-runes/tile3.png);
}
.tile4 {
  background: url(../assets/images/card-runes/tile4.png);
}
.tile5 {
  background: url(../assets/images/card-runes/tile5.png);
}
.tile6 {
  background: url(../assets/images/card-runes/tile6.png);
}
.tile7 {
  background: url(../assets/images/card-runes/tile7.png);
}
.tile8 {
  background: url(../assets/images/card-runes/tile8.png);
}
.tile9 {
  background: url(../assets/images/card-runes/tile9.png);
}
.tile10 {
  background: url(../assets/images/card-runes/tile10.png);
}
.tile11 {
  background: url(../assets/images/card-runes/tile11.png);
}
.tile12 {
  background: url(../assets/images/card-runes/tile12.png);
}
.tile13 {
  background: url(../assets/images/card-runes/tile13.png);
}
.tile14 {
  background: url(../assets/images/card-runes/tile14.png);
}
.tile15 {
  background: url(../assets/images/card-runes/tile15.png);
}
.tile16 {
  background: url(../assets/images/card-runes/tile16.png);
}
.tile17 {
  background: url(../assets/images/card-runes/tile17.png);
}
.tile18 {
  background: url(../assets/images/card-runes/tile18.png);
}
.tile19 {
  background: url(../assets/images/card-runes/tile19.png);
}
.tile20 {
  background: url(../assets/images/card-runes/tile20.png);
}
.tile21 {
  background: url(../assets/images/card-runes/tile21.png);
}
.tile22 {
  background: url(../assets/images/card-runes/tile22.png);
}
.tile23 {
  background: url(../assets/images/card-runes/tile23.png);
}
.tile24 {
  background: url(../assets/images/card-runes/tile24.png);
}
.tile25 {
  background: url(../assets/images/card-runes/tile25.png);
}
.tile26 {
  background: url(../assets/images/card-runes/tile26.png);
}
.tile27 {
  background: url(../assets/images/card-runes/tile27.png);
}
.tile28 {
  background: url(../assets/images/card-runes/tile28.png);
}
.tile29 {
  background: url(../assets/images/card-runes/tile29.png);
}
.tile30 {
  background: url(../assets/images/card-runes/tile30.png);
}
.tile31 {
  background: url(../assets/images/card-runes/tile31.png);
}
.tile32 {
  background: url(../assets/images/card-runes/tile32.png);
}
.tile33 {
  background: url(../assets/images/card-runes/tile33.png);
}
.tile34 {
  background: url(../assets/images/card-runes/tile34.png);
}
.tile35 {
  background: url(../assets/images/card-runes/tile35.png);
}
.tile36 {
  background: url(../assets/images/card-runes/tile36.png);
}
.tile37 {
  background: url(../assets/images/card-runes/tile37.png);
}
.tile38 {
  background: url(../assets/images/card-runes/tile38.png);
}
.tile39 {
  background: url(../assets/images/card-runes/tile39.png);
}
.tile40 {
  background: url(../assets/images/card-runes/tile40.png);
}
.tile41 {
  background: url(../assets/images/card-runes/tile41.png);
}
.tile42 {
  background: url(../assets/images/card-runes/tile42.png);
}
.tile43 {
  background: url(../assets/images/card-runes/tile43.png);
}
.tile44 {
  background: url(../assets/images/card-runes/tile44.png);
}
.tile45 {
  background: url(../assets/images/card-runes/tile45.png);
}
.tile46 {
  background: url(../assets/images/card-runes/tile46.png);
}
.tile47 {
  background: url(../assets/images/card-runes/tile47.png);
}
.tile48 {
  background: url(../assets/images/card-runes/tile48.png);
}
.tile49 {
  background: url(../assets/images/card-runes/tile49.png);
}
.tile50 {
  background: url(../assets/images/card-runes/tile50.png);
}
.tile51 {
  background: url(../assets/images/card-runes/tile51.png);
}
.tile52 {
  background: url(../assets/images/card-runes/tile52.png);
}
</style>
