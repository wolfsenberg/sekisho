# Sekisho brand assets

This directory keeps the assets used by repository documentation, color values, and the sources required to regenerate the full brand kit.

## Kept in the repository

| Files | Purpose |
| --- | --- |
| `logos/lockup-black.svg`, `logos/lockup-white.svg` | The horizontal logo used in the README |
| `social/github-dark.svg`, `social/github-dark.png` | The selected GitHub social preview |
| `source/Geist.ttf`, `source/OFL.txt` | The wordmark font and its license for reproducible exports |
| `tokens.json` | Brand colors, typography, and spacing rules |

The canonical symbol is the Sekisho seal, [`docs/assets/logo.svg`](../docs/assets/logo.svg): a square vermilion hanko with a framed, seal-script S cut through it, drawn as one even-odd path. [`logo-dark.svg`](../docs/assets/logo-dark.svg) uses the brighter dark-ground vermilion, and [`logo-small.svg`](../docs/assets/logo-small.svg) is the frameless version for sizes below 24px. Keep the geometry unchanged. The sole maintained generator lives in `tools/brand/`.

## Usage

- Use the black lockup on light backgrounds and the white lockup on dark backgrounds.
- Preserve proportions and keep at least one quarter of the mark canvas clear in ordinary layouts.
- Use the framed seal at 24px or larger; below that, use the frameless seal. The favicon files already do.
- The seal stays vermilion (`#D9442A` on light grounds, `#E5573B` on dark). Monochrome seal exports exist for single-colour printing.
- The wordmark uses outlined Geist 620 with -0.03em tracking, matching the application typeface.

Sekisho assets follow the repository's [AGPL-3.0-only license](../LICENSE). The included Geist font retains its separate [SIL OFL 1.1 license](source/OFL.txt).

## Generate the full kit

From the repository root:

```sh
make brand
```

The complete set is written to **`dist/brand/`**, which is ignored by Git. The command also refreshes only the two README lockups, the selected social preview pair, and `tokens.json` in this directory.

Open `dist/brand/preview.html` to browse all exports. The export includes PNG variants, avatars, app/browser icons, banners, social cards, presentation covers, wallpapers, stickers, usage guides, and a portable generator. The distribution guide template lives in `tools/brand/guide.md`.

Keep complete export archives as release attachments or separate downloads. Do not commit `dist/` or copy its generated files back into this directory. No release attachment is published by the build command.

The generator's integration test uses synthetic assets:

```sh
pnpm --dir tools/brand test
```
