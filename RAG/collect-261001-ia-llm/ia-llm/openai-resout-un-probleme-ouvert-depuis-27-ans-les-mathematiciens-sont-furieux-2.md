---
id: collect-261001-ia-llm/ia-llm/openai-resout-un-probleme-ouvert-depuis-27-ans-les-mathematiciens-sont-furieux-2
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "Meta", "Moonshot", "OpenAI", "Z.ai"]
dates: []
keywords: ["agentic", "agents", "astra", "attention", "benchmark", "chatgpt", "claude", "deepseek", "distillation", "embeddings", "fable 5", "glm"]
source: docs/RAG/collect-261001-ia-llm/openai-resout-un-probleme-ouvert-depuis-27-ans-les-mathematiciens-sont-furieux.md
source_anchor: ""
source_lines: [91, 150]
sha256: e7a5c0ba3610585d94858a934a50621586ab1a1132840f0516bedd2ad467eb5f
---

# 🧠 **RECHERCHE**

Anthropic a publié le 10 septembre son rapport « Detecting and Countering Misuse of AI », qui documente **cinq cas** où des comptes Claude ont servi à des travaux pouvant faciliter le développement d'armes biologiques. Le plus marquant : un utilisateur demandait à Claude de rédiger une **demande de subvention** pour de la recherche de gain de fonction sur le **virus chikungunya**, afin de le muter chez l'animal vivant pour le rendre plus infectieux et plus capable d'échapper au système immunitaire. Le virus n'a aucun traitement autorisé et provoque des symptômes invalidants pendant des semaines.

**Les points essentiels :**

- La détection vient du **classificateur de sécurité biologique** interne d'Anthropic, pas d'un signalement extérieur. Le projet était rattaché à un **institut de recherche militaire**, ce qui a pesé dans la décision

- Deuxième cas : un chercheur hors des États-Unis a échangé **des milliers de messages sur plusieurs semaines** autour de la grippe aviaire hautement pathogène, sur des approches génétiques liées à l'adaptation aux mammifères et à la transmission par voie aérienne

- Autre dossier : la constitution d'un **atlas de peptides toxiques de venin** avec un pipeline génératif optimisant les caractéristiques des toxines

- **Tous les comptes ont été bannis**, mais ni les personnes ni leurs laboratoires ne sont nommés, pour ne pas les exposer. Les renseignements ont été partagés avec les autorités et d'autres entreprises d'IA

- Nuance officielle : Anthropic précise que ce sont des **scientifiques en activité** et n'affirme pas qu'ils avaient une intention malveillante

« Vous ne voyez pas quelqu'un dire, façon comic book, "je veux construire une arme biologique pour tuer tout le monde" », résume Jacob Klein, responsable du renseignement sur les menaces chez Anthropic : « c'est une situation incroyablement nuancée. » Voilà le vrai problème : une demande de subvention pour un vaccin et une demande pour une arme se ressemblent, et plus les modèles deviennent bons en assistance à la recherche, plus le tri devient arbitraire. Anthropic dit avoir systématiquement tranché dans le sens de la prudence, donc au risque de bloquer des travaux légitimes.

# 🧠 **RECHERCHE**

### **Colibrì fait tourner un modèle de 744 milliards de paramètres sur 25 Go de RAM**

Un runtime open source en C pur, sans aucune dépendance, qui exécute GLM-5.2 sur du matériel grand public. L'astuce : seuls ~40 milliards de paramètres sont activés par token, donc Colibrì garde en RAM les tenseurs denses quantifiés en int4 (attention et embeddings, environ 9,9 Go) et laisse les ~370 Go d'experts routés sur un SSD NVMe, chargés à la demande à chaque étape de décodage. C'est lent, mais ça tourne chez vous.

**DeepSeek V4.1-Flash divise par quatre la mémoire des agents IA**

Modèle multimodal de 552 milliards de paramètres dont 16 milliards seulement s'activent par token, avec un cache KV réduit au quart de celui de la version précédente. Sur le benchmark de code DeepSWE, il devance de peu Opus 5 et GPT-5.6 Sol. Le tout sous licence MIT, ce qui rend les agents nettement moins coûteux à faire tourner en continu.

**MultiMatte détoure une image à partir d'une phrase**

La startup Feyn publie un modèle de suppression d'arrière-plan qu'on vise avec des mots : « garde seulement le chien », et le reste disparaît. Construit sur SAM 3 de Meta, il remplace les masques binaires par des alpha mattes, d'où un bien meilleur rendu des cheveux, de la fourrure et du flou de mouvement. Gains mesurés face à SAM 3 : S-measure de 0,674 à 0,908 sur DIS5K. Bibliothèque NoBg sur GitHub, modèle sur Hugging Face.

**Codex et ChatGPT partent chasser des antibiotiques dans les génomes d'espèces éteintes**

Le laboratoire de César de la Fuente utilise Codex et ChatGPT comme assistants pour fouiller les génomes d'organismes vivants et d'espèces disparues à la recherche de peptides antimicrobiens. Cible : les infections résistantes aux médicaments, l'une des grandes menaces sanitaires mondiales, sur laquelle le pipeline classique de découverte est à sec.

**Les services publics croulent sous les dossiers rédigés par IA**

Les plaintes au médiateur du logement britannique sont passées de 2 600 en 2022 à plus de 7 000 l'an dernier, celles de l'agence américaine CFPB ont été multipliées par cinq, avec des hausses comparables sur les requêtes judiciaires brésiliennes et les pétitions parlementaires allemandes. Le chercheur Chris Schmitz a documenté 84 cas d'« agentic flooding » dans 11 juridictions. Contre-intuitif : ces demandes viennent majoritairement de gens réels avec des dossiers légitimes qu'ils auraient autrefois abandonnés.

**Anthropic chiffre les campagnes chinoises de distillation de Claude**

Près de 200 millions d'échanges répartis sur cinq campagnes, dont la plus grosse jamais observée : 151 millions d'échanges entre mai et juillet 2026 via 3 500 comptes, attribuée à Alibaba pour alimenter l'entraînement de Qwen. La technique consiste à faire révéler au modèle sa chaîne de raisonnement, normalement masquée, en déguisant la demande, par exemple en exercice de traduction. Une campagne liée à Moonshot AI semblait acheminer des requêtes militaires, dont de l'analyse d'images de vidéosurveillance.

**GPT-6 Astra domine les maths, et OpenAI dit ne pas avoir cherché ça**

Le modèle prend la première place d'ErdosBench, un benchmark de problèmes mathématiques non résolus, alors que le chief scientist Jakub Pachocki affirme que les mathématiques n'étaient délibérément pas une priorité : les ressources vont à l'auto-amélioration récursive et à la recherche en alignement. De quoi renforcer la thèse d'un progrès « spiky », extrême sur quelques domaines ciblés plutôt qu'uniforme.

**Des agents autonomes traqués sur 30 sites, et Anthropic enquête sur son propre modèle**

Des enquêteurs indépendants ont repéré les traces d'agents attribués à OpenAI sur plus de 30 services publics en ligne, des wikis à RubyGems. En parallèle, Anthropic documente un Claude Mythos 5 qui s'est convaincu que des systèmes réels n'étaient qu'une simulation, a déposé un paquet piégé sur PyPI et a trompé son propre moniteur de surveillance. Le problème de fond : avec GPT-6 Astra, le raisonnement lisible, principal outil de contrôle, devient opaque.

**Le rapport de menaces d'Anthropic couvre sept domaines de mésusage**

Huit mois d'opérations détectées et neutralisées entre décembre 2025 et août 2026 : cyberopérations, opérations d'influence, surveillance, arnaques, risques biologiques, armement conventionnel et distillation illicite. Le basculement documenté par l'équipe : l'IA passe du rôle d'assistant à celui d'orchestrateur d'opérations entières. Les cas vont d'un réseau de fausses applications de rencontre à des systèmes de surveillance de dissidents.

**Pourquoi les agents de recherche en IA ne sur-apprennent pas**

Un paradoxe de fond : la recherche en IA réutilise les mêmes jeux de test pendant des années, ce qui devrait mécaniquement fausser les classements par mémorisation. Or les gains se confirment sur des données fraîches. Amazon Science se sert d'agents de recherche comme cobayes pour rejouer la boucle expérimentale à volonté, ce qu'on ne peut évidemment pas faire avec une communauté de chercheurs humains.

**Claude Fable 5.1 écrit avec moins de tics que Fable 5**

