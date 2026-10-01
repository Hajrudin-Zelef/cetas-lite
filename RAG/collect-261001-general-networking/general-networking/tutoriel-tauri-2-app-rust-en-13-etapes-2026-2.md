---
id: collect-261001-general-networking/general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026-2
title: "Sous Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["copyright"]
source: docs/RAG/collect-261001-general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026.md
source_anchor: ""
source_lines: [78, 239]
sha256: 41f8dac9e5d131f8e326670ad51aef151e58264c29afa687db7da5f37c0ac33e
---

# Sous Linux/macOS

```
noteforge/
├── src/                       # Frontend React/TypeScript
│   ├── App.tsx
│   ├── main.tsx
│   └── styles.css
├── src-tauri/                 # Backend Rust
│   ├── Cargo.toml             # Dépendances Rust
│   ├── tauri.conf.json        # Configuration centrale
│   ├── build.rs               # Script de build
│   ├── capabilities/          # Permissions par fenêtre
│   │   └── default.json
│   ├── icons/                 # Icônes multi-plateformes
│   └── src/
│       ├── main.rs            # Point d'entrée
│       └── lib.rs             # Logique applicative
├── package.json
├── vite.config.ts
└── index.html
```
Le fichier le plus important est `src-tauri/tauri.conf.json`. Il définit l’identifiant applicatif, les fenêtres à créer, la politique de sécurité du contenu (CSP), les cibles de bundle et la configuration des plugins. Le second fichier critique est `src-tauri/capabilities/default.json`, nouveau dans Tauri 2, qui remplace l’ancien système d’allowlist par un mécanisme de capabilities granulaire. Chaque permission y est associée à une fenêtre spécifique et peut filtrer les arguments transmis aux commandes.

## Étape 4 : Configuration de tauri.conf.json pour la production

Remplacez le contenu de `src-tauri/tauri.conf.json` par cette configuration adaptée à un usage production. Notez la définition des cibles de bundle (MSI Windows, DMG macOS, DEB et AppImage Linux), la fenêtre principale avec ses dimensions raisonnables, et la CSP stricte qui interdit le chargement de scripts distants.

```
{
  "$schema": "https://schema.tauri.app/config/2",
  "productName": "NoteForge",
  "version": "0.1.0",
  "identifier": "com.noteforge.app",
  "build": {
    "beforeDevCommand": "npm run dev",
    "devUrl": "http://localhost:1420",
    "beforeBuildCommand": "npm run build",
    "frontendDist": "../dist"
  },
  "app": {
    "windows": [
      {
        "title": "NoteForge",
        "width": 1100,
        "height": 720,
        "minWidth": 800,
        "minHeight": 600,
        "resizable": true,
        "fullscreen": false,
        "decorations": true,
        "transparent": false
      }
    ],
    "security": {
      "csp": "default-src 'self' ipc: http://ipc.localhost; img-src 'self' data: asset: http://asset.localhost; style-src 'self' 'unsafe-inline'"
    }
  },
  "bundle": {
    "active": true,
    "targets": ["msi", "deb", "appimage", "dmg"],
    "icon": [
      "icons/32x32.png",
      "icons/128x128.png",
      "icons/icon.icns",
      "icons/icon.ico"
    ],
    "category": "Productivity",
    "shortDescription": "Gestionnaire de notes chiffrées local-first",
    "copyright": "2026 NoteForge"
  }
}
```
La directive CSP exclut les scripts inline et les sources tierces, ce qui bloque la quasi-totalité des injections XSS. Si votre frontend utilise Tailwind CSS injecté en runtime ou un loader Vite spécifique, ajustez `style-src` mais évitez d’autoriser `'unsafe-inline'` sur `script-src`.

## Étape 5 : Ajout des plugins Tauri 2 (SQL, FS, Notification, Store)

Les plugins Tauri 2 sont des caisses Rust publiées sur crates.io et accompagnées d’un paquet npm pour le binding TypeScript. Pour notre gestionnaire de notes, nous installons quatre plugins officiels stabilisés depuis la version 2.4.0 du 20 mars 2025, consolidée par les correctifs 2.4.1 du 1er avril 2025 puis 2.5.0 du 15 avril 2025 (Tauri Releases) : `tauri-plugin-sql` pour SQLite, `tauri-plugin-fs` pour la lecture/écriture de fichiers, `tauri-plugin-notification` pour les notifications système, et `tauri-plugin-store` pour le stockage clé-valeur persistant – ce dernier a atteint sa version 2.0.0 dès la sortie stable du 2 octobre 2024 et sa page de release a encore été mise à jour en juillet 2026. Tout l’écosystème suit un versionnage cohérent : le crate `tauri-plugin-http` affiche lui aussi une v2.0.0 stable depuis octobre 2024, confirmée par le changelog 2026, tandis que `tauri-plugin-global-shortcut` partage cette même bascule du 2 octobre 2024 vers la branche v2. L’écosystème continue d’évoluer au même rythme soutenu : le plugin de stockage sécurisé `tauri-plugin-stronghold` est passé en v2.3.0 le 25 juin 2025 puis en v2.3.1 le 27 octobre 2025, tandis que `tauri-plugin-upload`, dédié aux téléversements de fichiers volumineux, a suivi une cadence similaire avec sa v2.3.1 du 20 juillet 2025 puis sa v2.3.2 du 27 octobre 2025, d’après les pages de releases officielles Tauri.

```
# Côté Rust : ajouter dans src-tauri/Cargo.toml
cd src-tauri
cargo add tauri-plugin-sql --features sqlite
cargo add tauri-plugin-fs
cargo add tauri-plugin-notification
cargo add tauri-plugin-store
cd ..
# Côté JavaScript : binding officiel
npm install @tauri-apps/plugin-sql @tauri-apps/plugin-fs \
  @tauri-apps/plugin-notification @tauri-apps/plugin-store
```
Enregistrez ensuite ces plugins dans `src-tauri/src/lib.rs`. C’est le point central qui démarre le runtime Tauri et les enchaîne dans le builder :

```
// src-tauri/src/lib.rs
use tauri_plugin_sql::{Migration, MigrationKind};
#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    let migrations = vec![Migration {
        version: 1,
        description: "create_notes_table",
        sql: "CREATE TABLE IF NOT EXISTS notes (
                id INTEGER PRIMARY KEY AUTOINCREMENT,
                title TEXT NOT NULL,
                body  TEXT NOT NULL,
                created_at INTEGER NOT NULL,
                updated_at INTEGER NOT NULL
              );",
        kind: MigrationKind::Up,
    }];
    tauri::Builder::default()
        .plugin(tauri_plugin_fs::init())
        .plugin(tauri_plugin_notification::init())
        .plugin(tauri_plugin_store::Builder::default().build())
        .plugin(
            tauri_plugin_sql::Builder::default()
                .add_migrations("sqlite:noteforge.db", migrations)
                .build(),
        )
        .invoke_handler(tauri::generate_handler![
            crate::commands::save_note,
            crate::commands::list_notes,
            crate::commands::delete_note
        ])
        .run(tauri::generate_context!())
        .expect("erreur au démarrage de Tauri");
}
mod commands;
```
## Étape 6 : Écrire ses premières commandes Rust invoquées depuis le frontend

Le pont entre votre frontend et Rust se nomme **invoke**. Côté Rust, vous décorez une fonction avec `#[tauri::command]`. Côté JS, vous appelez `invoke('nom_commande', { args })`. Tauri sérialise automatiquement les paramètres en JSON via `serde`. Créez le fichier `src-tauri/src/commands.rs` avec trois commandes asynchrones pour notre gestionnaire de notes :

```
// src-tauri/src/commands.rs
use serde::{Deserialize, Serialize};
use tauri::AppHandle;
#[derive(Serialize, Deserialize, Debug)]
pub struct Note {
    pub id: Option<i64>,
    pub title: String,
    pub body: String,
    pub created_at: i64,
    pub updated_at: i64,
}
#[tauri::command]
pub async fn save_note(_app: AppHandle, note: Note) -> Result<i64, String> {
    // Logique d'insertion via tauri-plugin-sql en frontend
    // ou via sqlx ici si on choisit le backend pur Rust.
    println!("Sauvegarde demandée pour : {}", note.title);
    Ok(note.id.unwrap_or(0))
}
#[tauri::command]
pub async fn list_notes() -> Result<Vec<Note>, String> {
    Ok(vec![])
}
#[tauri::command]
pub async fn delete_note(id: i64) -> Result<(), String> {
    println!("Suppression note id = {}", id);
    Ok(())
}
```
Côté frontend, l’appel est limpide :

