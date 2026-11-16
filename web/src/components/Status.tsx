import { cn } from '@/lib/cn'

// The word carries the state; colour only marks states that need attention.
const tones = {
  success: 'text-foreground',
  warning: 'text-warning',
  error: 'text-error',
  info: 'text-info',
  neutral: 'text-muted-foreground',
}

export function Status({
  kind,
  children,
}: {
  kind: keyof typeof tones
  children: React.ReactNode
}) {
  return (
    <span
      className={cn(
        'inline-flex w-fit items-center whitespace-nowrap font-mono text-xs leading-5 font-medium',
        tones[kind],
      )}
    >
      {children}
    </span>
  )
}
