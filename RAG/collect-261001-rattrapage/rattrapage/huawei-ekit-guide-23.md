---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-23
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["sol", "throughput"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1954, 2063]
sha256: 3de9e4a772348b5deb76ff7dfaaa2e9fcacdceaa4282df91047c49aef9f73efc
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- Aucune intervention ne doit dépendre d'**une seule personne** : tout est dans le dossier client (181) + le coffre (102) + ce guide.
- Passation écrite : sites sensibles du moment, MAJ en cours, incidents ouverts, « ne pas toucher à ».
- Teste une fois par an : « si je suis injoignable 1 semaine, qui fait quoi ? » — si la réponse est « personne », corrige.
- C'est aussi ta **valorisation** : une entreprise dont le savoir est documenté vaut plus qu'une entreprise « dans la tête du chef ».

## 183. Fin de contrat : le transfert propre

1. **Préavis** : comme écrit au contrat (jamais de coupure brutale).
2. **Transfert des accès** : tenant cloud au client ou à son nouveau prestataire, avec inventaire écrit (comptes, rôles, ce qui est transféré).
3. **Exports** : configurations, plans, fiches sites, historique des incidents — le client a payé pour, c'est à lui.
4. **Retrait de tes accès** : le jour du transfert, pas 6 mois après (ni 6 mois avant).
5. **Clôture** : PV de transfert signé, factures soldées, retour des équipements de secours.
Un départ propre = un client qui te **recommande** même en partant. Un départ sale = une réputation durablement abîmée.

## 184. Relation distributeur : ton allié n°1

- Ton **Gold Partner eKit** : stock, prix, RMA, firmwares, formations, escalade vers Huawei. C'est ton support niveau 1 — cultive la relation.
- À négocier : remises par volume, **stock de secours** dédié, délais RMA écrits, accès aux formations, un interlocuteur nommé (pas un ticket anonyme).
- Remonte les bugs firmware **avec des preuves** (logs, captures, procédure de reproduction) : un bug bien documenté est corrigé, un « ça marche pas » est ignoré.
- Demande les **roadmaps produits** : savoir 6 mois à l'avance qu'un modèle est remplacé t'évite de chiffrer du obsolète.

## 185. Extension de garantie et fin de vie : anticiper

- Note la **date d'achat + durée de garantie** de chaque équipement dans la fiche site — un RMA refusé pour « garantie expirée depuis 2 mois » alors que tu aurais pu l'étendre, c'est rageant.
- **Fin de vie / fin de support** d'un modèle : le distributeur l'annonce — prévois le renouvellement 12 mois avant (budget client).
- Règle de renouvellement : un switch/AR se garde **5-7 ans**, un AP **4-6 ans** (le Wi-Fi évolue vite : un AP Wi-Fi 6 en 2030 sera un frein). Planifie le refresh dans la roadmap client au lieu de le subir.

## 186. Annexe — modèle de fiche d'intervention

```
FICHE D'INTERVENTION n° __________
Date/heure : __________  Client/site : __________
Technicien : __________  Type : [ ] install [ ] depannage [ ] maintenance [ ] audit
Motif : __________________________________________
Constats : __________________________________________
Actions realisees : __________________________________________
Materiel remplace (SN ancien/nouveau) : __________
Config modifiee : [ ] oui (sauvegarde : __________) [ ] non
Tests effectues : __________________________________________
Temps passe : __________  Deplacement : __________
Reste a faire : __________________________________________
Signature client : __________  Signature technicien : __________
```
Une fiche par intervention, classée dans le dossier client (181). C'est ta preuve, ta mémoire et ta base de facturation.

## 187. Annexe — check-list de réception du câblage (avant de brancher l'actif)

- [ ] Chaque brin testé au testeur (continuité, paires, longueur < 90 m)
- [ ] Étiquettes aux deux extrémités, lisibles, conformes au plan
- [ ] Prises murales fixées, caches posés, pas de câble en tension
- [ ] Chemins de câbles fermés, pas de câble au sol dans les zones de passage
- [ ] Terre de la baie vérifiée (< 5 ohms visé)
- [ ] Distances PoE respectées (brins AP/caméras < 90 m)
- [ ] Photos du cheminement avant fermeture des faux plafonds (sinon c'est perdu)
- [ ] Plan de brassage à jour et affiché dans la baie
**Ne branche jamais l'actif sur un câblage non réceptionné** : 80 % des « pannes eKit » sont des pannes de câblage.

## 188. Annexe — lexique anglais-français (pour lire les docs constructeur)

| Anglais (docs) | Français | Note |
|---|---|---|
| Onboarding | Rattachement au cloud | Le geste n°1 |
| Uplink | Lien montant (vers le cœur/Internet) | — |
| Downlink | Lien descendant (vers les AP/clients) | — |
| Throughput | Débit utile | ≠ débit radio théorique |
| Forwarding rate (pps/Mpps) | Capacité de commutation en paquets/s | Le vrai muscle d'un switch |
| Switching capacity | Capacité de commutation (Gbit/s) | — |
| Egress / Ingress | Sortant / entrant | — |
| Power budget | Budget PoE total du switch | À ne jamais dépasser |
| Perpetual / Fast PoE | PoE maintenu au reboot / réalim. rapide | — |
| Heat dissipation | Dissipation thermique | Pour le dimensionnement clim |
| MTBF / MTTR | Temps moyen entre pannes / de réparation | Ordres de grandeur constructeur |
| Stacking | Empilement logique de switchs | Propriétaire : même marque |
| Zero-touch provisioning | Mise en service sans intervention | L'idéal eKit via scan SN |
| Captive portal | Portail captif | — |
| Site survey | Étude de couverture radio | À pied, avec mesures |
| RMA | Retour garantie | Via le distributeur |

## 189. Annexe — lire une fiche technique sans se faire piéger

- **« Jusqu'à »** = valeur labo, divise par 2 en conditions réelles pour ton dimensionnement mental.
- **Débit Wi-Fi agrégé** (ex. 1,775 Gbit/s) = somme théorique des bandes, jamais atteinte par un seul client.
- **Débit pare-feu brut** vs **débit tout activé** : seul le 2e compte (section 7).
- **« Recommandé pour X utilisateurs »** : c'est pour de la bureautique légère — divise par 2 si visio intensive.
- **Température de fonctionnement** : vérifie qu'elle colle au local (un switch à -5/+50 °C dans un local à 55 °C l'été = panne).
- **Garantie** : la fiche dit rarement la durée — demande au distributeur, par écrit.
- **« Cloud-managed »** : vérifie *ce qui* est gérable dans le cloud (tout ? juste le monitoring ?) — le diable est dans le détail, et dans la version du firmware.

## 190. Annexe — les 7 péchés capitaux du déploiement eKit

1. **L'orgueil** : « pas besoin de survey, je connais » → zones blanches.
2. **L'avarice** : 1 AP pour 300 m² « pour économiser » → tout le monde rame, on en rajoute 2 après (plus cher).
3. **La paresse** : pas de sauvegarde, pas de fiche site → 6 mois plus tard, on redécouvre tout.
4. **La gourmandise** : 5 SSID « au cas où » → airtime bouffé par les beacons.
5. **La colère** : reset usine en pleine journée « pour voir » → site coupé, client furieux.
6. **L'envie** : du Wi-Fi 7 partout « comme le concurrent » → budget explosé pour un besoin Wi-Fi 6.
7. **La luxure** : (on la remplace par) **l'approximation** : PoE « à peu près », adressage « on verra » → la dette technique se paie toujours, avec intérêts.

## 191. Ce qu'il faut surveiller en 2026-2027 (veille)

- **Disponibilité réelle** de l'AR281, des AP772/AP772E, du F700D dans ton pays (annoncés ≠ stockés).
- **Bande 6 GHz** : évolution réglementaire locale (impact direct sur l'AP673H).
- **Fonctions SNC** : le cloud s'enrichit vite (troubleshooting 2.0 aujourd'hui, quoi demain ?) — suis les release notes.
- **Conditions de licence** : le « gratuit sans licence » d'aujourd'hui — vérifier qu'il dure, surtout sur les fonctions IA/USG.
- **Wi-Fi 8 (802.11bn)** : en préparation dans l'industrie — ne fige pas un parc pour 10 ans sur du Wi-Fi 6 sans trajectoire.
- **Tes concurrents locaux** : que proposent-ils ? À quel prix ? (Le meilleur argumentaire se construit en connaissant l'adversaire.)
- **Ce guide** : relis-le une fois par an et mets à jour les sections marquées « à vérifier » — c'est ta dette documentaire.

## 192. Bibliographie commentée (d'où vient ce guide)

