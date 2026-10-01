---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-5
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Cohere", "Google", "Meta", "Microsoft", "OpenAI", "Sakana", "United States", "xAI"]
dates: []
keywords: ["agents", "agi", "attention", "cohere", "embedding", "embeddings", "gpu", "grok", "llama", "mai", "research", "sakana"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [267, 325]
sha256: ee626fd5ae2b95cb9ea7f4703ae076d430cc5d2f33b596e5744437cec8647304
---

# IA — Le grand dossier

**Yann LeCun (né 1960, France).**
- **1989** : applique la rétropropagation aux réseaux convolutionnels (CNN) aux Bell Labs ; **1998** : LeNet-5 (avec Bottou, Bengio, Haffner) lit les chèques bancaires — première application industrielle majeure des CNN.
- **2013** : rejoint Facebook pour fonder et diriger **FAIR** (Facebook AI Research), devenu le labo ouvert le plus productif du monde (PyTorch, LLaMA plus tard).
- **Prix Turing 2018.** Professeur à NYU, Chief AI Scientist de Meta.
- Position publique : **le grand sceptique du scaling LLM**. Pour LeCun, les LLM sont une impasse (« les LLM ne mèneront pas à l'AGI »), l'avenir est aux **world models** et à l'apprentissage auto-supervisé type JEPA (Joint Embedding Predictive Architecture). Il qualifie les discours sur le risque existentiel de « prématurés » et s'oppose à la régulation restrictive de la recherche ouverte. C'est le « camp d'en face » de Hinton/Bengio sur la sûreté — tout en restant ami avec eux (dixit LeCun lui-même).

**Yoshua Bengio (né 1964, France / Canada).**
- Université de Montréal, directeur scientifique de **Mila** (Québec). Travaux fondateurs : **word embeddings** (modèle neuronal de langage, 2003 — l'ancêtre du pré-entraînement), mécanismes d'attention, théorie du deep learning.
- **Prix Turing 2018.**
- Position : aligné avec Hinton sur les risques (soutien à la régulation, signataire d'appels), mais reste un chercheur académique actif (GFlowNets, etc.). A dirigé le rapport scientifique international sur la sûreté de l'IA (AI Safety Report) commandé après le sommet de Bletchley Park (2023).

### 2.2. L'équipe « Attention Is All You Need » (Google, 2017)

Les huit auteurs, ordre du papier : **Ashish Vaswani, Noam Shazeer, Niki Parmar, Jakob Uszkoreit, Llion Jones, Aidan N. Gomez, Łukasz Kaiser, Illia Polosukhin** (arXiv:1706.03762, NeurIPS 2017).

- **Ashish Vaswani** : premier auteur, à l'époque chez Google Brain. Quitte Google en 2021, cofonde **Adept AI Labs** (agents IA, levée de ~350 M$ en 2023 — « à vérifier » pour le montant exact).
- **Noam Shazeer** : vétéran de Google (depuis 2000), inventeur du Transformer avec Vaswani ; cofonde **Character.AI** en 2021 avec Daniel De Freitas (chatbots de personnages, succès grand public 2023), puis **retourne chez Google en 2024** dans le cadre d'un accord rapporté par la presse à ~2,7 milliards $ (montant « à vérifier ») — cas d'école de l'« acqui-hire » des hyperscalers.
- **Aidan N. Gomez** : cofonde **Cohere** (2020, Toronto) — labo « entreprise-first », modèles Command, RAG et embeddings.
- **Llion Jones** : cofonde **Sakana AI** (Tokyo, 2023) avec David Ha — labo qui explore l'IA « inspirée de la nature » et l'automatisation de la recherche.
- **Illia Polosukhin** : cofonde **NEAR Protocol** (blockchain, 2017) — trajectoire crypto/IA décentralisée.
- **Niki Parmar, Jakob Uszkoreit, Łukasz Kaiser** : carrières entre Google, startups et recherche (Kaiser a notamment travaillé sur le raisonnement et les architectures efficaces ; Uszkoreit a cofondé **Inceptive** — biotech IA — « à vérifier »).

**Leçon :** le papier le plus cité de l'histoire de l'IA moderne a été écrit par une équipe dont **plus aucun membre n'est resté chez Google** — la valeur a fui vers l'écosystème.

### 2.3. Les fondateurs et dirigeants de labos

**Sam Altman (né 1985).**
- Ex-président de Y Combinator (2014-2019). Cofondateur d'**OpenAI** (décembre 2015, avec Elon Musk, Ilya Sutskever, Greg Brockman et d'autres — voir 2.4).
- Pilote le virage **capped-profit de 2019** et le partenariat Microsoft (1 Md$ en 2019, puis ~10 Md$ en janvier 2023).
- **17 novembre 2023** : limogé par le board d'OpenAI (« il n'a pas été constamment franc dans ses communications »), réintégré le **22 novembre 2023** après la révolte des ~770 employés (plus de 700 signent une lettre de démission). Le board avait même exploré une **fusion avec Anthropic** (confirmé sous serment par Ilya Sutskever en octobre 2025, dans le cadre du procès Musk vs Altman).
- Stratège du scaling maximaliste : « l'AGI est atteignable avec le paradigme actuel » (position publique récurrente). A confirmé publiquement que GPT-4 avait coûté **plus de 100 millions de dollars** à entraîner.

**Ilya Sutskever (né 1986).**
- Co-auteur d'AlexNet (2012), co-inventeur des **seq2seq** (2014, avec Vinyals et Le). Chief Scientist d'OpenAI de la fondation à **mai 2024**.
- Membre du board lors du limogeage d'Altman (nov. 2023), puis regret public (« je regrette ma participation »).
- **Juin 2024** : fonde **Safe Superintelligence Inc. (SSI)** — superintelligence sûre, sans produit commercial à court terme. Levée d'environ 1 Md$ en septembre 2024 (montant rapporté — « à vérifier »).
- Figure du « camp prudence » interne à OpenAI ; son départ marque la victoire du camp accélérationniste.

**Dario Amodei (né 1983) & Daniela Amodei.**
- Dario : ex-VP Recherche d'OpenAI (2016-2020), spécialiste du scaling et de l'interprétabilité. Daniela : ex-VP Safety & Policy.
- **2021** : fondent **Anthropic** avec une dizaine d'ex-OpenAI (Jack Clark, Chris Olah, Sam McCandlish, Tom Henighan, Jared Kaplan entre autres — Kaplan est le premier auteur des lois de scaling 2020 !).
- Motif du départ : désaccord sur la **direction sûreté vs vitesse** (détails : le « plan Brockman » de levée incluant des États rivaux, promesses de gouvernance non tenues — relaté par la presse US, « à vérifier » dans le détail).
- Octobre **2024** : Dario publie l'essai **« Machines of Loving Grace »** — vision optimiste et datée (il y envisage une « IA digne d'un prix Nobel dans la plupart des disciplines » d'ici 2026-2027 — prédiction, pas fait).
- Positionnement : « scale fast AND safe » — accélération avec garde-fous, IA constitutionnelle, publication du **Responsible Scaling Policy**.

**Demis Hassabis (né 1976).**
- Prodige des échecs, co-designer de *Theme Park* à 17 ans, PhD neurosciences à UCL.
- **2010** : cofonde **DeepMind** (avec Shane Legg et Mustafa Suleyman) avec la mission « solve intelligence ».
- **2014** : rachat par Google (~400-650 M£/$ selon les sources — fourchette « à vérifier » au dollar près).
- AlphaGo (2016), AlphaZero (2017), AlphaFold 2 (2020), **Nobel de chimie 2024**.
- Depuis avril **2023** : CEO de **Google DeepMind** (fusion DeepMind + Google Brain). Fonde aussi **Isomorphic Labs** (2021, drug discovery).
- Position : le « troisième pôle » — ni maximaliste pur (Altman) ni alarmiste (Hinton) : AGI via percées scientifiques + planification, pas seulement du scaling de LLM.

**Elon Musk.**
- Cofondateur d'OpenAI (2015), départ du board en **2018** (conflit d'intérêts Tesla + désaccord stratégique).
- **Juillet 2023** : fonde **xAI** (« comprendre la vraie nature de l'univers »), lance **Grok** (intégré à X). Construit le cluster **Colossus** (Memphis, 100k+ GPU H100 annoncés en 2024 — « à vérifier » le chiffre exact opérationnel).
- **2024-2025** : procès contre OpenAI/Altman (abandon de la mission non-profit allégué).
- Citation vérifiée (livestream, début 2025, rapportée par TechCrunch) : « on a maintenant épuisé à peu près la somme cumulée des connaissances humaines dans l'entraînement de l'IA » — argument du **mur des données** repris par le camp du ralentissement.

