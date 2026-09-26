---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/19-elixir-1-19
title: "19. Elixir 1.19"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: []
dates: ["2025-10-16", "2025-12-23", "2025-12-27", "2026-07", "2026-07-23", "2026-09", "2026-09-10"]
keywords: ["advisory", "bedrock", "inference", "memory", "research"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [408, 482]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 1680e47cb4910a4d2337df371543642a9ae93572adedbd51c83fda52da46be4c
---

# 19. Elixir 1.19

## 19. Elixir 1.19

- Elixir 1.19 release post dated 2025-10-16 [secondary].
- Enhanced type inference and type checking for anonymous functions and protocols [secondary].
- Up to 4× faster compilation for large projects (vendor-reported figure; methodology not reviewed) [vendor-reported].
- Ecosystem snapshot (advisory, not an official compatibility matrix): Elixir 1.19 with OTP 28, Phoenix 1.8, LiveView 1.1, Ecto 3.13, and Bandit [secondary].
- Sources: https://github.com/elixir-lang/elixir-lang.github.com/blob/HEAD/src/content/blog/elixir-v1-19-0-released.md, https://github.com/sztheory/mailglass/blob/HEAD/prompts/elixir-plug-ecto-phoenix-system-design-best-practices-deep-research.md, https://elixirmerge.com/p/overview-of-elixir-v1-19-release-and-related-updates [secondary].

---

## 20. Lua 5.5 / 5.4.9

- Lua 5.5.0 released 2025-12-27 (per NuGet metadata) [secondary]; Lua 5.5.1 released 2026-07-23 (per the lua/lua GitHub releases page) [secondary]; Lua 5.4.9 released 2026-09-10 (maintenance of the 5.4 line) [secondary].
- New in 5.5: explicit global variable declarations (avoiding implicit-global errors) [secondary].
- For-loop variables are now read-only, reducing unintended side effects [secondary].
- Compact arrays: ~60% memory-footprint reduction for large tables, per the release coverage [secondary].
- Garbage collection: new generational GC mode plus incremental major GC cycles, reducing pause times for long-running/real-time applications [secondary].
- Float values print with enough digits for exact read-back (serialization safety) [secondary].
- String/UTF-8 updates: expanded `utf8.offset`, support for external strings using non-Lua-managed memory [secondary].
- New C API utilities: `luaL_openselectedlibs`, `luaL_makeseed`; `LUA_NOBUILTIN` build option to avoid `__builtin_expect` in the Lua API; the `lua.c` interpreter dynamically loads readline when available [secondary].
- Tooling: lua-language-server 3.16.3 (2025-12-23) upgraded to Lua 5.5 with ~10% memory reduction; 3.18.x (2026-04) added Lua 5.5 syntax support (`local <close>`, `<const>`, `global` identifier rules) [secondary].
- Lua remains MIT-licensed, PUC-Rio originated (1993), embed-first with intentional incompatible changes between major releases [secondary].
- Sources: https://github.com/lua/lua/releases, https://linuxiac.com/lua-5-5-released-with-incremental-garbage-collection-and-compact-arrays/, https://www.nuget.org/packages/lua/5.4.6, https://github.com/luals/lua-language-server/blob/HEAD/changelog.md, https://archlinux.org/packages/extra/x86_64/lua/ [secondary].

---

## 21. Adoption rankings 2026 — TIOBE, Stack Overflow, GitHub Octoverse

### 21.0 Non-comparability warnings (read first)

- TIOBE measures search-engine/query activity about languages (popularity of *discussion*), Stack Overflow surveys self-selected developers about languages they *use*, and GitHub Octoverse measures repository/contributor *activity* on one platform. These are different populations answering different questions — ranks are not interchangeable and deltas across sources are not meaningful [secondary].
- Survey respondents are self-selected; weighting is sometimes applied for demographic skew (Stack Overflow notes this explicitly) [secondary].
- AI-generated and AI-assisted code inflates repository-activity metrics without necessarily reflecting human language preference [secondary].

### 21.1 TIOBE Index — September 2026

| Rank | Language | Rating | YoY change | Tag |
|---|---|---|---|---|
| 1 | Python | 17.76% | -8.22% | [secondary] |
| 2 | C | 10.28% | +1.63% | [secondary] |
| 3 | C++ | 8.67% | -0.13% | [secondary] |
| 4 | Java | 7.54% | -0.81% | [secondary] |
| 5 | C# | 4.22% | -2.16% | [secondary] |
| 6 | JavaScript | 2.76% | -0.46% | [secondary] |
| 7 | Visual Basic | 2.55% | -0.28% | [secondary] |
| 8 | SQL | 2.16% | +0.29% | [secondary] |
| 9 | R | 1.69% | +0.27% | [secondary] |
| 10 | Rust | 1.34% | +0.33% | [secondary] |
| 11 | Fortran | 1.24% | -0.25% | [secondary] |
| 12 | Go | 1.10% | -1.22% | [secondary] |
| 13 | Delphi/Object Pascal | 1.08% | -1.18% | [secondary] |
| 14 | PHP | 1.04% | -0.21% | [secondary] |
| 15 | Scratch | 0.99% | -0.19% | [secondary] |
| 16 | Assembly language | 0.89% | -0.15% | [secondary] |
| 17 | Ada | 0.85% | -0.42% | [secondary] |
| 18 | Swift | 0.83% | +0.10% | [secondary] |
| 19 | Objective-C | 0.81% | +0.32% | [secondary] |
| 20 | COBOL | 0.78% | -0.13% | [secondary] |

- Table source: https://www.tiobe.com/tiobe-index (via secondary write-ups: https://www.techrepublic.com/article/news-tiobe-september-2026-julia-nears-top-20/ and https://www.swapupdate.in/tiobe-index-september-2026-julia-nears-top-20/) [secondary].
- Notes: the top 10 held identical positions from August to September 2026; Python fell from 18.53% to 17.76% but kept a large lead; C++ widened its gap over Java (1.13 points); Rust held No. 10 for a third consecutive month (its first-ever top-10 run, entering July 2026); Fortran sits at No. 11 only 0.10 points behind Rust [secondary].
- Outside the top 20: Julia ranked No. 21 at 0.74% (nearing a top-20 return; TIOBE CEO Paul Jansen said Julia is taking some of MATLAB's numerical/scientific territory), MATLAB fell to No. 27, Ruby fell to No. 22, Perl to No. 23 [secondary].
- TIOBE methodology: ratings track programming-language popularity using search activity across major websites and search engines, oriented to skilled engineers, courses, and third-party vendors [secondary].
- Very-long-term context (12-month averages): 2026 top-10 ordering was Python, C++, C, Java, C#, JavaScript, Visual Basic, Go, Delphi/Object Pascal, SQL [secondary].

### 21.2 Stack Overflow Developer Survey

- 2025 survey: JavaScript remained the most-used language at 63.61% of respondents (87,585 answers), HTML/CSS second at ~53%, Python third at 49.28% — Python had been climbing for three years [secondary].
- 2026 survey: Python reportedly overtook JavaScript — 38% of developers named Python vs 36% JavaScript, ending JavaScript's run in the top spot (held since 2014); Python jumped ~7 points in one cycle, the biggest move on the language chart in five years [secondary].
- ⚠️ The 2026 figures (38 vs 36) come from a secondary write-up (https://python.plainenglish.io/python-has-finally-taken-the-spot-from-javascript-most-people-get-the-career-math-wrong-9df8ec4ad08a) that itself warns: the survey is self-selected (~49,000 respondents across 177 countries), it measures *usage* not hiring demand, and the question/methodology differs from prior years — do not treat "Python overtook JavaScript" as a like-for-like time series with the 2025 63.61% figure [secondary].
- 2026 directional signals: Rust and Go kept climbing; JS/TS remain the web bedrock; ~8 in 10 developers report using GPT-family models for development tasks [secondary]: https://dev.to/dhruvjoshi9/what-devs-are-actually-learning-in-2025-stack-overflow-dora-insights-17j3.
- Learning-methodology note: online resources (docs, Stack Overflow, blogs) dominate; tutorial videos fell off the podium [secondary].

### 21.3 GitHub Octoverse

