---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-12
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Anthropic", "Hugging Face", "MiniMax", "OpenAI", "SGLang", "Unsloth", "vLLM"]
dates: ["2026-08-02", "2026-09-27"]
keywords: ["agentic", "agents", "attention", "benchmark", "benchmarks", "claude", "datacenter", "fine-tuning", "fp8", "gguf", "gpu", "gqa"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1080, 1172]
sha256: 59a8304617e74405e519693f8b6eeb200581028fb9a0839d8b3121d7620d35f3
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Noms exacts** : MiniMax-M2 (23 octobre 2025), M2.1 (22–23 décembre 2025), M2.5 (février 2026 — une
  source indique le 12/02, à confirmer), M2.7 (annoncé 18 mars 2026, poids open-sourcés le 12 avril 2026).
  Statut : disponible (API + poids ouverts).
- **Architecture** : **MoE sparse** — **230B paramètres totaux, 10B actifs par forward** (256 experts fins,
  8 actifs, sigmoid routing avec expert biases). 62 blocs décodeur, hidden size 3 072, vocab 200 064.
  **Full attention GQA dans toutes les couches** (48 query heads, 8 KV heads, RoPE) — retour au full
  attention après l'hybride Lightning/full de MiniMax-01. Module **MTP** : 1 module au pré-entraînement,
  étendu à 3 par copie de poids pour le speculative decoding. Pré-entraînement : 29,2T tokens.
- **Contexte** : **192K–205K natif** (192K selon arXiv 2605.26494 ; 204.8K input / 176.9K output selon
  Artificial Analysis ; limite par séquence 196K selon un profil d'intégration).
- **Fonctionnalités** : text in/out, function calling, structured output, reasoning toujours actif. Design
  orienté **coding & agents** (multi-file edits, code-run-fix loops, toolchains long-horizon
  shell/browser/retrieval/code runners). Pas de vision.
- **Poids** : ouverts (HuggingFace `MiniMaxAI/MiniMax-M2*`). **Nom exact de la licence : non vérifié**
  (controverse de licence autour de M2.7 mentionnée sans détail par awesomeagents).
- **Inférence** : poids ~220 Go ; ~240 Go de KV cache par 1M tokens (note fournisseur citée par un profil
  d'intégration) ; MTP (3 têtes) pour speculative decoding ; serveurs recommandés vLLM (4×H100 en FP8) ou
  SGLang, parsers d'outils `minimax_m2` / `minimax_m2_append_think`. Sampling : temperature=1.0,
  top_p=0.95, top_k=40.
- **Benchmarks** : M2 (oct 2025, fournisseur) — SWE-bench Verified 69.4, ArtifactsBench 66.8, τ²-Bench
  77.2, GAIA text-only 75.7, BrowseComp 44.0, FinSearchComp-global 65.5. M2.7 — SWE-Pro 56.22 % (parité
  GPT-5.3-Codex), Terminal Bench 2 57.0 %, VIBE-Pro 55.6 %, GDPval-AA 1495 ELO (#1 open-source), MLE Bench
  Lite 66.6 % medal rate. M2.5 → M2.7 : taux d'hallucination AA-Omniscience de 88 % à **34 %** (meilleur
  que Claude Sonnet 4.6 à 46 %). Artificial Analysis : #1 composite open-source.
- **Prix API** : **$0,30 / 1M input, $1,20 / 1M output** (tarif cité par Magica, sept 2026).
- **Déploiement** : vLLM, SGLang (parsers dédiés), Transformers, GGUF (communauté), KTransformers, ATOM
  (ROCm), MLX-LM, Ollama (docs tierces).
- **Fine-tuning** : non vérifié au 27/09/2026.

## 86. MiniMax-M2.7 et le « self-evolution harness »

La nouveauté conceptuelle de M2.7 (mars 2026) est le **« self-evolution harness »** : mémoire court terme,
self-feedback, self-optimization — le modèle gère **30–50 % du workflow RL interne** de MiniMax
(VentureBeat). Autrement dit, MiniMax utilise M2.7 pour entraîner ses propres successeurs : la boucle
« modèle qui s'améliore lui-même » n'est plus un slogan mais une part mesurée du pipeline. C'est aussi
avec M2.7 qu'est apparue une « controverse de licence » (awesomeagents) dont le détail n'a pas été vérifié
— prudence donc sur la réutilisation commerciale des poids M2.7 sans lecture directe de la licence.

## 87. MiniMax-M3 (juin 2026)

- **Nom exact** : MiniMax-M3. **Sortie : 1er juin 2026** (annonce MiniMax API docs). Statut : disponible
  (modèle de langue courant de la série M).
- **Architecture** : **MoE**. Nombre total de paramètres : **non vérifié** — le checkpoint
  `MiniMaxAI/MiniMax-M3-MXFP8` fait ~444 Go (ce qui suggère ~440B paramètres en MXFP8 ≈ 1 octet/param,
  mais c'est une inférence, pas une spec).
- **Contexte** : **1M tokens** (contexte long).
- **Fonctionnalités** : agentic reasoning, tool use, coding, **chat multimodal en input** (texte +
  image/vision en entrée), long-context. Pas de vision en sortie.
- **Poids** : ouverts (HF `MiniMaxAI/MiniMax-M3-MXFP8` notamment). Nom de licence : non vérifié.
- **Inférence** : self-hosting marqué **expérimental** par MiniMax (revue doc : 26 août 2026). Baseline :
  8×B200, image SGLang pinnée par digest, `--reasoning-parser auto --tool-call-parser auto --tp 8`.
  Support : SGLang, vLLM, Transformers (`minimax_m3_vl`), KTransformers, Unsloth, ATOM (ROCm), MLX-LM.
  Sampling : temperature=1.0, top_p=0.95 (pas de top_k).
- **Benchmarks** : avec le framework **MaxProof** (generative-verifier RL + evolutionary search), M3
  dépasse le **seuil médaille d'or humaine sur IMO 2025 et USAMO 2026** (blog MiniMax, 9 juin 2026). Pas
  de chiffres de benchmarks publics tiers détaillés dans les sources trouvées.
- **Prix API** : non vérifié. **Fine-tuning** : non vérifié.

## 88. MiniMax-H3 (juillet 2026, vidéo + audio génératifs)

- **Nom exact** : MiniMax-H3. Annoncé à **WAIC 2026 le 17 juillet 2026** ; **API le 31 juillet 2026** ;
  **poids ouverts le 3 août 2026**. Statut : disponible (API + poids ouverts partiels).
- **Architecture** : **omni-modal génératif (vidéo+audio)** — **H3-Omni-Transformer dense 33B single-stream**
  qui prédit conjointement les latents vidéo et audio ; encodages spécifiques par modalité packés en
  séquence multimodale ; H3-VisualVAE et H3-AudioVAE (espaces latents séparés) ; sparse attention annoncée
  native (update ultérieure). Deux checkpoints Base distillés CFG task-specific : **FL2VA** (texte /
  first-or-last-frame conditioning) et **Ref2VA** (références multimodales : jusqu'à 9 images, 3 clips
  vidéo, 3 clips audio, max 12 fichiers).
- **Sorties** : vidéo **4–15 secondes à 24 FPS, audio stéréo 32 kHz natif** ; Base = 768p, workflow hébergé
  complet jusqu'à 2K.
- **Fonctionnalités** : text→vidéo, image→vidéo, vidéo native avec audio stéréo natif, séquençage
  multi-shot, prompt adherence, text rendering, character consistency, motion transfer.
- **Poids** : **ouverts** (checkpoints H3-Base) mais voir section 89 — **restriction territoriale**.
  Les composants H3-Context-IR et H3-Regenerate-2K restent hébergés ; le workflow 2K complet n'est pas
  en poids ouverts.
- **Benchmarks** : aucun benchmark public standard trouvé (modèle génératif — pas de benchmarks LLM
  classiques). **Prix API** : non vérifié. **Fine-tuning** : non vérifié.
- **Déploiement** : ComfyUI (guide Spheron), API MiniMax. VRAM : du GPU consommateur au datacenter selon
  Spheron.

## 89. MiniMax : la licence territoriale restrictive du H3

C'est la licence la plus inhabituelle du volume : la `minimax-h3-community-license-agreement` (datée
**2026-08-02**) **exclut explicitement les États-Unis, l'UE, le Royaume-Uni et la Corée du Sud** du
déploiement local sans autorisation individuelle. Autrement dit : poids « ouverts » mais pas pour tout le
monde — une première qui inverse la logique habituelle (d'habitude, c'est l'export depuis la Chine qui est
restreint, pas l'usage en Occident). Lecture géopolitique : MiniMax se protège contre la concurrence
frontale sur son propre terrain génératif tout en gardant l'API comme porte d'entrée mondiale. Pour un
utilisateur hors zones exclues, les poids Base 768p sont auto-hébergeables ; le workflow 2K complet
reste de toute façon hébergé.

## 90. humain-m3 : le dérivé saoudien (septembre 2026)

