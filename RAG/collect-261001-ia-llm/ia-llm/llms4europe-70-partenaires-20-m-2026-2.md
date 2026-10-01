---
id: collect-261001-ia-llm/ia-llm/llms4europe-70-partenaires-20-m-2026-2
title: "llms4europe-70-partenaires-20-m-2026"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Google", "Mistral", "OpenAI"]
dates: []
keywords: ["agents", "amd", "chatgpt", "fine-tuning", "gemini", "mistral", "open source", "research", "valuation"]
source: docs/RAG/collect-261001-ia-llm/llms4europe-70-partenaires-20-m-2026.md
source_anchor: ""
source_lines: [35, 60]
sha256: c76cecbdbe16519aca781b95aab469398f7bffd4971958c4cc5790909edcb949
---

# llms4europe-70-partenaires-20-m-2026

Pour financer cette adaptation sectorielle, le projet ouvre des appels dits FSTP, pour Financial Support to Third Parties. Concrètement, des PME et organismes tiers non membres du consortium initial pourront candidater au printemps 2026 pour fournir des données linguistiques ou contribuer au fine-tuning et à l’évaluation des modèles. C’est un mécanisme classique de financement en cascade dans les programmes européens, déjà utilisé sur d’autres initiatives Digital Europe, et qui permet d’élargir la base de contributeurs sans multiplier les partenaires officiels du consortium.

## LLMs4Europe face à OpenEuroLLM : deux projets, une même galaxie

La confusion entre LLMs4Europe et OpenEuroLLM, dont *Tech Insider* avait déjà détaillé les ambitions, est fréquente et compréhensible : les deux projets partagent la même galaxie institutionnelle, coordonnée par l’ALT-EDIC, et sont parfois présentés côte à côte dans les documents de la Commission. Mais leurs missions diffèrent nettement. OpenEuroLLM se concentre sur l’entraînement de modèles de fondation multilingues à part entière. Le projet est coordonné par l’Université Charles, à Prague, sous la direction de Jan Hajič, avec Peter Sarlin (AMD Silo AI, Finlande) comme co-responsable, et rassemble vingt institutions de recherche, entreprises et centres de calcul EuroHPC. Son budget total atteint 37,4 millions d’euros, dont 20,65 millions financés par l’UE, et le projet revendique l’ambition de produire la première famille de modèles open source couvrant l’ensemble des langues officielles, actuelles et futures, de l’Union.

Selon un document de présentation officiel de la Commission, OpenEuroLLM est décrit comme “a pan-European initiative that unites leading AI startups and research organisations to create an open, multilingual Large Language Model”, soit une initiative paneuropéenne réunissant des start-ups et organismes de recherche en IA de premier plan pour créer un grand modèle de langage ouvert et multilingue (European Commission, STEP Stories). Le même document précise que le projet s’appuie sur EuroHPC, l’infrastructure européenne de calcul haute performance, pour entraîner des modèles alignés sur les valeurs et la réglementation européennes.

LLMs4Europe, à l’inverse, ne prétend pas entraîner de modèle de fondation depuis zéro. La présentation de la Commission sur la Language Data Space le classe explicitement dans deux catégories d’action : le soutien à la collecte de données linguistiques et le soutien au fine-tuning de grands modèles de langage existants. Autrement dit, là où OpenEuroLLM construit la base, LLMs4Europe l’adapte aux usages professionnels et sectoriels. Une note de présentation attribuée à Petr Sojka, dans le cadre du consortium OpenEuroLLM, résume l’ambition commune de cette galaxie de projets : “Our goal: Open Multilingual European Generative Foundational LLM” (présentation EOSC CZ), soit littéralement l’objectif d’un grand modèle de langage génératif, fondationnel, ouvert et multilingue pour l’Europe. La même présentation insiste sur l’ouverture complète des données (“Open Source (in full) including fully inspectable data”), un point sur lequel la Commission communique volontiers pour se différencier des modèles fermés des grands laboratoires américains.

## Le LLM Institutionnel et la Language Data Space, la pièce qui manquait

Un troisième chantier complète ce puzzle institutionnel. Le 16 juillet 2026, la direction générale de la traduction de la Commission a publié en libre accès son LLM Institutionnel, un modèle interne jusqu’alors réservé aux traducteurs et fonctionnaires européens. Toute entité juridique établie dans un pays de l’UE peut désormais le télécharger via l’European Language Data Space, la même infrastructure de données que doit alimenter LLMs4Europe. Ce modèle n’a pas vocation à rivaliser avec un ChatGPT ou un Gemini en usage grand public, mais il illustre la même logique politique : réutiliser des actifs déjà développés en interne par les institutions européennes plutôt que de payer des licences à des fournisseurs étrangers.

L’articulation entre ces trois briques, LLM Institutionnel, OpenEuroLLM et LLMs4Europe, dessine une stratégie à trois étages assez cohérente sur le papier. Un modèle interne réutilisable pour les administrations, un modèle de fondation ouvert pour les acteurs privés et académiques, et un dispositif de fine-tuning sectoriel pour transformer ces bases génériques en outils opérationnels dans l’énergie, les télécoms ou la santé publique. Reste que ces trois projets sont pilotés par des structures différentes, avec des calendriers différents, ce qui pose la question classique de la coordination entre silos administratifs européens, un défi que la Commission a déjà rencontré sur d’autres dossiers technologiques.

## Pourquoi maintenant ? L’AI Act, la souveraineté et la dépendance américaine

Le calendrier de communication de LLMs4Europe n’est pas neutre. Il intervient trois semaines et demie après l’entrée en application, le 2 août 2026, des obligations de transparence de l’article 50 de l’AI Act, qui impose désormais aux fournisseurs de chatbots et de systèmes génératifs opérant sur le marché européen de signaler qu’un utilisateur échange avec une IA et d’étiqueter les contenus générés artificiellement, deepfakes compris. Ces obligations s’accompagnent de sanctions pouvant atteindre 3 % du chiffre d’affaires mondial en cas de manquement. La même date marque aussi l’entrée en application des règles pour les systèmes à haut risque de l’annexe III du règlement, avec obligation d’enregistrement dans la base de données européenne et de marquage CE pour les entreprises françaises et européennes concernées.

Ce contexte réglementaire renforce mécaniquement l’argument en faveur de modèles développés selon des règles européennes dès leur conception, plutôt qu’adaptés a posteriori à la conformité. Un LLM entraîné et fine-tuné sous supervision d’un consortium public européen, avec des données dont la provenance est documentée, limite les zones d’incertitude juridique qui pèsent sur les déploiements de modèles étrangers dans des usages à haut risque, notamment dans les services publics, l’un des cinq secteurs ciblés par LLMs4Europe. La France a d’ailleurs illustré cette tension récemment en écartant un projet impliquant OpenAI au profit de solutions jugées plus souveraines, un épisode que *Tech Insider* avait couvert en détail.

## Mistral AI et le secteur privé européen : alliés ou rivaux de Bruxelles ?

La stratégie publique de Bruxelles ne se déploie pas dans le vide : elle coexiste avec un secteur privé européen déjà structuré autour de Mistral AI. La start-up française a confirmé en août 2026 des projets d’investissement pour construire jusqu’à un gigawatt de capacité de calcul en Europe d’ici 2030, dans le cadre de sa propre stratégie de souveraineté, avec l’introduction de “Regional Endpoints” permettant de contrôler géographiquement où s’exécute l’inférence, ainsi qu’un niveau de service prioritaire pour ses clients entreprise. L’État français a par ailleurs choisi de déployer les modèles Mistral pour équiper jusqu’à un million d’agents publics, un contrat que *Tech Insider* avait détaillé au moment de l’annonce.

