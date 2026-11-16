import {
  useEffect,
  useEffectEvent,
  useRef,
  useState,
  type PointerEvent,
} from 'react'
import { useTranslation } from 'react-i18next'
import { setMascotEnabled, useMascotEnabled } from '@/lib/mascot'
import { sprites, type Sprite } from './mascotSprites'

/*
 * Seki, the checkpoint fox: a 16×14 pixel sprite that lives along the bottom of
 * the dashboard. Foxes guard gates in Japanese folklore, which suits a gateway.
 * Position is written straight to the element every frame; React renders only
 * when the sprite, facing, mood or speech changes, a few times a second at most.
 */

const fills: Record<string, string> = {
  R: 'var(--seal)',
  W: 'var(--mascot-chest)',
  O: 'color-mix(in oklab, var(--seal) 55%, black)',
  E: '#111113',
}

// Merge horizontal runs into one rect each so a frame is ~40 rects, not ~150.
function Pixels({ rows }: { rows: string[] }) {
  const rects: React.ReactNode[] = []
  rows.forEach((row, y) => {
    let x = 0
    while (x < row.length) {
      const ch = row[x]
      let end = x + 1
      while (end < row.length && row[end] === ch) end++
      if (ch !== '.')
        rects.push(
          <rect
            key={`${x}-${y}`}
            x={x}
            y={y}
            width={end - x}
            height={1}
            fill={fills[ch]}
          />,
        )
      x = end
    }
  })
  return <>{rects}</>
}

const scale = 3
const width = 16 * scale
const height = 14 * scale
const gap = 12
const walkSpeed = 42 // px per second
const runSpeed = 130
const gravity = 2200 // px per second²
const greetedKey = 'sekisho-mascot-greeted'

type Mode =
  | 'idle'
  | 'walk'
  | 'run'
  | 'sit'
  | 'rest'
  | 'sleep'
  | 'wake'
  | 'hi'
  | 'crouch'
  | 'jump'
  | 'held'
  | 'fall'
const airborne = (mode: Mode) => mode === 'jump' || mode === 'fall'
// Moods the small mind may leave on its own; scripted moods finish first.
const free: Mode[] = ['idle', 'walk', 'run', 'sit', 'rest']

function pick<T>(choices: [T, number][]): T {
  let roll = Math.random()
  for (const [choice, weight] of choices) {
    if ((roll -= weight) <= 0) return choice
  }
  return choices[choices.length - 1][0]
}

export function Mascot() {
  const enabled = useMascotEnabled()
  return enabled ? <Fox /> : null
}

function Fox() {
  const { t } = useTranslation()
  const element = useRef<HTMLDivElement>(null)
  const position = useRef({ x: 0, y: 0, vx: 0, vy: 0 })
  const drag = useRef({ dx: 0, dy: 0, startX: 0, startY: 0, moved: false })
  const modeRef = useRef<Mode>('idle')
  const facingRef = useRef<1 | -1>(-1)
  const script = useRef<number[]>([])
  const reduce = useRef(false)
  const [mode, setModeState] = useState<Mode>('idle')
  const [facing, setFacing] = useState<1 | -1>(-1)
  const [frame, setFrame] = useState(0)
  const [blink, setBlink] = useState(false)
  const [heart, setHeart] = useState(false)
  const [speech, setSpeech] = useState<{
    text: string
    align: 'start' | 'center' | 'end'
  } | null>(null)

  const setMode = (next: Mode) => {
    modeRef.current = next
    setModeState(next)
  }
  const turn = (next: 1 | -1) => {
    facingRef.current = next
    setFacing(next)
  }
  const floor = () => window.innerHeight - height - gap
  const place = () => {
    const { x, y } = position.current
    if (element.current)
      element.current.style.transform = `translate3d(${x}px, ${y}px, 0)`
  }
  const stopScript = () => {
    script.current.forEach((timer) => window.clearTimeout(timer))
    script.current = []
  }
  // Run a short sequence of moods, e.g. yawn, then sit, then idle.
  const play = (steps: [Mode, number][], then?: () => void) => {
    stopScript()
    let at = 0
    for (const [step, duration] of steps) {
      if (at === 0) setMode(step)
      else script.current.push(window.setTimeout(() => setMode(step), at))
      at += duration
    }
    script.current.push(
      window.setTimeout(() => {
        script.current = []
        if (then) then()
        else setMode('idle')
      }, at),
    )
  }
  const say = (text: string, duration: number) => {
    const x = position.current.x
    const align =
      x < 80 ? 'start' : x > window.innerWidth - width - 80 ? 'end' : 'center'
    setSpeech({ text, align })
    script.current.push(window.setTimeout(() => setSpeech(null), duration))
  }
  const leap = (vy: number, vx: number) => {
    position.current.vy = vy
    position.current.vx = vx
    setMode('jump')
  }
  const sayHello = () => {
    turn(position.current.x > window.innerWidth / 2 ? -1 : 1)
    play([['hi', 2600]], () => setMode('sit'))
    say(t('mascotHello'), 2600)
  }
  const wakeUp = (greet: boolean) => {
    play([
      ['wake', 1000],
      ['sit', 600],
    ])
    if (greet) say(t('mascotWake'), 1600)
  }
  const jumpForJoy = () => {
    setHeart(true)
    window.setTimeout(() => setHeart(false), 1100)
    stopScript()
    setSpeech(null)
    if (reduce.current) return
    setMode('crouch')
    script.current.push(
      window.setTimeout(() => {
        // The sequence is over once it leaps; free the small mind again.
        script.current = []
        leap(-560, 0)
      }, 140),
    )
  }

  // Effects call the latest handlers without re-running when they change.
  const greet = useEffectEvent(() => {
    // Mark the visit greeted only once the greeting really happens, so a remount
    // before the timer fires still gets its hello.
    try {
      sessionStorage.setItem(greetedKey, '1')
    } catch {
      // Without session storage it may greet again on the next visit.
    }
    sayHello()
  })
  const wakeQuietly = useEffectEvent(() => wakeUp(false))
  const hop = useEffectEvent(() => leap(-520, facingRef.current * 90))

  // Start near the bottom right, out of the way of page actions.
  useEffect(() => {
    reduce.current =
      typeof window.matchMedia === 'function' &&
      window.matchMedia('(prefers-reduced-motion: reduce)').matches
    position.current.x = Math.max(0, window.innerWidth - width - 96)
    position.current.y = floor()
    place()
    const onResize = () => {
      const p = position.current
      p.x = Math.min(Math.max(0, p.x), window.innerWidth - width)
      if (modeRef.current !== 'held' && !airborne(modeRef.current))
        p.y = floor()
      place()
    }
    window.addEventListener('resize', onResize)
    // Say hello once per visit, shortly after arriving.
    let greeting = 0
    try {
      if (!sessionStorage.getItem(greetedKey))
        greeting = window.setTimeout(greet, 1200)
    } catch {
      // Without session storage the fox simply stays quiet.
    }
    return () => {
      window.removeEventListener('resize', onResize)
      window.clearTimeout(greeting)
      stopScript()
    }
  }, [])

  // Physics, walking and running: an animation frame loop that runs only while
  // the fox is moving, so a resting or sleeping fox costs nothing per frame.
  const moving = mode === 'walk' || mode === 'run' || airborne(mode)
  useEffect(() => {
    if (!moving) return
    let raf = 0
    let last = performance.now()
    const tick = (now: number) => {
      const dt = Math.min(0.05, (now - last) / 1000)
      last = now
      const p = position.current
      const current = modeRef.current
      if (current === 'walk' || current === 'run') {
        const speed = current === 'run' ? runSpeed : walkSpeed
        p.x += facingRef.current * speed * dt
        if (p.x <= 0 || p.x >= window.innerWidth - width) {
          p.x = Math.min(Math.max(0, p.x), window.innerWidth - width)
          turn(facingRef.current === 1 ? -1 : 1)
        }
        place()
      } else if (airborne(current)) {
        p.vy += gravity * dt
        p.y += p.vy * dt
        p.x = Math.min(Math.max(0, p.x + p.vx * dt), window.innerWidth - width)
        if (p.y >= floor()) {
          p.y = floor()
          p.vy = 0
          p.vx = 0
          setMode('idle')
        }
        place()
      }
      raf = requestAnimationFrame(tick)
    }
    raf = requestAnimationFrame(tick)
    return () => cancelAnimationFrame(raf)
  }, [moving])

  // Sprite frames: a two-step walk, a quicker gallop, a slow wave.
  useEffect(() => {
    const rate =
      mode === 'walk' ? 170 : mode === 'run' ? 90 : mode === 'hi' ? 280 : 0
    if (!rate || reduce.current) return
    const timer = window.setInterval(() => setFrame((value) => value + 1), rate)
    return () => window.clearInterval(timer)
  }, [mode])

  // Blink now and then while awake.
  useEffect(() => {
    if (mode === 'sleep' || mode === 'held' || mode === 'wake') return
    let timer = 0
    const schedule = () => {
      timer = window.setTimeout(
        () => {
          setBlink(true)
          timer = window.setTimeout(() => {
            setBlink(false)
            schedule()
          }, 140)
        },
        2500 + Math.random() * 3500,
      )
    }
    schedule()
    return () => window.clearTimeout(timer)
  }, [mode])

  // A small mind: wander, sit, get the zoomies, hop, and doze off.
  useEffect(() => {
    if (reduce.current) return
    if (mode === 'sleep') {
      // A long nap, then it wakes up by itself without fuss.
      const timer = window.setTimeout(
        wakeQuietly,
        25000 + Math.random() * 20000,
      )
      return () => window.clearTimeout(timer)
    }
    if (!free.includes(mode) || script.current.length) return
    const wait =
      mode === 'walk'
        ? 3000 + Math.random() * 4000
        : mode === 'run'
          ? 1400 + Math.random() * 1200
          : mode === 'rest'
            ? 5000 + Math.random() * 5000
            : 1800 + Math.random() * 2500
    const timer = window.setTimeout(() => {
      const next = pick<Mode | 'hop'>(
        mode === 'idle'
          ? [
              ['walk', 0.5],
              ['sit', 0.18],
              ['run', 0.12],
              ['hop', 0.1],
              ['rest', 0.1],
            ]
          : mode === 'walk'
            ? [
                ['idle', 0.45],
                ['sit', 0.2],
                ['run', 0.17],
                ['hop', 0.18],
              ]
            : mode === 'run'
              ? [
                  ['idle', 0.65],
                  ['hop', 0.35],
                ]
              : mode === 'sit'
                ? [
                    ['idle', 0.5],
                    ['rest', 0.5],
                  ]
                : [
                    ['sleep', 0.6],
                    ['sit', 0.4],
                  ],
      )
      if ((next === 'walk' || next === 'run') && Math.random() < 0.4)
        turn(facingRef.current === 1 ? -1 : 1)
      if (next === 'hop') {
        setMode('crouch')
        script.current.push(
          window.setTimeout(() => {
            script.current = []
            hop()
          }, 160),
        )
      } else setMode(next)
    }, wait)
    return () => window.clearTimeout(timer)
  }, [mode])

  const onPointerDown = (event: PointerEvent<HTMLDivElement>) => {
    event.currentTarget.setPointerCapture?.(event.pointerId)
    const p = position.current
    drag.current = {
      dx: event.clientX - p.x,
      dy: event.clientY - p.y,
      startX: event.clientX,
      startY: event.clientY,
      moved: false,
    }
  }
  const onPointerMove = (event: PointerEvent<HTMLDivElement>) => {
    const d = drag.current
    if (!event.buttons && !d.moved) return
    if (
      !d.moved &&
      Math.hypot(event.clientX - d.startX, event.clientY - d.startY) > 4
    ) {
      d.moved = true
      stopScript()
      setSpeech(null)
      setMode('held')
    }
    if (!d.moved) return
    const p = position.current
    const nextX = event.clientX - d.dx
    if (nextX !== p.x) turn(nextX > p.x ? 1 : -1)
    p.x = Math.min(Math.max(0, nextX), window.innerWidth - width)
    p.y = Math.min(Math.max(0, event.clientY - d.dy), floor())
    place()
  }
  const onPointerUp = () => {
    const d = drag.current
    if (!d.moved) {
      const current = modeRef.current
      if (current === 'sleep' || current === 'rest') wakeUp(true)
      else if (current === 'hi' || current === 'sit') jumpForJoy()
      else sayHello()
      return
    }
    d.moved = false
    if (reduce.current) {
      position.current.y = floor()
      place()
      setMode('idle')
      return
    }
    position.current.vy = 0
    position.current.vx = 0
    setMode('fall')
  }

  const sprite: Sprite = (() => {
    switch (mode) {
      case 'walk':
        return frame % 2 ? 'step' : 'stand'
      case 'run':
        return frame % 2 ? 'runIn' : 'runOut'
      case 'hi':
        return frame % 2 ? 'waveDown' : 'waveUp'
      case 'sit':
        return blink ? 'sitBlink' : 'sit'
      case 'wake':
        return 'yawn'
      case 'rest':
      case 'crouch':
        return 'crouch'
      case 'sleep':
        return 'curl'
      case 'jump':
      case 'fall':
      case 'held':
        return 'runOut'
      default:
        return blink ? 'blink' : 'stand'
    }
  })()

  return (
    <div
      ref={element}
      data-mascot=""
      data-mood={mode}
      className="pointer-events-none fixed top-0 left-0 z-40 select-none print:hidden"
      style={{ width, height }}
    >
      {speech && (
        <p
          aria-hidden="true"
          className={
            'mascot-speech absolute bottom-full mb-3 w-max max-w-56 border-2 border-foreground bg-popover px-2 py-1 font-mono text-[0.6875rem] leading-4 text-foreground ' +
            (speech.align === 'start'
              ? 'left-0'
              : speech.align === 'end'
                ? 'right-0'
                : 'left-1/2 -translate-x-1/2')
          }
        >
          {speech.text}
        </p>
      )}
      {heart && (
        <svg
          aria-hidden="true"
          viewBox="0 0 7 6"
          width={21}
          height={18}
          shapeRendering="crispEdges"
          className="mascot-heart absolute -top-6 left-1/2"
        >
          <Pixels
            rows={[
              '.RR.RR.',
              'RRRRRRR',
              'RRRRRRR',
              '.RRRRR.',
              '..RRR..',
              '...R...',
            ]}
          />
        </svg>
      )}
      {mode === 'sleep' && (
        <span
          aria-hidden="true"
          className="mascot-z absolute top-2 right-0 font-mono text-xs font-medium text-muted-foreground"
        >
          z
        </span>
      )}
      <div
        role="img"
        aria-label={t('mascotLabel')}
        title={t('mascotLabel')}
        onPointerDown={onPointerDown}
        onPointerMove={onPointerMove}
        onPointerUp={onPointerUp}
        onPointerCancel={onPointerUp}
        className={
          'pointer-events-auto h-full w-full touch-none ' +
          (mode === 'held' ? 'cursor-grabbing -rotate-6' : 'cursor-grab')
        }
      >
        <svg
          viewBox="0 0 16 14"
          width={width}
          height={height}
          shapeRendering="crispEdges"
          className={mode === 'sleep' ? 'mascot-breathe' : undefined}
          style={{ transform: facing === 1 ? undefined : 'scaleX(-1)' }}
        >
          <Pixels rows={sprites[sprite]} />
        </svg>
      </div>
    </div>
  )
}

export function MascotSetting() {
  const { t } = useTranslation()
  const enabled = useMascotEnabled()
  return (
    <label className="flex min-h-11 cursor-pointer items-center gap-2 text-sm">
      <input
        type="checkbox"
        role="switch"
        className="peer sr-only"
        checked={enabled}
        onChange={(event) => setMascotEnabled(event.target.checked)}
      />
      <span
        aria-hidden="true"
        className="relative h-5 w-9 shrink-0 rounded-full bg-muted-foreground/30 transition-colors duration-200 after:absolute after:left-0.5 after:top-0.5 after:size-4 after:rounded-full after:bg-white after:shadow-[0_1px_3px_rgb(0_0_0/0.35)] after:transition-transform after:duration-200 after:ease-(--ease-drawer) peer-checked:bg-primary peer-checked:after:translate-x-4 peer-focus-visible:ring-2 peer-focus-visible:ring-ring peer-focus-visible:ring-offset-2 peer-focus-visible:ring-offset-background motion-reduce:transition-none motion-reduce:after:transition-none"
      />
      <span>{t('mascotShow')}</span>
    </label>
  )
}
