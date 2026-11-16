import { cn } from '@/lib/cn'

// Canonical geometry lives in docs/assets/logo.svg and logo-small.svg.
const framed =
  'M38 16H218A22 22 0 0 1 240 38V218A22 22 0 0 1 218 240H38A22 22 0 0 1 16 218V38A22 22 0 0 1 38 16ZM48 34A14 14 0 0 0 34 48V208A14 14 0 0 0 48 222H208A14 14 0 0 0 222 208V48A14 14 0 0 0 208 34ZM48 42H208A6 6 0 0 1 214 48V208A6 6 0 0 1 208 214H48A6 6 0 0 1 42 208V48A6 6 0 0 1 48 42ZM170 69H102A31 31 0 0 0 71 100V108A31 31 0 0 0 102 139H154A9 9 0 0 1 163 148V156A9 9 0 0 1 154 165H82V187H154A31 31 0 0 0 185 156V148A31 31 0 0 0 154 117H102A9 9 0 0 1 93 108V100A9 9 0 0 1 102 91H170Z'
// Below 24px the frame channel closes up, so small marks drop it.
const compact =
  'M38 16H218A22 22 0 0 1 240 38V218A22 22 0 0 1 218 240H38A22 22 0 0 1 16 218V38A22 22 0 0 1 38 16ZM178 59H98A37 37 0 0 0 61 96V108A37 37 0 0 0 98 145H158A9 9 0 0 1 167 154V162A9 9 0 0 1 158 171H76V197H158A35 35 0 0 0 193 162V154A35 35 0 0 0 158 119H98A11 11 0 0 1 87 108V96A11 11 0 0 1 98 85H178Z'

/** The Sekisho seal: a vermilion hanko with a seal-script S cut through it. */
export function Logo({
  className = 'size-8',
  small = false,
}: {
  className?: string
  small?: boolean
}) {
  return (
    <svg
      viewBox="0 0 256 256"
      aria-hidden="true"
      className={cn('text-seal', className)}
    >
      <path
        fill="currentColor"
        fillRule="evenodd"
        d={small ? compact : framed}
      />
    </svg>
  )
}
