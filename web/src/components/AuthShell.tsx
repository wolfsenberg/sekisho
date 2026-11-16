import type { ReactNode } from 'react'
import { LanguageSelect, ThemeSelect } from './Preferences'
import { Logo } from './Logo'

/*
 * Sign-in and setup present a travel pass (tegata): a quiet ruled panel that the
 * checkpoint seal stamps once on arrival. The ground carries faint horizontal
 * ruling of a ledger page, fading out above and below so it never competes with fields.
 */
export function AuthShell({ children }: { children: ReactNode }) {
  return (
    <div className="relative isolate flex min-h-svh flex-col overflow-hidden">
      <div
        aria-hidden="true"
        className="pointer-events-none absolute inset-0 -z-10 bg-[linear-gradient(to_bottom,var(--hairline)_1px,transparent_1px)] [mask-image:radial-gradient(ellipse_55%_50%_at_50%_48%,black,transparent_80%)] bg-[size:100%_44px] opacity-30"
      />
      <header className="flex min-h-16 items-center justify-between gap-3 px-5 sm:px-8">
        <div className="flex items-center gap-2.5">
          <Logo className="size-7" />
          <span className="text-[0.9375rem] font-semibold tracking-[-0.02em]">
            Sekisho
          </span>
        </div>
        <div className="flex items-center gap-1">
          <LanguageSelect compact />
          <ThemeSelect compact />
        </div>
      </header>
      <main className="flex flex-1 items-center justify-center px-5 pt-6 pb-16">
        <section className="tegata relative w-full max-w-[26rem] rounded-xl border border-hairline bg-card px-6 pt-8 pb-7 shadow-[0_24px_60px_-28px_rgb(60_40_20/0.28),0_1px_3px_0_rgb(60_40_20/0.08)] dark:shadow-[0_30px_80px_-30px_rgb(0_0_0/0.55),0_2px_8px_-2px_rgb(0_0_0/0.25)] sm:px-8 sm:pt-10 sm:pb-8">
          <Logo className="seal-stamp-slow absolute -top-5 end-6 size-12 sm:end-8" />
          {children}
        </section>
      </main>
    </div>
  )
}
