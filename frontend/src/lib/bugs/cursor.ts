import { ref } from 'vue'
import { randomExponential } from 'd3-random'

const isCursorLoading = ref(false)

export const useCursorLoading = () => {
  return { isCursorLoading: isCursorLoading }
}

const toggleCursorLoading = () => {
  isCursorLoading.value = !isCursorLoading.value

  const delay = isCursorLoading.value
    ? randomExponential(1 / 1000)()
    : randomExponential(1 / 5000)()
  setTimeout(toggleCursorLoading, delay)
}

toggleCursorLoading()
