---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/3-node-js-22-24-26-lifecycle-and-runtime-features
title: "3. Node.js — 22/24/26 lifecycle and runtime features"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["Anthropic"]
dates: ["2025-05-06", "2025-06", "2026-09", "2026-10", "2027-04-30", "2028-04-30"]
keywords: ["acquisition", "benchmark", "benchmarks", "exploit"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [97, 167]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 5cea0a4c4e3ad3e0dba507196142ab56918cd15560fc347a5a758eeeeceff32b
---

# 3. Node.js — 22/24/26 lifecycle and runtime features

## 3. Node.js — 22/24/26 lifecycle and runtime features

### 3.1 Release lines and lifecycle (as understood at cutoff)

- Node.js 22: Maintenance LTS, expected end-of-life 2027-04-30 [secondary].
- Node.js 24 "Krypton": Active LTS, initial release 2025-05-06, expected EOL 2028-04-30 [secondary].
- Node.js 26: Current release line in September 2026, with a planned transition to Active LTS in October 2026; exact October date conflicted among mirrored release schedules, so the date is reported as "October 2026" only [secondary]. ⚠️ Conflict preserved: sources disagreed on the exact day of the LTS promotion.
- Release policy: odd-numbered Node.js lines do not become LTS; they serve as "Current" lines that lead into the next even-numbered LTS [secondary].
- Mirrored-schedule sources: https://github.com/viroj168/release-node, https://github.com/rashworld-max/Release, https://github.com/zwindler/101-ways-to-deploy-kubernetes/pull/177 [secondary].
- Node.js runtime supports `node --experimental-strip-types` as a path to run TypeScript directly without a separate compiler [secondary].

### 3.2 Non-comparability warning — runtime performance claims

- Do not repeat "99% Node compatibility" style claims or synthetic benchmark speedups as established fact; observed sources repeated them without methodology, and no independent benchmark was reviewed in this session [secondary].

---

## 4. Deno 2.x

- Deno 2.x works with `package.json`, npm specifiers, and private npm registries; it supports import maps and workspaces [secondary].
- JSR (JavaScript Registry) integration is a first-class package source alongside npm [secondary].
- Permission controls (`--allow-*` model) remain the core security posture of the runtime [secondary].
- The global cache removes the requirement for a local `node_modules` directory by default [secondary].
- Sources: https://github.com/hackylabs/deep-redact/blob/HEAD/_bmad-output/implementation-artifacts/5-3-verify-the-deno-baseline-path-and-installation-documentation-lockstep.md, https://github.com/lockgraph/lockgraph/blob/HEAD/spec/pm/deno.md, https://github.com/kt3k/changelog/blob/HEAD/src/posts/2026-06-29_denoland-deno.md [secondary].

---

## 5. Bun 1.3.x

- Secondary sources describe Bun 1.3.x with built-in SQL/database clients and frontend development tooling [secondary].
- Bun 1.3.14 was described as the final release implemented in Zig [secondary].
- ⚠️ Acquisition/rewrite rumor: reports of an Anthropic acquisition of Bun and a subsequent Rust rewrite need cautious provenance — the evidence observed was a single satirical-register piece at https://www.theregister.com/devops/2026/05/14/anthropics-bun-rust-rewrite-merged-at-speed-of-ai/5240381 [secondary]. Do not present the acquisition as confirmed fact.
- Other sources: https://dev.to/last9/is-bun-production-ready-in-2026-a-practical-assessment-181h, https://www.I-Programmer.info/news/98-languages/18587-bun-13-adds-frontend-development-support.html [secondary].

---

## 6. npm / pnpm / Yarn / Corepack — JS package managers and the npm registry

### 6.1 npm (bundled with Node.js)

- npm is the default package manager bundled with Node.js (first released 2010); `package-lock.json` arrived in npm v5 and a large performance overhaul in v7+ [secondary].
- One secondary source stated "as of 2026, npm v10 ships with Node.js 22 LTS" [secondary]; ⚠️ this conflicts with the broader expectation that Node 22 ships npm v11 — treat the exact bundled npm major as `[unverified]` for the 22.x line.
- npm lock file: `package-lock.json`; native workspace support since npm v7 [secondary].

### 6.2 pnpm

- pnpm uses a content-addressable global store: if 10 projects use React 19, only one copy lives on disk (hard-linked), which sources claim saves roughly 70–80% of disk space versus npm [secondary].
- Strict dependency isolation: a package can only access what is listed in its own `package.json`; phantom dependencies are blocked at the filesystem level, unlike npm/Yarn hoisting [secondary].
- Consistently benchmarked as faster than npm for cached installs and monorepo setups [secondary].
- Monorepo ergonomics: `pnpm --filter` and `pnpm -r` (recursive) [secondary].
- **Security model (v10+):** lifecycle scripts (`postinstall`, `preinstall`) are no longer run by default — a large reduction in supply-chain attack surface [secondary].
- pnpm v11 introduced a `minimumReleaseAge` setting (defaulting to 1 day): newly published packages are blocked for 24 hours to give the community time to spot compromised versions [secondary].
- Compatibility: pnpm installs any package from the npm registry; failures of well-maintained packages are "extremely rare" per the source, and the common cause is a package relying on npm hoisting with an incomplete dependency list [secondary].
- Sources: https://dev.to/hamidrazadev/pnpm-vs-npm-vs-yarn-which-package-manager-should-you-actually-use-3ke3, https://dev.to/_d7eb1c1703182e3ce1782/npm-vs-pnpm-vs-yarn-package-manager-showdown-2026-benchmarks-2c38, https://tech-insider.org/pnpm-vs-npm-2026/ [secondary].

### 6.3 Yarn and Corepack

- Yarn (Berry, v4+) remains a maintained alternative with workspace/Plug'n'Play lineage; no major version change was verified in this session — treat the current Yarn major as `[unverified]` [unverified].
- The `packageManager` field in `package.json` (managed by Corepack) is the recommended mechanism to pin a package manager per project so teams stay consistent [secondary].
- Mixing npm and pnpm (or Yarn) in one project is strongly discouraged because each generates its own lock file (`package-lock.json` vs `pnpm-lock.yaml` vs `yarn.lock`) and resolution can diverge [secondary].

### 6.4 npm registry scale (estimates conflict — reported as observed)

- The npm registry was described as "more than 2.5 million packages" in an SQLI supply-chain analysis [secondary]: https://www.sqli.com/int-en/insights/pnpm-vs-shai-hulud-npm-security.
- Another source claimed "over 5 million packages" in a security-best-practices repo [secondary]: https://github.com/bingecode/npm-security-best-practices.
- A June 2025 estimate pegged roughly 1.8 million "available" packages [secondary]: https://tech-insider.org/pnpm-vs-npm-2026/.
- ⚠️ Conflict preserved: sources use different counting methods (total published names vs. currently available versions). Do not quote a single registry-size number without its source and method [secondary].
- JavaScript's ubiquity makes the registry a prime supply-chain target; attackers exploit weak maintainer account security and hijack dormant/orphaned packages with transitive reach [secondary].

---

