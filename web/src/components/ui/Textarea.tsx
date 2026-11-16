import type { ComponentProps } from 'react'
import { cn } from '@/lib/cn'

export function Textarea({ className, ...props }: ComponentProps<'textarea'>) {
  return (
    <textarea
      className={cn(
        'w-full min-w-0 rounded-lg border border-transparent bg-control px-3 py-2 text-base outline-none transition-[border-color,box-shadow,background-color] duration-150 placeholder:text-muted-foreground hover:bg-control-hover focus-visible:border-ring focus-visible:bg-control focus-visible:ring-[3px] focus-visible:ring-ring/25 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm',
        className,
      )}
      {...props}
    />
  )
}
