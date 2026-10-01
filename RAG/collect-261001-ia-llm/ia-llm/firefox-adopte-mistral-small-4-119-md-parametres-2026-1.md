---
id: collect-261001-ia-llm/ia-llm/firefox-adopte-mistral-small-4-119-md-parametres-2026-1
title: "firefox-adopte-mistral-small-4-119-md-parametres-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Microsoft", "Mistral", "OpenAI", "Perplexity"]
dates: []
keywords: ["mistral", "apache", "attention", "copilot", "distribution", "embedding", "gemini", "mixture of experts", "open source", "perplexity", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/firefox-adopte-mistral-small-4-119-md-parametres-2026.md
source_anchor: ""
source_lines: [1, 36]
sha256: de38b07d9854ec7ac0e3a191f93ac8c60f69ed6258eae1cceddeb9fb0ee2c089
---

# firefox-adopte-mistral-small-4-119-md-parametres-2026

Le 16 septembre 2026, Mozilla a changé la manière dont des millions de Français vont utiliser leur navigateur. En s’associant à Mistral AI pour intégrer le modèle Mistral Small 4 dans Firefox Smart Window, l’éditeur de Firefox fait un choix inhabituel : plutôt que de suivre Google ou Microsoft vers un fournisseur d’IA américain, il mise sur une entreprise parisienne. La France devient le premier marché hors Amérique du Nord à recevoir cette fonctionnalité en version bêta, avec un support natif en français. Ce choix arrive alors que Mistral vient tout juste de boucler une levée de fonds de 3 milliards d’euros, portant sa valorisation au-delà de 21 milliards d’euros. Deux annonces, un même message : l’IA européenne cherche une vitrine grand public, et Firefox pourrait bien être celle-là.

## Firefox Smart Window : le pari de Mozilla sur l’IA embarquée

Firefox Smart Window n’est pas un simple bouton de résumé de page. C’est un mode de navigation à part entière, distinct des fenêtres classiques et de la navigation privée, dans lequel un assistant IA peut s’appuyer sur les onglets ouverts, l’historique sélectionné et des mémoires configurées par l’utilisateur pour aider à comparer des produits, résumer une recherche ou reprendre un travail interrompu. La fonctionnalité a d’abord été lancée en anglais aux États-Unis et au Canada en août 2026, avant que Mozilla n’annonce, le 16 septembre, son arrivée en France avec une interface entièrement traduite.

Dans son communiqué officiel, Mozilla précise que Mistral Small 4 rejoint la liste des modèles disponibles pour les utilisateurs de Smart Window aux États-Unis et au Canada, tout en étendant l’accès bêta et le support en français aux utilisateurs français (Mozilla, communiqué officiel). Sur le plan produit, Mozilla insiste sur le fait que Smart Window reste multi-modèles : Mistral Small 4 devient une option recommandée, pas un remplacement exclusif des autres fournisseurs d’IA déjà proposés.

La page produit de Mozilla confirme que la France marque la première expansion de Smart Window en dehors de l’Amérique du Nord, et que d’autres marchés européens suivront plus tard dans l’année (Mozilla, page Firefox Smart Window). Ce séquençage n’est pas anodin : Mozilla teste sa fonctionnalité IA la plus ambitieuse depuis des années sur un marché où sa part de navigateur reste supérieure à sa moyenne mondiale.

## Mistral Small 4 : la fiche technique du modèle qui s’installe dans votre navigateur

Mistral Small 4 n’est pas un modèle sorti la semaine dernière pour l’occasion. Mistral AI l’a présenté le 16 mars 2026 comme la synthèse de trois familles de modèles maison : Magistral pour le raisonnement, Pixtral pour le traitement d’images et Devstral pour le code agentique. Sur le plan architecture, il s’agit d’un mélange d’experts (Mixture of Experts) composé de 128 experts dont 4 sont activés par jeton traité, pour un total de 119 milliards de paramètres et 6 milliards de paramètres actifs par jeton, soit environ 8 milliards en comptant les couches d’embedding et de sortie.

Le modèle dispose d’une fenêtre de contexte de 256 000 jetons, suffisante pour analyser plusieurs onglets ouverts, un historique de navigation ou un document long sans le tronquer. Mistral revendique aussi un paramètre de “reasoning_effort” ajustable : les utilisateurs peuvent basculer entre des réponses rapides à faible latence et un mode de raisonnement approfondi selon la tâche. Sur les gains de performance, Mistral annonce une réduction de 40 % du temps de complétion de bout en bout et un débit trois fois supérieur à celui de Mistral Small 3, sur une configuration optimisée pour le débit.

Autre choix stratégique : Mistral Small 4 est distribué sous licence Apache 2.0, une licence open source permissive. Pour une intégration navigateur grand public comme Smart Window, ce choix de licence pèse dans la balance face à des modèles fermés comme ceux d’OpenAI ou de Google, puisqu’il permet à Mozilla d’auditer, d’adapter et potentiellement d’auto-héberger le modèle plutôt que de dépendre uniquement d’une API propriétaire.

## Pourquoi Mozilla a choisi Mistral plutôt qu’OpenAI ou Google

Mozilla explique avoir sélectionné Mistral Small 4 après avoir évalué ses performances spécifiquement pour Smart Window, en accordant une attention particulière à la qualité multilingue. Selon l’éditeur, les deux entreprises ont traité l’adaptation multilingue et multiculturelle comme une caractéristique centrale du modèle, plutôt que comme un ajustement a posteriori d’une expérience pensée d’abord pour l’anglais.

Anthony Enzor-DeMeo, directeur général de Mozilla Corporation, a justifié ce partenariat en expliquant qu’un navigateur ne devrait pas devenir un entonnoir à sens unique vers un seul fournisseur d’IA, et que Firefox doit rester un espace où plusieurs acteurs de l’IA peuvent entrer en concurrence, open source compris (Mozilla, communiqué officiel). De son côté, Arthur Mensch, cofondateur et directeur général de Mistral AI, a présenté ce partenariat comme la rencontre de deux défenseurs de l’open source, dont l’objectif commun est d’apporter confidentialité, contrôle et choix à la navigation web assistée par IA.

Le message de fond est clair : Mozilla redoute qu’une poignée d’entreprises finisse par contrôler simultanément le navigateur, le moteur de recherche, le modèle d’IA et les services associés. En choisissant un partenaire européen et open source plutôt qu’un des trois grands laboratoires américains fermés, Mozilla cherche à préserver une forme de pluralisme technologique, tout en s’associant à une entreprise qui a besoin d’un canal de distribution grand public pour exister face à OpenAI, Google et Anthropic.

## La France, premier marché hors Amérique du Nord

Le choix de la France comme premier marché européen n’a rien d’un hasard commercial. C’est le pays d’origine de Mistral AI, fondée à Paris en 2023, et un terrain où l’argument de la souveraineté numérique résonne particulièrement auprès des institutions publiques comme des entreprises. Mozilla annonce vouloir étendre Smart Window à d’autres pays européens plus tard dans l’année, sans donner de calendrier précis pour l’Allemagne, l’Espagne ou l’Italie.

Concrètement, les utilisateurs français de Firefox peuvent désormais rejoindre la bêta de Smart Window avec une interface intégralement en français, plutôt que de composer avec un assistant fonctionnant uniquement en anglais. Ce détail compte : la plupart des assistants IA intégrés à des navigateurs concurrents proposent d’abord une expérience anglophone avant une localisation complète, ce qui laisse un boulevard commercial à qui livre une expérience française native dès le lancement bêta.

## Chrome, Edge, Opera, Brave : le champ de bataille des navigateurs IA

Firefox n’est pas seul sur ce terrain. Google intègre progressivement Gemini à Chrome, Microsoft pousse Copilot dans Edge, tandis qu’Opera mise sur son assistant Aria et que Brave propose Leo. Perplexity, de son côté, a lancé son propre navigateur, Comet, entièrement pensé autour de son moteur de réponse IA. Chacun de ces acteurs poursuit la même idée : transformer le navigateur en point d’entrée vers un assistant IA, plutôt que de laisser cette fonction aux applications tierces ou aux moteurs de recherche.

