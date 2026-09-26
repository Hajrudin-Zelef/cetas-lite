---
id: collect-240926-datacamp/datacamp/un-stack-tech-de-vibe-coding-pragmatique-pour-livrer-vite-2
title: "un-stack-tech-de-vibe-coding-pragmatique-pour-livrer-vite"
domain: datacamp
role: reference
task: reference
actors: ["Anthropic", "Stripe"]
dates: []
keywords: ["claude", "copilot", "mcp", "open source", "reasoning"]
source: docs/RAG/clean_en/datacamp/un-stack-tech-de-vibe-coding-pragmatique-pour-livrer-vite.md
source_anchor: ""
source_lines: [123, 277]
sha256: e1436e63feae1f25a0138aa133e3ca50a3dc5b6d3309d54c6ee036e84e68e14b
---

# un-stack-tech-de-vibe-coding-pragmatique-pour-livrer-vite

Supabase serves as the backend foundation by offering a managed Postgres database with dashboards and backup options, built-in authentication, file and blob storage, realtime, and scheduled tasks.

Resend handles transactional email, Stripe handles billing and subscriptions, and PostHog and Sentry provide analytics and error visibility through their free tiers.

Testing stays light with Vitest and Playwright, while CI/CD runs through GitHub Actions with database migration syncing.

- Frontend: Next.js + Tailwind + shadcn/ui on Vercel
- Backend: Next.js API + server actions
- Database: Supabase Postgres (includes dashboard + backup options)
- Auth: Supabase Auth (email/password, magic links, OAuth)
- File / blob storage: Supabase Storage (blob storage equivalent)
- Email: Resend (transactional)
- Realtime: Supabase Realtime (optional)
- Scheduled tasks / Cron: Supabase Scheduled Functions (or GitHub Actions cron)
- Analytics: PostHog (free tier)
- Error tracking: Sentry (free tier)
- Search: Algolia (small free tier) or Meilisearch Cloud (small plans)
- Payments: Stripe
- Testing: Vitest + Playwright + local Supabase test utilities
- Deployment: Vercel
- CI/CD: GitHub Actions + Supabase migrations

Ideal when: you want a central dashboard for most backend needs, generous free tiers, and a stack that scales without pulling you into infra early.

### 3. The Control Stack: 100% open source and no lock-in

This stack is for those who want total control, zero vendor dependency, and fully local operation. Every component is open source and runs via Docker Compose, which lets anyone clone the repo and start a complete SaaS environment with a single command.

Next.js runs the frontend and backend locally, while Postgres provides a self-hosted relational database.

Schema management is handled by Drizzle locally. Authentication is handled with Better Auth, to keep identity and access entirely under your control.

MinIO replaces cloud blob storage with a local S3-compatible service. Mailpit captures outgoing emails in development, and Redis powers caching and background jobs via BullMQ.

Meilisearch handles fast full-text search, and optional observability tools like Grafana, Prometheus, and Loki can be added for increased visibility.

Deployment remains flexible via Coolify or Dokku on a VPS, with GitHub Actions to build and deploy Docker images over SSH.

- Frontend: Next.js + Tailwind + shadcn/ui (runs locally)
- Backend / API: Next.js route handlers
- Database: Postgres (Docker)
- Migrations / ORM: Drizzle (local)
- Auth: Better Auth
- File / blob storage: MinIO (S3-compatible, Docker)
- Email: Mailpit (catches emails locally) + optional SMTP container
- Cache / queues: Redis (Docker)
- Background jobs: BullMQ (Node) + Redis
- Search: Meilisearch (Docker)
- Analytics: Plausible (Docker)
- Observability: Grafana + Prometheus + Loki (optional, but possible in Docker)
- Tests: Vitest + Playwright (local), plus Testcontainers if needed
- Deployment: Coolify (self-hosted PaaS) or Dokku on a VPS
- CI/CD: GitHub Actions building Docker images + deployment via SSH

Ideal when: you want a fully offline SaaS stack, shareable with `docker compose up`, without vendor lock-in, while remaining production-ready.

## How to start building your SaaS

Getting started doesn't need to be long or complicated. The goal: remove friction, create momentum, and ship something concrete as early as possible. The steps below are intentionally simple and work with each of the three vibe coding stacks.

### 1. Choose your stack and create the accounts first

Start by choosing the stack you'll use. Don't overthink it. You can always migrate later, but changing tools mid-development breaks momentum.

- Sprint Stack: Vercel, Neon, Clerk, Stripe
- Leverage Stack: Vercel, Supabase, Resend, Stripe, PostHog, Sentry
- Control Stack: Docker, Docker Compose, Git. No hosted accounts required

Having accounts ready from the start avoids context switching once building begins.

### 2. Set up accounts and tools early

After choosing your stack, set up the key services you'll need from day one. These are usually authentication, database access, email sending, payments, and deployment. Creating these accounts early lets you integrate them naturally as development progresses, rather than grafting them on afterward.

Depending on your stack, you'll generally need to create accounts and generate API keys for:

- Deployment: Vercel
- Database: Neon or Supabase
- Authentication: Clerk or Supabase Auth
- Email: Resend
- Payments: Stripe
- Analytics: PostHog
- Error tracking: Sentry

Once the accounts are created, generate the necessary API keys and URLs and store them in environment variables. Keep these values in a local `.env` file and never commit secrets. Only version safe configurations, such as example files or public keys.

Doing this early smooths your development and avoids breaking changes as launch approaches.

### 3. Start your project with a starter

Don't start from an empty repository. Use a starter template that already includes your framework, a base layout, and essential configuration.

A good starter should give you:

- A working frontend layout
- Auth wiring
- A basic routing and API structure

You'll save hours of setup and focus on product logic rather than boilerplate.

### 4. Set up Claude Code as your development copilot

Claude Code is most effective if you treat it as an interactive pair programmer, not just a code generator. Using the Claude Code extension in VS Code gives you fast feedback, solid context, and a more conversational way to build.

When starting a new project, always begin in plan mode and make sure reasoning is enabled. Claude then thinks through the architecture, data flows, and edge cases before writing code. You get a cleaner structure and fewer rewrites later.

Also connect the essential MCP servers early. At a minimum, enable:

- Web search and scraping for up-to-date references
- GitHub search for patterns and examples
- Specific MCPs (Supabase, Vercel, Stripe, Postgres)

Claude can then reason with real documentation and real-world implementations, rather than "guessing."

Before any code is written, always provide a clear and detailed overview of what you're building. Be explicit about:

- The product idea and primary use case
- The chosen tech stack
- Constraints and non-goals

This context anchors the AI from the start and prevents misunderstandings.

Example starting prompt:

```
Build a minimal, production-ready course-selling SaaS using Next.js + Supabase + Stripe + Vercel.
Must include marketing pages, auth, Stripe subscriptions (Checkout + webhooks + portal), dashboard, and protected course content.
Keep it minimal, secure (server-only secrets, webhook verification), and deployable.
```
### 5. Deploy early and keep the app online

Deploy as soon as your app runs locally.

Connect your repository to the deployment platform (Vercel), enable preview deployments, and push frequently. A live URL lets you catch environment issues early and makes it easier to share progress and gather feedback.

The earlier your app is online, the better your decisions will be.

### 6. Iterate from real feedback, not assumptions

As soon as real users are using your product, let their behavior guide your roadmap.

Focus on:

- Where users drop off
- What confuses them
- Which features they actually use

Fix the most painful point first, then repeat. Small iterations add up faster than big rewrites.

### 7. Add monitoring, tests, automations, and behavior tracking

Once your SaaS is live and being used, visibility becomes more important than new features. This is the time to make the product reliable, predictable, and easier to scale.

Start by adding monitoring and error tracking so you're alerted before your users are. Pair that with basic automated tests to protect critical paths (signup, login, payment) as you iterate.

