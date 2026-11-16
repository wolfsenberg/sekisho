import { useSyncExternalStore } from 'react'

// A per-browser convenience: whether Seki the pixel fox wanders the dashboard.
const key = 'sekisho-mascot'
const listeners = new Set<() => void>()
// Used only when storage is blocked, so the switch still works for this page.
let fallback = true

export function mascotEnabled() {
  try {
    return localStorage.getItem(key) !== 'off'
  } catch {
    return fallback
  }
}

export function setMascotEnabled(enabled: boolean) {
  fallback = enabled
  try {
    if (enabled) localStorage.removeItem(key)
    else localStorage.setItem(key, 'off')
  } catch {
    // Storage can be blocked; the choice then lasts for this page only.
  }
  listeners.forEach((listener) => listener())
}

function subscribe(listener: () => void) {
  listeners.add(listener)
  const onStorage = (event: StorageEvent) => {
    if (event.key === key) listener()
  }
  window.addEventListener('storage', onStorage)
  return () => {
    listeners.delete(listener)
    window.removeEventListener('storage', onStorage)
  }
}

export function useMascotEnabled() {
  return useSyncExternalStore(subscribe, mascotEnabled, () => true)
}
