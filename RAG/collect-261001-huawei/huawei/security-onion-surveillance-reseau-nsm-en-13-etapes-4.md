---
id: collect-261001-huawei/huawei/security-onion-surveillance-reseau-nsm-en-13-etapes-4
title: "Calculer le hash SHA256 de l'ISO téléchargée"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "incident", "open source"]
source: docs/RAG/collect-261001-huawei/security-onion-surveillance-reseau-nsm-en-13-etapes.md
source_anchor: ""
source_lines: [197, 253]
sha256: 7cc48b88d24fa623a86227244f07c5a86ec376a7e6ca311b234ee469f2501bea
---

# Calculer le hash SHA256 de l'ISO téléchargée

L’argument central de Security Onion est le temps d’intégration : monter individuellement Suricata, Zeek, Wazuh puis les faire dialoguer avec une stack Elastic demande généralement plusieurs semaines de travail à un ingénieur sécurité expérimenté. Security Onion réduit cet effort à quelques heures, au prix d’une flexibilité un peu moindre sur la personnalisation fine de chaque composant pris isolément. Pour une comparaison plus large des solutions SIEM commerciales, consultez notre analyse Sentinel vs Splunk vs Elastic Security.

Le second argument, moins évident au premier abord, concerne la portabilité des compétences. Un analyste formé sur Security Onion maîtrise en réalité Suricata, Zeek, Wazuh et les bases d’Elasticsearch de façon transférable : ces compétences restent valables même si l’organisation bascule un jour vers un SIEM commercial qui embarque tout ou partie de ces mêmes moteurs open source en coulisses, ce qui est le cas d’un nombre croissant de solutions propriétaires construites elles-mêmes sur des fondations Elastic ou Suricata sous licence.

## 5 pièges courants à éviter

- **Sous-dimensionner la RAM.** Installer avec moins de 24 Go provoque des crashs répétés d’Elasticsearch au bout de quelques heures d’utilisation, une fois que le volume de logs commence à remplir les index en mémoire.
- **Inverser les interfaces management et monitoring.** C’est l’erreur numéro un des débutants : la console web devient inaccessible et le capteur ne voit passer aucun trafic utile.
- **Oublier le mode promiscuous sur le commutateur virtuel.** Sous VMware ou Proxmox, si le port group de l’interface de monitoring n’autorise pas le mode promiscuous, Security Onion ne recevra jamais la copie du trafic réseau, même avec un port SPAN correctement configuré côté switch physique.
- **Négliger la calibration des règles Suricata après l’installation.** Les signatures par défaut génèrent souvent un volume élevé de faux positifs les premières semaines ; sans ajustement via le module Playbook, les analystes se noient rapidement sous les alertes non pertinentes.
- **Exposer la console SOC directement sur Internet.** Cette interface d’administration centralise l’accès à l’ensemble des données de sécurité collectées ; elle doit systématiquement rester derrière un VPN ou un pare-feu restrictif, jamais accessible publiquement.

## Exemple de sortie attendue après une installation réussie

Une fois l’installation terminée et tous les services démarrés, la commande so-status doit afficher un résultat similaire à celui-ci, confirmant que l’ensemble de la stack fonctionne correctement :

```
$ sudo so-status
so-elasticsearch    running (healthy)
so-logstash         running (healthy)
so-kibana           running (healthy)
so-soc              running (healthy)
so-suricata         running (healthy)
so-zeek              running (healthy)
so-wazuh-manager     running (healthy)
so-redis             running (healthy)
Statut global : OPERATIONNEL - 8/8 services actifs
```
Si un ou plusieurs services affichent “degraded” ou “stopped” au lieu de “running (healthy)”, ne poursuivez pas la configuration avant d’avoir résolu le problème : chaque service dépend souvent des autres, et une base instable génère des comportements erratiques difficiles à diagnostiquer plus tard.

## Dépannage : 8 problèmes fréquents et leurs solutions

- **La console SOC affiche une page blanche après connexion.** Videz le cache du navigateur et vérifiez que le service so-soc est bien “running” via so-status ; un redémarrage du service suffit dans la majorité des cas.
- **Aucune alerte n’apparaît après plusieurs heures.** Vérifiez que l’interface de monitoring reçoit effectivement du trafic avec sudo tcpdump -i eth1 -c 20 ; si aucun paquet n’apparaît, le problème vient du port SPAN ou du mode promiscuous, pas de Security Onion lui-même.
- **Elasticsearch refuse de démarrer avec une erreur de mémoire virtuelle.** Le paramètre noyau vm.max_map_count doit être ajusté à au moins 262144 ; l’installateur le configure normalement automatiquement, mais une réinstallation ou une migration peut le réinitialiser.
- **La mise à jour soup échoue avec une erreur de dépôt.** Vérifiez la connectivité Internet du serveur et la résolution DNS ; les dépôts Security Onion doivent être joignables en HTTPS sortant sans proxy transparent qui intercepte le TLS.
- **Le disque se remplit rapidement malgré une politique de rétention configurée.** Vérifiez le volume réel de trafic capturé avec so-elastic-index-size ; un trafic plus élevé que prévu consomme l’espace plus vite que la politique de rétention par âge ne peut le libérer.
- **Les signatures Suricata ne se mettent pas à jour automatiquement.** Contrôlez le statut du service de mise à jour des règles dans le module Playbook et vérifiez que le serveur a bien accès aux flux de règles externes configurés.
- **Un nœud sensor distant n’apparaît pas dans le manager.** Vérifiez que le port de communication entre sensor et manager (généralement le port de registration Salt) n’est pas bloqué par un pare-feu intermédiaire entre les deux sites.
- **La console devient lente après plusieurs semaines d’utilisation.** C’est généralement un signe que les index Elasticsearch ont atteint une taille critique par rapport à la RAM allouée ; ajoutez de la mémoire ou réduisez la durée de rétention des données.

## Astuces avancées pour aller plus loin

Une fois la plateforme stabilisée, plusieurs axes permettent d’en tirer davantage de valeur opérationnelle. D’abord, intégrez des flux de renseignement sur les menaces (threat intelligence) externes pour enrichir automatiquement les alertes avec un contexte de réputation IP ou de domaine, via les intégrations disponibles dans le module Playbook.

Ensuite, configurez des règles de corrélation personnalisées qui combinent des événements Wazuh (comportement anormal sur un poste) avec des événements réseau Zeek ou Suricata, pour détecter des scénarios d’attaque en plusieurs étapes que chaque outil pris isolément ne verrait pas. C’est l’un des véritables avantages d’une stack intégrée par rapport à des outils cloisonnés.

Enfin, pour les organisations soumises à des obligations réglementaires, exportez régulièrement les journaux d’audit et les statistiques d’alertes vers un stockage froid externe, séparé du serveur Security Onion lui-même. En cas d’incident majeur touchant le serveur de surveillance, ces exports garantissent que les preuves collectées avant l’incident restent exploitables pour l’investigation post-mortem, un point souvent négligé mais essentiel pour la conformité NIS2 et pour toute procédure judiciaire ultérieure.

## Ressources et documentation officielle

Pour approfondir chaque composant intégré dans Security Onion, plusieurs sources font autorité. La documentation officielle Security Onion détaille l’ensemble des scénarios de déploiement, y compris les architectures distribuées à grande échelle. Le dépôt GitHub du projet reste la référence pour le téléchargement de l’ISO et le suivi des notes de version.

Pour les moteurs de détection sous-jacents, le site officiel de Suricata publie régulièrement les mises à jour de signatures, tandis que le projet Zeek documente en détail son langage de scripting pour créer des détections personnalisées. Côté supervision des hôtes, la documentation Wazuh couvre la configuration avancée des agents. Enfin, pour comprendre le fonctionnement interne de la couche d’indexation, la documentation officielle de la suite Elastic reste incontournable pour optimiser les performances de recherche à grande échelle.

## Cas d’usage concrets : PME, collectivités et établissements scolaires

