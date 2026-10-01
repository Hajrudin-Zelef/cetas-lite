---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-14
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["exploit", "incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [2112, 2226]
sha256: 5df18c650f91b336cfe6b0bbd9fc2a45bf0d0b8a8ca177181c5f2051c09662f5
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

1. **Community SNMP « public » en lecture sur tout le parc** — avec des
   ACL inexistantes. N'importe qui cartographie votre réseau.
2. **SNMP v2c là où v3 est possible** — par paresse. Le trafic de
   supervision en clair sur le réseau de management.
3. **Trap-target oublié** sur les nouveaux équipements — eSight ne voit
   les pannes qu'au polling suivant (ou jamais).
4. **Timeouts trop courts** vers les sites distants — fausses alertes
   « down » en série, puis on ne croit plus les alertes.
5. **Mot de passe SNMP v3 changé** en masse sans mettre à jour le profil
   eSight — découverte morte du jour au lendemain.
6. **Changement d'IP du serveur eSight** sans recenser les références
   (trap-targets, ACL, firewall, DNS) — tout casse (cas n°15).
7. **Aucune sauvegarde testée** — on découvre le jour du sinistre que
   les backups échouaient depuis 4 mois.
8. **Base jamais purgée** — eSight rame, puis se fige (cas n°12-13).
9. **Notifications jamais testées** — le relais SMTP a changé il y a
   6 mois, personne n'est prévenu depuis.
10. **Tout le monde admin** — 12 comptes administrators, zéro traçabilité.
11. **Acquittement sans commentaire** — « quelqu'un s'en occupe »…
    personne ne s'en occupe.
12. **Masquages permanents** — des pannes invisibles pendant des mois.
13. **200 alarmes non acquittées en permanence** — le NOC ne regarde plus
    la console. Recalibrer (section 50).
14. **Topologie jamais nettoyée** — équipements fantômes, carte mensongère.
15. **Pas de runbook** — chaque incident se réinvente à 3h du matin.
16. **eSight non supervisé** — quand il tombe, silence radio général.
17. **Upgrade sans snapshot ni rollback** — la mise à jour qui transforme
    un vendredi en week-end de crise (cas n°17).
18. **Modules installés ≠ modules licenciés** — fonctionnalités en erreur
    au démarrage, diagnostic obscur.
19. **Rétention infinie des données de performance** — disque plein,
    rapports vides, eSight à genoux.
20. **Investir dans eSight sans regarder iMaster NCE** — déployer un NMS
    classique aujourd'hui sans feuille de route 3-5 ans, c'est s'exposer
    à une migration subie (section 5).

---

# 20. GLOSSAIRE

## 114. Glossaire eSight / supervision

- **AC (Access Controller)** : contrôleur Wi-Fi qui pilote les points d'accès.
- **Acquittement** : prise en charge d'une alarme par un exploitant
  (« je m'en occupe ») — n'éteint pas le défaut.
- **Agrégation** : regroupement d'alarmes répétitives identiques en une
  seule alarme avec compteur.
- **AP (Access Point)** : point d'accès Wi-Fi.
- **B/S (Browser/Server)** : architecture où l'administration se fait
  via navigateur, sans client lourd.
- **Clear** : fin d'une alarme (le défaut a disparu).
- **Community SNMP** : « mot de passe » partagé de SNMP v1/v2c (faible
  sécurité — préférer SNMP v3).
- **Corrélation** : rattachement automatique d'alarmes conséquences à
  une alarme cause racine.
- **FCAPS** : les 5 domaines de la gestion réseau — Fault, Configuration,
  Accounting, Performance, Security.
- **Flapping** : bascule répétée up/down d'un lien ou d'un équipement.
- **Hot standby** : secours à chaud — un nœud passif prêt à prendre le
  relais immédiatement.
- **IPFIX** : standard d'export de flux IP (NetStream en est proche).
- **Masquage** : inhibition temporaire de la remontée d'alarmes
  (maintenance, dérangement connu) — toujours avec date de fin.
- **MIB** : base d'information de management — dictionnaire des objets
  supervisés via SNMP.
- **MPLS VPN** : VPN d'opérateur basés sur MPLS (supervision en édition
  Standard/Professional).
- **NE (Network Element)** : équipement réseau supervisé.
- **NetStream** : technologie Huawei d'export de flux de trafic
  (équivalent NetFlow), exploitée par le NTA d'eSight.
- **NMS** : Network Management System — système de supervision réseau.
- **Northbound** : interface d'eSight vers un système de niveau supérieur.
- **NTA (Network Traffic Analyzer)** : analyse du trafic dans eSight
  (édition Standard+).
- **Polling** : interrogation périodique des équipements par eSight.
- **PRA** : Plan de Reprise d'Activité.
- **RPO / RTO** : perte de données max acceptable / durée max de
  reprise — objectifs du PRA.
- **Sévérité** : criticité d'une alarme (Critical, Major, Minor,
  Warning, Info).
- **SLA** : accord de niveau de service — objectifs mesurés (disponibilité…).
- **SNMP** : Simple Network Management Protocol — protocole de supervision.
- **Southbound** : interface d'eSight vers les équipements supervisés.
- **Storm (tempête d'alarmes)** : avalanche d'alarmes suite à un
  incident majeur.
- **Syslog** : protocole de remontée de journaux d'événements.
- **Toggling** : alarme intermittente qui apparaît/disparaît en boucle.
- **Trap** : notification spontanée envoyée par l'équipement vers eSight.
- **v3 authPriv** : mode SNMP v3 avec authentification + chiffrement
  (le niveau à viser).

---

# 21. QUIZ

## 115. Quiz : 10 questions pour valider

**Q1.** Quelles sont les quatre fonctions historiques (FCAPS partielles)
assurées par eSight ?
**Q2.** Quelle édition minimale faut-il pour la gestion WLAN et le NTA ?
**Q3.** En SNMP, quelle est la différence fondamentale entre v2c et v3,
et lequel faut-il privilégier ?
**Q4.** Un équipement n'est pas découvert alors qu'il répond au ping.
Citez trois vérifications dans l'ordre.
**Q5.** Que signifie « acquitter » une alarme, et que ne signifie-t-il pas ?
**Q6.** À quoi servent respectivement l'agrégation, le masquage et la
corrélation des alarmes ?
**Q7.** Pourquoi faut-il une date de fin à tout masquage d'alarme ?
**Q8.** Que sauvegarder pour pouvoir reconstruire eSight après sinistre ?
**Q9.** Que faire en premier face à une tempête d'alarmes ?
**Q10.** eSight est-il en fin de vie face à iMaster NCE ? Que répondre
honnêtement à votre direction ?

### Réponses

