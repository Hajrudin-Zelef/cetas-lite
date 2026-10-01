---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-21
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Apple", "DeepSeek", "Huawei", "LongCat", "Meituan", "MiniMax", "Moonshot", "Nvidia", "SGLang", "Xiaomi", "Z.ai", "vLLM"]
dates: ["2026-08-05"]
keywords: ["agent", "amd", "apache", "ascend", "attention", "aws", "benchmarks", "compute", "cyber", "datacenter", "deepseek", "exploit"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [2064, 2166]
sha256: f491912055827ef1e771ec02ec510e638c30c3d80fa449dc0f05e013560beebf
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

Zhipu a **retenu les poids du GLM-5.3 deux semaines** (14 → 28 août 2026) pour « durcissement sécurité »
avant publication. C'est la première fois qu'un labo chinois ouvert admet publiquement un délai
sécurité sur une release — et c'est lié au positionnement « cyber » du modèle (découverte de vulns,
exploit scripting). Le message implicite : à partir d'un certain niveau de capacité offensive, la
publication immédiate n'est plus tenable, même en MIT. C'est un précédent : on peut s'attendre à ce que
les futurs modèles « cyber-capables » subissent des délais similaires, voire des restrictions. Pour la
veille : une annonce sans poids immédiats n'est plus un retard, c'est peut-être une revue sécurité.

## 184. GLM : le Coding Plan (Lite/Pro/Max/Team)

Zhipu vend GLM-5.2/5.3 via le **GLM Coding Plan** (tiers Lite/Pro/Max/Team) en plus de l'API au token
($1,40/$4,40 par 1M). C'est le modèle économique « abonnement développeur » (forfait mensuel pour les
usages de code intensifs) face au « pay-as-you-go » de DeepSeek et au « Token Plan » d'Alibaba. Trois
philosophies de monétisation coexistent donc en Chine : le token spot (DeepSeek), l'abonnement dev
(Zhipu), la plateforme entreprise (Alibaba/QwenWork). Pour un usage régulier et prévisible (une équipe
dev), le forfait bat le token ; pour un usage sporadique ou batch, c'est l'inverse. Le choix du plan
tarifaire est devenu un vrai sujet d'architecture.

## 185. MiMo : 7 plateformes de chips et 100T tokens gratuits

Le lancement du MiMo-V2.5 (22 avril 2026) cumulait deux gestes industriels : **adaptation jour 1 à 7
plateformes de chips chinoises** (T-Head, Kunlun, Enflame, Muxi, Tianshu Zhixin...) + AWS, et un programme
d'incitation de **100T tokens gratuits pendant 30 jours**. Le premier geste dit : l'inférence domestique
est une priorité produit, pas un portage après-coup. Le second dit : Xiaomi achète des parts de marché
API au prix du compute — 100T tokens, c'est de quoi faire tourner une PME entière pendant un mois. Les
deux ensemble : une stratégie d'écosystème où le modèle est un produit d'appel pour la plateforme (MiMo
Open Platform, Studio, Desktop). C'est l'équivalent chinois du « free tier agressif » des hyperscalers.

## 186. MiMo : l'encodeur vision 681M et le décodeur MTP 5 couches

Zoom technique sur le V2.6-Pro-RL : **encodeur vision de 681M paramètres** (conséquent — c'est la taille
d'un petit LLM à lui seul), **encodeurs audio** dédiés, et **décodeur MTP à 5 couches**. L'omnimodalité
« native » a un coût architectural réel : près d'un milliard de paramètres juste pour voir et entendre,
avant même le backbone MoE (1,02T/42B, 384 experts routés dont 8 actifs par token, 70 couches en attention
hybride sliding-window + globale). Le décodeur MTP à 5 couches, lui, sert le décodage spéculatif — la
vitesse est pensée dès l'architecture. C'est un modèle « système » : chaque modalité et chaque
optimisation a son budget de paramètres propre.

## 187. Ling : l'adoption infra (vLLM PR #51045, parsers SGLang)

Le Ling-3.0-flash a obtenu le **support vLLM via la PR #51045 (mergée le 2026-08-05)** et des **parsers
dédiés SGLang (`ling3` reasoning/tool-call)**. Deux semaines après la sortie des poids (23 juillet), le
modèle était servable sur les deux moteurs de référence — c'est le rythme « day-0 » devenu standard pour
les gros ouverts chinois. Le point technique : supporter l'hybride KDA + MLA + 512 experts + MTP dans
vLLM/SGLang n'est pas trivial (kernels d'attention linéaire, routage d'experts, speculative decoding) ;
le fait que ce soit fait en deux semaines montre la maturité de l'écosystème d'intégration autour des
modèles chinois. Pour l'ops : vérifier la version minimale (vLLM post-PR #51045) avant de planifier un
déploiement.

## 188. Ling : le DGX Spark comme cible edge

Ant documente le déploiement du Ling-3.0-flash sur **DGX Spark (GB10)** en INT4 (72 Go) et valide le
tiny en BF16 sur la même machine ; le runbook communautaire Hy3 tourne en NVFP4 sur 2× DGX Spark
(~21,8 tok/s single-stream). Le DGX Spark — la « machine IA de bureau » NVIDIA — est devenu en 2026 la
**référence edge** des labos chinois : si ton 124B passe en INT4 sur un GB10, tu as un argument
« workstation » crédible. C'est un changement de paradigme discret : les frontier ouverts ne sont plus
réservés au datacenter, ils visent le bureau d'ingénieur. Pour un RAG local exigeant, le couple
« DGX Spark + Ling-3.0-flash INT4 » est l'option la plus sérieuse du volume côté chinois.

## 189. LongCat : dNaViT, la voix 24 kHz et la vidéo absente

Zoom sur le LongCat-Next : tokenizer visuel **dNaViT** (RVQ à 8 couches — voir glossaire), tokenizers
audio RVQ natifs, **synthèse vocale et clonage de voix en 24 kHz streaming**, compréhension + **génération
d'images**. Et une absence notable : **pas de génération vidéo** (compréhension seule) — là où MiniMax-H3
fait de la vidéo+audio native 4–15 s à 24 FPS. Deux écoles du multimodal chinois : Meituan fait du
« discret autoregressif unifié » (DiNA) avec des moyens contenus (~70B), MiniMax fait du « gros
transformer génératif » (33B dense vidéo+audio). Le premier est déployable (build w8a8_int8 ~90 Go sur
DGX Spark), le second est spectaculaire. Selon le besoin (agent vocal vs studio vidéo), le choix est
tranché.

## 190. MiniMax : la recette de sampling (et pourquoi elle compte)

MiniMax publie ses réglages recommandés — fait rare et précieux : **M2 : temperature=1.0, top_p=0.95,
top_k=40** ; **M3 : temperature=1.0, top_p=0.95, sans top_k**. La température 1,0 (comme Kimi) est le
réglage « agentique » : diversité maximale, pas de réponses tièdes. Le top_k=40 du M2 resserre
l'échantillonnage sur les 40 tokens les plus probables ; son abandon sur M3 suggère que le post-training
a rendu le modèle assez calibré pour s'en passer. Retenir : **les benchmarks d'un labo sont mesurés avec
ses réglages** — servir un M2 à température 0,2 « pour être sûr » et comparer aux chiffres officiels
n'a aucun sens. Toujours partir des réglages recommandés, puis ajuster.

## 191. MiniMax : ATOM, ROCm et l'inférence AMD

La liste de déploiement MiniMax inclut **ATOM (ROCm)** — le stack AMD — aux côtés de vLLM, SGLang,
KTransformers et MLX-LM (Apple Silicon). C'est le seul labo du volume à citer explicitement ROCm dans ses
supports, et ça compte : l'inférence des ouverts chinois n'est pas un monopole CUDA. Entre ROCm (AMD),
les NPU Ascend (tutoriel Hy3), les 7 plateformes chinoises (MiMo-V2.5) et MLX (Apple Silicon), les modèles
ouverts chinois de 2026 sont les plus **portables en silicium** du marché. Pour une stratégie
d'infrastructure anti-lock-in, c'est un argument de poids face aux modèles fermés calés sur une seule
plateforme.

## 192. Tencent : le support day-0 (vLLM hy_v3, SGLang FP8)

Hy3 est arrivé avec un **support day-0** exemplaire : **SGLang** (FP8 + speculative decoding) et
**vLLM ≥ 0.23** avec architectures natives `hy_v3` / `hy_v3_mtp`. Des archs nommées dans vLLM, c'est le
signe d'une intégration poussée (pas un simple mapping de config). Ajouter les variantes communautaires
NVFP4 (`kodelow/Hy3-NVFP4-W4A16`, 181 Go) et le **tutoriel SGLang officiel sur Ascend NPU** (910B, 2×
Atlas 800I A2) : Tencent a soigné les trois couches — NVIDIA (SGLang/vLLM), communauté (NVFP4), domestique
(Ascend). C'est la check-list du lancement ouvert réussi en 2026 : poids + licence claire (Apache 2.0) +
moteurs day-0 + hardware domestique. Les labos qui ratent une case (poids « coming soon » de LongCat-2.0,
licence non vérifiée du K3) paient en adoption.

## 193. Face-à-face : les scores SWE-bench (code)

