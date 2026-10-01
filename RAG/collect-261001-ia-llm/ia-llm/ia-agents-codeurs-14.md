---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-14
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "claude", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [2275, 2434]
sha256: 94062b003007aafcb3ad23240840bed084df37ea52cf74d6ef68b402d0169aff
---

# Les agents codeurs IA

## 91. Conventions : faire écrire l'agent « comme l'équipe »

L'agent n'a pas de goût — il a ton AGENTS.md. Exemple de conventions qui
changent tout (à adapter) :

```markdown
# AGENTS.md — conventions d'équipe (extrait)

## Style
- Python : ruff, 100 colonnes max, type hints obligatoires.
- Nommage : français métier (onduleur, baie), anglais technique (retry, timeout).
- Commentaires : le POURQUOI, jamais le QUOI.

## Messages de commit (conventionnel)
- `feat:`, `fix:`, `docs:`, `refactor:`, `test:`, `chore:`
- Ex. : `feat: ajoute le retry SNMP avec backoff exponentiel`

## Tests
- Un fichier test_*.py par module, pytest, couverture > 80 % sur le nouveau code.
- Interdit : `assert True`, tests sans assertion, `time.sleep` dans les tests.

## Doc
- Toute nouvelle fonction publique : docstring + exemple.
- README mis à jour si l'usage change.
```

Astuce : fais générer **par l'agent** la première version de ces conventions
(« propose-moi un AGENTS.md à partir de nos 3 derniers projets »), puis
**corrige-la à la main**. C'est plus rapide que partir de zéro, et la
correction t'oblige à expliciter tes vrais standards.

## 92. Sécurité (1/3) : ce qu'on ne confie JAMAIS à un agent

Liste non négociable — à afficher au-dessus du poste :

1. 🔒 **Les secrets** : mots de passe, clés API, clés SSH, tokens. L'agent
   travaille avec des **exemples fictifs** et des `.env.example`. Un secret
   collé dans un prompt part chez le fournisseur et dort dans les logs de
   session.
2. 🔒 **La production sans garde-fou** : aucun agent ne touche la prod
   directement. Le chemin, c'est : agent → branche → PR → revue humaine →
   CI → déploiement contrôlé. Pas d'exception « c'est urgent ».
3. 🔒 **L'exécution aveugle** : `--dangerously-skip-permissions`,
   `--ask-for-approval never`, Full-auto… ces interrupteurs existent pour
   des sandboxes jetables, pas pour ton poste de travail.
4. 🔒 **Le réseau sans contrôle** : un agent qui peut `curl` peut exfiltrer.
   En phase de test, coupe le réseau (sandbox) quand la tâche ne l'exige pas.
5. 🔒 **Les suppressions** : `rm -rf`, `DROP`, écrasement de fichiers —
   toujours avec validation explicite, jamais en auto.
6. 🔒 **L'installation de dépendances** : `pip install`, `npm install`, PPA
   obscurs… chaque nouvelle dépendance = une chaîne d'approvisionnement à
   valider par un humain (typosquatting, paquets malveillants).
7. 🔒 **Les décisions irréversibles** : rotation de clés, suppression de
   sauvegardes, changement de pare-feu — l'agent **propose**, l'humain
   **exécute**.

## 93. Sécurité (2/3) : durcir chaque outil

| Outil | Réglage clé | Où |
|---|---|---|
| Claude Code | Permissions par défaut (pas `--dangerously-skip-permissions`), hooks de garde | `~/.claude/settings.json`, CLAUDE.md |
| Codex CLI | Rester en **Suggest/Auto-edit**, Full-auto = sandbox | flags `--ask-for-approval` |
| OpenCode | Agent **Plan** par défaut (Tab), modèles figés | `opencode.json` |
| Cursor | Privacy mode, règles `.cursor/rules/securite.mdc` | Settings + repo |
| Cline | Auto-approve : lecture seule auto, **tout le reste manuel** | Panneau Cline |
| Kilo Code | Permissions allow/ask/deny, worktrees pour isoler | `kilo.jsonc` |
| OpenClaw | Bind 127.0.0.1, user sans sudo, crons lecture seule | config + checklist section 70 |
| Tous | `.gitignore` béton, `git status` propre avant session | repo |

Le **fichier `.gitignore` de combat** (à copier dans chaque repo) :

```gitignore
# Secrets — JAMAIS versionnés, JAMAIS sous les yeux d'un agent
.env
.env.*
*.pem
*.key
id_rsa*
*.p12
*credentials*
*secrets*

# Système & build
__pycache__/
node_modules/
.venv/
*.log
```

Et le **bloc « sécurité » à mettre dans TOUS tes AGENTS.md** :

```markdown
## Sécurité (non négociable)
- Ne lis, ne crée, ne modifie JAMAIS de fichier contenant un secret.
  Utilise .env.example et des valeurs fictives (ex. : "motdepasse-fictif").
- Ne lance JAMAIS : rm -rf, sudo, pip install, curl | bash, sans validation
  humaine explicite dans cette session.
- N'envoie JAMAIS de données hors de la machine (pas de curl/post vers
  l'extérieur) sauf demande explicite.
- Avant toute action destructive, affiche ce que tu vas faire et attends "go".
```

## 94. Sécurité (3/3) : confidentialité des données et entraînement

Ton code **voyage** : chaque prompt part chez le fournisseur du modèle. Ce
qu'il faut savoir en 2026 :

- **Claude Code / Anthropic** : politique de non-entraînement sur les
  données API par défaut (l'API n'entraîne pas) — mais vérifie les
  conditions « à vérifier » le jour J, elles ont déjà changé par le passé.
- **Modèles gratuits / tiers gratuits** (offres « free » d'OpenRouter, Kilo,
  Trae…) : la règle quasi générale est que **les prompts gratuits peuvent
  servir à l'entraînement** — c'est le prix du gratuit. Ne jamais y mettre
  de code sensible ou client.
- **Cursor** : indexation du repo sur **son** infrastructure (cloud only).
  Le privacy mode promet « rien stocké », mais le code transite par leurs
  serveurs. Pour du code soumis au secret (client, R&D), tranche
  explicitement avant d'indexer.
- **BYOK via API** : chez la plupart des fournisseurs, l'usage **API payant**
  n'alimente pas l'entraînement — c'est l'argument confidentialité n°1 du
  BYOK (OpenCode, Cline, Continue en API).
- **100 % local (Ollama)** : rien ne sort. Le seul niveau « secret défense »,
  au prix d'un modèle moins brillant.

Checklist confidentialité par projet (à cocher) :

- [ ] Classification du repo : public / interne / sensible / secret
- [ ] « Sensible » et au-dessus : **pas d'offre gratuite**, API payante ou local
- [ ] Indexation cloud (Cursor) : oui/non, décidé et noté
- [ ] Logs de session (`~/.claude/`, `~..codex/`) : contiennent du code en
      clair → exclus des backups non chiffrés, nettoyés périodiquement

## 95. Coût et contrôle des tokens (1/2) : comprendre la facture

**Anatomie d'une session** : chaque tour de boucle agentique renvoie
l'historique + les nouveaux outils au modèle. Une session, c'est :

```
input  = prompt + fichiers lus + sorties d'outils + historique (grossit à chaque tour)
output = raisonnement + code généré
```

Ordres de grandeur (modèles « Sonnet/GPT-5 »-like, sept. 2026 — **à
vérifier**, les tarifs bougent) :

| Session type | Tokens input | Tokens output | Coût API indicatif |
|---|---|---|---|
| Question simple + 2 fichiers | ~10 k | ~2 k | **~0,05 $** |
| Petite feature + tests | ~100 k | ~20 k | **~0,50–1 $** |
| Feature moyenne (30 min) | ~500 k | ~80 k | **~2–5 $** |
| Grosse session autonome (1–2 h) | 2–5 M | 300–800 k | **~10–40 $** |

Le **cache de prompt** change tout : les fournisseurs facturent le
*cache read* ~10× moins cher que l'input normal. Les bons agents (Claude
Code en tête) cachent automatiquement le contexte stable (ton AGENTS.md,
les gros fichiers relus) — d'où l'intérêt de sessions qui **relisent les
mêmes fichiers** plutôt que d'en ouvrir 50 différents.

## 96. Coût et contrôle des tokens (2/2) : les 10 leviers d'économie

