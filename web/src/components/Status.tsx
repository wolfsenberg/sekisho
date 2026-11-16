import { CircleAlert, CirclePause, Info, TriangleAlert } from 'lucide-react'
import { cn } from '@/lib/cn'

// A passed check is stamped: the success badge carries a small vermilion seal.
function Seal({ className }: { className?: string }) {
  return (
    <span
      className={cn(
        'inline-block size-[7px] shrink-0 -rotate-6 rounded-[1.5px] bg-seal',
        className,
      )}
    />
  )
}

const variants = {
  // The seal alone carries approval, so the stamped chip itself stays neutral ink.
  success: { icon: Seal, color: 'text-foreground ring-hairline!' },
  warning: { icon: TriangleAlert, color: 'bg-warning-muted text-warning' },
  error: { icon: CircleAlert, color: 'bg-error-muted text-error' },
  info: { icon: Info, color: 'bg-info-muted text-info' },
  neutral: { icon: CirclePause, color: 'bg-muted text-muted-foreground' },
}

export function Status({
  kind,
  children,
}: {
  kind: keyof typeof variants
  children: React.ReactNode
}) {
  const { icon: Icon, color } = variants[kind]
  return (
    <span
      className={cn(
        'inline-flex w-fit items-center gap-1.5 whitespace-nowrap rounded-sm px-2 py-0.5 text-xs leading-5 font-medium ring-1 ring-current/15 ring-inset',
        color,
      )}
    >
      <Icon
        className={kind === 'success' ? 'mx-[3px]' : 'size-3.5'}
        aria-hidden="true"
      />
      {children}
    </span>
  )
}
