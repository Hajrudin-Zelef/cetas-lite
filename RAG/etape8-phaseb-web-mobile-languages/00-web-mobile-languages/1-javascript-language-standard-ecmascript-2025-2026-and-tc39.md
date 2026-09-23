---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/1-javascript-language-standard-ecmascript-2025-2026-and-tc39
title: "1. JavaScript language standard — ECMAScript 2025 / 2026 and TC39 proposals"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2025-03", "2025-06", "2025-08", "2026-03"]
keywords: ["research"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [46, 96]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 3de9b7918401efb732cd651131e1881db9a188b816dd0ec76ed6c332613fd373
---

# 1. JavaScript language standard — ECMAScript 2025 / 2026 and TC39 proposals

## 1. JavaScript language standard — ECMAScript 2025 / 2026 and TC39 proposals

### 1.1 ECMAScript 2025 (ES2025) — ratified June 2025

- ES2025 was approved by TC39 and finalized in June 2025 as the year's ECMAScript edition [secondary].
- Iterator helpers: a built-in global `Iterator` with lazy functional operators (`map`, `filter`, `take`, `drop`, etc.) that wrap any iterator without creating intermediate arrays; the implementation evaluates element-wise rather than producing "working arrays" at each stage [secondary].
- Set methods: `union`, `intersection`, `difference`, `symmetricDifference`, `isSubsetOf`, `isSupersetOf`, `isDisjointFrom` — collection operations on `Set` without manual filtering logic [secondary]. ⚠️ Non-comparability note: some sources file Set methods under ES2024 (Stage 4 reached in 2024) while others bundle them into ES2025 write-ups; treat the edition boundary for Set methods as `[unverified]`.
- `Promise.try()`: static method that runs a synchronous function and wraps the result in a Promise, intended to simplify mixed sync/async error handling [secondary].
- Import attributes: JSON module imports (`import data from "./x.json" with { type: "json" }`) standardized in ES2025; browser support remained uneven at mid-2026, so runtime JSON loading via `fetch` was still the safer choice in shipped browser code [secondary].
- Regular-expression upgrades: duplicate named capturing groups (previously a `SyntaxError`), regex pattern modifiers (inline control of flags within subexpressions), and `RegExp.escape()` for safe interpolation of strings into patterns [secondary].
- `Float16Array` typed array plus `DataView` `getFloat16`/`setFloat16` and `Math.f16round`: 16-bit half-precision floats aimed at graphics and ML workloads [secondary].
- `Error.isError()` and `Atomics.pause()` were reported as part of the ES2025 feature list [secondary]; `DurationFormat` (Intl) was reported alongside as well [secondary].
- Exact sources: https://pawelgrzybek.com/whats-new-in-ecmascript-2025/ and https://www.infoworld.com/article/4021944/ecmascript-2025-the-best-new-features-in-javascript.html [secondary].

### 1.2 ECMAScript 2026 trajectory (proposals, not yet ratified)

- Temporal (date/time API): reached Stage 4, i.e. accepted for ECMAScript 2026; at mid-2026 it shipped in Firefox 139+ and Chrome 144 (early 2026), while Safari support was still partial/flagged — NOT cross-browser safe without a polyfill or explicit support check [secondary].
- Upsert (object property update-or-insert) and other proposals were still "in the mix" at the time of the ES2025 write-ups and had not been ratified [secondary].
- Proposal-stage caveat: a feature reaching Stage 4 is accepted for the next edition, but features at Stage 3 or below can change or be rejected; do not treat ES2026 candidate features other than Temporal as settled [secondary].

### 1.3 Baseline availability guidance (mid-2026)

- ES modules run natively in all current browsers (`<script type="module">`, static `import`/`export`, dynamic `import()`); top-level `await` works in modules but blocks the module graph [secondary].
- `fetch`, `AbortController`/`AbortSignal`, `structuredClone`, `Array.prototype.at()`, `Object.groupBy()`/`Map.groupBy()`, `Array.fromAsync()`, `Promise.withResolvers()`, `String.prototype.replaceAll` are reported as broadly available in current browsers [secondary].
- Recommended caution: JSON module imports standardized in ES2025 but with uneven browser support; verify before use [secondary].

---

## 2. TypeScript — 5.8, 5.9, and the 6.0/Go-rewrite claims

### 2.1 TypeScript 5.8 (March 2025)

- `--erasableSyntaxOnly`: a compiler flag aimed at type-stripping workflows so that TypeScript code can be checked for syntax erasability [secondary].
- Node 18 module mode and `require()` of ESM under `nodenext` module settings were headline changes [secondary].
- Exact source: https://www.untergletscher.com/en/blog/typescript-6-0-new-features-dx-guide-2026 (5.8 recap in the same piece) [secondary].

### 2.2 TypeScript 5.9 (August 2025)

- `import defer` (deferred imports) added as a language feature [secondary].
- `--module node20` module resolution mode added [secondary].
- Minimal `tsc --init` output and expandable editor hovers were developer-experience additions [secondary].
- Exact source: https://progosling.com/en/dev-digest/2025-08/typescript-5-9-release [secondary].

### 2.3 TypeScript 6.0 and the Go rewrite — claims requiring care

- Several secondary sources claimed TypeScript 6.0 shipped in March 2026 as "the last JavaScript-based compiler", with TypeScript 7 as a native Go rewrite in preview [secondary].
- ⚠️ These claims were NOT verified against Microsoft's official release page in this research session; tag every TS 6.0/7 statement `[secondary]` or `[unverified]` and do not present the Go-rewrite timeline as established fact.
- Exact source of the claim: https://www.untergletscher.com/en/blog/typescript-6-0-new-features-dx-guide-2026 [secondary].

---

