import type { Meta, StoryObj } from '@storybook/vue3'

import InputDialog from '../components/InputDialog.vue'

const meta = {
  title: 'InputDialog',
  component: InputDialog,
  tags: ['autodocs'],
  argTypes: {
    title: { control: 'string' },
    message: { control: 'string' },
    default: { control: 'string' },
    confirmButtonText: { control: 'string' },
    cancelButtonText: { control: 'string' },
    onConfirm: { action: 'clicked' },
  },
  args: { },
} satisfies Meta<typeof InputDialog>

export default meta
type Story = StoryObj<typeof meta>

export const DefaultButtonTexts: Story = {
  args: {
    title: 'Some Title',
    message: 'Some prompt for giggles',
  },
};

export const NonDefaultButtonTexts: Story = {
  args: {
    title: 'Some Title',
    message: 'Some prompt for giggles',
    confirmButtonText: 'Confirm!',
    cancelButtonText: 'Cancel!',
  },
};

export const DefaultValue: Story = {
  args: {
    title: 'Some Title',
    message: 'Some prompt for giggles',
    default: 'a default value',
  },
};
