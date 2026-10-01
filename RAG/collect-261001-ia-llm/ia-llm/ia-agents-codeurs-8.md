---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-8
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "arr", "attention", "claude", "mcp", "pricing"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1284, 1439]
sha256: c223b8ba0d14edb224224d074d86134f6644c89f46fea9b4beb37c9429af9b58
---

# Les agents codeurs IA

```
1. Ctrl+I → panneau Composer.
2. Prompt : « Ajoute une option --format json à la commande inventaire,
   mets à jour le README et les tests. »
3. Cursor propose les fichiers touchés → tu valides le périmètre.
4. Il génère les diffs fichier par fichier → tu relis (diff agrégée).
5. Apply → les modifs s'appliquent, tu peux checkpoint/annuler.
```

**Agent (`Ctrl+L` puis mode Agent, ou onglet Agent)** — l'autonome :

```
1. Bascule en mode Agent.
2. Prompt avec objectif + garde-fous :
   « Ajoute la persistance SQLite au script d'inventaire.
     Contraintes : ne touche pas à cli.py, tests pytest verts à la fin,
     arrête-toi et résume si un choix d'architecture se présente. »
3. L'agent travaille : tu vois les outils appelés, les diffs se construisent.
4. Review : chaque fichier a un diff à accepter/rejeter.
5. Background : pour les longues tâches, « Run in background » → tu continues
   à coder pendant qu'il bosse.
```

**Checkpoints** : Cursor versionne l'état avant chaque action d'agent —
tu peux revenir en arrière d'un cran (`Restore checkpoint`). C'est ton
`git stash` d'urgence quand tu n'as pas commité.

## 44. Cursor : premier workflow commenté

Scénario : **refactorer** un script monolithique en modules, avec revue.

```
1. Ouvre /tmp/lab-agent dans Cursor. Attends la fin de l'indexation.
2. Crée .cursor/rules/python.mdc (section 42) — 5 lignes de conventions.
3. Sélectionne le gros fichier, Ctrl+I (Composer) :
   « Découpe ce script en 3 modules (collecte, formatage, cli) sans changer
     son comportement. Propose d'abord la découpe, attends ma validation. »
   → Composer propose la découpe. VALIDE ou CORRIGE avant l'exécution.
4. Il génère les diffs → relis chaque diff (clic sur fichier).
5. Passe en Agent pour la suite :
   « Ajoute les tests pytest manquants pour les 3 modules. Lance-les.
     Si un test échoue, corrige le code, pas le test, sauf si le test
     est objectivement faux — dans ce cas explique-moi pourquoi. »
6. Vérifie : terminal intégré → pytest -q. git diff → relecture finale.
```

Ce que ce workflow t'apprend : **Cursor excelle quand tu alternes les
couches** (Composer pour le périmètre précis, Agent pour l'autonomie
cadrée), et il punit le « prompt unique magique » (diffs énormes,
irrelisibles, à jeter).

## 45. Cursor : prix (vérifiés sept. 2026 — À VÉRIFIER, ça a bougé en 2026)

⚠️ **Les grilles Cursor ont été restructurées en 2026** et les sources se
contredisent sur le détail. Chiffres recoupés (sept. 2026), **à confirmer
sur cursor.com/pricing avant tout engagement** :

| Plan | Prix indicatif | Contenu indicatif |
|---|---|---|
| **Hobby** | 0 € | Quota limité (complétions + requêtes premium lentes) — pour tester |
| **Pro** | **~20 $/mois** | Tab illimité, requêtes rapides incluses, agents |
| **Pro+** | **~60 $/mois** | ~3× l'usage Pro |
| **Ultra** | **~200 $/mois** | ~20× l'usage, fonctions prioritaires |
| **Teams/Business** | **~40 $/utilisateur/mois** | Facturation centrale, SSO, analytics, privacy mode |
| **Bugbot** (add-on) | ~40 $/utilisateur/mois | Revue de PR par IA avec règles |

Points d'attention :

- La facturation passe par des **crédits d'usage** : le « Pro à 20 $ »
  inclut une enveloppe d'usage API (~20 $) — au-delà, ça se paie. Surveille
  le dashboard d'usage.
- **Étudiants** : Pro gratuit 1 an (selon l'offre en cours — à vérifier).
- Le modèle **Composer** (maison) est optimisé pour être rapide et
  économique dans l'éditeur : l'utiliser par défaut réduit la facture.
- Certaines sources (fév. 2026) évoquent un plan **Individual à 60 $/mois**
  remplaçant l'ancien Pro à 20 $ — **contradiction non tranchée** :
  vérifie la page pricing le jour où tu t'abonnes.

## 46. Cursor : points forts / faiblesses honnêtes

**Points forts**

- L'expérience éditeur la plus **polie** : indexation repo, diffs agrégées,
  checkpoints, background agents — tout est pensé pour ne pas te ralentir.
- **Tab** excellent (renforcé par Supermaven) : le gain quotidien le plus
  tangible.
- **Composer** (modèle maison) : rapide, bien réglé pour le code in-éditeur.
- Règles projet (`.cursor/rules/`) + MCP : le cadrage d'équipe est simple.
- Écosystème : Bugbot (revue PR), Cloud Agents, intégrations CI.

**Faiblesses honnêtes**

- **Enfermement éditeur** : ton investissement (règles, habitudes) ne te
  suit pas dans le terminal ni ailleurs. Les CLI sont agnostiques de
  l'éditeur.
- **Facturation par crédits** : moins lisible qu'un forfait ou qu'une
  facture API au token. Des utilisateurs se font surprendre.
- **Confidentialité** : ton repo est **indexé sur l'infra Cursor** (cloud
  only, pas d'option on-prem en 2026). Le « privacy mode » promet de ne
  rien stocker, mais le code transite par leurs serveurs — à évaluer
  (section 99), surtout pour du code client sensible.
- Fork VS Code : tu dépends du rythme d'Anysphere pour les nouveautés de
  l'éditeur lui-même.
- Tentation du « tout-Agent » : l'outil rend si facile de déléguer qu'on
  oublie de comprendre — dangereux pour un sysadmin responsable d'une prod.

## 47. Cursor : bonnes pratiques

1. **Règles projet dès le jour 1** (`.cursor/rules/`) : c'est 80 % de la
   qualité des sorties.
2. **Composer pour le précis, Agent pour l'autonome** — ne demande jamais à
   l'Agent ce que Composer fait en 30 secondes.
3. **Valide le périmètre de fichiers AVANT** la génération (Composer le
   propose : ne clique pas « tout accepter » par flemme).
4. **Checkpoints + git** : checkpoint pour l'urgence, commit git pour le
   durable.
5. **Modèle par défaut = Composer** (rapide/pas cher) ; bascule sur
   Claude/GPT pour l'architecture et le debug profond.
6. **Relis les diffs comme si un stagiaire les avait écrites** — parce que
   c'est exactement ce qui s'est passé (section 100).
7. **Dashboard d'usage chaque semaine** tant que tu n'as pas calibré ta
   consommation.

## 48. Cursor : usage en équipe

```bash
.cursor/
  rules/
    python.mdc        # standards partagés (versionnés !)
    ansible.mdc       # tes conventions Ansible d'équipe
    securite.mdc      # interdictions (secrets, prod, sudo…)
  mcp.json            # serveurs MCP d'équipe (docs internes, tickets)
```

Conventions d'équipe :

1. **Les règles sont versionnées et relues** comme du code (PR pour modifier
   `.cursor/rules/`).
2. **Un Bugbot bien réglé** vaut mieux que trois règles floues : encodez vos
   standards de revue dedans.
3. **Privacy mode activé par défaut** + charte : quel code a le droit d'être
   indexé (pas les repos clients sensibles sans accord).
4. **Budget par développeur** : alerte à 80 % de l'enveloppe mensuelle.

## 49. Cursor : verdict Zelef

**Le meilleur choix si tu veux coder dans un éditeur sans te poser de
questions.** Pour un sysadmin qui écrit du Python/Ansible au quotidien,
le duo **Tab + Composer** fait gagner un temps fou sur le boilerplate.
En revanche, pour tes usages « RAG » (ingestion, nettoyage de corpus en
batch), un CLI scriptable (OpenCode/Codex) est plus adapté que Cursor.
Et la question de l'indexation cloud doit être tranchée avant d'y mettre du
code sensible. Idéal en **complément** d'un CLI, pas en remplacement.

## 50. Récapitulatif Partie A+B (les 5 grands)

