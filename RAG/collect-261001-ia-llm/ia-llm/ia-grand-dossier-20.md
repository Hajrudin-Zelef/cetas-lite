---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-20
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "OpenAI", "Stability AI", "xAI"]
dates: []
keywords: ["agents", "apache", "arr", "benchmarks", "chatgpt", "claude", "compute", "deepseek", "diffusion", "distillation", "embeddings", "fine-tuning"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1281, 1361]
sha256: 3977baea44d629255b810dbe2a34ec08cf87c2483019741e309d25cc3d83843f
---

# IA — Le grand dossier

| Date | Modèle | Labo | Paramètres | Poids | Pourquoi ça compte |
|---|---|---|---|---|---|
| 06/2018 | GPT-1 | OpenAI | 117M | Fermés | Pré-entraînement génératif + fine-tuning |
| 10/2018 | BERT-Large | Google | 340M | Ouverts | L'encodeur roi 2018-2021 ; ancêtre des embeddings |
| 02/2019 | GPT-2 | OpenAI | 1,5B | Fermés→ouverts (fin 2019) | Génération cohérente ; la « staged release » |
| 05/2020 | GPT-3 | OpenAI | 175B | Fermés (API) | Few-shot ; 3,14×10^23 FLOP ; ~4,6 M$ |
| 01/2021 | DALL-E | OpenAI | 12B | Fermés | Texte→image par Transformer |
| 01/2022 | InstructGPT | OpenAI | 1,3B-175B | Fermés | RLHF : 1,3B préféré à 175B |
| 03/2022 | Chinchilla | DeepMind | 70B | Fermés | 70B/1 400B tokens bat Gopher 280B |
| 04/2022 | DALL-E 2 | OpenAI | « à vérifier » | Fermés | Saut photoréaliste |
| 07/2022 | Midjourney v3 | Midjourney | « à vérifier » | Fermés | Le produit image grand public |
| 08/2022 | Stable Diffusion 1.x | Stability AI / LMU | ~1B (UNet) | **Ouverts** | La démocratisation de l'image |
| 11/2022 | ChatGPT (3.5) | OpenAI | « à vérifier » | Fermés | 100M utilisateurs en 2 mois |
| 02/2023 | LLaMA | Meta | 7B-65B | Ouverts (fuite) | L'événement open-weight |
| 03/2023 | GPT-4 | OpenAI | Non publié | Fermés | Multimodal ; zéro détail publié |
| 03/2023 | Claude 1 | Anthropic | Non publié | Fermés | Constitutional AI en produit |
| 07/2023 | Llama 2 (+Chat) | Meta | 7B-70B | Ouverts (licence comm.) | RLHF documenté en ouvert |
| 07/2023 | Claude 2 | Anthropic | Non publié | Fermés | Contexte 100k |
| 09/2023 | Mistral 7B | Mistral | 7,3B | **Apache 2.0** | Efficacité européenne |
| 12/2023 | Gemini 1.0 Ultra | Google DeepMind | Non publié | Fermés | Nativement multimodal |
| 12/2023 | Mixtral 8x7B | Mistral | 47B (13B actifs) | **Apache 2.0** | MoE en ouvert |
| 02/2024 | Gemini 1.5 Pro | Google DeepMind | Non publié | Fermés | Contexte 1M tokens |
| 02/2024 | Sora | OpenAI | Non publié | Fermés | Vidéo 1 minute réaliste |
| 04/2024 | Llama 3 (8B/70B) | Meta | 8B / 70B | Ouverts (licence comm.) | Sur-entraînement assumé |
| 05/2024 | GPT-4o | OpenAI | Non publié | Fermés | Omni : texte+image+audio natif |
| 07/2024 | Llama 3.1 405B | Meta | 405B | Ouverts (licence comm.) | 15 000B tokens ; ~37 tok/param |
| 09/2024 | o1-preview / o1-mini | OpenAI | Non publié | Fermés | **Test-time compute en produit** |
| 09/2024 | Qwen2.5 (dont Coder 32B) | Alibaba | 0,5B-72B | **Apache 2.0** | Le champion open du code |
| 12/2024 | o1 (complet) | OpenAI | Non publié | Fermés | 83 % AIME 2024 |
| 12/2024 | DeepSeek-V3 | DeepSeek | 671B (37B actifs) | **MIT** | 5,576 M$ le run ; MLA+MoE+FP8 |
| 12/2024 | Gemini 2.0 Flash | Google DeepMind | Non publié | Fermés | Ère « agentique » |
| 01/2025 | DeepSeek-R1 | DeepSeek | 671B (37B actifs) | **MIT** | **Niveau o1 en ouvert** ; R1-Zero (pur RL) |
| 01/2025 | Qwen2.5-Max | Alibaba | Non publié | API | — |
| 02/2025 | Grok 3 | xAI | Non publié | Partiel | Cluster Colossus |
| 04/2025 | o4-mini | OpenAI | Non publié | Fermés | Petit raisonnant efficace |
| 04/2025 | Llama 4 (Scout/Maverick) | Meta | Non publié (« à vérifier ») | Ouverts (licence comm.) | Multimodal natif |
| 05/2025 | Veo 3 | Google DeepMind | Non publié | Fermés | Vidéo + audio natif |
| 06/2025 | Magistral | Mistral | 24B (« à vérifier ») | **Apache 2.0** | Raisonnement en ouvert européen |
| 08/2025 | GPT-5 | OpenAI | Non publié | Fermés | « Intelligence unifiée » ; lancement mitigé |
| 09/2025 | Qwen3 | Alibaba | 0,6B-235B (MoE) | **Apache 2.0** | Hybride thinking/non-thinking |
| 2026 | DeepSeek-V4 | DeepSeek | « à vérifier » | **MIT** (attendu) | -75 % prix API définitifs (mai 2026) |
| 2026 | GPT-5.5, Claude Opus 4.7, Mistral Large 3, Kimi K2... | Divers | Non publié | Mixte | Versions « à vérifier » au jour près |

**Lecture de la frise :** 2018-2020 = le texte se scale ; 2021-2022 = l'image bascule (diffusion) + l'alignement (RLHF) ; 2023 = l'ouverture contre-attaque ; 2024 = le raisonnement et le multimodal ; 2025 = l'efficacité et les agents ; 2026 = l'inférence comme champ de bataille économique.

---

## 4.8. Les mathématiques du ralentissement : comprendre les exposants

> Pour ceux qui veulent vérifier les affirmations des deux camps par le calcul. Niveau : terminale scientifique.

**Les exposants de Kaplan (2020).** La perte décroît en loi de puissance :

```
L(N) = (Nc/N)^αN   avec αN ≈ 0,076
L(D) = (Dc/D)^αD   avec αD ≈ 0,095
L(C) = (Cc/C)^αC   avec αC ≈ 0,050
```

Interprétation : **multiplier les paramètres par 10** (à données fixées) réduit la perte d'un facteur 10^0,076 ≈ **1,19** (soit -19 % de perte). Multiplier le compute par 10 : facteur 10^0,05 ≈ **1,12** (-12 %). Les gains sont **logarithmiques en échelle linéaire** : pour diviser la perte par 2, il faut multiplier le compute par ~10^6 (un million de fois) ! C'est la raison mathématique des rendements décroissants : **la courbe est concave dès le départ**, on l'a juste masquée en regardant des benchmarks (échelle non linéaire) plutôt que la loss.

**Pourquoi les benchmarks ont masqué le ralentissement.** La loss décroît doucement, mais la **performance sur tâche** suit souvent une **sigmoïde** : longtemps plate, puis brusque montée (« émergence »), puis plateau. Entre 2020 et 2023, on était sur la pente raide de plusieurs sigmoïdes (MMLU, GSM8K...) — d'où l'impression d'accélération. En 2024-2025, ces benchmarks **saturent** (> 90-95 %) : on retombe sur la pente douce de la loss. Le « mur » ressenti est en partie un **artefact de saturation des instruments de mesure**, pas seulement un arrêt du progrès sous-jacent.

**Le calcul du mur des données (Epoch AI, simplifié).**
Stock estimé : S ≈ 3×10^17 tokens (300 000 Md). Consommation annuelle des labos frontières (2024-2025) : de l'ordre de 10^16-10^17 tokens/an cumulés (« à vérifier » l'estimation exacte, l'ordre de grandeur est celui des publications Epoch AI). Durée de vie : S / consommation ≈ **3 à 10 ans** → 2026-2032 selon les hypothèses de croissance. Les leviers qui repoussent l'échéance : répétition des données (rendement exponentiellement décroissant — Muennighoff 2023), données synthétiques filtrées, multimodal (la vidéo YouTube seule représente des ordres de grandeur au-delà du texte — « à vérifier » la quantification exacte).

**Le calcul qui justifie le test-time compute (Snell et al., 2024, simplifié).**
Si accuracy(c) ≈ 1 - A·c^(-β) avec β ≈ 0,3-0,5 (selon tâche), alors multiplier le compute d'inférence par 100 multiplie (1 - accuracy) par 100^(-0,4) ≈ 1/6,3 : **l'erreur est divisée par ~6 pour ×100 de compute**. Comparez au training : pour le même gain via Kaplan (αC ≈ 0,05), il faudrait ×10^16 de compute d'entraînement — impossible. **C'est le calcul qui a convaincu l'industrie en 2024 : à gain égal, l'inférence est infiniment plus rentable que le training une fois les rendements du training épuisés.**

**La synthèse en une formule (formulation de l'auteur) :**
```
Progrès ≈ f(training_compute) + g(inference_compute) + h(efficacité_algo)
```
- 2012-2023 : f dominait.
- 2024-2026 : f s'aplatit (rendements décroissants + mur des données), **g et h prennent le relais**.
- Le débat « mur vs continuation » = « f est-il tout ? » Non. « g+h compensent-ils ? » Jusqu'ici : oui (o1, R1, distillation).

---

## 2.6. Où sont-ils maintenant ? (tableau de suivi des figures clés)

