---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-7
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "Samsung", "TensorRT-LLM", "United States", "vLLM", "xAI"]
dates: []
keywords: ["agents", "agi", "apache", "attention", "chatgpt", "claude", "compute", "deepseek", "distribution", "fp8", "gemini", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [399, 445]
sha256: f23ec6782dd6b0b762abb7f4c864dcfffd2f60cf2ed5a4eb5a802190150da348
---

# IA — Le grand dossier

**OpenAI — le tempo de l'industrie.**
Fondée en décembre 2015 comme non-profit (« construire une AGI bénéfique à toute l'humanité ») par Sam Altman, Elon Musk, Ilya Sutskever, Greg Brockman et d'autres. Le virage **capped-profit de 2019** (plafond de rendement pour les investisseurs) permet de lever les milliards nécessaires au compute : 1 Md$ de Microsoft en 2019, ~10 Md$ en janvier 2023. Produits : ChatGPT (nov. 2022), API GPT, DALL-E, Sora, puis la lignée **o** (raisonnement). GPT-4.5 « Orion » (mars 2025) illustre le mur : énorme, coûteux, non-raisonnant — déprécié quelques mois plus tard. GPT-5 (août 2025) tente l'« intelligence unifiée » avec un succès mitigé (voir 1.12). OpenAI reste le labo qui **donne le tempo**, mais son avance technologique s'est réduite : l'écart avec le n° 2 se mesure en mois, plus en années.

**Anthropic — la confiance comme produit.**
Fondée en 2021 par les frères Amodei et une douzaine d'ex-OpenAI. Levées massives auprès de Google et d'Amazon (plusieurs milliards — montants publics par tranches, « à vérifier » le cumul exact). Différenciation : **Constitutional AI**, **Responsible Scaling Policy** (publiée 2023, mise à jour ensuite : seuils de capacités déclenchant des mesures de sûreté), positionnement « l'IA sans pubs » (campagne Super Bowl 2026 anti-pubs dans l'IA, visant OpenAI). Claude est historiquement fort sur le contexte long, le code et l'« honnêteté » perçue. En 2026, xAI loue du compute à Anthropic (accord « Colossus » rapporté par la presse — « à vérifier ») : même les rivaux partagent l'infrastructure.

**Google DeepMind — l'empire intégré.**
Atouts uniques : **TPU maison** (v6/Trillium, Ironwood v7 — 2 à 4× plus efficaces que H100 sur charges Google selon des rapports internes relayés par la presse, « à vérifier » en indépendant), **données propriétaires** (YouTube pour le multimodal, Search pour le grounding), **distribution** (Android, Chrome, Workspace : des milliards d'utilisateurs). Gemini 1.5 (2024) a créé la catégorie « contexte 1M+ tokens » ; Gemini 2.0 (déc. 2024) a lancé l'ère « agentique » chez Google. Faiblesse historique : la lenteur produit (culture recherche), en cours de correction (rapprochement équipe Gemini app / DeepMind début 2025). Le Nobel 2024 de Hassabis donne au labo un prestige scientifique sans équivalent.

**Meta — l'open-weight comme stratégie.**
FAIR (2013, LeCun) a produit PyTorch — l'infrastructure logicielle de 90 % de la recherche mondiale (« à vérifier » la part, l'hégémonie est consensuelle). Avec **LLaMA (fév. 2023)** puis **Llama 2/3/4** en poids ouverts, Meta a choisi de **commoditiser le modèle** pour capter la valeur ailleurs (publicité, terminaux, données d'usage). Llama 3 405B (juillet 2024) : 15 000 milliards de tokens d'entraînement, ~37 tokens/paramètre — l'exemple canonique du **sur-entraînement volontaire** (voir 4.4). Pour un sysadmin/RAG : les Llama sont **les modèles self-hostables de référence**.

**xAI — le compute comme message.**
Fondée juillet 2023 par Musk après son départ d'OpenAI (2018) et le rachat de Twitter/X (2022). Grok est intégré à X (données temps réel du réseau social = différenciation). Le cluster **Colossus** à Memphis (100k+ H100 annoncés) est le symbole de la stratégie « scale hardware ». Positionnement idéologique : « IA qui dit la vérité », moins de garde-fous affichés — ce qui est autant un choix produit qu'un risque.

**DeepSeek — l'efficacité comme arme géopolitique.**
Voir 1.11-1.12 et 2.3. Points à ajouter : innovations publiées — **DeepSeekMoE** (experts partagés + routage fin), **MLA** (Multi-head Latent Attention : compression du cache KV, critique pour l'inférence longue), **Multi-Token Prediction**, **FP8** d'entraînement. Le rapport V3 (déc. 2024) détaille un coût de **5,576 M$** pour le run final — chiffre **du run final uniquement**, hors R&D et hors infra (les estimations « tout compris » montent à 1+ Md$ selon certains analystes — fourchette « à vérifier », à présenter comme telle). La leçon n'est pas « c'est pas cher », c'est : **l'écart d'efficacité entre labos est devenu un facteur géopolitique** (d'où les contrôles d'export US sur les H100/H800 vers la Chine, 2022-2023).

**Alibaba / Qwen — l'open chinois de référence.**
La famille **Qwen** (Tongyi Qianwen) est la plus adoptée hors Chine : Qwen2.5-Coder et Qwen3 rivalisent avec les meilleurs modèles ouverts en code et en multilingue, licence Apache 2.0. Alibaba joue la carte **écosystème développeur** (Hugging Face, Ollama, vLLM) — stratégie d'influence par l'adoption.

**Moonshot AI — le contexte long.**
Fondée en 2023 par Yang Zhilin (ex-Google Brain, Tsinghua). **Kimi** a été le premier chatbot grand public à proposer des contextes d'1M+ tokens côté chinois (2024). Spécialité : lecture de documents longs, agents de recherche.

**Mistral AI — le champion européen.**
Voir 2.3. Ajouts : **Mixtral 8x7B/8x22B** (MoE), **Codestral** (code), **Magistral** (raisonnement, 2025), **Le Chat** (assistant grand public, rebaptisé « Vibe » en 2026 selon la presse — « à vérifier »), partenariats industriels (ASML, Samsung, Airbus, BMW). Particularité : **Apache 2.0 véritable** (pas de licence « communautaire » restrictive comme Meta), ce qui en fait le choix open-weight le plus « propre » juridiquement pour une entreprise européenne.

**NVIDIA — le vendeur de pelles.**
Ne vend (presque) pas de modèles, vend **le terrain de jeu** : GPU + CUDA + bibliothèques (cuDNN, TensorRT, NCCL) + racks DGX + désormais du réseau (Mellanox/InfiniBand) et des « AI factories » clés en main. Marge brute ~75 % (« à vérifier » au trimestre près). Risque : la **désintermédiation par l'efficacité** (DeepSeek) et les chips maison des hyperscalers (TPU, Trainium, Maia).

### 3.3. Les dynamiques structurantes (2024-2026)

1. **Consolidation du talent vers les hyperscalers.** Suleyman (Microsoft), une partie de Character.AI (Google, 2024), Adept (Amazon, 2024 — « à vérifier » le périmètre exact de l'accord) : les labos indépendants peinent à financer le compute seuls.
2. **La guerre du compute.** Contrôles d'export US (H100/H800 → Chine, 2022-2023, renforcés ensuite), course aux datacenters (« Stargate » Microsoft/OpenAI, ~500 Md$ annoncés en janvier 2025 — montant d'annonce, « à vérifier » la réalisation), crise de l'énergie (les GPU arrivent dans des datacenters que le réseau électrique ne peut pas encore alimenter — voir 4.6).
3. **Ouvert vs fermé : le match nul.** Les poids ouverts (Llama, Qwen, DeepSeek, Mistral) atteignent 90-95 % des performances des fermés sur la plupart des tâches courantes ; les fermés gardent l'avantage sur la frontière (raisonnement extrême, multimodalité). Pour un RAG d'entreprise, **l'ouvert suffit presque toujours** — voir section 6.
4. **Le pivot inference.** Tous les labos déplacent l'investissement du training vers le serving (voir 4.7 et 6) : c'est là que se joue la marge désormais.

---

## 4. Le ralentissement du training / la fin du scaling naïf : le débat central

> C'est la section la plus importante de cette partie. Elle présente les **deux camps**, leurs arguments, leurs auteurs et leurs publications. Les chiffres sont sourcés ou marqués « à vérifier ». Les opinions sont attribuées.

### 4.1. Rappel : ce que disaient les lois de scaling

**Kaplan et al. (OpenAI, janvier 2020)** — « Scaling Laws for Neural Language Models » (arXiv:2001.08361). En entraînant des dizaines de modèles de tailles variées, l'équipe (Jared Kaplan, Sam McCandlish, Tom Henighan et al.) établit que la **perte (loss) suit des lois de puissance** en fonction de trois variables : nombre de paramètres N, volume de données D, compute C.

