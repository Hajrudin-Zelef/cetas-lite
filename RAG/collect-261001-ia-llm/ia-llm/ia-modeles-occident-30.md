---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-30
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Glasswing", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai", "xAI"]
dates: ["2026-07-30", "2026-09-27", "2027-02-04"]
keywords: ["agent", "agi", "apache", "astra", "chatgpt", "claude", "cyber", "deepseek", "fable 5", "gemini", "gemini 3.8", "gemini 4"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [2348, 2467]
sha256: 1651547cea6d3336661435efbdcd666bba4b7a2be1d1b483a7a06ca03a577a2e
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

## 176. FAQ rapide — 10 réponses en une phrase

1. **Quel est le meilleur modèle de code public ?** Claude Fable 5.1 (95,0 % SWE-bench Verified, $10/$50) — ou Opus 5.5 pour un meilleur rapport qualité/prix.
2. **Quel est le moins cher ?** GPT-6 Luna ($0,10/$0,50) côté fermé ; Nemotron 3.5 Lightning ($0,06/$0,24) côté ouvert.
3. **Quel modèle pour tourner en local ?** Nemotron 3.5 Lightning NVFP4 (22 Go) ou Gemma 4 26B-A4B quantifié (24 Go).
4. **« GPT Pro » existe ?** Non comme modèle : c'est l'abonnement + les variantes -Pro.
5. **Mythos est-il utilisable ?** Non publiquement : accès restreint via Project Glasswing (~40 orgs, NDA).
6. **Gemini 4 existe ?** Non au 27/09/2026 — la ligne Pro est restée à 3.1 Preview.
7. **Quel modèle a le plus grand contexte ?** Llama 4 Scout (10M, démontré) côté ouvert ; 1,05M (GPT-6) / 1M (Fable, Spark, Gemini, Laguna, Nemotron) côté standard.
8. **Quelle licence open-weight est la plus sûre ?** Apache 2.0 (Gemma 4, Granite, Mistral Small 4, Laguna XS.2).
9. **Lequel pour un agent de bureau ?** GPT-6 Astra, Claude Opus 4.8 ou Muse selon l'OS cible — trois écosystèmes incompatibles.
10. **Les prix vont-ils continuer à baisser ?** Trois baisses OpenAI en deux mois cet été + réponse Anthropic le jour même : la tendance est à la baisse, mais vérifiez le pricing du fournisseur avant de signer.

---

# PARTIE XVI — Index et clôture

## 177. Index alphabétique des modèles cités

| Modèle | §§ | Modèle | §§ |
|---|---|---|---|
| Claude Fable 5 / 5.1 | 38–41 | GPT-5.4 | 7–8 |
| Claude Haiku 4.5 / 5.5 (annoncé) | 60–61 | GPT-5.4-Cyber | 19 |
| Claude Mythos / 5 / 5.1 | 37 | GPT-5.4 Instant | 10 |
| Claude Opus 4.6 → 5.5 | 42–49 | GPT-5.4 mini / nano | 9 |
| Claude Sonnet 4.5 / 4.6 / 5 / 5.5 (annoncé) | 57–59 | GPT-5.5 | 11–12 |
| Codestral 2508 | 106 | GPT-5.6 Sol / Terra / Luna | 13–15 |
| Devstral 2 (déprécié) | 109 | GPT-6 Astra | 16–17 |
| Gemma 4 12B | 115 | GPT-6 Luna | 20 |
| Gemma 4 26B-A4B | 113 | GPT-6 Sol | 18 |
| Gemma 4 31B | 112 | Grok 4.1 Fast (retiré) | 91 |
| Gemma 4 E2B / E4B | 111 | Grok 4.3 → 4.7 | 83–87 |
| Gemini 3.1 Pro Preview | 64 | Grok Build | 90 |
| Gemini 3.5 / 3.6 / 3.7 / 3.8 Flash | 65–69 | Grok Code Fast 1 (retiré) | 91 |
| Gemini 3.x Cyber (Fairwind) | 70 | Grok Imagine / Aurora | 92 |
| Gemini Flash-Lite | 71 | Hermes 4 / 4.3-36B | 133–134 |
| Gemini Omni | 80 | Hermes Agent | 135 |
| GPT-5.3-Codex / LTS | 5–6 | Laguna M.1 | 121 |
| GPT-5.3-Codex-Spark | 10 | Laguna S 2.1 | 124 |
| Granite 4.1 / 4.2 | 116–117 | Laguna XS 2.1 | 123 |
| Leanstral | 103 | Laguna XS.2 | 120 |
| Llama 4 Maverick | 94 | Magistral (déprécié) | 109 |
| Llama 4 Scout | 93 | Ministral 3 | 107 |
| Mistral Large 3 | 108 | Mistral Medium 3.5 | 99 |
| Mistral Moderation | 102 | Mistral Small 4 | 98 |
| Mistral OCR | 105 | Muse (agent) | 25–28 |
| Muse Spark 1.0 → 1.3 | 29–33 | Nano Banana 2 / Pro / Lite | 74–76 |
| Nemotron 3 Nano / Super / Ultra | 125–128 | Nemotron 3.5 Lightning | 129 |
| Nemotron 3 Nano Omni | 130 | Nemotron Content Safety | 131 |
| Nemotron-Cascade-2 | 126 | Robostral Navigate | 100 |
| Shieldstral | 104 | Veo 3.1 | 78 |
| Voxtral Mini Transcribe 2 | 101 | Voxtral TTS | 101 |

## 178. Les prix triés par coût croissant (entrée, $/MTok)

| Prix entrée | Modèle(s) | Prix sortie |
|---|---|---|
| $0,06 | Nemotron 3.5 Lightning | $0,24 |
| ~$0,09 | Laguna S 2.1 | ~$0,18 |
| $0,10 | GPT-6 Luna · Muse Spark (Contributor) | $0,50 · $0,20 |
| $0,20 | GPT-5.6 Luna | $1,20 |
| $0,25 | Gemini 3.1 Flash-Lite (preview) | $1,50 |
| $0,30 | Gemini 3.5 Flash-Lite | $2,50 |
| $0,75 | Gemini 3.6/3.7/3.8 Flash (promo) | $3,75 |
| $1,00 | Nemotron 3 Ultra | $3,00 |
| $1,25 | Muse Spark 1.1–1.3 · Grok 4.3 | $4,25 · $2,50 |
| $1,50 | Mistral Medium 3.5 (?) · Gemini 3.5 Flash (standard) | $7,50 · $9,00 |
| $1,75 | GPT-5.3-Codex (?) | $14 |
| $2,00 | GPT-6 Sol · Gemini 3.1 Pro · Grok 4.5/4.6/4.7 · GPT-5.4 | $10 · $12 · $6 · $15 |
| $2,50 | GPT-5.4 (standard) | $15 |
| $3,00 | Claude Sonnet 5 (post-31/08 ?) | $15 |
| $4,00 | GPT-5.6 Sol (promo) · Claude Opus 5.5 | $20 · $20 |
| $5,00 | GPT-5.5 · Claude Opus 4.6–5 | $30 · $25 |
| $10,00 | GPT-6 Astra · Claude Fable 5/5.1 | $50 · $50 |
| $25,00 | Claude Mythos Preview (?) | $125 |

« ? » = source unique ou non vérifié au 27/09/2026. Promos datées : re-vérifier avant engagement.

## 179. Septembre 2026 : le mois le plus dense de l'année

| Date | Lancement |
|---|---|
| 01/09 | Claude Fable 5.1 (+ Mythos 5.1 restreint) |
| 02/09 | Gemini 3.8 Flash · Muse Spark 1.3 |
| 03/09 | GPT-6 Astra (premier « Critical » cyber) |
| 08–09/09 | Agent Muse (US) |
| 10–11/09 | ChatGPT Pro $200/mois suspendu aux nouvelles souscriptions |
| 21/09 | Grok 4.7 |
| 22/09 | **GPT-6 Sol + Luna vs Claude Opus 5.5 — guerre des prix** |
| 24/09 | Meta Connect : Realtime Avatar, lunettes, Mac desktop control, voix custom (fuite) |
| 24–26/09 | Rumeurs « aeon » (GPT-6) et « ChatGPT Pro Max $500 » — non vérifiées |
| 25/08 (annonce) | IBM Granite 4.2 (fin août, effet septembre) |

## 180. Les 10 chiffres à retenir

1. **$0,06/$0,24** — Nemotron 3.5 Lightning : le plancher open-weight.
2. **$0,10/$0,50** — GPT-6 Luna : le plancher fermé.
3. **$10/$50** — le plafond : GPT-6 Astra et Claude Fable 5/5.1.
4. **1,05M** — le contexte standard des GPT-5.6/6.
5. **10M** — Llama 4 Scout : le record ouvert démontré.
6. **99,9 % vs 62,7 %** — ARC-AGI-3 d'Astra selon le harness : le chiffre ne vaut que par son protocole.
7. **−80 %** — la baisse du prix de GPT-5.6 Luna le 30/07/2026.
8. **$100M** — les crédits Glasswing d'Anthropic pour la cybersécurité défensive.
9. **40** — les organisations avec accès à Claude Mythos.
10. **04/02/2027** — fin de garantie du seul modèle LTS du marché (GPT-5.3-Codex).

## 181. Sources brutes de ce volume

Les huit dossiers de recherche vérifiée (27/09/2026) sont conservés sous `~/workspace/research_ia/` : `gpt_muse.md`, `claude.md`, `gemini_grok_poolside.md`, `llama_mistral_gemma_granite.md`, `minimax_nemotron_hermes.md` (seules les sections NVIDIA et NousResearch ont servi ici). Les familles chinoises (MiniMax, HY3/Tencent, Kimi, DeepSeek, Qwen, GLM, MiMo) sont dans le **volume Chine** de l'encyclopédie.

## 182. Historique des révisions

| Version | Date | Contenu |
|---|---|---|
| 1.0 | 27/09/2026 | Création : 182 sections, 14 parties + glossaire (40 termes) + quiz (15 Q/R). Période couverte : février → 27/09/2026. Données « au 27/09/2026 ». |

*Fin du volume 2 — Encyclopédie des modèles IA : l'Occident. 182 sections.*

## 183. Errata et corrections intégrées pendant la rédaction

