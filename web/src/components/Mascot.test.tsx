import { act, fireEvent, render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { afterEach, beforeEach, expect, it } from 'vitest'
import { setMascotEnabled } from '@/lib/mascot'
import { Mascot, MascotSetting } from './Mascot'
import { sprites } from './mascotSprites'

beforeEach(() => {
  // Skip the once-per-visit greeting so each test starts from a quiet fox.
  sessionStorage.setItem('sekisho-mascot-greeted', '1')
})

afterEach(() => {
  localStorage.clear()
  sessionStorage.clear()
})

const click = (element: HTMLElement) => {
  fireEvent.pointerDown(element, { clientX: 10, clientY: 10, pointerId: 1 })
  fireEvent.pointerUp(element, { clientX: 11, clientY: 10, pointerId: 1 })
}

it('draws every sprite frame on the same 16 by 14 pixel grid', () => {
  for (const [name, rows] of Object.entries(sprites)) {
    expect(rows, name).toHaveLength(14)
    for (const row of rows) expect(row, name).toHaveLength(16)
  }
})

it('shows the fox by default', () => {
  render(<Mascot />)
  expect(screen.getByRole('img', { name: /Seki the fox/ })).toBeTruthy()
})

it('stays hidden when the person turned it off', () => {
  localStorage.setItem('sekisho-mascot', 'off')
  render(<Mascot />)
  expect(screen.queryByRole('img', { name: /Seki the fox/ })).toBeNull()
})

it('introduces itself when clicked, then jumps for joy on the next click', async () => {
  render(<Mascot />)
  const fox = screen.getByRole('img', { name: /Seki the fox/ })
  click(fox)
  expect(screen.getByText('Hi, my name is Seki!')).toBeTruthy()
  expect(fox.closest('[data-mascot]')?.getAttribute('data-mood')).toBe('hi')
  click(fox)
  // It crouches first, then leaps.
  const mood = () => fox.closest('[data-mascot]')?.getAttribute('data-mood')
  expect(mood()).toBe('crouch')
  await waitFor(() => expect(mood()).toBe('jump'))
  // After landing it goes back to living its life instead of freezing.
  await waitFor(() => expect(mood()).not.toBe('jump'), { timeout: 3000 })
  await waitFor(() => expect(mood()).not.toBe('idle'), { timeout: 6000 })
}, 15000)

it('greets once per visit when it first appears', async () => {
  sessionStorage.clear()
  render(<Mascot />)
  expect(
    await screen.findByText('Hi, my name is Seki!', {}, { timeout: 3000 }),
  ).toBeTruthy()
  expect(sessionStorage.getItem('sekisho-mascot-greeted')).toBe('1')
})

it('hides and shows the fox from the preference switch', async () => {
  const user = userEvent.setup()
  render(
    <>
      <MascotSetting />
      <Mascot />
    </>,
  )
  const toggle = screen.getByRole('switch', { name: 'Show Seki' })
  expect((toggle as HTMLInputElement).checked).toBe(true)
  await user.click(toggle)
  expect(screen.queryByRole('img', { name: /Seki the fox/ })).toBeNull()
  expect(localStorage.getItem('sekisho-mascot')).toBe('off')
  act(() => setMascotEnabled(true))
  expect(screen.getByRole('img', { name: /Seki the fox/ })).toBeTruthy()
})
