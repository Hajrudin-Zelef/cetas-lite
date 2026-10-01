---
id: collect-261001-general-networking/general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026-5
title: "Sous Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Intel"]
dates: []
keywords: ["apache", "benchmarks", "cyber", "distribution", "incident", "intel", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026.md
source_anchor: ""
source_lines: [545, 633]
sha256: c19f6bfd3899cfe3616878ba7d8f9aa796b34a04bcfa0194dfc0c335e197f706
---

# Sous Linux/macOS

Troisièmement, exposez l’icône système avec `tauri-plugin-tray` (renommé en `tray-icon` côté Rust) pour offrir une icône dans la barre des tâches macOS, le system tray Windows ou la status bar GNOME. Cela transforme votre application en quasi-démon qui réagit aux clics depuis l’icône.

Quatrièmement, pour les workflows DevOps européens, GitHub Actions propose des runners Linux, macOS et Windows que vous pouvez orchestrer avec l’action officielle `tauri-apps/tauri-action@v0`. Une seule définition YAML produit les binaires des trois plateformes et publie les artefacts signés sur GitHub Releases – idéal pour un cycle de release hebdomadaire automatisé.

## Cas d’usage réels : Tauri 2 en production

Plusieurs applications grand public et open source tournent en production sur Tauri 2 depuis fin 2024. **Ente Photos**, alternative chiffrée bout-en-bout à Google Photos, a migré son client desktop sur Tauri 2 en 2025 et expédie des binaires de moins de 12 Mo pour Windows, macOS et Linux. Le code est partagé avec leurs versions iOS et Android grâce au support mobile.

**AppFlowy**, alternative open source à Notion, a finalisé sa migration de Flutter Desktop vers Tauri 2 au premier trimestre 2025. L’équipe rapporte une réduction de 60 % de la consommation mémoire et un démarrage trois fois plus rapide qu’avec Electron sur leur ancienne version. Le binaire AppFlowy 2.x pèse environ 18 Mo contre 220 Mo pour la version Electron initiale.

**Relay**, client Git desktop pour GitHub, est passé sur Tauri 2 dès novembre 2024. Selon leurs benchmarks publiés en mars 2026, l’application démarre en 80 ms sur un MacBook Air M3 et consomme 38 Mo de RAM au repos contre 410 Mo pour GitHub Desktop (Electron). Pour une équipe française qui veut publier un client lourd à des milliers d’utilisateurs sans alourdir leur infrastructure de téléchargement, l’argument est imparable.

## Comparaison écosystèmes : Tauri 2, Electron, Flutter et .NET MAUI

| Framework | Langage backend | Frontend | Taille typique | Cibles | Licence | 
|---|---|---|---|---|---|
| Tauri 2.11 | Rust 1.81+ | HTML/CSS/JS (n’importe quel framework) | 3-8 Mo | Win, macOS, Linux, Android, iOS | MIT/Apache 2.0 | 
| Electron 33 | Node.js 22 | HTML/CSS/JS (Chromium embarqué) | 120-200 Mo | Win, macOS, Linux | MIT | 
| Flutter 3.22 | Dart 3.4 | Widgets Flutter (Skia/Impeller) | 30-60 Mo | Win, macOS, Linux, Android, iOS, Web | BSD-3 | 
| .NET MAUI 9 | C# / .NET 9 | XAML | 40-80 Mo | Win, macOS, Android, iOS | MIT | 
| Wails 3 | Go 1.23 | HTML/CSS/JS | 5-12 Mo | Win, macOS, Linux | MIT | 

Le tableau résume les forces et compromis de chaque approche. Tauri 2 reste le plus économe en taille et en mémoire, et le seul à offrir Rust comme langage backend, ce qui séduit les équipes qui veulent un seul langage pour leur backend cloud (axum, actix-web) et leur backend desktop. Flutter conserve l’avantage des widgets natifs identiques sur toutes plateformes, mais paye une taille de binaire deux à dix fois supérieure.

## Performance et benchmarks officiels

Les benchmarks comparatifs publiés sur le blog officiel Tauri en 2025 et confirmés par la communauté Dev.to en mars 2026 montrent un écart constant. Sur un Lenovo ThinkPad X1 Carbon Gen 11 (Intel Core i7-1365U, 32 Go RAM), une application Hello World Tauri 2 consomme **28 Mo de RAM** au repos contre **184 Mo pour Electron** et **67 Mo pour Flutter Desktop**. Le temps jusqu’à interactivité (TTI) est de **92 ms** pour Tauri, **1,8 s** pour Electron et **340 ms** pour Flutter.

Côté CPU, l’écart est moins spectaculaire car la majorité du temps de rendu se fait dans la WebView système, identique en performance brute à celle d’Electron. Là où Tauri brille, c’est sur l’empreinte mémoire de processus secondaires : Electron lance trois à cinq processus Chromium par fenêtre, contre un seul pour Tauri. Sur une machine modeste avec 8 Go de RAM, cela permet d’avoir trois applications Tauri ouvertes simultanément là où une seule application Electron consommerait déjà la moitié de la mémoire disponible.

## Distribution dans l’Union européenne et conformité

Pour les éditeurs français et européens, Tauri 2 simplifie plusieurs aspects réglementaires. D’abord, l’absence de Chromium embarqué élimine la collecte de données de télémétrie Google par défaut, un point sensible pour les organismes publics et les administrations. La CNIL et l’ANSSI publient depuis 2025 des guides recommandant explicitement les frameworks à WebView système pour les déploiements en environnement souverain.

Ensuite, la chaîne de signature et de notarization peut être réalisée intégralement via des autorités de certification européennes telles que **Certigna**, **Certinomis** ou **SwissSign** pour les binaires Windows. La Commission européenne, dans le cadre du règlement Cyber Resilience Act qui entre en vigueur en décembre 2027, impose aux éditeurs de logiciels desktop des exigences de sécurité par défaut : système de mises à jour signées, suivi des CVE, et notification d’incident sous 24 heures. Le système d’updater Tauri répond nativement aux deux premières exigences.

Enfin, la **Mission French Tech** a publié en mai 2026 un appel à projets pour soutenir les éditeurs français de logiciels bureau utilisant des stacks open source à empreinte légère. Tauri y est cité parmi les frameworks éligibles aux côtés de Flutter et Wails. Les candidatures sont ouvertes jusqu’à fin juin 2026.

## Mise en place CI/CD multi-plateforme avec GitHub Actions

Pour publier des binaires multi-plateformes à chaque tag git, l’action officielle `tauri-apps/tauri-action` orchestre la compilation sur les trois OS hébergés par GitHub. Voici un workflow minimal :

```
# .github/workflows/release.yml
name: Release Tauri
on:
  push:
    tags: ['v*']
jobs:
  build:
    strategy:
      fail-fast: false
      matrix:
        platform: [macos-latest, ubuntu-22.04, windows-latest]
    runs-on: ${{ matrix.platform }}
    steps:
      - uses: actions/checkout@v4
      - uses: dtolnay/rust-toolchain@stable
      - uses: actions/setup-node@v4
        with:
          node-version: 20
      - name: Install Linux deps
        if: matrix.platform == 'ubuntu-22.04'
        run: |
          sudo apt-get update
          sudo apt-get install -y libwebkit2gtk-4.1-dev libxdo-dev \
            libssl-dev libayatana-appindicator3-dev librsvg2-dev
      - run: npm ci
      - uses: tauri-apps/tauri-action@v0
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          TAURI_SIGNING_PRIVATE_KEY: ${{ secrets.TAURI_KEY }}
          TAURI_SIGNING_PRIVATE_KEY_PASSWORD: ${{ secrets.TAURI_KEY_PWD }}
        with:
          tagName: ${{ github.ref_name }}
          releaseName: 'NoteForge ${{ github.ref_name }}'
          releaseDraft: true
          prerelease: false
```
Avec ce workflow, un simple `git tag v0.2.0 && git push --tags` déclenche trois builds parallèles. Vingt-cinq minutes plus tard environ (essentiellement compression des binaires release Rust), votre release GitHub contient un MSI, deux DMG (Intel et ARM), un DEB et un AppImage signés.

## FAQ : questions fréquentes sur Tauri 2

### Faut-il connaître Rust pour développer avec Tauri 2 ?

Pas nécessairement pour démarrer. La majorité du code applicatif vit dans le frontend (React, Vue, Svelte, Angular ou même HTML pur). Vous touchez à Rust uniquement pour écrire des commandes `#[tauri::command]`, généralement quelques dizaines de lignes par projet. Pour un usage avancé (plugins personnalisés, intégrations système profondes), une connaissance de base de Rust devient nécessaire – environ deux à quatre semaines d’apprentissage pour un développeur expérimenté.

### Tauri 2 supporte-t-il les frameworks Vue, Svelte ou Angular ?

