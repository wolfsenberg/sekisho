import { expect, it } from 'vitest'
import { formatDuration } from './duration'

it.each([
  [0, '0 sec'],
  [59, '59 sec'],
  [60, '1 min'],
  [3661, '1 hr 1 min'],
  [90061, '1 day 1 hr'],
])('formats %i seconds into at most two useful units', (seconds, expected) => {
  expect(formatDuration(seconds, 'en')).toBe(expected)
})

it('uses localized duration units', () => {
  expect(formatDuration(90061, 'es')).toBe('1 d 1 h')
})
