<template>
  <ds-draggable
    :list="orderedCards"
    :class="$style.hand"
    group="cards"
    @start="drag = true"
    @end="drag = false"
  >
    <ds-card
      v-for="card in orderedCards"
      :key="card.globalRank"
      :class="[$style.card, selectedMap[card.globalRank] && $style.raised]"
      :card="card"
      :selectable="true"
      :show-face="true"
      @selected="val => onSelectedToggle(card, val)"
    />
  </ds-draggable>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';
import Draggable from 'vuedraggable';
import CardView from './Card.vue';
import { Card } from '~/lib/models';

interface Data {
  orderedCards: Card[];
  selectedMap: Record<number, boolean>;
}

export default Vue.extend({
  components: {
    'ds-card': CardView,
    'ds-draggable': Draggable,
  },

  props: {
    cards: {
      type: Array as PropType<Card[]>,
      required: true,
    },
  },

  data(): Data {
    return {
      orderedCards: [],
      selectedMap: {},
    };
  },

  watch: {
    cards: {
      immediate: true,
      handler(cards: Card[]) {
        if (this.orderedCards.length === 0 || this.orderedCards.length < this.cards.length) {
          // new hand has been dealt, so accept initial order from server
          this.orderedCards = cards.map(c => c);
          return;
        }

        // remove played cards while preserving player-desired order
        const ranks = cards.map(c => c.globalRank);
        this.orderedCards = this.orderedCards.filter(c => ranks.includes(c.globalRank));
      },
    },
  },

  methods: {
    onSelectedToggle(card: Card, selected: boolean) {
      this.selectedMap[card.globalRank] = selected;
      if (!selected) delete this.selectedMap[card.globalRank];
      this.$forceUpdate();
      this.$emit('selected', Object.keys(this.selectedMap).map(s => parseInt(s)));
    },
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
  margin-top: 40px;
}

.raised {
  margin-top: 0px !important;
}
</style>