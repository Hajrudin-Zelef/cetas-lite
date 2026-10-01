---
id: collect-261001-general-networking/general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026-5
title: "Redémarre l'agent et relance les scans SCA"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "arr", "agents", "mai", "open source"]
source: docs/RAG/collect-261001-general-networking/wazuh-siem-open-source-en-12-etapes-40-min-2026.md
source_anchor: ""
source_lines: [294, 359]
sha256: d7ef1df916c6ef67e6b276077a90129575622c8120577e791c6f66ba6a935c68
---

# Redémarre l'agent et relance les scans SCA

- **Sous-dimensionner la RAM.** En dessous de 8 Gio, l’indexeur OpenSearch se fait tuer par le noyau (OOM) et le tableau de bord devient inaccessible. C’est l’erreur numéro un.
- **Négliger la synchronisation NTP.** Un décalage d’horloge entre agent et manager fausse l’horodatage, désynchronise les flux et fait apparaître des agents « déconnectés » à tort.
- **Exposer l’interface à Internet.** Publier les ports 443, 55000 ou 9200 sans filtrage transforme votre SIEM en cible. Passez par un VPN.
- **Activer le FIM temps réel partout.** Surveiller`/var` ou`/` en`realtime` sature le CPU. Ciblez uniquement les répertoires sensibles.
- **Oublier la rétention des index.** Sans politique de cycle de vie, l’indexeur remplit le disque en quelques semaines et bloque l’ingestion.
- **Mélanger les versions.** Un agent dont la version est supérieure à celle du manager n’est pas garanti. Mettez toujours le manager à niveau en premier.

## Dépannage : 10 erreurs fréquentes et leurs solutions

Un problème lors de l’installation ou de l’enrôlement ? Ce tableau couvre les symptômes les plus courants rencontrés avec Wazuh et leur résolution.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Agent en « Never connected » | Ports 1514/1515 filtrés ou mauvaise IP manager | Ouvrir le pare-feu, vérifier `WAZUH_MANAGER` , redémarrer l’agent | 
| Agent « Disconnected » par intermittence | Décalage d’horloge (NTP) | Synchroniser l’heure sur agent et manager | 
| Installateur : « not enough RAM » | Moins de 8 Gio de mémoire | Augmenter la RAM à 8 Gio minimum | 
| « Failed to connect to indexer » (9200) | Service `wazuh-indexer` arrêté ou heap JVM trop petit | Redémarrer l’indexeur, ajuster `jvm.options` | 
| Tableau de bord inaccessible (443) | Service `wazuh-dashboard` down ou certificat invalide | Vérifier le service, régénérer les certificats | 
| Mot de passe admin perdu | Archive d’installation non conservée | Utiliser `wazuh-passwords-tool.sh` | 
| Agent connecté mais aucune alerte | Bloc `localfile` ou`syscheck` absent | Vérifier `ossec.conf` , redémarrer l’agent | 
| Indexeur qui consomme toute la RAM | Heap JVM mal réglé | Fixer le heap à 50 % de la RAM (max 32 Go) | 
| Disque plein en quelques semaines | Aucune politique de rétention d’index | Configurer l’Index State Management (ISM) | 
| Réponse active ne bannit pas l’IP | Règle non déclenchée ou droits iptables manquants | Vérifier `rules_id` et`active-responses.log` | 

En cas de doute, deux fichiers sont vos meilleurs alliés : `/var/ossec/logs/ossec.log` côté manager et agent, et `/var/ossec/logs/active-responses.log` pour la réponse active. La commande `sudo /var/ossec/bin/wazuh-control status` liste l’état de tous les démons.

## Astuces avancées pour aller plus loin avec Wazuh

Une fois les bases maîtrisées, Wazuh révèle une profondeur remarquable. Voici les pistes qui font la différence en environnement de production.

- **Cluster haute disponibilité.** Déployez plusieurs managers en mode cluster (un nœud*master* , plusieurs*workers* ) et répartissez la charge des agents pour tolérer les pannes et absorber les gros volumes.
- **Règles et décodeurs personnalisés.** Créez vos propres règles dans`/var/ossec/etc/rules/local_rules.xml` pour couvrir vos applications métier, et des décodeurs dans`local_decoder.xml` pour parser des formats de log exotiques.
- **Intégrations SOAR et alerting.** Le module*integrator* pousse les alertes vers Slack, un e-mail, TheHive ou Shuffle. Automatisez ainsi la création de tickets et l’orchestration de la réponse.
- **Ingestion des logs réseau.** Envoyez en Syslog (port 514) les journaux de votre pare-feu pfSense ou OPNsense vers Wazuh pour corréler événements réseau et endpoints.
- **API RESTful.** Pilotez Wazuh par programmation via l’API (port 55000) et un jeton JWT. Idéal pour l’automatisation et l’intégration CI/CD.

Exemple d’authentification à l’API et de récupération de la liste des agents actifs :

```
TOKEN=$(curl -u wazuh-wui:VOTRE_MOT_DE_PASSE -k -X POST \
  "https://10.0.0.10:55000/security/user/authenticate?raw=true")
curl -k -X GET "https://10.0.0.10:55000/agents?pretty=true&status=active" \
  -H "Authorization: Bearer $TOKEN"
```
Enfin, gardez votre déploiement à jour. Wazuh publie des correctifs réguliers sur la branche 4.14.x : la 4.14.3 (11 février 2026), la 4.14.4 (17 mars 2026), la 4.14.5 (23 avril 2026) puis la 4.14.6 (1er juillet 2026) – qui a par exemple retiré des options de transport SSL/TLS obsolètes du cluster et mis à jour ses bibliothèques eBPF, avec libbpf en version 1.7.0 et bpftool en version 7.7.0 – ont chacune apporté leur lot de correctifs cumulatifs, confirmant un rythme d’au moins trois publications sur le seul premier semestre 2026, jusqu’à la 4.14.7 du 29 juillet 2026, aujourd’hui la version stable recommandée. Suivez les notes de version officielles et planifiez vos montées de version, en commençant toujours par le manager.

## Foire aux questions sur Wazuh

### Wazuh est-il vraiment gratuit ?

Oui. Tous les composants de Wazuh (manager, indexeur, tableau de bord, agents) sont open source sous licence GPLv2 et totalement gratuits, sans limite d’agents ni de volume de données. Wazuh, Inc. propose par ailleurs une offre cloud gérée payante, optionnelle, pour ceux qui ne veulent pas auto-héberger : d’après le comparateur MSP Compared, son tarif d’entrée démarrait à 571 $/mois pour 100 agents en juin 2026 – un coût que la version auto-hébergée décrite dans ce tutoriel permet d’éviter entièrement.

### Quelle est la dernière version de Wazuh en 2026 ?

La dernière version stable est Wazuh 4.14.7, publiée le 29 juillet 2026. La branche 4.14 a introduit courant 2025-2026 la fonctionnalité IT Hygiene, un tableau de bord Microsoft Graph et le support des politiques SCA pour Windows Server 2025 et macOS 26 ; les utilisateurs de la branche de maintenance 4.10.x reçoivent eux aussi des correctifs réguliers, comme la version 4.10.4 du 21 mai 2026. Une version 5.0 est en développement : sa branche a été fusionnée dans la branche principale le 29 juin 2026, avec un premier bump de version vers l’agent 5.0.1 peu après ; selon un suivi de Tech Insider, elle restait toutefois encore au stade bêta à cette date de juin 2026 (après la préversion 5.0.0-beta2 de mai 2026), une étape clé mais pas encore définitive vers cette prochaine génération de Wazuh.

### Quelle différence entre Wazuh et Splunk ?

Splunk est une plateforme commerciale facturée à l’ingestion de données, très puissante mais coûteuse à grande échelle. Wazuh est open source et gratuit, avec des capacités SIEM et XDR intégrées (FIM, SCA, détection de vulnérabilités, réponse active). Pour un budget contraint ou une exigence de souveraineté, Wazuh est souvent le meilleur rapport valeur/coût.

### Combien d’agents un seul serveur Wazuh peut-il gérer ?

Un déploiement tout-en-un correctement dimensionné (8 vCPU, 8 Gio de RAM, 200 Go de disque) supervise confortablement jusqu’à 100 agents avec 90 jours de rétention. Au-delà, on répartit l’indexeur et le tableau de bord sur des machines dédiées, et l’on passe le manager en cluster pour atteindre plusieurs milliers d’agents – à titre de comparaison, la formule cloud gérée équivalente de l’éditeur facturait, en juin 2026 selon MSP Compared, 923 $/mois pour 250 agents et jusqu’à 1 467 $/mois pour 500 agents, un argument de poids en faveur du cluster auto-hébergé pour les parcs de grande taille.

### Wazuh fonctionne-t-il sur Windows ?

