<template>
  <div :class="$style.card" :style="style" @click="onClick">
    <slot name="info" />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref, PropType } from '@vue/composition-api';
import { Card, Suit } from '~/lib/models';

export default defineComponent({
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

  setup(props, { emit }) {
    const selected = ref(false);

    const suitYmap: Record<Suit, number> = {
      [Suit.Clubs]: 0,
      [Suit.Diamonds]: 120,
      [Suit.Hearts]: 240,
      [Suit.Spades]: 360,
    };
    const offsetX = computed(() => {
      if (!props.showFace) return 0;
      return 80 * props.card.faceValue;
    });
    const offsetY = computed(() => {
      if (!props.showFace) return [Suit.Spades, Suit.Clubs].includes(props.card.suit) ? 0 : 120;
      return suitYmap[props.card.suit];
    });
    const style = computed<Record<string, string>>(() => {
      return {
        cursor: props.selectable ? 'pointer' : 'default',
        backgroundPosition: `top -${offsetY.value}px left -${offsetX.value}px`,
      };
    });

    const onClick = () => {
      if (!props.selectable) return;
      selected.value = !selected.value;
      emit('selected', selected.value);
    };

    return { style, onClick };
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
