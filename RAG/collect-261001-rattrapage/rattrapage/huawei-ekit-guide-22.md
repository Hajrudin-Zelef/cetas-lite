---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-22
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1853, 1953]
sha256: bb12093e17b58e6dadca5e8ef7b977188633ba90be55c4e1000ff143c40458d6
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- Les switchs eKit annoncent l'**authentification 802.1X et MAC** (fiches S220/S310) : de quoi faire du contrôle d'accès filaire de base avec un serveur RADIUS.
- En pratique PME : le 802.1X complet (certificats, EAP-TLS) est **rarement déployé** sur eKit — c'est un marqueur typique de bascule vers le Datacom (section 81).
- Ce que tu peux faire en eKit : portail captif (invités), PSK par SSID, isolation, MAC filtering basique (faible sécurité, utile en complément).
- Si un client exige du 802.1X **Wi-Fi d'entreprise** : qualifie avec la grille (83) — c'est probablement un besoin Datacom (AirEngine + NCE).

## 172. STP/RSTP : éviter les boucles qui tuent un LAN

- Une **boucle** (deux câbles entre les mêmes switchs, ou un câble bouclé sur lui-même) = tempête de broadcast = site mort en secondes.
- Les AR eKit annoncent **RSTP** (fiches AR180) : laisse-le **activé** partout, ne le désactive jamais « pour aller plus vite ».
- En pratique : les boucles viennent d'un câble branché « pour tester » entre deux prises murales, ou d'un petit switch de bureau rebouclé. L'étiquetage (71) et les ports inutilisés désactivés (155) préviennent 90 % des cas.
- Symptôme typique : tout le site rame d'un coup, LED des switchs qui clignotent frénétiquement à l'unisson → cherche la boucle physique **avant** de toucher à la config.

## 173. Multicast / IGMP snooping : pour la TV sur IP

- TV IP d'hôtel (VLAN 70) = **multicast** : sans IGMP snooping, chaque flux TV est diffusé à tous les ports (le switch devient un hub).
- Vérifie que l'**IGMP snooping** est actif sur les switchs du VLAN TV (**à vérifier sur la documentation officielle** par modèle).
- Un VLAN TV **séparé** du VLAN chambres : un client ne doit ni voir ni perturber les flux TV.
- Dimensionnement : ~8-15 Mbit/s par flux HD (ordre de grandeur) × nombre de chaînes simultanées — à intégrer dans le budget de la dorsale.

## 174. NTP : l'heure juste, partout

- Des équipements **pas à l'heure** = logs inexploitables, certificats qui échouent, portail captif qui coince (cas n°7).
- Configure un **serveur NTP** sur l'AR (celui de l'opérateur ou un public — ex. fictif : `pool.ntp.org`) et propage-le en DHCP (option 42) aux clients qui l'acceptent.
- Vérifie l'heure et le **fuseau horaire** de chaque site dans le cloud (un site à cheval sur deux fuseaux = des alertes horodatées bizarres).
- À chaque dépannage : commence par vérifier l'heure des équipements — 5 % des « bugs mystérieux » viennent de là.

## 175. Syslog et journaux : la mémoire du réseau

- Les switchs eKit supportent **Syslog** (écosystème cité dans les fiches) : envoie les logs vers un serveur central (même un simple syslog sur ton NMS).
- En PME : au minimum, conserve les logs de la **passerelle/USG** (connexions, VPN, alertes sécurité) — durée selon la réglementation locale.
- Les logs sont ta **boîte noire** : quand le client dit « c'est tombé mardi », tu dois pouvoir dire « mardi 14:32, perte du lien WAN pendant 12 minutes ».
- Sans serveur syslog : exporte les logs avant chaque MAJ et à chaque incident majeur — un log effacé par un reboot est une preuve perdue.

## 176. Gestion des incidents : le processus (même à 3 personnes)

1. **Détection** : alerte cloud, appel client, rituel hebdo.
2. **Qualification** (5 min) : périmètre (qui est impacté ?), criticité (tout le site ? la caisse ?), workaround possible ?
3. **Ticket** : même simple (date, client, symptôme, criticité) — sans ticket, pas de suivi, pas de KPI.
4. **Diagnostic** : méthode section 91, une hypothèse à la fois.
5. **Résolution + vérification** : le client confirme que c'est bon (pas toi seul).
6. **Clôture** : cause racine écrite, fiche site mise à jour si besoin, cas ajouté au guide si nouveau.
7. **Retour d'expérience** (mensuel, 30 min en équipe) : qu'est-ce qui a cassé ce mois-ci ? Qu'est-ce qu'on change dans nos procédures ?
Un processus écrit, même d'une page, bat l'héroïsme désorganisé à tous les coups.

## 177. Astreinte : l'organiser sans y laisser sa santé

- **Plages** : définis les heures couvertes (ex. fictif : 7 h-22 h, 7 j/7 pour l'hôtellerie ; heures ouvrées pour les bureaux) — écrites au contrat.
- **Rotation** : jamais la même personne 2 semaines d'affilée. Un technicien épuisé fait des bêtises à 3 h du matin.
- **Escalade** : niveau 1 (technicien : cycle PoE, redémarrage, lecture cloud) → niveau 2 (toi : config, RMA, distributeur) → niveau 3 (distributeur/support Huawei).
- **Kit d'astreinte** : téléphone chargé, accès cloud, coffre, contacts (opérateurs, électricien, distributeur), équipement de secours pour les sites critiques.
- **Facturation** : interventions hors contrat = facturées (tarif écrit à l'avance). La gratuité systématique tue ton business.

## 178. SLA : les écrire pour de vrai

| Niveau | Exemple d'engagement (fictif, à adapter) | Pour qui |
|---|---|---|
| Bronze | Intervention sous 48 h ouvrées, sans astreinte | Bureau non critique |
| Silver | Sous 8 h ouvrées + astreinte téléphonique | Commerce, PME standard |
| Gold | Sous 4 h, 7 j/7 + équipement de secours sur site | Hôtel, clinique, industrie |

- Un SLA se **mesure** (KPI section 80) et se **rapporte** (rapport mensuel 1 page).
- Ne promets jamais un SLA que ton stock (105) et ton équipe ne peuvent pas tenir.
- Pénalités : à manier avec prudence — propose plutôt des **avoirs** en cas de non-respect, c'est plus sain commercialement.

## 179. Reporting client : la page mensuelle qui fidélise

```
RAPPORT MENSUEL - [Client - Site] - [Mois]
Disponibilite : 99,7 % (1 coupure WAN operateur, 22 min, hors notre perimetre)
Incidents : 2 (AP chambres debranche par le menage -> rebranche ; PSK invites tourne)
Interventions : 1 sur site, 3 a distance
Firmwares : a jour (versions : ...)
Alertes traitees : 5 (detail en annexe)
Recommandations : ajouter 1 AP aile ouest (densite en hausse), prevoir remplacement
                  onduleur (batteries 3 ans)
Prochain audit : [date]
```
Un client qui lit ça chaque mois **ne te quitte pas** — et il accepte tes recommandations parce qu'elles sont écrites et chiffrées.

## 180. Gérer un parc multi-clients : l'organisation du prestataire

- **1 tenant par client** (ou tenant prestataire bien cloisonné — section 32) : jamais de mélange.
- **Nommage global** : `[CLIENT]-[SITE]-[EQUIPEMENT]` partout (cloud, fiches, coffre, tickets) — quand tu gères 200 sites, la rigueur de nommage est ta mémoire externe.
- **Modèles de config** par vertical (boutique, hôtel, bureau) : 80 % commun, 20 % spécifique — c'est ça, l'industrialisation.
- **Calendrier partagé** : MAJ planifiées, audits, renouvellements de contrats — un oubli de MAJ sur 30 sites, c'est 30 sites vulnérables.
- **Facturation** : aligne la facturation maintenance sur les audits (le rapport mensuel justifie la facture).

## 181. Documentation d'exploitation : le dossier client type

```
DOSSIER CLIENT - [Client]
  01_Contrat/ (contrat, SLA, PV de recette)
  02_Sites/ (par site : fiche site, plan, photos, sauvegardes)
  03_Acces/ (references au coffre, registre des acces)
  04_Incidents/ (tickets, rapports mensuels)
  05_Maintenance/ (audits, MAJ, RMA)
```
5 dossiers, toujours les mêmes, pour tous les clients. Un nouveau technicien doit pouvoir prendre un dossier et intervenir **sans t'appeler** — c'est le test.

## 182. Passation entre techniciens (et le jour où tu pars en vacances)

