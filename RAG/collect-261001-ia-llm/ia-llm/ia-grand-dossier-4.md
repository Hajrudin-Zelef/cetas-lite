---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-4
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Mistral", "Nvidia", "OpenAI", "United States"]
dates: ["2026-09-27"]
keywords: ["agents", "attention", "aws", "benchmarks", "blackwell", "claude", "compute", "deepseek", "distillation", "gemini", "gpu", "inference"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [214, 266]
sha256: 249fba013975cb5810c45f70d4589f7fcdfd0ea39501442443d968b2db467f6a
---

# IA — Le grand dossier

**Autres jalons 2023 :**

- **Claude 2** (Anthropic, juillet 2023), fenêtre de contexte 100k tokens — à l'époque, 10× la concurrence.
- **Gemini 1.0** (Google DeepMind, décembre 2023) : premier modèle « nativement multimodal » de Google, avec une version Ultra qui dépasse GPT-4 sur certains benchmarks (chiffres du rapport technique Google, à prendre avec les réserves d'usage sur les benchmarks maison).
- **Mistral 7B** (septembre 2023) puis **Mixtral 8x7B** (décembre 2023) : le labo français prouve qu'on peut rivaliser avec des modèles bien plus gros grâce à l'architecture **MoE (Mixture of Experts)** — voir glossaire.

### 1.11. 2024 : multimodalité, agents, et le tournant du raisonnement

**Multimodalité native.** GPT-4o (mai 2024, « o » pour omni) traite texte, image et audio dans un seul modèle avec une latence conversationnelle (~320 ms). Gemini 1.5 (février 2024) introduit une fenêtre de contexte d'**1 million de tokens**, puis 2 millions — on peut lui faire ingérer un livre entier, un codebase complet, des heures de vidéo.

**Les agents autonomes.** Avec l'**appel de fonctions (function calling)**, l'exécution de code et la navigation web, les LLM deviennent des **agents** : ils planifient, utilisent des outils, itèrent. 2024 voit Devin (Cognition, mars 2024, « premier ingénieur logiciel IA » — marketing à nuancer), les « computer use » (Anthropic, octobre 2024 : Claude pilote un ordinateur via captures d'écran), et OpenAI o1-preview. En 2025, les agents deviennent le principal vecteur de valeur (opérateurs, deep research).

**Septembre 2024 — o1 (« Strawberry »).** OpenAI lance **o1-preview** puis **o1** (décembre 2024) : des modèles entraînés en **reinforcement learning** pour **« réfléchir » avant de répondre** — ils génèrent une longue chaîne de pensée interne (cachée) avant la réponse finale. Résultats : **83 % à l'AIME 2024** (olympiades de maths US) contre 13 % pour GPT-4o ; 89e percentile sur Codeforces ; niveau doctorat sur GPQA. **C'est le pivot paradigmatique : on ne scale plus seulement le training, on scale l'inférence** — le « test-time compute ». Voir section 4 en détail.

**Décembre 2024 — DeepSeek-V3.** Le labo chinois publie un modèle MoE de **671 milliards de paramètres (37 milliards actifs par token)** entraîné pour **~5,6 millions de dollars** (chiffre du rapport technique, run final uniquement — le coût total incluant R&D est débattu, voir section 4). Architecture : MLA (Multi-head Latent Attention), DeepSeekMoE. Le rapport est d'une transparence rare pour un modèle de ce niveau.

### 1.12. 2025 : le choc DeepSeek-R1 et la course au raisonnement

**20 janvier 2025 — DeepSeek-R1.** Poids ouverts (licence MIT), papier détaillé, performances **au niveau d'o1** sur maths et code (AIME 2024 : 79,8 %, MATH-500 : 97,3 %), pour un coût d'entraînement du raisonnement estimé à **~294 000 $** (ordre de grandeur cité dans la littérature secondaire — « à vérifier » sur le rapport primaire). La variante **R1-Zero** montre l'émergence du raisonnement par **pur RL sans supervision humaine** — « aha moment », auto-vérification spontanée.

**Réaction des marchés : le 27 janvier 2025, NVIDIA perd ~593 milliards de dollars de capitalisation en une séance (-17 %)**, la plus grosse perte journalière de l'histoire boursière US à l'époque. La thèse « plus de compute = plus de valeur » est publiquement ébranlée : si la frontière s'atteint pour 100× moins cher, que vaut le hardware ?

**Réponse des labos US :** OpenAI sort **o3** (annoncé décembre 2024, déployé 2025 ; 96,7 % à l'AIME 2024 selon les benchmarks publiés), **o4-mini** (avril 2025) ; Anthropic généralise le raisonnement dans Claude ; Google pousse Gemini « thinking ». Le paradigme « raisonnement = RL + test-time compute » devient le standard 2025-2026.

**Agents et « unified intelligence ».** 2025 voit les assistants « deep research », les opérateurs web, et chez OpenAI la promesse d'une **intelligence unifiée** (GPT-5, août 2025) fusionnant les lignées GPT et o — lancement jugé décevant par la presse spécialisée (modèle « un peu meilleur » mais sans rupture, routeur thinking/non-thinking défaillant les premiers jours), illustration du débat sur « le mur » (section 4).

### 1.13. 2026 (janvier-septembre) : efficacité, prix cassés, maturité

D'après la littérature technique disponible au 27/09/2026 :

- **Distillation massive.** Les capacités de raisonnement des grands modèles sont **distillées** dans des modèles de 1,5B à 70B paramètres (R1-Distill-Qwen, etc.) : le raisonnement « o1-class » tourne sur du matériel accessible.
- **Guerre des prix API.** DeepSeek casse les prix de façon permanente (-75 % sur V4-Pro en mai 2026) ; l'inférence devient « too cheap to meter » sur les modèles courants. Le coût marginal du token s'effondre — mais le coût **total** explose avec les volumes (agents qui consomment des millions de tokens).
- **Le débat training vs inference.** Deloitte/Lenovo et plusieurs analystes décrivent le **« 80/20 reversal »** : ~80 % du compute dépensé en training en 2024, projection de ~80 % en **inférence** fin 2026. Le centre de gravité économique bascule du côté du serving.
- **Chips spécialisés.** Au-delà des GPU NVIDIA (Blackwell, puis Rubin), les TPU Google (Ironwood v7), les Trainium/Inferentia d'AWS et les projets de silicium custom se multiplient : l'inférence à grande échelle pousse vers du matériel **dédié**, moins généraliste que le GPU.

**Le fil rouge 2012→2026 :** chaque révolution a déplacé le goulot d'étranglement — des données (2012) au parallélisme (2017), à l'échelle (2020), à l'alignement (2022), à l'ouverture (2023), au raisonnement (2024-2025), puis à l'**efficacité économique de l'inférence** (2026). La section 4 examine si le premier moteur — scaler le training — est en train de caler.

---

## 2. Les grands noms : qui a fait quoi

> Règle de cette section : chaque fiche indique le fait vérifié et sa source publique. Quand un détail biographique n'a pas pu être vérifié, c'est signalé.

### 2.1. Les trois « parrains » du deep learning

**Geoffrey Hinton (né 1947, Royaume-Uni / Canada).**
- Popularise la **rétropropagation** pour les réseaux multicouches (Rumelhart, Hinton & Williams, *Nature*, 1986) — l'algorithme sans lequel rien de ce qui suit n'existe.
- **2006** : les *deep belief networks* relancent l'idée qu'on peut entraîner des réseaux profonds (fin symbolique du « deuxième hiver de l'IA » pour les réseaux de neurones).
- **2012** : AlexNet avec ses doctorants Krizhevsky et Sutskever (voir 1.2).
- **Prix Turing 2018** (avec Bengio et LeCun), puis **prix Nobel de physique 2024** (avec John Hopfield) « pour leurs découvertes fondamentales permettant l'apprentissage automatique avec des réseaux de neurones artificiels ».
- **2023** : quitte Google pour alerter publiquement sur les risques de l'IA (entretien NYT, mai 2023). Depuis, il estime publiquement qu'il y a une probabilité significative que l'IA dépasse l'humain dans les 5 à 20 ans (chiffre qu'il a lui-même fait varier selon les interviews — « à vérifier » au cas par cas).
- Position : le plus « alarmiste » des trois parrains sur le risque existentiel.

