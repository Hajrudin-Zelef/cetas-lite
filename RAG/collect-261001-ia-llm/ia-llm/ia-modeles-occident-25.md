---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-25
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Z.ai", "xAI"]
dates: ["2026-09-27"]
keywords: ["agent", "agi", "astra", "benchmark", "benchmarks", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8", "glm", "gpt-5.6"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1948, 2002]
sha256: ffd8f67430727fad3caa2ebd7263cc1db9a4e50f3fb7f10e7c74494196f4e371
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| ID API | Modèle | ID API | Modèle |
|---|---|---|---|
| `gpt-5.6-sol` | GPT-5.6 Sol | `claude-fable-5` / `claude-fable-5-1` | Fable 5 / 5.1 |
| `gpt-5.6-terra` | GPT-5.6 Terra | `claude-opus-4-6` → `claude-opus-5-5` | Opus 4.6 → 5.5 |
| `gpt-5.6-luna` | GPT-5.6 Luna | `claude-sonnet-4-5` → `claude-sonnet-5` | Sonnet 4.5 → 5 |
| `gpt-6-astra` | GPT-6 Astra | `claude-haiku-4-5` | Haiku 4.5 |
| `gpt-6-sol` / `gpt-6-luna` | GPT-6 Sol / Luna | `gemini-3.1-pro-preview` | Gemini 3.1 Pro |
| `gpt-5.4` / `gpt-5.4-pro` | GPT-5.4 / Pro | `gemini-3.8-flash` | Gemini 3.8 Flash |
| `gpt-5.5` / `gpt-5.5-pro` | GPT-5.5 / Pro | `gemini-3.7-flash` / `gemini-3.6-flash` | Gemini 3.7 / 3.6 Flash |
| `gpt-5.3-codex` | GPT-5.3-Codex | `gemini-3.5-flash` (+ `-lite`) | Gemini 3.5 Flash |
| `gpt-5.3-codex-spark` | Codex-Spark | `grok-4.3` → `grok-4.7` | Grok 4.3 → 4.7 |
| `mistral-small-2603` | Mistral Small 4 | `grok-code-fast-1` | Grok Code Fast 1 (retiré) |
| `mistral-medium-3-5-26-04` | Mistral Medium 3.5 | `meta-llama/Llama-4-Scout-17B-16E` | Llama 4 Scout (HF) |
| `voxtral-mini-2602` | Voxtral Mini Transcribe 2 | `meta-llama/Llama-4-Maverick-17B-128E` | Llama 4 Maverick (HF) |
| `voxtral-tts-2603` | Voxtral TTS | `google/gemma-4-31b-it` | Gemma 4 31B (HF) |
| `mistral-moderation-2603` | Moderation 2603 | `google/gemma-4-26b-a4b-it` | Gemma 4 26B-A4B (HF) |
| `labs-leanstral-2603` | Leanstral | `ibm-granite/granite-4.2-8b-instruct` | Granite 4.2 8B (HF) |
| `codestral-2508` | Codestral 2508 | `nvidia/Nemotron-3-Super-120b-a12b` | Nemotron 3 Super (HF) |
| `devstral-2512` | Devstral 2 (déprécié) | `nvidia/NVIDIA-Nemotron-3.5-Lightning-30B-A3B` | Lightning (HF) |
| `magistral-medium-2509` | Magistral (déprécié) | `NousResearch/Hermes-4.3-36B` | Hermes 4.3-36B (HF) |
| `veo-3.1-generate-preview` | Veo 3.1 | `poolside/Laguna-S-2.1-NVFP4` | Laguna S 2.1 NVFP4 (HF) |
| `gemini-3-pro-image` | Nano Banana Pro | `mistralai/Mistral-Small-4-119B-2603` | Small 4 (HF) |
| `gemini-3.1-flash-image` | Nano Banana 2 | `mistralai/Mistral-Medium-3.5-128B` | Medium 3.5 (HF, gated) |

## 154. Méthode, sources et limites de ce volume

- **Méthode** : huit chercheurs (humains + IA) ont mené des recherches web vérifiées le 27/09/2026, avec la règle « rien d'inventé, tout ce qui n'est pas vérifiable est marqué ». Chaque fiche porte ses sources dans les fichiers de recherche bruts (`~/workspace/research_ia/`).
- **Sources primaires lues en direct** : rares — la plupart des pages officielles (openai.com, anthropic.com/news, deepmind.google) sont bloquées aux fetchers (Cloudflare, JS). Seule l'annonce Claude Opus 5 a été lue en lecture directe. Le reste repose sur des couvertures tierces (presse tech, repos GitHub de synthèse, docs miroirs) — les citations d'annonces officielles sont reprises de ces sources, pas lues en direct.
- **Qualité du corpus 2026** : une partie du web indexé 2026 est du contenu IA-généré agrégeant d'autres contenus (fermes de contenu, repos « ai-news »). Les dates y sont parfois incohérentes ; les sources officielles (blogs, docs, model cards HF) ont été privilégiées quand accessibles ; tout le reste est marqué secondaire/vendor.
- **Ce que ce volume ne fait pas** : il ne tranche pas les conflits de sources (ils sont signalés), ne prédit pas les sorties futures (Sonnet 5.5, Haiku 5.5, GPT-6 Terra, « Watermelon » restent « annoncés/non vérifiés »), et ne remplace pas une évaluation sur vos propres tâches — les benchmarks cités sont très majoritairement vendor-déclarés.
- **Périmètre** : les familles chinoises (MiniMax, HY3/Tencent, Kimi, DeepSeek, Qwen, GLM, MiMo) sont dans le **volume Chine** de l'encyclopédie, pas ici — sauf mention explicite (Hermes Agent supporte ces endpoints ; § 151 signale le trou de l'open-weight occidental en vidéo).
- **Fraîcheur** : toute donnée est « au 27/09/2026 ». Les prix API changent vite (trois baisses OpenAI en deux mois cet été) — vérifier le pricing du fournisseur avant tout engagement.

---

# PARTIE XIV — Dossiers transverses : ce que les fiches ne disent pas seules

## 155. Dossier : l'affaire des benchmarks — pourquoi les chiffres 2026 sont fragiles

Quatre affaires documentées dans les sources de ce volume montrent que les scores 2026 ne se comparent pas naïvement :

1. **Meta / Llama 4 (avril 2025, révélé 2026)** : Yann LeCun a admis au Financial Times que Meta avait utilisé **différentes versions de Maverick/Scout sur différents benchmarks** (« fudged a little bit »). Conséquence : tous les chiffres Llama 4 vendor sont suspects — d'où la réserve transversale § 93.
2. **OpenAI / GPT-6 Astra, ARC-AGI-3** : 99,9 % sous le harness OpenAI, **62,7 % sous le harness standardisé ARC** (intelligentliving.co). Le harness fait le score : un benchmark « résolu » peut n'être qu'un benchmark domestiqué.
3. **NVIDIA / Nemotron 3.5 Lightning, GPQA Diamond** : 75,44–75,57 % (NVIDIA) vs **63,0–68,8 %** mesurés indépendamment sur OpenRouter — 7 à 12 points d'écart. Même modèle, même benchmark, deux mondes.
4. **Google / Gemini 3.8 Flash, DeepSWE v1.1** : 73,7 % (secondtalent) vs **71,0 %** (Neowin) — écart non réconcilié entre deux sources tierces citant la même table de lancement.
5. **Anthropic / Opus 4.7, le tokenizer silencieux** : nouveau tokenizer = 1,0–1,35× tokens par requête, soit **+35–47 % de coût réel mesuré par la communauté** malgré un « prix inchangé ». Le benchmark ne mesure pas la facture.

Règles de lecture pour 2026 :
- Un score sans harness nommé ne vaut rien (exiger : harness, température, nombre de runs).
- « Vendor-déclaré » ≠ « indépendant » — les deux sont notés comme tels dans ce volume.
- Les index composites (Artificial Analysis Intelligence Index, LMArena Elo) changent de version : v4.3 donne Fable 5.1 ex æquo 1er (53), v4.3.2 donne Opus 5.5 1er (57,6). Citer la version ou ne rien citer.
- Le coût par tâche (tokens × prix) est la seule métrique économique honnête — voir § 140.

## 156. Dossier : retraits, EOL et dépréciations — le cimetière 2026

