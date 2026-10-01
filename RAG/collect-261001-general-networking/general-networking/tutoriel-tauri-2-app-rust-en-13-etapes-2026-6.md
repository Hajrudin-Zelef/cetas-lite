---
id: collect-261001-general-networking/general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026-6
title: "Sous Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026.md
source_anchor: ""
source_lines: [634, 666]
sha256: b193a10e9df699e9ce5965d593f3415841859172e8de45ba69b126aeacbb1a3a
---

# Sous Linux/macOS

Oui, Tauri est agnostique au framework frontend. Le template `create-tauri-app` propose React, Vue, Svelte, SolidJS, Angular, Preact, Yew (Rust) et HTML/CSS/JS vanilla. Vous pouvez aussi greffer Tauri sur un projet Vite, Next.js statique ou Nuxt 4 existant en pointant `frontendDist` vers le répertoire de build.

### Quelle est la différence entre Tauri 1.x et Tauri 2 ?

Tauri 2 introduit le support mobile (Android et iOS), un nouveau système de capabilities remplaçant l’allowlist, des plugins officiels stabilisés (SQL, Notification, Store, Updater, etc.), et une refonte de l’API JavaScript en modules ESM. La migration depuis Tauri 1.x s’effectue avec `npx tauri migrate` qui transforme automatiquement `tauri.conf.json` et la plupart des appels JS.

### Combien coûte un certificat de signature de code pour Windows et macOS ?

Un certificat Apple Developer coûte 99 € par an. Pour Windows, un certificat OV (Organization Validation) revient à 200-400 € par an chez Sectigo, DigiCert ou Certigna, et un EV (Extended Validation) qui supprime immédiatement les avertissements SmartScreen coûte 400-800 € par an. Pour Linux, aucune signature obligatoire – un AppImage signé GPG suffit.

### Tauri 2 est-il prêt pour la production en 2026 ?

Oui. La version stable 2.0 est sortie le 2 octobre 2024, suivie d’une cadence de releases quasi mensuelle sur le noyau lui-même : 2.5.1 le 21 avril 2025, 2.6.0 le 24 juin 2025, puis toute la série 2.9.x à l’automne 2025 dont la 2.9.1 le 22 octobre 2025, jusqu’à l’actuelle 2.11.1 (Tauri Releases). Le runtime sous-jacent suit le même rythme soutenu : `tauri-runtime-wry` a atteint la v2.8.1 le 25 août 2025, puis la v2.9.2 le 30 novembre 2025 et la v2.9.3 le 9 décembre 2025, d’après la page de releases du runtime Tauri. Plusieurs applications grand public l’utilisent en production (Ente Photos, AppFlowy, Relay), un audit de sécurité externe a été réalisé avant la sortie 2.0, et la Tauri Foundation gouverne le projet sous l’égide du Commons Conservancy.

### Quelle est la différence entre tauri-plugin-sql et SQLx ?

Le plugin officiel expose une API JavaScript simple et fonctionne immédiatement avec SQLite, MySQL ou PostgreSQL. SQLx est une librairie Rust avancée avec vérification SQL à la compilation et support du pooling. Pour un prototype ou un projet de petite taille, le plugin suffit largement. Pour une application bureau qui synchronise avec un Postgres distant et exécute des requêtes complexes, SQLx en Rust apporte plus de robustesse.

### Tauri 2 fonctionne-t-il avec WSL 2 sous Windows 11 ?

Oui, mais uniquement pour le développement de cibles Linux. Pour produire des binaires Windows depuis WSL, vous devez utiliser un cross-compiler ou exécuter `tauri build` directement sous Windows. La majorité des développeurs sous Windows préfèrent une stack 100 % Windows native (PowerShell, Visual Studio Build Tools).

## Verdict : pourquoi Tauri 2 mérite votre prochain projet bureau

En près de deux ans d’existence stable, jalonnés d’une dizaine de versions mineures publiées sans interruption jusqu’à la 2.9.5 du 9 décembre 2025 selon les Tauri Releases, puis jusqu’à la révision 2.11.1 de l’API atteinte en juillet 2026 d’après le blog ESB1995, Tauri 2 s’est imposé comme l’alternative crédible et mature à Electron pour les applications bureau modernes. Les chiffres parlent : binaires 30 à 40 fois plus légers, démarrage 20 à 30 fois plus rapide, empreinte mémoire 5 à 10 fois inférieure. Surtout, le framework couvre désormais cinq plateformes (Windows, macOS, Linux, Android, iOS) avec une seule base de code Rust + Web.

Pour une équipe française ou européenne qui démarre un projet bureau aujourd’hui, le choix Tauri 2 réduit les coûts d’infrastructure de distribution (binaires 40 fois plus petits = 40 fois moins de bande passante CDN), simplifie la conformité réglementaire (pas de runtime Chromium embarqué = moins d’enjeux RGPD et CRA), et donne accès à un langage backend partageable avec votre stack cloud (axum, actix-web, sqlx). Les arguments contre se réduisent à la courbe d’apprentissage Rust pour les commandes avancées et à un écosystème de plugins encore plus jeune que celui d’Electron – mais qui croît mensuellement avec le soutien actif de la Tauri Foundation.

Le projet NoteForge complet décrit dans ce tutoriel – gestionnaire de notes local-first avec SQLite, capabilities granulaires, notifications natives, updater Ed25519 et CI/CD GitHub Actions – vous donne un squelette opérationnel pour démarrer votre prochain projet desktop multi-plateforme. Bonne route avec Tauri.

### Couverture associée

Sources et références externes : documentation officielle Tauri 2, dépôt GitHub tauri-apps/tauri, guide sécurité Tauri, site officiel Rust, notice Wikipedia Tauri.
