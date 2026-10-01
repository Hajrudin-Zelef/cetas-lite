---
id: collect-261001-ia-llm/ia-llm/ia-modeles-chine-15
title: "ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Huawei", "Hugging Face", "LongCat", "Meituan", "MiniMax", "Moonshot", "Nvidia", "SGLang", "Z.ai"]
dates: []
keywords: ["agent", "agents", "ascend", "asic", "attention", "benchmarks", "cyber", "deepseek", "glm", "kimi", "kv cache", "moe"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_chine.md
source_anchor: ""
source_lines: [1351, 1468]
sha256: 2719dc76d1956400f88930808e828f9469806b09301ef33d75068214ba336879
---

# ENCYCLOPÉDIE DES MODÈLES IA — VOLUME 1 : LA CHINE

Le fil rouge technique du volume : **tout le monde a tué l'attention quadratique**. Gated DeltaNet
(Qwen), KDA (Kimi, réutilisé par Ant), DSA + IndexShare (Zhipu), attention hybride CSA/HCA/SWA (DeepSeek),
LongCat Sparse Attention (Meituan) — cinq noms pour une idée : mélanger couches linéaires à état borné
et couches d'attention sparse/dense pour un coût mémoire qui n'explose plus avec la séquence. C'est cette
brique, plus que le MoE, qui a rendu le 1M généralisé et les prix cassés possibles. Et elle circule :
Ant reprend le KDA de Moonshot en quelques semaines. L'open-weight chinois fonctionne comme un
« commons » d'architectures où l'innovation se diffuse à la vitesse des releases — l'inverse du modèle
fermé occidental où chaque labo garde sa recette.

## 103. Le pari des puces chinoises

Quatre faits convergent : GLM-5 entraîné **entièrement sur Huawei Ascend** (zéro NVIDIA, corroboré par
Reuters) ; GLM-5.3-Flash servi depuis **100 000 puces de fabrication chinoise** ; LongCat-2.0 entraîné
sur **~50 000 ASIC domestiques** ; MiMo-V2.5 adapté **jour 1 à 7 plateformes de chips chinoises** ;
tutoriel SGLang officiel pour Hy3 sur **Ascend 910B**. Ajouter le lancement conjoint puce+modèle
d'Alibaba (Zhenwu M890 + Qwen3.7-Max). La pile IA chinoise se découple méthodiquement de NVIDIA, du
silicium au serving. Pour l'observateur : les modèles ouverts chinois de 2026 sont déjà pensés pour
tourner sur silicium domestique — c'est un critère de robustesse géopolitique, pas un détail.

## 104. Où va Qwen4 (et la suite)

Le Qwen3.8-Flash est présenté par Alibaba comme une **« preview explicite de l'architecture Qwen4 »**
(Gated DeltaNet + Qwen Sparse Attention, table N-gram 51B). Le schéma habituel d'Alibaba : teaser
l'architecture sur un modèle efficace ouvert, puis décliner en flagship. Avec DeepSeek V4.1 Pro attendu
imminemment (rumeur 28–30 septembre 2026), Kimi K3 installé comme référence ouverte et GLM-5.3 en
embuscade « cyber », le Q4 2026 s'annonce comme un choc frontal : Qwen4 vs DeepSeek V4.1 Pro vs Kimi K4
(?). Rien de tout cela n'est vérifié — c'est de la lecture de trajectoire, pas de l'information.

## 105. Limites et incertitudes de ce volume

Rappel honnête de ce que ce volume ne garantit pas. Un, les **benchmarks sont vendor-reported** sauf
mention « indépendant » (AA Intelligence Index, LMArena, Artificial Analysis) : les chiffres constructeur
sont des documents marketing autant que techniques. Deux, les **prix API** viennent majoritairement de
sources tierces et bougent vite (les divergences GLM-5.3 : $1,40/$4,40 vs $0,50/$2,00 — en sont
l'exemple). Trois, les **architectures fines** (attention hybride DeepSeek, Engram 196B, actifs exacts du
K3) reposent sur des notes communautaires, pas des rapports officiels. Quatre, le corpus 2026 est
**pollué par du contenu IA-généré** qui se recopie : les dates et chiffres isolés sur une seule source
« ai-news » sont fragiles. Cinq, les **licences** doivent être vérifiées sur la carte Hugging Face avant
tout usage commercial (cas MiniMax-M2/M3, GLM-5.3, Qwen3.8-Max/Flash, Kimi K3). Ce volume est une carte
dressée en septembre 2026 : le terrain bouge chaque semaine.

## 106. Glossaire — mode d'emploi

Les 32 termes ci-dessous sont les briques de vocabulaire minimales pour lire les fiches sans se perdre.
Chaque terme est défini en une phrase dense, avec son contexte 2026 quand il compte. Les termes marqués
« voir section » renvoient aux fiches où ils sont illustrés.

## 107. Glossaire — MoE (Mixture-of-Experts)

Architecture où chaque token n'active qu'un sous-ensemble d'« experts » (petits réseaux FFN) au lieu de
tout le modèle : on peut avoir 2 800B paramètres totaux (Kimi K3) pour seulement 104B actifs par token.
Le ratio total/actifs (20–27× en 2026) est la mesure de la sparsité — et donc du coût réel d'inférence.

## 108. Glossaire — Paramètres actifs par token

Nombre de paramètres réellement utilisés pour générer un token. C'est lui, pas le total, qui détermine le
coût de calcul et la mémoire vive nécessaire : un Ling-3.0-flash (124B total / 5,1B actifs) se sert comme
un modèle ~5B, tout en « sachant » comme un 124B.

## 109. Glossaire — MLA (Multi-head Latent Attention)

Attention à cache latent compressé (popularisée par DeepSeek-V3) : au lieu de stocker clés et valeurs
pleines par tête, on stocke une représentation latente compacte. Divise le KV cache par ~10 à contexte
égal. Utilisée par DeepSeek, Kimi (K2.6, K3), Zhipu (GLM-5.x), Ant (Ling-3.0), Meituan (LongCat).

## 110. Glossaire — Attention linéaire (GDN, KDA)

Famille d'attentions en O(n) au lieu de O(n²) : Gated DeltaNet (Qwen), Kimi Delta Attention (Moonshot,
réutilisée par Ant). Elles remplacent le KV cache croissant par un **état récurrent de taille fixe**
(quelques dizaines à ~200 MiB). Le prix : une qualité légèrement différente sur les dépendances très
longues, d'où les architectures **hybrides** (mélange couches linéaires + attention standard).

## 111. Glossaire — DSA (Dynamic Sparse Attention) et IndexShare

Briques de Zhipu (GLM-5.2/5.3) : la DSA ne calcule l'attention que sur un sous-ensemble dynamique (top-k)
de positions passées ; **IndexShare** fait partager les mêmes indices top-k à des groupes de couches au
lieu de les recalculer. Objectif : diviser le coût à 1M de contexte. Voir section 48.

## 112. Glossaire — MTP (Multi-Token Prediction)

Tête auxiliaire qui prédit plusieurs tokens d'un coup au lieu d'un seul. Deux usages : (1) meilleur
apprentissage (signal plus riche), (2) **décodage spéculatif** — le modèle « brouillon » propose, le
modèle principal vérifie, ce qui accélère l'inférence (Ollama l'exploite en `draft-mtp`). Présent chez
Qwen, MiMo (décodeur MTP 5 couches sur V2.6-Pro), GLM (couche MTP améliorée), MiniMax-M2 (3 têtes),
Tencent Hy3 (+3.8B params MTP), NVIDIA Nemotron (hors volume).

## 113. Glossaire — KV cache

Mémoire où sont stockées les clés/valeurs des tokens déjà traités pour ne pas les recalculer. Son volume
croît avec la longueur du contexte : c'est **le** goulot du long contexte (ex. ~240 Go par 1M tokens pour
MiniMax-M2). Toutes les innovations 2026 (MLA, GDN/KDA, DSA, HiCache, Mooncake) visent à le réduire ou à
le réutiliser.

## 114. Glossaire — TTFT (Time To First Token)

Temps avant l'arrivée du premier token généré — la latence perçue par l'utilisateur. Critique pour les
agents (qui enchaînent les appels) : le HiCache d'Ant revendique -60 à -80 % de TTFT sur longues
conversations. Un modèle « rapide » en tok/s mais lent en TTFT reste désagréable en usage interactif.

## 115. Glossaire — Décodage spéculatif (speculative decoding)

Technique d'accélération : un petit modèle « draft » propose plusieurs tokens, le grand modèle les vérifie
d'un coup. Si le draft a raison souvent, gain de 1,5–2× en débit. Les têtes MTP servent de drafter intégré.
Supporté day-0 par SGLang sur plusieurs modèles du volume (Hy3, Kimi K3 via TokenSpeed).

## 116. Glossaire — YaRN

Méthode d'**extrapolation de contexte** : étendre au-delà de la fenêtre d'entraînement (ex. 256K → 1M)
en rééchelonnant les positions RoPE, sans réentraînement. Utilisée par Qwen (3.6, 3.8-27B, 3.8-Flash).
Réserve : un 1M « via YaRN » n'a pas la même qualité garantie qu'un 1M natif (entraîné à cette longueur).

## 117. Glossaire — Reasoning / thinking (modes)

Capacité du modèle à produire une chaîne de raisonnement avant la réponse. En 2026, c'est devenu
**réglable** : `reasoning_effort` (low/high/max chez GLM, Qwen), thinking désactivable (DeepSeek V4.1-Flash,
Ling-3.0), thinking forcé (Kimi K2.7 Code), trois niveaux `no_think`/`think_low`/`think_high` (Tencent Hy3).
Le « Thinking Preservation » (Qwen) conserve le raisonnement entre les tours d'un agent.

