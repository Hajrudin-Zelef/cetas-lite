---
id: collect-261001-ia-llm/ia-llm/il-raconte-a-chatgpt-son-projet-de-meurtre-l-ia-appelle-le-fbi-3
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "ByteDance", "Microsoft", "Nvidia", "OpenAI", "SpaceX", "Stripe"]
dates: []
keywords: ["agent", "agents", "attention", "chatgpt", "claude", "foundry", "nvidia", "opus 4"]
source: docs/RAG/collect-261001-ia-llm/il-raconte-a-chatgpt-son-projet-de-meurtre-l-ia-appelle-le-fbi.md
source_anchor: ""
source_lines: [154, 199]
sha256: 7da2165c1c0c63fa45b7bb8246bece90bd3980d93ba6d012f5c4ba4ec93e2cc7
---

# 🧠 **RECHERCHE**

Sur la réplication d'articles scientifiques jamais vus à l'entraînement, un agent de seulement **27 milliards de paramètres** devance Claude Opus 4.8 et GPT-5.5. Sa recette : il ne code pas, il pilote. Le modèle est entraîné à raisonner comme un chercheur, à décider quoi tester et dans quel ordre, puis il sous-traite l'implémentation à des modèles plus forts en code. Les auteurs proposent avec Replica un espace de tâches à grande échelle pour entraîner ce comportement. L'implication est majeure : un seul modèle géant n'est peut être pas nécessaire pour faire de la science.

# **🗞️PLUS D'ACTUALITÉS**

**Un ingénieur de Cursor offre 200 $ de crédits à qui résilie son abonnement Claude**

Le 17 août, un ingénieur de Cursor propose sur X, à titre strictement personnel et sans validation de sa hiérarchie, **200 dollars de crédits Ultra** à quiconque publie une capture d'écran de résiliation de son abonnement Claude. Le post devient viral en quelques heures, des développeurs confirment avoir reçu les crédits, et un bot ne parvient pas à écouler la file en 16 heures. L'opération tombe trois jours après le rachat de Cursor par SpaceX et juste après une nouvelle panne d'Anthropic, la 166e de l'année selon le décompte relayé par AI Secret. Pas une campagne marketing, juste un employé avec une carte de crédit et un timing redoutable.

**Claude Code garde +50% de limite hebdomadaire jusqu'au 31 août**

Anthropic prolonge sa promotion : la **limite d'usage hebdomadaire de Claude Code reste majorée de 50%** jusqu'au **31 août 2026**, sans aucune action à faire. L'offre couvre les plans Pro, Max, Team et les sièges Enterprise historiques, mais pas les comptes gratuits ni les sièges Enterprise facturés à la consommation. Elle s'applique automatiquement sur le CLI, les extensions IDE, le desktop et le web, sans changement de facturation. Attention au détail qui compte : seules les limites hebdomadaires sont concernées, les limites de 5 heures restent inchangées. Onze jours pour en profiter.

**Cursor lance Origin, un GitHub concurrent, le jour d'une panne mondiale de GitHub**

Cursor, désormais filiale de SpaceX, dégaine **Origin**, une plateforme d'hébergement de code qui reprend l'essentiel de ce que les développeurs font sur GitHub : parcourir et éditer un dépôt, collaborer, gérer les pull requests. La synchronisation est interopérable, on peut donc travailler en parallèle des deux côtés sans migrer. Cursor annonce des fonctionnalités pensées nativement pour les agents et un écosystème d'applications autour. Le timing est cruel pour Microsoft : GitHub a enregistré **257 pannes en un an** selon LeadDev, dont une interruption mondiale de plus de six heures le jour même du lancement d'Origin.

**Databricks publie sa méthode pour maîtriser la facture de l'IA de codage**

Avec les retours de Stripe, Coinbase, Uber et Ramp, Databricks documente le mur que rencontrent toutes les entreprises qui déploient l'IA de codage à grande échelle : **des coûts qui croissent exponentiellement** au point de menacer les gains de productivité obtenus. Le levier principal identifié n'est pas de réduire l'usage mais de changer de cible : suivre la "frontière d'efficacité", c'est à dire le meilleur rapport prix sur intelligence, plutôt que la frontière d'intelligence pure. Databricks a open-sourcé les deux outils internes qui appliquent cette logique, Omnigent et Unity AI Gateway.

**OpenAI ralentit volontairement ses modèles à cause de leurs capacités en cyberattaque**

OpenAI annonce renforcer le monitoring, l'alignement et la sécurité de ses modèles de pointe, et surtout acte un principe nouveau : les capacités offensives en cybersécurité deviennent un facteur qui **encadre le rythme de déploiement** des futurs modèles. Autrement dit, un modèle trop compétent pour attaquer des systèmes pourra voir sa sortie retardée ou restreinte. La note est signée de l'entreprise elle même, ce qui en fait un engagement public, à surveiller sur les prochaines sorties.

**L'humanoïde Superman d'Unitree court plus vite qu'Usain Bolt**

Unitree a présenté Superman, un robot humanoïde conçu en **trois mois**, qui atteint une vitesse de pointe de **12,66 m/s** sur des jambes de seulement 0,85 mètre, soit plus rapide que le record humain d'Usain Bolt. Il affiche aussi un **saut vertical de 2 mètres**, largement au dessus de tout ce qu'un athlète peut produire. L'entreprise avait déjà repris le record de vitesse à 10 m/s plus tôt cette année, puis engagé une machine au marathon de Pékin. Unitree affirme qu'il reste une marge de progression significative.

**La Chine laisse ByteDance et Tencent importer 10 000 puces H200 chacune**

Selon le Financial Times, Pékin a assoupli ses propres restrictions et laissé entrer sur le continent **10 000 puces NVIDIA H200 pour ByteDance et autant pour Tencent** ces dernières semaines, avec un plafond autorisé de 100 000 unités chacune. Objectif : permettre aux groupes chinois d'entraîner des modèles de pointe face aux Américains. Rappel du contexte : Washington avait interdit la vente des H200 à la Chine avant d'autoriser certains clients approuvés en décembre 2025. Subtilité révélatrice, Pékin demanderait aux entreprises de garder la majorité des puces hors du continent, livrées à Hong Kong, pour ne pas décourager sa filière de semi-conducteurs domestique.

**Warp Factories, une "usine logicielle" IA livrée clé en main**

Warp lance une infrastructure prête à l'emploi pour déployer et piloter une flotte d'agents IA le long des étapes classiques du développement : triage, spécification, implémentation, revue, vérification. Le système est agnostique côté modèle, il fonctionne avec Codex comme avec Claude Code, et s'intègre à Linear, Jira, Slack et Teams. La cible assumée, ce sont les petites structures qui n'ont pas les moyens de construire cette architecture elles mêmes. Warp affirme automatiser **30 à 35% de ses propres tâches de développement** chaque semaine grâce à ce dispositif.

**Des agents IA prennent en main la paperasse du transport routier**

Alvys, éditeur de logiciels de fret, lance Alvys Foundry, une plateforme d'agents intégrée directement à son système de gestion du transport. Elle embarque **plus de 20 modèles d'agents prêts à l'emploi** pour les tâches les plus ingrates du métier : suivi des expéditions, gestion documentaire, audit de factures et de tarifs, conformité, sinistres. Les opérateurs peuvent créer leurs propres agents en décrivant le besoin en langage naturel, ou à partir d'une procédure interne existante. Chaque agent se teste sur données simulées avant déploiement, puis se pilote, se monitore et se met en pause depuis la plateforme.

**L'IA ne guérit pas le "théâtre du travail"**

Thèse défendue et beaucoup commentée dans la communauté tech : le vrai problème des grandes entreprises n'est pas la lenteur d'exécution mais le fait de construire les mauvaises choses, et l'IA accélère précisément la seule moitié qui n'était pas le goulot. Le "théâtre du travail" désigne ces projets qui sonnent très bien dans un ticket Jira mais n'apportent rien au client final. L'auteur avance une ligne de fracture nette : au delà de **quatre niveaux hiérarchiques**, les gains de l'IA agentique deviennent faibles voire négatifs, alors qu'à **trois niveaux ou moins** l'accélération est réelle. L'IA excelle à exécuter, elle n'aide pas à décider quoi construire.

**ChatGPT Ads débarque dans 31 marchés européens**

