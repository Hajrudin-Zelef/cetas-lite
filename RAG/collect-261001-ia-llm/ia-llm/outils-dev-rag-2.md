---
id: collect-261001-ia-llm/ia-llm/outils-dev-rag-2
title: "Outils dev + ingénierie RAG (chunk & corpus)"
domain: ia-llm
role: reference
task: reference
actors: ["Huawei", "Microsoft", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["agent", "agents", "chatgpt", "copilot", "embedding", "embeddings", "mcp", "model context protocol"]
source: docs/RAG/collect-261001-ia-llm/outils_dev_rag.md
source_anchor: ""
source_lines: [168, 377]
sha256: ebd416257718f112da413844d317eb078b65900199ed9062aeb7ddebd26952fb
---

# Outils dev + ingénierie RAG (chunk & corpus)

```jsonc
{
  "RAG : embed batch OpenAI": {
    "prefix": "embbatch",
    "body": [
      "from openai import OpenAI",
      "client = OpenAI()",
      "",
      "def embed_batch(texts: list[str], model: str = \"text-embedding-3-small\") -> list[list[float]]:",
      "    resp = client.embeddings.create(model=model, input=texts)",
      "    return [d.embedding for d in resp.data]",
      "${0}"
    ],
    "description": "Embedding OpenAI en batch"
  },
  "RAG : main argparse": {
    "prefix": "pymain",
    "body": [
      "import argparse",
      "from pathlib import Path",
      "",
      "def main() -> None:",
      "    ap = argparse.ArgumentParser()",
      "    ap.add_argument(\"--input\", type=Path, required=True)",
      "    ap.add_argument(\"--output\", type=Path, required=True)",
      "    args = ap.parse_args()",
      "    ${0}",
      "",
      "if __name__ == \"__main__\":",
      "    main()"
    ],
    "description": "Squelette script CLI"
  }
}
```

Tape `embbatch` + `Tab` : le bloc se génère. Crée des snippets pour tes
patterns récurrents (connexion pgvector, retry, lecture manifest).

## 6. Remote-SSH : développer sur ton serveur distant (ton cas d'usage)

C'est **ton** usage : le code tourne sur le serveur (où sont les corpus, les
venv, les cron), tu édites depuis ton poste.

```bash
# 1. Côté poste : installe l'extension (voir §3)
code --install-extension ms-vscode-remote.remote-ssh

# 2. Configure ton SSH (~/.ssh/config) — vérifié, doc officielle
```

```ssh-config
Host rag-server
    HostName 192.0.2.10        # ou ton FQDN
    User zelef
    IdentityFile ~/.ssh/id_ed25519
    ForwardAgent yes            # pour que git/gh marchent côté serveur
```

```text
3. Ctrl+Shift+P → "Remote-SSH: Connect to Host…" → rag-server
4. Première connexion : VS Code installe VS Code Server dans ~/ sur le serveur
   (1 Go RAM mini requis, 2 Go + 2 CPU recommandés — vérifié).
5. En bas à gauche : "SSH: rag-server" = tu es sur le serveur.
```

**Points clés :**
- Les extensions s'installent **côté serveur** (onglet Extensions → « Installer
  sur SSH: rag-server »). Installe **Python, Pylance, Copilot** côté serveur ;
  le thème, lui, reste local.
- Le terminal intégré (`Ctrl+`` `) est un shell **du serveur** : tes
  `collect_*.py`, `tmux`, `psql` y tournent nativement.
- Port forwarding : pour exposer un dashboard local du serveur
  (ex. : une API FastAPI sur le port 8000 du serveur), clic droit dans
  l'onglet **Ports** → *Forward a Port*. Accessible en `localhost:8000`
  **sur ton poste**.
- **Piège classique :** ouvrir un dossier *local* puis lancer un script qui
  attend les chemins du serveur. Vérifie toujours l'indicateur en bas à
  gauche avant de lancer un traitement de corpus.

## 7. MCP dans VS Code : brancher des outils à l'agent

Le **Model Context Protocol** permet à l'agent d'appeler des outils externes
(base de données, API, système de fichiers). Config repo : `.vscode/mcp.json`
(vérifié sept 2026) :

```jsonc
{
  "servers": {
    "postgres-rag": {
      "command": "uvx",                       // à vérifier : paquet exact
      "args": ["mcp-server-postgres@latest"], // à vérifier : paquet exact
      "env": {
        "POSTGRES_URL": "${input:pg-url}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "pg-url",
      "description": "URL Postgres du RAG",
      "password": true
    }
  ]
}
```

> Les noms de paquets MCP évoluent vite : vérifie sur le registre MCP /
> GitHub avant d'utiliser cet exemple tel quel. Le **mécanisme**
> (`.vscode/mcp.json` + `mcp.servers`) est lui confirmé.

## 8. Agents distants & fenêtre Agents (nouveauté 2026, vérifié)

Depuis la **1.139 (23/09/2026, Stable)** :
- **Fenêtre Agents** : surface dédiée pour piloter plusieurs sessions
  d'agents en parallèle, revoir les diffs, comparer.
- **Agents distants (preview)** : une session tourne sur une machine distante
  (SSH ou Dev Tunnel) et **continue même si ton client se déconnecte**.
- **BYOK** : tu peux brancher tes propres clés API de modèles, y compris en
  environnement air-gapped.
- Les sessions de chat se synchronisent sur ton compte GitHub (historique
  interrogeable d'une machine à l'autre).

**Usage concret pour toi :** lance une session agent sur `rag-server` via
Remote-SSH, demande-lui « nettoie les headers/footers des 200 .md de
`corpus/brut/` », ferme ton laptop : la session continue.

## 9. Intégration avec les agents codeurs (Codex, Copilot CLI)

- **Codex (extension `openai.chatgpt`, vérifiée)** : sidebar → mode
  **Agent** → tâche en nommant les fichiers avec `@nom_fichier`.
  Auth distante (SSH) : forwarde le port `1455` (*Forward a Port*) puis
  `codex login` dans le terminal distant ; sinon copie `auth.json` depuis
  une connexion locale vers le dossier de l'extension côté distant.
  **Ne commite jamais `auth.json`.**
- **Copilot CLI** : utile quand tu es en SSH sans IDE
  (`github.copilot.chat.*` pour la config).
- **Bon réflexe** : fais toujours relire le diff (`git diff`) avant
  d'accepter le travail d'un agent sur tes scripts de collecte.

## 10. Checklist VS Code

- [ ] VS Code ≥ 1.139 installé, extensions vérifiées (§3)
- [ ] Profil « RAG » créé et exporté
- [ ] `settings.json` : format on save, trim whitespace, Copilot agent configuré
- [ ] Remote-SSH : `~/.ssh/config` + connexion testée au serveur
- [ ] Extensions **serveur** installées (Python, Copilot) côté distant
- [ ] `.vscode/mcp.json` versionné dans le repo (sans secrets)
- [ ] Snippets RAG créés (`embbatch`, `pymain`, …)
- [ ] Aucune clé API dans les settings synchronisés

---

## 11. GitHub : pourquoi c'est le socle de ton projet RAG

Ton pipeline RAG, c'est du **code** (scripts de collecte, chunking,
ingestion) + du **corpus** (données). GitHub te donne :

| Besoin | Outil GitHub | Pourquoi ça compte pour toi |
|---|---|---|
| Versionner le code | Repos Git | Historique, rollback, branches |
| Gros fichiers de corpus | Git LFS | Au-delà de ~100 Mo, Git seul casse |
| Qualité auto | GitHub Actions | Lint + tests à chaque push |
| Suivi des collectes | Issues (+ templates) | « Collecte Huawei lot 3 : 12/452 restantes » |
| Secrets (clés API) | Actions secrets | Jamais de clé en dur dans le code |
| Dev ponctuel dans le cloud | Codespaces | 60 core-hours/mois gratuits (vérifié sept 2026) |
| Doc vivante | README + wiki | Ton futur toi te remerciera |

> **Règle structurante :** le repo contient le **code** et les **petits**
> artefacts (configs, manifestes, docs). Le **corpus brut volumineux**
> (> quelques centaines de Mo) vit ailleurs (disque serveur, stockage
> objet) ou en Git LFS — voir §16.

## 12. Créer le repo (gh CLI + web)

La CLI officielle `gh` (à installer : `sudo apt install gh` sur
Debian/Ubuntu — paquet standard) :

```bash
# Authentification (une fois)
gh auth login

# Créer le repo privé + initialiser + pousser
mkdir rag-perso && cd rag-perso
git init -b main
gh repo create rag-perso --private --source=. --push
# --public si tu veux le partager ; --private par défaut pour toi

# Vérifier
gh repo view --web
```

Création manuelle (web) : github.com → *New repository* → nom,
visibilité **Private**, coche *Add a README* si tu pars de zéro.
Ensuite :

```bash
git remote add origin git@github.com:<ton-user>/rag-perso.git
git branch -M main
git push -u origin main
```

> `<ton-user>` : remplace par ton identifiant GitHub. Utilise l'URL SSH
> (`git@...`) avec ta clé Ed25519 déjà en place (cf. §6 `ForwardAgent`).

## 13. Structure de repo recommandée pour ton projet RAG

Arborescence cible (adapte les noms à tes scripts existants) :

