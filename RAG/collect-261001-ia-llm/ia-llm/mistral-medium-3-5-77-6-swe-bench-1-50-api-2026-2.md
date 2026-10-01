---
id: collect-261001-ia-llm/ia-llm/mistral-medium-3-5-77-6-swe-bench-1-50-api-2026-2
title: "mistral-medium-3-5-77-6-swe-bench-1-50-api-2026"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Apple", "DeepSeek", "EU", "Google", "Huawei", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["mistral", "agent", "agents", "ascend", "aws", "claude", "deepseek", "gpu", "opus 4", "sandbox", "valuation"]
source: docs/RAG/collect-261001-ia-llm/mistral-medium-3-5-77-6-swe-bench-1-50-api-2026.md
source_anchor: ""
source_lines: [34, 75]
sha256: 6f8f87e3f8f66e1c7fc94ee037fb838cc7efcc7ec42667f53841259618fd881b
---

# mistral-medium-3-5-77-6-swe-bench-1-50-api-2026

Le deuxième volet de l’annonce est **Vibe**, l’agent de codage à distance de Mistral. Vibe existait déjà depuis fin 2025 sous forme d’interface en ligne de commande propulsée par Devstral 2. Avec la mise à jour du 29 avril, Vibe abandonne Devstral 2 au profit de Medium 3.5 et bascule sur une architecture *remote-first* : l’agent s’exécute désormais dans une sandbox cloud isolée, accède au dépôt Git via une connexion sécurisée, et peut être piloté à distance depuis n’importe quel terminal, y compris un téléphone mobile.

Ce positionnement entre directement en concurrence avec **Claude Code** d’Anthropic, **Cursor** et le récent agent autonome de **Replit**. La différence stratégique tient en deux points. D’abord, Vibe propulsé par Medium 3.5 affiche un coût d’exploitation divisé par dix par rapport à Claude Code, ce qui ouvre les agents de codage à distance à des équipes intermédiaires (PME tech, ESN, agences). Ensuite, Vibe peut être déployé en self-hosted sur un cluster on-premises de quatre GPU, une option inédite chez les concurrents américains qui imposent tous une exécution cloud propriétaire.

Concrètement, Vibe accepte des instructions en langage naturel – par exemple « ajoute la pagination cursor-based à l’endpoint /products et écris les tests Jest correspondants » – et produit une *pull request* complète, testée, documentée, prête à être révisée. Lors des démonstrations internes diffusées par Mistral, l’agent a résolu 8 issues GitHub réelles consécutives sur un dépôt Django de 47 000 lignes, sans intervention humaine, en moins de 38 minutes au total.

## Le Chat Work mode : multi-étapes et productivité de bureau

Le troisième volet est le **Work mode** introduit dans **Le Chat**, l’assistant grand public et professionnel de Mistral. Work mode permet à Le Chat d’orchestrer des tâches multi-étapes complexes : recherche web, navigation sur des sites authentifiés, manipulation de documents bureautiques (Word, Excel, Google Workspace), génération de présentations, envoi d’emails via un compte connecté. Là où le mode standard de Le Chat répondait à une question, Work mode *exécute* une mission.

La promesse fonctionnelle ressemble à celle du *Computer Use* d’Anthropic ou de l’*Operator* d’OpenAI, mais avec une différenciation européenne assumée : l’ensemble des données traitées par Work mode transitent par des centres de données situés dans l’Union européenne, ce qui rend la fonctionnalité immédiatement compatible avec le RGPD et le futur AI Act dont l’application a été repoussée au 2 décembre 2027 dans le cadre du Digital Omnibus. Pour les directions juridiques des grandes entreprises françaises, c’est un argument différenciant majeur, et c’est précisément la cible visée par Mistral.

### Réactions du marché et premiers déploiements

Dans les heures qui ont suivi l’annonce, plusieurs grandes entreprises européennes ont indiqué publiquement leur intérêt pour l’évaluation de Medium 3.5 et Work mode dans leurs pilotes en cours. La capitalisation boursière des cotés européens de la chaîne IA – ASML, ASM International, BE Semiconductor – a affiché des hausses modérées en séance, signe que le marché interprète positivement la consolidation de l’écosystème logiciel européen autour d’un champion crédible. Le titre Microsoft, en revanche, est resté stable, indiquant que les investisseurs ne perçoivent pas encore Medium 3.5 comme une menace directe pour OpenAI sur le marché américain.

## Pourquoi le timing du 29 avril 2026 est stratégique pour Mistral

Le calendrier du lancement n’est pas neutre. Le 29 avril 2026 intervient **onze jours** après le séisme Apple Q2 2026 publié le 18 avril (111,2 Md$ de revenus, +17 %), **deux semaines** après les résultats d’Alphabet Q1 2026 marqués par la surprise du segment Google Cloud (+63 %), et **douze jours** après la révélation par le Wall Street Journal du manque à gagner de revenus 2026 chez OpenAI dans le contexte du méga-pari Stargate à 1 400 milliards de dollars. Dans ce paysage où chaque acteur américain communique des chiffres titanesques, Mistral choisit délibérément de jouer la carte du produit utilisable, du modèle ouvert et du coût maîtrisé.

Le lancement coïncide également avec la dernière ligne droite de **VivaTech 2026**, le salon technologique européen qui se tiendra à Paris Porte de Versailles à la mi-juin. Mistral a confirmé qu’Arthur Mensch tiendrait la keynote d’ouverture du salon. Dévoiler Medium 3.5 six semaines avant ce rendez-vous médiatique permet au modèle d’accumuler des retours d’usage, des intégrations partenaires et des cas clients publics qui pourront être mis en avant à VivaTech. La séquence rappelle la stratégie pré-Dreamforce de Salesforce ou pré-re:Invent d’AWS.

## Le positionnement souverain : argument différenciant pour l’Europe

L’angle **souveraineté** est désormais au cœur du discours commercial de Mistral. Dans un contexte où Bruxelles durcit l’application des règles numériques européennes en 2026 et où la perspective de représailles tarifaires américaines contre les services numériques continentaux pèse sur les arbitrages d’achat des grandes entreprises et administrations, l’origine française de Mistral, la possibilité d’auto-hébergement, et la disponibilité des poids ouverts forment un trio d’arguments puissant. Selon les analystes du *French Tech Journal*, le baromètre VivaTech 2026 indique que la souveraineté est devenue la première préoccupation tech des dirigeants européens, devant la productivité.

Mistral capitalise sur ce momentum réglementaire. La mise à disposition des poids du modèle dès le jour du lancement, sous une licence MIT modifiée permissive pour l’inférence en interne, satisfait les exigences les plus strictes des départements DSI publics. Les administrations françaises, encadrées par le programme *French Tech Souverain* et les directives de la DINUM, disposent depuis longtemps de freins à l’achat de licences cloud-only américaines ; Medium 3.5 leur offre une option crédible, performante, à un prix prédictible.

## Le pari économique : 1,50 $ contre 15 $ par tâche d’agent

La grille tarifaire de Medium 3.5 – 1,50 $ en entrée et 7,50 $ en sortie par million de tokens – est l’arme commerciale principale de Mistral. À titre de comparaison, Claude Opus 4.7 facture 15 $ en entrée et 75 $ en sortie, soit un ratio de 10× sur la sortie. GPT-5.5 d’OpenAI s’établit à 5 $ en entrée et 15 $ en sortie, soit un ratio de 2× sur la sortie pour des performances comparables sur SWE-Bench. Seul DeepSeek V4 affiche un tarif inférieur (0,55 $ / 1,74 $), mais avec une exécution principalement sur infrastructure Huawei Ascend en Chine, ce qui est incompatible avec les politiques d’achat de la plupart des entreprises occidentales.

### Tableau du coût total de possession pour un déploiement type

| Cas d’usage | Volume mensuel | Mistral Medium 3.5 | Claude Opus 4.7 | GPT-5.5 | Économie vs Claude | 
|---|---|---|---|---|---|
| Chatbot support N1 (50K conv.) | 500 M tokens E / 100 M S | 1 500 $ | 15 000 $ | 4 000 $ | −90 % | 
| Agent codage équipe 50 dev. | 1 Md tokens E / 200 M S | 3 000 $ | 30 000 $ | 8 000 $ | −90 % | 
| Analyse documents juridiques | 5 Md tokens E / 50 M S | 7 875 $ | 78 750 $ | 25 750 $ | −90 % | 
| Recherche scientifique R&D | 2 Md tokens E / 300 M S | 5 250 $ | 52 500 $ | 14 500 $ | −90 % | 
| Self-hosted (4× H100) | Illimité | ~12 000 $ (amort.) | Non disponible | Non disponible | n/a | 

