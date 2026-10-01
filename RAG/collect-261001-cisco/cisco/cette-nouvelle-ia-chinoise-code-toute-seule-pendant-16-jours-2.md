---
id: collect-261001-cisco/cisco/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours-2
title: "🧠 **RECHERCHE**"
domain: cisco
role: reference
task: reference
actors: ["Anthropic", "Google", "Hugging Face", "Microsoft", "Nvidia", "OpenAI"]
dates: []
keywords: ["agent", "agentic", "agents", "attention", "claude", "cyber", "distillation", "gemini", "gpt-live", "nvidia", "open source", "open-weight"]
source: docs/RAG/collect-261001-cisco/cette-nouvelle-ia-chinoise-code-toute-seule-pendant-16-jours.md
source_anchor: ""
source_lines: [65, 119]
sha256: b5a3819a62eaaf2f1bfbcb64cad38fb75b61ec4af6559c40b9455871c6eabd35
---

# 🧠 **RECHERCHE**

**Le mode vocal de Claude passe sous Opus et agit sur vos applications**

Jusqu'ici cantonné à Haiku, rapide mais limité pour du raisonnement soutenu, le mode vocal de Claude peut désormais tourner sous **Opus, Sonnet ou Haiku**, avec bascule en cours de conversation sans perte de contexte entre l'oral et l'écrit. Plus intéressant encore, la voix déclenche des actions réelles sur **Gmail, Calendar, Slack, Canva et Notion** : déplacer une réunion, rédiger un email, résumer un fil, créer un document, toujours après demande d'autorisation. Le mode vocal couvre maintenant **10 langues**.

**Microsoft Research publie Orchard, son framework open source pour entraîner des agents**

Orchard repose sur Orchard Env, un service d'environnements réutilisables qui permet d'entraîner et d'évaluer des agents directement dans les harnais où ils tourneront pour de vrai : Codex, OpenClaw, ZeroClaw. La même infrastructure sert pour du développement logiciel, de la navigation web et de l'assistance personnelle. Le modèle Orchard-SWE atteint **69,7 % sur SWE-bench Verified** (73 % avec reranking par modèle de valeur) avec seulement **3 milliards de paramètres actifs**, soit dix fois moins que les systèmes propriétaires comparables. Données et méthodes d'entraînement sont publiées.

**Des chercheurs militaires chinois distillent les modèles américains**

Une enquête de Reuters portant sur plus de **80 articles et brevets chinois** montre que des chercheurs liés à l'armée utilisent les sorties de modèles OpenAI et Anthropic pour entraîner des systèmes plus petits, tournant en local. Applications visées : surveillance, opérations cyber, analyse de code et usages tactiques. La technique, la distillation, transfère des capacités ciblées, mais Reuters rappelle qu'elle ne reproduit ni la pleine puissance d'un modèle frontière, ni la capacité de calcul nécessaire pour en entraîner un.

**Gérer la mémoire d'un agent comme un cycle de vie, pas comme un stockage**

Ce papier propose l'"Agentic Context Management", un cadre en cinq temps pour décider ce qu'un agent doit **garder, récupérer, partager, préparer et compresser** dans son contexte. Le constat de départ est concret : les agents en production échouent moins par manque d'intelligence que par saturation, quand vieux messages, descriptions d'outils et sorties obsolètes engorgent la fenêtre. Stocker davantage ne suffit pas, il faut des règles de pertinence, d'accès, de timing et de compression.


# **🗞️PLUS D'ACTUALITÉS**

**Le PDG de Hugging Face : "la Chine domine clairement les modèles ouverts"**

Clément Delangue a déclaré lundi sur CNBC que la Chine était en train de gagner la course à l'IA, et qu'elle pourrait rattraper la frontière technologique américaine dès la fin 2026 ou en 2027. Il pointe un écosystème chinois de collaboration ouverte face à des labos américains qui "construisent en silos". Il est aussi revenu sur le piratage de Hugging Face par des agents OpenAI échappés d'un environnement d'entraînement le mois dernier, tout en assurant que la collaboration entre les deux entreprises reste saine. Microsoft, Palantir et Nvidia ont par ailleurs signé une lettre appelant à ne pas restreindre les modèles open-weight.

**Un robot centaure à tête de bouc, avec des tronçonneuses à la place des mains**

La startup californienne Satyress a construit Threehalves, une machine de plus d'1m80 qui combine un torse humanoïde et une base à quatre pattes freinée par friction. Ses bras se terminent par un connecteur rapide alimenté en 12v, 18v ou 48v, capable d'accueillir différents outils, dont des tronçonneuses. La cible : fronts d'incendie, zones industrielles toxiques, glissements de terrain et fouilles de décombres. Ce n'est pas un robot autonome, un opérateur le pilote au joystick.

**Des drones au poivre déployés dans des écoles américaines**

La Floride, la Géorgie et le Colorado financent des programmes pilotes, jusqu'à **557 000 dollars**, pour équiper des établissements de drones destinés à intercepter un tireur actif en une quinzaine de secondes. Capables d'atteindre **60 mph** et de traverser une fenêtre, ils sont pilotés à distance depuis Austin par des opérateurs de Mithril Defense, et non autonomes. Des experts en sécurité scolaire s'alarment du risque d'erreur d'identification et du fait que le dispositif détourne l'attention des vrais sujets.

**L'IA a conquis le code, elle attaque le drive**

Après les débuts catastrophiques et viraux de 2023, les IA vocales se sont fiabilisées et s'installent massivement dans la restauration rapide américaine. Taco Bell en a équipé **plus de 890 files de drive**, soit plus de 10 % de ses restaurants aux États-Unis. Dairy Queen déploie dans **25 États** avec l'objectif de couvrir tout son parc, et White Castle fait tourner son assistante "Julia" sur 12 % de ses sites, systématiquement installée dans les nouvelles ouvertures. McDonald's avait lancé le mouvement dès 2019 en rachetant Apprente.

**Anthropic a supprimé 80 % du prompt système de Claude Code, sans rien perdre**

Un ingénieur d'Anthropic explique que l'entreprise a retiré **plus de 80 %** des instructions système de Claude Code sans aucune dégradation sur les évaluations de code. La raison de la longueur d'origine : les premiers modèles avaient besoin de garde-fous rigides pour éviter les suppressions de fichiers ou les commentaires non professionnels. L'équipe savait que ces règles seraient parfois contre-productives, mais le compromis valait le coup. Avec des modèles capables de juger seuls d'une situation, il ne tient plus.

**OpenAI dévoile GPT-Live, la voix continue sans tour de parole**

OpenAI a publié les coulisses de GPT-Live, un système d'interaction vocale continue développé en six mois. Il repose sur un modèle de parole dit "turnless" et une architecture à faible latence, censée supprimer les coupures caractéristiques des assistants vocaux classiques, où il faut attendre la fin d'une phrase pour obtenir une réponse. L'objectif affiché est une conversation qui ressemble à une vraie conversation.

**Fish Audio clone une voix en quelques secondes, en direct**

La startup a lancé publiquement son modèle phare S2.1 Pro, dont les fondateurs ont fait la démonstration en clonant leur propre voix en temps réel devant leur audience. L'annonce s'accompagne d'une levée de **52 millions de dollars** en seed, pour clore sa première année d'existence. À rapprocher des nouvelles obligations européennes d'étiquetage entrées en vigueur cette semaine.

**Gemini Spark navigue dans Chrome avec vos comptes connectés**

L'assistant agentique de Google s'intègre désormais au navigateur et peut "utiliser vos comptes connectés et mots de passe enregistrés pour gérer les corvées du web". Les exemples cités par Google : rechercher des options de vol et amorcer une réservation, ou prendre rendez-vous pour visiter un appartement. Google annonce des défenses contre l'injection de prompt et assure que Spark ne finalise jamais un paiement sans validation. Déploiement d'abord aux États-Unis, tandis que Google AI Pro est désormais disponible dans plus de 160 pays.

**Google retire la génération d'images de Google Earth au bout de 24 heures**

Lancée le 30 juillet, la fonctionnalité Nano Banana 2 permettait de générer des images ancrées dans les vues satellite, aériennes et 3D de Google Earth : reconstitutions historiques, infographies éducatives, projets immobiliers, visualisation d'aménagements urbains. Dès le 31 juillet, Google faisait machine arrière après avoir vu circuler des captures d'écran contraires à ses règles. Les images étaient pourtant filigranées et n'apparaissaient pas dans l'expérience principale. "Les gens font une confiance particulière à Google Earth pour une vue fiable du monde", explique l'entreprise.

