---
id: collect-261001-general-networking/general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026-1
title: "Sous Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 77]
sha256: d818628db7a93310e85dbd786e79dedff90feeee51dc4e0bc036bef0075bbb78
---

# Sous Linux/macOS

Publié le 10 avril 2026, mis à jour en septembre 2026 – Construire une application bureau multi-plateforme en 2026 ne ressemble plus à ce qu’elle était il y a deux ans. Avec **Tauri 2.11**, dernier maillon d’une lignée qui a vu le noyau franchir la version 2.6.0 le 24 juin 2025, puis 2.7.0 le 20 juillet 2025 et 2.8.1 le 18 août 2025, avant d’enchaîner une série de correctifs à l’automne – 2.9.2 le 29 octobre 2025, 2.9.3 le 13 novembre 2025, 2.9.4 le 30 novembre 2025 et 2.9.5 le 9 décembre 2025, d’après les Tauri Releases – jusqu’à la révision 2.11.1 de `@tauri-apps/api`, atteinte en juillet 2026 selon le blog ESB1995, sorti dans la lignée de la version stable 2.0 d’octobre 2024, vous générez désormais un binaire de 3 à 8 Mo qui démarre en moins de 100 ms et consomme entre 25 et 45 Mo de RAM au repos – là où Electron pèse 120 à 200 Mo et engloutit 150 à 300 Mo de mémoire vive. Ce **tutoriel Tauri 2** de 13 étapes vous guide pas à pas dans la création d’une application bureau complète : un gestionnaire de notes chiffré utilisant Rust, SQLite via le plugin officiel, et un frontend React. Toutes les commandes ont été testées sur Windows 11, macOS Sonoma et Ubuntu 24.04.

Le projet Tauri, gouverné par la **Tauri Foundation** sous l’égide du Commons Conservancy et licencié MIT/Apache 2.0, dépasse les **70 000 étoiles GitHub** et compte plus de **4 200 contributeurs**. Le paquet `@tauri-apps/cli` dépasse les 250 000 téléchargements hebdomadaires sur npm, tandis que `@tauri-apps/api` approche 180 000. Pour les équipes françaises et européennes, l’argument est aussi réglementaire : Tauri ne télécharge pas Chromium à l’exécution, ce qui simplifie la conformité RGPD et évite les inquiétudes liées au tracking embarqué dans les moteurs propriétaires.

## Pourquoi choisir Tauri 2 plutôt qu’Electron en 2026

Tauri 2 réutilise le moteur de rendu web installé sur le système d’exploitation : **WebView2** (basé sur Edge/Chromium) sous Windows, **WKWebView** sous macOS et iOS, **WebKitGTK 6.0+** sous Linux, et le composant Android System WebView. Cette stratégie élimine l’embarquement d’un runtime Node.js et d’un binaire Chromium complet à l’intérieur de chaque application. Le résultat se mesure : un binaire Tauri 2 typique tient entre 3 et 8 Mo, contre 120 à 200 Mo pour une application Electron équivalente – un facteur 40 à la baisse documenté par la documentation officielle et confirmé par le tutoriel comparatif publié sur rustify.rs en mars 2026.

Le second axe est la sécurité. Le backend de Tauri est écrit en Rust, langage à mémoire sûre qui élimine les classes entières de vulnérabilités. Surtout, le système de capacités introduit en version 2.0 – un fichier `capabilities/default.json` qui déclare explicitement quelles commandes peuvent être appelées depuis le frontend, par quelle fenêtre et avec quels arguments – applique le principe du moindre privilège par défaut. Un audit externe complet a été réalisé avant la sortie stable de Tauri 2.0 et toutes les conclusions ont été corrigées avant publication, comme indiqué sur le blog officiel d’octobre 2024.

Troisième argument : la couverture mobile. Depuis Tauri 2.0, la même base de code Rust + Web cible **Windows, macOS, Linux, Android et iOS**. Vous écrivez vos commandes Rust une fois, et la CLI génère un APK signé pour Android, un IPA pour iOS, un MSI pour Windows, un DMG pour macOS et un AppImage ou DEB pour Linux. Pour un éditeur SaaS européen qui veut accompagner son application web d’un client bureau et d’une application mobile sans tripler ses équipes, l’économie est immédiate.

## Tauri 2 vs Electron : tableau comparatif détaillé

| Critère | Tauri 2.11.1 (avril 2026) | Electron 33.x | Écart | 
|---|---|---|---|
| Taille du binaire de base | 3 à 8 Mo | 120 à 200 Mo | ~40× plus léger | 
| RAM au démarrage (repos) | 25 à 45 Mo | 150 à 300 Mo | ~6× plus économe | 
| Temps de démarrage à froid | <100 ms | 2 000 à 4 000 ms | ~30× plus rapide | 
| Langage backend | Rust 1.81+ | Node.js 22.x | Mémoire sûre vs JIT | 
| Moteur de rendu | Système (WebView2/WKWebView/WebKitGTK) | Chromium embarqué | Système vs intégré | 
| Cibles mobiles | Android et iOS natifs | Aucune (Capacitor requis) | Mobile inclus | 
| Système de permissions | Capabilities + ACL fines | contextIsolation manuelle | Sécurité par défaut | 
| Étoiles GitHub | 70 000+ | 116 000+ | Communauté plus jeune | 
| Licence | MIT / Apache 2.0 | MIT | Permissive | 

## Étape 1 : Prérequis et versions à installer

Avant de démarrer, votre poste doit disposer de la chaîne d’outils suivante. Les versions ci-dessous sont celles validées par l’équipe Tauri pour la branche 2.11.x en avril 2026 – ne descendez pas en dessous, sous peine de rencontrer des erreurs de compilation obscures sur les caisses (crates) `tao` et `wry`.

- **Rust 1.81.0 ou supérieur** via`rustup` sur le canal stable (vérifiez avec`rustc --version` ).
- **Node.js 20.18.0 LTS ou supérieur** , npm 10.8+ ou pnpm 9+.
- **Cargo** installé automatiquement par rustup.
- **WebView2 Runtime** sur Windows 10 et 11 (préinstallé depuis Windows 11 22H2).
- **WebKitGTK 6.0+** sous Linux :`apt install libwebkit2gtk-4.1-dev build-essential curl wget file libxdo-dev libssl-dev libayatana-appindicator3-dev librsvg2-dev` sur Ubuntu 24.04.
- **Xcode Command Line Tools** sur macOS Sonoma ou supérieur (`xcode-select --install` ).
- Pour le ciblage mobile : Android Studio Hedgehog 2023.1+, JDK 17, NDK 26+, Xcode 15+ pour iOS.

Installez Rust si ce n’est pas déjà fait :

```
# Sous Linux/macOS
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
source "$HOME/.cargo/env"
rustup update stable
# Vérification
rustc --version
# Sortie attendue : rustc 1.81.0 (eeb90cda1 2026-XX-XX)
cargo --version
node --version
# Sortie attendue : v20.18.0 ou plus récent
```
## Étape 2 : Création du squelette de projet avec create-tauri-app

L’outil `create-tauri-app` (CTA) est l’équivalent Tauri de `create-react-app` ou `npm create vite`. Il génère un projet bi-compartiments – un répertoire `src/` pour le frontend web et un répertoire `src-tauri/` pour le backend Rust – avec toute la configuration câblée. Nous allons créer une application appelée `noteforge`, un gestionnaire de notes chiffrées local-first.

```
# Création interactive
npm create tauri-app@latest
# Réponses à donner :
# ? Project name › noteforge
# ? Identifier › com.noteforge.app
# ? Choose which language to use for your frontend › TypeScript / JavaScript
# ? Choose your package manager › npm
# ? Choose your UI template › React
# ? Choose your UI flavor › TypeScript
# ? Would you like to setup the project for mobile › No (pour l'instant)
cd noteforge
npm install
npm run tauri dev
```
La première compilation peut durer entre 5 et 12 minutes selon votre machine, car cargo doit télécharger et compiler l’intégralité de la chaîne `tao`, `wry`, `webkit2gtk-sys` ou `webview2-com`, ainsi que les nombreuses caisses tierces. Les compilations suivantes sont incrémentales et descendent à 2-5 secondes en mode développement grâce au cache de cargo. Une fenêtre Tauri devrait s’ouvrir affichant le template React par défaut.

## Étape 3 : Anatomie du projet Tauri 2

Comprendre la structure générée est essentiel avant de coder. Voici les fichiers clés et leur rôle :

