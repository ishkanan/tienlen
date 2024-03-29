<template>
  <Draggable
    :list="orderedCards"
    :class="$style.hand"
    group="cards"
    delay="150"
    delay-on-touch-only="true"
    @start="drag = true"
    @end="drag = false"
  >
    <CardView
      v-for="card in orderedCards"
      :key="card.globalRank"
      :class="[$style.card, selectedMap[card.globalRank] && $style.raised]"
      :card="card"
      :selectable="true"
      :show-face="true"
      @selected="(val) => onSelectedToggle(card, val)"
    />
  </Draggable>
</template>

<script lang="ts">
import { defineComponent, ref, watch, PropType } from '@vue/composition-api';
import Draggable from 'vuedraggable';
import CardView from './Card.vue';
import { Card } from '~/lib/models';

export default defineComponent({
  components: { CardView, Draggable },

  props: {
    cards: {
      type: Array as PropType<Card[]>,
      required: true,
    },
  },

  setup(props, { emit }) {
    const orderedCards = ref<Card[]>([]);
    const selectedMap = ref<Record<number, boolean>>({});

    watch(
      () => props.cards,
      (val) => {
        if (orderedCards.value.length === 0 || orderedCards.value.length < props.cards.length) {
          // new hand has been dealt, so accept initial order from server
          orderedCards.value = val.slice();
          return;
        }

        // remove played cards while preserving player-desired order
        const ranks = val.map((c) => c.globalRank);
        orderedCards.value = orderedCards.value.filter((c) => ranks.includes(c.globalRank));
      },
      { immediate: true },
    );

    const onSelectedToggle = (card: Card, selected: boolean) => {
      selectedMap.value[card.globalRank] = selected;
      if (!selected) delete selectedMap.value[card.globalRank];
      orderedCards.value = orderedCards.value.slice();
      emit(
        'selected',
        Object.keys(selectedMap.value).map((s) => parseInt(s)),
      );
    };

    return { orderedCards, selectedMap, onSelectedToggle };
  },
});
</script>

<style lang="postcss" module>
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
