---
id: etape8-phaseb-web-mobile-languages/00-web-mobile-languages/7-supply-chain-security-shai-hulud-and-registry-hardening
title: "7. Supply-chain security — Shai-Hulud and registry hardening"
domain: step-8-phase-b-web-mobile-languages-runtimes-and-toolchains
role: deep-dive
task: reference
actors: ["Hugging Face", "OpenAI"]
dates: ["2025-01", "2025-10-07", "2025-11", "2026-09-22", "2026-10"]
keywords: ["acquisition", "agent", "benchmarks", "consumer", "inference", "latency", "research", "throughput"]
source: docs/RAG/etape8_phaseB_web_mobile_languages.md
source_anchor: ""
source_lines: [168, 236]
section: "Step 8 — Phase B — Web & Mobile Languages, Runtimes and Toolchains"
sha256: 19e19ec700b3b3919a97957ad172a1ba02da72f9b700963b5f31d4fc3eb288f9
---

# 7. Supply-chain security — Shai-Hulud and registry hardening

## 7. Supply-chain security — Shai-Hulud and registry hardening

- Shai-Hulud is a self-replicating worm targeting the npm registry: (1) a phishing campaign impersonating the npm registry harvests maintainer credentials; (2) malicious code is injected into affected packages and triggers on the consumer's next `npm install`; (3) the payload scans for secrets (tokens, env vars, API keys), exfiltrates them, and publishes them to a public GitHub repo under the victim's account; (4) it self-spreads via stolen credentials of maintainers/contributors [secondary].
- In November 2025 a more aggressive variant appeared that leveraged the `preinstall` hook for earlier execution and attempted to delete the victim's home directory if exfiltration failed [secondary].
- A documented victim was `@ctrl/tinycolor`, a package with several million weekly downloads [secondary].
- Source: https://www.sqli.com/int-en/insights/pnpm-vs-shai-hulud-npm-security [secondary].
- npm registry defenses: mandatory 2FA for new scoped packages and scoped publish tokens [secondary]; legacy packages and low adoption of the new controls remain a gap [secondary].
- PyPI-side measures: `pip-audit`, virtual environments, requirements pinning; maintainers urged to use 2FA and upload-scoped API tokens; typosquatting (e.g. "requestts", "urllibs") is a persistent vector [secondary]: https://forum.gnoppix.org/t/the-next-wave-of-supply-chain-attacks-npm-pypi-and-docker-hub-incidents-set-the-stage-for-2026/2983.
- Practical mitigations circulated: scrutinize changelogs, use `npm audit`, lock dependencies (`package-lock.json`), and prefer built-in runtime modules (`node:test`, `node --watch`, `node --env-file`, native `fetch`) over third-party packages where equivalent [secondary]: https://github.com/bingecode/npm-security-best-practices.

---

## 8. Python — 3.13 / 3.14, free threading, JIT

### 8.1 Release status at cutoff

- Python 3.14 was released 2025-10-07 according to multiple secondary sources [secondary].
- Python 3.13 introduced experimental free-threading (no-GIL) builds and an experimental JIT [secondary].
- Free-threaded mode in 3.14 is supported but remains opt-in, not the default interpreter [secondary].
- Python 3.15 was expected October 2026 (one release cycle after 3.14); not released by the 2026-09-22 cutoff — do not describe 3.15 features as shipped [unverified].

### 8.2 Python 3.14 highlights

- PEP 779: free-threading support continued; PEP 750: template strings; PEP 649: deferred evaluation of annotations; PEP 734: multiple interpreters in the stdlib; `compression.zstd` added; external debugger interface [secondary].
- Third-party benchmarks claim a remaining ~5–10% average single-thread penalty for free-threaded builds with multi-core gains, but methodologies vary widely — keep the penalty figure `[independent]`-at-best and never mix benchmarks with different workloads [secondary].
- Sources: https://github.com/mr-pylin/python-workshop/blob/HEAD/CHANGELOG.md, https://github.com/byronwilliamscpa/rag-processor/blob/HEAD/docs/PYTHON_COMPATIBILITY.md, https://dev.to/dmaxdev/python-314-free-threading-real-benchmarks-real-breakage-real-code-3m5 [secondary].

### 8.3 Python in AI/ML workloads (context)

- Python's dominance in AI/ML rests on the framework ecosystem: TensorFlow, PyTorch, Hugging Face, NumPy, pandas, scikit-learn are all Python-first [secondary].
- The AI boom is the primary driver of Python's TIOBE #1 position and its Octoverse ranking (see §21) [secondary].

---

## 9. Python tooling — uv, Ruff, ty, Poetry, and the Astral acquisition

### 9.1 uv — the momentum leader in Python package/project management

- uv is a Rust-based Python package and project manager from Astral (the same team as Ruff/ty); it is positioned as a replacement for pip + pip-tools + pipx + poetry + pyenv + virtualenv in a single binary [secondary].
- Project mode: `uv init`, `uv add`, `uv lock`, `uv sync`; pip-compatible mode: `uv pip install` [secondary].
- `uv.lock` pins every transitive dependency with hashes for reproducible environments across machines and CI [secondary].
- Python version management: `uv python install 3.12` downloads prebuilt CPython binaries (via the python-build-standalone project) in seconds instead of compiling from source (pyenv default: 5–15 minutes); `uv python pin 3.12` locks a project to a version [secondary].
- `uvx` runs ephemeral tools without installing them (e.g. `uvx ruff check`); `uv tool install` persists global tools [secondary].
- PEP 723 script support: `uv run script.py` reads inline `# /// script` metadata and creates an ephemeral environment [secondary].
- Speed claims: 10–100× faster installs than pip; 8–115× depending on cache state; `uv sync` on a warm cache takes milliseconds [secondary]. (Vendor-adjacent claims; treat throughput multiples as `[secondary]`.)
- Adoption signals (early 2026, all secondary): ~36,100 GitHub stars; roughly 13% of all PyPI package downloads served to uv; Wagtail project CI reported 66% uv vs 34% pip [secondary].
- Competitive landscape: Poetry ~31K stars (Poetry 2.0, January 2025, finally adopted PEP 621 `[project]` table), PDM ~7.5K stars (standards-focused), Hatch ~6K stars (no lockfile support yet) [secondary].
- New projects increasingly default to uv; existing Poetry teams are advised staying is reasonable [secondary].
- Sources: https://github.com/djbclark/stayturgid/blob/HEAD/docs/research/python-tooling-uv-ty-ruff.md, https://github.com/rgesteves5/ai-arch-toolkit/blob/HEAD/research/modern_python_2015_16.md, https://github.com/jajupmochi/agent-harness/blob/HEAD/tooling/python-uv-ruff/README.md, https://medium.com/@naveenr3830/stop-using-pip-try-uv-%EF%B8%8F-10-100x-faster-dependency-installs-8eebbdaa9587 [secondary].

### 9.2 Astral acquisition by OpenAI (early 2026)

- OpenAI acquired Astral (the team behind uv, Ruff, and ty) in early 2026; the engineering team joined OpenAI's Codex division [secondary].
- The strategic read offered by the source: AI coding assistants need deep control over environment setup, linting, and type inference infrastructure — uv/ruff/ty becoming part of the Codex stack enables AI-generated code to auto-create virtualenvs, install dependencies, and validate types with low latency [secondary].
- ⚠️ The acquisition report came from an AI-news aggregator (https://aihaberleri.org), not from OpenAI's or Astral's official channels; treat the acquisition as `[secondary]` and the "Codex auto-environment" scenario as `[unverified]` speculation.
- Source: https://aihaberleri.org (OpenAI Acquires Astral) [secondary].

### 9.3 Ruff, ty, and the lint/format consolidation

- Ruff implements 800+ lint rules natively in Rust, replacing Flake8, isort, pyupgrade, autoflake, and pydocstyle [secondary].
- `ruff format` achieves >99.9% compatibility with Black while running ~30× faster [secondary].
- Projects such as FastAPI, pandas, Django, and Pydantic have switched to Ruff [secondary].
- Pylint retains an edge in deep cross-file analysis (~200 rules Ruff does not cover); many teams now treat Ruff plus a type checker as sufficient [secondary].
- Black remains maintained, but `ruff format` is described as the recommended choice for new projects [secondary].
- ty is Astral's modern Python type checker, built for speed and developer experience [secondary].
- Canonical modern stack per practitioner guides: uv (package + Python version management) + Ruff (lint + format), replacing `pip + venv + black + flake8 + isort` [secondary].

---

