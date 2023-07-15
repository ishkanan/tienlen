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
  <div>
    <Draggable
      :list="orderedCards"
      item-key="globalRank"
      class="hand"
      group="cards"
      delay="150"
      delay-on-touch-only="true"
      @start="drag = true"
      @end="drag = false"
    >
      <template #item="{element}: {element: Card}">
        <div :class="{
            card: true,
            raised: selectedMap[element.globalRank]
          }"
        >
          <CardView
            :card="element"
            :selectable="true"
            :show-face="true"
            @selected="(val) => onSelectedToggle(element, val)"
          />
        </div>
      </template>
    </Draggable>
  </div>
</template>

<style scoped>
.hand {
  justify-content: center;
  display: grid;
  grid-template-columns: repeat(auto-fit,  minmax(10px, max-content));
}

.card {
  padding-top: 30px;
  border-radius: 8px;
}

.raised {
  padding-top: 0px !important;
}
</style>
