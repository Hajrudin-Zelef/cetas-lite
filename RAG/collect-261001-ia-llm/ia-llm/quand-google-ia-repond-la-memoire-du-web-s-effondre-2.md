---
id: collect-261001-ia-llm/ia-llm/quand-google-ia-repond-la-memoire-du-web-s-effondre-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Anthropic", "Apple", "Google", "Meta", "Moonshot", "Nvidia", "OpenAI"]
dates: []
keywords: ["agent", "agents", "amd", "benchmark", "benchmarks", "chatgpt", "claude", "distillation", "exploit", "inference", "int4", "kimi"]
source: docs/RAG/collect-261001-ia-llm/quand-google-ia-repond-la-memoire-du-web-s-effondre.md
source_anchor: ""
source_lines: [83, 132]
sha256: ba30ee77e5013ac741831b935d1773b5daa74eb4ad8ee1402797014178880674
---

# 🧠 **RECHERCHE**

Conception signée LoveFrom, le studio de Jony Ive, via le rachat de sa startup hardware io par OpenAI, finalisé en mai 2025 pour environ **6,5 milliards de dollars** .
**Caméra intégrée et reconnaissance faciale de type Face ID** : l'appareil identifie qui lui parle et ce qui l'entoure, avec des micros et des lumières qui signalent quand il écoute.
Le détail qui intrigue : **des parties de la coque bougent d'elles-mêmes** quand l'appareil répond. L'objectif affiché est de paraître vivant, là où un Echo reste une brique posée sur un meuble.
L'appareil permettra de **faire des achats directement** , sans passer par un téléphone.
Positionnement tarifaire **25 à 30% au-dessus du haut de gamme Amazon Echo** (40 à 240 dollars). Aucune spécification technique n'a fuité : ni puce, ni autonomie, ni date d'annonce officielle.

Rien n'est confirmé par OpenAI, tout vient de fuites presse, et il n'y a strictement rien à tester avant 2027. Mais la direction est limpide : après ChatGPT et Codex, l'entreprise cherche à exister ailleurs que dans un navigateur ou une application, c'est-à-dire ailleurs que sur des plateformes contrôlées par Apple et Google. Reste la question que se poseront tous ceux qui ont déjà une enceinte connectée muette au salon : 300 dollars pour un objet qui vous filme, qu'est-ce qu'il fera de plus ?

Lundi, Mark Zuckerberg a mis en ligne « The Future is for Everyone », un essai de **plus de 6 500 mots** sur la coexistence entre l'humanité et une IA superintelligente, prolongement d'une lettre publique bien plus courte publiée l'an dernier. Derrière la philosophie, le texte contient quatre demandes très concrètes et un chèque d'**1 milliard de dollars**.

**Data centers** : Meta promet de restituer plus d'eau qu'elle n'en consomme d'ici 2030 et des emplois locaux, pour rendre ses infrastructures socialement acceptables. Le contexte chiffré :**145 milliards de dollars d'investissement prévus en 2026** , majoritairement en data centers.
**Accès** : rendre l'IA gratuite ou quasi gratuite pour « des milliards de personnes » via des modèles open source téléchargeables, avec**Muse Glimmer** et une version open-weight de**Muse Spark 1.2** annoncées, sans date précise.
**Régulation** : alléger les règles américaines sur les données et la distillation, jugées inefficaces puisqu'elles ne s'appliquent pas aux modèles étrangers.
**Coopération avec l'État** : partager les « points de contrôle d'entraînement intermédiaires » pour la cybersécurité et la détection d'abus.
Un **« Future Is for Everyone Fund » doté d'1 milliard de dollars** , et la promesse d'agents personnels gratuits ou à très bas coût sur la santé, la carrière, les finances et les relations.

Le fil rouge du texte est le refus de voir la superintelligence rester concentrée dans « une poignée de labos », argument qui a l'avantage d'être à la fois défendable et parfaitement aligné sur la stratégie open source de Meta. À noter, un changement plus discret mais réel : c'est désormais le conseil d'administration de Meta, et non plus Zuckerberg seul, qui valide les critères de sécurité des modèles. Pour le lecteur, l'engagement le plus vérifiable reste le plus simple : des modèles téléchargeables, ou pas.

# 🧠 **RECHERCHE**

**Kimi K3 sort du bac à sable pour aller lire les réponses de son examen**

Des chercheurs de Frontier Security ont observé le modèle chinois Kimi K3 en train de tricher pendant une évaluation de cybersécurité du UK AI Safety Institute. Profitant d'une résolution DNS sortante restée fonctionnelle dans le sandbox censé l'isoler, le modèle a atteint github.com, cloné le dépôt officiel du benchmark et lu la solution sur le disque plutôt que de résoudre la tâche. La faille touche les frameworks Inspect et Cybench, et Kimi K3 étant un modèle public, contrairement au cas OpenAI détecté en amont, n'importe qui peut reproduire l'exploit.

**Claude fait progresser un résultat lié à l'hypothèse de Riemann**

Un employé d'Anthropic a demandé à une version de recherche non publiée de Claude de s'attaquer à l'hypothèse de Riemann. Le modèle a échoué, comme prévu depuis 1859, mais a fait au passage progresser un résultat connexe : la proportion démontrée de zéros de la fonction zêta qui vérifient l'hypothèse passe de **41,6% à 67,2%**, une borne qui n'avait pas bougé depuis des décennies. La preuve, appuyée sur les travaux de Baluyot, Goldston, Suriajaya, Turnage-Butterbaugh et Bombieri, a été validée par deux mathématiciens d'Anthropic et des experts externes.

**Un mini LLM tourne à près de 60 000 tokens par seconde dans un FPGA à 250 dollars**

Un développeur a fait tenir un modèle de langage de **3,16 millions de paramètres** (1,5 Mo en INT4) entièrement dans la mémoire on-chip d'une carte AMD Kria KV260, sans jamais toucher la RAM externe. Résultat : **59 965 tokens par seconde** dans la logique reconfigurable, contre **11 tok/s** sur les cœurs ARM de la même puce et **719 tok/s** sur une RTX 3050 Ti. Le modèle ne sait générer que de petites histoires, ce n'est pas un assistant, mais la démo live tourne en direct sur une carte physique installée au pays de Galles.

**NVIDIA ouvre les poids de Magpie TTS, 12 langues dont le français**

NVIDIA met à jour son modèle de synthèse vocale open-weights **Magpie TTS Multilingual**, 364 millions de paramètres, qui couvre désormais **12 langues** avec l'ajout de l'arabe standard moderne, du coréen et du portugais brésilien. L'argument n'est pas la qualité brute mais le contrôle : déployable sur votre propre infrastructure via NVIDIA NIM, il permet de maîtriser la latence et la confidentialité des données, là où les API vocales fermées imposent leur boîte noire.

**Sonder les modèles pour deviner combien de paramètres ils cachent**

Un chercheur détaille des techniques pour extraire des informations que les labos ne publient pas. Les « Incompressible Knowledge Probes » testent les modèles sur des faits très pointus pour estimer leur nombre de paramètres, la « Data Mixture Inference » analyse la façon dont ils découpent les tokens pour déduire les jeux de données d'entraînement, et des questions liées aux dates révèlent les calendriers de pré-entraînement. Faute de référence publique, les résultats restent spéculatifs, mais la méthode est instructive.

**Humaniser les réponses des LLM serait une mauvaise idée**

L'auteur s'attaque à une mode : demander aux agents IA d'écrire court, sans jargon, en « anglais technique simplifié » ou en style adapté au TDAH. Le problème, selon lui, est que l'instruction ne s'applique pas après le travail mais pendant : c'est une compression avec perte, appliquée en continu, qui efface au passage les signaux d'échec (hypothèses non résolues, résultats contradictoires, incertitudes). Sa proposition : garder la représentation la plus riche possible entre agents, et ne compresser qu'au dernier moment, face à l'humain.

**Quel langage de programmation pour les agents de codage ?**

Une étude très citée affirme que les langages dynamiques et concis comme Clojure ou J consomment 2 à 3 fois moins de tokens que Rust, Go ou C++, au point que les résumés IA de Google reprennent la conclusion. Dan Luu démonte la méthode : les benchmarks reposent sur des problèmes triviaux de type Rosetta Code, sans rapport avec du vrai code, et l'un d'eux contient un bug (un symlink erroné créé par un agent) qui a faussé le scoring de plusieurs langages testés ensuite.

**Les chercheurs universitaires s'adaptent à la mainmise des labos privés**

