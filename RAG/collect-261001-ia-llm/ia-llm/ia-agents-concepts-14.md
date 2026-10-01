---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-14
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agent", "agents", "arr", "benchmark", "benchmarks", "distribution", "incident", "leaderboard", "tool use"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [2096, 2232]
sha256: 1d2f38e3befe2b0818b875eeffa0c8ec009e260adc83abf5436b5cb029b04c5c
---

# Concepts : agents IA, agentic, autonomie

## 98. Ce que LangGraph ne fait pas à ta place (checklist)

Migrer vers LangGraph ne dispense pas de :

- [ ] **budgets durs** : `max_steps` via un compteur dans le State + arête
      conditionnelle vers END ; coût via cumul des `usage` dans le State ;
- [ ] **liste blanche d'outils** : le vérificateur (section 96) ou des `@tool`
      déjà confinés (sections 81-82 — réutilise ton code, il est framework-agnostic) ;
- [ ] **logs JSONL** : un nœud « logger » ou un callback sur chaque transition ;
- [ ] **tests d'injection** : inchangés (section 88) ;
- [ ] **sandboxing** : l'outil shell reste le même, avec les mêmes montages (section 52).

Le framework structure la boucle ; **la sécurité reste dans tes outils et tes
nœuds garde-fous.** Un graphe élégant avec des outils naïfs, c'est juste un
incident bien architecturé.

---

# PARTIE I — Évaluer un agent : benchmarks et tests maison

## 99. Pourquoi évaluer (et ce que les benchmarks ne disent pas)

Tu ne peux pas améliorer ce que tu ne mesures pas — et tu ne peux pas faire
confiance à ce que tu n'as pas mesuré **toi-même**. Les benchmarks publics
servent à comparer des modèles/harness entre eux ; ils ne disent pas si TON
agent réussira TON travail, avec TES outils, sur TES systèmes.

Limites connues des benchmarks académiques (relevées 2026) :
1. **contamination** : les questions/réponses publiques ont pu entrer dans les
   données d'entraînement → scores gonflés (d'où SWE-bench Verified, SWE-bench-Live) ;
2. **décalage de distribution** : tes tâches ≠ leurs tâches ;
3. **sur-optimisation** : un harness réglé pour un benchmark ne transfère pas
   forcément ;
4. **conditions incohérentes** : sous-ensembles, nombre d'essais et jeux d'outils
   différents → scores non comparables entre publications.

Bonne pratique : utilise les benchmarks pour **choisir un modèle/harness**,
puis construis ton **eval set maison** (section 103) pour valider TON agent.

## 100. Les benchmarks publics : le tableau (sept 2026, à vérifier)

Scores relevés dans des revues/synthèses 2026 — **ordres de grandeur, à revérifier**,
les classements bougent vite :

| Benchmark | Domaine | Ce qu'il mesure | Référence de score (frontier, 2026) |
|---|---|---|---|
| GAIA | assistant généraliste (466 questions multi-outils) | réponse exacte | ~75 % (niveau 1) — SOTA agents |
| GAIA2 | environnements dynamiques | pass@1 | ~42 % (GPT-5 high) |
| SWE-bench Verified | correction de vrais bugs GitHub | tests du patch | ~46-80 % selon harness (fourchette large !) |
| Terminal-Bench | 89 tâches terminal (sysadmin, sécu, data) | assertions d'état | ~70 % (tier 1) → ~30 % (tier 3) |
| OSWorld | usage d'un vrai OS desktop (369 tâches) | état du système | < 25 % (non vérifié) → ~73 % (version « verified ») |
| WebArena | navigation web réaliste | état fonctionnel | ~35-52 % |
| τ-bench / τ²-bench | service client + **respect des règles** | succès + conformité | variable selon politique |
| AgentBench | 8 environnements (OS, DB, web, jeux) | score normalisé | 30-70 % selon env. |
| BFCL | **qualité du function calling** | appels corrects | référence pour le tool use |
| AgentDojo | **résistance aux injections** | succès + attaques bloquées | benchmark sécurité |
| HAL | leaderboard holistique (Princeton) | **précision vs coût** | Pareto accuracy/coût |
| BrowseComp | recherche profonde d'info | réponse + preuve | retrieval difficile |
| TheAgentCompany | entreprise simulée (GitLab, chat, cloud) | crédit partiel long-horizon | tâches longues réalistes |

Deux à retenir pour toi : **Terminal-Bench** (le plus proche du sysadmin :
le harness y fait varier le score de 30-50 points à modèle égal — la preuve que
ton `agent.py` compte autant que ton modèle) et **τ-bench** (le seul qui note
le **respect des règles**, pas juste le résultat — essentiel pour un agent
d'exploitation).

## 101. Lire un score sans se faire avoir

- **pass@k vs pass^k** : `pass@k` = proba de réussir **au moins 1 fois sur k**
  (optimiste, mesure le potentiel) ; `pass^k` (τ-bench) = proba de réussir
  **les k fois** (pessimiste, mesure la fiabilité). Pour de la prod, c'est
  `pass^k` qui compte : un agent qui réussit 1 fois sur 3 est inutilisable en
  astreinte.
- **single-pass ≠ prod** : un score « 45 % » en une passe ne veut pas dire
  « 45 % de réussite en prod » (pas de retry, environnement propre, scaffolding
  du benchmark).
- **le harness est noté avec le modèle** : quand un leaderboard dit « modèle X :
  82 % », lis « modèle X + leur harness + leurs outils : 82 % ». Change un
  élément, le score change.
- **coût par point** : HAL (Princeton) a raison de tracer précision vs coût.
  Un agent à 80 % qui coûte 2 €/tâche bat un agent à 85 % qui coûte 20 €/tâche,
  sauf si l'échec coûte cher (alors c'est l'inverse — fais le calcul).

## 102. Les métriques qui comptent en prod (les tiennes)

Pour chaque agent en service, suis ces 6 métriques — pas les scores académiques :

| Métrique | Définition | Seuil d'alerte (exemple) |
|---|---|---|
| Taux de succès | tâches réussies (vérifiées) / tâches totales | < 90 % sur 50 tâches → investigation |
| Fiabilité (`pass^3`) | réussite sur 3 exécutions identiques | < 80 % → non déployable en autonome |
| Coût / tâche | tokens × prix + temps machine | > budget alloué → optimiser ou arrêter |
| Latence p95 | durée murale, 95e percentile | > SLA du cas d'usage → paralléliser/simplifier |
| Taux d'intervention | % de tâches nécessitant une validation/correction humaine | > 30 % → le « gain » est illusoire |
| Incidents | actions hors périmètre / garde-fou déclenché | **> 0** → revue immédiate |

Ajoute : **conformité** (l'agent a-t-il respecté les règles ? style τ-bench)
dès que l'agent touche à des processus encadrés (tickets, changements, sauvegardes).

## 103. Tests maison : construire ton eval set (méthode)

Un eval set maison = 20-50 tâches réelles, chacune avec un **correcteur automatique**.
Méthode en 6 étapes :

1. **Collecte** : prends 30 tâches que tu fais vraiment (bilan disque, analyse de
   log, génération de rapport, recherche doc, tri de ticket…).
2. **Formalisation** : pour chacune, un objectif en une phrase + un état initial
   reproductible (VM snapshot, dossier de test).
3. **Correcteur** : un script qui dit OK/KO **sans LLM** (le fichier existe-t-il ?
   le service répond-il ? le rapport contient-il les 5 champs ?). Les correcteurs
   LLM-juges sont un second choix (biais, coût).
4. **Difficulté** : tag chaque tâche (facile/moyen/difficile) pour lire où l'agent
   échoue.
5. **Sécurité** : inclus 5 tâches **piégées** (fichier avec fausse instruction,
   demande ambiguë dangereuse, site/outils qui mentent) → l'agent doit refuser
   ou demander. Une tâche piégée réussie par l'agent = **échec de sécurité**.
6. **Rituel** : lance l'eval set à chaque changement (modèle, prompt, outil,
   version). Note les 6 métriques (section 102). C'est ta CI d'agent.

Exemple de correcteur (bash, 5 lignes) :

```bash
#!/bin/bash
# correcteur : l'agent devait produire /srv/agent-work/rapport.md avec 5 sections
f=/srv/agent-work/rapport.md
[ -f "$f" ] || { echo "KO: fichier absent"; exit 1; }
n=$(grep -c '^## ' "$f")
[ "$n" -ge 5 ] || { echo "KO: $n sections (< 5)"; exit 1; }
echo "OK: rapport conforme ($n sections)"
```

## 104. Le harnais d'éval minimal (Python)

```python
"""eval.py — lance l'agent sur N tâches, corrige, résume."""
import json, subprocess, time

