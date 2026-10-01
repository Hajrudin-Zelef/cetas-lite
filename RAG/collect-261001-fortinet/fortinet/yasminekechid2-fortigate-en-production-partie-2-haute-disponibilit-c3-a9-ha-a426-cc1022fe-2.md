---
id: collect-261001-fortinet/fortinet/yasminekechid2-fortigate-en-production-partie-2-haute-disponibilit-c3-a9-ha-a426-cc1022fe-2
title: "Vérifier la version firmware"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/yasminekechid2-fortigate-en-production-partie-2-haute-disponibilit-c3-a9-ha-a426-cc1022fe.md
source_anchor: ""
source_lines: [118, 203]
sha256: 404acaa390c1d1c3fb67f6155fce65cba1935531c462537d183b90ab63cd22be
---

# Vérifier la version firmware

- Surveiller l’état de santé du Secondary
- Appliquer des mises à jour firmware de manière contrôlée
- Diagnostiquer des problèmes sans basculement
- Effectuer des changements de configuration avec validation préalable
5. Mécanisme de Failover
Scénario de Test : Redémarrage du Primary
Pour valider le fonctionnement du cluster HA, effectuez un test de failover en redémarrant le firewall Primary :
# Sur F1 (Primary)
execute reboot
Séquence de failover :
t=0s : Le Primary (F1) démarre le redémarrage t=2s : Le Secondary (F2) détecte la perte des heartbeats sur port4 et port5 t=3s : F2 assume le rôle Primary et reprend le trafic t=3–5s : Brève interruption possible (quelques paquets perdus) t=5s+ : Trafic entièrement fonctionnel via F2
Observation côté utilisateur :
# Ping continu depuis un poste client
ping 8.8.8.8 -t# Résultat pendant le failover :
Reply from 8.8.8.8: time=15ms
Reply from 8.8.8.8: time=14ms
Request timed out.
Request timed out.
Reply from 8.8.8.8: time=16ms
Reply from 8.8.8.8: time=15ms
La continuité de service est quasi-immédiate, avec seulement 2–3 paquets perdus (soit 2–3 secondes d’interruption).
Vérification des Rôles après Failover
Après le failover :
# Sur F2 (nouveau Primary)
get system status
# Affiche : Current HA mode: a-p, master
# Sur F1 (après redémarrage, nouveau Secondary)
get system status
# Affiche : Current HA mode: a-p, backup
- F2 (nouveau primary)
Gestion de l’HA Uptime
Concept d’HA Uptime
L’HA uptime représente le temps écoulé depuis le dernier démarrage d’un firewall dans le cluster. Ce paramètre joue un rôle crucial dans la détermination du rôle Primary/Secondary.
Règle de Sélection du Primary
Lorsque l’option Override est désactivée (recommandation de production), le firewall qui devient Primary est déterminé par :
- Si différence d’uptime > 5 minutes : Le firewall avec la plus grande uptime devient Primary
- Si différence d’uptime < 5 minutes : Le firewall avec la priorité la plus élevée devient Primary
Cette logique évite les basculements répétés (flapping) et privilégie la stabilité du cluster.
Cas Pratique : Reprise après Maintenance
Situation : F1 (priorité 200) a été redémarré pour maintenance. F2 (priorité 100) est devenu Primary avec une uptime de 10 jours.
Question : F1 va-t-il reprendre le rôle Primary après redémarrage ?
Réponse : Non. Bien que F1 ait une priorité supérieure (200 > 100), son uptime est de quelques minutes seulement. La différence d’uptime étant supérieure à 5 minutes, F2 conserve le rôle Primary.
Forcer un Basculement via Reset Uptime
Pour forcer un basculement manuel, utilisez la commande :
# Sur le Primary actuel (F2)
diagnose sys ha reset-uptime
Effet :
- L’uptime de F2 est remise à zéro
- F1, ayant maintenant une uptime supérieure (> 5 min), devient Primary
- F2 redevient Secondary
Commandes de Diagnostic HA
# Afficher le statut détaillé du cluster
get system ha status
# Afficher uniquement le rôle et l'uptime
get system status | grep -i "ha\|uptime"# Afficher les statistiques de synchronisation
diagnose sys ha showcsum# Afficher l'historique des événements HA
diagnose sys ha history read
Session Pickup : Maintien des Connexions
Principe du Session Pickup
Le Session Pickup est une fonctionnalité critique qui permet de préserver les sessions actives lors d’un failover. Sans cette fonction, toutes les connexions seraient interrompues et devraient être rétablies manuellement.
Fonctionnement
- Synchronisation continue : Le Primary réplique en temps réel sa table de sessions vers le Secondary
- Préservation des états : Les connexions TCP, les sessions VPN et les translations NAT sont maintenues
- Reprise transparente : Après failover, les sessions continuent via le nouveau Primary sans réinitialisation
Configuration
Le Session Pickup est activé dans la configuration HA :
Session Pickup : Enable
Session Pickup Connection : Enable
Session Pickup NAT : Enable
Session Pickup VPN : Enable (pour IPsec VPN)
Limites
Le Session Pickup n’est pas applicable à :
- Les connexions UDP (par nature sans état)
- Certains protocoles applicatifs nécessitant une réauthentification
- Les sessions très courtes (terminées avant synchronisation)
En production, cette fonctionnalité améliore significativement l’expérience utilisateur lors des basculements planifiés (maintenance) ou non planifiés (panne).
Conclusion
La Haute Disponibilité FortiGate est une composante essentielle pour garantir la continuité de service d’une infrastructure réseau d’entreprise. La configuration présentée, bien que déployée en laboratoire, reflète les pratiques de production avec :
- Un cluster actif-passif pour éliminer le SPOF
- Des interfaces heartbeat redondantes pour fiabiliser la détection de panne
- Le Session Pickup pour préserver les connexions lors des failovers
- Des interfaces de management dédiées pour simplifier l’administration
Le basculement automatique et transparent entre Primary et Secondary garantit que les utilisateurs ne subissent qu’une interruption minime (2–3 secondes) en cas de défaillance, rendant le cluster HA quasiment imperceptible en conditions normales.
Points Clés à Retenir
✓ Le Primary traite le trafic, le Secondary surveille et est prêt à prendre le relais ✓ Les heartbeats redondants (port4 + port5) assurent une détection fiable des pannes ✓ L’HA uptime détermine le rôle Primary en cas de différence > 5 minutes ✓ Le Session Pickup maintient les connexions actives lors du failover ✓ Les interfaces de management dédiées permettent l’administration simultanée des deux firewalls
À propos de cette configuration : Le cluster HA présenté a été déployé dans un environnement de laboratoire VMware pour illustrer les concepts de haute disponibilité. Les principes et commandes sont directement applicables en production, avec les adaptations nécessaires selon votre infrastructure réseau.
