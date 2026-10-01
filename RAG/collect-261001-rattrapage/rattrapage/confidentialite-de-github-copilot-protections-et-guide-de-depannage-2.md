---
id: collect-261001-rattrapage/rattrapage/confidentialite-de-github-copilot-protections-et-guide-de-depannage-2
title: "confidentialite-de-github-copilot-protections-et-guide-de-depannage"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot", "agent", "agents", "attribution"]
source: docs/RAG/collect-261001-rattrapage/confidentialite-de-github-copilot-protections-et-guide-de-depannage.md
source_anchor: ""
source_lines: [86, 206]
sha256: 0dd041018a6f211ec5e60a4074773c92e69d57438062a859ff983c84aac68b45
---

# confidentialite-de-github-copilot-protections-et-guide-de-depannage

Pour vérifier que les exclusions fonctionnent : ouvrez un fichier exclu et demandez à Copilot Chat d'« expliquer ce fichier ». Si Chat fournit une réponse pertinente sur son contenu, l'exclusion ne s'applique pas. C'est le signal pour recharger l'extension et vérifier la syntaxe de la règle.

### Limitations à connaître

En utilisant les exclusions de contenu GitHub Copilot, gardez à l'esprit :

- Copilot CLI, le mode Agent et les Cloud Agents ne respectent pas les exclusions (mi-2026).
- Fuite sémantique : des infos de typage et définitions au survol issues de fichiers exclus peuvent encore influer indirectement sur les suggestions.
- Les symlinks et systèmes de fichiers distants ne sont pas couverts.

## Référencement de code et filtre de duplication

GitHub Copilot aide à comprendre la provenance du code proposé en référençant et en liant la source. Lorsque vous acceptez une suggestion de ce type, Copilot enregistre l'URL de la source et sa licence.

Vous pouvez ainsi décider d'utiliser ou non l'extrait et du type d'attribution à fournir.

### Fonctionnement du filtre de duplication

Lorsqu'une suggestion est générée, Copilot la compare à du code public connu. Si elle correspond à un dépôt public au-delà d'un seuil de similarité, elle est bloquée ou signalée avec attribution.

En cas d'acceptation, Copilot enregistre :

- La date et l'heure d'acceptation
- Le fichier où la suggestion a été ajoutée
- Un extrait du code ajouté
- La licence et l'URL de la source

Vous pouvez voir les références de code directement dans l'IDE lorsqu'une correspondance est signalée. Dans VS Code, elles apparaissent dans le panneau de sortie Copilot aux côtés de la suggestion.

### Ce que le filtre ne détecte pas

Le filtre de duplication ne fait pas correspondre :

- Les courts extraits et motifs trop génériques pour être signalés.
- Le code restructuré ou partiellement modifié par rapport à une source.

Son objectif est de repérer les correspondances textuelles (identiques ou quasi identiques), pas la similarité conceptuelle.

### Protection PI et garanties contractuelles

GitHub offre une protection en propriété intellectuelle (PI) aux clients Copilot Business et Enterprise. Si une suggestion entraîne une réclamation, GitHub prend en charge la défense juridique.

Deux conditions s'appliquent :

1. Le filtre de duplication doit être activé
2. Vous devez disposer d'une offre éligible

Les utilisateurs Free et Pro ne sont pas couverts. Il s'agit d'une garantie commerciale : elle n'empêche pas l'apparition de correspondances, mais couvre le risque juridique le cas échéant.

## Gérer les politiques Copilot sur GitHub.com

La page de politique Copilot sur GitHub.com est l'endroit où les admins org et Enterprise pilotent ce que Copilot peut faire pour l'équipe.

### Activer et désactiver des fonctionnalités Copilot

Les admins peuvent activer/désactiver indépendamment les complétions de code, Chat, la relecture de code, l'intégration GitHub CLI et le mode agent. Cette granularité est utile pour un déploiement progressif ou pour restreindre certaines capacités à des équipes spécifiques.

L'attribution des licences se fait également ici : vous choisissez qui a accès, par utilisateur ou par groupe.

### Configurer les modèles d'IA autorisés

Copilot prend en charge plusieurs modèles sous-jacents, et les administrateurs peuvent restreindre ceux autorisés pour l'organisation. Vous pouvez le verrouiller sur un modèle précis ou autoriser toutes les options disponibles et laisser le choix aux développeurs.

Pour les environnements publics et réglementés, GitHub propose des options conformes au programme FedRAMP (Federal Risk and Authorization Management Program).

Vérifiez les paramètres de politique Copilot dans l'onglet Copilot de votre organisation pour voir ce qui est disponible avec votre offre.

### Audit et conformité

Les indicateurs d'usage de Copilot, y compris les taux de complétion, le nombre d'utilisateurs actifs et les tendances par fonctionnalité et modèle, sont disponibles sous l'onglet Insights au niveau Enterprise et organisation (Insights > Copilot usage).

La politique « Copilot usage metrics » doit être activée pour accéder au tableau de bord. Des détails par membre sont disponibles via export NDJSON.

Les données de sièges et de licences sont distinctes et accessibles sous Org Settings > Copilot > Access.

Pour un tour d'horizon détaillé de la nouvelle Usage Metrics API et d'autres fonctionnalités avancées, n'hésitez pas à lire notre guide GitHub Copilot Enterprise.

Notez que le journal d'audit n'inclut pas les données de session client telles que les invites ; une solution personnalisée est nécessaire pour cela. Pour conserver l'historique au-delà de 180 jours ou mettre en place des alertes d'anomalie, GitHub recommande de diffuser le journal d'audit vers une plateforme SIEM via la fonctionnalité de streaming intégrée.

## Dépanner GitHub Copilot

Lorsque Copilot cesse de fonctionner, la cause est presque toujours l'une des suivantes. Vérifiez ces points avant d'ouvrir un ticket.

### Suggestions absentes ou interrompues

Commencez par l'icône d'état Copilot dans la barre d'état de l'IDE. Une diagonale sur l'icône indique qu'une exclusion de contenu est active pour le fichier courant.

Si l'icône paraît normale mais qu'aucune suggestion n'apparaît, contrôlez, dans l'ordre :

- Mettez à jour votre IDE et l'extension Copilot.
- Vérifiez que votre abonnement est actif et que votre compte dispose d'une licence attribuée.
- Contrôlez les règles d'exclusion pour le fichier et le repo, et testez votre connexion réseau.

Les configurations de proxy et de VPN bloquent fréquemment en silence : l'IDE doit joindre les serveurs Copilot sur l'infrastructure GitHub et certains proxys d'entreprise les bloquent sans message explicite.

### Exclusion de contenu inattendue

Après l'ajout ou la modification d'exclusions, un délai allant jusqu'à 30 minutes peut s'appliquer dans les IDE où les paramètres sont déjà chargés.

Pour appliquer les changements immédiatement :

- Dans VS Code, ouvrez la **Command Palette** et lancez**Developer: Reload Window** .
- Dans les IDE JetBrains et Visual Studio, fermez puis rouvrez l'application.
- Dans Vim/Neovim, aucune action : les exclusions sont récupérées automatiquement à chaque ouverture de fichier.

Après rechargement, testez explicitement : ouvrez le fichier exclu et demandez à Chat de l'expliquer. Si Chat décrit le contenu du fichier, l'exclusion ne s'applique pas ; revérifiez la syntaxe de la règle.

Notez que trois fonctionnalités Copilot ne gèrent pas les exclusions : Copilot CLI, l'agent de codage Copilot (agent cloud) et le mode Agent dans Copilot Chat en IDE. Si vous observez des accès inattendus avec ces modes, ce n'est pas une mauvaise configuration.

### Problèmes d'authentification et de jetons

Si Copilot est indisponible dans VS Code malgré une session ouverte, déconnectez-vous via l'icône Accounts en bas à gauche, rechargez la fenêtre (F1 > Developer: Reload Window dans VS Code), puis reconnectez-vous.

Dans Visual Studio, vérifiez que le compte GitHub connecté correspond à celui disposant d'une licence Copilot, rafraîchissez les informations d'identification si besoin ou supprimez/réajoutez le compte et redémarrez Visual Studio.

### Limitation de taux

Le modèle de facturation à l'usage de Copilot implique des capacités différentes par offre, et les modèles premium consomment cette capacité plus vite que les modèles de base.

Si les suggestions s'interrompent en cours de session ou si Copilot Chat renvoie des erreurs, passer en sélection de modèle automatique (ou vers un modèle à multiplicateur plus faible) peut résoudre le problème le temps que la fenêtre d'usage se réinitialise.

