<script setup lang="ts">
import { ref, watch } from 'vue'
import Draggable from 'vuedraggable'
import CardView from './Card.vue'
import { type Card } from '../lib/models'

const props = defineProps<{
  cards: Card[],
}>()

const emit = defineEmits<{
  (e: 'selected', value: number[]): void
}>()

const orderedCards = ref<Card[]>([])
const selectedMap = ref<Record<number, Boolean>>({})
const drag = ref(false)

watch(
  () => props.cards,
  (val) => {
    if (orderedCards.value.length === 0 || orderedCards.value.length < props.cards.length) {
      // new hand has been dealt, so accept initial order from server
      orderedCards.value = val.slice()
      return
    }

    // remove played cards while preserving player-desired order
    const ranks = val.map((c) => c.globalRank)
    orderedCards.value = orderedCards.value.filter((c) => ranks.includes(c.globalRank))
  },
  { immediate: true },
)

const onSelectedToggle = (card: Card, selected: Boolean) => {
  selectedMap.value[card.globalRank] = selected
  if (!selected) delete selectedMap.value[card.globalRank]
  orderedCards.value = orderedCards.value.slice()
  emit(
    'selected',
    Object.keys(selectedMap.value).map((s) => parseInt(s)),
  )
}
</script>

<template>
  <Draggable
    :list="orderedCards"
    class="hand"
    group="cards"
    delay="150"
    delay-on-touch-only="true"
    @start="drag = true"
    @end="drag = false"
  >
    <CardView
      v-for="card in orderedCards"
      :key="card.globalRank"
      :class="{
        card: true,
        raised: selectedMap[card.globalRank]
      }"
      :card="card"
      :selectable="true"
      :show-face="true"
      @selected="(val) => onSelectedToggle(card, val)"
    />
  </Draggable>
</template>

<style scoped>
.hand {
  width: 100%;
  display: flex;
  justify-content: center;
}

.card {
  margin-top: 30px;
}

.raised {
  margin-top: 0px !important;
}
</style>
