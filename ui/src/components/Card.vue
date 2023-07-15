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

const imageSrc = computed(() => {
  const prefix = "/src/assets/images/cards/tile"

  if (!props.showFace) {
    return redCard.value
      ? `${prefix}-back-red.png`
      : `${prefix}-back-black.png`
  }

  return `${prefix}${props.card.globalRank}.png`
})

const onClick = () => {
  if (!props.selectable) return
  selected.value = !selected.value
  emit('selected', selected.value)
}
</script>

<template>
  <img
    :src="imageSrc"
    class="tile"
    :class="{ selectable: props.selectable }"
    @click="onClick"
  />
  <slot name="info" />
</template>

<style scoped>
.tile {
  border-radius: 8px;
  cursor: default;
}

.selectable {
  cursor: pointer;
}
</style>
