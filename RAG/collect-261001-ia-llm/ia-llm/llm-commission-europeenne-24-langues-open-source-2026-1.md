---
id: collect-261001-ia-llm/ia-llm/llm-commission-europeenne-24-langues-open-source-2026-1
title: "llm-commission-europeenne-24-langues-open-source-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Cohere", "EU", "Google", "Mistral", "OpenAI", "United States"]
dates: []
keywords: ["agents", "chatgpt", "gemini", "incident", "llama", "mistral", "open source"]
source: docs/RAG/collect-261001-ia-llm/llm-commission-europeenne-24-langues-open-source-2026.md
source_anchor: ""
source_lines: [1, 41]
sha256: bdbad9049261f5387ac2b203c1ec9b67891f1ad4388368b166ae795ea801d47e
---

# llm-commission-europeenne-24-langues-open-source-2026

La Commission européenne a mis en ligne le 16 juillet 2026 son propre grand modèle de langage, baptisé sobrement **EU Institutional LLM**. Développé par la Direction générale de la traduction (DGT), ce modèle ouvert couvre les 24 langues officielles de l’Union et vise un public précis : les administrations, chercheurs et entreprises basés dans un pays membre. L’annonce passe presque inaperçue au milieu de l’actualité IA de l’été, éclipsée par le bras de fer entre la France et OpenAI ou par l’entrée en vigueur de l’AI Act le 2 août. Elle marque pourtant un tournant : Bruxelles ne se contente plus de réguler l’intelligence artificielle, elle en construit désormais elle-même.

Ce lancement s’inscrit dans un calendrier chargé pour la Commission, entre le paquet de souveraineté technologique présenté le 3 juin 2026, la sélection du consortium EUROPA le 19 juin, et l’entrée en application des obligations de transparence de l’AI Act début août. Trois initiatives, trois échelles, un même objectif affiché : réduire la dépendance européenne aux modèles américains. Voici ce que révèlent les chiffres, les dates et les prises de position officielles.

## Un nouveau modèle IA made in Bruxelles

L’EU Institutional LLM existe en deux versions : un modèle de base et une variante instruct, optimisée pour suivre des consignes. La DGT le présente comme un modèle ouvert, mais avec une nuance importante : l’accès n’est pas universel. Seules les entités légales établies dans un pays de l’UE peuvent le télécharger, via l’European Language Data Space, la plateforme d’échange de données linguistiques de la Commission. Ce choix distingue clairement le projet des modèles ouverts au sens large, type Llama ou Mistral, accessibles à n’importe qui dans le monde.

Sur le plan technique, la transparence reste limitée. Aucun document officiel consulté à ce jour ne précise le nombre de paramètres du modèle, ni le modèle de base sur lequel il a été entraîné. Pour une administration qui prépare, au même moment, l’application des obligations de documentation de l’AI Act, l’écart entre le discours et la pratique interroge. La DGT justifie sa démarche par un objectif de service public plutôt que de performance brute : comprendre la terminologie, la législation et le contexte politique propres à l’Union.

## Pourquoi la Commission investit dans les langues européennes

La justification officielle tient en une formule frappante : éviter “l’extinction numérique” des langues moins représentées dans les corpus d’entraînement des grands modèles commerciaux. Le maltais, l’irlandais ou le letton pèsent peu face à l’anglais, au chinois ou à l’espagnol dans les jeux de données utilisés par OpenAI, Google ou Anthropic. Un modèle entraîné à parité sur les 24 langues officielles répond à un besoin que les acteurs privés n’ont, jusqu’ici, pas jugé prioritaire.

La DGT insiste sur un second usage : la compréhension fine du jargon juridique et administratif propre aux institutions européennes. Un modèle généraliste américain traduit correctement une phrase, mais peine parfois à restituer la nuance d’un règlement ou d’une directive. C’est ce créneau, plus étroit qu’un modèle frontière, que vise l’EU Institutional LLM : un outil de travail pour les traducteurs, juristes et fonctionnaires de l’Union, pas un concurrent direct de ChatGPT ou Gemini.

## GPT@EC, l’autre pilier de la stratégie IA interne

L’EU Institutional LLM ne sort pas de nulle part. Depuis le 22 octobre 2024, la Commission dispose déjà d’un outil d’IA générative interne, GPT@EC, développé par la Direction générale de l’informatique (DIGIT) et ouvert à l’ensemble du personnel. Présenté comme un assistant capable d’aider à la rédaction, au résumé et à la génération d’idées, GPT@EC s’appuie sur plusieurs modèles de langage, sans se limiter à une seule technologie propriétaire. Un document de la Commission daté du 8 octobre 2025 annonçait déjà une première publication open source de l’outil prévue courant 2026.

La distinction entre les deux projets compte. GPT@EC fonctionne comme une couche d’orchestration, une interface sécurisée qui peut mobiliser différents modèles selon la tâche. L’EU Institutional LLM, lui, est un modèle en tant que tel, publié et téléchargeable. Ensemble, ils dessinent une architecture à plusieurs niveaux : un assistant pour les agents publics d’un côté, une brique linguistique ouverte de l’autre.

## EU Institutional LLM, EUROPA, OpenEuroLLM : trois projets, trois échelles

La Commission ne mise pas sur un seul cheval. Trois initiatives coexistent aujourd’hui, avec des objectifs et des budgets très différents. OpenEuroLLM, mis en avant dès le 3 février 2025, dispose d’un budget total de 37,4 millions d’euros, dont 20,6 millions financés par le programme Digital Europe. EUROPA, sélectionné le 19 juin 2026 comme lauréat du Frontier AI Grand Challenge, vise un modèle de plus de 400 milliards de paramètres couvrant nativement les 24 langues officielles, avec pour récompense un accès stratégique pouvant atteindre 2,5 % de la capacité totale d’EuroHPC pendant un an, plutôt qu’une subvention en numéraire.

| Projet | Porteur | Statut (août 2026) | Échelle / budget | Objectif principal | 
|---|---|---|---|---|
| EU Institutional LLM | DG Traduction (DGT) | Publié le 16 juillet 2026 | Paramètres non divulgués | Traduction et usage institutionnel, 24 langues | 
| GPT@EC | DG Informatique (DIGIT) | Déployé depuis oct. 2024, open source prévu en 2026 | Assistant multi-modèles | Productivité interne des agents de la Commission | 
| OpenEuroLLM | Consortium académique/industriel | Lancé fév. 2025, Sceau STEP obtenu | 37,4 M€ (dont 20,6 M€ Digital Europe) | Modèle ouvert multilingue européen | 
| EUROPA | Consortium sélectionné par la Commission | Désigné lauréat le 19 juin 2026 | 400+ Md paramètres, jusqu’à 2,5 % d’EuroHPC | Modèle frontière européen concurrent des géants US | 

Cette juxtaposition traduit une méthode plus qu’une improvisation : miser sur plusieurs paris technologiques et budgétaires en parallèle, du modèle institutionnel déjà livré à l’ambition frontière encore en construction. Reste une question de fond, souvent posée par les analystes du secteur : cette dispersion entre quatre initiatives distinctes accélère-t-elle vraiment la course, ou dilue-t-elle des moyens déjà modestes face aux investissements américains et chinois ?

## Le déclencheur politique : la France écarte OpenAI

Le contexte français explique en partie l’accélération du calendrier européen. Le 18 août 2026, le ministre du Budget David Amiel a annoncé que l’État français recourrait désormais à des “fournisseurs d’intelligence artificielle souverains”, citant Mistral, pour tester la sécurité de ses systèmes publics, tout en excluant explicitement OpenAI de ces missions sensibles. Cette décision fait suite à une fuite de données touchant environ 700 000 contribuables au sein de l’administration fiscale française, un incident qui a poussé Paris à décider que les systèmes traitant des données sensibles migreraient vers des fournisseurs d’IA certifiés souverains, Mistral et l’allemand Aleph Alpha en tête de liste.

Le parallèle avec l’EU Institutional LLM n’est pas anodin. Dans les deux cas, une administration publique choisit consciemment un modèle développé ou hébergé en Europe plutôt qu’une solution américaine, pour des raisons de contrôle des données autant que de principe politique. La différence tient à l’échelle : la France mise sur un acteur commercial déjà mature, Mistral, quand la Commission construit son propre outil en interne, via la DGT.

