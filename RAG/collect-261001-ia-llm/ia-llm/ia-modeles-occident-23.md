---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-23
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Glasswing", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "Poolside", "United States", "xAI"]
dates: ["2026-04-02", "2026-07-15", "2026-09-24", "2026-09-27", "2027-02-04"]
keywords: ["agent", "agents", "apache", "astra", "attention", "bedrock", "benchmarks", "claude", "consumer", "cyber", "fable 5", "gemini"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1790, 1882]
sha256: 73174f04123045691be16fc512ddfe40380df699730690481df026cef4b097af
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

**Q11.** Qu'est-ce que « Veo 4 » ?
**R11.** Il n'existe pas : Google a lancé Gemini Omni à l'I/O 2026 au lieu d'un Veo 4. La version vidéo officielle actuelle est Veo 3.1.

**Q12.** Quelles sont les trois variantes cyber restreintes des labs, et leurs programmes ?
**R12.** Claude Mythos (Project Glasswing, Anthropic), GPT-5.x-Cyber (Daybreak, OpenAI), Gemini Flash Cyber (Fairwind Program, Google).

**Q13.** Quelle est la rupture de licence de Gemma 4 ?
**R13.** Gemma 4 (02/04/2026) passe en Apache 2.0, abandonnant les Gemma Terms of Use custom des générations 1–3.

**Q14.** Citez deux modèles open-weight à moins de $0,10/M tokens d'entrée et leur usage typique.
**R14.** Nemotron 3.5 Lightning 30B-A3B ($0,06/$0,24, agents always-on) et Laguna S 2.1 (~$0,09/$0,18, code agentique). GPT-6 Luna ($0,10/$0,50) est fermé mais au même plancher.

**Q15.** Que s'est-il passé le 22 septembre 2026 ?
**R15.** Double lancement et guerre des prix : OpenAI sort GPT-6 Sol ($2/$10) et Luna ($0,10/$0,50) ~2 h avant qu'Anthropic ne lance Claude Opus 5.5 ($4/$20, −20 %) avec System Card publiée.

---

# PARTIE XIII — Guide pratique : choisir son modèle (sept. 2026)

> Ces sections croisent les fiches pour répondre à « lequel pour quel usage ». Aucune donnée nouvelle inventée : tout renvoie aux fiches et tableaux précédents.

## 143. Choisir pour le code agentique

Le code agentique est le segment le plus disputé de 2026. Repères vérifiés (SWE-bench Verified / Pro, revendications fournisseurs) :

| Modèle | SWE-bench Verified | SWE-bench Pro | Prix in/out par MTok | Poids | Verdict usage |
|---|---|---|---|---|---|
| Claude Fable 5 | 95,0 % | 80,3 % | $10/$50 | fermés | le meilleur public, prix frontier |
| Claude Opus 5.5 | — | 89,9 % | $4/$20 | fermés | meilleur rapport capacité/prix fermé |
| Claude Opus 4.8 | 88,6 % | 69,2 % | $5/$25 | fermés | solide, fallback garde-fous |
| GPT-6 Sol | — | — (DeepSWE 68,8 %) | $2/$10 | fermés | coding/agents à prix cassé |
| GPT-5.3-Codex | — | 56,8 % | $1,75/$14 ? | fermés | LTS jusqu'au 04/02/2027 — valeur sûre entreprise |
| GPT-5.4 | ~80 % | 57,7 % | $2,50/$15 | fermés | computer use + code |
| Mistral Medium 3.5 | 77,6 % | — | $1,50/$7,50 ? | **ouverts (Mod. MIT)** | meilleur code ouvert européen |
| Laguna S 2.1 | — (multilingue 78,5 %) | 59,4 % | ~$0,09/$0,18 | **ouverts (OpenMDW-1.1)** | code ouvert le moins cher |
| Gemini 3.8 Flash | — (DeepSWE 73,7 %/71,0 %) | — | $0,75/$3,75 promo | fermés | coding rapide pas cher |
| Grok 4.6 | — (DeepSWE 65,9 %) | — | $2/$6 | fermés | alternative Cursor/Bedrock |

Règles de décision :
- **Budget illimité, résultat max** : Fable 5.1 ou Opus 5.5 (Claude Code en défaut Opus sur les plans Max/Pro).
- **Entreprise, stabilité contractuelle** : GPT-5.3-Codex **LTS** (garanti jusqu'en fév. 2027) — le seul modèle du marché avec un engagement de longévité explicite.
- **Self-host / souveraineté** : Laguna S 2.1 (spécialiste code, 1M contexte) ou Mistral Medium 3.5 (le plus polyvalent des ouverts).
- **Volume / CI** : Gemini 3.8 Flash ou GPT-6 Luna selon l'écosystème déjà en place.
- Piège : les scores SWE-bench sont mesurés avec des harness différents (Anthropic, OpenAI, Poolside) — ne comparez jamais deux chiffres sans vérifier le harness.

## 144. Choisir pour les agents autonomes long-horizon

| Besoin | Candidats vérifiés | Points d'attention |
|---|---|---|
| Sessions de plusieurs heures, auto-vérification | Grok 4.7 (RL pondéré multi-heures, harness Grok Bot natif), Nemotron 3 Ultra 550B-A55B (planification multi-turn, subagents) | Grok 4.7 : prix inchangé vs 4.6 malgré base plus large |
| Centaines de subagents parallèles | Claude Opus 4.8 (Dynamic Workflows, research preview), GPT-5.6 Sol (mode `ultra`) | Dynamic Workflows encore en research preview |
| Agent de bureau (computer use) | GPT-6 Astra (1,9× plus rapide que Sol sur Mind2Web), Claude Opus 4.8 (OSWorld 83,4 %), Muse (Mac desktop control, 24/09/2026) | trois écosystèmes incompatibles : choisir selon l'OS cible |
| Harness open source auto-évolutif | Hermes Agent (MIT, GEPA, fév. 2026), agent `pool` + Shimmer (Poolside), Grok Build (source-available, 15/07/2026) | Hermes Agent : 40–118 skills selon version ; vérifier la version |
| Orchestration multi-agents managée | Muse Spark 1.1+ (subagent delegation, goal conditioning), Muse (agent produit, Secure VM) | Muse : US only au lancement |

Le point aveugle de 2026 : **l'anti-injection**. Un agent qui browse est un agent attaquable. Seul Muse documente publiquement une architecture dédiée (Sentinel, Secure VM, 5 couches, bounty $300K). Pour les autres, les mitigations existent (garde-fous Claude, Frontier Safety Framework Google) mais le design n'est pas public au même niveau.

## 145. Choisir pour le RAG et le self-hosting local

Les flagships fermés sont exclus par construction (pas de poids). Le podium ouvert occidental au 27/09/2026 :

| Modèle | Licence | Contexte | Atout RAG/local | Limite |
|---|---|---|---|---|
| Gemma 4 31B / 26B-A4B | **Apache 2.0** | 256K | function calling natif, JSON strict, NVFP4 officiel, #3/#6 ouverts Arena | pas d'API payante (AI Studio/Gemini API) |
| Mistral Medium 3.5 | Modified MIT | 256K | dense 128B, toggle reasoning, FIM, OCR | carve-out gros revenus ; prix API à confirmer |
| Mistral Small 4 | Apache 2.0 | 256K | MoE 119B/6B actifs, unifié, NVFP4 officiel | prix API non vérifié |
| Granite 4.2 | Apache 2.0 | 512K (cible) | tool calling OpenAI-format, poids signés, doc 5 phases | benchmarks indépendants absents |
| Nemotron 3.5 Lightning | OpenMDW-1.1 | 1M | 22 Go en NVFP4, MTP, recettes NeMo complètes | licence NVIDIA à lire par repo |
| Laguna S 2.1 | OpenMDW-1.1 | 1M | 71 Go NVFP4, trajectoires d'éval publiées | spécialiste code, moins généraliste |
| Hermes 4.3-36B | permissive | 128K–512K ? | consumer hardware, steerability | divergence 128K/512K non expliquée |
| Llama 4 Scout | Llama Community (custom) | **10M** | needle-in-haystack démontré à 10M | licence custom, plus de suivi Meta |

Recommandation par profil :
- **PME / labo, un seul GPU** : Gemma 4 26B-A4B quantifié (24 Go) ou Nemotron 3.5 Lightning NVFP4 (22 Go).
- **Entreprise régulée** : Granite 4.2 (poids signés, traçabilité) — le seul à documenter 5 phases de pré-training.
- **Gros corpus (10M+ tokens)** : Llama 4 Scout reste le seul ouvert occidental démontré à 10M — mais sans suivi Meta depuis le pivot Muse.
- **RAG multilingue FR** : Mistral (FR natif) ou Gemma 4 (140+ langues).

## 146. Choisir pour le multimodal (vision, audio, vidéo)

| Modalité | Options vérifiées | Notes |
|---|---|---|
| Vision + raisonnement (entrée image) | GPT-6 Astra, Claude Fable 5/5.1, Gemini 3.8 Flash, Gemma 4 (ouvert), Nemotron 3 Nano Omni (ouvert, texte/image/vidéo/audio) | Opus 4.7 : vision 3× résolution (2 576 px) |
| Audio natif (ASR, transcription) | Gemma 4 E2B/E4B (ouverts), Voxtral Mini Transcribe 2 (Mistral) | E2B/E4B pensés on-device |
| TTS / clonage vocal | Voxtral TTS (Mistral, zero-shot cloning), Grok 4.3 Custom Voices ($4,20/M caractères) | IDs Gemini 3.8 TTS non vérifiés au-delà du guide de migration |
| Génération d'images | Nano Banana 2 (workhorse, $0,067/1K), Nano Banana Pro ($0,134, Adobe), Nano Banana 2 Lite ($0,0336), Grok Imagine 2.0 (Arena #2) | SynthID sur toutes les images Google |
| Génération vidéo | Veo 3.1 (audio natif, 9:16, 148 s extension), Gemini Omni Flash (max 10 s/clip, édition conversationnelle), Grok Imagine (vidéo) | API entreprise Omni non lancée au 27/09/2026 |
| Vision + action (robotique) | Robostral Navigate (8B VLA, single-camera RGB) | annoncé uniquement, pas d'API publique |

## 147. Choisir pour le budget : le plancher des prix

Prix par MTok (in/out), offres les moins chères vérifiées au 27/09/2026 :

