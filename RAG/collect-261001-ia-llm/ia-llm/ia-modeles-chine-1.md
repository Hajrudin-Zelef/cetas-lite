---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-1
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Hugging Face", "LongCat", "Meituan", "MiniMax", "Moonshot", "Nvidia", "SGLang", "Together AI", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agents", "attention", "benchmarks", "claude", "datacenter", "deepseek", "fine-tuning", "fp4", "fp8", "glm", "kimi", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1, 103]
sha256: 52159224f5b960eda7b700da6a327af2dc0c6cdd44046b2daa8624095b3d06af
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE
## Février 2026 → 27 septembre 2026

*Volume 1/3. Ce volume couvre les laboratoires chinois : DeepSeek, Alibaba (Qwen), Moonshot (Kimi),
Zhipu (GLM), Xiaomi (MiMo), Ant Group (Ling/BaiLing), Meituan (LongCat), MiniMax, Tencent (Hunyuan).
Les modèles américains (NVIDIA Nemotron, NousResearch Hermes) et le reste de l'Occident sont dans les
volumes 2 et 3. Date de référence des faits : **27 septembre 2026**. Toute spécification non vérifiable
est marquée « non vérifié au 27/09/2026 » et aucun chiffre n'est inventé.*

---

## 1. Ce que couvre ce volume

Ce volume recense les modèles de langage et multimodaux sortis, annoncés ou retirés entre février 2026
et le 27 septembre 2026 par neuf acteurs chinois. Chaque famille a sa partie : portrait de l'acteur,
fiches modèle par modèle, encadrés techniques, puis un tableau comparatif de synthèse. La dernière partie
croise tout : grand tableau Chine, timeline, analyse (« que retenir »), glossaire de 32 termes et quiz.

Périmètre strict : seules les informations vérifiées par la recherche documentaire de septembre 2026 sont
reprises. Les fiches signalent honnêtement les trous (modèles annoncés mais jamais sortis, specs
introuvables, contradictions entre sources). Les modèles antérieurs à février 2026 ne sont rappelés que
lorsqu'ils éclairent la lignée 2026 (ex. MiMo-7B d'avril 2025, LongCat-Flash d'août 2025).

## 2. Comment lire une fiche modèle

Chaque fiche suit le même plan : nom exact (avec ID API et repo Hugging Face quand ils sont connus),
date de sortie, statut au 27/09/2026 (disponible / retiré / annoncé), architecture (type, paramètres
totaux et actifs par token), contexte, fonctionnalités, fine-tuning, poids et licence, benchmarks clés,
prix API, déploiement. Les chiffres de benchmarks sont **vendor-reported** (fournis par le constructeur)
sauf mention « indépendant ». Les prix API viennent souvent de sources tierces (agrégateurs, fournisseurs)
et sont signalés comme tels quand le tarif officiel n'a pas été consulté directement.

## 3. Les mentions d'incertitude : mode d'emploi

Trois niveaux sont utilisés partout dans ce volume. **« Non vérifié au 27/09/2026 »** : l'information
circule (presse, notes communautaires) mais n'a pas été confirmée par une source primaire ou officielle.
**« Divergence »** : deux sources sérieuses donnent des chiffres différents, impossible de trancher.
**« Non trouvé »** : après plusieurs requêtes de recherche, aucune trace — le modèle n'existe probablement
pas sous ce nom. Ces mentions ne sont pas des coquetteries : sur un corpus 2026 largement pollué par du
contenu IA-généré qui se recopie, elles sont la seule protection contre la fausse précision.

---

## 4. DeepSeek : portrait de l'entreprise

DeepSeek, laboratoire chinois fondé par Liang Wenfeng (adossé au fonds High-Flyer), a défini le standard
de l'open-weight agressif : modèles géants Mixture-of-Experts (MoE), poids ouverts sous licence MIT,
prix API cassés, rapports techniques publics. En 2026, la famille V4 (sortie le 24 avril 2026) prolonge
la lignée V3/V3.2 avec un changement architectural majeur : abandon du MLA pur au profit d'une attention
hybride. DeepSeek reste l'acteur qui force tout le marché chinois à aligner ses prix vers le bas.

## 5. DeepSeek V4 : fiche générale

- **Nom exact** : DeepSeek V4 (préview le 24 avril 2026), décliné en deux variantes : **V4-Pro** et **V4-Flash**.
- **Statut au 27/09/2026** : V4-Pro disponible (version « V4 Pro 0813 » en GA depuis le 13 août 2026).
  V4-Flash retiré le 10 septembre 2026, remplacé par V4.1-Flash (routes API redirigées).
- **Architecture** : MoE sparse, **~1,6T paramètres totaux / 49B actifs par token** (V4-Pro). Attention
  **hybride** (CSA stride 4 + HCA stride 128 + SWA ~128) — la famille V4 abandonne le MLA pur de V3 pour
  cette attention hybride (sources GitHub tierces ; non vérifié par documentation officielle DeepSeek).
  MegaMoE (expert parallelism), résiduels « mHC », optimiseur Muon (+ AdamW partiel), FP4 QAT pour les
  poids MoE. Entraînement sur **32T+ tokens**.
- **Contexte** : **1M tokens** (fenêtre native), sortie max **384K tokens**.
- **Fonctionnalités** : reasoning (modes de pensée unifiés), tool use / agents (« Interleaved Thinking »),
  vision (le V4-Flash avait une variante vision expérimentale — non vérifié dans le détail).
- **Poids** : **ouverts, licence MIT** (model cards Hugging Face `deepseek-ai/DeepSeek-V4-Flash` et Pro).
  Formats BF16 / FP8 / FP4 natifs (NVFP4).
- **Inférence** : attention hybride conçue pour réduire le KV cache (~10 % de V3.2 selon la model card) ;
  Together AI (mai 2026) documente 3 layouts de cache simultanés (CSA compressé / HCA dense ~8K entrées /
  SWA exact) ; capacité ~3,7M tokens par nœud HGX B200 avec politique optimisée. Déploiement cible :
  vLLM ou SGLang (forks DeepSeek), DeepGEMM + FlashInfer + DeepEP, EP sur NVLink (NVL72). Chiffres exacts
  par source officielle : non vérifiés. Serveurs datacenter : 8× H100/H200 ou 16+ × B100/B200.
- **Benchmarks clés** (vendor-reported) : V4 Pro Max — Codeforces **3206**, SWE-bench Verified **80,6 %**,
  SimpleQA 57,9 %, Chinese-SimpleQA 84,4 % ; bat Claude Opus 4.8 sur Terminal Bench 2.1.
- **Prix API** (V4-Pro) : **$0,435 / 1M tokens input** (cache miss), **$0,003625 / 1M** (cache hit),
  **$0,87–0,93 / 1M output**. Concurrency limitée à 500 requêtes (Pro).
- **Fine-tuning** : l'API DeepSeek propose la personnalisation — non vérifié au 27/09/2026 pour V4
  précisément. Support Ollama : non vérifié.

## 6. DeepSeek V4-Pro (0813)

- **Nom exact** : DeepSeek V4-Pro (version GA « V4 Pro 0813 », 13 août 2026).
- **Statut au 27/09/2026** : disponible, mais en **dépréciation ordonnée**. L'annonce DeepSeek du
  10 septembre 2026 (WeChat) indique que V4.1-Flash « a comprehensively surpassed V4 Pro » ; à partir du
  14 septembre 2026 (04:00 UTC), les requêtes `deepseek-v4-pro` devaient être servies par V4.1-Flash au
  tarif Flash, « jusqu'au lancement de V4.1 Pro ». **Contradiction non tranchée** : benchlm.ai (26 sept.
  2026) affirme que cette reroute a été **annulée** et que V4 Pro continue d'être servi ; des notes de
  recherche du 11 septembre affirment l'inverse. Impossible de trancher sans accès direct à l'API.
- **Architecture** : MoE ~1,6T total / 49B actifs (voir fiche V4). Mêmes benchmarks et prix que la fiche
  générale : Codeforces 3206, SWE-bench Verified 80,6 %, $0,435/$0,87–0,93 par 1M.
- **Poids** : ouverts, MIT.

## 7. DeepSeek V4-Flash (retiré)

- **Nom exact** : DeepSeek V4-Flash. Sortie : 24 avril 2026. **Retiré : 10 septembre 2026** (routes API
  redirigées vers V4.1-Flash le même jour).
- **Architecture** : MoE, **284B total / 13B actifs**, 256 experts routés + 1 partagé, top-6. NVFP4 natif.
- **Contexte** : 1M tokens. **Poids** : ouverts, MIT.
- **Rôle historique** : le « petit » efficace de la famille V4, positionné prix contre les tiers Flash
  concurrents (Qwen3.8-Flash, GLM-5.3-Flash). Son retrait au profit de V4.1-Flash illustre le rythme
  infernal des remplacements chez DeepSeek en 2026 : à peine 4,5 mois de vie commerciale.

## 8. DeepSeek V4.1-Flash

