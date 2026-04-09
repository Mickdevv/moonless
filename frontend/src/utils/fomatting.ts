// export function formatPrice(amount: number) {
//   return `$${(amount / 100).toFixed(2)}`
// }
//
export function formatDateTime(date: Date | string) {
  let t = date
  if (typeof t == 'string') {
    t = new Date(t)
  }
  return `${t.toLocaleDateString('fr-FR', {
    day: '2-digit',
    month: 'long',
    year: 'numeric',
    hour: '2-digit',
    minute: '2-digit',
  })}`
}
export function formatDate(date: Date | string) {
  let t = date
  if (typeof t == 'string') {
    t = new Date(t)
  }
  return t.toLocaleDateString()
}
