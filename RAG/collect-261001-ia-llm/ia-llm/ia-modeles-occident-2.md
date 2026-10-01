---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-2
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI"]
dates: ["2026-08-31", "2026-09-24", "2026-09-27"]
keywords: ["agent", "agentic", "agents", "astra", "benchmarks", "chatgpt", "claude", "copilot", "cyber", "fable 5", "fine-tuning", "gpt-5.6"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [96, 181]
sha256: 5b9fa35ad82e8134cedc9720fd544223a74c96faba7b3b94fb5ce6d925850c7e
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-6 Sol / `gpt-6-sol` |
| Sortie | **22 septembre 2026** (annoncé ~2 h avant Claude Opus 5.5 le même jour) |
| Statut | Disponible — API, ChatGPT Work, Codex (plans payants) ; remplace GPT-5.6 Sol |
| Architecture | **non vérifié au 27/09/2026** |
| Contexte | 1,05 M tokens, 128K max output (sources tierces : pages modèles OpenAI) |
| Features | text + image in, text out ; reasoning, tool use, agents (coding, agent tasks) ; prompt caching amélioré (90 % de remise cached input, meilleurs hit rates) ; Responses API, Chat Completions, structured outputs, streaming, Batch |
| Fine-tuning | **NON supporté** (assistants/realtime/live/fine-tuning listés comme non supportés) |
| Poids | fermés |
| Benchmarks clés (OpenAI-déclarés) | AutomationBench (bat Claude Opus 5 à 9 % du coût) ; DeepSWE v1.1 : **68,8 %** (vs Claude Fable 5 : 69,9 %) à ~80 % moins cher par tâche ; AA Coding Agent Index : 55 → 57 ; AA Omniscience Test : fabrication de faits 92 % → 60 % |
| Prix API | **$2 in / $10 out** par MTok (50 % moins que GPT-5.6 Sol) |
| Déploiement | API OpenAI, ChatGPT Work, Codex ; rollout ChatGPT standard en cours |

## 7. GPT-6 Luna — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-6 Luna / `gpt-6-luna` |
| Sortie | **22 septembre 2026** |
| Statut | Disponible ; remplace GPT-5.6 Luna ; **accessible aussi aux utilisateurs Free via l'app desktop** (limite initiale) |
| Architecture | **non vérifié au 27/09/2026** |
| Contexte | 1,05 M tokens, 128K max output (pages modèles OpenAI, sources tierces) |
| Features | high-volume, tâches répétitives ; tool use, agents légers ; même stack que Sol optimisée coût/latence ; AA Coding Agent Index : 41 (−2 pts vs 5.6 Luna) ; AA Omniscience : fabrication 93 % → 77 % |
| Fine-tuning | **NON supporté** (mêmes exclusions que Sol) |
| Poids | fermés |
| Prix API | **$0,10 in / $0,50 out** par MTok (50 % moins que 5.6 Luna ; ~25× moins cher que le tier Sol le plus cher) |
| Déploiement | API, ChatGPT Work, Codex, app desktop gratuite |

## 8. GPT-6 Terra : non sorti au 27/09/2026 ❌

La série GPT-6 lancée le 22 septembre 2026 ne contient que **Sol et Luna**. **Pas de GPT-6 Terra sorti** au 27/09/2026 ; l'existence d'un futur GPT-6 Terra est **non vérifiée au 27/09/2026** (« No GPT-6 version shipped; unclear if one is coming » — explainx.ai). Le tier Terra reste donc, fin septembre 2026, exclusivement GPT-5.6 Terra.

## 9. « GPT Pro » : PAS un modèle standalone ❌ (terme ambigu — ce qu'il recouvre vraiment)

Après trois angles de recherche (annonces de modèles, tiers d'abonnement, docs API), conclusion : **il n'existe aucun modèle standalone appelé « GPT Pro »**. Le terme recouvre trois réalités distinctes :

1. **ChatGPT Pro = OFFRE d'abonnement (pas un modèle)** — deux tiers en sept. 2026 :
   - ChatGPT Pro **$100/mois (« Pro 5x »)**, lancé le **9 avril 2026** (5× les limites Plus) ;
   - ChatGPT Pro **$200/mois (« Pro 20x »)** — **nouvelles souscriptions et upgrades SUSPENDUS les 10–11 septembre 2026** car la demande GPT-6 Astra dépassait la capacité (Thibault Sottiaux, OpenAI ; existants non affectés).
   - Tiers complets : Free / Go $8 / Plus $20 / Pro $100 / Pro $200 / Business / Enterprise.
2. **Variantes « -Pro » des modèles** : GPT-5.4 Pro, GPT-5.5 Pro, GPT-5.6 Pro (`reasoning.mode: "pro"`), GPT-6 Astra Pro, historiquement GPT-5 Pro (août 2025). Ce sont des variantes haute-précision de modèles existants, pas un produit « GPT Pro ».
3. **Rumeur non confirmée (sept. 2026) : « ChatGPT Pro Max » $500/mois** (TestingCatalog via X, 24/09/2026) — **non vérifié au 27/09/2026**, lié au DevDay OpenAI.

## 10. GPT-5.3 — la ligne coding (pas de flagship général) ✅

| Champ | Valeur vérifiée |
|---|---|
| Noms exacts | **GPT-5.3-Codex** (`gpt-5.3-codex`), **GPT-5.3-Codex-Spark**, **GPT-5.3 Instant** |
| Sortie | **5 février 2026** (GPT-5.3-Codex) ; 9 fév. 2026 GA dans GitHub Copilot ; **12 fév. 2026** (GPT-5.3-Codex-Spark, research preview) ; **3 mars 2026** (GPT-5.3 Instant) |
| Statut | disponible ; **GPT-5.3-Codex désigné LTS le 18 mars 2026** (garanti jusqu'au 4 fév. 2027 pour Business/Enterprise Copilot) — **premier modèle LTS OpenAI** |
| Architecture | **non vérifié au 27/09/2026** — « premier modèle combinant les training stacks Codex + GPT-5 » |
| Contexte | Spark : 128K (déclaré par OpenAI) ; Codex : non vérifié précisément |
| Features | agentic coding (génération → agent de codage généraliste steerable), computer use ; Spark = variante petite/rapide « pair programmer » temps réel |
| Poids | fermés |
| Benchmarks clés (déclarés) | SWE-bench Pro 56,8 % ; Terminal-Bench 2.0 77,3 % ; ~25 % plus rapide sur tâches agentiques (Codex) |
| Prix API | Codex : $1,75 / $14 (source unique epic-skills — **non vérifié au 27/09/2026**) ; Instant : **~$0,30 / $1,20** par MTok |
| Déploiement | Codex (Codex-only au lancement), GitHub Copilot, API (Responses) |
| **IMPORTANT** | aucune trace d'un « GPT-5.3 » flagship général standalone — la ligne 5.3 est coding (Codex) + Instant. Un « GPT-5.3 » tout court sans suffixe = **non vérifié comme produit distinct au 27/09/2026**. |

## 11. GPT-5.3-Codex-Spark : le « pair programmer » temps réel ✅

Research preview lancée le 12 février 2026. Variante petite et rapide de GPT-5.3-Codex, pensée comme copilote temps réel : latence basse, contexte 128K déclaré par OpenAI. Positionnement : là où le Codex standard réfléchit longuement, Spark répond vite. Prix : **non vérifié au 27/09/2026**. Poids fermés, disponible dans Codex puis Copilot.

## 12. GPT-5.3 Instant : le tier pas cher ✅

Sorti le 3 mars 2026. Le modèle d'usage quotidien basse latence de la génération 5.3 : prix ~$0,30 / $1,20 par MTok. Le blogueur cité par la recherche revendique « 70 % d'hallucinations en moins » pour la génération Instant en général — **non vérifié au 27/09/2026**. Poids fermés.

## 13. GPT-5.4 — fiche ✅

| Champ | Valeur vérifiée |
|---|---|
| Nom exact / ID API | GPT-5.4 / `gpt-5.4` (variantes : GPT-5.4 Pro, mini, nano, -Cyber) |
| Sortie | **5 mars 2026** (GPT-5.4 + Pro) ; **17 mars 2026** (mini + nano) ; 14 avril 2026 (GPT-5.4-Cyber, Trusted Access for Cyber) |
| Statut | disponible API ; **retiré de Codex via login ChatGPT le 31 août 2026** (redirection vers GPT-5.6 Terra/Luna) |
| Architecture | **non vérifié au 27/09/2026** |
| Contexte | **1 050 000 tokens**, **128K max output** ; knowledge cutoff : **31 août 2025** |
| Features | **premier modèle généraliste avec computer use natif** (Playwright, souris/clavier sur screenshots) ; intègre le coding de GPT-5.3-Codex + raisonnement avancé + workflows agentiques ; GPT-5.4 Thinking (toggle ChatGPT) ; ChatGPT for Excel/Sheets (bêta, lancé avec) |
| Fine-tuning | **non vérifié au 27/09/2026** |
| Poids | fermés |
| Benchmarks clés (déclarés) | SWE-bench Pro 57,7 % ; SWE-bench Verified ~80 % ; OSWorld 75 % (vs 47,3 % pour 5.2 ; humain 72,4 %) ; WebArena-Verified 67,3 % ; Online-Mind2Web 92,8 % (screenshot) ; hallucinations individuelles −33 %, taux d'erreur full-response −18 %, usage tokens −47 % sur certaines tâches |
| Prix API | GPT-5.4 : **$2,50 / $15** ; mini : $0,75 / $4,50 ; nano : $0,20 / $1,25 ; **Pro : $30 / $180** (Responses API uniquement) — batch/flex à 50 % |
| Déploiement | ChatGPT (Plus/Team/Pro), Codex (jusqu'au 31/08/2026), API OpenAI |

## 14. Les variantes GPT-5.4 : Pro, mini, nano ✅

