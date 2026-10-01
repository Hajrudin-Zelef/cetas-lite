---
id: collect-261001-automatisation-infra/automatisation-infra/mcp-lsp-protocoles-2
title: "MCP, LSP, code-server, websearch — Guide pratique"
domain: automatisation-infra
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-07-28"]
keywords: ["mcp", "agent", "claude"]
source: docs/RAG/collect-261001-automatisation-infra/mcp_lsp_protocoles.md
source_anchor: ""
source_lines: [136, 318]
sha256: 3e0a86f6033bdc37fdd9e0573af6856e52aed6fc85260b15d366e4ef395b11f4
---

# MCP, LSP, code-server, websearch — Guide pratique

Requête du client :
```json
{"jsonrpc":"2.0","id":7,"method":"tools/call",
 "params":{"name":"read_file","arguments":{"path":"/srv/docs/ups.md"}}}
```

Réponse du serveur :
```json
{"jsonrpc":"2.0","id":7,"result":{
  "content":[{"type":"text","text":"# Onduleurs...\n..."}],
  "isError":false}}
```

💡 Le champ `isError: true` signale une erreur **métier** (fichier introuvable)
tout en restant une réponse JSON-RPC valide. Les erreurs **protocole**
(méthode inconnue, JSON invalide) utilisent le champ `error` standard
JSON-RPC avec ses codes (-32700, -32601, ...).

⚠️ Règle d'or du transport stdio : **jamais** de `print()` de debug sur
`stdout` côté serveur — ça corrompt le protocole. Les logs vont sur `stderr`.

## 7. Les trois primitives : vue d'ensemble

| Primitive | Question qu'elle répond | Contrôlée par | Exemple |
|---|---|---|---|
| **Tools** | « Que puis-je **faire** ? » | Le modèle décide d'appeler | `query(sql)`, `read_file(path)`, `browser_snapshot()` |
| **Resources** | « Quelles **données** puis-je lire ? » | L'utilisateur/le host choisit d'attacher | `doc://onduleurs/maintenance`, `pg://schema/tables` |
| **Prompts** | « Quelle **tâche** pré-packagée lancer ? » | L'utilisateur choisit | `diagnostic_ups`, `revue_code` |

Différence fondamentale : un **tool** est une action que le modèle déclenche
lui-même pendant sa réflexion. Une **resource** est un contenu que l'humain
(ou le host) glisse dans le contexte — comme une pièce jointe. Un **prompt**
est un template d'instructions avec des paramètres, que l'utilisateur
sélectionne explicitement (« /diagnostic_ups »).

## 8. Tools : anatomie complète

Un tool = un nom + une description + un schéma d'entrée JSON Schema + un
handler. Ce que voit le modèle dans `tools/list` :

```json
{
  "name": "search_docs",
  "title": "Rechercher dans la documentation",
  "description": "Recherche plein-texte dans le corpus documentaire local (guides, manuels). Retourne les extraits les plus pertinents avec leur source.",
  "inputSchema": {
    "type": "object",
    "properties": {
      "query": {"type": "string", "description": "Mots-clés de recherche"},
      "limit": {"type": "integer", "description": "Nombre max de résultats", "default": 5}
    },
    "required": ["query"]
  },
  "annotations": {
    "readOnlyHint": true,
    "destructiveHint": false,
    "idempotentHint": true,
    "openWorldHint": false
  }
}
```

🔒 Les **annotations** sont des indices de sécurité que le host peut exploiter :
- `readOnlyHint: true` → ne modifie rien (pas de confirmation nécessaire) ;
- `destructiveHint: true` → peut détruire (suppression, envoi) → le host
  **doit** demander confirmation ;
- `idempotentHint: true` → rejouer l'appel est sans danger ;
- `openWorldHint: false` → l'outil n'interagit pas avec le monde extérieur
  (pas de réseau) → périmètre fermé, plus sûr.

💡 La **description** est le vrai « contrat » avec le modèle : c'est elle qui
détermine quand il appellera ton tool. Sois précis, donne un exemple
d'usage, dis ce que ça ne fait PAS. Une mauvaise description = un tool
jamais appelé ou appelé à contresens.

## 9. Tools : ce que renvoie un appel

Le résultat d'un `tools/call` contient toujours un tableau `content`
(blocs typés) et optionnellement `structuredContent` (données validées) :

```json
{
  "content": [
    {"type": "text", "text": "3 résultats pour « onduleur bypass » :\n1. ..."}
  ],
  "structuredContent": {
    "results": [
      {"file": "onduleurs_ups_guide.md", "score": 0.94, "extrait": "..."}
    ]
  },
  "isError": false
}
```

Types de blocs possibles : `text`, `image` (base64), `audio`, `resource`
(embarque une resource lue), `resource_link` (pointe vers une URI sans
l'embarquer — économise des tokens 💡).

Règle : `content` (lisible) est obligatoire car c'est ce que le modèle
consomme ; `structuredContent` est un bonus pour les hosts qui veulent
parser proprement.

## 10. Resources : principe, URIs, templates

Une resource expose des **données adressables par URI**, avec un type MIME :

```json
{
  "uri": "doc://guides/onduleurs_ups_guide.md",
  "name": "Guide onduleurs",
  "title": "Guide onduleurs / UPS",
  "description": "Guide terrain 1600 lignes : dimensionnement, maintenance",
  "mimeType": "text/markdown"
}
```

Les URIs sont libres (`doc://`, `pg://`, `file:///`, `https://`...) mais
doivent être stables et non ambiguës. Deux formes :
- **Resources directes** : URI fixe, listées par `resources/list`.
- **Templates** : `doc://guides/{nom}` — le client remplit `{nom}`.
  Idéal pour « N documents du même type » sans les lister un par un.

💡 Cas Zeef typique : ton corpus RAG (les 40+ guides `.md`) exposé en
`doc://guides/{nom}` → n'importe quel agent MCP peut lire ta doc sans
copier de fichiers.

## 11. Resources : lecture et abonnements

Lecture via `resources/read` → renvoie `contents[]` (texte ou blob base64).
En spec 2026-07-28, l'ancien couple `resources/subscribe` / SSE est remplacé
par un flux unique **`subscriptions/listen`** : le client ouvre un stream
et reçoit les notifications de changement (`notifications/resources/updated`).

En pratique : la plupart des serveurs n'implémentent pas les abonnements.
C'est une capacité optionnelle — ne la promets pas si ton backend ne peut
pas détecter les changements (inotify, trigger BDD...).

## 12. Prompts : les workflows pré-packagés

Un prompt = un template nommé avec des arguments. `prompts/list` :

```json
{"name": "diagnostic_ups",
 "description": "Guide l'agent dans un diagnostic d'onduleur étape par étape",
 "arguments": [{"name": "symptome", "description": "Symptôme observé",
                "required": true}]}
```

`prompts/get` avec `{"symptome": "bypass permanent"}` renvoie une liste de
messages (rôles `user`/`assistant`) qui constituent le **début de
conversation**. C'est l'équivalent d'un slash-command `/diagnostic_ups`
qui injecte une procédure experte.

💡 Pour Zeef : transforme tes checklists terrain (commissioning onduleur,
recette copieur) en prompts MCP → chaque agent qui se connecte à ton
serveur hérite de tes procédures.

## 13. Quand utiliser quoi : tableau de décision

| Besoin | Primitive | Pourquoi |
|---|---|---|
| L'agent doit interroger ta BDD | Tool (`query`) | Action paramétrée, le modèle choisit quand |
| L'agent doit lire un doc précis | Resource (`doc://...`) | Donnée statique, l'humain l'attache |
| Lancer ta procédure de diagnostic | Prompt (`diagnostic_ups`) | Workflow figé, l'humain le déclenche |
| L'agent doit agir sur le monde (API, shell) | Tool | Effet de bord → annotations + confirmation |
| Exposer un catalogue (fichiers, tables) | Resources + templates | Navigable, listable |
| Donner du contexte « toujours utile » | Resource | Pas un appel, une pièce jointe |

⚠️ Erreur fréquente : tout mettre en tools. Si c'est de la donnée passive,
c'est une resource : ça coûte moins de tokens et c'est plus clair pour
le modèle.

## 14. Transport stdio : les règles d'or

Le transport **stdio** = le client lance ton serveur comme un processus fils
et dialogue via stdin/stdout. C'est le transport du développement local et
de 90 % des serveurs personnels.

```
client (Claude Code) ──stdin──► [serveur MCP] ──stdout──► client
                        ◄──stderr (logs uniquement) ──────
```

