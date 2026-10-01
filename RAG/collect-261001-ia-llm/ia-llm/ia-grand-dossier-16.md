---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-16
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Meta", "Nvidia", "OpenAI", "xAI"]
dates: []
keywords: ["agent", "agi", "chatgpt", "compute", "deepseek", "exploit", "gpu", "mai", "nvidia", "rlhf", "training", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1014, 1080]
sha256: 85c6b8864b766d3b27063a5b620bd00cb90735c9caa2dea82ddba84275efad35
---

# IA — Le grand dossier

**12.5. Rich Sutton — le père de la « Bitter Lesson ».**
Pionnier du **reinforcement learning** (avec Andrew Barto, livre de référence 1998/2018), professeur à Alberta, chercheur chez DeepMind puis Keen Technologies (avec John Carmack). Son essai **« The Bitter Lesson » (mars 2019)** — 800 mots — est devenu le texte sacré du camp « continuation » : les méthodes qui exploitent le calcul général battent toujours, à terme, les connaissances humaines codées en dur. C'est la justification philosophique du scaling.

**12.6. Gwern Branwen — l'essayiste du scaling.**
Auteur anonyme/pseudonyme, son essai **« The Scaling Hypothesis » (2020, mis à jour en continu sur gwern.net)** est la défense la plus argumentée de la thèse « scaler suffit » : courbes, réfutations des objections une par une, prédictions chiffrées. Très influent dans les labos (cité en interne chez OpenAI/Anthropic selon des témoignages publics). Position : les critiques du scaling « n'émettent pas de prédictions falsifiables » et ont toujours eu tort. Le représentant intellectuel du maximalisme, sans être salarié d'un labo.

**12.7. John Schulman — l'homme du PPO.**
Cofondateur d'OpenAI (2015, ex-UC Berkeley), inventeur du **PPO (2017)** — l'algorithme de RL derrière le RLHF, donc derrière ChatGPT. Architecte de l'alignement chez OpenAI, parti chez **Anthropic en août 2024** (sujet : alignement). Son parcours illustre la circulation des talents-clés entre les deux rivaux.

**12.8. Noam Brown — du poker au raisonnement.**
Chercheur CMU : **Libratus (2017)** bat les meilleurs joueurs de poker — jeu à information imparfaite, exploit jugé aussi dur que le go par d'autres voies. Rejoint OpenAI (via FAIR) et devient l'un des piliers du projet **o1/Strawberry** (raisonnement par RL). Trajectoire emblématique : **les techniques du jeu (search + RL) migrent vers le langage** — c'est exactement la convergence DeepMind × OpenAI décrite en 1.3 et 4.5.

**12.9. Jason Wei — le raisonnement par prompting.**
Chez Google Brain, co-auteur de **Chain-of-Thought (2022)** puis des travaux sur l'émergence ; passé chez OpenAI sur la lignée o. Ses threads et papiers ont popularisé l'idée que **le raisonnement est d'abord une propriété émergente du prompting**, avant de devenir une propriété entraînée (RL). Pour le praticien : ses techniques de prompting restent utiles même sans modèle « raisonnant ».

**12.10. Chris Olah — l'interprétabilité.**
OpenAI puis Anthropic (cofondateur). Ses travaux sur les **« circuits »** des réseaux (distill.pub, puis l'équipe d'interprétabilité d'Anthropic) cherchent à **comprendre ce que les neurones représentent vraiment** — la seule voie sérieuse vers des garanties de sûreté autres que les tests comportementaux. En 2024-2025, Anthropic publie des avancées (features « Golden Gate », dictionnaires sparse) : on commence à **lire dans les pensées** des modèles, partiellement. Enjeu sysadmin indirect : c'est de là que viendront les futurs outils d'audit des modèles déployés.

---

## 13. Les voix du débat : qui a dit quoi (citations et positions vérifiées)

> Opinions attribuées à leurs auteurs, avec la source publique. Ce ne sont pas des faits, ce sont des positions — à citer comme telles.

| Auteur | Position | Source publique |
|---|---|---|
| Sam Altman (OpenAI) | Le training de GPT-4 a coûté **« plus de 100 millions de dollars »** ; l'AGI est atteignable avec le paradigme actuel en continuant à scaler. | Déclarations publiques 2023 (coût) ; interviews récurrentes (AGI). |
| Elon Musk (xAI) | « On a épuisé à peu près **la somme cumulée des connaissances humaines** dans l'entraînement de l'IA. » | Livestream, début 2025 (rapporté par TechCrunch). |
| Dario Amodei (Anthropic) | Essai **« Machines of Loving Grace »** (oct. 2024) : une IA « digne d'un prix Nobel dans la plupart des disciplines » est envisageable d'ici 2026-2027 ; optimisme conditionné à la sûreté. | Publication Anthropic, octobre 2024. |
| Yann LeCun (Meta) | Les LLM sont une **impasse vers l'AGI** ; l'avenir = world models + JEPA ; le risque existentiel est « prématuré ». | Interviews et posts publics récurrents (2023-2026). |
| Geoffrey Hinton | Après son départ de Google (mai 2023), il alerte : probabilité significative que l'IA dépasse l'humain sous 5-20 ans (fourchette variable selon les interviews). | NYT mai 2023, puis interviews 2023-2025. |
| Rich Sutton | **« The Bitter Lesson »** (2019) : seules deux méthodes scalent indéfiniment avec le compute — **l'apprentissage et la recherche (search)**. | Essai, mars 2019. |
| Hyung Won Chung (OpenAI) | « **Compute is getting cheaper faster than we are becoming better researchers** » ; l'histoire des architectures = suppression des biais humains. | Stanford CS25, 2024. |
| Gwern Branwen | Les courbes de scaling **ne fléchissent pas** ; les critiques n'émettent pas de prédictions falsifiables et ont toujours eu tort. | « The Scaling Hypothesis », gwern.net, 2020→. |
| Pablo Villalobos / Epoch AI | Le stock de texte public de qualité (ordre ~300 000 Md tokens) sera **épuisé entre 2026 et 2032** au rythme actuel. | Epoch AI, 2022, mises à jour ensuite. |
| Ben Thompson (Stratechery) | Sur DeepSeek (janv. 2025) : le choc vient moins du coût que de la **fin des certitudes** (la Chine ne pouvait pas rivaliser, les coûts resteraient hauts) ; l'efficacité ne tue pas la demande. | « DeepSeek FAQ », Stratechery, janvier 2025. |
| Ilya Sutskever | Fonde **Safe Superintelligence Inc.** (juin 2024) : la superintelligence sûre exige un labo dédié, sans pression produit. Témoignage sous serment (oct. 2025) confirmant les discussions de fusion OpenAI-Anthropic en nov. 2023. | Annonces 2024 ; procès Musk vs Altman, 2025. |
| Jensen Huang (NVIDIA) | Le **« AI factory »** : les datacenters deviennent des usines à intelligence ; la demande de compute reste structurelle. | Keynotes GTC 2024-2026. |

**Comment utiliser ce tableau :** quand un article affirme « les experts disent que... », revenir ici et demander **quel** expert, **quand**, avec **quelles données**. C'est la méthode anti-bullshit de cette partie du dossier.

---

## 14. Cas pratiques chiffrés pour sysadmin (5 scénarios)

> Des scénarios réalistes avec des ordres de grandeur. Hypothèses explicites à chaque fois — à recalibrer avec vos tarifs.

### 14.1. Scénario 1 — RAG documentaire interne (le cas Zelef)

**Besoin :** 200 techniciens interrogent une base de ~50 000 pages (guides, manuels, procédures). 5 000 questions/jour, réponses de ~500 tokens, contexte RAG de ~4 000 tokens par requête.

**Calcul tokens/jour :** 5 000 × (4 000 in + 500 out) = 22,5M tokens/jour ≈ **675M tokens/mois**.

| Option | Coût mensuel estimé | Remarques |
|---|---|---|
| API modèle moyen (~2,5 $/1M in, ~10 $/1M out) | 600M×2,5 + 75M×10 ≈ **2 250 $** | Simple, mais données chez un tiers |
| API DeepSeek (~0,27 $/1M in, ~1,1 $/1M out) | ≈ **245 $** | 10× moins cher, mêmes réserves de confidentialité |
| Self-hosting 8B Q4 (1× RTX 4090, ~2 000 $ amortis 3 ans + élec) | ≈ **80-120 $/mois** | Données on-premise ; ~55 $/mois d'amortissement + ~30 $ d'élec (300 W × 24h × 30j × 0,15 €/kWh ≈ 32 €) |

**Verdict :** à ce volume, le self-hosting est **10-20× moins cher** que l'API généraliste et garde les données en interne. Le vrai coût n'est pas le GPU : c'est **la qualité du corpus** (le scraping/nettoyage que Zelef fait déjà) et l'**évaluation**.

### 14.2. Scénario 2 — Agent de supervision qui « réfléchit »

**Besoin :** un agent analyse chaque nuit 200 alarmes, avec raisonnement (10k tokens de thinking par alarme) sur un modèle raisonnant API à ~10 $/1M out.

**Calcul :** 200 × 10 000 = 2M tokens/nuit ≈ 60M/mois → **~600 $/mois** rien que pour le thinking.

