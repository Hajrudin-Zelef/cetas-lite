---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-10
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [825, 950]
sha256: 8edb137ce8384b8b69a8514c5c1f823c447c9c0603a171f8718f962de93ce349
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

```
1. INFRA PASSIVE : baie, câblage, prises, terre, étiquetage. (Tester chaque brin.)
2. ÉLECTRIQUE : onduleur en place, prises ondulées identifiées.
3. TÊTE : modem/ONT opérateur -> AR (WAN). Valider INTERNET sur l'AR d'abord.
4. DISTRIBUTION : switch(s) derrière l'AR. Valider : l'AP de test obtient une IP.
5. ACCÈS : AP un par un, vérifiés dans l'app (en ligne, bon site).
6. CONFIG : VLAN, SSID, DHCP, portail captif, règles.
7. TESTS : chaque SSID, chaque VLAN, isolation, débit, roaming à pied.
8. SAUVEGARDE : export config + fiche site + photos.
9. TRANSMISSION : former le contact client (2-3 gestes : redémarrer, lire une alerte).
10. SUIVI : inspection cloud à J+2 et J+7.
```

**Pourquoi cet ordre :** chaque étape valide la précédente. Si tu configures les SSID avant d'avoir Internet, tu ne sauras jamais si « ça ne marche pas » vient du Wi-Fi ou de l'uplink.

## 71. Étiquetage et documentation physique : le professionnalisme visible

- Chaque câble : étiquette aux **deux extrémités** (`B12-P07` = baie 1, prise 2, port 7 du switch — exemple fictif de convention).
- Chaque AP : étiquette discrète avec son nom logique (`Etage1-AP02`).
- Chaque port du switch : correspond au plan (ports 1-4 = AP, 5-20 = PC…).
- Dans la baie : un **schéma plastifié** du brassage + la fiche site (section 65).
- Le client voit la différence entre « des câbles » et « une installation ». C'est aussi ce qui te fait rappeler.

## 72. Mise à la terre et parafoudre : ne pas négliger

- Les switchs eKit ont une protection interne (±6-7 kV d'après les fiches S310), mais en zone orageuse ou réseau électrique instable : **parafoudre en tête** + **terre < 5 ohms** (cf. ton guide onduleurs).
- Les **AP761 extérieurs** : mise à la terre du mât, parafoudre PoE sur la ligne si exposition (toit, zone dégagée).
- Un AP grillé par la foudre, c'est un RMA refusé (« surtension » = exclusion de garantie typique). Protège, et documente la protection.

## 73. Plan de nommage des équipements (convention proposée)

```
[SITE]-[ETAGE]-[TYPE][NUM]
Exemples fictifs :
  HOTELPALM-RDC-AR01      (passerelle)
  HOTELPALM-E1-SW01       (switch etage 1)
  HOTELPALM-E1-AP01..AP10 (AP chambres etage 1)
  HOTELPALM-EXT-AP01      (AP761 piscine)
```
Avantages : tri alphabétique = tri géographique dans l'app ; une alerte « HOTELPALM-E1-AP07 hors ligne » est immédiatement localisable.

## 74. Check-list de mise en service (à cocher sur site)

- [ ] Internet OK sur l'AR (ping 8.8.8.8 — exemple fictif d'IP publique de test — + DNS résout)
- [ ] Tous les équipements **verts** dans l'app, sur le **bon site**
- [ ] Chaque SSID testé avec un vrai client (connexion + Internet + débit)
- [ ] Isolation invités vérifiée (pas d'accès au VLAN staff, pas de ping inter-clients)
- [ ] Portail captif testé sur Android ET iPhone
- [ [ ] Imprimante et équipements filaires OK
- [ ] Caméras : image OK sur le NVR, PoE stable
- [ ] Téléphones IP : appel interne + externe
- [ ] Firmwares notés, config sauvegardée, fiche site remplie
- [ ] Photos baie + AP + étiquettes
- [ ] Client formé aux 3 gestes de base, numéro d'astreinte communiqué
- [ ] Inspection cloud planifiée à J+2

## 75. Les 10 commandements du déploiement eKit (pense-bête)

1. **Internet d'abord** : rien ne se configure avant que l'AR ait son uplink.
2. **Un site = un site** : nomme, cloisonne, documente.
3. **Pré-stage au bureau** : scanne et prépare avant d'aller sur site.
4. **PoE calculé, pas supposé** : additionne, marge 30 %.
5. **3 SSID max** : staff, invités, technique si besoin.
6. **VLAN invités isolé** : toujours, sans exception.
7. **Photos des SN avant** de fixer les AP au plafond.
8. **Teste en marchant** : le roaming ne se valide pas assis.
9. **Sauvegarde le jour J** : config + fiche + photos.
10. **Reviens à J+2/J+7** (via le cloud) : un site qu'on ne revoit jamais finit par tomber en panne un dimanche.

---

## 76. Supervision via eKit : ce que tu vois vraiment

D'après les fiches techniques, la plateforme cloud supervise **l'état du réseau, l'état des équipements et l'état des connexions clients (STA)** sur tous les sites du tenant. En pratique, organise ta supervision autour de 4 vues :
1. **Vue santé globale** : tous les sites, code couleur. Ton rituel du matin (5 min).
2. **Vue site** : équipements en ligne/hors ligne, charge, alertes du site.
3. **Vue équipement** : détail d'un AP/switch/AR (état, clients associés, version firmware, uptime — niveau de détail **à vérifier** selon version).
4. **Vue clients** : qui est connecté où (utile pour « le Wi-Fi est lent dans la salle de réunion » → 30 clients sur 1 AP).

## 77. Les alertes qui comptent (et celles qu'on peut ignorer)

| Alerte | Priorité | Action |
|---|---|---|
| Équipement hors ligne | **Haute** | Vérifier cloud (tout le site ? → uplink/électrique ; un seul AP ? → PoE/câble), puis cycle PoE, puis déplacement |
| Nouvel équipement détecté | Moyenne | Normal après ajout ; suspect sinon (intrusion ?) |
| Taux d'utilisation CPU/mémoire élevé | Moyenne | Surveiller, planifier remplacement si récurrent |
| Client DHCP épuisé (scope plein) | Moyenne | Élargir le scope ou réduire le bail |
| Interférences / changement de canal fréquent | Basse | Survey radio si récurrent |
| Firmware disponible | Basse (info) | Planifier selon section 42 |

**Règle :** une alerte sans action définie = du bruit. Pour chaque type d'alerte, écris qui fait quoi (ta procédure d'astreinte, section 101).

## 78. Supervision eKit vs vrai NMS : le comparatif honnête

| Fonction | eKit (app/SNC) | Vrai NMS (Zabbix, PRTG, NCE…) |
|---|---|---|
| État en ligne/hors ligne | Oui | Oui |
| Alertes de base | Oui | Oui, bien plus fines |
| Historique de performance (débit, erreurs) | Limité (**à vérifier**) | Oui, graphes long terme |
| NetFlow / sFlow / analyse de trafic | Non (à ma connaissance) | Oui |
| Supervision d'équipements tiers | Non | Oui (SNMP générique) |
| Corrélation d'événements | Non | Oui (règles) |
| API / webhooks / intégration ticketing | À vérifier sur la documentation officielle | Oui |
| Cartographie réseau auto | Basique (à vérifier) | Oui |
| Coût | Inclus (annoncé gratuit) | Licence / infra à prévoir |

**Conclusion :** eKit suffit pour **surveiller** 5-50 sites PME. Dès que tu veux **analyser** (pourquoi c'est lent tous les jours à 14 h ?), il te faut un NMS complémentaire — même simple (un Zabbix qui ping les équipements eKit en SNMP si les modèles l'exposent — **à vérifier** par modèle).

## 79. Compléter eKit avec un NMS léger : la bonne combinaison

- Garde eKit pour : onboarding, config, alertes « équipement hors ligne », gestion multi-sites.
- Ajoute un NMS (Zabbix/PRTG/LibreNMS — tu as des guides complets dans ton workspace) pour : ping/SNMP des passerelles et switchs, graphes de débit des uplinks, disponibilité sur 30 jours (utile pour prouver un SLA ou diagnostiquer un problème récurrent).
- Les switchs eKit supportent **SNMPv1/v2c/v3** d'après les fiches (S220, S310, S620) : c'est ta porte d'entrée NMS. Active SNMPv3 (authentifié/chiffré), jamais v1/v2c en production si tu peux l'éviter.
- Documente les **community strings / credentials SNMP** dans ton coffre (section 102), pas dans la fiche site.

## 80. Les KPI à suivre par site (tableau de bord mensuel client)

| KPI | Cible (ordre de grandeur) | Source |
|---|---|---|
| Disponibilité équipements | > 99,5 % (PME) | Cloud eKit / NMS |
| Temps moyen de détection panne | < 15 min | Alertes cloud |
| Temps moyen de résolution | < 4 h (ouvré) | Tes tickets |
| Clients Wi-Fi max simultanés / AP | < 30-40 (bureautique) | Vue clients cloud |
| Taux de réussite portail captif | > 95 % | Tests / retours |
| Firmwares à jour | 100 % des sites (N-1 max) | Cloud eKit |

