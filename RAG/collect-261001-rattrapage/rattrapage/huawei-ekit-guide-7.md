---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-7
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "capex"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [477, 605]
sha256: b231d442d6c56987661fde560f3386c3c12e152b3a6672901a1f03bda787e8d4
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

```
                    [Internet Fibre Operateur]
                               |
                          (ONT/Box operateur en mode bridge)
                               |  (WAN - PPPoE ou DHCP)
                    +---------------------+
                    |  AR180 Pro (passerelle)|
                    |  192.168.10.1 (LAN)   |
                    |  DHCP + DNS relay     |
                    +----------+----------+
                               | (LAN trunk vers switch)
                    +----------+----------+
                    | S310-24 ports PoE+   |
                    | (ex: S310-24P4S)     |
                    +--+---+---+---+--+---+
                       |   |   |   |  |
        +--------------+   |   |  |  +--------------+
        |                  |   |  |                 |
   [AP361 #1]         [AP361 #2] |           [AP361 #3]
   Open space         Salle reunion|           Bureaux fermes
   (plafond)          (plafond)    |
                              [PC filaires: ports 5-20, VLAN 10]
                              [Imprimante: port 21, VLAN 10]
                              [Cameras: ports 22-23, VLAN 40, PoE]

VLAN:
  VLAN 1   (defaut, eviter pour l'usage) 
  VLAN 10  STAFF      192.168.10.0/24  (PC + Wi-Fi staff)
  VLAN 20  GUESTS     192.168.20.0/24  (Wi-Fi invites, isole, portail captif)
  VLAN 40  CAMERAS    192.168.40.0/24  (si cameras)
```

## 48. Scénario A — Liste matériel

| Qté | Équipement | Rôle | Notes |
|---|---|---|---|
| 1 | AR180 Pro | Passerelle (routage, DHCP, Wi-Fi de secours) | Gère jusqu'à 8 AP / 32 équipements : de la marge |
| 1 | S310-24P4S (ou 24 ports PoE+ équivalent eKit) | Switch d'accès PoE | Vérifier la référence exacte dispo chez ton distributeur |
| 3 | AP361 | Wi-Fi intérieur | 1 / ~100-150 m² en open space placo |
| 1 | Onduleur 1000-1500 VA | Protection électrique | AR + switch + ONT dessus (cf. ton guide onduleurs) |
| — | Câblage Cat6, goulottes, baie 6U | Infra passive | 25 prises RJ45 murales |
| 2 | Caméras IP (option) | Vidéosurveillance | VLAN 40 dédié |

**Budget PoE (vérification) :** 3× AP361 (~9 W) = 27 W + 2 caméras (~10 W) = ~47 W. Un switch 24 ports PoE+ (~190-370 W selon modèle) est large. **Toujours** calculer : ne suppose jamais.

## 49. Scénario A — Configs clés (via app eKit / cloud)

1. **Accès Internet AR** : PPPoE (identifiants opérateur — fictifs ici : `client@fai.xx` / mot de passe à saisir, jamais noté en clair dans ce guide) ou DHCP selon l'opérateur. DNS : ceux de l'opérateur ou publics.
2. **VLAN** : créer 10 (staff), 20 (invités), 40 (caméras). Le lien AR↔switch en **trunk** (tous les VLAN taggés).
3. **Ports switch** : ports 1-3 (AP) en trunk avec VLAN natif 10 ; ports 5-20 en access VLAN 10 ; port 21 (imprimante) access VLAN 10 ; ports 22-23 access VLAN 40 + PoE activé.
4. **SSID** : `BUREAU-Staff` (WPA2/WPA3, VLAN 10, SSID non diffusé en option) ; `BUREAU-Invites` (portail captif, VLAN 20, isolation client, limite 10 Mbit/s/client — valeurs fictives d'exemple).
5. **DHCP** : un scope par VLAN sur l'AR (10 : 192.168.10.50-200 ; 20 : 192.168.20.50-200, bail court 4 h ; 40 : 192.168.40.10-50).
6. **Tests** : client filaire (ping passerelle + Internet), client Wi-Fi staff, client invité (portail OK ? isolation OK ? pas d'accès au VLAN 10 ?), débit speedtest à 3 endroits.

## 50. Scénario B — Hôtel 40 chambres : cadrage

**Besoin type :** hôtel 3 étoiles, 4 étages × 10 chambres + réception + restaurant + piscine extérieure. 60-80 clients simultanés le soir, portail captif avec code chambre, 8 caméras, 5 téléphones IP (réception), TV IP (option).
**Deux architectures possibles :**
- **Classique cuivre** : 1 AR280 (ou USG6000F-S125 si besoin pare-feu) + 2× S310-48P4S + ~20 AP361 (1 pour 2 chambres) + 2 AP761 (piscine/extérieur).
- **MiniFTTO** : OLT + 1 F700D par chambre (fibre). Plus cher à l'achat, imbattable en esthétique et en évolutivité — à proposer en neuf/rénovation lourde uniquement.

## 51. Scénario B — Topologie ASCII (version cuivre)

```
                    [Internet Fibre 1-2 Gbit/s]
                               |
                    +---------------------+
                    | AR280 ou USG6000F-S |
                    | (pare-feu + routage) |
                    +----------+----------+
                               | (trunk 10G si possible)
              +----------------+----------------+
              |                                 |
   [S310-48P4S #1 - Etages 1-2]      [S310-48P4S #2 - Etages 3-4 + RDC]
      |      |      |      |              |        |         |
   [AP361] x10 (1 pour 2 chambres,   [AP361] x10   [AP761] x2
    couloir, plafond)                 + RDC:      (piscine, terrasse,
    VLAN 30 (chambres)                 reception    exterieur IPxx)
                                       restaurant
                                       VLAN 30/50
   [Cameras: VLAN 40]                [Telephones IP: VLAN 60]
   [TV IP: VLAN 70 si present]

VLAN:
  VLAN 10  MGMT/STAFF    10.10.10.0/24   (personnel, recep.)
  VLAN 30  CHAMBRES      10.10.30.0/24   (clients, portail captif, isole)
  VLAN 40  CAMERAS       10.10.40.0/24   (VMS/NVR)
  VLAN 50  RESTAURANT    10.10.50.0/24   (caisse, Wi-Fi clients resto)
  VLAN 60  VOIX          10.10.60.0/24   (telephones IP, QoS prioritaire)
  VLAN 70  TV-IP         10.10.70.0/24   (si TV sur IP, multicast)
```

**Règles radio hôtel :** puissance d'émission **modérée** (mieux vaut plus d'AP à faible puissance que peu d'AP à fond — sinon les clients s'accrochent à l'AP du couloir d'en face), canaux 1/6/11 en 2,4 GHz répartis par étage, 5 GHz en priorité (band steering si dispo).

## 52. Scénario B — Variante MiniFTTO (eKitOptix)

```
   [Internet] -> [OLT eKitOptix] -> [Splitter optique 1:32]
                                              |
                    +------------+------------+------------+
                    |            |                         |
              [F700D Ch.101] [F700D Ch.102] ... [F700D Ch.140]
              (1 fibre / chambre : Wi-Fi 7 + 2-4 ports GE + TV/tel.)

Avantages : 1 seul cable (fibre) par chambre, aucun switch d'etage,
            Wi-Fi 7 homogene, esthetique (boitier mural discret).
Contraintes : OLT + splitters + soudures fibre (competence fibre requise),
              cout initial superieur, ecosysteme ferme.
```

**Quand proposer quoi :** rénovation légère / existant cuivre → version cuivre (AP361). Construction neuve ou rénovation lourde avec faux plafonds ouverts → chiffrer les deux et laisser le client choisir avec un comparatif honnête (CAPEX vs OPEX vs esthétique).

## 53. Scénario B — Liste matériel (version cuivre)

| Qté | Équipement | Rôle |
|---|---|---|
| 1 | AR280 ou USG6000F-S125 | Tête de réseau (si pare-feu exigé par le client → USG) |
| 2 | S310-48P4S (380 W PoE+) | Distribution étages |
| 20 | AP361 | Chambres (1/2 chambres) + couloirs |
| 2 | AP761 | Extérieur (piscine, terrasse) — prévoir parafoudre et terre |
| 8 | Caméras IP | VLAN 40, vers NVR |
| 5 | Téléphones IP | VLAN 60 |
| 1 | Onduleur 2000-3000 VA | Tête + switchs (autonomie 15-30 min mini) |
| — | Câblage Cat6, 60+ prises | Chambres + paliers |

**Budget PoE :** 20× 9 W = 180 W + 2× 17,7 W = 35,4 W + 8× 10 W = 80 W + 5× 5 W = 25 W → **~320 W**. Deux S310-48P4S (380 W chacun) : OK avec marge si bien réparti (ne mets pas tout sur le même switch).

## 54. Scénario B — Configs clés

