# Frontiers Meetup January 2026 — Project Summary (LLM Seed)

This document summarizes the presentation planning + implementation decisions so far. It is intended as a seed for future LLM prompting and for onboarding new contributors.

## High-level goal

A story-driven, honest comparison of two real builds for the Regncon festival website:

- **2024:** “Heavy frontend” — Next.js + Firebase (Firestore), with strong interest in learning Server Components.
- **2025:** “Grug brain” — Go + Datastar + templ + SQLite (+ NATS), deliberately exploring a simpler architecture and lower accidental complexity.

The talk uses **Good / Bad / Ugly** framing by topic, but keeps the tone humble and non-ideological on-slide. Spicier opinions live mostly in presenter notes.

---

## Presentation style principles

- Story-first pacing with intentional “beats” (setup → punchline → reveal → demo).
- Slides: minimal text, strong readability, whitespace; keep technical blame/analysis in presenter notes.
- Prefer **one strong idea per slide**.
- Use “impact slides” (image-only) for humor/emotional emphasis (e.g., KO meme, explosion slide).
- No persistent header/footer; the Agenda slide is shown only when needed.

---

## Core slide decisions completed

### Lobby welcome slide (Slide 1)
Purpose: displayed while people mingle (pizza/beer/networking). First impression must communicate:
- What talk is about (comparison of stacks, honest experience).
- “Starting soon” messaging.
- Tech logos should be small and tasteful (Frontiers + Regncon); tech logos optional.

### Why this talk (Start-with-Why slide)
- Humble, story-first, non-spicy.
- **One bullet only**:
  > “We want to share a real experience: what worked, what didn’t, what surprised us.”
- Presenter notes emphasize: not here to change minds; stories travel; some claims backed with demos/data.

### Agenda slide (progress indicator)
- Reusable slide component taking a `current AgendaStep`.
- Steps (updated):  
  1) Who we are + what we built  
  2) 2024 stack  
  3) 2025 stack  
  4) Good/Bad/Ugly (by topic)  
  5) **What’s next (2026 → 2027)** (replaces “Takeaways”)

### Regncon context slides
- What is Regncon (festival description).
- Who we are (dev group culture).
- Different viewpoints slide (humor: “prove me wrong” vs “React/VSC tattoos”).
- Project vision slide includes Norwegian statement:
  - Vision: *“Vi gjør vårt beste at alle deltagere skal få spille minst et spill som de er veldig interessert i å spille.”*
- Note: project is real production with real users; repeated “production-ready” failures are part of the story/foreshadowing.

---

## 2024 section: what we built + how it went

### 2024 snapshot
- Browse events by **pulje** (time slots).
- Participants mark interest levels.
- Submit interest (core workflow).
- Admins edit/publish + assign participants based on interest.
- Notes explain “pulje” and emphasize core flow.

### 2024 how did it go
- Underestimated work.
- Shipped not production-ready.
- Crunch (coding during festival).
- Usable in the end.
- Notes foreshadow later deep dives:
  - Overused Server Components.
  - Firebase is great when used correctly; mixing with Server Components contributed to performance issues (including interest selection).

---

## 2025 section: what we built + how it went

### 2025 snapshot (mirrors 2024 for fairness)
- Same core workflow as 2024.
- Admin ticketing support completed.
- UX/design better due to larger team + UX specialist.
- Allocation feature **not completed**:
  - On-slide uses strike-through: ~~Assign participants based on interest~~
- Notes explicitly: comparison is fair because workflow remained same; also admits “finished less overall”.

### 2025 how did it go (story pacing)
Three-slide story beat:

1) **“So… how did it go in 2025?”** (question-only, pause)
2) **Explosion image slide** (no text on slide; title added on top of image):
   - “Regncon servers less than 24 hours before start”
3) **Real “How did it go in 2025?”** summary:
   - Underestimated work.
   - Not production-ready.
   - Crunch (coding during festival).
   - Unsuccessful: not fully usable.
   - Survived with manual Excel backup process.

Key incident details (for later Hosting/Deployment topic):
- “Actual traffic” was **several** concurrent users (not dozens).
- **Litestream replication (VFS package) was immature**, replication failed, **data loss** occurred.
- Backup plan: Excel; veterans stayed calm; team was prepared.

---

## Interactive features implemented: Polling

A generic polling system exists:

- DB table `poll_votes(invite_key, poll_key, session_id, option_key, voted_at_unix)`
- Primary key ensures 1 vote per session per poll per invite key; upsert allows changing vote.
- Results page shows totals for **local key**, **remote key**, and **sum** (two invite keys).
- Kahoot-inspired UI: color-coded options and geometric badges.

### Poll used: “Datastar checkbox trick”
- Poll key: `dom-checkbox-limit`
- Setup slide explains “1 million checkboxes in React meme”, asks:
  - “How many can a simple Datastar + Go + SQLite app handle on a cheap VPS?”
- Poll options:
  - `2 (DOM is slow)`
  - `~10,000 (then it struggles)`
  - `~452,600 (measured on this laptop)` (reasonable-sounding anchor)
  - `1,000,000,000` (correct answer; includes Dr Evil image)
- Reveal slide: “One billion.” (Dr Evil image)
- Demo host slide exists as placeholder for the interactive checkbox demo component.

Chaos is desired in voting UX (no attempt to reduce “spoilers”).

### Poll used: “Regncon concurrency reality check”
- Setup slide claims Next.js+Firebase could likely scale to ~10,000 concurrent users (for this app).
- Poll asks: “How many concurrent users does Regncon have?”
- Options:
  1) Mostly 0, sometimes 1 (**correct**)
  2) 20–30 (that’s when you start ‘thinking about scaling’)
  3) 200 (there were 200 people at the festival)
  4) One billion
- Reveal slide is separate: “Answer: Mostly 0, sometimes 1.”

---

## Tech explainer slides added

- What is Next.js
- What is Firebase
- What is Go
- What is templ (planned as image/code screenshot slide)
- What is Datastar (BYOB backend, hypermedia framework, ~10 KiB script; exact size varies)
- What is SQLite
- What is NATS (JetStream optional persistence)

---

## “Good / Bad / Ugly (by topic)” slides added so far

### Styling
- **Good:** keep styling close to UI code (JSX/TSX locality).
- **Good:** MUI great for React.
- **Bad:** templ styling was a major pain point (details in notes).

### Tooling / DX
- **Good:** React tooling is amazing.
- **Bad:** templ tooling was terrible (LSP, go-to-definition, editor essentials were hard).
- **Good:** Go tooling is strong/reliable.
- Additionally: a standalone **KO impact slide** is planned for React tooling win.

### LLM assistance
- **Bad:** templ is newer/evolving → weaker LLM support.
- **Good:** React has strong LLM support.
- **Bad:** easy to generate huge amounts of React code (trap).
- Slide includes generated “obviously AI-generated dev chaos” image.
- Presenter notes include the line:
  - “We can’t have a presentation in 2026 without talking about AI.”

### Complexity
- Essential vs accidental complexity.
- Emphasize real costs: speed of development, cost, maintenance, “time to spaghetti”.
- Slide prefixes include Bad/Ugly outcomes (hard to feel early; shows up later).

### Stack summary slides (short, punchy)
**2025 stack slide**
- Good: Speeeeeeeeeed
- Good: Raw SQL
- Bad: templ was immature
- **Catastrophic:** Fly.io + Litestream replication was immature
- Notes: Datastar maturity mentioned as not significantly hampering dev; deep dive later.

**2024 stack slide**
- Good: Server Components
- Good: Firebase
- Bad: Firebase local development
- Notes keep deeper critique (overuse of RSC) for later.

---

## Language section plan (reduced slides, high impact)

Decisions:
- Use a standalone **Street Fighter KO** impact slide for Go (image-only).
- Use a standalone KO impact slide in Tooling/DX for React (balance).
- Reduce slide count:
  1) KO impact slide (Go)
  2) Go “wins” slide with **two-column bullet list** (intentionally intense)
  3) Go is not perfect slide (quirks + code examples)
  4) Error handling slide (web handler code example)
  5) JS/TS wins slide + **Good: Go prop drilling / explicit params** + Bad: React prop drilling

---

## “What’s next” final slide (replaces takeaways)

- **2026:** no full rewrite for October.
- **May 2026:** finish/harden the 2025 project to production-ready.
- **October 2026:** use hardened 2025 site for the festival.
- **2027:** decide next experiment:
  - Svelte + Supabase, or
  - retro 90s RPG makeover on current stack (gold plating + easter eggs).

---

## Security slide idea (planned)
- A high-impact slide about a critical Next.js CVE:
  - CVE-2025-55182 (CVSS 10.0)
- Goal: big-impact typography; optional background imagery.
- (Implementation requires verifying details and patch guidance when finalizing.)

---

## Assets referenced (paths used in slides)
- `/static/memes/dr-evil-one-billion.png` (used in poll option + reveal slide)
- `/static/slides/regncon-server-explosion.png` (image-only explosion slide)
- `/static/slides/ai-chaos-developer.png` (LLM slide image)
- `/static/slides/ko-go.png` (Go KO impact slide)
- `/static/slides/ko-tooling-react.png` (React Tooling KO impact slide)

---

## Key implementation notes (engineering)
- Favor inline styles for single-use slide styling.
- Promote to shared CSS classes only when reused across multiple slides.
- Keep slides in **one file per slide** (locality).
- Polling uses session-based identity from existing session state (`LoadOrCreateState`) and broadcasts updates via KV/NATS helper (`BroadcastUpdate`).
- Two invite keys exist (local vs remote); results show each + sum.

---

## Open items / TODOs
- Integrate 2025 snapshot + 2025 how-did-it-go slides into the deck (if not already).
- Implement the actual “1 billion checkboxes” demo component and wire it into the demo host slide.
- Add “What is Datastar” slide into deck near the Datastar poll sequence (if not already).
- Add the concurrency poll sequence into the deck at the right moment in Stack/Scaling discussion.
- Decide final placement and wording for the CVE impact + fix slides.
- Update “Agenda” calls in deck to use the final step `AgendaWhatsNext` and place the “What’s next” slide last.

---

## Slide/component inventory (names used so far)
(Useful for quick searching in the codebase)

- LobbyWelcome
- WhyThisTalk
- Agenda(currentStep, isPresenter)
- WhatIsRegncon
- ProjectVision
- WhoWeAre
- DifferentViewpoints
- Build2024Philosophy
- WhatIsNextJS
- WhatIsFirebase
- Build2024Snapshot
- Build2024HowDidItGo
- Build2025Philosophy
- GrugBrainMeme
- Build2025HowWeChoseTheStack
- WhatIsGo
- WhatIsTempl
- WhatIsDatastar
- DatastarCheckboxTrickSetup
- DatastarCheckboxTrickReveal
- DatastarCheckboxTrickDemo (host placeholder)
- PollSlide(db, inviteKey, sessionID, poll)
- PollResultsSlide(db, poll, localKey, remoteKey)
- RegnconConcurrencyPollSetup
- RegnconConcurrencyPollReveal
- Build2025Snapshot
- Build2025HowDidItGoQuestion
- Build2025ServerExplosion
- Build2025HowDidItGo
- StylingGoodBad
- ToolingAndDXGoodBad
- LLMAssistanceGoodBad
- ComplexityEssentialVsAccidental
- Stack2024GoodBad
- Stack2025GoodBad
- WhatsNext2026_2027
- LanguageGoKOSlide
- LanguageGoWins
- LanguageGoNotPerfect
- LanguageGoErrorHandling
- LanguageJSTSWins
- ToolingDXKOSlide

---
