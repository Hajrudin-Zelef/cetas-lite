---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-26
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [2251, 2357]
sha256: 94fe6b575c29f146912fb61ac4f9642edd7a906d7b0baa0edee6657bc5af356f
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- **Profil** : curiosité + rigueur > diplômes. Un technicien qui documente tout et teste méthodiquement vaut mieux qu'un « expert » qui improvise.
- **Parcours** : maquette (210) → binôme sur site (3-4 déploiements) → site simple seul (boutique) → multi-sites → astreinte.
- **Ce guide** = son manuel : fais-le lire (au moins les sections 16-30, 70-75, 91-100) avant le premier site seul.
- **Erreur à laisser faire** : un technicien qui n'a jamais planté un site ne sait pas dépanner. Laisse-le se tromper **sur la maquette**, pas chez le client — et débriefe chaque incident (176).

## 212. Sous-traitants : les cadrer pour garder la qualité

- **Ce que tu sous-traites** : câblage, fibre, électricité — **jamais** la configuration sans supervision (c'est ton savoir-faire et ta responsabilité).
- **Contrat écrit** : périmètre, délais, réception (187), garantie de la pose, pénalités de retard.
- **Réception systématique** : tu testes le câblage **avant** de payer le solde — sans exception.
- **Ne sous-traite jamais** à quelqu'un moins rigoureux que toi sur la documentation : un câbleur qui n'étiquette pas te coûte 3× son prix en dépannage futur.

## 213. Qualité : les 5 indicateurs de ton service

1. **Taux de sites « verts »** (cloud) : objectif > 98 % en heures ouvrées.
2. **Délai moyen de résolution** par criticité (vs SLA, section 178).
3. **Retours sur site** dans les 30 jours après installation : objectif < 5 % (un retour = un défaut de recette).
4. **Documentation** : 100 % des sites avec fiche à jour (audit trimestriel).
5. **Satisfaction client** : un appel ou un mini-questionnaire 1 mois après chaque déploiement — 5 minutes qui valent de l'or commercial.
Mesure-les, affiche-les en interne, améliore le plus mauvais chaque trimestre.

## 214. Litiges : les éviter, les gérer

- **Prévention** : devis détaillé signé, PV de recette (160), contrat de maintenance écrit, tout par écrit (les promesses orales n'existent pas).
- **En cas de litige** : dossier complet (181) + fiches d'intervention (186) + PV signé = ta défense. Sans écrits, c'est parole contre parole.
- **Négociation** : propose toujours une solution technique avant de parler argent (un AP ajouté gracieusement coûte moins cher qu'un procès et sauve la relation).
- **Ne jamais** : couper un service pour faire pression (illégal et suicidaire commercialement), ni insulter qui que ce soit par écrit.
- **Assurance** : RC pro à jour, avec une couverture adaptée à ton activité (un dégât des eaux causé par ta baie, ça arrive).

## 215. Cap sur l'autonomie : ta feuille de route 12 mois

- **Mois 1-2** : maquette montée (210), 1er site pilote eKit, ce guide lu par toute l'équipe.
- **Mois 3-4** : 3-5 sites en production, rituel cloud hebdo en place (43), modèles de config par vertical.
- **Mois 6** : 10+ sites, premier rapport mensuel client (179), stock de secours constitué (105), 1er audit annuel planifié.
- **Mois 9** : NMS complémentaire en place (79), SLA proposés (178), formation équipe terminée (158).
- **Mois 12** : revue annuelle — KPI (213), mise à jour de ce guide (sections « à vérifier » levées), roadmap produits avec le distributeur (184), objectif de croissance fixé.
Dans 12 mois, tu ne « poses plus du Wi-Fi » : tu exploites un parc. C'est là que la marge et la réputation se construisent.

---

## 216. Topologie type — site industriel léger / atelier de production

```
   [Internet Fibre]                [4G secours - routeur dedie]
          |                                   |
   +------+----------------+------------------+------+
   |              [USG6000F-S125]                    |
   |         (pare-feu + VPN siege)                 |
   +------+----------------+------------------+------+
          | (trunk, fibre si distance)
   +------+------+                                  (bureaux)
   | [S310-24P4X] |                         +-------+--------+
   |  (atelier)  |                         | [S220-8P4S]    |
   +--+--+--+---+                         |  (bureaux)   |
      |  |  |                             +--+--+---+----+
   [AP761] [Cameras] [Automates*]            |  |   |
   (quais,  (VLAN 40)  (VLAN 50,            [AP361] [PC]
    exterieur)          isoles)             bureaux

VLAN 10 BUREAUX / 20 INVITES / 40 CAMERAS / 50 AUTOMATES / 60 VOIX
* Automates : coordonner avec l'integrateur industriel (jamais touche sans lui).
Regles : poussiere = baie filtree ; vibrations = fixation renforcee ;
         coupures elec frequentes = onduleur + redemarrage auto ; ATEX = voir §144.
```

## 217. Tableau — ce que chaque usage consomme (pour dimensionner)

| Usage | Débit indicatif par flux (ordre de grandeur) | Sensible à |
|---|---|---|
| Navigation web / mail | 1-5 Mbit/s | Latence modérée |
| Visio HD (1 participant) | 2-4 Mbit/s | Latence + gigue (jitter) |
| Visio 4K / salle immersive | 15-25 Mbit/s | Latence + gigue |
| Streaming vidéo HD | 5-8 Mbit/s | Débit descendant |
| Streaming 4K | 20-30 Mbit/s | Débit descendant |
| VoIP (1 appel) | 0,1 Mbit/s | Latence + gigue (prioritaire !) |
| Sauvegarde réseau (1 poste) | Variable (pics) | Ne pas saturer l'uplink en journée |
| Caméra IP 1080p (1 flux) | 4-8 Mbit/s | Débit montant vers NVR |
| Caméra IP 4K (1 flux) | 12-20 Mbit/s | Débit montant vers NVR |
| Mise à jour OS (1 PC) | Pics à 50-100 Mbit/s | Planifier hors heures de pointe |

**Usage** : multiplie par le nombre de flux simultanés pour dimensionner l'uplink et les dorsales. Une salle de réunion 10 personnes en visio = ~30 Mbit/s réservés.

## 218. Glossaire complémentaire (20 termes)

- **Airtime** : temps d'antenne — la ressource radio partagée. Chaque SSID/client en consomme.
- **Beacon** : trame de signalisation émise par chaque SSID (~10/s) — d'où la règle des 3 SSID max.
- **DFS** : canaux 5 GHz partagés avec les radars — l'AP doit les libérer si radar détecté (micro-coupures possibles).
- **Tx Power** : puissance d'émission — baisser > augmenter en dense.
- **RSSl / SNR** : niveau de signal / rapport signal-bruit — le SNR fait la qualité, pas le nombre de barres.
- **Co-channel interference** : deux AP sur le même canal qui se gênent — d'où l'alternance 1/6/11.
- **Hidden node** : deux clients qui ne s'entendent pas mais brouillent l'AP — typique en entrepôt.
- **Jitter** : variation de la latence — tue la voix et la visio.
- **MOS** : note de qualité vocale (1-5) — vise > 4 en téléphonie IP.
- **LLDP** : protocole de découverte de voisinage — utile pour voir « qui est branché où ».
- **LACP (802.3ad)** : agrégation de liens — à vérifier par modèle côté eKit.
- **IGMP snooping** : optimisation multicast (TV IP) — section 173.
- **DHCP snooping** : protection contre les serveurs DHCP pirates — annoncé sur S220.
- **Sticky MAC** : mémorisation des MAC par port — sécurité filaire de base.
- **Port isolation** : isolation L2 entre ports — pour les ports invités filaires.
- **Rogue AP** : AP non autorisé (celui du voisin ou celui « branché pour dépanner ») — à détecter et neutraliser.
- **Heatmap** : carte de couverture radio — issue d'un survey, pas d'une promesse.
- **Backhaul** : lien de collecte (filaire idéalement) — un AP en mesh sans fil divise le débit.
- **Failover** : bascule automatique (WAN, VPN) — à tester pour de vrai, pas sur papier.
- **Baseline** : mesures de référence d'un site sain — speedtests, ping, charge — pour comparer quand « c'est lent ».

## 219. La valise du technicien eKit (check-list matériel)

