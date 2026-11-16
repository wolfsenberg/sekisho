// Seki's pixel sprites: every frame is 16 columns by 14 rows.
// R coat · W chest and tail tip · O legs (shaded coat) · E eye, nose and mouth.
const blank = '................'
const head = [
  '...........R..R.',
  '...........RRRR.',
  '..........RRRRRR',
  '..........RRRERR',
  '..........RRRRRE',
  '..........WWWW..',
  'RR.......RRRWW..',
  'RRR..RRRRRRRWW..',
  'WRRRRRRRRRRRRR..',
  'WWRRRRRRRRRRRR..',
  '.W.RRRRRRRRRRR..',
]
const sleepyHead = head.map((row, index) =>
  index === 3 ? '..........RRRRRR' : index === 4 ? '..........RRERRE' : row,
)
const sitting = [
  '........R..R....',
  '........RRRR....',
  '.......RRRRRR...',
  '.......RRRERR...',
  '.......RRRRRE...',
  '........WWWW....',
  '........RWWR....',
  'RR......RWWR....',
  'RRR....RRWWR....',
  'WRR...RRRWWR....',
  'WWRR.RRRRRRR....',
  '.WWRRRRRRRRR....',
  '.....RRRRRRO....',
  '.....OO..O.O....',
]
const edit = (rows: string[], changes: Record<number, string>) =>
  rows.map((row, index) => changes[index] ?? row)

export const sprites = {
  stand: [...head, '...O.O.....O.O..', '...O.O.....O.O..', '...O.O.....O.O..'],
  step: [...head, '....OO.....OO...', '...O..O...O..O..', '..O....O.O....O.'],
  blink: [
    ...sleepyHead,
    '...O.O.....O.O..',
    '...O.O.....O.O..',
    '...O.O.....O.O..',
  ],
  // The gallop: legs flung out, then gathered under the body.
  runOut: [...head, '..O.........OO..', '.O............O.', 'O...............'],
  runIn: [...head, '....OO....OO....', '.....OO..OO.....', blank],
  crouch: [...head, '...OOOOOOOOOO...', blank, blank],
  sit: sitting,
  sitBlink: edit(sitting, { 3: '.......RRRRRR...', 4: '.......RRERRE...' }),
  // The wave: a red paw raised beside the face, then held out to the side.
  waveUp: edit(sitting, {
    4: '.......RRRRRE.O.',
    5: '........WWWW.RO.',
    6: '........RWWRR...',
  }),
  waveDown: edit(sitting, { 6: '........RWWRRRO.' }),
  yawn: edit(sitting, {
    3: '.......RRRRRR...',
    4: '.......RRERRE...',
    5: '........WWEW....',
  }),
  // Curled up asleep, tail wrapped round with its white tip showing.
  curl: [
    blank,
    blank,
    blank,
    blank,
    blank,
    blank,
    '...........R..R.',
    '....RRRR...RRRR.',
    '..RRRRRRRRRRRRRR',
    '.RRRRRRRRRRRRRRR',
    '.RRRRRRRRRRREERE',
    'WRRRRRRRRRRRRRR.',
    'WWWRRRRRRRRRRR..',
    '.WWWWWWOOOOOO...',
  ],
}
export type Sprite = keyof typeof sprites
