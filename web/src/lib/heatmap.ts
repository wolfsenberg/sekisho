export const heatmapTones = [
  'bg-muted',
  'bg-seal/20',
  'bg-seal/40',
  'bg-seal/65',
  'bg-seal',
] as const

export function heatmapTone(value: number | null, maximum: number) {
  if (value === null) return 'bg-muted/40'
  return heatmapTones[
    value === 0 ? 0 : Math.min(4, Math.ceil((value / Math.max(1, maximum)) * 4))
  ]
}
