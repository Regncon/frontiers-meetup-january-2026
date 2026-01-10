# Frontiers Meetup January 2026 — Agent Project Seed

This document is a concise, living summary of the presentation and its implementation. It is intended to seed future LLM discussions and onboarding.

---

## Purpose and tone

- Story-driven, honest comparison of two real Regncon builds.
- 2024: “Heavy frontend” (Next.js + Firebase / Firestore).
- 2025: “Grug brain simple” (Go + Datastar + templ + SQLite + NATS).
- On-slide tone is humble and factual; spicier takes live in presenter notes.
- Slides aim for one strong idea each, minimal text, strong readability.

---

## Narrative outline (high level)

1) Opening: Lobby welcome, why this talk.  
2) Who we are + Regncon context + different viewpoints.  
3) 2024 stack: what it is, snapshot, how it went.  
4) 2025 stack: philosophy, tech explainers, poll + demo beats.  
5) Good/Bad/Ugly by topic (stack, complexity, language, styling, tooling/DX, LLMs).  
6) What’s next (2026 → 2027).

The detailed outline lives in `presentationOutline.md` and the expanded planning notes live in `project_summary.md`.

---

## Slide ordering and source of truth

- Slide ordering lives in `pages/root/active_slide.templ` (`deck()`).
- Most slide content is in `slides/*.templ` with intent/design notes in top-of-file comments.
- Some older placeholder slides exist but are not in the deck:
  - `slides/slide_who_are_we.templ`
  - `slides/slide_what.templ`
  - `slides/slide_wellcome.templ`
  - `slides/slide_intro_next_js.templ`

---

## Interactive features

### Polling

- Poll definitions live in `pages/root/poll.templ`.
- Poll votes stored in table `poll_votes(invite_key, poll_key, session_id, option_key, voted_at_unix)`.
- Poll results are split by invite key (local vs remote) and summed in `pages/root/poll_results.templ`.
- Poll UI is Kahoot-inspired: 4 options, color + badge per option.

### Live updates

- Datastar SSE drives live patching for the root page.
- Updates are broadcast by re-saving state via JetStream KV (see `helpers/nats_helpers.go`).

### Presenter mode

- Presenter controls live at `/presenter` (keyed by `PRESENTER_KEY` env var).
- Presenter state is stored in a cookie session.

### Emoji reactions

- Emoji counter uses `emoji_counter` table and broadcasts updates through KV + SSE.

---

## Runtime architecture

- Single Go server (`main.go`) using Chi router + templ templates.
- Embedded NATS with JetStream KV bucket `presentation` (TTL 1h) for state + broadcasts.
- SQLite DB `presentation.db` with WAL + pragmas set at startup.
- Sessions:
  - Invite session stores audience (`local` vs `remote`) by URL key.
  - Separate presenter session for controls.
  - Per-session state ID stored in cookies for KV/SSE.
- Access model:
  - `/` shows “need a key” page.
  - `/{inviteKey}/...` gives the deck; invite key must match env vars.

---

## Data and schema

- `schema.sql` defines `emoji_counter` and `slide_state`.
- `poll_votes` schema is in a comment block in `pages/root/poll.templ` (not in `schema.sql` yet).
- `slide_state` has an additional `updated_at` column in `schema.sql` that is not used in code.

---

## Assets and paths

Slide comments reference assets in `/static/slides/...` and `/static/memes/...`, but actual files currently live directly under `/static/` (e.g., `/static/one_billion.webp`, `/static/explosion.webp`, `/static/ko-go_meme.webp`).

Potential missing assets referenced by slides:
- `/static/brand/frontiers-logo.svg`
- `/static/brand/regncon-logo.svg`
- `/static/tech/nextjs.svg`
- `/static/tech/go.svg`

If these assets are added, confirm paths in the slide templates.

---

## Build and run

- `task build:templ` runs `go tool templ generate`.
- `task build` builds the production binary.
- `task start` runs templ watch + Go hot reload (air).

---

## Open items / gaps

- Poll table creation is not in `schema.sql` (only in a comment block).
- Some slide intent comments mention asset paths that do not exist yet.
- The demo host slide for the “1 billion checkboxes” demo is a placeholder.
- The SQLite schema and DB bootstrap process are manual; no auto-migrations.

