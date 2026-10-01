---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-8
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Google", "Lambda", "Meta", "OpenAI", "xAI"]
dates: []
keywords: ["benchmarks", "compute", "deepseek", "gemini", "llama", "moe", "training"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [446, 503]
sha256: a6822ad80afe7848688f8b3ec12b244fb6dcf4af74f2b1dac5df60a511f2c82b
---

# IA — Le grand dossier

```
L(N) ∝ N^(-0,076)    L(D) ∝ D^(-0,095)    L(C) ∝ C^(-0,050)
```

Conclusion opérationnelle de Kaplan : **à budget compute fixé, mieux vaut un très gros modèle peu entraîné qu'un petit modèle très entraîné**. D'où GPT-3 : 175B paramètres pour ~300B tokens (1,7 token/paramètre).

**Hoffmann et al. (DeepMind, mars 2022)** — « Training Compute-Optimal Large Language Models » (arXiv:2203.15556), le papier **Chinchilla**. En corrigeant un biais méthodologique de Kaplan (les petits modèles testés n'étaient pas entraînés à convergence), l'équipe montre que **N et D doivent scaler à parts égales** :

```
N_opt ∝ C^0,5    et    D_opt ∝ C^0,5    →    D ≈ 20 × N  (règle des 20 tokens/paramètre)
```

Démonstration : **Chinchilla (70B paramètres, 1 400B tokens)** bat **Gopher (280B paramètres, 300B tokens)** sur presque tous les benchmarks, à compute égal. GPT-3 était donc **sous-entraîné d'un facteur ~10**.

**Ce que les deux papiers ont en commun :** dans les deux cas, **plus de compute = meilleure performance**, de façon lisse et prévisible. C'est cette promesse — le « scaling naïf » — qui a justifié des centaines de milliards d'investissement. Le débat de 2024-2026 porte sur **ce qui se passe quand on continue** : la courbe continue-t-elle, s'aplatit-elle, ou bute-t-elle sur des murs ?

### 4.2. Le camp « le scaling du training ralentit » — arguments et auteurs

**Argument 1 : les rendements décroissants sont mesurables.**
Sur les benchmarks standard (MMLU, GSM8K, HumanEval), chaque doublement de compute apporte **moins** que le précédent. L'exemple emblématique : **GPT-4.5 « Orion » (mars 2025)** — supposé être GPT-5 vu le saut d'échelle, il n'apporte pas de gain « surprenant » pour un coût très élevé, et est déprécié quelques mois plus tard. **GPT-5 (août 2025)** est « un peu meilleur » mais sans avance dominante, avec un routeur défaillant au lancement. La règle implicite « +1 au numéro = rupture » est brisée. (Analyse : presse spécialisée 2025 ; faits de lancement vérifiés, l'interprétation « mur » est celle des commentateurs.)

**Argument 2 : le mur des données.**
- **Epoch AI (Pablo Villalobos et al., 2022, mis à jour ensuite)** : le stock de **texte public de qualité** filtrable est estimé à ~300 000 milliards de tokens (ordre de grandeur) ; au rythme de consommation des labos frontières, il serait **épuisé entre 2026 et 2032**. Les labos frontières subissent déjà des contraintes sur leurs budgets de tokens uniques.
- **Elon Musk (livestream, début 2025)** : « on a épuisé à peu près la somme cumulée des connaissances humaines dans l'entraînement ».
- **Data Provenance Initiative (MIT, 2024)** : sur trois grands datasets (C4, RefinedWeb, Dolma), **5 % de toutes les données et 25 % des sources de haute qualité ont été restreintes en un an** (robots.txt, CGU, paywalls : NYT, Reddit, Stack Overflow...).
- Correctifs théoriques : **Muennighoff et al. (2023)** montrent que la valeur des tokens **répétés** décroît exponentiellement (demi-vie apprenable) ; **Lovelace et al. (2026)** ajoutent une pénalité de surparamétrisation : sous contrainte de données, mieux vaut **plus d'époques** que plus de paramètres. Le sur-entraînement volontaire (Llama 3 8B : 15 000B tokens pour 8B paramètres, soit ~1 875 tokens/paramètre, ~100× Chinchilla) est la réponse industrielle à ce mur.

**Argument 3 : le coût du compute devient prohibitif.**
Ordres de grandeur vérifiés ou consensuels :

| Modèle | Année | Compute (FLOP) | Coût estimé du run final | Source |
|---|---|---|---|---|
| Transformer originel | 2017 | ~10^19 | ~930 $ | Littérature (ordre de grandeur) |
| GPT-3 (175B) | 2020 | 3,14×10^23 | ~4,6 M$ | Lambda Labs (estimation) |
| GPT-4 | 2023 | « à vérifier » (non publié) | **> 100 M$** | Sam Altman, déclaration publique |
| DeepSeek-V3 (671B MoE) | 2024 | ~2,8×10^24 (« à vérifier ») | **5,576 M$** (run final) | Rapport technique DeepSeek |
| GPT-5 (génération) | 2025 | « à vérifier » | ~500 M$ (analystes) | Estimations d'analystes — « à vérifier » |

Le coût **total de développement** est typiquement **2 à 3× le run final** (expériences, échecs, itérations). Structure des coûts (Epoch AI, sur GPT-3/OPT/GPT-4/Gemini Ultra) : **hardware 47-65 %, personnel 29-49 %, électricité 2-6 %**. Oui : l'électricité ne représente que quelques pourcents — le silicium et les chercheurs coûtent bien plus cher que les électrons.

**Argument 4 : le « overthinking » et les limites du test-time scaling.**
Même le nouveau paradigme (raisonnement à l'inférence, voir 4.5) montre des **rendements décroissants** : au-delà d'un budget de « thinking tokens », les modèles entrent dans des boucles de vérification dégénératives (« overthinking ») — études arXiv 2025-2026, débattues sur les forums ML. D'où le pivot vers le **RLVR** (Reinforcement Learning with Verifiable Rewards) : récompenses vérifiables (tests unitaires, démonstrateurs formels) plutôt que raisonnement libre.

**Argument 5 : les signaux économiques.**
Deloitte/Lenovo et analystes (2026) : le **« 80/20 reversal »** — ~80 % du compute en training en 2024, projection de ~80 % en **inférence** fin 2026. Les clusters généralistes d'entraînement risquent le **surdimensionnement** (« stranded capacity ») faute de raccordement électrique. Goldman Sachs / Bank of America (2026) pointent le décalage entre ~400 Md$ d'infra et ~100 Mds$ de revenus logiciels incrémentaux (« à vérifier » les chiffres exacts par banque, l'ordre de grandeur du débat est public).

**Porte-voix du camp ralentissement :** Epoch AI (Villalobos), Elon Musk, une partie de la presse spécialisée (ex. : « The Wall », cronica-ia, mars 2025), les analystes financiers 2026, et — avec des nuances — **Yann LeCun** (pour qui le paradigme lui-même est à bout de souffle, pas seulement son scaling).

### 4.3. Le camp « le scaling continue » — arguments et auteurs

**Argument 1 : « on nous annonce la fin du scaling depuis 2015, on a toujours eu tort ».**
C'est la position de **Gwern Branwen** (essayiste, « The Scaling Hypothesis », 2020, mis à jour) : les courbes de scaling **ne montrent aucun signe de fléchissement** quand on les trace correctement ; les critiques ont prédit la fin à chaque génération et ont été réfutées. Le tableau des critiques récurrentes (« c'est de la force brute », « rendements décroissants », « trop cher ») avec leurs réfutations est devenu un classique du genre.

**Argument 2 : le scaling a changé d'axe, pas de nature.**
- **Test-time compute (Snell et al., 2024**, « Scaling LLM Test-Time Compute Optimally ») : la performance suit des **lois d'échelle à l'inférence** — plus le modèle « réfléchit » (tokens de raisonnement), meilleur il est, selon une loi de puissance. OpenAI o1/o3 et DeepSeek-R1 en sont la démonstration produit.
- **Hyung Won Chung (OpenAI, Stanford CS25, 2024)** : l'histoire des architectures Transformer est une **suppression progressive des biais inductifs humains** au profit de méthodes générales qui exploitent l'échelle (« Compute is getting cheaper faster than we are becoming better researchers »).
- Les tenants de ce camp notent que **les modèles 2024-2025 (o1, o3, R1) ont produit des sauts discontinus** (AIME 13 % → 83 % → 96,7 %) sans rupture d'échelle de training — preuve que le paradigme n'est pas épuisé, il **mute**.

