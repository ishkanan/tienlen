<script lang="ts">
import Vue from 'vue';

interface Data {
  windowHeight: number;
  windowWidth: number;
}

export default Vue.extend({
  data(): Data {
    return {
      windowHeight: window.innerHeight,
      windowWidth: window.innerWidth,
    };
  },

  beforeDestroy() {
    window.removeEventListener('resize', this.handleWindowResize);
  },

  mounted() {
    window.addEventListener('resize', this.handleWindowResize);
  },

  methods: {
    handleWindowResize() {
      this.windowHeight = window.innerHeight;
      this.windowWidth = window.innerWidth;
    },
  },
});
</script>