---
id: collect-261001-cisco/cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026-5
title: "Génère un secret de session Graylog (obligatoire, 16+ caractères)"
domain: cisco
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["arr", "aws", "open source"]
source: docs/RAG/collect-261001-cisco/graylog-7-1-centraliser-ses-logs-en-11-etapes-2026.md
source_anchor: ""
source_lines: [322, 388]
sha256: d5e3e235a940f3459354b9a3af6775494aef19fffd98f5c267e8ecc7192fc81c
---

# Génère un secret de session Graylog (obligatoire, 16+ caractères)

- **Utiliser le tag Docker “latest”** — une mise à jour majeure non testée peut casser la compatibilité MongoDB/OpenSearch du jour au lendemain. Figez toujours un numéro de version précis comme 7.1.8.
- **Oublier le réglage vm.max_map_count** — OpenSearch refuse de démarrer sans cet ajustement noyau, une des causes d’échec les plus fréquentes en environnement conteneurisé.
- **Exposer le port 9000 directement sur Internet sans TLS ni pare-feu** — l’interface d’administration devient alors une cible directe pour du brute-force ou de l’exploitation de vulnérabilité.
- **Ne définir aucune politique de rétention d’index** — le disque se remplit progressivement jusqu’à un arrêt brutal du service par manque d’espace.
- **Confondre GELF et Syslog** — envoyer un flux Syslog vers un input configuré en GELF ne génère aucune erreur visible, mais aucun message n’apparaît jamais dans la recherche.
- **Ignorer les logs du conteneur Graylog lors d’un problème** — la commande docker compose logs -f graylog révèle presque toujours la cause exacte d’un échec de démarrage.
- **Sous-dimensionner le heap OpenSearch par précaution budgétaire** — allouer 512 Mo à un moteur censé indexer plusieurs gigaoctets par jour provoque des ralentissements puis des pertes de messages silencieuses sous forte charge.
- **Négocier les identifiants d’accès en clair dans le fichier docker-compose.yml versionné sur Git** — utilisez systématiquement un fichier .env exclu du dépôt via .gitignore, comme dans l’exemple de ce tutoriel.

## Dépannage : problèmes courants et solutions

**Le conteneur OpenSearch s’arrête immédiatement après le démarrage.** Vérifiez que `vm.max_map_count` vaut bien 262144 avec `sysctl vm.max_map_count`, et que la limite `memlock` est illimitée dans le fichier compose.

**Graylog reste bloqué sur “Waiting for indexer to become available”.** Le serveur Graylog a démarré trop tôt. Ajoutez un délai ou un mécanisme de retry, ou redémarrez simplement le conteneur Graylog après confirmation qu’OpenSearch répond sur le port 9200.

**Impossible de se connecter avec le mot de passe admin.** Le hash SHA-256 doit correspondre exactement au mot de passe en clair, sans retour à la ligne parasite. Régénérez-le avec `echo -n` (sans le “-n”, un caractère invisible s’ajoute et casse le hash).

**Les messages Syslog n’apparaissent jamais dans la recherche.** Vérifiez avec `tcpdump -i any port 1514` sur le serveur Graylog que les paquets arrivent bien, puis contrôlez que l’input Syslog est en statut “Running” et non “Failed”.

**L’interface web répond très lentement après quelques semaines d’utilisation.** C’est généralement un signe que la mémoire allouée à OpenSearch (`OPENSEARCH_JAVA_OPTS`) est sous-dimensionnée par rapport au volume de logs ingéré. Augmentez le heap à 4 Go si le serveur dispose de 16 Go de RAM ou plus.

**Erreur “too many open files” dans les logs OpenSearch.** La limite de fichiers ouverts du système n’a pas été appliquée. Vérifiez `/etc/security/limits.conf` et redémarrez la session ou le serveur pour que la limite prenne effet.

**Les pipelines ne modifient jamais les messages malgré une règle active.** La règle doit être explicitement rattachée à un pipeline, et ce pipeline doit être connecté au flux (“Stream”) concerné dans “System > Pipelines > Manage rules”. Une règle non connectée à un pipeline ne s’exécute jamais.

**Le volume de stockage Docker grossit anormalement vite.** Vérifiez la politique de rotation d’index dans “System > Indices” : un index configuré sans limite de taille ou de nombre continue de croître indéfiniment.

**Les alertes par e-mail ne partent jamais.** La configuration SMTP se règle dans le fichier de configuration serveur ou via les variables d’environnement `GRAYLOG_TRANSPORT_EMAIL_ENABLED` et associées, absentes par défaut du fichier Compose de ce tutoriel. Ajoutez-les explicitement si vous comptez notifier par e-mail plutôt que par webhook.

## Astuces avancées pour aller plus loin

Une fois la base opérationnelle, plusieurs optimisations méritent d’être explorées. Le Graylog Sidecar, en version 1.5.4, permet de déployer et gérer à distance des collecteurs (Filebeat, Winlogbeat) sur des dizaines de machines depuis l’interface Graylog centrale, sans se connecter en SSH sur chaque poste individuellement. Pour les architectures multi-sites, le Graylog Forwarder, passé en version 7.6 le 4 septembre 2026, relaie désormais les logs collectés sur des sites distants ou des environnements isolés vers le cluster central sans exposer directement OpenSearch à Internet. Les content packs Illuminate 7.0.7 fournissent des règles de détection prêtes à l’emploi pour des sources courantes (pare-feu Palo Alto, Microsoft 365, AWS CloudTrail), un gain de temps considérable comparé à l’écriture manuelle de chaque pipeline.

Pour les environnements qui dépassent un seul serveur, Graylog supporte un déploiement en cluster avec plusieurs nœuds de traitement partageant le même cluster OpenSearch, ce qui répartit la charge d’ingestion et élimine le point de défaillance unique. Enfin, pour les organisations qui doivent démontrer leur conformité NIS2 ou répondre à un audit ANSSI, l’export régulier des dashboards et des règles d’alerte vers un dépôt de configuration versionné (Git) constitue une preuve documentée de la posture de détection en place, un exercice que la directive NIS2 demande explicitement aux entités essentielles et importantes.

## Foire aux questions

**Graylog est-il vraiment gratuit ?**

L’édition Graylog Open est open source et gratuite pour un usage illimité. Les fonctions avancées (data tiering étendu, contenus de sécurité premium, support entreprise) relèvent des éditions commerciales Graylog Enterprise et Graylog Security, dont le tarif se négocie directement avec l’éditeur.

**Peut-on installer Graylog sans Docker ?**

Oui, une installation native sur Debian ou Ubuntu est possible via les paquets APT officiels, mais elle demande de gérer manuellement les versions Java, MongoDB et OpenSearch, ce qui multiplie les risques d’incompatibilité. Docker Compose reste la méthode la plus fiable pour une première installation.

**Combien de RAM faut-il vraiment pour un usage en production ?**

8 Go suffisent pour un lab ou un test avec quelques sources. Au-delà de 5 Go de logs ingérés par jour, comptez 16 Go minimum, majoritairement consommés par le heap OpenSearch.

**Graylog remplace-t-il un antivirus ou un pare-feu ?**

Non. Graylog centralise et analyse des logs déjà générés par d’autres équipements ; il ne bloque rien par lui-même. Il se combine avec des outils actifs comme un pare-feu ou une solution de blocage comportemental pour une protection complète.

**Quelle est la différence entre Graylog Open et Graylog Security ?**

Graylog Open couvre la centralisation, la recherche et les pipelines de base. Graylog Security ajoute les investigations automatisées, la détection comportementale avancée et les content packs de sécurité premium sous licence commerciale.

**Faut-il migrer d’Elasticsearch vers OpenSearch avant d’installer Graylog 7.1 ?**

Oui si vous partez d’une ancienne installation. Graylog a définitivement abandonné le support d’Elasticsearch au profit d’OpenSearch depuis la version 5.0 ; toute nouvelle installation en 2026 doit directement utiliser OpenSearch 2.19.x.

**Combien de temps faut-il pour déployer cette architecture complète ?**

Comptez environ 90 minutes pour suivre les onze étapes de ce tutoriel sur un serveur neuf, en incluant les tests de chaque input et la création d’une première alerte fonctionnelle.

**Graylog fonctionne-t-il pour surveiller des logs Windows ?**

