---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-8
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "OpenAI", "Xiaomi", "xAI"]
dates: ["2026-09-27"]
keywords: ["agentic", "agents", "attention", "aws", "benchmarks", "cyber", "deepseek", "distillation", "gguf", "gpu", "grok", "grok 4"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [691, 796]
sha256: 79aafb1e32be3be98fa6a2b440c713f5c34da8fa8d672f3b06ed28612cb8856d
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

L'équipe **MiMo** de Xiaomi est dirigée par **Luo Fuli**, ex-DeepSeek et core dev de DeepSeek-V2 — ce qui
explique la filiation technique visible (MoE sparse, MLA, MTP). Xiaomi alterne les phases : open-weight
généreux (7B en 2025, V2.5, V2.6), parenthèse propriétaire (V2-Pro/Omni/TTS, mars 2026, API/produits
Xiaomi uniquement), puis retour à l'ouvert. La marque de fabrique 2026 de MiMo : l'**omnimodalité native**
(texte + image + vidéo + audio dans un seul modèle) et une transparence inédite — le post-training RL de
V2.6 a été **retransmis en direct** sur dashboard public.

## 52. MiMo-7B (avril 2025, la fondation)

- **Nom exact** : MiMo-7B (variantes : Base, SFT, RL, RL-Zero). **Sortie : 30 avril 2025**. Statut :
  disponible, open weights. Antérieur à la période, mais c'est la fondation de la lignée.
- **Architecture** : dense **7B** paramètres, optimisé reasoning. Pré-entraînement sur **25T tokens**
  (dont 200B dédiés reasoning). **MTP (Multiple Token Prediction)**.
- **Contexte** : non vérifié au 27/09/2026. **Poids** : **ouverts, MIT**.
- **Benchmarks** (Xiaomi, MiMo-7B-RL, temp 0,6) : MATH-500 95,8 %, AIME 2024 68,2 %, AIME 2025 55,4 %,
  LiveCodeBench v5 57,8 % / v6 49,3 %, GPQA Diamond 54,4 %, MMLU-Pro 58,6 % ; présenté comme > OpenAI
  o1-mini en maths/code.

## 53. MiMo-VL-7B et MiMo-Audio-7B (fiches prudentes)

- **MiMo-VL-7B** : vision. Date : **4 juin 2025** selon une timeline tierce (logeshwaran.org) — **non
  corroborée par une source primaire**. Statut : disponible, open weights. Specs détaillées : non
  vérifiées au 27/09/2026.
- **MiMo-Audio-7B** : audio. Date : **29 décembre 2025** (même timeline tierce — non vérifiée). Statut :
  disponible, open weights. Specs : non vérifiées.
- Ces deux modèles existent dans la lignée multimodale MiMo, mais la recherche n'a pas permis d'établir
  des fiches solides : ils sont cités pour mémoire, pas pour usage décisionnel.

## 54. MiMo-V2-Flash (décembre 2025, 1ère génération V2)

- **Nom exact** : MiMo-V2-Flash. **Sortie : 17 décembre 2025**. Statut : disponible, open weights.
- **Architecture** : MoE, **309B total / 15B actifs**. Ratio d'attention hybride 5:1 (source tierce, non vérifié).
- **Poids** : **ouverts, MIT**. Specs fines : non vérifiées au 27/09/2026.
- C'est le gabarit « Flash » de Xiaomi : ~300B totaux pour ~15B actifs, le format efficacité qui sera
  repris par V2.5 puis V2.6-Flash-RL.

## 55. MiMo-V2-Pro / V2-Omni / V2-TTS (la parenthèse fermée)

- **Noms exacts** : MiMo-V2-Pro, MiMo-V2-Omni, MiMo-V2-TTS. **Sortie : 18 mars 2026** (Spring launch event,
  Lei Jun). Statut : disponible via API/produits Xiaomi.
- **Architecture** : V2-Pro — **>1T total / 42B actifs** (MoE). V2-Omni / V2-TTS : specs non vérifiées.
- **Poids** : **fermés / propriétaires** — la seule parenthèse fermée de Xiaomi, refermée dès V2.5 (avril 2026).
- **Contexte / prix / benchmarks** : non vérifiés au 27/09/2026 — les specs de ces modèles propriétaires
  sont quasi introuvables publiquement.

## 56. MiMo-V2.5 / V2.5-Pro (avril 2026, retour à l'ouvert)

- **Noms exacts** : MiMo-V2.5, MiMo-V2.5-Pro. **Sortie : 22 avril 2026**. Statut : disponible, open weights.
- **Architecture** : MoE — V2.5 : **310B total** ; V2.5-Pro : **1,02T total** (actifs non vérifiés).
- **Poids** : **ouverts, MIT**. Fait industriel : adapté **jour 1 à 7 plateformes de chips chinoises**
  (T-Head, Kunlun, Enflame, Muxi, Tianshu Zhixin...) + AWS ; programme d'incitation **100T tokens gratuits**
  (30 jours).
- **Contexte / prix / benchmarks** : non vérifiés au 27/09/2026.
- Lecture : Xiaomi pousse l'inférence sur silicium domestique dès le jour de sortie — l'alignement
  modèle-puce chinoise est une stratégie d'écosystème, pas un accident.

## 57. MiMo-V2.6-Pro-RL (septembre 2026)

- **Nom exact** : MiMo-V2.6-Pro-RL. **Poids publiés le 21 septembre 2026**, annonce le 22 (décalage fuseau
  Chine). Statut : disponible, open weights.
- **Architecture** : MoE sparse, **1,02T total / 42B actifs**, 70 couches (attention hybride sliding-window
  + globale), **384 experts routés, 8 actifs/token**, encodeur vision 681M, encodeurs audio, décodeur MTP
  5 couches.
- **Contexte** : **1M tokens**. **Omnimodal natif** : texte + image + vidéo + audio dans un seul modèle.
- **Fonctionnalités** : agents/tool use (focus), raisonnement, méthode « You Only RL Once » (un seul run RL
  mixte code/général/vision/cyber + distillation on-policy).
- **Poids** : **ouverts, MIT** (BF16 sur Hugging Face ; Pro-RL = 573,49 Go de poids, 155 fichiers).
- **Inférence** : non vérifié au 27/09/2026 — pas de données publiques trouvées. Pro demande 2 nœuds /
  32 GPU selon la commande de serving Xiaomi.
- **Benchmarks** : AA Intelligence Index **46** (indépendant, Artificial Analysis — top open-weights, à
  égalité avec Grok 4.7 selon aiweekly). Vendor-reported : DeepSWE v1.1 71,9 (Opus 5 : 74,0),
  AutomationBench 53,1, Toolathlon-Verified 76,9, OSWorld-Verified 82,0, Terminal Bench 2.1 89,9,
  CyberGym 94,0, ExploitBench 47,9.
- **Prix API** : Pro **$0,435 / 1M input, $0,87 / 1M output** (tiers — mêmes chiffres que DeepSeek V4-Pro,
  coïncidence tarifaire du marché).
- **Déploiement** : API Xiaomi (MiMo Open Platform), MiMo Studio (chat web), MiMo Desktop ; poids HF.

## 58. MiMo-V2.6-Flash-RL et Pro-UltraSpeed

- **MiMo-V2.6-Flash-RL** : **309B total / 15B actifs**, contexte 1M, omnimodal, MIT. Prix API :
  **$0,14 / $0,28** par 1M (tiers). Le « Flash » de la génération : même ratio ~20× entre total et actifs
  que le V2-Flash de décembre 2025, la recette Xiaomi de l'efficacité est stable.
- **MiMo-V2.6-Pro-UltraSpeed** : ~1T paramètres, **API uniquement** — la variante vitesse du Pro, non
  ouverte. Specs fines : non vérifiées.

## 59. MiMo-V2.6-Distill-Qwen-9B (le petit malin)

- **Nom exact** : MiMo-V2.6-Distill-Qwen-9B. Sortie avec la famille V2.6 (21–22 septembre 2026).
- **Architecture** : **9B dense**, SFT de Qwen3.5-9B, distillé pour la recherche agentic-RL.
- **Poids** : ouverts, MIT. GGUF officiel (ggml-org, Q4_K_M) — **16 Go VRAM suffisants**.
- **Déploiement** : Ollama / LM Studio. C'est le seul modèle chinois « trillion-class lineage » à tourner
  sur une machine bureautique : distillé depuis un 1T, il embarque une partie du savoir agentic du Pro-RL
  dans 9B. Pour un labo local ou un RAG embarqué, c'est la porte d'entrée la plus accessible de tout ce volume.

## 60. MiMo : le RL retransmis en direct (transparence inédite)

Fait sans précédent dans l'industrie : Xiaomi a **retransmis en direct le post-training RL** de V2.6 sur
un dashboard public (mimo.xiaomi.com/rl/) du 15 au 20 septembre 2026. Chiffres publiés : coût RL **~$3,47M**
($2,62M Pro + $854K Flash), 30 steps, ~750K trajectoires, <6 jours. **7 000+ environnements RL et le code
d'entraînement publiés**. C'est la première fois qu'un labo frontier ouvre le capot du RL à ce niveau :
on connaît désormais le coût marginal du post-training d'un modèle trillion-class (~3,5M$), et c'est
dérisoire face au pré-entraînement. Le message implicite : la barrière n'est plus le RL, c'est le
pré-entraînement et les données.

## 61. MiMo : la controverse Anthropic (septembre 2026)

