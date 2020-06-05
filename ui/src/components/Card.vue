<template>
  <div
    :class="$style.card"
    :style="style"
    @click="onClick"
  >
    <slot name="info"/>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue';
import { Card, Suit } from '~/lib/models';

const suitYmap: Record<Suit, number> = {
  [Suit.Spades]: 360,
  [Suit.Clubs]: 0,
  [Suit.Diamonds]: 120,
  [Suit.Hearts]: 240,
};

interface Data {
  selected: boolean;
}

export default Vue.extend({
  props: {
    card: {
      type: Object as PropType<Card>,
      required: true,
    },
    selectable: {
      type: Boolean,
      required: true,
    },
    showFace: {
      type: Boolean,
      required: true,
    },
  },

  data(): Data {
    return {
      selected: false,
    };
  },

  computed: {
    offsetX(): number {
      if (!this.showFace) return 0;
      return 80 * this.card.faceValue;
    },
    offsetY(): number {
      if (!this.showFace) return [Suit.Spades, Suit.Clubs].includes(this.card.suit) ? 0 : 120;
      return suitYmap[this.card.suit];
    },
    style(): Record<string, unknown> {
      return {
        cursor: this.selectable ? 'pointer' : 'default',
        backgroundPosition: `top -${this.offsetY}px left -${this.offsetX}px`,
      };
    },
  },

  methods: {
    onClick() {
      if (!this.selectable) return;
      this.selected = !this.selected;
      this.$emit('selected', this.selected);
    },
  },
});
</script>

<style lang="postcss" module>
.card {
  width: 80px;
  height: 120px;
  background: url(../assets/images/cards.png);
}
</style>