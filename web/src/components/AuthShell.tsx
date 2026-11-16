import type { ReactNode } from 'react'
import { LanguageSelect, ThemeSelect } from './Preferences'
import { Logo } from './Logo'

/*
 * Sign-in and setup present a travel pass (tegata): a quiet panel on a plain
 * ground that the checkpoint seal stamps once on arrival.
 */
export function AuthShell({ children }: { children: ReactNode }) {
  return (
    <div className="flex min-h-svh flex-col">
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
        <section className="tegata relative w-full max-w-[26rem] rounded-xl border border-hairline bg-card px-6 pt-8 pb-7 sm:px-8 sm:pt-10 sm:pb-8">
          <Logo className="seal-stamp-slow absolute -top-5 end-6 size-12 sm:end-8" />
          {children}
        </section>
      </main>
    </div>
  )
}
