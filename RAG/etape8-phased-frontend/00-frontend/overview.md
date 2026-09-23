---
id: etape8-phased-frontend/00-frontend/overview
title: "Step 8 — Phase D: Frontend & Web Platform (Dev Angle)"
domain: step-8-phase-d-frontend-web-platform-dev-angle
role: deep-dive
task: reference
actors: []
dates: ["2024-04", "2024-12-05", "2025-10", "2025-12", "2025-12-18", "2026-04", "2026-07", "2026-08-06", "2026-08-12", "2026-09-22", "2026-10", "2026-10-21"]
keywords: ["benchmark", "research"]
source: docs/RAG/etape8_phaseD_frontend.md
source_anchor: ""
source_lines: [1, 48]
section: "Step 8 — Phase D: Frontend & Web Platform (Dev Angle)"
sha256: e7cd0a732399a32488e031ec109f7e3d56b4b6e1411207172e4196873b2e61a1
---

# Step 8 — Phase D: Frontend & Web Platform (Dev Angle)

- **Scope:** React/Next.js ecosystem, Vue/Nuxt, Svelte/SvelteKit, Angular, SolidJS/Qwik/Astro, HTML-CSS 2026 platform, Tailwind CSS v4, styling state, build tooling, rendering patterns, WebAssembly, HTMX/Alpine, state management, accessibility + Core Web Vitals, 2026 ecosystem surveys.
- **Research date / cutoff:** 2026-09-22. English.
- **Method:** read-only web research (search + page fetch). Append-only writing, single writer; no other workspace files modified.
- **Provenance legend:** `[official]` = vendor/framework official docs or release notes; `[vendor-reported]` = vendor blog/PR claims; `[independent]` = third-party benchmark/study; `[secondary]` = press, blogs, docs mirrors, community research; `[unverified]` = single-source or unconfirmed.
- **Non-goals (covered elsewhere):** ops/deployment angle of the same tools (Step 7), backend languages and AI/ML stacks (other Step 8 phases), package-manager security (Step 7G).

## D1. React 19.x — current stable line

- React 19.0.0 stable released **2024-12-05** (RC April 2024) — first major since React 18, built around Actions, Server Components and the Compiler [secondary].
- Patch train through 2026: **19.2.5** reported April 2026 [secondary]; **19.2.8** reported as latest stable 19.x on 2026-08-12 [secondary]; 19.3 only in canary at that point [secondary].
- **Server Components (RSC) stable** — the stable RSC payload format is not HTML but a serialized data structure (objects, arrays, Dates, Maps, Promises, new React trees, Server Function references) [secondary].
- **Actions API stable** — `useActionState`, `useOptimistic`, `useFormStatus`, native `<form action={...}>`; pending state, error handling and optimistic updates first-class [secondary].
- **`use()` hook stable** — reads promises and context conditionally [secondary].
- **ref as a prop** — function components accept `ref` directly; `forwardRef` no longer needed [secondary].
- **Document metadata** — `<title>`, `<meta>`, `<link>` rendered in components and hoisted to the document head [secondary].
- **Resource loading** — stylesheet precedence via `<link precedence>` and async `<script>` deduplication built in [secondary].
- **React Compiler stable (2025)** — automatic memoization replacing most manual `useMemo` / `useCallback` / `React.memo` [secondary].
- **`<Activity />` component** (React 19.2) — manages hidden/visible UI while preserving component state [secondary].
- **Improved hydration** — tolerates unexpected elements injected by browser extensions/third-party scripts; single consolidated hydration error with diff [secondary].
- Concurrent rendering is the default model (`startTransition`, `useDeferredValue`, Suspense compose naturally) [secondary].
- Security note: a **React2Shell / RSC RCE cluster** was tracked in December 2025, distinct from the Next.js image-pipeline issue [secondary].
- TypeScript pairing: `@types/react` 19.x line (19.2.14 cited) [secondary]; TypeScript 5.x required for Next 15.5 (rejects TS >= 7.0) [secondary].
- Migration guidance commonly cited: install `react` and `react-dom` together at the same version; mismatched versions cause subtle runtime errors [secondary].

## D2. Next.js 15 / 16 — App Router maturity

- **Next 16.x stable line released October 2025**; 16.3.x is the current stable line in 2026 [secondary].
- Next 15 line: 15.5.23 (2026-08-06) reported; **15.x security support ends 2026-10-21** — plan migration before October 2026 for greenfield or long-lived apps [secondary].
- Next 15 headline features: async request APIs (`cookies()`, `headers()`, `params()`), React 19 support, stable Turbopack, smarter caching, codemod CLI (`npx next codemod`) [secondary].
- **Cache Components (beta)** — unifies experimental caching (Dynamic IO, `use cache`, Partial Prerendering) behind one flag [secondary].
- **Next.js 16.1 (2025-12-18)** — Turbopack file-system caching stable and on by default for `next dev`; experimental built-in bundle analyzer (`next experimental-analyze`); `next dev --inspect` [secondary].
- Next 16 stabilizes **Turbopack builds** (`next build --turbopack` beta, 2-5x faster builds claimed by press) and ships Cache Components for hybrid static/dynamic rendering [secondary].
- Next 16 breaking changes: Node.js 20+ required (Node 18 deprecated), AMP deprecated, select `next/image` API changes [secondary].
- Node.js Middleware promoted to stable (introduced experimental in 15.2); Deployment Adapters in alpha for custom build/deploy targets [secondary].
- **Security, July 2026 release:** two issues shipped — (1) AVIF image config issue via transitive native dep (`sharp` → `libheif`); (2) **CVE-2026-75604**, unauthenticated RCE on **Windows-hosted** servers using Pages + App Router without Cache Components (path traversal CWE-22), CVSS 3.1 **9.0**, GHSA-p293-qw3h-jr36 [secondary].
- Fixes: `next@15.5.24` and `next@16.3.3`; Vercel-hosted apps protected without upgrade; self-hosted Docker/VM/Windows must bump [secondary].
- AI-assisted development is a first-class Next.js theme: Vercel AI SDK positioning for LLM apps; DevTools browser-log forwarding aimed at AI debugging workflows [secondary].
- TanStack Start is cited as a modern alternative to Next.js (file-based routing + RSC + server functions) [secondary].

## D3. React Native / Expo (2026 status)

- React Native remains the primary cross-platform native path for React teams; **Expo SDK** is the default managed workflow in 2026 guides [secondary].
- The **New Architecture (Fabric renderer + TurboModules)** is the production default; legacy bridge apps are migration candidates [unverified].
- Not a web technology per se — included here because the React hiring pool and component-model knowledge transfer directly [secondary].
- Alternatives cited in 2026 comparisons: Ionic/Capacitor (web-view hybrid) [secondary].
- React Native Web / Tamagui-style universal UI is a niche for code-sharing between web and native [unverified].
