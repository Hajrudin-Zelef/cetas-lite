---
id: collect-261001-general-networking/general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026-3
title: "Sous Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-tauri-2-app-rust-en-13-etapes-2026.md
source_anchor: ""
source_lines: [240, 420]
sha256: e42c29f0581bae1a51d5f9e2c7c895e260211a07655ee2b27bccab01c3bfc4f1
---

# Sous Linux/macOS

```
// src/api.ts
import { invoke } from "@tauri-apps/api/core";
export interface Note {
  id?: number;
  title: string;
  body: string;
  created_at: number;
  updated_at: number;
}
export async function saveNote(note: Note): Promise<number> {
  return await invoke<number>("save_note", { note });
}
export async function listNotes(): Promise<Note[]> {
  return await invoke<Note[]>("list_notes");
}
export async function deleteNote(id: number): Promise<void> {
  return await invoke<void>("delete_note", { id });
}
```
## Étape 7 : Persistance avec SQLite via tauri-plugin-sql

Plutôt que de réinventer la couche de persistance en Rust, le plugin `tauri-plugin-sql` expose directement une API `Database` côté JavaScript. La base SQLite est créée dans le répertoire applicatif du système (par exemple `~/.local/share/com.noteforge.app/` sous Linux), avec gestion automatique des migrations versionnées que nous avons déclarées à l’étape 5.

```
// src/db.ts
import Database from "@tauri-apps/plugin-sql";
let dbInstance: Database | null = null;
export async function getDb(): Promise<Database> {
  if (!dbInstance) {
    dbInstance = await Database.load("sqlite:noteforge.db");
  }
  return dbInstance;
}
export async function insertNote(title: string, body: string): Promise<number> {
  const db = await getDb();
  const now = Date.now();
  const result = await db.execute(
    "INSERT INTO notes (title, body, created_at, updated_at) VALUES ($1, $2, $3, $4)",
    [title, body, now, now]
  );
  return result.lastInsertId ?? 0;
}
export async function fetchNotes() {
  const db = await getDb();
  return await db.select<
    { id: number; title: string; body: string; created_at: number; updated_at: number }[]
  >("SELECT * FROM notes ORDER BY updated_at DESC");
}
export async function removeNote(id: number) {
  const db = await getDb();
  await db.execute("DELETE FROM notes WHERE id = $1", [id]);
}
```
N’oubliez pas d’autoriser le plugin SQL dans le fichier de capabilities. Sans cela, l’appel échouera silencieusement avec une erreur `permissions denied`.

## Étape 8 : Système de capabilities et sécurisation des permissions

Le système de capabilities est l’innovation sécurité majeure de Tauri 2. Il remplace l’ancienne `tauri.allowlist` (Tauri 1.x) par des fichiers JSON déclaratifs scoped par fenêtre, permettant un contrôle d’accès granulaire à l’IPC. Modifiez `src-tauri/capabilities/default.json` :

```
{
  "$schema": "../gen/schemas/desktop-schema.json",
  "identifier": "default",
  "description": "Permissions par défaut pour la fenêtre principale",
  "windows": ["main"],
  "permissions": [
    "core:default",
    "fs:default",
    "fs:allow-app-read",
    "fs:allow-app-write",
    "notification:default",
    "notification:allow-notify",
    "store:default",
    {
      "identifier": "sql:allow-execute",
      "allow": [{ "name": "sqlite:noteforge.db" }]
    },
    {
      "identifier": "sql:allow-select",
      "allow": [{ "name": "sqlite:noteforge.db" }]
    },
    {
      "identifier": "sql:allow-load",
      "allow": [{ "name": "sqlite:noteforge.db" }]
    }
  ]
}
```
Notez la granularité : nous n’autorisons l’accès SQL que pour la base nommée `sqlite:noteforge.db`, pas n’importe quelle base. De même, l’accès au système de fichiers est limité au scope `$APPDATA` via `fs:allow-app-read/write`. Si plus tard vous ajoutez une fenêtre `settings`, vous créerez un fichier `capabilities/settings.json` séparé avec ses propres permissions, beaucoup plus restrictives.

## Étape 9 : Construction de l’interface React et liaison aux commandes

Place au frontend. Remplacez le contenu de `src/App.tsx` par une interface minimaliste de gestion de notes. Le composant utilise les hooks React standards et appelle nos fonctions `insertNote`, `fetchNotes`, `removeNote` définies à l’étape 7.

```
// src/App.tsx
import { useEffect, useState } from "react";
import { insertNote, fetchNotes, removeNote } from "./db";
import { sendNotification, isPermissionGranted, requestPermission }
  from "@tauri-apps/plugin-notification";
import "./styles.css";
interface NoteRow {
  id: number;
  title: string;
  body: string;
  created_at: number;
  updated_at: number;
}
export default function App() {
  const [notes, setNotes] = useState<NoteRow[]>([]);
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const reload = async () => setNotes(await fetchNotes());
  useEffect(() => { reload(); }, []);
  const onSave = async () => {
    if (!title.trim()) return;
    await insertNote(title, body);
    setTitle(""); setBody("");
    await reload();
    let granted = await isPermissionGranted();
    if (!granted) granted = (await requestPermission()) === "granted";
    if (granted) sendNotification({ title: "NoteForge", body: "Note sauvegardée" });
  };
  return (
    <main className="container">
      <h1>NoteForge</h1>
      <section className="editor">
        <input value={title} onChange={e => setTitle(e.target.value)} placeholder="Titre" />
        <textarea value={body} onChange={e => setBody(e.target.value)} placeholder="Contenu" />
        <button onClick={onSave}>Sauvegarder</button>
      </section>
      <ul className="notes">
        {notes.map(n => (
          <li key={n.id}>
            <strong>{n.title}</strong>
            <p>{n.body}</p>
            <button onClick={async () => { await removeNote(n.id); reload(); }}>
              Supprimer
            </button>
          </li>
        ))}
      </ul>
    </main>
  );
}
```
Relancez `npm run tauri dev`. Vous devriez voir l’interface, créer une note, recevoir une notification système native, et la note apparaître dans la liste – persistante entre les redémarrages grâce à SQLite.

## Étape 10 : Communication événementielle Rust ↔ Frontend

Au-delà des commandes synchrones, Tauri propose un bus d’événements asynchrone bidirectionnel. C’est utile pour notifier le frontend quand un travail long en Rust progresse, ou pour pousser des messages depuis le backend sans que le frontend ait demandé quoi que ce soit. Côté Rust, vous émettez avec `app.emit` ; côté JS, vous écoutez avec `listen`.

```
// src-tauri/src/commands.rs – ajout
use tauri::{AppHandle, Emitter};
#[tauri::command]
pub async fn long_task(app: AppHandle) -> Result<(), String> {
    for i in 0..=100 {
        std::thread::sleep(std::time::Duration::from_millis(40));
        app.emit("task-progress", i).map_err(|e| e.to_string())?;
    }
    Ok(())
}
```
```
// src/progress.ts
import { listen } from "@tauri-apps/api/event";
import { invoke } from "@tauri-apps/api/core";
export async function startWithProgress(onTick: (n: number) => void) {
  const unlisten = await listen<number>("task-progress", e => onTick(e.payload));
  await invoke("long_task");
  unlisten();
}
```
Pensez à enregistrer la commande `long_task` dans `generate_handler!` dans `lib.rs`. L’oubli de cet enregistrement est la première erreur que rencontrent les nouveaux développeurs Tauri.

## Étape 11 : Build de production et signature des artefacts

Une fois l’application terminée, lancez la build de production. La commande `tauri build` compile en mode release (LTO activé, optimisations agressives), exécute le bundler approprié à votre plateforme et signe l’exécutable si vous avez configuré une clé.

