---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/guides-guide-carp-haute-disponibilite-opnsense-pdf-d2e0a144-2
title: "OPNsense-1 (MASTER)"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/guides-guide-carp-haute-disponibilite-opnsense-pdf-d2e0a144.md
source_anchor: ""
source_lines: [169, 220]
sha256: 1b4da6e64c086f36e07e9293d54e70279f7be89e1236d043007e7d0e5aec7b5f
---

# Etape 1 : Verifier que le backup est sain
#    System > High Availability > Status
#    BACKUP : CARP BACKUP, pfsync OK, config synced
# Etape 2 : Basculer le trafic vers le BACKUP
#    MASTER : System > HA > Forcefully become BACKUP
#    -> BACKUP devient MASTER, prend toutes les VIPs
# Etape 3 : Mettre a jour l'ancien MASTER (maintenant BACKUP)
#    System > Firmware > Updates > Upgrade
#    -> Redemarrage normal, aucun impact sur le trafic
# Etape 4 : Verifier le retour de l'ancien MASTER
#    -> Il revient en mode BACKUP automatiquement
#    -> Verifier sync states + config sync OK
# Etape 5 : Repeter pour l'autre noeud
#    Resultat : 2 noeuds mis a jour, 0 seconde de downtime
8. Monitoring CARP avec Grafana
Integration avec la stack Grafana + InfluxDB du Billet 9 pour monitorer l'etat CARP en temps reel.
Collecte des metriques CARP via Telegraf
# Dans telegraf.conf (sur le serveur de monitoring) :
[[inputs.http]]
  urls = ["http://OPNSENSE-VIP/api/diagnostics/interface/getVipStatus"]
  method = "GET"
  username = "telegraf_user"
  password = "VOTRE-TOKEN-API"
  data_format = "json"
  name_suffix = "_carp"
# Requete Flux Grafana — etat CARP :
from(bucket: "opnsense")
  |> range(start: -1h)
  |> filter(fn: (r) => r._measurement == "http_carp")
  |> filter(fn: (r) => r._field == "status")
Panneaux recommandes (Dashboard Grafana CARP)
 Etat CARP actuel (MASTER/BACKUP) pour chaque VIP — Gauge
 Historique des basculements — Timeline/Log panel
 Compteur pfsync : etats synchronises/seconde — Time series
 Latence lien SYNC — Gauge (alert si > 10ms)
 Nombre d'etats firewall synchronises — Stat panel
9. Prochaines etapes — SIEM Wazuh
Guide Pratique — Haute Disponibilite OPNsense CARP
BOTUM INC.
www.botum.ca  |  contact@botum.ca
Page 6

Vous avez maintenant une infrastructure OPNsense en haute disponibilite, resiliente aux pannes et aux
maintenances. La prochaine etape de la serie est le Billet 11 : integration d'un SIEM avec Wazuh pour
centraliser les logs securite et detecter les intrusions.
Article complet : blog.botum.ca/opnsense-carp-haute-disponibilite
Site : www.botum.ca | 
 — Canada
Guide Pratique — Haute Disponibilite OPNsense CARP
BOTUM INC.
www.botum.ca  |  contact@botum.ca
Page 7
