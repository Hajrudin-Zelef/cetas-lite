---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-29
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Cerebras", "DeepSeek", "EU", "Fireworks AI", "Glasswing", "Google", "Groq", "Hugging Face", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nebius", "Nvidia", "OpenAI", "OpenRouter", "Poolside", "Together AI", "United States", "Xiaomi", "Z.ai", "xAI"]
dates: ["2026-05-01", "2026-09-27"]
keywords: ["agent", "agents", "astra", "aws", "bedrock", "benchmark", "chatgpt", "claude", "copilot", "cyber", "deepseek", "fable 5"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [2233, 2347]
sha256: 0160210d211defcf06f440cd2673b9cb401ef72472160265a6633ccc3c39f226
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Modèle | Contexte annoncé | Niveau de preuve au 27/09/2026 |
|---|---|---|
| GPT-5.6 / GPT-6 | 1,05M | **vérifié** — chiffre officiel répété sur toute la famille |
| Muse Spark 1.1+ | 1M | **vérifié** — pricing officiel à 1M |
| Claude Fable 5/5.1 | 1M | **vérifié** — pricing + docs |
| Claude Opus 4.6 | 1M (bêta) | bêta via header ; sortie 128K standard |
| Gemini 3.x Flash | 1M entrée / 64K sortie | **vérifié** — constant sur toute la série |
| Grok 4.3 | 1M | revendiqué (lancement 01/05/2026) |
| Laguna S 2.1 | 1M | **vérifié** — entraîné 131K→1M (RoPE 1M), NVFP4 1M |
| Nemotron 3 (Nano/Super/Ultra/Lightning) | 1M | **vérifié** — constant famille ; Nano : fiche HF dit 262K (divergence § 165-17) |
| Llama 4 Maverick | 1M | revendiqué ; chiffres sous réserve LeCun (§ 93) |
| Llama 4 Scout | **10M** | needle-in-haystack démontré par Meta |
| Mistral Medium 3.5 / Small 4 | 256K | **vérifié** |
| Gemma 4 | 256K (128K E2B/E4B) | **vérifié** |
| Granite 4.2 | 512K (cible) | objectif d'entraînement, pas une mesure |
| Hermes 4.3-36B | 128K–512K ? | **divergence non expliquée** (§ 165-15) |

Leçon : « 1M de contexte » est devenu un standard marketing en 2026, mais la preuve varie du benchmark public (Scout, Laguna) à la simple fiche produit (Grok 4.3). Pour du long-contexte critique, exiger le test needle-in-haystack **sur vos documents**.

## 170. Multilinguisme : qui parle quoi

| Modèle | Couverture vérifiée |
|---|---|
| Gemma 4 | **140+ langues** (le plus large de ce volume) |
| Hermes 4 | **40 langues** |
| Nemotron 3 Nano/Super | **38 langues** |
| Grok 4.3 Custom Voices | 28 langues (voix) |
| Mistral Medium 3.5 | dizaines de langues ; **80+ langages de programmation** (FIM) |
| Mistral Small 4 | dizaines de langues ; **80+ langages** (FIM) |
| Voxtral TTS | 9 langues (clonage vocal) |
| Poolside Laguna | 20 langues naturelles + 100 langages de code |
| Granite 4.2 | multilingue documenté (détail non vérifié) |
| Grok 4.6 | multilingue (détail non vérifié) |

Pour un RAG francophone : Mistral (français natif de l'éditeur), Gemma 4 (140+ langues), Hermes 4. Dans tous les cas, la qualité réelle sur corpus français technique doit être mesurée — aucun benchmark public 2026 ne la documente dans ce volume.

## 171. Tool calling et protocoles agents — matrice de support

| Modèle / harness | Function/tool calling | MCP | A2A / ACP | Format |
|---|---|---|---|---|
| Claude Opus 4.6+ / 5.x / Fable | ✅ (natif) | ✅ natif (Opus 5) | — | Anthropic |
| GPT-5.x / 6.x | ✅ (natif) | ✅ (Astra) | — | OpenAI Responses |
| Gemini 3.x Flash | ✅ (natif) | — | — | Gemini API |
| Grok 4.6 | ✅ (natif) | ✅ | ✅ (A2A) | xAI |
| Grok 4.1 Fast | ✅ (amélioré) | ✅ | — | xAI |
| Mistral Medium 3.5 / Small 4 | ✅ (natif) | — | — | Mistral |
| Gemma 4 (26B-A4B/31B) | ✅ (natif, JSON strict) | — | — | ouvert |
| Granite 4.2 | ✅ (natif) | — | — | **OpenAI-compatible** |
| Nemotron 3/3.5 | ✅ (enterprise tool calling) | — | — | ouvert |
| Laguna S 2.1 | ✅ (trajectoires publiées) | — | ✅ (ACP via `pool`) | ouvert |
| Hermes Agent | ✅ (model-agnostic) | — | — | MIT |
| Muse Spark 1.1+ | ✅ (subagent delegation, goal conditioning) | — | — | Meta Model API |

**MCP** (Model Context Protocol, protocole ouvert d'Anthropic) est le standard de facto 2026 pour brancher outils et sources : supporté nativement par Opus 5, Astra, Grok 4.1 Fast/4.6. Vérifier la version du protocole supportée avant d'investir dans un connecteur.

## 172. API et plateformes managées par famille

| Famille | API / plateformes vérifiées |
|---|---|
| OpenAI GPT | OpenAI API (Responses), ChatGPT, Azure OpenAI, AWS Bedrock (Astra), GitHub Copilot, Codex |
| Anthropic Claude | Anthropic API, AWS Bedrock (**Zero Data Retention**), Google Vertex AI, Azure AI Foundry, GitHub Copilot |
| Google Gemini | Gemini API, AI Studio, Google Vertex AI, Gemini app, Firebase |
| xAI Grok | xAI API, xAI Console, X, Grok app, SuperGrok, OpenRouter, AWS Bedrock (Grok 4.6), Cursor, GitHub Copilot |
| Meta Muse Spark | **Meta Model API** (depuis Spark 1.1), agent Muse (US) |
| Meta Llama 4 | Hugging Face, Together AI, Fireworks AI, Groq, Cerebras, Nebius, Snowflake, Databricks, Azure, Bedrock |
| Mistral | La Plateforme (API), Le Chat / Vibe, Hugging Face, partenaires cloud |
| Google Gemma | **AI Studio + Gemini API (gratuit)**, Hugging Face, Ollama, Kaggle |
| IBM Granite | watsonx, **Granite AI Studio (NLX, gratuit)**, Hugging Face, Ollama |
| Poolside | **Poolside API** (XS.2/M.1 free tiers fermés juil. 2026), OpenRouter, Hugging Face |
| NVIDIA Nemotron | **NVIDIA NIM**, **Token Factory** (pay-per-token), **Build.nvidia.com**, Hugging Face |
| NousResearch Hermes | OpenRouter, Hugging Face, **Nous API**, endpoints tiers |

## 173. Disponibilité cloud : Bedrock, Vertex, Azure — qui est où

| Modèle | AWS Bedrock | Google Vertex | Azure |
|---|---|---|---|
| GPT-6 Astra | ✅ | — | ✅ (OpenAI + GitHub) |
| Claude Fable 5/5.1, Opus 4.6→5.5, Sonnet 5, Haiku 4.5 | ✅ (Zero Data Retention) | ✅ | ✅ (AI Foundry) |
| Gemini 3.x | — | ✅ | — |
| Grok 4.6 | ✅ | — | — |
| Llama 4 Scout/Maverick | ✅ | — | ✅ |
| Mistral (Small 4, Medium 3.5…) | partenaires | — | ✅ (sélection) |
| Gemma 4 | — | ✅ (AI Studio / Gemini API) | — |
| Granite 4.x | — | — | — (watsonx/IBM Cloud) |

Pour les déploiements soumis à des contraintes cloud (contrats existants, régions), ce tableau prime sur les préférences de benchmark : un excellent modèle indisponible sur votre cloud est un mauvais choix.

## 174. Ce que couvre le volume Chine (pour ne pas chercher ici)

Les familles suivantes sont documentées dans le **volume Chine** de l'encyclopédie, pas dans celui-ci :
DeepSeek (V4 et dérivés), Qwen (série 3.x), GLM (Zhipu), Kimi (Moonshot AI — K2, K3, agent teams), MiniMax (M1/M2, H3), MiMo (Xiaomi), HY3 (Tencent Hunyuan). Les ponts avec ce volume sont signalés explicitement : Hermes Agent supporte les endpoints Qwen/GLM/Kimi/MiniMax (§ 135), et § 151 note le trou de l'open-weight occidental en vidéo face à MiniMax-H3.

## 175. Checklist « avant de choisir » — 20 questions

1. Le modèle est-il **disponible** (pas retiré/EOL/annoncé) à la date de mise en prod ? (§ 156)
2. Quelle est la politique de **dépréciation** écrite du fournisseur ? (§ 157)
3. Le prix affiché inclut-il le **tokenizer** réel (tokens/requête mesurés) ? (§ 158)
4. Le **prompt caching** est-il actif et au-delà de quel seuil ? (§ 158)
5. Y a-t-il un **palier long-contexte** qui double la facture ? (§ 158)
6. L'**effort de raisonnement** est-il réglable (low→max) ? (§ 139)
7. Le benchmark cité est-il **vendor-déclaré ou indépendant**, avec quel harness ? (§ 155)
8. La **licence** autorise-t-elle mon usage commercial (seuils, carve-out) ? (§ 160)
9. Les **poids** sont-ils nécessaires (self-host, audit) ou l'API suffit-elle ? (§ 145)
10. Le modèle est-il dispo sur **mon cloud** (Bedrock/Vertex/Azure) ? (§ 173)
11. La **résidence des données** est-elle garantie (EU si besoin) ? (§ 152)
12. Le **Zero Data Retention** est-il possible ? (§ 152)
13. Le **MCP** est-il supporté pour mes connecteurs ? (§ 171)
14. Le **computer use** est-il requis, et sur quel OS ? (§ 144)
15. L'usage est-il **cyber/offensif** → quel programme d'accès (Glasswing/Daybreak/Fairwind) ? (§ 159)
16. Le **watermarking** (Fable 5.1, SynthID) pose-t-il problème pour mon cas d'usage ? (§ 152)
17. Le modèle **bascule-t-il** sur les requêtes sensibles de mon domaine ? (garde-fous, § 51)
18. Quel est le **coût par tâche** mesuré, pas par token ? (§ 140)
19. Existe-t-il une **version LTS** pour figer la stack ? (§ 157 — seul GPT-5.3-Codex)
20. Ai-je vérifié le **pricing le jour de la signature** (baisses trimestrielles) ? (§ 140)

