---
id: collect-261001-general-networking/general-networking/fr-review-dell-powerprotect-one-9dec1a23-2
title: "fr-review-dell-powerprotect-one-9dec1a23"
domain: general-networking
role: reference
task: reference
actors: ["Google", "Microsoft", "Oracle"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/fr-review-dell-powerprotect-one-9dec1a23.md
source_anchor: ""
source_lines: [13, 39]
sha256: 46a6a7210a5af7d8327f29584c6ec20d25f13d8c59c9e8252c790f07a810a86b
---

# fr-review-dell-powerprotect-one-9dec1a23

Le contrôle d'accès s'étend à plusieurs méthodes d'authentification. L'intégration de l'authentification unique (SSO) inclut Okta, Microsoft Entra ID, PingOne et RSA SecurID. Parallèlement, l'authentification multifacteur est prise en charge via des fournisseurs TOTP tels que Google Authenticator, Microsoft Authenticator, Authy, Duo et LastPass. Le contrôle d'accès basé sur les rôles limite les actions d'administration aux utilisateurs autorisés, et la journalisation des audits enregistre l'activité sur la plateforme à des fins de vérification de conformité et d'analyse forensique.
Le verrouillage de la période de rétention est le mécanisme d'immuabilité le plus directement lié à la protection contre les ransomwares. Une fois appliqué, les données de sauvegarde ne peuvent être ni supprimées ni modifiées avant l'expiration de cette période, même par des administrateurs disposant de privilèges élevés. La plateforme prend en charge les modes de gouvernance et de conformité, permettant aux organisations de se conformer aux exigences réglementaires qui imposent des contrôles d'immuabilité plus stricts. Combiné aux flux de travail de détection d'anomalies qui signalent les modifications suspectes dans les données de sauvegarde, le verrouillage de la période de rétention offre à PowerProtect One les bases nécessaires pour une restauration propre après une attaque de ransomware, sans dépendre de l'environnement de production compromis.
Gestion PowerProtect One
Jour 1 : Déploiement
Le déploiement initial du dispositif PowerProtect One est rapide et guidé. Lors de nos tests, le système était opérationnel en moins de 10 minutes après l'application de la configuration. La configuration initiale privilégie la mise en ligne des services essentiels plutôt qu'un long processus de configuration.
La configuration initiale comprend la configuration réseau de l'hôte, les paramètres horaires et l'accès iDRAC au matériel sous-jacent. Dell fournit un compte iDRAC en lecture seule, que nous avons trouvé très utile. Il permet de surveiller l'état du matériel et de recevoir des alertes sans risque de modification ou d'arrêt accidentel, un aspect crucial dans les environnements gérant un parc de systèmes Dell.
Lors de la première connexion, la plateforme propose à l'utilisateur un processus de démarrage. Ce dernier détaille les notifications par e-mail, l'assistance automatique, les paramètres de sécurité et la gestion des licences. Bien que ces éléments puissent être configurés ultérieurement, leur paramétrage initial permet d'éviter des erreurs importantes et de rendre rapidement l'appliance opérationnelle.
Jour 2 : Opérations
Tableau de bord unifié
Les opérations quotidiennes s'articulent autour de ce que Dell appelle un tableau de bord unifié. À partir de là, PowerProtect One peut gérer de nombreux systèmes enregistrés, offrant aux administrateurs une visibilité et un contrôle complets sur tous les systèmes connectés de l'environnement.
En pratique, le tableau de bord affiche les informations système les plus importantes que vous consulterez probablement en premier. L'activité des tâches est clairement indiquée pour tous les systèmes, y compris les sauvegardes en cours, terminées et ayant échoué. L'état du système est ventilé par services, protection, stockage et sécurité, ce qui facilite le repérage des problèmes sans avoir à naviguer dans plusieurs menus.
La capacité est également facile à suivre, grâce à l'affichage clair de l'utilisation des niveaux actifs et de l'espace disponible. Le tableau de bord présente également le total des actifs protégés, les anomalies récentes et l'efficacité de la réduction des données, offrant ainsi un contexte précieux sur le comportement quotidien des systèmes au sein de l'environnement.
La navigation utilise une arborescence imbriquée sur la gauche. À l'usage, cela garantit une navigation intuitive. Lors de l'exploration d'un système ou d'une alerte spécifique, quelques clics suffisent pour passer d'une vue à l'autre sans perdre sa position.
Création d'unités de stockage
Le stockage dans PowerProtect One s'articule autour d'unités de stockage, qui servent de conteneurs principaux pour les données de sauvegarde selon différentes politiques. Les unités de stockage sont créées directement depuis l'onglet Infrastructure et engendrent peu de surcharge.
Chaque unité de stockage peut être configurée avec des quotas souples et stricts, ainsi que des limites de flux, afin de contrôler le débit. Le verrouillage de la rétention est disponible et simple à appliquer, permettant ainsi de garantir l'intégrité des sauvegardes pendant une période définie. Lors de nos tests, cette fonctionnalité s'est avérée facile à activer et n'a pas complexifié le flux de travail.
Des optimisations spécifiques à la charge de travail sont également disponibles, notamment pour les environnements Oracle. Les unités de stockage peuvent servir aux sauvegardes internes ou être exposées en externe via des intégrations telles que DD Boost. PowerProtect One prend en charge les niveaux de stockage actif et cloud, permettant ainsi de stocker les données localement ou de les étendre au stockage cloud selon les besoins.
Création de politiques
La création des politiques s'effectue dans l'onglet Protection et suit un processus simple. Une fois les ressources ajoutées au système, elles peuvent être associées à une politique et se voir attribuer un objectif de sauvegarde défini.
Par défaut, la création d'une stratégie crée également une nouvelle unité de stockage, mais il est possible de sélectionner des unités existantes si nécessaire. Le verrouillage de la rétention peut être appliqué à cette étape, garantissant ainsi que les sauvegardes restent inchangées jusqu'à l'expiration de la période de rétention.
La fréquence et le type de sauvegarde sont configurables, notamment les sauvegardes complètes synthétiques et les fenêtres d'exécution définies. Les durées de conservation peuvent être ajustées selon les besoins. Des options supplémentaires telles que la réplication, le stockage sécurisé, la hiérarchisation vers le cloud et l'archivage sont disponibles dans la même configuration de stratégie.
Une fois une politique créée, sa progression peut être suivie via la vue des tâches. En pratique, cela permet de vérifier facilement que les politiques s'exécutent comme prévu sans avoir à naviguer entre plusieurs sections de l'interface.
Planification intelligente
La planification est flexible sans ajouter de complexité inutile. Les administrateurs peuvent définir le moment d'exécution des sauvegardes, la fréquence des sauvegardes complètes synthétiques et la durée de conservation des données, le tout grâce à un flux de travail simple et facile à ajuster.
Les fenêtres d'exécution permettent d'éviter les conflits avec les charges de travail de production, tandis que les paramètres d'optimisation permettent aux équipes de privilégier les performances ou la capacité selon le cas d'utilisation. Cette flexibilité contribue également au respect des SLA, qui imposent l'exécution des tâches de sauvegarde dans des délais précis. En ajustant les planifications et l'utilisation des ressources, les administrateurs peuvent mieux aligner les opérations de sauvegarde sur les objectifs de restauration définis et les attentes de l'entreprise.
Globalement, les commandes offrent un bon équilibre entre flexibilité et simplicité, ce qui facilite l'adaptation des horaires sans engendrer de frais généraux inutiles.
Détection d’Anomalies
