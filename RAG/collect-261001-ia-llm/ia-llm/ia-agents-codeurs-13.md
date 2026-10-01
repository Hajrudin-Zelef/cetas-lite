---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-13
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: []
keywords: ["agent", "agents", "apache", "aws", "chatgpt", "claude", "copilot", "gemini", "mcp", "pricing"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [2116, 2274]
sha256: 3a0b5c0bd0ebd656ce72b30a9adddc7af8600933cebac1c3809991b09cbb0cce
---

# Les agents codeurs IA

- **Gemini CLI** (Google) : le CLI agentique de Google, branché sur Gemini.
  Logique si tu es déjà client Google Cloud / amateur des longs contextes
  Gemini.
- **Kiro** (AWS) : l'IDE agentique d'AWS (fork VS Code) + **Kiro CLI**.
  Argument prix : accès aux modèles Claude via la facturation AWS, annoncé
  comme nettement moins cher que l'abonnement direct (ordre de grandeur
  évoqué : ~1/10 — **à vérifier**, dépend de ton contrat AWS).
- **Devin** (Cognition) : « l'ingénieur logiciel autonome » — pionnier des
  agents qui travaillent seuls des heures. Plans Pro ~20 $/mois, Max
  ~200 $/mois (à vérifier). Positionné « équipe virtuelle » plus
  qu'assistant.
- **Jules** (Google) : agent cloud **asynchrone** — tu lui assignes des
  tâches (web, CLI, API, labels GitHub), il travaille dans une VM Google
  Cloud et livre des **PR**, avec un agent « critique » qui relit avant de
  présenter. Offre gratuite : ~15 tâches/jour (à vérifier). Le « collègue
  virtuel » le plus mature côté cloud.

## 84. Récapitulatif plugins : que choisir selon le besoin

| Besoin | Choix 2026 | Pourquoi |
|---|---|---|
| Complétion pure, rapide | **Supermaven** ou Tab de Cursor | Vitesse, 0 friction |
| Complétion privée (code ne sort pas) | **Tabnine** / Continue + Ollama local | Local-first |
| Rester sur VS Code + OSS | **Continue** | BYOK, Apache 2.0 |
| Entreprise standardisée GitHub | **Copilot** | Intégration, prix, conformité |
| Écosystème AWS | **Amazon Q** | Connaissance AWS native |
| IDE JetBrains (PyCharm…) | **JetBrains AI** | Indexation sémantique |
| Éditeur rapide + mes agents | **Zed** + ton CLI | ACP, Rust, OSS |
| Tester à moindre coût | **Trae** (offre gratuite) | Généreux en free tier |
| CLI git-first sobre | **Aider** | Commits auto, offline possible |

## 85. Stack recommandée « Zelef » (sysadmin + RAG)

Une stack cohérente, sans doublon, calibrée pour ton profil :

```
COMPLÉTION ......... Supermaven (ou Copilot si déjà abonné GitHub)
ÉDITEUR/AGENT ...... Kilo Code (OSS) ou Cursor (confort) — PAS les deux
TERMINAL/CI ........ OpenCode (BYOK) + Codex CLI si ChatGPT Plus
APPRENTISSAGE ...... Cline (2 semaines, tout valider)
VIE PERSO/INFRA .... OpenClaw (durci, lecture seule d'abord)
```

Coût mensuel indicatif de cette stack : **0–40 $** selon les abonnements
déjà détenus (le détail par outil est dans chaque fiche prix).

---

# PARTIE E — TRANSVERSAL : BIEN UTILISER LES AGENTS

---

## 86. Grand comparatif final (septembre 2026)

| Critère | OpenCode | Claude Code | Codex CLI | Cursor | Cline | Kilo Code | OpenClaw |
|---|---|---|---|---|---|---|---|
| Type | CLI/TUI | CLI | CLI | Éditeur | Extension VS Code | Ext.+CLI | Agent perso |
| Licence | MIT | Fermé | Apache 2.0 | Propriétaire | MIT | MIT | MIT |
| Modèles | Tous (BYOK) | Claude only | GPT only | Multi + maison | BYOK | Tous (BYOK) | Tous (BYOK) |
| Prix outil | 0 € | 20 $/mois ou API | Inclus Plus 20 $ ou API | ~20 $/mois* | 0 € | 0 € | 0 € |
| Plan mode | Tab | Shift+Tab | /plan | Périmètre validé | Plan/Act | Architect | N/A (autonome) |
| MCP | Oui | Oui (client+serveur) | Oui (client) | Oui | Oui (manuel) | Oui (marketplace) | Oui |
| Skills | Oui | Oui | Non | Oui | Modes JSON | Oui (standard) | ClawHub |
| Scriptable CI | `opencode run` | `claude "…"` | **`codex exec`** | Cloud Agents | CLI 2.0 | `kilo` | Crons |
| Mémoire projet | AGENTS.md | CLAUDE.md/AGENTS.md | AGENTS.md | .cursor/rules | Modes perso | AGENTS.md | Mémoire persistante |
| Verrouillage | Aucun | Anthropic | OpenAI | Anysphere | Aucun | Aucun | Aucun |
| Idéal pour | Liberté modèle, terminal | Référence généraliste | Déjà abonné Plus, CI | Confort éditeur | Apprendre, contrôle | OSS multi-surfaces | Vie/infra autonome |

\* Cursor : grille restructurée en 2026 — **à vérifier** sur cursor.com/pricing.

## 87. Comment choisir selon le besoin (guide express)

```
BESOIN : écrire du code au quotidien dans un éditeur
└── Je veux le max de confort → Cursor
└── Je veux de l'OSS dans VS Code → Kilo Code (ou Cline pour apprendre)

BESOIN : automatiser en terminal / CI / cron
└── Déjà abonné ChatGPT Plus → Codex CLI (inclus, codex exec)
└── Je veux choisir mes modèles → OpenCode (BYOK)
└── Je veux LA référence, budget OK → Claude Code

BESOIN : comprendre et contrôler (phase d'apprentissage)
└── Côté éditeur → Cline (tout valider, 2 semaines)
└── Côté terminal → Aider (git-first, sobre)

BESOIN : un assistant de vie / une vigie infra
└── OpenClaw (durci, lecture seule d'abord)

BESOIN : juste de la complétion pendant la frappe
└── Supermaven / Copilot / Tabnine (ne paie QUE ça)
```

## 88. Matrice de décision : abonnement vs API (BYOK)

|  | Abonnement (20 $/mois type) | API au token (BYOK) |
|---|---|---|
| Facture | Fixe, prévisible | Variable, à surveiller |
| Plafond | « Fair use » flou, bridage possible | Pas de plafond, tu paies tout |
| Petit usage (< 1 h/sem) | Cher au final | **Gagnant** (quelques $) |
| Usage régulier quotidien | **Gagnant** (amorti vite) | 20–80 $/mois possibles |
| Usage en rafales (sprints) | Bridage possible au mauvais moment | **Gagnant** (paie les pics) |
| Multi-outils, 1 clé | Non (1 abo = 1 outil) | **Gagnant** (1 clé, N outils) |
| Tranquillité mentale | **Gagnant** | Alertes à configurer |

Règle : **commence en abonnement** (prévisible), **passe en BYOK** quand tu
sais mesurer ta conso (section 98) — ou quand tu veux une seule clé pour
plusieurs outils.

## 89. Usage en équipe : la charte minimale

Fais signer (même informellement) ces 10 règles avant de généraliser un
agent dans l'équipe :

1. **AGENTS.md versionné** dans chaque repo, relu comme du code.
2. **Plan mode obligatoire** au-delà de ~15 min de travail estimé.
3. **Aucun commit IA sans relecture humaine du diff** (section 100).
4. **Modèle par défaut d'équipe** + règle d'escalade (qui a le droit à Opus ?).
5. **Interdictions écrites** : secrets, prod, `rm -rf`, dépendances sauvages.
6. **Skills d'équipe** plutôt que prompts individuels qui se dupliquent.
7. **Budget** : enveloppe mensuelle par personne + alerte à 80 %.
8. **Confidentialité** : quels repos ont le droit d'être indexés/envoyés où.
9. **Journal** : l'agent résume chaque feature dans `docs/decisions/`.
10. **Un seul écrivain** : un agent en écriture à la fois par repo.

## 90. Revue du code généré : la méthode

Le code d'un agent se relit **comme celui d'un prestataire brillant mais
junior** : vite, bien, et avec des angles morts.

```
REVUE EN 4 PASSES (10–15 min pour une feature moyenne) :

PASSE 1 — Le diff brut
  $ git diff --stat          # l'ampleur : >500 lignes = méfiance
  $ git diff                 # lire CHAQUE ligne, pas en diagonale

PASSE 2 — Les tests
  $ pytest -q                # vert ? VRAIMENT vert (pas de tests skippés en douce)
  # Cherche : assert True, tests vides, mocks qui mockent le bug

PASSE 3 — Les angles morts de l'IA
  - Gestion d'erreurs : que se passe-t-il si le réseau tombe ? le fichier manque ?
  - Secrets : une clé en dur ? un mot de passe en clair dans un log ?
  - Dépendances : une lib ajoutée sans raison ? (pip freeze avant/après)
  - Sécurité : injection (SQL, shell), permissions de fichiers, sudo

PASSE 4 — Le subagent relecteur (double filet)
  > utilise le subagent reviewer sur cette branche
  # Il voit ce que ton œil fatigué rate. Toi tu tranches.
```

Signaux d'alerte (« à jeter et recommencer ») :

- Le diff touche des fichiers hors périmètre sans justification.
- Des tests ont été **affaiblis** pour passer (seuils baissés, cas supprimés).
- Du code dupliqué au lieu d'une fonction (l'agent a la flemme de refactorer).
- Des commentaires qui décrivent un comportement **différent** du code.

