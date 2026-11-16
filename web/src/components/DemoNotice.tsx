import { useMutationState, useQuery } from '@tanstack/react-query'
import { FlaskConical } from 'lucide-react'
import { useTranslation } from 'react-i18next'
import { authOptions } from '@/lib/auth'
import { ApiError } from '@/lib/request'

export function DemoNotice() {
  const { t } = useTranslation()
  const { data } = useQuery(authOptions())
  const restrictions = useMutationState({
    filters: { status: 'error' },
    select: (mutation) =>
      mutation.state.error instanceof ApiError &&
      mutation.state.error.code === 'demo_read_only',
  })
  if (!data?.demo) return null
  const blocked = restrictions.some(Boolean)
  return (
    <aside
      aria-label={t('demoTitle')}
      className="flex items-center gap-2.5 border-b border-border/70 bg-card/60 px-4 py-2 text-[0.8125rem] leading-5 md:px-6"
    >
      <FlaskConical
        aria-hidden="true"
        className="size-3.5 shrink-0 text-muted-foreground"
      />
      <p className="min-w-0">
        <span className="font-medium">{t('demoTitle')}</span>{' '}
        <span
          role={blocked ? 'alert' : undefined}
          className="text-muted-foreground"
        >
          {t(blocked ? 'demoReadOnly' : 'demoDescription')}
        </span>
      </p>
    </aside>
  )
}
