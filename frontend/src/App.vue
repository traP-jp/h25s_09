<script setup lang="ts">
import ErrorToast from '@/components/ErrorToast.vue'
import TheFooter from '@/layouts/footer/TheFooter.vue'
import TheHeader from '@/layouts/header/TheHeader.vue'
import TheSidebar from '@/layouts/sidebar/TheSidebar.vue'
import { useBreakpoints } from '@vueuse/core'
import { RouterView } from 'vue-router'

const breakpoints = useBreakpoints({
  mobile: 0,
  tablet: 768,
})

const isDesktop = breakpoints.greaterOrEqual('tablet')
</script>

<template>
  <TheHeader />
  <TheSidebar v-if="isDesktop" />
  <TheFooter v-else />
  <RouterView />
  <ErrorToast />
</template>
