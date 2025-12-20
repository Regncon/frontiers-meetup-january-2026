# Presentation outline

## Opening
- Welcome
- Why we're giving this talk (the motivation / "why")

## What this talk is
- Agenda
- Who we are (Regncon devs / team intro)
- Different people, different perspectives (Dev 1 vs Dev 2)
    - Cant see the value of heavy frameworks, prove me wrong
    - Wants to tattoo React and Vs Code on their booty

##  Intro: **Next.js + Firebase**
- What is Next.js?
- What is Firebase?

## Intro: **Go + Datastar + SQLite (+ NATS / supporting pieces)**
- What is Datastar?
    - 1 million checkboxes demo
    - Question: 
        - Dont spoil the answer if you know it already
        - How many checkboxes can the regular dom handle with datastar?
            - 1 billion
    - 1 billion checkboxes demo
- Datastar demo
- What is NATS?
- What is SQLite?

## The 2024 build (the "heavy" approach)
- 2024 philosophy: heavy framework, batteries included
- Server side components vs client side components

## The 2025 build (the "grog brain" approach)
- 2025 philosophy: minimal framework, custom pieces
- Contrast with 2024. Full heavy framework vs minimal groug brain approach

## Comparison: "The Good, The Bad, The Ugly" (by topic)

### Stack
- Speed ("speeeeed")
- Build size
- Data layer notes (e.g., raw SQL)
- Hosting/deployment notes (and what felt mature vs not)
- Scaling considerations
    - Question: How many concurrent users does Regncon haeve?
        - About 200 pariticipants at the festival
        - Mostly 0 sometimes 1
- Sqlite
    - Raw SQL
    - Local development experience
- Signals
    - 
- Nats 
    - Smuuuuuth
- Server side components
- Firebase
    - NoSql lack of structure
- Jest worked super smooth
- Routing
    - Next.js file based routing vs custom grog brain routing
    - Chi router in Go


### Language
- Trust / confidence in the code
- Interfaces & structs (types that feel "real")
- Stable syntax
- Productivity / "gets out of your way"
- Pain points (example: Go date parsing / time formatting ergonomics)

### Structure & complexity
- "No magic" vs framework abstraction
- Low accidental complexity
- Cohesion, coupling, separation of concerns
- Where discipline is required (e.g., JS not being opinionated  easy to drift)
- Risks: obfuscation / complexity creep

### Styling
- UI library (MUI) vs custom CSS
- Component styling approaches
- Issues seen with `<style>` tags in templates:
  - Duplicated CSS in the DOM when reused
  - Fragile placement/selector issues (e.g., IDs needing to exist at the top level)
  - Ending up with a mixed approach (component `<style>` + global CSS files)

### Tooling / DX
- Logging visibility (difficulty spotting log output)
- Falling back to `console.log`
- Editor ergonomics:
  - Harder navigation (e.g., can't ctrl-click into templ components)
  - Weaker syntax highlighting (types not visually distinct)
- Hot reload setup feeling immature / fiddly

### LLM assistance
- Where LLMs helped a lot
- Where LLMs struggled (edge cases, project-specific conventions)

## Closing
- Summary / takeaways
- Questions & discussion (including "what if we had 1 million concurrent users?"-style prompts)
