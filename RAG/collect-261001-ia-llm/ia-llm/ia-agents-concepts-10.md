---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-10
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI", "Z.ai", "vLLM"]
dates: []
keywords: ["agent", "agents", "apache", "claude", "copilot", "glm", "open source", "parameters", "sandbox", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [1361, 1519]
sha256: 6603f379c23a6178b2673ceaa16c0d36078c9b5802acd518b533a36038442f55
---

# Concepts : agents IA, agentic, autonomie

- Versions desktop **3.11.2** (relevée début sept. 2026) et mentions de **3.14**
  côté docs d'intégrateurs : le rythme de release est rapide, vérifie la tienne.
- **Open-sourcing récent** : Z.ai a publié le code sous **licence Apache 2.0**
  (dépôt `zai-org/ZCode` rapporté) avec les trois interfaces (desktop/web/CLI).
  C'est un signal positif (audit possible), mais « open source depuis peu » =
  l'historique de maintenance reste à construire.
- Il existe aussi un fork communautaire « local-first » branché sur LM Studio
  (sans compte Z.ai) : intéressant si tu veux tester sans créer de compte,
  mais c'est du code tiers — même prudence que pour tout fork.

Site officiel rapporté : `zcode.z.ai` (à vérifier — ne télécharge que depuis
le site officiel ou le dépôt officiel, jamais depuis un lien de forum).

## 74. ZCode vs la concurrence (positionnement)

| Outil | Écosystème | Interface | Modèle de référence | Code |
|---|---|---|---|---|
| ZCode | Z.ai / GLM | desktop + web + CLI | GLM-5.x | Apache 2.0 (récent) |
| Claude Code | Anthropic | terminal | Claude | propriétaire |
| Codex CLI | OpenAI | terminal | GPT | open source (MIT, 2026) |
| Open Interpreter | indépendant | terminal | au choix | MIT (Python) / Rust (réécriture) |
| OpenCode | indépendant | terminal | au choix | open source |

ZCode se distingue par : le **centrage tâche** (tu confies un objectif, tu suis
l'avancement, tu interviens à distance) plutôt que le centrage fichier/éditeur,
et l'optimisation pour **GLM** (si ton équipe est « GLM-first », c'est la voie native).

## 75. ZCode pour un sysadmin : intérêt et limites

Intérêt :
- si tu veux un **harnais agentique clé en main** sans écrire ton harness Python,
  ZCode est un candidat crédible, surtout couplé à un modèle local ou à GLM ;
- les skills / `AGENTS.md` permettent d'y injecter tes runbooks (synergie avec
  ton RAG : tes docs deviennent des skills) ;
- le pilotage à distance des longues tâches parle à un chef de service
  (lancer un inventaire le soir, relire le matin).

Limites :
- exécution de code locale = **mêmes risques** qu'Open Interpreter (section 64) :
  sandbox, compte dédié, relecture ;
- modèle chinois, données : si tu branches le cloud Z.ai, tes prompts partent
  chez Zhipu — à cadrer selon ta politique de données (modèle local ou provider
  européen si besoin) ;
- jeunesse de l'open-sourcing : audite avant usage pro, épingle les versions.

## 76. Prise en main ZCode (marche à suivre prudente)

1. Télécharge **uniquement** depuis le site officiel ou le dépôt officiel
   (vérifie l'URL deux fois — homonymes et typosquatting existent).
2. Installe dans une **VM de test** d'abord (montage 3, section 52), pas sur ton poste.
3. Démarre avec un **provider local** (LM Studio / Ollama en mode compatible OpenAI)
   ou un provider dont tu maîtrises la politique de données.
4. Configure `AGENTS.md` avec tes règles : périmètre, interdictions, budgets.
5. Teste sur des tâches en lecture seule pendant une semaine (niveau 2).
6. Ajoute tes runbooks comme skills, un par un, en mesurant.
7. Ne monte en autonomie qu'avec la doctrine de la partie C.

## 77. Synthèse parties D-E-F : quel outil pour quel besoin

| Besoin | Piste | Pourquoi |
|---|---|---|
| Langage naturel → code local, contrôle total | Open Interpreter | simple, éprouvé, tu vois tout |
| Harnais agentique clé en main, sessions longues | ZCode | 3 interfaces, skills, pilotage distant |
| Agent fichiers pour non-tech (macOS/Anthropic) | Claude Cowork | pas d'exécution de code (preview — à revérifier) |
| Agent transverse en entreprise Microsoft | Copilot Cowork | gouvernance M365, facturation à l'usage |
| Agent perso local-first bricolable | CoWork OS / agent maison | contrôle, mais audit requis |
| Apprendre vraiment comment ça marche | Partie G (Python maison) | 150 lignes, zéro magie |

Règle générale : **l'outil ne remplace pas la doctrine.** Quel que soit le harnais,
les garde-fous de la partie C s'appliquent (budgets, approbations, sandbox, logs).

---

# PARTIE G — Construire son propre agent en Python

## 78. L'agent minimal : ce qu'il faut (et rien de plus)

Un agent ReAct maison tient en ~150 lignes. Les pièces :

```
agent.py
├── 1. config : modèle, budgets, dossiers autorisés
├── 2. outils : run_shell (liste blanche), find_files, read_file
├── 3. schémas d'outils (JSON) déclarés au modèle
├── 4. boucle : appel modèle → exécute outils → réinjecte → recommence
├── 5. budgets : étapes / tokens / temps (dur)
├── 6. approbation humaine pour les outils sensibles
└── 7. logs JSONL
```

Pas de framework, pas de magie. Quand tu auras fait tourner ça, tu comprendras
ce que les frameworks t'apportent (et ce qu'ils te cachent).

## 79. Prérequis : environnement Python

```bash
python3 -m venv ~/.venvs/agent && source ~/.venvs/agent/bin/activate
pip install openai   # SDK officiel ; toute API compatible OpenAI convient
export OPENAI_API_KEY="..."      # ou ta variable / ton endpoint local
export AGENT_MODEL="gpt-4.1-mini"  # à vérifier : nom exact chez ton provider
```

Notes :
- le code ci-dessous utilise l'API **Chat Completions** (`chat.completions.create`)
  avec `tools=` : c'est l'interface de function calling la plus standard et la
  plus portable (fonctionne aussi avec les endpoints compatibles OpenAI :
  Ollama, LM Studio, vLLM, etc. — noms de modèles à vérifier).
- `temperature=0` partout : déterminisme maximal pour de l'admin.
- tout tourne en local sauf l'appel au modèle.

## 80. Étape 1 : squelette, prompt système, déclaration d'outils

```python
"""agent.py — agent ReAct minimaliste (boucle + outils + budgets)."""
import json, os
from openai import OpenAI

MODEL = os.environ.get("AGENT_MODEL", "gpt-4.1-mini")  # à vérifier
WORKDIR = os.path.abspath(os.environ.get("AGENT_WORKDIR", "/srv/agent-work"))

client = OpenAI()  # lit OPENAI_API_KEY (ou OPENAI_BASE_URL pour un endpoint local)

SYSTEM = """Tu es un agent sysadmin. Tu travailles par cycles :
THOUGHT (raisonnement court) -> ACTION (un appel d'outil) -> OBSERVATION (résultat).
Règles :
- N'appelle qu'un outil à la fois quand l'étape suivante dépend du résultat.
- Ne devine jamais : vérifie (ex : après une écriture, relis le fichier).
- Si tu bloques 3 étapes sur la même piste, change de stratégie ou conclus.
- Réponse finale : résumé des actions + résultats VÉRIFIÉS + ce qui reste à faire.
- Interdictions : aucune donnée inventée, aucune action hors des outils fournis."""

TOOL_SCHEMAS = [
    {"type": "function", "function": {
        "name": "run_shell",
        "description": "Exécute UNE commande simple (sans pipe, sans redirection) "
                       "parmi une liste blanche de binaires. Lecture seule en pratique.",
        "parameters": {"type": "object",
            "properties": {"command": {"type": "string",
                "description": "Ex: 'df -h', 'systemctl status nginx'"}},
            "required": ["command"]}}},
    {"type": "function", "function": {
        "name": "find_files",
        "description": "Cherche des fichiers par motif, confinée à WORKDIR.",
        "parameters": {"type": "object",
            "properties": {
                "root": {"type": "string", "description": "Dossier de départ"},
                "pattern": {"type": "string", "description": "Motif glob, ex '*.log'"}},
            "required": ["root"]}}},
    {"type": "function", "function": {
        "name": "read_file",
        "description": "Lit un fichier texte (tronqué à 8000 caractères), confiné à WORKDIR.",
        "parameters": {"type": "object",
            "properties": {"path": {"type": "string"}},
            "required": ["path"]}}},
]
```

## 81. Étape 2 : l'outil `run_shell` sécurisé (liste blanche)

