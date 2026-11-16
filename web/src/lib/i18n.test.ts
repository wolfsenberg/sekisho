import { expect, it, vi } from 'vitest'
import { waitFor } from '@testing-library/react'

const loadSpanish = vi.hoisted(() => vi.fn())
vi.mock('@/locales/es.json', async (original) => {
  loadSpanish()
  return original()
})

it('does not load the Spanish dictionary for an English session', async () => {
  localStorage.setItem('sekisho-language', 'en')
  vi.resetModules()
  loadSpanish.mockClear()
  const { i18n } = await import('./i18n')
  await waitFor(() => expect(i18n.isInitialized).toBe(true))
  expect(i18n.resolvedLanguage).toBe('en')
  expect(loadSpanish).not.toHaveBeenCalled()
})

it.each(['es', 'ja', 'ar'])(
  'restores %s before completing initialization',
  async (language) => {
    localStorage.setItem('sekisho-language', language)
    vi.resetModules()
    const { i18n } = await import('./i18n')
    await waitFor(() => expect(i18n.resolvedLanguage).toBe(language))
    expect(document.documentElement.lang).toBe(language)
    expect(document.documentElement.dir).toBe(language === 'ar' ? 'rtl' : 'ltr')
  },
)

it('falls back to English for an unsupported stored language', async () => {
  localStorage.setItem('sekisho-language', 'not-a-language')
  vi.resetModules()
  const { i18n } = await import('./i18n')
  if (!i18n.isInitialized)
    await new Promise((resolve) => i18n.on('initialized', resolve))
  expect(i18n.resolvedLanguage).toBe('en')
})
