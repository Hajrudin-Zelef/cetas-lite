---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-5
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: ["2026-09-10"]
keywords: ["agent", "agents", "arr", "claude", "cost", "mcp", "open source", "pricing"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [719, 913]
sha256: 18323014e23680a2e3196266806f921a7e258d538da628e28c1b89100b64aae6
---

# Les agents codeurs IA

## 24. Claude Code : premier workflow commenté

Scénario : ajouter une commande `inventaire` à un CLI Python, **méthode
plan d'abord**.

```bash
cd /tmp/lab-agent
claude
```

```
# 1. Cadrage du projet (une seule fois par repo)
> /init
# → génère CLAUDE.md. CORRIGE-LE : commandes de test, interdictions (cf. section 6).

# 2. Plan mode : Shift+Tab, puis
> [Plan mode] Voici le besoin : ajouter une sous-commande `inventaire`
> qui liste les équipements depuis un CSV. Explore le code existant et
> propose un plan détaillé (fichiers touchés, structure, tests).
> N'écris aucun code pour l'instant.

# 3. Lis le plan. Demande des ajustements si besoin :
> ton plan prévoit de modifier cli.py directement ; préfère un nouveau
> module inventaire.py importé dans cli.py. Mets à jour le plan.

# 4. Sors du plan mode (Shift+Tab), puis exécution :
> implémente le plan validé. Après chaque fichier modifié, lance
> `pytest -q` et corrige jusqu'au vert.

# 5. Contrôle : /cost pour voir ce que la session a coûté (route API),
#    git diff pour relire chaque ligne avant commit.
```

Réflexes à garder :

- **`Esc`** dès que l'agent part dans une direction bizarre — corriger tôt
  coûte 10× moins cher que corriger tard.
- **`/compact`** quand la session dépasse ~30 min : résume le contexte,
  évite la dérive.
- **Un commit git avant chaque grosse étape** : `git diff` reste le meilleur
  outil de revue d'un agent.

## 25. Claude Code : subagents, skills, hooks, MCP

C'est ici que Claude Code se distingue pour un usage avancé.

**Subagents** (`.claude/agents/`) : des agents spécialisés avec leur propre
prompt système et leurs propres outils autorisés.

```markdown
# .claude/agents/reviewer.md
---
name: reviewer
description: Relit un diff et rend un verdict bloquant/non-bloquant
tools: Read, Grep, Glob   # pas d'écriture, pas de shell
---
Tu es un relecteur exigeant. On te donne un diff.
Rends : 1) verdict BLOQUANT ou OK, 2) la liste des problèmes par sévérité,
3) pour chaque problème bloquant, le patch exact proposé.
Ne modifie jamais de fichier toi-même.
```

Usage : `> utilise le subagent reviewer sur le diff de la branche feature/inventaire`.

**Skills** (`.claude/skills/`) : savoir-faire réutilisables, invoqués par
`/skill` ou automatiquement.

```bash
.claude/skills/ansible-review/SKILL.md   # ta checklist de relecture de playbooks
.claude/skills/rag-ingest/SKILL.md       # ta procédure d'ingestion RAG
```

**Hooks** : scripts shell déclenchés par événements (après écriture d'un
fichier, avant un commit…). Exemple : relancer le lint après chaque
modification Python.

```json
// ~/.claude/settings.json — extrait (schéma indicatif, vérifier la doc)
{
  "hooks": {
    "PostToolUse": [
      { "matcher": "Edit|Write", "command": "ruff check ${file} || true" }
    ]
  }
}
```

**MCP** : Claude Code est **client ET serveur** MCP — le support le plus
mature du marché en 2026.

```bash
# Ajouter un serveur MCP (ex. : filesystem sur ton dossier de docs)
claude mcp add docs -- npx -y @modelcontextprotocol/server-filesystem ~/workspace/user/files
# Lister
claude mcp list
```

> 💡 Pour ton RAG : un serveur MCP « docs » branché sur ton corpus permet à
> l'agent de citer tes propres guides pendant qu'il code. C'est le pont entre
> tes deux chantiers.

## 26. Claude Code : bonnes pratiques (terrain)

1. **Toujours `/init` + relecture du CLAUDE.md** sur un nouveau repo.
2. **Plan mode par défaut** pour toute tâche > 10 minutes. Le réflexe
   `Shift+Tab` doit devenir un automatisme.
3. **Prompts avec critères d'arrêt** : « …et arrête-toi dès que les tests
   passent, résume ce que tu as changé » — sinon l'agent « améliore »
   indéfiniment (et facture).
4. **Donne les chemins, pas des descriptions vagues** : `@app/main.py:120`
   bat « la fonction qui gère le login ».
5. **Interdis explicitement** dans CLAUDE.md : pas de `rm -rf`, pas de
   modification des migrations, pas de nouvelle dépendance sans validation.
6. **Revue systématique** : `git diff` + subagent `reviewer` avant chaque
   commit (section 100).
7. **Sessions courtes** : au-delà de 45–60 min, `/compact` ou nouvelle
   session avec un résumé — la qualité dérive avec la longueur.
8. **Ne colle jamais de secrets dans le prompt** : ils partent chez Anthropic
   et restent dans l'historique de session (`~/.claude/`).

## 27. Claude Code : prix (vérifiés sept. 2026)

| Route | Prix | Notes |
|---|---|---|
| Claude **Pro** | **20 $/mois** (17 $/mois en annuel) | Claude Code **inclus**, usage « raisonnable » |
| Claude **Max** | **100 $/mois** (5× l'usage Pro) ou **200 $/mois** (20×) | Pour usage intensif ; Code + chat partagent le quota |
| Plan gratuit | 0 € | **N'inclut PAS** Claude Code (vérifié 10/09/2026) |
| **API** (clé) | Au token, ex. **Sonnet 4.6 : 3 $ / 15 $ par million de tokens** (input/output) ; cache : ~3,75 $ écriture / 0,30 $ lecture | Tarifs officiels Anthropic ; les familles Opus/Haiku ont leurs propres grilles — **à vérifier** sur platform.claude.com/docs/en/about-claude/pricing |

Repères de coût en API (ordres de grandeur 2026) :

- Petite session (question + 2–3 fichiers) : **quelques centimes**.
- Feature moyenne avec tests : **0,50–3 $**.
- Grosse session autonome d'1 h sur Sonnet : **5–15 $**.
- Sur Opus : **3–5× plus cher** que Sonnet.

> 💰 `/cost` en cours de session + alertes de facturation sur la console
> Anthropic : les deux garde-fous à activer dès la route API.

## 28. Claude Code : points forts / faiblesses honnêtes

**Points forts**

- L'agent le plus **poli et collaboratif** : il pose des questions, admet
  l'incertitude, conteste un besoin flou — comportement de « pair programmer ».
- Tient les **longues sessions** mieux que la concurrence (contexte 1M+
  tokens, moins de dérive).
- Écosystème MCP le plus mature (client + serveur), skills, hooks, subagents.
- Cinq surfaces avec le même moteur (terminal, VS Code, JetBrains, desktop,
  web + `--teleport`).
- Doc et communauté énormes : un problème = quelqu'un l'a déjà eu.

**Faiblesses honnêtes**

- **Verrouillage Anthropic total** : un seul fournisseur, un seul jeu de
  modèles. Panne ou hausse de prix = pas de plan B dans l'outil.
- Le CLI est **fermé** (pas open source) — tu ne peux pas l'auditer ni le
  forker.
- En abonnement, les limites d'usage (« fair use ») sont floues et ont
  changé plusieurs fois en 2026 : un mois tu es large, le suivant tu es
  bridé.
- Tendance à la **verbosité** : il explique beaucoup, ça coûte des tokens
  de sortie. À cadrer (« sois concis » dans CLAUDE.md).
- Ton code transite par les serveurs d'Anthropic : à intégrer dans ton
  analyse de confidentialité (section 99).

## 29. Claude Code : usage en équipe (conventions)

Pour que 3–5 personnes utilisent Claude Code sans chaos :

```bash
# À versionner dans chaque repo
CLAUDE.md                  # ou AGENTS.md + symlink
.claude/
  agents/
    reviewer.md            # relecteur bloquant partagé
    test-writer.md         # génère les tests selon VOS standards
  skills/
    commit-msg/SKILL.md    # format de messages de commit de l'équipe
  settings.json            # hooks partagés (lint auto après écriture)
```

Conventions à écrire noir sur blanc :

1. **Modèle par défaut d'équipe** (ex. : Sonnet pour le quotidien, Opus
   uniquement sur demande motivée — question de coût).
2. **Plan mode obligatoire** au-delà d'une estimation de 15 min de travail.
3. **Aucun commit généré sans `git diff` relu par un humain** (section 100).
4. **Skills d'équipe > prompts individuels** : quand deux personnes
   réécrivent le même long prompt, c'en est un skill.
5. **Journal des sessions** : un `docs/decisions/` où l'agent résume ce qu'il
   a fait et pourquoi (prompté systématiquement en fin de feature).

## 30. Claude Code : verdict Zelef

