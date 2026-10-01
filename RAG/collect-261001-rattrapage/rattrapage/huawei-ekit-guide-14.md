---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-14
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27", "2026-09-28"]
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1215, 1320]
sha256: ada03c6146a1deca9a6c92f417b9b61ac740eb12c26586f0150804f185e3cbb5
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

## 106. Registre des accès et traçabilité (modèle)

```
REGISTRE DES ACCES - [Client]
Date       | Systeme      | Compte/Role      | Action                 | Par
2026-09-27 | eKit cloud   | admin (proprio)  | Creation site + onboard| Zelef
2026-09-28 | AR180 Pro    | admin local      | Changement mdp defaut  | Zelef
...        | ...          | ...              | ...                    | ...
```
Tenu à jour = preuve de professionnalisme en cas d'audit ou de litige. 2 minutes par intervention, valeur inestimable.

## 107. Pense-bête de poche — déploiement (à imprimer)

```
eKit DEPLOIEMENT EXPRESS
1. Internet sur l'AR d'abord (ping + DNS)
2. Ordre : modem -> AR -> switch -> AP
3. Scan SN AVANT de fixer les AP au plafond
4. 3 SSID max | VLAN invites isole | PSK 20+ car.
5. PoE : additionner + 30% de marge
6. Tester en marchant (roaming), pas assis
7. Photos : baie, SN, AP, etiquettes
8. Sauvegarde + fiche site le jour J
9. Client forme (3 gestes), astreinte communiquee
10. Revoir le cloud a J+2 et J+7
```

## 108. Pense-bête de poche — dépannage (à imprimer)

```
eKit DEPANNAGE EXPRESS
1. Ouvrir l'app : QUEL equipement est rouge ?
2. Tout le site rouge -> electrique / operateur / AR (pas les AP !)
3. 1 AP rouge -> PoE / cable (cycle PoE a distance d'abord)
4. Wi-Fi associe mais pas d'Internet -> IP ? passerelle ? DNS ? portail ?
5. "C'est lent" -> MESURER (chiffres, pas d'impression)
6. Apres MAJ qui casse -> sauvegarde d'avant-MAJ + downgrade
7. Une hypothese, un test. Noter. Ne jamais changer 3 choses a la fois.
8. Si on ne comprend pas en 30 min : escalader (distributeur / support)
```

## 109. Pense-bête de poche — les chiffres à retenir

```
AP361 : Wi-Fi 6, 1,775 Gbit/s, PoE af ~9 W, 1x GE
AP761 : Wi-Fi 6 EXT, PoE at conseille, 17,7 W max, 1x GE + 1x SFP
S310-48P4S : 48x GE PoE+, 4x SFP 1G, 380 W, 104 Gbit/s
S310-48P4X : 48x GE PoE+, 4x SFP+ 10G, 380 W, 176 Gbit/s
S220-8P4S : 8x GE + 4x SFP, 125 W PoE+
AR180/Pro : Wi-Fi 7, 4-5x GE + 1x 2.5GE WAN, 100-150 terminaux, 8-16 AP geres
AR280 : 4x GE + 1x 2.5GE, 16 AP, 32 equip. geres, 16 tunnels IPsec
USG6000F-S125/150/200 : ~600/1000/2500 utilisateurs reco.
Regles : 1 AP / 100-150 m2 (bureau placo) | 30-40 clients/AP max
         PoE : marge 30% | Cable : Cat6, 90 m max
         192.168.1.0/24 et 192.168.0.0/24 : INTERDITS en pro
```

## 110. Glossaire eKit (les 30 termes à connaître)

- **eKit** : marque distribution Huawei pour les PME (lancée 2023).
- **eKitEngine** : gamme datacom eKit (AR, AP, switchs, USG).
- **eKitOptix (MiniFTTO)** : fibre jusqu'au bureau/chambre (F700D, FG736).
- **eKitStor / IdeaHub** : stockage / écrans collaboratifs (hors réseau).
- **SNC (SME Network Center)** : cloud de gestion des réseaux eKit.
- **Onboarding** : déclaration/rattachement d'un équipement au cloud.
- **Site** : unité d'organisation dans l'app (1 bâtiment = 1 site).
- **Tenant** : compte client/fournisseur dans le cloud, contient les sites.
- **Fit / Fat / Cloud** : modes de fonctionnement d'un AP (contrôleur / autonome / cloud).
- **WAC** : contrôleur Wi-Fi (fonction intégrée aux AR280/AR281).
- **SSID** : nom du réseau Wi-Fi diffusé.
- **PSK** : clé pré-partagée (mot de passe Wi-Fi).
- **WPA2 / WPA3** : protocoles de chiffrement Wi-Fi.
- **Portail captif** : page d'authentification avant accès Internet (invités).
- **VLAN** : réseau local virtuel (cloisonnement logique, 802.1Q).
- **Trunk / Access** : port multi-VLAN taggé / port mono-VLAN.
- **PoE (af/at/bt)** : alimentation par le câble réseau (15,4 / 30 / 60-90 W).
- **Perpetual / Fast PoE** : PoE maintenu pendant reboot / réalimentation rapide.
- **HOUP** : dépôt en ligne des mises à jour Huawei (switchs).
- **MU-MIMO / OFDMA / Beamforming** : technos Wi-Fi 6 d'efficacité radio.
- **MLO / 320 MHz / 4096-QAM** : technos Wi-Fi 7 (débit, latence).
- **Band steering** : orientation des clients vers le 5 GHz.
- **Roaming (802.11r/k/v)** : passage fluide d'un AP à l'autre.
- **STA** : station = client Wi-Fi.
- **IPS / AV / URL filtering** : prévention d'intrusion / antivirus / filtrage web (USG).
- **IPsec / SSL VPN** : tunnels chiffrés site-à-site / accès distant.
- **NVR** : enregistreur vidéo réseau (fonction intégrée à l'AR281).
- **iGuard (AP673H)** : détection de caméras espion par IA (gestion locale uniquement).
- **Wi-Fi Shield** : fonctions anti-écoute mises en avant par Huawei sur le Wi-Fi 7.
- **RMA** : retour matériel en garantie via le distributeur.
- **Gold Partner** : distributeur agréé eKit (ton support niveau 1).

## 111. Quiz — 10 questions pour valider (réponses en section 112)

1. Un AP761 consomme jusqu'à 17,7 W. Sur quel standard PoE faut-il l'alimenter, et pourquoi pas en 802.3af simple ?
2. Ton S310-48P4S a un budget PoE de 380 W. Tu y branches 20 AP361 (~9 W), 2 AP761 (~17,7 W) et 8 caméras (~10 W). Le budget tient-il avec 30 % de marge ?
3. Dans l'app, tout un site passe au rouge d'un coup mais le client dit que « les fichiers en local marchent encore ». Où cherches-tu en premier ?
4. Cite les 3 modes de fonctionnement d'un AP eKit et dis dans quel cas tu choisis chacun.
5. Pourquoi ne faut-il jamais utiliser 192.168.1.0/24 sur un site professionnel ?
6. Un client veut 5 SSID « pour chaque service ». Que lui réponds-tu, et combien en proposes-tu ?
7. Après une mise à jour firmware, un site est en vrac. Quelles sont tes 3 premières actions, dans l'ordre ?
8. À partir de quels signaux dois-tu proposer de basculer du eKit vers du Datacom entreprise ? (cite-en 3)
9. L'AP673H détecte les caméras espion. Pourquoi cette fonction ne remonte-t-elle pas dans le cloud ?
10. Un hôtel veut du Wi-Fi dans 40 chambres en construction neuve. Quelles 2 architectures peux-tu proposer, et quel est l'arbitrage ?

## 112. Quiz — réponses

