# Presentation outline
These are the first draft of the text for the slides.
Each ## represents a main slide heading.
The first level of - represents a slide.
The next levels of - represent text on that slide.

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
- How did it go?

## The 2025 build (the "grog brain" approach)
- 2025 philosophy: minimal framework, custom pieces
- Contrast with 2024. Full heavy framework vs minimal groug brain approach
- How did it go?

## Comparison: "The Good, The Bad, The Ugly" (by topic)

### Stack
- Speed ("speeeeed")
- Build size
- Scaling considerations
    - Question: How many concurrent users does Regncon haeve?
        - About 200 pariticipants at the festival
        - Mostly 0 sometimes 1
- Sqlite
    - Raw SQL
    - Excellent local development experience
- Datastar
    - Performance
    - Simplicity
    - Was still under heavy development
        - That's ok, its 14k big 
- Signals
    - 
- Nats 
    - Smuuuuuth
- Server side components
- Firebase
    - NoSql lack of structure
    - Works well when used correctly
    - Terrible local development experience
- Jest worked super smooth
- Routing
    - Next.js file based routing 
    - Chi router in Go
- Boilerplate
    - Next.js 
    - Go
- Have to use web components for JavaScript
- Caching
    - Next.js built in caching
    - Go doesn't need it

### Hosting / deployment
- Vercel owns your soul
- Firebase hosting
- fly.io war very immature
    - Sqlite replication issues
    - Wrong documentation
    - Wrong tool for the job
        - Immutable docker containers was the wrong choice for our stack

### Language
- This is where Go shines
    - There is really no comparison

- Go is a real programming language
    - Real types
    - Structs 
    - Interfaces

- Go is opinionated
    - And that was really good
    - All go code looks the same no matter experience level

- Typesript is a superset of JavaScript
    - JavaScript is a scripting language
    - Dynamic types
    - Prototypal inheritance

- Updates and upgrades
    - Go is nearly completely painless to upgrade
        - Stability is one of the main features of the language
            - Other languages have solved this problem too
        - Stable syntax
    - React has major paradigm shifts every few years
        - Huge breaking changes
        - Major rewrites
        - Completely different ways of structuring code

- Go is productive
    - "gets out of your way"
    - Can be boring to write

- Go has some quirks
    - No language is perfect
    - It was built for web
    - Export vs capital letter
    - Abriviations everywhere
        - Optimized for typing
    - WTFS For some F***ing reason, go time.Parse cant take YYYY-MM-DD like a normal programming language.     parseLayout := "2006-01-02"     birthDate, err := time.Parse(parseLayout, born)

- Next.je has some quirks
    - Image component

- Build size
    - Go produces a single binary
    - Next.js needs 400mb of node modules to print "hello world"

- JSX/TSX is super smooth
    - Writing HTML in Go is a pain.
        - We tried to replicate TSX in Go with templ but it was not great

- Prop drilling 
    - React prop drilling is bad 
    - Go prop drilling is is good


### Structure & complexity
- "No magic" vs framework abstraction
- Low accidental complexity
- Cohesion, coupling, separation of concerns
- Where discipline is required (e.g., JS not being opinionated  easy to drift)
- Risks: obfuscation / complexity creep

### Styling
- UI library (MUI) vs custom CSS
- Next.js was challenging without a UI library
- Component styling approaches
- Issues seen with `<style>` tags in templates:
  - Duplicated CSS in the DOM when reused
  - Fragile placement/selector issues (e.g., IDs needing to exist at the top level)
  - Ending up with a mixed approach (component `<style>` + global CSS files)

### Tooling / DX
- The tooling for next.js is superb
- The tooling for templ was immature
- Logging visibility (difficulty spotting log output)
- Hot reload setup feeling immature / fiddly

### LLM assistance
- Where LLMs helped a lot
- Where LLMs struggled (edge cases, project-specific conventions)

## Closing
- Summary / takeaways
- Questions & discussion (including "what if we had 1 million concurrent users?"-style prompts)
