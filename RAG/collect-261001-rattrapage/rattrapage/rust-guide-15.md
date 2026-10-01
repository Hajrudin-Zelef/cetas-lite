---
id: collect-261001-rattrapage/rattrapage/rust-guide-15
title: "Guide Rust ultra-complet — Langage système sûr, rapide et moderne"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/rust_guide.md
source_anchor: ""
source_lines: [3347, 3365]
sha256: b965d6362bb7e34ca19af860953a0aa7b7a506dfb1961696a4272d13ed16d942
---

# Guide Rust ultra-complet — Langage système sûr, rapide et moderne

| Crate | Usage |
|---|---|
| `clap` | CLI |
| `serde` / `serde_json` / `toml` / `csv` | Formats de données |
| `tokio` / `reqwest` | Async + HTTP |
| `tracing` | Logs structurés |
| `anyhow` / `thiserror` | Erreurs |
| `sysinfo` | CPU/RAM/disques/processus (cross-platform) |
| `walkdir` | Parcours récursif |
| `notify` | Watch de fichiers (inotify) |
| `cron` (crate `cron`) | Planification façon cron |
| `nix` | Appels Unix idiomatiques |
| `prometheus` | Métriques Prometheus |

---

*Fin du guide. Le borrow checker est ton ami le plus exigeant : s'il râle, c'est qu'il vient
de t'éviter un incident de prod. Bon code, et que tes binaires soient statiques et tes
panics absentes.* 🦀
