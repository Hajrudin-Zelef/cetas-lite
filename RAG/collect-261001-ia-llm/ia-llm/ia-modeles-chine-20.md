---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-20
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Ant", "Anthropic", "DeepSeek", "Fireworks AI", "Huawei", "Hugging Face", "LongCat", "Meituan", "MiniMax", "Moonshot", "Nvidia", "OpenRouter", "SGLang", "TensorRT-LLM", "Together AI", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agent", "agents", "arr", "ascend", "asic", "attention", "benchmarks", "blackwell", "datacenter", "deepseek", "distribution", "embedding"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1957, 2063]
sha256: f0debdbfd335b756d1c38dd52a6cd18910f93360ec80e7157f72a2f3d1ce9564
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

D'après les données vérifiées du volume : **16 Go VRAM** → MiMo-V2.6-Distill-Qwen-9B (Q4_K_M, Ollama/LM
Studio). **24 Go (1× GPU)** → Qwen3.8-27B en 4-bit (18 Go Q4_K_M, MTP + décodage spéculatif), Ling-3.0-tiny
(SGLang, DGX Spark). **8× H200** → Kimi K2.6 en INT4 natif (baseline Moonshot). **4 nœuds / 32 GPU** →
Kimi K3 (minimum pratique cité). **2 nœuds / 32 GPU** → MiMo-V2.6-Pro-RL (commande Xiaomi). **DGX Spark
(GB10)** → Ling-3.0-flash en INT4 (72 Go), Hy3 en NVFP4 (~21,8 tok/s single-stream sur 2× DGX Spark, test
communautaire). Règle : les MoE sparse se servent au prorata des **actifs**, pas du total — un 124B/5,1B
(Ling-3.0) est bien plus servable que son total ne le suggère, à condition d'avoir les bons kernels
(vLLM/SGLang récents, support des attentions linéaires).

## 174. Veille : comment suivre après le 27/09/2026

Les points chauds à surveiller : **DeepSeek V4.1 Pro** (rumeur 28–30 sept. — vérifier le changelog
api-docs.deepseek.com), la **release open-weight d'humain-m3**, **Qwen4** (teasé via 3.8-Flash), la
**licence finale du Kimi K3** (carte HF `moonshotai/Kimi-K3`), les **cartes HF de Qwen3.8-Max/Flash et
GLM-5.3** (licences à confirmer), l'issue de la **controverse Anthropic–Xiaomi**, et l'évolution des
**prix GLM-5.3** ($1,40/$4,40 vs $0,50/$2,00 — une baisse serait un signal marché fort). Sources
primaires à bookmarker : blogs labos (minimax.io/blog, api-docs.deepseek.com/news), Hugging Face
(orgs `Qwen`, `moonshotai`, `zai-org`, `inclusionAI`, `meituan-longcat`, `MiniMaxAI`), Artificial Analysis
et LMArena pour les mesures indépendantes. Et garder le réflexe de ce volume : **un chiffre sans source
primaire reste une rumeur bien habillée**.

## 175. Traçabilité des sources

Ce volume est une synthèse rédigée à partir de quatre dossiers de recherche documentaire (septembre
2026) : `deepseek_mimo_ling_longcat.md` (DeepSeek, Xiaomi, Ant, Meituan), `qwen_kimi_glm.md` (Alibaba,
Moonshot, Zhipu), `minimax_nemotron_hermes.md` (MiniMax — sections Nemotron/Hermes exclues, volume
occidental), `llama_mistral_gemma_granite.md` (section HY3 = Tencent Hunyuan 3 uniquement). Les sources
primaires citées dans ces dossiers : communiqués (Business Wire, Reuters, Caixin, Bloomberg via Zhipu),
docs API officielles (api-docs.deepseek.com), model cards Hugging Face, blogs labos, papiers arXiv
(2603.00729, 2603.27538, 2605.26494, 2509.18883). Les sources secondaires : Artificial Analysis, LMArena,
OpenRouter, agrégateurs (best-of-ai, benchlm.ai), presse tech, notes communautaires GitHub. Les URLs
exactes figurent dans les dossiers de recherche ; ce volume n'invente aucune URL.

## 176. Ce que le volume 2 couvrira (pour mémoire)

Ce volume 1 s'arrête aux frontières de la Chine. Le **volume 2 (Occident)** reprendra : NVIDIA Nemotron
(Nano 2 12B, Nemotron 3 Nano/Super/Ultra, Cascade-2, 3.5 Lightning, Content Safety) et NousResearch
Hermes (Hermes 4, 4.3-36B, Hermes Agent, GEPA, Psyche, atropos) — documentés dans les mêmes dossiers de
recherche mais exclus ici par consigne éditoriale. Le **volume 3** croisera les deux hémisphères :
comparatif Chine–Occident, cartes des licences mondiales, et la question qui traverse tout 2026 —
l'open-weight peut-il rester ouvert quand le frontier coûte des centaines de millions à entraîner ?

---

*Fin du volume 1 (provisoire — recompté à la fin).* Vérification : `wc -l` et 100+ sections `## N.` exigés par la maquette.*

## 177. DeepSeek : Together AI et les trois layouts de cache

La documentation d'inférence la plus fine sur V4 vient de **Together AI** (mai 2026) : l'attention hybride
impose de gérer **trois layouts de KV cache simultanés** — CSA compressé, HCA dense (~8K entrées), SWA
exact. Capacité annoncée : **~3,7M tokens par nœud HGX B200** avec politique de cache optimisée. C'est
concret : servir V4 à 1M de contexte n'est pas qu'une question de poids, c'est une question de moteur
d'inférence capable de jongler avec trois formats de cache hétérogènes — d'où les forks vLLM/SGLang
maison de DeepSeek, DeepGEMM + FlashInfer + DeepEP, et l'expert-parallelism sur NVLink (NVL72). Pour un
hébergeur tiers, supporter V4 correctement demande un travail d'intégration réel, pas un simple
`vllm serve`.

## 178. DeepSeek : la concurrency limitée à 500 requêtes

Détail commercial qui en dit long : l'API **V4-Pro limite la concurrence à 500 requêtes simultanées**.
Ce n'est pas une limite technique du modèle, c'est du **capacity management** : DeepSeek rationne
l'accès à son datacenter au lieu de faire la queue. Combiné à la tarification peak/off-peak, ça dessine
un fournisseur qui gère la pénurie plutôt que l'abondance — paradoxal pour le champion des prix bas.
En pratique : pour un workload à fort parallélisme (batch, évaluations), prévoir le retry/backoff et,
surtout, des fournisseurs tiers (OpenRouter, DeepInfra, Novita) qui mutualisent leurs propres quotas.

## 179. Qwen : Megatron-Bridge et le paradoxe NVIDIA

La liste de déploiement de Qwen3.5 inclut **Megatron-Bridge (NVIDIA NeMo)** — le tooling d'entraînement/
inférence de NVIDIA — et NVIDIA a mis Kimi K2.6 « on a faster path to Blackwell inference » (variante
NVFP4 du 13 mai 2026). Paradoxe apparent : pendant que la Chine se découple du silicium NVIDIA (Ascend,
ASIC domestiques), **le software stack NVIDIA reste la lingua franca** du serving des modèles chinois
ouverts (vLLM, SGLang, TensorRT-LLM, NeMo). Le découplage est matériel, pas logiciel — et c'est logique :
un modèle ouvert a intérêt à tourner partout, y compris sur GPU NVIDIA. La vraie ligne de fracture n'est
pas « NVIDIA vs Chine », c'est « training domestique possible sans NVIDIA » (GLM-5, LongCat-2.0).

## 180. Qwen : 201 langues et un vocabulaire de 250K

Qwen3.5 revendique **201 langues** et un vocabulaire de **~250K tokens** — parmi les plus larges du volume.
C'est un choix de tokenization qui coûte cher (la table d'embedding grossit) mais qui paie sur le
multilingue et sur les langues à morphologie riche. À comparer : Kimi K2.6 (vocab 163 840), MiniMax-M2
(200 064). Pour un usage francophone avec documents multilingues (le cas typique d'un RAG d'entreprise
en Afrique de l'Ouest, par exemple), un gros vocabulaire multilingue réduit la fragmentation des tokens
— donc le coût et la latence. C'est un critère sous-coté face aux benchmarks anglophones.

## 181. Kimi : la température fixée à 1,0

Détail d'API qui surprend : sur l'API Moonshot, la **température de Kimi K2.6 est fixée à 1,0**, non
modifiable. C'est un choix assumé pour les workloads agentiques : à température 1,0, le modèle explore
plus (diversité des trajectoires), ce qui est souhaitable quand on lance des essaims d'agents ou des
rollouts RL — mais moins quand on veut des réponses déterministes (extraction, classification). Le
raisonnement est retourné dans le champ **`reasoning_content`** séparé, ce qui permet de le logger ou de
l'afficher sans polluer la réponse. Deux décisions produit qui montrent que Moonshot designe pour les
agents d'abord, le chat ensuite.

## 182. Kimi : Cloudflare Workers AI, la distribution edge

Kimi K2.7 Code était **disponible sur Cloudflare Workers AI dès le jour J** (`@cf/moonshotai/kimi-k2.7-code`,
12 juin 2026). C'est un canal de distribution sous-estimé : servir un MoE 1T/32B via le réseau edge de
Cloudflare, c'est rapprocher l'inférence de l'utilisateur sans gérer de GPU. Pour Moonshot, c'est aussi
une façon de contourner la barrière des 1,4 To de poids du K3 : l'utilisateur n'a jamais besoin de
télécharger. La leçon distribution : en 2026, **un modèle ouvert se juge aussi à son graphe de
distribution** (API first-party, OpenRouter, Cloudflare, Fireworks, together...) — la disponibilité
multi-canal vaut presque autant que les benchmarks.

## 183. GLM : les deux semaines de retenue (sécurité)

