---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-24
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [2064, 2154]
sha256: c2b8c601dcb3c29e51d6bdc787d667824904c1f3024f076cdb4e1ceddb61c8b6
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- **Fiches techniques constructeur** (e-file.huawei.com, distributeurs) : AR180/AR280, AP361/AP761/AP266/AP673H, S220/S310/S620, USG6000F-S — la source de vérité pour les specs.
- **Communiqués Huawei** (MWC 2025, Intelligent Office 2.0, e.huawei.com) : positionnement, nouveautés, chiffres de la gamme.
- **ekit.huawei.com** : portail officiel — à consulter avant chaque chiffrage.
- **support.huawei.com** : guides d'installation, release notes — la lecture obligatoire avant chaque MAJ.
- **Tes guides workspace** : `onduleurs_ups_guide.md` (énergie), guides supervision/réseau (NMS complémentaire), `proxmox_guide.md` (si NMS auto-hébergé).
- **La règle d'or** : une info de ce guide contredite par une fiche technique récente ? **La fiche a raison.** Signale-le et mets le guide à jour.

## 193. Historique des versions du guide

| Version | Date | Contenu |
|---|---|---|
| 1.0 | 2026-09-27 | Création : 195 sections, recherche web de cadrage (gamme vérifiée : AR180/Pro/280/281, AP361/761/266/572/673E/673H/772, S220/S310/S620, USG6000F-S, eKitOptix, app eKit, SNC). Sections marquées « à vérifier sur la documentation officielle » là où la source manquait. |

(Tiens ce tableau à jour à chaque mise à jour significative — un guide sans historique est un guide dont on ne sait plus ce qui est frais.)

## 194. Anti-sèche ultime — une page pour tout retenir

```
eKit = PME simple, app+cloud gratuits annonces, 70+ pays, depuis 2023.
AR : 180 (petit) / 180 Pro (8 AP) / 280 (16 AP, tete PME) / 281 (2026, 5-en-1).
AP : 361 (base Wi-Fi 6) / 761 (ext.) / 572-673E (Wi-Fi 7 dense) /
     673H iGuard (spycam, local only) / 772-772E (ext. Wi-Fi 7).
Switch : S220-8P4S (125 W) / S310-48P4S/4X (380 W, SFP/SFP+) / tout-optique J.
USG6000F-S : S125/150/200 = ~600/1000/2500 users (debit tout active !).
Methodes : Internet d'abord | modem->AR->switch->AP | scan SN avant plafond |
           3 SSID max | VLAN invites isole | PoE +30% | tester en marchant |
           sauvegarde jour J | PV signe | revoir cloud J+2/J+7.
Depannage : app d'abord | site rouge = amont | AP rouge = PoE/cable |
            cycle PoE a distance | une hypothese, un test.
Limites : pas de 802.1X pousse, pas de routage dyn., pas de redondance :
          -> Datacom (NCE/S600) quand ca devient serie (grille §83).
Or : documenter (fiche site, registre, photos), sauvegarder, auditer 1x/an,
     contrat de maintenance avec KPI, distributeur = allie.
```

## 195. Dernier mot

Un réseau eKit bien déployé, c'est 90 % de **méthode** et 10 % de matériel. Les AP et les switchs font ce qu'on leur demande — c'est la préparation (survey, adressage, PoE calculé), la rigueur (sauvegarde, fiche site, étiquetage) et le suivi (rituel cloud, audits, KPI) qui font la différence entre « le Wi-Fi marche » et « on ne pense jamais au Wi-Fi ». Ce guide t'a donné la méthode ; le terrain fera le reste. Bon déploiement, chef.

---

## 196. Procédure — migrer un SSID sans couper les utilisateurs

1. Crée le **nouveau SSID** à côté de l'ancien (ex. fictif : `BUREAU-New` à côté de `BUREAU`), même VLAN, nouveau PSK robuste.
2. Migre les utilisateurs **par vagues** (un service par jour) : ils se connectent au nouveau, tu vérifies.
3. Quand tout le monde est migré : **désactive** l'ancien SSID (ne le supprime pas tout de suite — 1 semaine de battement).
4. Surveille les alertes « client qui cherche l'ancien SSID » (il reste toujours un oublié).
5. Supprime l'ancien SSID, mets à jour la fiche site et le coffre.
**Jamais** de changement de PSK « à chaud » sur le SSID de production un lundi à 9 h : c'est 20 personnes qui t'appellent en même temps.

## 197. Procédure — remplacer une passerelle AR (panne ou upgrade)

1. **Sauvegarde** complète : config cloud du site + export local + fiche site (la config WAN/PPPoE est critique — identifiants opérateur dans le coffre).
2. Prépare la **nouvelle AR** : onboard sur le site **avant** de toucher à l'ancienne si possible (pré-staging).
3. Fenêtre de maintenance annoncée (coupure Internet du site pendant l'opération).
4. Permute : WAN (opérateur) → nouvelle AR → LAN vers switch. Vérifie : Internet, DHCP, DNS, VPN s'il y en a.
5. La nouvelle AR récupère la config du site via le cloud → **vérifie quand même** : SSID, VLAN, règles, portail.
6. Tests complets (check-list 74, version courte), puis RMA de l'ancienne (151).
7. Mets à jour fiche site (nouveau SN, nouveau firmware) et registre.

## 198. Procédure — ajouter un 2e accès Internet (secours)

- **Objectif** : si la fibre tombe, le site bascule sur un 2e lien (4G/5G ou 2e opérateur) — au minimum pour les fonctions critiques (caisses, TPE).
- **Prérequis** : une passerelle avec 2 ports WAN configurables (AR280 : ports LAN/WAN commutables ; vérifier le basculement automatique sur le modèle exact — **à vérifier sur la documentation officielle**).
- **Logique** : WAN1 = fibre (prioritaire), WAN2 = 4G (secours) ; bascule sur détection de panne (ping de supervision) ; retour automatique au rétablissement (avec hystérésis pour éviter le yoyo).
- **Pièges** : la 4G a un débit et un quota limités → en secours, **brider** les invités et le streaming ; tester la bascule **pour de vrai** (débrancher la fibre) 1×/trimestre ; noter la consommation data.
- **Alternative simple** : un routeur 4G indépendant branché en secours manuel (moins élégant, mais robuste et pas cher).

## 199. Procédure — durcir un site existant (sans tout reconstruire)

1. **Audit** (155) : état des lieux écrit, priorités.
2. Mots de passe : changer tous les défauts, PSK robustes uniques, rotation des invités.
3. VLAN : isoler les invités (souvent le point le plus critique et le plus rapide à corriger).
4. Ports switch inutilisés : **désactiver**. SSID temporaires : supprimer.
5. Firmwares : monter à N/N-1 (procédure 42).
6. SNMP : passer en v3, changer les strings. Telnet : désactiver si présent.
7. Sauvegarde + documentation : fiche site à jour, coffre à jour.
8. USG : si le client a des données sensibles et n'a qu'une AR, proposer l'ajout (argumentaire section 7).
Un durcissement se fait **en journée** (pas de coupure si bien mené), VLAN par VLAN, avec rollback possible à chaque étape.

## 200. Procédure — déménager un site (transfert de locaux)

1. **Avant** : sauvegarde complète, photos de la baie, inventaire SN, préavis opérateur (résiliation/transfert du lien).
2. **Pendant** : démonter en notant (quel AP allait où — les étiquettes servent ici), protéger les AP (cartons d'origine si gardés).
3. **Nouveau local** : refaire le **survey** (les murs ne sont pas les mêmes !), adapter le plan (nombre/position des AP), câblage neuf réceptionné (187).
4. **Remise en service** : ordre section 70, comme un nouveau site — parce que c'en est un.
5. **Après** : nouveau nom de site si besoin dans le cloud (ou nouveau site + transfert des équipements), fiche site neuve, PV de recette (160).
Un déménagement = un nouveau déploiement déguisé. Ne le chiffre jamais « comme une simple dépose/repose ».

## 201. Négocier avec le distributeur : les 8 leviers

