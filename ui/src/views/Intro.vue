<template>
  <div :class="$style.viewport">
    <h1>Welcome to Tiến lên online!</h1>
    <div :class="$style.controls">
      <input v-model="name" type="text" maxlength="35" placeholder="Enter your name..." />
      <button :disabled="!canSubmit" @click="handleConnect">Connect</button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, ref } from '@vue/composition-api';
import { joinGame } from '~/lib/socket';
import { setTitle } from '~/lib/utils';
import { game, ConnectionState } from '~/store/game';

export default defineComponent({
  setup() {
    const name = ref('');
    const nameLatin1Check = /[\u0020-\u007e\u00C0-\u00ff]+/g;
    const validName = computed(() => {
      const matches = name.value.match(nameLatin1Check);
      return (matches ?? []).length === 1 || name.value.length === 0;
    });
    const disconnected = computed(() => game.connState === ConnectionState.NotConnected);
    const canSubmit = computed(() => validName.value && disconnected.value);

    const handleConnect = () => {
      game.name = name.value;
      joinGame({ name: name.value });
    };

    return { name, canSubmit, handleConnect };
  },

  mounted() {
    setTitle('Tiến lên (Thirteen)');
    game.events = [];
  },
});
</script>

<style lang="postcss" module>
.viewport {
  width: 600px;
  height: 200px;
  background-color: rgba(48, 112, 16, 0.7);
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
  border: 4mm ridge rgba(170, 50, 50, 0.6);
  border-radius: 8%;
  display: flex;
  flex-direction: column;
  justify-content: space-evenly;
  align-items: center;
  padding: 1%;

  & h1 {
    color: white;
    margin: 0;
  }

  & .controls {
    display: flex;
    flex-direction: row;
    justify-content: center;
    width: 90%;
  }

  & input {
    width: 75%;
    border-radius: 15px;
    padding: 10px;
    font-size: 18px;
    margin-right: 20px;
    transition: all ease-in-out 0.2s;
  }
}
</style>
