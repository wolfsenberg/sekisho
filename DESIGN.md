---
name: Sekisho
description: A sumi-ink ledger for an internal AI gateway, where approval is a vermilion seal.
colors:
  primary: "#C2361C"
  primary-foreground: "#FFFFFF"
  seal: "#C2361C"
  ring: "#C2361C"
  background: "#F6F4EF"
  surface: "#FCFBF8"
  popover: "#FCFBF8"
  sidebar: "#EFECE5"
  sidebar-accent: "#E4DFD4"
  control: "#EBE7DF"
  control-hover: "#E2DDD2"
  secondary: "#ECE8E0"
  accent: "#EBE6DC"
  muted: "#EEEAE2"
  foreground: "#1A1916"
  muted-foreground: "#5F5A50"
  border: "#E2DDD3"
  hairline: "#D6D0C4"
  input: "#8C867B"
  destructive: "#A8283A"
  success: "#2A7343"
  success-muted: "#E8F3EB"
  warning: "#8A5A0B"
  warning-muted: "#FBF1DE"
  error: "#B4233A"
  error-muted: "#FBE8EB"
  info: "#2A5DA8"
  info-muted: "#E6EEF9"
  scrollbar: "#CFC9BD"
  dark-primary: "#C9432A"
  dark-primary-foreground: "#FBF8F2"
  dark-seal: "#E5573B"
  dark-ring: "#E5573B"
  dark-background: "#121110"
  dark-surface: "#181715"
  dark-popover: "#1C1B19"
  dark-sidebar: "#0D0C0B"
  dark-sidebar-accent: "#1F1D1B"
  dark-sidebar-border: "#24221F"
  dark-control: "#22201D"
  dark-control-hover: "#2A2825"
  dark-secondary: "#24221F"
  dark-accent: "#24221F"
  dark-muted: "#1F1E1B"
  dark-foreground: "#ECE8DF"
  dark-muted-foreground: "#A39E93"
  dark-border: "#2B2925"
  dark-hairline: "#34322D"
  dark-input: "#67625A"
  dark-destructive: "#A8283A"
  dark-success: "#8CC69B"
  dark-success-muted: "#15241A"
  dark-warning: "#E7B95E"
  dark-warning-muted: "#2A2214"
  dark-error: "#F1899A"
  dark-error-muted: "#2E171C"
  dark-info: "#94B6EA"
  dark-info-muted: "#16202F"
  dark-scrollbar: "#2F2D29"
  brand-seal: "#D9442A"
typography:
  title:
    fontFamily: "Geist Variable, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "26px"
    fontWeight: 600
    lineHeight: 1.3846
    letterSpacing: "-0.03em"
  auth-title:
    fontFamily: "Geist Variable, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "24px"
    fontWeight: 600
    lineHeight: 1.3333
    letterSpacing: "-0.025em"
  body:
    fontFamily: "Geist Variable, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "14px"
    fontWeight: 400
    lineHeight: 1.4286
    fontFeature: "'ss01', 'cv11'"
  label:
    fontFamily: "Geist Variable, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "12px"
    fontWeight: 500
    lineHeight: 1.6667
  wordmark:
    fontFamily: "Geist Variable, -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif"
    fontSize: "15px"
    fontWeight: 600
    letterSpacing: "-0.02em"
  mono:
    fontFamily: "IBM Plex Mono, ui-monospace, SFMono-Regular, Menlo, Consolas, monospace"
    fontSize: "14px"
    fontWeight: 400
  mono-label:
    fontFamily: "IBM Plex Mono, ui-monospace, SFMono-Regular, Menlo, Consolas, monospace"
    fontSize: "12px"
    fontWeight: 400
    lineHeight: 1.6667
  mono-chip:
    fontFamily: "IBM Plex Mono, ui-monospace, SFMono-Regular, Menlo, Consolas, monospace"
    fontSize: "11px"
    fontWeight: 500
    lineHeight: 1.8182
rounded:
  sm: "4px"
  md: "6px"
  lg: "8px"
  xl: "12px"
  2xl: "16px"
  full: "9999px"
spacing:
  xs: "4px"
  sm: "8px"
  md: "16px"
  lg: "24px"
  xl: "32px"
  page-x: "32px"
  page-x-narrow: "20px"
  page-y: "28px"
components:
  button-primary:
    backgroundColor: "{colors.primary}"
    textColor: "{colors.primary-foreground}"
    rounded: "{rounded.lg}"
    padding: "8px 16px"
    height: "36px"
  button-outline:
    backgroundColor: "{colors.control}"
    textColor: "{colors.foreground}"
    rounded: "{rounded.lg}"
    padding: "8px 16px"
    height: "36px"
  button-outline-hover:
    backgroundColor: "{colors.control-hover}"
  button-secondary:
    backgroundColor: "{colors.control}"
    textColor: "{colors.foreground}"
    rounded: "{rounded.lg}"
    height: "36px"
  input:
    backgroundColor: "{colors.control}"
    textColor: "{colors.foreground}"
    rounded: "{rounded.lg}"
    padding: "4px 12px"
    height: "36px"
  input-hover:
    backgroundColor: "{colors.control-hover}"
  nav-item-active:
    backgroundColor: "{colors.sidebar-accent}"
    textColor: "{colors.foreground}"
    rounded: "{rounded.lg}"
    height: "36px"
  status-success:
    backgroundColor: "{colors.control}"
    textColor: "{colors.foreground}"
    typography: "{typography.mono-chip}"
    rounded: "{rounded.full}"
    padding: "2px 10px 2px 8px"
  status-warning:
    backgroundColor: "{colors.warning-muted}"
    textColor: "{colors.warning}"
    typography: "{typography.mono-chip}"
    rounded: "{rounded.full}"
    padding: "2px 10px 2px 8px"
  status-error:
    backgroundColor: "{colors.error-muted}"
    textColor: "{colors.error}"
    typography: "{typography.mono-chip}"
    rounded: "{rounded.full}"
    padding: "2px 10px 2px 8px"
  tooltip:
    backgroundColor: "{colors.foreground}"
    textColor: "{colors.background}"
    rounded: "{rounded.lg}"
    padding: "6px 10px"
  menu:
    backgroundColor: "{colors.popover}"
    rounded: "{rounded.xl}"
    padding: "4px"
  panel:
    backgroundColor: "{colors.surface}"
    rounded: "{rounded.xl}"
    padding: "24px"
  dialog:
    backgroundColor: "{colors.popover}"
    rounded: "{rounded.2xl}"
    padding: "24px"
---

# Design System: Sekisho

## Overview

**Creative North Star: "The Checkpoint Ledger"**

Sekisho is the checkpoint that inspects every pass. The interface is a sumi-ink ledger: warm charcoal grounds (or washi paper in the light theme), kinari off-white ink, ruled with 1px hairlines. One shu-vermilion seal ink marks approval acts and nothing else. The operator should always be able to see where they are, what passed and what was stopped; approving something feels like stamping a pass.

It is an operational interface first: dense, legible, built from adapted Shadcn Admin primitives for long sessions in dim rooms beside terminals and editors. The world is carried by material (ink, paper, hairline, seal) rather than decoration. Its finish is native-app calm at the level of macOS and iOS: controls are filled tones rather than outlined boxes, cards are soft 12px sheets, menus grow out of the control that opened them, and the toolbar is translucent. The ledger's data and small labels are set in IBM Plex Mono; its prose stays in Geist. The dark theme is the default (`defaultTheme="dark"`); the light theme is a full washi-paper variant with its own tuned values, not an inversion. The implementation lives in `web/src/styles.css` and `web/src/components/`; update the tokens and this document together when an intentional design change is made. There are no marketing illustrations or decorative charts.

**Key Characteristics:**
- Warm sumi charcoal grounds in three steps: darkest sidebar, canvas, raised paper-dark panels.
- One vermilion seal ink reserved for approval acts; everything else is ink and paper.
- Hairlines (border, hairline, input) structure in-flow surfaces; only floating layers cast a shadow.
- Filled controls one tone off their surface (`control`), never outlined boxes.
- Soft native corners: 8px controls, 12px cards and menus, 16px dialogs; capsule status chips; 1.5px seal marks.
- Geist for interface and prose; IBM Plex Mono for the ledger voice (figures, IDs, keys, timestamps, chips, column heads, labels); both self-hosted.
- One authored motion, the hanko stamp; everything else is short state feedback on native curves.

## Colors

A warm, low-chroma sumi/kinari neutral world with a single shu-vermilion ink and four retuned semantic hues (moss, yamabuki, carmine-pink, ai-blue). Light values are the plain keys above; dark values carry the `dark-` prefix.

### Primary
- **Shu Seal Vermilion** (`primary` / `seal` / `ring`): primary action buttons, the active-route seal marker, the focus ring and caret, text selection (32% mix), checkbox and radio accent, the on state of toggle switches, completed setup-step squares, the brand seal in the app, the usage chart line and the activity heatmap ink. In dark mode the button fill (`dark-primary`) is slightly deeper than the seal ink (`dark-seal`) so the brighter seal reads as ink on charcoal while white button text keeps contrast.
- **Brand Seal Vermilion** (`brand-seal`): the exported brand assets only (README lockups, favicon set, `docs/assets/logo.svg`). Dark-ground exports use the same value as `dark-seal`.

### Neutral
- **Washi Paper / Sumi Night** (`background`, `dark-background`): the canvas.
- **Raised Paper** (`surface`, `popover`; dark `dark-surface`, `dark-popover`): cards, panels, dialogs, menus; popovers sit one step above cards in dark mode.
- **Sidebar Sumi** (`sidebar`, `dark-sidebar`): the navigation rail, the darkest ground in dark mode; `sidebar-accent` is the hover and active-row surface.
- **Kinari Ink / Sumi Ink** (`foreground`, `dark-foreground`) and **Faded Ink** (`muted-foreground`): body text and secondary metadata.
- **Control Tone** (`control`, `control-hover`; dark `dark-control`, `dark-control-hover`): the fill of every pop-up button, outline and secondary button, input, textarea and the success chip. It sits one tone off its surface so a control reads as a filled object, not a box; hover steps one tone further.
- **Secondary, Accent, Muted** grounds: menu-item focus (`accent`), neutral chips and empty heatmap cells (`muted`).
- **Three hairlines**: `border` (structural dividers, in-flow card edges, the header and demo-banner rules at 70%), `hairline` (a slightly stronger rule for the tegata panel), `input` (the ruled fields on the tegata pass, the neutral outcome-strip segment).

### Semantic
- **Moss** (`success`), **Yamabuki** (`warning`), **Carmine-Pink** (`error`), **Ai-Blue** (`info`), each with a `-muted` ground. `destructive` is reserved for destructive controls; error text and badges use `error`, which has a lighter dark-theme foreground. Paused or unconfigured states remain neutral.

Quota bars and remaining percentages use success at 20% or more, warning below 20%, and error below 5%. Unknown quota stays neutral and has no bar. Numeric labels remain visible.

Keep text contrast at least 4.5:1 and focus indicators at least 3:1 against adjacent colors. Filled controls are identified by their tone, label and position rather than a 3:1 edge; their focus state supplies the 3:1 seal edge. Subtle structural dividers are not interactive control boundaries.

### Named Rules
**The One Seal Rule.** Vermilion is the ink of approval: primary actions, the active route, focus, the brand seal, chart ink and activity density. It never decorates, never fills a panel, and never stands in for a status colour.

**The Dimmed Seal Rule.** A disabled primary action keeps its identity but loses its strength: vermilion at 35% with its text at 60%, and no ink-press shadow. Other disabled buttons fade to 50% opacity.

**The Icon And Word Rule.** Status colour never travels alone. Every status chip pairs an icon with text; every warning or error paragraph (`p.text-warning`, `p.text-error`) hangs a 16px triangle or circle icon in its inline-start margin through CSS, so wrapped lines keep one clean edge.

## Typography

**Body Font:** Geist Variable (with -apple-system, BlinkMacSystemFont, Segoe UI, sans-serif)
**Mono Font:** IBM Plex Mono 400/500/600 (with ui-monospace, SFMono-Regular, Menlo, Consolas, monospace)

**Character:** a clean, slightly technical grotesque that matches the wordmark, with stylistic sets `ss01` and `cv11` enabled, against IBM Plex Mono as the ledger's old-tech accent: the voice of entries, figures and labels, never of sentences.

Both faces are self-hosted through `@fontsource-variable/geist` and `@fontsource/ibm-plex-mono`, bundled with the app. No external font service is used.

### Hierarchy
- **Title** (600, 26px, 36px line height, -0.03em, balanced wrap): page titles.
- **Auth title** (600, 24px, -0.025em, balanced wrap): setup and sign-in headings on the pass.
- **Dialog title** (600, 18px, tight leading).
- **Section heading**: body size, medium weight.
- **Body** (400, 14px, 20px line height): descriptions and most controls. Page descriptions use a 24px line height and a 72ch measure.
- **Label** (500, 12px): secondary metadata in Geist.
- **Mono label** (IBM Plex Mono 400, 12px, 20px line height, muted ink): column heads (`th`), definition labels (`dt`) and field labels (`.field-label`), applied globally.
- **Mono chip** (IBM Plex Mono 500, 11px, 20px line height): status chips and sidebar group labels (group labels at 55% ink with +0.01em tracking).
- **Mono** (IBM Plex Mono): API keys, request IDs, device codes, regex and configuration text, timestamps (every `time` element), and metric figures that opt in with `font-mono`.

### Named Rules
**The Ledger Voice Rule.** IBM Plex Mono is the voice of the ledger's entries and labels: figures that opt in, IDs, keys, timestamps, status chips, column heads, definition labels, field labels and sidebar group labels. Prose (titles, descriptions, sentences, buttons, menu items) is never mono. Tables, `time`, `output`, `kbd` and `code` set `tabular-nums` globally.

**The Sentence Case Rule.** Every label, including the mono ones, is sentence case. No uppercase labels, and no negative tracking on mono; negative tracking belongs only to Geist titles and the wordmark.

Keep all visible UI copy and accessibility labels in the translation dictionaries. Supported languages are English, Japanese, Korean, Spanish, French, German, Portuguese, Russian and Arabic. Allow translations to wrap naturally, including longer European-language phrases; do not truncate important status or recovery information. In right-to-left layouts, code and preformatted text stay left-to-right and headings drop negative tracking.

## Layout

The desktop shell follows the Shadcn Admin layout: a 272px sidebar (darkest sumi) with the seal and wordmark top-left above the workspace selector, a sticky, translucent toolbar header (52px minimum height, canvas at 80% under a strong backdrop blur and saturation boost, a 70%-strength border rule beneath, 16px side padding, 24px from 768px), and a fluid main area capped at 1152px. When the instance runs in demo mode, a one-line demo banner sits directly under the header: 13px text on card at 60% with a flask icon, a medium-weight title and muted explanation, and a 70%-strength bottom rule. Main content uses 32px horizontal padding on larger screens, 20px on small screens, and 28px vertical padding. The sidebar ends with a compact account trigger containing initials, username, and actual account role, and its scrolling navigation fades out over the last 24px. Always show real instance data.

Request history, audit records, members, API keys and model prices use the full available main width. Reuse `TableViewport` for a viewport-bounded scrolling area with a sticky header and `Pagination` outside it, so horizontal scrolling remains accessible with long result pages. Show the filtered total, visible record range, current and total pages, 10/25/50/100-row choices, bounded page numbers and direct page entry. Use the existing custom dropdown pattern for page size. On narrow screens, retain the total, page size, previous/next controls and direct page entry while hiding the numbered strip; model prices retain their existing stacked rows. Filters and page-size changes return to page one; deleted or expired records clamp to the last valid page. Page changes reset vertical table scrolling while preserving the horizontal position. On short screens or with expanded filters, allow the main area to scroll rather than clipping controls or collapsing the table.

Below 768px, navigation becomes a 288px modal sheet opened by the header trigger, and the header shows the seal beside the trigger. Main content takes the full width. Lists of details and toolbars wrap or stack; the service details become a single column below 640px.

Personnel team member and allowed-pool checklists and allocation scheme fields use a single column on narrow screens, expanding into two columns at 640px; settlement fields expand into three. Member allowance rows keep the amount and unit together while long usernames wrap. Keep these forms inline and bounded, with divided rows and wrapping actions rather than nested cards.

Use spacing to group related information. The home page leads with personal usage and keeps account, client, or gateway access information concise. Administrators see a compact gateway prompt only while setup is incomplete. Full gateway setup guidance and instance health live on the administrator **Instance status** page. Show readable uptime and the last successful check time there. A failed refresh keeps the previous readings visible, adds an explicit warning, and marks current service health as unknown. Do not add dashboard metrics without real data. Language, theme, and a settings shortcut appear as icons at the top right on authenticated pages, including mobile. The sidebar groups common pages above administrator tools, including a separate **System settings** page for instance-wide configuration. The header settings shortcut opens personal **Preferences**. Give icons accessible names and hover descriptions, and reflect the selected theme with a sun, moon, or monitor. Setup and login share the same language and theme icon controls in a quiet 64px header. The Preferences form uses triggers with visible text.

System settings uses a collapsible sidebar branch with **Workspace alerts**, **Time zone**, a single platform-owner **Client versions** entry at `/admin/settings/clients`, and **Backup and restore**. Use the existing submenu primitives, an expansion chevron, and a subtle indentation guide. Highlight the active child with the same seal marker as top-level routes; keep the parent as a keyboard-operable disclosure button with `aria-expanded` and `aria-controls`. Selecting a child closes mobile navigation. A direct child URL opens its branch. The original system-settings URL and saved Codex, Grok, and Claude settings URLs redirect to Client versions. That page stacks all three providers using `ClientVersion` with provider-specific copy, independent form state, and separate save/check actions. Keep its flat divided sections, neutral effective/latest-version summaries, labeled manual-version fields, automatic-check checkboxes, wrapping actions, and secondary check metadata. The summaries use two columns on wide screens and stack on narrow screens; preserve localized loading, validation, success, failure, and disabled-action feedback.

Coarse-pointer targets (buttons, selects, menu buttons, menu items) expand to at least 44px, and inputs use 16px text below 768px and on touch devices. Table rows answer a fine pointer with a quiet 2.5% ink tint over 120ms; touch never sees a stuck hover.

## Elevation & Depth

Surfaces are flat at rest and structured by tonal steps and hairlines: sidebar, canvas and raised paper-dark panels in dark mode; washi canvas, lighter card paper and a deeper sidebar in light mode. Shadows appear only where something genuinely sits above the page. Menus, submenus, dialogs and sheets share one float shadow (`--float-shadow`): a 1px ring at 6% in place of a drawn border, a long soft drop and a short contact shadow, tuned separately per theme. The tegata pass keeps its own warm drop shadow, and the primary button carries a faint ink-press (inset top highlight over a 1px drop). Modal layers dim the page with 55% black and a light 2px backdrop blur. The header is the one translucent layer: content scrolls under it blurred. Shadow values are carried in `.impeccable/design.json`.

### Named Rules
**The Ruled Page Rule.** Separate content with a 1px hairline before reaching for a shadow or a nested card. Cards that sit in the page flow carry a border, not a shadow.

**The Floating Layer Rule.** Only what floats above the page (menus, dialogs, sheets, the tegata pass) casts a shadow, and floating layers use the float shadow's 1px ring instead of a hairline border.

## Shapes

Soft, native corners built from an 8px base radius (`--radius`): buttons, inputs, textareas, nav rows, menu items and tooltips use 8px; cards, panels, menus and the tegata pass use 12px; dialogs use 16px; status chips and toggle switches are full capsules; small close and icon affordances keep 6px. Seal markers are 7–8px squares with 1.5–2px corners set 6 degrees off true; heatmap cells keep 2px and outcome bars 1px. The seal itself is a square hanko; square forms with slight rotation are the system's one recurring silhouette. Scrollbar thumbs are fully rounded. Keep icons from Lucide with consistent stroke treatment; do not use emoji as interface icons.

## Components

### Logo and brand

The mark is the Sekisho seal: a square vermilion hanko with a framed, seal-script S cut through it, drawn as one even-odd path in a 256px square viewBox. `docs/assets/logo.svg` is canonical (brand vermilion for light grounds); `docs/assets/logo-dark.svg` uses the same path with the brighter dark-ground vermilion; `docs/assets/logo-small.svg` is the frameless seal for sizes below 24px, where the frame channel would close up. In the app, `Logo` renders the same geometry in the current `seal` ink (`small` switches to the frameless path). Keep the geometry unchanged and the seal vermilion; monochrome exports exist only for single-colour printing.

The web app ships a frameless vermilion seal favicon SVG, a multi-size ICO fallback and an Apple touch icon in `web/public/`, all generated from the canonical geometry. The repository keeps only the README lockups (`brand/logos/lockup-black.svg` for light backgrounds, `lockup-white.svg` for dark), the selected dark social preview, `tokens.json`, and the Geist source font in `brand/`. Run `make brand` to generate the full kit in Git-ignored `dist/brand/`. Follow `brand/README.md` for clear space (one quarter of the mark canvas) and format selection.

The graphic wordmark uses Geist 620 with -0.03em tracking, exported as outlined paths. In the interface, the wordmark is live Geist at 15px, 600, -0.02em beside a 28px seal. Hovering the sidebar brand link tilts the seal -6 degrees over 150ms.

### Hanko stamp and motion

The hanko stamp is the product's one authored motion: a square seal presses down from a slight lift and settles 6 degrees off true, like ink on a pass. Both variants use the ease-out curve `cubic-bezier(0.23, 1, 0.32, 1)`.
- **Active-route stamp** (220ms): the 7px nav seal marker stamps in from 1.9x scale and -14 degrees.
- **Sign-in stamp** (420ms after a 120ms delay): the 48px seal on the tegata pass drops 6px from 1.18x scale with a brief 2px blur, slightly overshoots to 0.98x, and settles.

Other motion is state feedback only, on two curves: the ease-out above and the iOS drawer curve `cubic-bezier(0.32, 0.72, 0, 1)` (`--ease-drawer`).
- **Buttons** press to 0.97 scale over 150ms ease-out; inputs ease border, ring and fill over 150ms.
- **Menus** grow from their trigger (Radix transform origin): fade plus scale from 0.96 over 150ms on open, fade plus scale to 0.97 over 100ms on close.
- **Tooltips** appear in 125ms from 0.97 and leave in 100ms.
- **Dialogs** fade and scale from 0.96 over 200ms (close 150ms) over the blurred scrim.
- **Sheets** slide on the drawer curve, 300ms in and 200ms out.
- **Sidebar** width and group-label fades take 250ms on the drawer curve; toggle-switch thumbs travel over 200ms on the drawer curve.

A global `prefers-reduced-motion` rule collapses all animation and transition durations to 0.01ms. There are no looping decorative effects.

### Buttons and selectors

Use `Button` for primary, outline, secondary, ghost, destructive and link actions. Buttons are 36px tall (32px small, 40px large, 36px square icon) with 8px corners and 14px medium Geist text.
- **Primary:** vermilion seal ink with white text and the faint ink-press shadow; hover lightens to 90%. Disabled primaries dim to vermilion at 35% with 60% text and no shadow (The Dimmed Seal Rule).
- **Outline and secondary:** the same filled look, the `control` tone with no border, stepping to `control-hover` on hover; outline adds a barely visible inset top highlight. The variant names survive from Shadcn but neither draws an outline.
- **Ghost:** transparent, `control` fill on hover; used for header icon controls.
- **Destructive:** destructive fill, white text, hover at 90%.
- **Focus:** a 2px ring in the seal colour, offset 2px from the canvas.

Language and theme choices use the adapted Shadcn Admin `DropdownMenu` with radio semantics, a checkmark on the selected item, neutral `accent` focus backgrounds on 8px item corners, and theme-aware 12px popover surfaces lifted by the float shadow. Do not use browser-native select popups. Text triggers retain a 16px Lucide chevron and a clear gap after the selected label; header triggers show only the language or current-theme icon. Preserve keyboard navigation, Escape dismissal, and return focus to the trigger.

Desktop controls are compact. Preserve loading and disabled states, and never show an enabled action that cannot perform useful work.

### Inputs / Fields

Inputs are 36px tall and textareas grow with content; both have 8px corners, the filled `control` tone with a transparent border, and step to `control-hover` on hover. Focus keeps the fill, turns the border to the seal and adds a 3px seal halo at 25%. Invalid fields use the destructive border and ring; disabled fields fade to 50%. Field labels sit above in the mono label voice (`.field-label`). Inside the tegata pass, fields become ruled lines (see Authentication).

### Navigation

Workspace content rules live at `/admin/content` in the Administration navigation. The page places refresh/configure or cancel/save actions in the header, with workspace policy on the left and a bounded rule list/editor on the right at desktop widths. Mode edits use a native keyboard-operable three-option radio control; rule types keep the shared dropdown. On narrow screens, policy metadata uses two compact columns and status badges sit beside wrapping rule names. Scope and detailed limits live in a disclosure below the primary controls. The page reuses the system typography and flat divided rows. A focused inline editor changes the workspace mode and up to 50 rules; saved patterns are never loaded into the browser, blank edits retain them, and literal terms use masked inputs, and newly entered regex drafts use a local syntax-highlighted native textarea. The decorative token layer is hidden from assistive technology, scrolls with the textarea and never rewrites input or inserts highlighted HTML; backend Go RE2 remains authoritative. Each rule has a disclosure for transient sample testing. Rule controls are disabled while a test runs; success and failure both clear the sample, leaving only match feedback or a localized error. Observe mode explicitly explains that matching content still goes upstream. Request details display bounded policy versions and rule IDs, with an administrator-only recovery link to the rules page.

Navigation follows persisted roles. The upper **General** group contains the shared overview with personal usage, personal API keys, personal request history, and Model prices (`/pricing`), available to both administrators and members; the lower **Administration** group contains subscription accounts, API channels, resource groups, all request records, team usage, instance status, and member management and is hidden entirely for members. The overview shows a personal workspace for members and a compact gateway prompt for administrators when needed. The former personal usage URL redirects to the overview. API key management is common to both roles and shows only the current user’s records. Revoked keys are hidden by default; a labeled Show revoked switch stays available in loading, empty, and failed views. Paused and expired keys remain visible. Filtering happens before server pagination, and changing the filter or revoking a key returns to the first page. The filtered empty state offers creation or viewing past keys; successful revocation has persistent inline feedback. Client configuration is opened on demand from a secondary Setup guide button beside Create key, using a right-side sheet with a fixed header and scrollable content. The closed guide takes no space below the key list and does not poll connection state; creation displays the full secret in a one-time dialog, and revocation requires confirmation. Both roles retain personal preferences in the header, including a password-change action. Password dialogs explain that all browser sessions will end while gateway keys remain valid. Member request limits and password reset use an overflow menu beside the existing row actions; dialogs restore focus to that trigger. Direct access to an administrator route shows a localized access-denied page before management content mounts. Member management uses a paginated table and a focused creation sheet; status changes show clear pending and error feedback.

Reuse `SidebarProvider`, `Sidebar`, `SidebarMenuButton`, and `SidebarTrigger`. Nav rows are 36px with 8px corners, 16px Lucide icons and ink at 72% opacity; group labels are IBM Plex Mono 11px at 55% ink, sentence case. The active route takes the sidebar-accent surface, medium weight, full ink, `aria-current`, and a 7px square vermilion seal at the row's inline end that stamps in over 220ms. Close the mobile sheet after navigation and retain keyboard focus handling from Radix. The account trigger opens a Shadcn Admin-style menu beside the desktop sidebar or above the trigger on mobile. Show the real username, account role, and a sign-out action; keep pending and failed sign-out feedback visible in the sidebar.

### Authentication: the tegata pass

Sign-in and setup present a travel pass (tegata): a centred raised-paper panel up to 416px wide with a hairline edge, 12px corners and a soft, warm-tinted drop shadow. The 48px vermilion seal sits over the panel's top inline-end corner and stamps once on arrival. Fields on the pass are ruled lines rather than boxes: a single bottom hairline, transparent ground, no inline padding; focus turns the rule to the seal with a 1px seal underline, and invalid fields turn it to the error colour. The page ground carries faint horizontal ledger ruling (hairline every 44px at 30% opacity) masked to an ellipse so it fades out above and below the pass. Setup progress shows two small square seals, filled vermilion for completed steps and a hairline outline for the pending one.

Optional Turnstile verification appears before the sign-in button and uses the active theme, flexible desktop width and compact mobile layout. Password and passkey controls wait for a valid challenge token; failed submissions refresh it, and widget loading failures offer an inline retry. The platform-owner Turnstile page under System settings uses the existing flat form layout, a verification checkbox, Site Key and masked Secret Key fields, persistent save feedback, and an empty secret field after saving. Explain that leaving the secret blank preserves it and that the login hostname must be configured in Cloudflare.

First-time setup starts with a centered welcome surface, the seal, a brief introduction, and a two-step outline of administrator creation and workspace entry. **Start setup** opens the administrator form on its own route; the form offers a return action until submission is in progress. Move focus to the page heading when moving between setup screens, and sign in automatically after successful creation. Initialized instances skip the welcome flow.

Setup and login share the workspace's compact icon controls, custom dropdown menus, typography, and palette. Keep the header transparent and borderless, with the 28px seal and wordmark at the start, and the pass focused on the current step. Forms retain visible field labels and localized inline errors. The private navigation is rendered only after authentication; the sidebar shows the current user and a sign-out action.

### Status

Use the shared `Status` component for success, warning, error, information, or neutral badges: filled capsules (full radius, no ring) with IBM Plex Mono 11px medium text and a 14px icon before the text. Warning, error and info use their muted ground and semantic ink with a triangle, circle-alert or info icon; neutral uses the muted ground with a pause icon. Success is a stamped pass: a neutral ink chip on the `control` tone with a 7px vermilion seal (rotated -6 degrees) in place of an icon. A temporary loading state uses a labeled spinner and respects reduced motion.

### Outcome strip

The usage metric area shows an outcome strip above the outcome sentence: a full-width row of 6px bars with 1px corners and 2px gaps, sized proportionally to completed (success, with a check icon above), errors plus rejected (error, with an x icon above), and incomplete plus canceled (input-grey, no icon). Empty parts are omitted. The strip is `aria-hidden`; the localized sentence beside it carries the exact counts.

### Data, errors, and empty states

System data comes from the backend. Show a loading state while connecting, an actionable error with retry on failure, and an explicit empty state for unavailable account integration. Never fill these states with invented account counts or usage charts.

The account connection form presents Codex, Claude, and Antigravity as three equal-width provider choices with their original full-color brand logos. Only Codex is selectable; the other two choices remain visible but disabled. Preserve a neutral selected border/background (ink border, muted fill), a checkmark, keyboard focus, and accessible pressed state. Existing Claude and Antigravity accounts remain visible with a temporarily disabled status and can be deleted. Provider brand marks are an intentional exception to the single-ink rule; do not recolor them with interface status tokens or the seal.

Network proxy selection in account connection and rebinding uses a searchable, keyboard-operable dropdown. Include the default network route as a choice; keep long proxy names readable and the menu bounded on narrow screens.

Subscription accounts use a compact list: identity and plan, quota windows, account status with token expiry, and icon actions align in columns when the list has enough width. Narrow containers stack the same content in reading order. Keep each quota window within its column; a single window must not occupy a separate full-width panel. Snapshot time remains secondary, and icon actions retain accessible names and hover descriptions.

Account pools use a compact list with names, enablement, account/member counts, and a clear edit action. No pool has a protected default role. Newly imported accounts show Unassigned with a link to account-pool setup until explicitly assigned. Pool editors use a bounded account checklist. Member grants live in the Members row action; personal key creation shows only authorized pool choices through a custom dropdown. Explain revocation and shared pool membership before applying changes.

Pool access is granted directly to members; ordinary sharing needs no allowance. The member list includes the workspace owner, administrators, and ordinary members. The owner remains protected from status and role changes. Joining an existing login uses its username and an explicit workspace role. Changing a role must preserve the membership's enabled state; creating a new member or joining an existing one offers an immediate pool-grant action. Personal key creation requires an explicit pool choice, shows each pool's account count, and warns when a selected pool is empty. The overview marks key setup ready only when an active key can reach a ready account through its current pool and allowance binding.

Optional resource allowances use dedicated pools and present percentage shares, direct internal USD amounts, direct token limits, and configurable time-window USD limits as mutually exclusive choices. Share mode chooses a Token/internal USD unit and either a calendar total or up to eight elapsed-window totals before assigning common member percentages. Time-window mode accepts up to eight whole-hour or whole-day conditions with shared USD defaults for selected members and per-condition personal overrides; each member is metered independently, and an empty limit is unlimited. Daily/monthly rules choose a reset clock or date and clock in the instance time zone; duration windows begin at activation and renew independently by elapsed time. Pool and reset-period selectors reuse the custom `DropdownMenu` radio pattern with a filled trigger, selected checkmark, and chevron. Keep fixed pool choices visible but disabled when editing. Group rule details, allowance fields, member selection, and activation into flat sections with dividers; keep the incumbent background. Put name and pool side by side on wide screens, arrange mode choices in two columns, align member inputs, and provide a select-all checkbox for each mode. Show M, USD, or % beside inputs and allocated/unallocated percentages below share inputs. For priced rules, snapshot prices from models supported by the selected pool without asking for prices in the allowance form. Keep internal budgets distinct from upstream capacity and cash balances.

Allocation balances use divided rows with tabular allowance, used, and remaining values. Show reset timestamps in the configured instance time zone independently of status. Time-window balances group every configured condition under one member, identify which condition blocks admission, show an unlimited condition without a zero remaining amount, and derive the member status from the most restrictive limited condition. Ratio balances use the selected amount or token unit after multiplying the total by the member's percentage. Use the shared `Status` component: unavailable is neutral, new-request risk pauses are warning, exhausted is error even when reconciliation is pending, and active is success. Current-window active/pending counts, provisional reservations, and admission headroom explain a risk pause without presenting reserved capacity as spent usage; pending entries outside a window do not change that window's status. Current scheme details remain distinct from a secondary next-cycle mode/effective-date notice. The current-member editor retains usage and captured reservations and exposes only member inputs. Ordinary budget/price changes stay scheduled without restarting intervals; structural changes require explicit new-period confirmation. Access enablement stays a separate control. Forms retain reviewed version conditions through background refresh and offer explicit refresh/review after conflicts.

The allocation edit form shows the server-calculated effective date before saving. For a new time-window rule scheduled to start next period, preview the date from its longest entered condition. Surface live price-coverage warnings near the current rule and in the edit form: current catalog prices win, previously saved prices cover gaps, and models missing from a revision's saved rates remain blocked until a new revision takes effect. Keep catalog-price gaps, saved-rate gaps, and temporarily unavailable account catalogs distinct; recompute these warnings on reads instead of storing them in rule config. Keep the distinct blank-value explanations beside direct member inputs and time-window condition inputs.

An uncovered-model warning includes the server-calculated edit effective date, so administrators can see when saving a price update will actually permit that model. For an already scheduled revision, use its saved effective date and past-tense scheduling copy.

Administrator allocation reports expose refresh and inline Token correction forms beside pending request IDs, with explicit M units and error feedback. Personal allocations reuse the balance presentation without editing or settlement controls, with visible loading, failure, empty, and refresh states.

Account scheduling metadata stays compact beside account status. Settings use a focused dialog; limits and cooldown controls do not replace authentication status. General → Requests shows the current user’s calls, with key/pool metadata and a result filter. Administration → All requests adds caller/subscription-account information and an account filter. Both tables use bounded pagination and explicit unknown token values; the personal page never fetches subscription-account management data. Keep source content and credentials out of logs and UI.

### Usage views

The personal usage section on the overview and the team usage page share the established request-table layout. A flat metric strip shows real counts, completion rate, mean duration and reported input/output tokens before personal allowances, followed by the outcome strip. Custom selectors choose the period, breakdown and request/input-token/output-token metric. Today and the 7/30/90-day periods use calendar-day boundaries in the configured instance time zone, including 23/25-hour days and skipped midnight. Daily trends use an interactive line chart drawn in seal ink (2px vermilion line, hollow canvas-filled points that fill vermilion when inspected, dashed border-colour gridlines and muted axis labels); model, pool and member comparisons use horizontal bars ordered by request count. Use responsive SVG geometry with readable axis labels, pointer/touch and keyboard inspection, and gaps for unknown token reports. Exact tables expand on demand through View details. Display the collection start in the configured instance time zone, hide unobserved earlier days, label partial token totals, and preserve the unknown state instead of inventing zeros. Personal usage adds a compact read-only limit summary; team usage also offers member breakdowns.

Personal usage adds an All keys selector beside the period. Load key metadata only when opened, include revoked keys with a status label, and keep every result page reachable with retry feedback. Selecting a key updates the metric strip, all chart breakdowns, detail tables, and hourly heatmap together; keep the selection when the period changes. Display the separate collection start for attributed key data, preserving earlier unattributed usage in All keys. Personal allowances and member limits retain their own user-level scope.

The full-width Activity by time card sits below the primary usage chart and above expanded detail tables in both usage views. Use a Monday–Sunday by 00–23 grid in the configured instance time zone, with borderless square cells (2px corners) drawn as seal-ink density: zero is the muted ground, then vermilion at 20%, 40%, 65% and full strength by quartile of the period maximum; uncollected or unreported hours use a fainter muted fill (40%). Show visible exact values on pointer/touch or keyboard selection. It follows the selected period and metric without changing with the primary breakdown. Keep labels readable and scroll only the grid horizontally on narrow screens; the detail footer remains outside the scrolling area. Show the independent hourly collection start and the grid label in the aggregation time zone returned by the server; never present pre-collection or future hours as zero. Refresh the activity when the instance time zone changes. Preserve historical hourly activity where local hours align with UTC hours; use minute aggregates where offsets or date boundaries split a stored hour, so earlier calls cannot leak into the selected period. Show the independent minute-collection start; legacy UTC daily totals remain preserved and are not fabricated into local dates. Use the returned report zone for all chart/table dates and CSV metadata, and use the next day boundary rather than assuming 24 hours when filtering uncollected days.

Personal key rows offer a compact Change group icon button beside the current group, retaining a localized tooltip and accessible name. Its focused dialog loads only authorized choices, includes allowance names, and explains that the key remains unchanged, existing usage stays in the original group, and switching requires a new conversation. A missing allowance binding has its own status and a Restore access action with confirmation; unavailable membership instead provides administrator guidance. Keep loading, retry, pending, failure and persistent success feedback explicit. New allowance forms explain automatic binding of selected members’ existing unrevoked keys and loss of access for excluded members before saving.

### Model catalogs

Model catalogs use compact **Models** actions on accounts, pools and personal keys. Account dialogs distinguish unknown support, a known empty catalog, stale saved data and refresh failure, and show the last successful observation time. Pool catalogs are read-only derivations of enabled accounts and existing policy. Keep searches and bounded lists inside the dialog. Include hidden capabilities in browser catalogs with a compact neutral “Hidden model” text badge beside the native ID; place them after visible models, keep alphabetical order within each group, and count and search them normally without treating the visibility label as an availability result. Preserve client discovery and picker filtering. CC Switch uses an accessible custom combobox over the selected key’s catalog; do not retrieve secrets to populate model choices. Show native model IDs without provider prefixes, with identical IDs appearing once across providers.

Model rows show upstream-reported capabilities in a compact definition list below the native ID. Numeric limits use localized tabular figures; unknown values remain explicitly unknown and reported false values read as unsupported. Keep existing search, hidden-model badges and model-test actions. Explain conservative resource-group merging beside the catalog, without suggesting that capability metadata grants access.

### Overlays

Menus and dialogs are raised paper (popover, one step above card in dark mode) lifted by the float shadow with no drawn border: menus have 12px corners and 4px inner padding; dialogs have 16px corners and 24px padding over a 55% black scrim with a 2px blur. Sheets slide from the edge on the popover ground with a hairline on their inner edge and the float shadow. Tooltips invert to ink on canvas with 8px corners, 12px text and a small rounded arrow.

### Backup and restore

Keep export and restore in two flat sections under System settings. State that archives contain credentials. Verification shows creation time, Sekisho version and database size; restoration prepares a new directory and shows the path with explicit stop/configure/restart steps. An existing target blocks overwriting. Disable dependent actions until verification succeeds, clear verification when the file changes or is checked again, and retain honest error/permission/loading states in every supported language.

### Request diagnostics

Keep account/result filters in the compact request toolbar. Additional model, request ID, caller and local-time filters expand inline and use the existing custom dropdowns. Show first-output latency as secondary text under total duration; null measurements remain an em dash. Show a compact, selectable request ID with a copy action below each start time; use an em dash when older records have no ID. The start-time button opens a focused request details dialog with a selectable/copyable correlation ID. On narrow screens, filters stack and the data table scrolls inside its own container. Quota exhaustion uses a warning status with a short explanation that existing conversations retain their account.

Place routing explanations in a flat divided section inside request details. Show the recorded strategy, final candidate count and nonzero aggregate exclusion counts with wrapping labels. Older or unrouted records explicitly state that no recorded explanation is available. Personal views never add account identities.

Quota usage pace uses neutral wrapping text below the existing reset information. A keyboard- and touch-accessible disclosure explains its observation time and reference semantics. Preserve remaining-quota bar colors, and omit pace when the snapshot or window cannot support the comparison.

### First-call walkthrough

The overview offers a compact, optional entry to a floating walkthrough. Highlight the existing control for the current step and position a readable, dismissible panel nearby, keeping it inside the viewport. Keep the underlying page usable. Hide the panel while a dialog or menu is open, preserve newly created keys until the user closes their dialog, and resume afterward. Back and skip stay available; advancing requires server-confirmed completion. Unknown progress never implies success. Keep keyboard focus, Escape dismissal, localization, and narrow-screen placement consistent with the rest of the application.

### API channels and resource groups

Subscription accounts (`/accounts`) and API channels (`/channels`) are separate administrator surfaces. Subscription creation offers OAuth or credential JSON for subscription providers only. Channel creation uses a dedicated `/channels/new` page with platform, upstream protocol, name, base URL, masked API key and network proxy. Searchable, keyboard-operable platform and protocol dropdowns share the proxy picker interaction. Platform search matches the displayed name, platform ID and preset endpoint, resets on reopening and preserves the current selection until confirmed. Selecting a platform fills the default endpoint and protocol; Custom keeps the endpoint editable and offers all supported upstream protocols. Long platform lists scroll within the menu. Platform choices and channel rows use locally bundled brand icons, with a neutral plug for Custom. Key replacement stays in a focused dialog with an empty secret field and read-only platform, protocol and endpoint. Channel lists show endpoint, protocol, actual credential state, observed activity and resource-group assignment, without subscription quota or token-expiry UI. Missing runtime observations stay unknown.

Resource groups retain `/groups` and existing key bindings. The editor presents subscription accounts and API channels in separate bounded lists, followed by routing preference and model policy. New groups prefer subscriptions and leave API fallback unchecked. Mixed groups explain the fallback condition; API-only groups explicitly use their configured channels. Source removal and routing restrictions preserve conversation identity by failing affected continuations instead of moving them. Navigation, member grants, key configuration, diagnostics and onboarding use the resource-group terminology. Initial activation offers both subscription and API-channel setup.

### Allowance scheduling controls

Keep pending cancellation and postponement inline with the existing allowance row. Initial-only rules disable cancellation and explain pause/reschedule alternatives. Preview the server-computed successor date before confirming a move; keep the reviewed version/time together. Delayed creation likewise shows the server start time before confirmation. Before the first activation, enable initial member correction in the full editor and explain that the reviewed time and anchor stay fixed. For an approved pending new structure, offer the existing new-period confirmation again to edit its members; ordinary budget/price edits keep it unchecked. Withdrawing optional reconfirmation shows the saved member instruction rather than a locked unsaved draft. Show shared-budget blocking without exposing another member's usage. A failed confirmation keeps the reviewed draft and requires explicit review rather than automatic rebasing.

## Do's and Don'ts

### Do:
- **Do** reuse the semantic tokens and adapted Shadcn Admin primitives; change `styles.css` and this file together.
- **Do** reserve vermilion for approval acts: primary actions, the active-route seal, focus, the brand seal, chart ink and activity density.
- **Do** dim disabled primaries to vermilion at 35% with 60% text and no shadow.
- **Do** fill controls with the `control` tone; let focus, not a resting outline, draw the seal edge.
- **Do** set figures that opt in, IDs, keys, timestamps, status chips, column heads, field labels and sidebar group labels in IBM Plex Mono, sentence case.
- **Do** pair every status colour with an icon and text; a success state is stamped with the seal.
- **Do** verify the dark default and the light variant, keyboard navigation, narrow layouts, right-to-left Arabic, and the supported languages (English, Japanese, Korean, Spanish, French, German, Portuguese, Russian, Arabic).
- **Do** keep the hanko stamp as the one authored motion and respect `prefers-reduced-motion`.
- **Do** distinguish a planned capability from a working action.

### Don't:
- **Don't** use vermilion as a decorative fill, a panel background, or a substitute for a status colour.
- **Don't** add decorative colour gradients, glows, vivid dashboard fills, or arbitrary accent colours; the only patterned ground is the faint ledger ruling behind the tegata pass.
- **Don't** load fonts from an external service; Geist and IBM Plex Mono ship with the app.
- **Don't** set prose in mono, uppercase a label, or apply negative tracking to mono text.
- **Don't** draw outlined boxes around buttons or fields, or give in-flow cards a shadow.
- **Don't** alter the seal geometry, recolour it outside the vermilion pair, or use the framed seal below 24px.
- **Don't** display fake business records or performance claims.
- **Don't** rely on color, hover, or icons alone for essential information.
- **Don't** expose SDK configuration details in routine team workflows unless users need them to make a decision.
