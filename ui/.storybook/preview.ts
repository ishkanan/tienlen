import type { Preview } from '@storybook/vue3'

import '../src/assets/css/global.css'
import '../src/assets/css/bulma-0.9.4.min.css'

const preview: Preview = {
  parameters: {
    actions: { argTypesRegex: '^on[A-Z].*' },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/,
      },
    },
  },
}

export default preview
