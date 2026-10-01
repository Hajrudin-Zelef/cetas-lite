---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-7
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "MiniMax", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Xiaomi", "Z.ai"]
dates: ["2026-06-16", "2026-08-14", "2026-08-27", "2026-08-28", "2026-09-27"]
keywords: ["agent", "agents", "attention", "benchmarks", "claude", "cyber", "deepseek", "exploit", "fine-tuning", "glm", "inference", "kimi"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [596, 690]
sha256: 818f966b4b7ebd3e44f6ffa147e7d0a3a2b26aea699aea60b9f50243ef0537fe
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

- **Nom exact** : GLM-5.2 (ID API : `glm-5.2` ; HF `zai-org/GLM-5.2`). **13 juin 2026** (GLM Coding Plan),
  **16 juin 2026** (API + poids ouverts). Statut : disponible, poids ouverts.
- **Architecture** : MoE — **~744–753B total / ~40B actifs** par token (incohérences mineures entre
  outlets : 744B, 750B, 753B — Z.ai n'a publié aucun papier d'architecture) ; **256 experts routés +
  1 partagé** ; architecture `glm_moe_dsa`.
- **Attention** : **MLA + DSA (Dynamic Sparse Attention)** + nouveauté **IndexShare** (des groupes de
  couches partagent les indices top-k de l'attention sparse → coût réduit à 1M de contexte).
- **Contexte** : **1M tokens** (solide/utilisable, vs 200K sur GLM-5.1) ; sortie max ~128–131K tokens.
- **Fonctionnalités** : **texte uniquement** (pas de vision) ; coding-first / agents long-horizon ;
  **effort de raisonnement réglable** (`reasoning_effort` : non-thinking → High → Max) ; couche MTP
  améliorée (décodage spéculatif) ; harnais officiel **ZCode** livré avec ; compatible Claude Code, Cline,
  OpenCode, Roo Code, Goose, Crush, OpenClaw, Kilo Code dès le jour J.
- **Poids** : **ouverts, licence MIT** (pleine, permissive — « no regional limits »). HF `zai-org/GLM-5.2`, ModelScope.
- **Benchmarks** : TerminalBench v2.1 **78 %**, GPQA Diamond **89 %**, SWE-bench Pro **62,1 %** (devant
  GPT-5.5 et Claude sur ce bench), AA Intelligence Index v4.1 **51** (n°1 de l'index à sa sortie, devant
  MiniMax-M3 44, DeepSeek V4 Pro 44, Kimi K2.6 43), ProofBench 30 %+ (11 pts devant le n°2),
  GDPval-AA v2 agent : **1524 Elo** (devant GPT-5.5 : 1514). N°1 open-weight sur FrontierSWE, PostTrainBench,
  SWE-Marathon (revendication Z.ai).
- **Prix API** : **$1,40 / 1M input, $4,40 / 1M output, $0,26 / 1M cached input** (= prix GLM-5.1 repris à
  l'identique) ; dispo via GLM Coding Plan (Lite/Pro/Max/Team) et API.
- **Déploiement** : Ollama (référencé), HF Inference Providers, API Z.ai (OpenAI-compatible), self-hosted.
- **Fine-tuning** : non vérifié au 27/09/2026 — aucune API de fine-tuning Z.ai ni recette communautaire
  vérifiée trouvée.

## 46. GLM-5.3 (le « cyber »)

- **Nom exact** : GLM-5.3 (HF `zai-org/GLM-5.3`). **14 août 2026** (API/Coding Plan), **poids ouverts le
  28 août 2026** (retenus 2 semaines pour durcissement sécurité). Statut : disponible, poids ouverts
  (licence exacte à confirmer sur la carte HF — la famille est MIT, mais la carte 5.3 n'a pas été vérifiée
  directement).
- **Architecture** : **même base pré-entraînée que GLM-5.2** (743–744B total / ~40B actifs — Z.ai : « chaque
  gain vient du post-training, aucun changement architectural »). **Divergence** : une source tierce
  (codepick) indique 730B / ~70B actifs — non résolu.
- **Attention** : MLA + DSA + IndexShare (hérité de 5.2). **Contexte : 1M tokens**.
- **Fonctionnalités** : positionné **« frontier coding with emergent cyber capabilities »** — post-training
  avec environnements de découverte de vulnérabilités → capacités cyber émergentes (analyse de vulns,
  exploit scripting, élévation de privilèges, chaînes d'exploitation multi-étapes ; 2 436 vulns trouvées
  sur 269 projets open-source, 53 CVE — revendication Z.ai). **Thinking non désactivable** :
  `thinking.type: "disabled"` n'est plus supporté (la requête échoue) ; `reasoning_effort` = low/high/max
  (défaut max).
- **Benchmarks** : Terminal-Bench 3.0 : **4,6 % → 28,3 %** (×6 vs 5.2), DeepSWE v1.1 : 46,2 % → **66,9 %**,
  CyberGym : 77,2 % → **84,5 %**, ExploitBench : 24,4 % → **54,4 %**, HumanEval 94,5 %, MBPP 91,2 %,
  LiveCodeBench 68,3 %, LMArena Elo 1498, AA Intelligence 63,8. **Divergence massive non résolue sur
  SWE-bench Verified : 57,8 %** (codepick, « données officielles ») **vs 82,4 %** (matrice ifnodoraemon) —
  l'un des deux chiffres (ou les deux) est faux ou mal labellisé.
- **Prix API** : **$1,40/$4,40** (annonce 14/08, Coding Plan) **vs $0,50/$2,00** (matrice ifnodoraemon) —
  divergence ; possiblement une baisse de prix entre août et septembre, non vérifié.
- **Déploiement** : API Z.ai, GLM Coding Plan, poids HF (28/08/2026).

## 47. GLM-5.3-Flash (« Ox Alpha »)

- **Nom exact** : GLM-5.3-Flash (HF `zai-org/GLM-5.3-Flash`) — apparu en stealth sous le nom **« Ox Alpha »**
  (`stealth/ox-alpha` sur OpenRouter, 20 août, gratuit 1 semaine), confirmé par Zhipu à Bloomberg le
  **26 août 2026**, lancement officiel + poids le **27 août 2026**.
- **Statut** : disponible, poids ouverts, **licence MIT**.
- **Architecture** : MoE — **320B total / 18B actifs** par token.
- **Contexte** : **1M tokens, nativement multimodal** (texte + image + vidéo en entrée).
- **Fonctionnalités** : tier « Flash » vitesse/débit (matrice tierce : 260 tok/s en Flash vs 115 pour le
  5.3 standard) ; coding/agents à bas coût.
- **Fait industriel** : servi initialement depuis un cluster de **100 000 puces de fabrication chinoise**
  (déclaration Zhipu à Bloomberg) — la démonstration que l'inférence frontier peut tourner sans NVIDIA.
- **Benchmarks** : non vérifié (aucun bench publié dans les sources consultées au 27/09).
- **Prix API** : **$0,15 / 1M input, $0,50 / 1M output** (prix catalogue).
- **Déploiement** : API Z.ai officielle ; poids HF.

## 48. GLM : MLA + DSA + IndexShare, explication

La signature technique de Zhipu combine trois briques. **MLA** (attention latente, style DeepSeek) pour
compresser le KV cache. **DSA** (Dynamic Sparse Attention) : au lieu d'attentionner sur tous les tokens
passés, le modèle sélectionne dynamiquement un sous-ensemble (top-k) de positions pertinentes — l'attention
devient sparse et le coût par token chute sur les très longues séquences. **IndexShare** (nouveauté 5.2) :
des groupes de couches **partagent les mêmes indices top-k** au lieu de les recalculer chacun, ce qui
divise encore le coût à 1M de contexte. C'est l'équivalent Zhipu du GDN de Qwen et du KDA de Kimi : trois
laboratoires, trois noms, une même direction — tuer le coût quadratique de l'attention.

## 49. GLM : la parenthèse « cyber » du 5.3

Le positionnement « emergent cyber capabilities » du GLM-5.3 est une première assumée : un constructeur
qui markete la découverte de vulnérabilités et l'exploit scripting comme feature. Z.ai a retenu les poids
deux semaines pour « durcissement sécurité » avant publication — aveu implicite du dilemme dual-use. Les
chiffres (2 436 vulnérabilités trouvées, 53 CVE, ExploitBench ×2,2) sont vendor-reported et invérifiables
indépendamment, mais la trajectoire est claire : les frontier ouverts de 2026 savent attaquer, et les
labos chinois ne s'en cachent plus. Pour un RSSI, c'est une donnée à intégrer : l'outillage offensif
open-weight est désormais au niveau des équipes de recherche.

## 50. GLM : tableau comparatif

| Modèle | Date | Archi | Params (total/actifs) | Contexte | Prix API ($/1M in-out) | Licence |
|---|---|---|---|---|---|---|
| GLM-5.2 | 16/06/2026 | MoE, MLA + DSA + IndexShare | ~744–753B / ~40B | 1M | 1,40 / 4,40 | MIT |
| GLM-5.3 | 14/08/2026 (poids 28/08) | MoE, même base 5.2 (post-training) | ~744B / ~40B (div. : 730B/70B) | 1M | 1,40 / 4,40 (div. : 0,50/2,00) | MIT (à confirmer carte HF) |
| GLM-5.3-Flash | 27/08/2026 | MoE | 320B / 18B | 1M (multimodal) | 0,15 / 0,50 | MIT |

## 51. Xiaomi MiMo : portrait d'équipe

