---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-9
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Huawei", "Microsoft", "Nvidia", "OpenAI", "Oracle", "Samsung", "TSMC", "United States", "vLLM"]
dates: []
keywords: ["agents", "ascend", "awq", "capex", "claude", "compute", "datacenter", "deepseek", "distillation", "gemini", "gguf", "gptq"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [504, 569]
sha256: 7898806c2a24ad516e1ec0431975aee440cd162a3f3201983ae8d2e67804b82d
---

# IA — Le grand dossier

**Argument 3 : les données ne sont pas finies si on les fabrique.**
- **Données synthétiques** : les modèles génèrent leurs propres données d'entraînement (distillation, R1 qui s'auto-améliore). Risque de « model collapse » (Shumailov et al., 2023 — « à vérifier » la portée exacte : l'effondrement est démontré en l'absence de données fraîches, pas une fatalité).
- **RL + environnements** : comme AlphaGo Zero (2017), les modèles de raisonnement apprennent par **auto-jeu et vérificateurs** (maths formelles, code compilé) — **sans limite de données humaines**. C'est la thèse du « Bitter Lesson » de **Rich Sutton (2019)** : seules deux méthodes scalent indéfiniment avec le compute — **l'apprentissage et la recherche (search)**.
- **Multimodalité** : la vidéo (YouTube), les capteurs, la robotique ouvrent des stocks de données **ordres de grandeur supérieurs** au texte.

**Argument 4 : l'efficacité compense le coût.**
DeepSeek-V3/R1 prouve qu'on peut atteindre la frontière pour **10 à 100× moins cher** qu'anticipé. Si le coût par unité d'intelligence chute de 90 %, le scaling économique continue même si le scaling brut ralentit. C'est l'argument **Jevons** : moins cher = plus d'usages = plus de demande totale de compute.

**Porte-voix du camp continuation :** Sam Altman, Dario Amodei (« Machines of Loving Grace », oct. 2024), Gwern, Rich Sutton (Bitter Lesson), les équipes o1/o3 d'OpenAI, Ben Thompson (Stratechery, analyse DeepSeek : l'efficacité ne tue pas la demande).

### 4.4. Efficacité vs taille : les quatre leviers qui changent la donne

**Levier 1 : Mixture of Experts (MoE).**
Au lieu d'activer tout le modèle pour chaque token, on route chaque token vers **quelques « experts »** (sous-réseaux) parmi des dizaines/centaines. DeepSeek-V3 : **671B paramètres totaux, 37B actifs par token** (~5,5 %). Mixtral 8x7B : 47B totaux, ~13B actifs. **Effet : capacité d'un très grand modèle, coût d'inférence d'un modèle moyen.** Le MoE est devenu l'architecture standard des modèles frontières ouverts 2024-2026.

**Levier 2 : la distillation.**
On entraîne un **petit modèle (élève)** à imiter un **grand modèle (professeur)** — y compris ses traces de raisonnement (R1-Distill). Résultat : des modèles de **1,5B à 70B** paramètres atteignent des performances « o1-class » sur maths/code. **Pour le RAG : c'est LA technique qui rend le raisonnement déployable en entreprise** (voir section 6).

**Levier 3 : le sur-entraînement (over-training).**
Chinchilla donne l'optimum **training-compute** (20 tokens/paramètre). Mais le training est un coût **unique**, l'inférence un coût **récurrent**. L'optimum **économique total** pousse à entraîner des **petits modèles sur énormément de données** : Llama 3 8B sur 15 000B tokens (~1 875 tok/param, ~100× Chinchilla). En 2025, la moyenne des modèles ouverts est à ~300 tokens/paramètre (~15× Chinchilla). **Traduction sysadmin : les modèles que vous servez sont de plus en plus petits et de plus en plus « denses en connaissance ».**

**Levier 4 : la quantification.**
Passer les poids de FP16/BF16 à **INT8, INT4, voire INT2** : divise par 2 à 8 la VRAM nécessaire, avec une perte de qualité faible à modérée (selon la méthode : GPTQ, AWQ, GGUF). **C'est ce qui fait tourner un 70B sur 2 GPU prosumer ou un 8B sur un laptop.** Écosystème : llama.cpp, Ollama, vLLM.

### 4.5. Le nouveau paradigme : inference scaling / test-time compute

**Le principe.** Au lieu d'investir tout le compute **avant** (training), on en dépense **pendant** la réponse : le modèle génère des milliers de tokens de « réflexion » interne (chaîne de pensée cachée), se vérifie, explore plusieurs pistes, puis répond.

**La lignée :**
- **2022** : Chain-of-Thought prompting (Wei et al.) — le prompting qui révèle le potentiel.
- **Sept. 2024** : **o1** (OpenAI, ex-« Strawberry ») — le RL appliqué au raisonnement, 83 % AIME 2024.
- **Janv. 2025** : **DeepSeek-R1** — la réplication ouverte, R1-Zero (pur RL, émergence spontanée du raisonnement).
- **2025-2026** : généralisation (o3, Claude raisonnant, Gemini thinking), **RLVR** (récompenses vérifiables), distillation du raisonnement.

**Pourquoi c'est un changement de paradigme économique :**
- Le training est un coût **fixe et spéculatif** (on paie avant de savoir si ça marche).
- L'inférence est un coût **variable et pilotable** : paramètre `reasoning_effort` (low/medium/high) — on paie la réflexion **seulement quand la question le mérite**.
- **Conséquence : le budget compute se déplace du CAPEX (clusters d'entraînement) vers l'OPEX (flotte d'inférence).** C'est exactement le « 80/20 reversal ».

**Limites connues (2025-2026) :** rendements décroissants du thinking pur (« overthinking »), latence (secondes à minutes), coût par requête complexe. D'où les garde-fous : **plafonds dynamiques de tokens**, routage thinking/non-thinking (le routeur de GPT-5, défaillant au lancement — symptôme que le routage est un problème dur).

### 4.6. Les murs physiques : énergie, chips, géopolitique

**Énergie.** Un datacenter IA frontière consomme **100 MW à 1 GW** (ordres de grandeur publics 2024-2026). La contrainte n'est plus le silicium seul : c'est le **raccordement électrique** (délais de 2-4 ans aux US — « à vérifier » selon les régions) et le **refroidissement**. D'où le retour du nucléaire dans les annonces (SMR, redémarrage de centrales — annonces 2024-2025, « à vérifier » l'état d'avancement projet par projet).

**Chips.** Dépendance quasi-totale à **TSMC** (gravure) et **SK Hynix/Micron/Samsung** (HBM — la mémoire à large bande, goulot 2024-2026). Les contrôles d'export US vers la Chine (2022-2023+) ont créé **deux écosystèmes** : NVIDIA bridé (H800) vs Huawei Ascend et SMIC côté chinois (« à vérifier » les performances relatives, sujet mouvant).

**Géopolitique du compute.** « Stargate » (OpenAI/Microsoft/Oracle, ~500 Md$ annoncés janvier 2025), plans souverains (UE, Émirats, Arabie saoudite). Le compute est devenu une **ressource stratégique** au même titre que le pétrole — avec les mêmes logiques d'alliance et de blocus.

### 4.7. Synthèse : que conclure en septembre 2026 ?

| Question | Réponse la plus honnête (état de la recherche) |
|---|---|
| Le scaling du **pré-entraînement** continue-t-il comme avant ? | **Non.** Rendements décroissants mesurables + mur des données + coûts : le scaling naïf « 10× le compute tous les 2 ans » ne produit plus les sauts d'avant. |
| L'IA a-t-elle atteint un plafond ? | **Non.** Le progrès s'est **déplacé** : test-time compute, RL, efficacité (MoE, distillation, quantification), multimodalité, agents. Les sauts 2024-2025 (o1, R1) le prouvent. |
| Les deux camps ont-ils partiellement raison ? | **Oui.** Le camp « mur » a raison sur le **training** ; le camp « continuation » a raison sur le **système global** (entraînement + inférence + algo). |
| Que doit retenir un décideur/sysadmin ? | **Le centre de gravité passe du training à l'inférence.** On n'achète plus (seulement) des clusters d'entraînement ; on dimensionne des **flottes de serving**, on optimise le **coût par token**, on déploie des **petits modèles distillés**. Voir section 6. |

**Formule mémotechnique :** *« La fin du scaling du training n'est pas la fin du progrès — c'est la fin du gaspillage. »* (formulation de l'auteur du dossier, pas une citation).

---

## 5. Timeline détaillée 2012 → 2026

> Convention : **P** = papier/recherche, **M** = modèle/produit, **E** = événement industriel/géopolitique, **I** = infrastructure.

