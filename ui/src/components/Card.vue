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
import WindowSizable from '~/mixins/WindowSizable.vue';

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
  mixins: [WindowSizable],

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
    scaleFactor(): number {
      return this.$data.windowWidth <= 1100 ? 0.75 : 1;
    },
    offsetX(): number {
      if (!this.showFace) return 0;
      return (80 * this.scaleFactor) * this.card.faceValue;
    },
    offsetY(): number {
      if (!this.showFace) return [Suit.Spades, Suit.Clubs].includes(this.card.suit) ? 0 : 120 * this.scaleFactor;
      return suitYmap[this.card.suit] * this.scaleFactor;
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

@media (max-width: 1100px) {
  .card {
    width: 60px;
    height: 90px;
    background: url(../assets/images/cards-mobile.png);
  }
}
</style>