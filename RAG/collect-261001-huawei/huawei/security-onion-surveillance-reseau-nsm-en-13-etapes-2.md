---
id: collect-261001-huawei/huawei/security-onion-surveillance-reseau-nsm-en-13-etapes-2
title: "Calculer le hash SHA256 de l'ISO téléchargée"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/security-onion-surveillance-reseau-nsm-en-13-etapes.md
source_anchor: ""
source_lines: [61, 123]
sha256: fb2c41164e67ab25f803eba4814f98244055249ef4ed5b2be19a00d7c3d95124
---

# Calculer le hash SHA256 de l'ISO téléchargée

Bootez la VM sur l’ISO. L’écran d’accueil de Security Onion propose deux chemins : une installation graphique guidée (recommandée pour une première prise en main) ou une installation sur une base Rocky Linux 9 Minimal déjà en place, pour les administrateurs qui préfèrent maîtriser chaque paquet du système de base.

Pour ce tutoriel, choisissez l’installation via ISO complète. L’assistant vous demande successivement : le fuseau horaire (Europe/Paris), la configuration réseau de l’interface management (IP statique fortement recommandée en production), le nom d’hôte, et un compte administrateur local avec mot de passe robuste. Ce compte servira à la fois pour l’accès SSH et pour la première connexion à la console web.

## Étape 4 : choisir le type d’installation Security Onion

Une fois le système de base Rocky Linux installé, le script setup de Security Onion démarre automatiquement. C’est ici que se joue la décision la plus importante du déploiement : le type de nœud.

| Type de nœud | Rôle | Cas d’usage typique | 
|---|---|---|
| Standalone | Manager + capteur sur une seule machine | Lab, PME, test de faisabilité, petit réseau | 
| Manager | Gère la configuration, la console SOC et l’orchestration du cluster | Déploiement distribué multi-sites | 
| Search node | Nœud dédié à l’indexation et la recherche Elasticsearch | Scaling horizontal quand le volume de logs augmente | 
| Sensor | Capture le trafic, exécute Suricata et Zeek | Chaque site distant à surveiller | 

Pour une première installation ou un environnement de moins de 200 postes, le mode Standalone est le bon choix : il regroupe tous les rôles sur une seule machine et simplifie considérablement la maintenance. Les déploiements distribués (Manager séparé des Sensors) ne se justifient qu’à partir d’une volumétrie de trafic ou d’un nombre de sites qui dépasse la capacité d’un seul serveur, typiquement au-delà de plusieurs centaines de mégabits par seconde de trafic soutenu.

## Étape 5 : configurer l’interface de capture réseau

L’installateur vous demande ensuite quelle interface réseau utiliser pour le management (accès web, SSH) et laquelle utiliser pour la capture (monitoring). C’est une confusion fréquente : inverser les deux rend la console web inaccessible et empêche toute capture de trafic utile.

```
# Vérifier après installation quelles interfaces sont détectées
ip link show
# Vérifier que l'interface de monitoring est bien en mode promiscuous
ip -d link show eth1 | grep -i promisc
# Consulter le statut des services de capture
sudo so-status
```
Le script so-status est votre premier réflexe après chaque redémarrage ou modification de configuration : il liste l’état (running, stopped, degraded) de chaque conteneur Docker composant la stack (Suricata, Zeek, Wazuh, Elasticsearch, Logstash, Kibana, SOC). Une installation saine affiche tous les services en statut “running” au bout de 10 à 15 minutes après le premier démarrage, le temps que les images Docker soient extraites et que les index Elasticsearch soient initialisés.

## Étape 6 : finaliser l’installation et le premier démarrage

Une fois toutes les questions répondues, l’installateur télécharge et configure automatiquement l’ensemble des composants : Suricata pour la détection d’intrusion signature-based, Zeek pour l’analyse de protocoles et la génération de logs réseau enrichis, Wazuh pour la collecte de logs système et la détection sur les hôtes, puis la suite Elastic (Elasticsearch, Logstash, Kibana) pour l’indexation et la visualisation. Cette phase dure généralement entre 20 et 45 minutes selon la vitesse de votre connexion Internet et la puissance du CPU alloué.

À la fin du processus, le serveur redémarre et affiche une invite de connexion classique. Notez l’adresse IP de l’interface management affichée à l’écran : c’est par cette adresse, en HTTPS, que vous accéderez à la console SOC.

## Étape 7 : première connexion à la console Security Onion (SOC)

Ouvrez un navigateur et rendez-vous sur https://IP-DE-VOTRE-SERVEUR. Le certificat TLS généré par défaut est auto-signé, votre navigateur affichera donc un avertissement de sécurité au premier accès : c’est normal, acceptez l’exception ou importez le certificat racine généré par l’installateur si vous souhaitez éviter l’avertissement sur tous vos postes d’administration.

Connectez-vous avec le compte administrateur créé pendant l’installation. La console SOC s’ouvre sur un tableau de bord (Dashboards) qui agrège les métriques clés : nombre d’alertes par sévérité, volume de trafic capturé, top des adresses IP sources et destinations, statut des capteurs. C’est le point d’entrée unique pour toute l’exploitation quotidienne, vous n’avez normalement plus besoin d’accéder séparément à Kibana ou aux interfaces natives de chaque outil.

## Étape 8 : explorer les alertes et lancer une première chasse aux menaces

Le menu Alerts regroupe toutes les détections remontées par Suricata (signatures réseau) et Wazuh (règles hôtes). Chaque alerte affiche sa sévérité, la règle qui l’a déclenchée, les adresses IP et ports concernés, ainsi qu’un lien direct vers le paquet réseau brut capturé au moment de l’événement grâce à l’intégration avec l’outil de capture complète (full packet capture). Un clic sur une alerte ouvre une vue détaillée qui affiche non seulement le paquet incriminé, mais aussi l’ensemble de la session réseau qui l’entoure, ce qui permet de reconstituer le contexte complet d’une tentative d’intrusion sans quitter la console.

Une bonne pratique pour les premières semaines d’exploitation consiste à trier les alertes par sévérité décroissante et à traiter en priorité les événements classés critiques et élevés, plutôt que de tenter de traiter l’intégralité du flux dès le premier jour. Le volume d’alertes générées par une installation fraîchement configurée peut atteindre plusieurs centaines par jour sur un réseau d’entreprise de taille moyenne, avant même le travail de calibration des règles évoqué plus loin dans ce guide.

Le module Hunt permet d’interroger manuellement l’ensemble des logs indexés, qu’ils proviennent de Zeek, de Suricata ou de Wazuh, avec une syntaxe de requête proche de celle de Kibana. C’est l’outil central pour la threat hunting proactive : rechercher une adresse IP suspecte sur les dernières 24 heures, isoler tous les flux DNS vers un domaine donné, ou croiser des logs d’authentification avec des connexions réseau sortantes anormales.

```
# Exemple de requête dans le module Hunt pour isoler
# tout le trafic DNS sortant vers un domaine suspect
event.dataset:dns AND dns.query.name:"domaine-suspect.example"
# Rechercher toutes les alertes de sévérité critique
# sur les dernières 24 heures
event.severity_label:"critical" AND @timestamp:[now-24h TO now]
```
## Comprendre le rôle exact de chaque composant intégré

Avant d’aller plus loin dans la configuration, il est utile de comprendre précisément ce que fait chaque brique de la stack, car les alertes que vous verrez dans la console SOC proviennent de moteurs très différents dans leur logique de fonctionnement.

