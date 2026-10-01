---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-18
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "Huawei", "LongCat", "Meituan", "Moonshot", "Nvidia", "OpenAI", "Together AI", "Xiaomi", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["agent", "agents", "ascend", "asic", "attention", "benchmarks", "claude", "datacenter", "decode", "deepseek", "distribution", "glm"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1742, 1844]
sha256: d2bc1d4c648628e04c01010b09ca8d67f273341124168c53f426539fa464e544
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

Fait remarquable pour le labo le plus influent du volume : **aucun rapport technique officiel DeepSeek V4
n'a été trouvé** au 27/09/2026. Tout ce que l'on sait de l'attention hybride (CSA+HCA+SWA), de MegaMoE,
des résiduels « mHC », de l'optimiseur Muon et de l'« Engram memory » 196B du V4.1-Flash vient de notes
communautaires GitHub, de documentations d'hébergeurs (Together AI) et de model cards. DeepSeek publie les
poids (MIT) mais garde le « comment » pour lui — l'inverse d'Anthropic/OpenAI qui publient des system
cards sans les poids. C'est la doctrine DeepSeek 2026 : **l'open-weight comme distribution, pas comme
documentation**. Pour la recherche, c'est frustrant ; pour l'utilisateur, c'est transparent — le modèle
parle de lui-même.

## 156. Qwen : le snapshot comme méthode de release

Alibaba a normalisé en 2026 la release par **snapshot daté** : Qwen3.8-Max-0902 succède au snapshot du
3 août sans changement de nom commercial, mais avec un AA Index qui passe de 40 à 45, une 1re place à
Code Arena WebDev, et un coût par tâche qui double ($2,67 → $5,41, ~108k tokens/tâche). Le snapshot n'est
ni un patch ni une version majeure : c'est un **recalibrage continu** (post-training, réglages de
reasoning). Conséquences pratiques : (1) les benchmarks ont une date de péremption de quelques semaines ;
(2) la facture API d'un même ID peut varier fortement d'un mois à l'autre ; (3) en production, il faut
épingler le snapshot et rejouer les évaluations à chaque changement. DeepSeek fait pareil (« V4 Pro 0813 »),
mais Alibaba l'a industrialisé.

## 157. Qwen : QwenWork, Wukong et la conquête de l'entreprise

Les modèles ne sont que la moitié de la stratégie Alibaba : **QwenWork** (plateforme de travail, citée au
déploiement du 3.8-Max), **Wukong** (plateforme entreprise, intégrée au 3.6-Plus) et l'**app Qwen**
grand public forment un étau entreprise. Le 3.8-Max est explicitement positionné « workplace tools » par
la presse (marketbusinessnews, août 2026), et le snapshot 0902 vise le coding d'entreprise (1er Code Arena
WebDev). La mention « China state-law risk » dans un article TechTimes sur QwenWork rappelle l'enjeu
compliance : un modèle d'entreprise chinois opère sous droit chinois — point de vigilance pour un
déploiement hors de Chine (données, juridiction). Alibaba ne vend plus des modèles : il vend une suite
bureautique IA.

## 158. Kimi : Mooncake et la disaggregation

Le support day-0 du Kimi K3 cite **Mooncake**, le système d'inférence de Moonshot : **réutilisation du KV
cache cross-instance** et **disaggregation PD/EPD** (séparer le prefill — traitement du prompt — du
decode — génération —, voire l'encoding). L'idée : dans une ferme servant des milliers de requêtes, les
préfixes de contexte se répètent (system prompts, documents RAG) ; Mooncake les partage entre instances
au lieu de les recomposer. Combiné au KDA (état borné), c'est l'arme du 1M « plusieurs fois moins cher ».
C'est aussi un signal : en 2026, **l'inférence est un problème de systèmes distribués**, pas seulement de
poids — et les labos chinois investissent dans la couche serving autant que dans les modèles.

## 159. Kimi : l'essaim de 300 agents

La signature produit de Kimi K2.6 : des **essaims jusqu'à 300 agents** coordonnés sur **4 000 étapes**
(revendication Moonshot, relayée par Spheron). Ce n'est pas du marketing creux : c'est le prolongement
direct du design « agent natif » — un modèle entraîné au tool use multi-acteurs, avec un champ
`reasoning_content` dédié et une température API fixée à 1,0 (pour la diversité des trajectoires). K2.7
Code pousse la logique : sessions de **>12 h en continu**, thinking forcé, -30 % de tokens de
raisonnement sur le code. La question ouverte : qui orchestre 300 agents en production réelle ? Au
27/09/2026, c'est une capacité démontrée plus qu'un usage banalisé — mais elle fixe le plafond de ce que
l'open-weight sait faire en 2026.

## 160. GLM : ZCode, le harnais maison

Avec GLM-5.2, Zhipu livre **ZCode**, son harnais agent officiel « style Claude Code », compatible dès le
jour J avec Claude Code, Cline, OpenCode, Roo Code, Goose, Crush, OpenClaw et Kilo Code. La stratégie est
double : fournir l'outil de référence (pour que l'expérience GLM soit optimale) **et** garantir
l'interopérabilité (pour capter les utilisateurs des harnais existants). C'est la même logique que la
compatibilité « protocole Anthropic » de Qwen (section 23) : en 2026, **le harnais est un champ de
bataille** — celui qui contrôle la boucle agent contrôle la distribution du modèle. Zhipu est le seul
labo chinois du volume à livrer harnais + modèle + Coding Plan (Lite/Pro/Max/Team) en pack intégré.

## 161. GLM-5 (février 2026) : le fondateur hors-période

Antérieur à la période du volume (sortie : **11 février 2026**), GLM-5 mérite sa fiche car il fonde la
lignée 5.2/5.3. **MoE 744B total / 40B actifs**, 28,5T tokens de pré-entraînement, **DeepSeek Sparse
Attention (DSA)**, contexte **200K** (137K selon une table tierce — divergence), **licence MIT**.
**Entraîné entièrement sur chips Huawei Ascend (MindSpore), zéro NVIDIA** — fait géopolitique majeur,
corroboré par Reuters via Maxime Labonne. Benchmarks (vendor) : SWE-bench Verified **77,8 %** (SOTA
open-source au lancement), Terminal Bench 2.0 56,2, AIME 2026 92,7 %, GPQA-Diamond 86,0 %. GLM-5.2/5.3 en
sont l'évolution directe (même base pré-entraînée pour 5.2→5.3, passage à 1M de contexte, IndexShare).
Sans GLM-5, pas de « moment Zhipu » de l'été 2026.

## 162. MiMo : Studio, Desktop et l'Open Platform

Xiaomi distribue ses modèles via trois canaux : l'**API MiMo Open Platform** (facturation au token),
**MiMo Studio** (chat web) et **MiMo Desktop** (client local). Le V2.6-Pro-RL demande **2 nœuds / 32 GPU**
selon la commande de serving Xiaomi — c'est du datacenter, pas du bureau. À l'autre bout : le
**Distill-Qwen-9B** tourne sous Ollama/LM Studio avec 16 Go VRAM. Entre les deux, le Flash-RL (309B/15B)
est le compromis servable. Cette échelle de distribution (API → datacenter self-hosted → bureau) est la
plus complète du volume avec Qwen : Xiaomi pense déjà en « gamme de déploiement », pas en modèle unique.

## 163. Ling : Fin, anatomie d'un vertical

Le **Ling-3.0-flash-Fin** (9 septembre 2026) est le seul modèle **vertical métier** du volume : base
3.0-flash (124B/5,1B) fine-tunée finance, co-développée avec des institutions financières (CICC investment
banking). Cas d'usage revendiqués : recherche investissement, conformité, modélisation de valorisation,
reporting réglementaire. Benchmarks dédiés (Ant) : FinFIRST, FinSearchComp Verified, FinCRAFT,
FinanceAgent v1.1/v2, APEX-Agents, SpreadsheetBench v1/v2, τ³-Banking. Deux lectures : (1) la
différenciation 2026 se joue sur les déclinaisons sectorielles autant que sur les généralistes ; (2) Ant,
adossé à un empire fintech, fait ce que personne d'autre ne peut faire aussi crédiblement — transformer
son accès aux workflows financiers en données de post-training. Licence exacte : non vérifiée.

## 164. LongCat : 50 000 ASIC et le pari domestique

Le **LongCat-2.0** (30 juin 2026) revendique un entraînement et un serving sur **supercalculateurs d'ASIC
chinois domestiques** (cluster ~50 000 cartes — claim Meituan), avec « Super kernels » et L2-cache weight
prefetching maison. C'est le cas le plus extrême du découplage : non seulement sans NVIDIA, mais sur une
infrastructure conçue autour du modèle (1,6T paramètres, 1M natif, LongCat Sparse Attention). Scores
revendiqués au niveau GPT-5.5 (SWE-bench Pro 59,5 vs 58,6). Réserves : benchmarks vendor-reported,
disponibilité effective des poids non vérifiée au 27/09/2026. Mais le signal stratégique est clair —
même un acteur « non-IA » comme Meituan peut monter une pile complète domestique.

