export function formatPrice(amount: number) {
  return `$${(amount / 100).toFixed(2)}`
}
