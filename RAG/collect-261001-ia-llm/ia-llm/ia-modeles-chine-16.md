---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-16
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Huawei", "LongCat", "Meituan", "MiniMax", "Moonshot", "Nvidia", "SGLang", "United States", "Xiaomi", "Z.ai", "vLLM"]
dates: []
keywords: ["agent", "agents", "apache", "ascend", "attention", "attribution", "benchmark", "benchmarks", "blackwell", "claude", "compute", "datacenter"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1469, 1601]
sha256: 0c7efef1073914a11e7aa0a519cd60b6c8e0ea8d9abca87e0d9f9fac72a25128
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

## 118. Glossaire — Agent natif vs tool calling

**Tool calling** : le modèle sait appeler des outils (fonctions) — brique de base. **Agent natif** : le
modèle est entraîné pour des boucles autonomes longues (planifier, agir, observer, se corriger) —
Qwen3.7-Max revendique 35 heures continues, Kimi K2.6 des essaims de 300 agents / 4 000 étapes. La
différence se joue dans le post-training (RL sur environnements), pas dans l'architecture.

## 119. Glossaire — RL (post-training) et DORA

Le post-training par apprentissage par renforcement (RLVR, GRPO...) aligne le modèle sur les tâches
cibles (code, agents, maths). Fait 2026 : Xiaomi a **retransmis son RL en direct** (coût ~$3,47M pour
V2.6) ; Meituan utilise **DORA** (rollouts asynchrones, >3× speedup) ; MiniMax fait gérer 30–50 % de son
RL par M2.7 lui-même (self-evolution harness).

## 120. Glossaire — Distillation

Entraîner un petit modèle sur les sorties d'un grand (ex. MiMo-V2.6-Distill-Qwen-9B : 9B distillé depuis
la lignée trillion-class). Légale entre modèles ouverts ; **controversée** quand elle pompe les sorties
d'une API commerciale contre ses conditions d'utilisation (accusation d'Anthropic contre Xiaomi,
sept. 2026 — non tranchée).

## 121. Glossaire — Quantification (INT4, FP8, NVFP4, GGUF)

Réduction de la précision des poids pour servir moins cher : FP8 (8 bits), INT4 (4 bits, ex. Kimi livré
nativement en INT4), NVFP4 (format 4 bits NVIDIA Blackwell), **GGUF** (format communautaire llama.cpp/
Ollama, ex. Q4_K_M). Un 124B en BF16 (~250 Go) passe à ~70 Go en FP4 — c'est ce qui rend les gros
ouverts servables hors datacenter.

## 122. Glossaire — vLLM / SGLang / Ollama

Les trois moteurs d'inférence open-source de référence : **vLLM** (le standard datacenter,PagedAttention),
**SGLang** (le challenger, fort sur le structured output et le speculative decoding — support day-0 de
Kimi K3, Hy3, Qwen3.8-Flash), **Ollama** (le local simple, GGUF, support day-0 du Qwen3.8-27B). Le support
day-0 par ces moteurs est devenu un critère de lancement aussi important que les benchmarks.

## 123. Glossaire — Open-weight vs open-source

**Open-weight** : les poids sont publiés (on peut télécharger et servir le modèle), mais le code
d'entraînement et souvent les données restent fermés. **Open-source** (au sens OSI strict) exigerait
tout. En 2026, « open-source » est utilisé abusivement pour « open-weight » dans 90 % de la presse :
ce volume dit « poids ouverts » quand c'est le cas précis, et donne la licence à chaque fois.

## 124. Glossaire — Licences (MIT, Apache 2.0, Modified MIT)

**MIT** et **Apache 2.0** : permissives, usage commercial inclus (DeepSeek, Zhipu, Xiaomi, Tencent,
Qwen — c'est le standard chinois 2026). **Modified MIT** (Moonshot) : MIT + clause d'attribution
au-delà de 100M MAU / $20M mensuels (voir section 38). **Licence communautaire territoriale**
(MiniMax-H3) : usage local interdit aux US/UE/UK/Corée du Sud sans autorisation — l'exception qui
confirme la règle permissive.

## 125. Glossaire — Snapshot (ex. Qwen3.8-Max-0902)

Version datée d'un même modèle (ici : 2 septembre 2026), qui peut changer comportement, benchmarks et
**coût par tâche** (le 0902 parle ~108k tokens/tâche, soit 2× plus cher que le 3 août). En 2026, le
modèle n'est plus une release figée mais un flux : toujours épingler le snapshot en production.

## 126. Glossaire — Peak / off-peak (tarification)

Tarification horaire inventée par DeepSeek : prix doublés en « peak » (01:00–04:00 et 06:00–10:00 UTC,
jours ouvrés hors fériés chinois). Le V4.1-Flash passe de $0,15/$0,60 à $0,30/$1,20. C'est la
financiarisation du compute : le token a un prix spot qui suit la charge du datacenter.

## 127. Glossaire — Cache hit / cache miss

Facturation selon que le préfixe de la requête est déjà en cache : DeepSeek V4-Pro facture $0,435/M en
cache miss contre **$0,003625/M** en cache hit (120× moins). Architecturer ses prompts avec des préfixes
stables (system prompt, contexte RAG figé) devient un levier économique direct.

## 128. Glossaire — SWE-bench (Verified, Pro)

Benchmark de **résolution de vraies issues GitHub** — la mesure reine du code en 2026. **Verified** :
500 tâches validées humainement ; **Pro** : version plus dure et plus récente. Scores du volume :
Kimi K3 85,6 % (Verified), MiMo-V2.6-Pro 71,9 % (DeepSWE v1.1), Qwen3.8-27B 61,7 % (Pro). Tous
vendor-reported sauf mention contraire.

## 129. Glossaire — Terminal-Bench / OSWorld / Toolathlon

Benchmarks **agentiques** (pas Q&R) : Terminal-Bench (tâches en terminal réel), OSWorld (contrôle d'un OS
via GUI), Toolathlon (usage d'outils enchaînés). Ils mesurent ce que les labos vendent en 2026 : l'agent
autonome. Scores notables : MiMo-V2.6-Pro 89,9 (Terminal-Bench 2.1) et 82,0 (OSWorld-Verified), GLM-5.2 78 %
(TerminalBench v2.1).

## 130. Glossaire — AA Intelligence Index / LMArena

**Artificial Analysis Intelligence Index** : score composite indépendant (0–100) agrégeant les benchmarks —
la mesure « tierce » la plus citée du volume (Kimi K3 : 65 ; GLM-5.3 : 63,8 ; Qwen3.7-Max : ~57 ;
MiMo-V2.6-Pro : 46 ; DeepSeek V4.1-Flash : 40). **LMArena** : classement par votes humains en aveugle
(Kimi K3 : Elo 1512 ; Qwen3.8-Max : 2e Vision Arena). Ce sont les deux garde-fous contre les chiffres
constructeur.

## 131. Glossaire — DiNA (Discrete Native Autoregression)

Paradigme de Meituan (LongCat-Next) : toutes les modalités (texte, image, audio) sont converties en
**tokens discrets dans un espace partagé**, puis un unique autoregresseur les prédit. Alternative aux
architectures « backbone texte + encodeurs greffés » : ici, la multimodalité est native dès le tokenizer.

## 132. Glossaire — RVQ (Residual Vector Quantization)

Technique de tokenization (ici : visuelle et audio) qui quantifie un signal en plusieurs couches
résiduelles successives — ex. le tokenizer visuel **dNaViT** de LongCat-Next (RVQ 8 couches). Plus de
couches = plus de fidélité, au prix de plus de tokens par image.

## 133. Glossaire — ScMoE (Shortcut-Connected MoE)

Variante MoE de Meituan : des connexions « shortcut » entre experts permettent de **recouvrir le calcul
et la communication** inter-experts (le goulot du MoE distribué). Complétée par des experts
« zero-computation » (routage à coût nul). École d'optimisation système plutôt que d'architecture
d'attention.

## 134. Glossaire — Omnimodal

Un seul modèle qui comprend **et** génère (ou au moins comprend) texte + image + vidéo + audio : MiMo-V2.6
(texte+image+vidéo+audio), LongCat-Next (texte+vision+audio, génération d'images), GLM-5.3-Flash
(texte+image+vidéo en entrée), Ming-Flash-Omni-2.0 (parole+audio+musique). Le multimodal « greffé »
(encodeur vision ajouté) vs « natif » (DiNA) est la ligne de fracture technique.

## 135. Glossaire — Harnais agent (ZCode, OpenClaw, Claude Code)

Le **harnais** est le logiciel autour du modèle : boucle d'outils, gestion du contexte, exécution de
code. **ZCode** (Zhipu, livré avec GLM-5.2), **Claude Code** (Anthropic, dont le protocole est implémenté
par Qwen), **OpenClaw** (open-source, multi-fournisseurs). En 2026, le harnais compte autant que le modèle :
Qwen3.7-Max est évalué « cross-scaffold » (Claude Code, OpenClaw, Qwen Code) parce que l'agentique se joue
à ce niveau.

## 136. Glossaire — Frontier / frontier-closed

**Frontier** : un modèle au niveau de l'état de l'art mondial (quelques mois d'avance max). En 2026,
plusieurs ouverts chinois sont frontier (Kimi K3, GLM-5.3, Qwen3.8-Max). **Frontier-closed** : la doctrine
Alibaba d'avril–mai 2026 (Qwen3.6-Max, 3.7-Max/Plus) — garder le frontier fermé pour le monétiser, ouvrir
le reste. Doctrine abandonnée en août avec l'ouverture du 3.8-Max.

## 137. Glossaire — Ascend / silicium domestique

