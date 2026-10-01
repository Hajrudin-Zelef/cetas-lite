---
id: collect-261001-general-networking/general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026-4
title: "Sous Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["attention", "distribution"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026.md
source_anchor: ""
source_lines: [421, 544]
sha256: 02a1019b42f26da82ba5d0450596c4acc3f0b8c977fae22e4c13583af10361f2
---

# Sous Linux/macOS

```
# Build par défaut pour la plateforme courante
npm run tauri build
# Cibler explicitement une plateforme
npm run tauri build -- --target x86_64-pc-windows-msvc
npm run tauri build -- --target aarch64-apple-darwin
npm run tauri build -- --target x86_64-unknown-linux-gnu
# Sortie attendue (sous Linux)
# Compiling noteforge v0.1.0
# Finished `release` profile [optimized] target(s) in 3m 24s
# Bundling NoteForge_0.1.0_amd64.deb
# Bundling NoteForge_0.1.0_amd64.AppImage
```
Sur macOS, la signature et la notarization Apple nécessitent un certificat Developer ID et le téléchargement de votre `app-specific password` dans le trousseau. Ajoutez à votre environnement :

```
export APPLE_SIGNING_IDENTITY="Developer ID Application: Votre Société (TEAM_ID)"
export APPLE_ID="[email protected]"
export APPLE_PASSWORD="@keychain:AC_PASSWORD"
export APPLE_TEAM_ID="ABCDE12345"
npm run tauri build
```
Sous Windows, fournissez un certificat de signature de code (.pfx) et son mot de passe via les variables `TAURI_SIGNING_PRIVATE_KEY` et `TAURI_SIGNING_PRIVATE_KEY_PASSWORD`. Pour Linux, aucune signature n’est exigée par défaut, mais un AppImage signé GPG renforce la confiance des distributions.

## Étape 12 : Activation du mécanisme de mise à jour automatique (updater)

Le plugin officiel `tauri-plugin-updater`, passé en version **2.10.0 en février 2026** avec l’ajout des options `no_proxy` et de configurations TLS non sécurisées pour les environnements de test, permet à votre application de se mettre à jour silencieusement depuis un endpoint que vous hébergez (S3, R2, GitHub Releases, ou votre propre serveur). L’updater vérifie une URL renvoyant un manifeste JSON signé Ed25519 et télécharge le binaire mis à jour.

```
# Génération de la clé de signature
npm install -D @tauri-apps/cli
npx tauri signer generate -w ~/.tauri/noteforge.key
# Sortie :
# Public key : dW50cnVzdGVkIGNvbW1lbnQ6IG1pbmlzaWduIH...
# Conservez le fichier ~/.tauri/noteforge.key précieusement et ajoutez la clé publique
# dans tauri.conf.json > plugins > updater > pubkey
# Ajout du plugin
cd src-tauri && cargo add tauri-plugin-updater
cd .. && npm install @tauri-apps/plugin-updater
```
Configurez l’updater dans `tauri.conf.json` :

```
"plugins": {
  "updater": {
    "endpoints": [
      "https://updates.noteforge.fr/{{target}}/{{current_version}}"
    ],
    "pubkey": "dW50cnVzdGVkIGNvbW1lbnQ6IG1pbmlzaWduIHB1YmxpYyBrZXk6..."
  }
}
```
Côté JavaScript, déclenchez la vérification au démarrage :

```
import { check } from "@tauri-apps/plugin-updater";
import { relaunch } from "@tauri-apps/plugin-process";
const update = await check();
if (update?.available) {
  console.log(`Mise à jour disponible : ${update.version}`);
  await update.downloadAndInstall();
  await relaunch();
}
```
## Étape 13 : Compilation pour Android et iOS

L’une des promesses majeures de Tauri 2 est la prise en charge native d’Android et iOS depuis la même base de code. Pour activer le support mobile, initialisez les projets natifs et installez les SDK requis.

```
# Android – prérequis : JDK 17, Android Studio, NDK 26+
export ANDROID_HOME="$HOME/Library/Android/sdk"
export NDK_HOME="$ANDROID_HOME/ndk/26.1.10909125"
npm run tauri android init
npm run tauri android dev    # Lance sur émulateur ou appareil USB
npm run tauri android build  # Génère l'APK signé
# iOS – prérequis : Xcode 15+, certificat Apple Developer
npm run tauri ios init
npm run tauri ios dev
npm run tauri ios build      # Génère l'IPA
```
Le ciblage mobile expose la macro `#[cfg_attr(mobile, tauri::mobile_entry_point)]` que nous avons utilisée à l’étape 5. Vos commandes Rust marchent telles quelles sur mobile, à condition de ne pas dépendre de caisses spécifiques bureau (par exemple `tray-icon` ou `arboard` pour le presse-papier desktop). Les plugins SQL, FS, Notification et Store fonctionnent sur les cinq cibles.

## Pièges fréquents et comment les éviter

Voici les huit erreurs les plus courantes signalées par la communauté Tauri en 2025-2026, avec leur résolution :

- **Compilation initiale extrêmement lente** : la première`cargo build` prend 5 à 12 minutes. C’est normal. Utilisez`sccache` pour partager le cache entre projets et activez`cargo install cargo-binstall` pour pré-compiler les outils Rust.
- **Erreur “webkit2gtk-4.1 not found”** sous Linux : votre distribution n’a pas la bonne version. Sur Ubuntu 22.04, installez aussi`libwebkit2gtk-4.0-dev` . Sur Fedora 40+, le paquet s’appelle`webkitgtk6.0-devel` .
- **Permission denied lors de l’invoke** : la commande n’est pas listée dans`capabilities/default.json` . Ajoutez le permission identifier exact retourné dans l’erreur.
- **Polices Tailwind invisibles** : la CSP par défaut bloque le chargement de polices distantes. Utilisez des polices locales ou ajoutez la directive`font-src` appropriée.
- **tauri build échoue avec “missing field bundle”** : votre`tauri.conf.json` est au format Tauri 1.x. Lancez`npx tauri migrate` pour convertir automatiquement vers le schéma 2.x.
- **Le DMG macOS refuse de s’ouvrir** : Gatekeeper bloque les binaires non signés. Signez avec votre Developer ID ou indiquez aux utilisateurs`xattr -dr com.apple.quarantine NoteForge.app` .
- **Erreur “frontendDist not found”** : Vite a généré sa sortie dans`build/` au lieu de`dist/` . Adaptez`build.frontendDist` dans`tauri.conf.json` ou modifiez la sortie Vite.
- **Plantage immédiat sur Windows ARM64** : compilez explicitement avec`--target aarch64-pc-windows-msvc` . Les binaires x64 émulés via Prism plantent souvent sur les Surface Pro X.

## Dépannage avancé : huit problèmes et leurs solutions

| Symptôme | Cause probable | Solution | 
|---|---|---|
| ``error: linker `cc` not found`` | Toolchain C absente | Installer build-essential (Linux), Xcode CLT (macOS), Visual Studio Build Tools (Windows) | 
| Fenêtre Tauri blanche au démarrage | devUrl pointe vers un port libre | Vérifier que Vite tourne sur 1420 ; ajuster `devUrl` dans tauri.conf.json | 
| Notifications jamais reçues | Permission OS refusée | Sous macOS : Réglages > Notifications > NoteForge ; sous Windows : Focus Assist désactivé | 
| `Database is locked` | Connexions concurrentes | Ouvrir une seule instance via `tauri-plugin-single-instance` | 
| Hot reload ne fonctionne pas | Ports différents entre Vite et Tauri | Aligner `devUrl` et le port Vite (1420 par défaut) | 
| Build mobile : `SDK 35 not found` | SDK Android obsolète | Mise à jour via Android Studio > SDK Manager | 
| iOS : `provisioning profile invalid` | Bundle ID non enregistré | Créer un App ID dans Apple Developer Portal | 
| Crash après cargo update | Versions Rust et plugins désynchronisées | Verrouiller via `cargo update -p tauri --precise 2.11.1` | 

## Astuces avancées pour aller plus loin

Une fois votre application bureau opérationnelle, plusieurs optimisations valent le détour. Premièrement, activez la **Profile-Guided Optimization (PGO)** sur la build release pour gagner 10 à 20 % de performance sur les chemins critiques. Ajoutez dans `src-tauri/Cargo.toml` :

```
[profile.release]
codegen-units = 1
lto = true
opt-level = "z"
panic = "abort"
strip = true
```
Avec ces réglages, un binaire Rust nu est typiquement réduit de 15 à 25 % supplémentaires. Pour aller encore plus loin, **UPX** peut compresser l’exécutable Windows à environ 50 % de sa taille initiale, mais attention : certains antivirus signalent les binaires UPX comme suspects.

Deuxièmement, pour les applications gourmandes en E/S disque, préférez **SQLx** directement en Rust (avec `tokio`) plutôt que `tauri-plugin-sql`. Vous gagnez la vérification SQL à la compilation et les requêtes préparées, au prix d’un peu plus de code. Le plugin reste idéal pour le prototypage et les cas simples.

