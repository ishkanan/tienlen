<template>
  <transition name="fade">
    <div id="vueConfirm" :class="$style.overlay" @click="handleClickOverlay">
      <transition name="zoom">
        <div :class="$style.container">
          <span :class="$style.textgrid">
            <h4 v-if="title" :class="$style.title">{{ title }}</h4>
            <p v-if="message" :class="$style.text">{{ message }}</p>
          </span>
          <div :class="$style.btngrid">
            <button
              :class="[$style.btn, $style.btnleft]"
              @click.stop="(e) => handleClickButton(e, true)"
            >
              {{ yesButtonText }}
            </button>
            <button :class="$style.btn" @click.stop="(e) => handleClickButton(e, false)">
              {{ noButtonText }}
            </button>
          </div>
        </div>
      </transition>
    </div>
  </transition>
</template>

<script lang="ts">
import { defineComponent } from '@vue/composition-api';

interface event {
  target?: {
    id: string;
  };
}

export default defineComponent({
  props: {
    title: {
      type: String,
      default: 'Confirm',
    },
    message: {
      type: String,
      default: '',
    },
    yesButtonText: {
      type: String,
      default: 'Yes',
    },
    noButtonText: {
      type: String,
      default: 'No',
    },
  },

  setup(_, { emit }) {
    const handleClickButton = (event: event, confirm: boolean) => {
      if (event.target?.id === 'vueConfirm') return;
      emit('confirm', confirm);
    };

    const handleClickOverlay = (event: event) => {
      if (event.target?.id === 'vueConfirm') emit('confirm', false);
    };

    return { handleClickButton, handleClickOverlay };
  },
});
</script>

<style lang="postcss" module>
:root {
  --title-color: black;
  --message-color: black;
  --overlay-background-color: #0000004a;
  --container-box-shadow: #0000004a 0px 3px 8px 0px;
  --base-background-color: #ffffff;
  --button-color: black;
  --button-background-color: #ffffff;
  --button-border-color: #e0e0e0;
  --button-background-color-disabled: #f5f5f5;
  --button-background-color-hover: #f5f5f5;
  --button-box-shadow-active: inset 0 2px 0px 0px #00000014;
  --input-background-color: #ebebeb;
  --input-background-color-hover: #dfdfdf;
  --font-size-m: 24px;
  --font-size-s: 18px;
  --font-weight-black: 1100;
  --font-weight-bold: 600;
  --font-weight-medium: 500;
  --font-weight-normal: 400;
  --font-weight-light: 300;
}
/**
* Dialog
*/
.overlay *,
.overlay *:before,
.overlay *:after {
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  text-decoration: none;
  -webkit-touch-callout: none;
  -moz-osx-font-smoothing: grayscale;
  margin: 0;
  padding: 0;
}
.title {
  color: var(--title-color);
  padding: 0 1rem;
  width: 100%;
  font-weight: var(--font-weight-black);
  text-align: center;
  font-size: var(--font-size-m);
  line-height: initial;
  margin-bottom: 15px;
}
.text {
  color: var(--message-color);
  padding: 0 1rem;
  width: 100%;
  font-weight: var(--font-weight-medium);
  text-align: center;
  font-size: var(--font-size-s);
  line-height: initial;
}
.overlay {
  background-color: var(--overlay-background-color);
  width: 100%;
  height: 100%;
  transition: all 0.1s ease-in;
  left: 0;
  top: 0;
  z-index: 999999999999;
  position: fixed;
  display: flex;
  justify-content: center;
  align-items: center;
  align-content: baseline;
}
.container {
  background-color: var(--base-background-color);
  border-radius: 1rem;
  width: 400px;
  height: auto;
  display: grid;
  grid-template-rows: 1fr max-content;
  box-shadow: var(--container-box-shadow);
}
.textgrid {
  padding: 1.2rem;
}
.btngrid {
  width: 100%;
  display: grid;
  grid-template-columns: 1fr 1fr;
  border-radius: 0 0 1rem 1rem;
  overflow: hidden;
}
.btn {
  border-radius: 0 0 1rem 0;
  color: var(--button-color);
  background-color: var(--button-background-color);
  border: 0;
  font-size: 1.1rem;
  border-top: 1px solid var(--button-border-color);
  cursor: pointer;
  outline: none;
  min-height: 50px;
}
.btn:hover {
  background-color: var(--button-background-color-hover);
}
.btn:disabled {
  background-color: var(--button-background-color-disabled);
}
.btn:active {
  box-shadow: var(--button-box-shadow-active);
}
.btnleft {
  border-radius: 0;
  border-right: 1px solid var(--button-border-color);
}
/**
* Transition
*/
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.21s;
}
.fade-enter,
.fade-leave-to {
  opacity: 0;
}
.zoom-enter-active,
.zoom-leave-active {
  animation-duration: 0.21s;
  animation-fill-mode: both;
  animation-name: zoom;
}
.zoom-leave-active {
  animation-direction: reverse;
}
@keyframes zoom {
  from {
    opacity: 0;
    transform: scale3d(1.1, 1.1, 1.1);
  }
  100% {
    opacity: 1;
    transform: scale3d(1, 1, 1);
  }
}
</style>
