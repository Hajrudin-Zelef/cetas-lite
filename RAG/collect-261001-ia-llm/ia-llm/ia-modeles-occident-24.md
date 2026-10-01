---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-24
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "EU", "Glasswing", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "United States", "xAI"]
dates: ["2026-02-04", "2026-03-23", "2026-09-24", "2026-09-27", "2026-12-31"]
keywords: ["agent", "agents", "agi", "apache", "astra", "attention", "bedrock", "benchmarks", "chatgpt", "claude", "copilot", "cyber"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1883, 1947]
sha256: a90171ec17244793e112d95d0e17c9479b3ae824febb2853dd6656f2a4691c91
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Segment | Modèle | Prix in/out | Conditions |
|---|---|---|---|
| Le moins cher (fermé) | GPT-6 Luna | **$0,10/$0,50** | API, ChatGPT Work, Codex ; Free via app desktop |
| Le moins cher (ouvert, API) | Nemotron 3.5 Lightning | **$0,06/$0,24** | Token Factory |
| Code ouvert pas cher | Laguna S 2.1 | ~$0,09/$0,18 | OpenRouter (−10 %) |
| Flash Google | Gemini 3.8 Flash | $0,75/$3,75 | promo jusqu'au 31/12/2026 |
| Ultra-low via consentement data | Muse Spark 1.1+ « Contributor » | **$0,10/$0,20** | opt-in : requêtes utilisées pour l'entraînement |
| Gratuit (limité) | Muse (agent) | gratuit ≤100M tok/sem | US only ; puis $20/$100/mois |

Stratégies d'économie vérifiées :
1. **Cache** : 90 % de remise en lecture (OpenAI, Anthropic standard) ; **97,5 % chez Fable 5.1** ($0,25/M).
2. **Batch** : −50 % (OpenAI Batch/Flex, Anthropic batch).
3. **Effort réglable** : baisser le thinking (low/minimal) sur les tâches simples — tous les labs le proposent en 2026.
4. **Routage** : Haiku 4.5 ($1/$5) ou Flash-Lite ($0,25–$0,30/$1,50–$2,50) en première passe, escalade vers le flagship si besoin.
5. **Long-contexte** : attention aux paliers — OpenAI ×2/×1,5 au-delà de 272K ; xAI ×2 au-delà de 200K (Grok 4.6).

## 148. Choisir pour le raisonnement frontier

| Modèle | Signaux vérifiés | Prix | Accès |
|---|---|---|---|
| GPT-6 Astra | FrontierMath Tier 4 : 98 % ; ARC-AGI-3 : 99,9 % (harness OpenAI) / 62,7 % (standard) ; Agents' Last Exam 59,3 % | $10/$50 | API, Azure, Bedrock, Copilot |
| Claude Fable 5.1 | Terminal-Bench-Science 52,6 % ; Vals 68,83 % ; AA Index v4.3 ex æquo 1er (53) | $10/$50 | API, Bedrock, Vertex, Azure |
| Claude Opus 5.5 | SWE-bench Pro 89,9 % ; Terminal-Bench 4.0 66,4 % ; AA v4.3.2 1er (57,6) | $4/$20 | toutes plateformes |
| Muse Spark 1.3 | DeepSWE 75,4 % ; Terminal-Bench 2.1 88,8 % ; AA 62 | $1,25/$4,25 | Meta Model API, agent Muse |
| Nemotron-Cascade-2 | IMO 2025 + IOI 2025 gold (revendiqué NVIDIA) | non vérifié | poids ouverts |
| Gemini 3 Deep Think | raisonnement parallèle profond, maths/sciences | non vérifié | AI Ultra + early access entreprise |

Note de méthode : les benchmarks « frontier » 2026 sont presque tous **vendor-déclarés** et mesurés avec des harness maison (l'écart ARC-AGI-3 d'Astra : 99,9 % vs 62,7 % selon le harness, l'illustre). Pour un choix d'entreprise, exiger une évaluation sur **vos** tâches avec **votre** harness — les classements publics (AA Index, LMArena) sont des proxys, pas des preuves.

## 149. Cybersécurité : qui fait quoi (et qui est restreint)

| Programme | Lab | Modèles concernés | Accès | Statut |
|---|---|---|---|---|
| Project Glasswing | Anthropic | Mythos, Mythos 5, Mythos 5.1 | ~40 orgs vérifiées, NDA, $100M crédits + $4M dons | actif |
| Daybreak / Daybreak Red | OpenAI | GPT-5.4-Cyber, 5.5-Cyber, 5.6-Cyber | trusted access, contrôlé | actif |
| Fairwind Program | Google | Gemini 3.5 / 3.8 Flash Cyber | gouvernements, infra critiques, mainteneurs | actif |
| Preparedness « Critical » | OpenAI | GPT-6 Astra | capacités offensives en accès contrôlé | premier « Critical » de l'histoire |

Lecture : en 2026, la cybersécurité offensive est devenue **le** critère de classement frontier — et simultanément le motif de non-publication (Mythos). Pour un RSSI, la question n'est plus « quel modèle détecte le mieux » mais « quel programme d'accès correspond à mon niveau d'habilitation ». Les variantes publiques (Fable 5/5.1, Astra standard) **basculent ou refusent** les requêtes offensives ; seules les variantes restreintes les traitent.

## 150. Voix, transcription et TTS : l'état au 27/09/2026

- **Transcription** : Voxtral Mini Transcribe 2 + Realtime (Mistral, 04/02/2026) ; ASR natif sur Gemma 4 E2B/E4B (ouverts, on-device). Specs détaillées Voxtral : non vérifiées.
- **TTS / clonage** : Voxtral TTS (Mistral, 23/03/2026, zero-shot voice cloning, 9 langues — licence exacte non vérifiée) ; Grok 4.3 Custom Voices (1 min d'enregistrement → clone en <2 min, $4,20/M caractères, $3/h realtime, 80+ voix, 28 langues).
- **Gemini TTS** : IDs `gemini-3.8-flash-tts` / `gemini-3.8-flash-lite-tts` mentionnés uniquement dans un guide de migration — existence non vérifiée au-delà.
- **Muse** : voix custom en fuite (timeline voix, tech-insider) + Realtime Avatar (Connect 24/09/2026) — produit, pas modèle API.
- Point de vigilance : le clonage vocal zero-shot (Voxtral, Grok) pose des questions de consentement ; aucun des labs ne documente publiquement de garde-fou anti-impersonation au 27/09/2026.

## 151. Image et vidéo génératives : qui mène

- **Image** : Google domine en profondeur de gamme (Nano Banana 2 Lite $0,0336 → Pro $0,134, intégré Adobe). xAI (Imagine 2.0, moteur Aurora) est #2 des arènes (1320 text-to-image, 1439 editing) derrière GPT-Image-2 d'OpenAI. Toutes les images Google portent **SynthID**.
- **Vidéo** : Veo 3.1 reste la référence officielle (audio natif synchronisé, 9:16 natif, extension 148 s) ; Gemini Omni Flash change le paradigme (édition conversationnelle post-génération, max 10 s/clip) mais son API entreprise n'est toujours pas lancée 4 mois après l'I/O.
- **Modèles ouverts occidentaux image/vidéo** : quasi absents de ce volume en 2026 — le seul omni-modal ouvert notable est MiniMax-H3 (volume Chine). C'est un trou béant de l'écosystème open-weight occidental.

## 152. Conformité et résidence des données : points de vigilance

- **Zero Data Retention** : défaut sur Bedrock pour Claude Opus 5.5 ; « Enterprise Frontier Safeguards » (données sur infra client) en déploiement phasé automne 2026 (Fable 5/5.1).
- **Gemini Omni** : pas de région EU sélectionnable / résidence EU garantie via Gemini API au 27/09/2026 — bloquant pour certains cas d'usage européens.
- **Muse** : US only au lancement ; credentials dans `hatch-authd` (jamais lisibles par l'agent) — le design le plus auditable côté gestion des secrets.
- **Poids signés** : Granite 4.2 (IBM) — seul modèle à signer cryptographiquement ses poids, argument pour les secteurs régulés.
- **Licences** : Apache 2.0 (Gemma 4, Granite, Mistral Small 4, Laguna XS.2) = usage commercial sans friction ; Modified MIT (Mistral Medium 3.5, Kimi — volume Chine) = carve-out au-delà d'un seuil de revenus ; Llama Community License = custom, restrictions d'usage ; licences NVIDIA = à vérifier **par repo** (trois licences différentes dans la même famille).
- **Watermarking** : Fable 5.1 watermarke statistiquement le texte généré ; SynthID sur images/vidéos Google — la traçabilité devient native.

## 153. Index des IDs API (référence rapide)

