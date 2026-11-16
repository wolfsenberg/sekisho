import { useEffect, useRef } from 'react'
import { Link } from '@tanstack/react-router'
import { ArrowRight } from 'lucide-react'
import { useTranslation } from 'react-i18next'
import { AuthShell } from '@/components/AuthShell'
import { Button } from '@/components/ui/Button'

const steps = [
  ['setupAccountStep', 'setupAccountDetail'],
  ['setupWorkspaceStep', 'setupWorkspaceDetail'],
] as const

export function Setup() {
  const { t } = useTranslation()
  const heading = useRef<HTMLHeadingElement>(null)

  useEffect(() => {
    heading.current?.focus()
  }, [])

  return (
    <AuthShell>
      <h1
        ref={heading}
        tabIndex={-1}
        className="pe-16 text-2xl font-semibold tracking-[-0.025em] text-balance outline-none"
      >
        {t('setupWelcome')}
      </h1>
      <p className="mt-3 text-sm leading-6 text-muted-foreground">
        {t('setupWelcomeDescription')}
      </p>
      <ol className="my-8 space-y-5 border-y border-border py-6">
        {steps.map(([title, description], index) => (
          <li key={title} className="flex gap-4">
            <span
              className="mt-0.5 flex size-5 shrink-0 items-center justify-center rounded-sm border border-hairline font-mono text-[0.6875rem] text-muted-foreground"
              aria-hidden="true"
            >
              {index + 1}
            </span>
            <div>
              <h2 className="text-sm font-medium">{t(title)}</h2>
              <p className="mt-1 text-sm leading-6 text-muted-foreground">
                {t(description)}
              </p>
            </div>
          </li>
        ))}
      </ol>
      <Button asChild className="h-11 w-full">
        <Link to="/setup/admin">
          {t('startSetup')}
          <ArrowRight aria-hidden="true" />
        </Link>
      </Button>
    </AuthShell>
  )
}
