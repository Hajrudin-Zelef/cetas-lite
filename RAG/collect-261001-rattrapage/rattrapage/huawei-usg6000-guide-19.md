---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-19
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["license", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2699, 2819]
sha256: 221abfc0ed18c4fad226f95191f85eef08333e75622032b881ae2394ada0a8f0
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

Paire HA installée il y a 2 ans, jamais basculée → le jour de la panne, le standby a une **config obsolète** ou un **heartbeat mort**. **Prévention** : test semestriel au planning (section 101), supervision du standby (section 106).

## 177. Erreur 9 : les logs qui partent nulle part

Syslog configuré une fois, serveur changé depuis, **personne ne s'en rend compte** → enquête impossible. **Prévention** : supervision de la réception des logs (section 160), test après chaque changement.

## 178. Erreur 10 : le compte admin partagé « équipe-réseau »

Un compte pour 5 personnes → impossible de savoir **qui** a fait quoi, mot de passe qui ne change jamais. **Prévention** : comptes **nominatifs** (section 13), traçabilité (section 138).

## 179. Erreur 11 : la règle « temporaire » de 2019

`permit any any` ajoutée « pour dépanner vite fait » → toujours là 7 ans après. **Prévention** : règle des 90 jours (section 128), time-range avec expiration pour les dépannages.

## 180. Erreur 12 : oublier la zone local

On protège trust/untrust/dmz mais on laisse **untrust→local** ouvert → le firewall lui-même est attaquable (SSH, web). **Prévention** : politiques local explicites dès l'initialisation (section 23).

## 181. Erreur 13 : le certificat SSL VPN expiré

Portail avec certificat expiré un lundi matin → 50 tickets. **Prévention** : supervision de l'expiration (section 112), renouvellement à J-30 (section 66).

## 182. Erreur 14 : la licence UTM expirée en silence

Signatures non mises à jour depuis 4 mois, **personne ne le voit** → protection théorique. **Prévention** : alertes J-60/J-30/J-7 (section 119), contrôle mensuel (section 127).

## 183. Erreur 15 : le firmware d'un autre modèle

« USG6630 et USG6650, c'est presque pareil » → flashé le mauvais firmware → boîtier en rideau. **Prévention** : vérifier la **référence exacte** sur l'étiquette + `display version`, jamais « à peu près » (section 124).

## 184. Erreur 16 : couper l'accès en durcissant la zone local

Nouvelle règle local trop stricte appliquée **à distance** → plus de SSH → déplacement en baie. **Prévention** : tester depuis une session existante, garder la **console à portée**, prévoir une règle de secours (section 23).

## 185. Erreur 17 : le DNS oublié dans les politiques

Tout est autorisé sauf... le DNS → « Internet ne marche pas » alors que le réseau est OK. **Prévention** : la règle DNS sortante fait partie du **socle** (section 35), tester par IP **et** par nom (section 156).

---
---

# BLOC Q — PENSE-BÊTE, GLOSSAIRE, QUIZ, POUR ALLER PLUS LOIN

## 186. Pense-bête de poche : une page à imprimer

```
┌─ USG6000 PENSE-BÊTE ─────────────────────────────────┐
│ CONSOLE : 9600-8-N-1                                 │
│ ZONES   : local(100) trust(85) dmz(50) untrust(5)     │
│ ORDRE PAQUET : DNAT → routage → politique → SNAT → UTM│
│ POLITIQUE : 1ère règle qui matche gagne ; défaut=deny │
│ NAT : exemption no-nat AVANT easy-ip (VPN !)         │
│ VPN : ike sa (ph1) → ipsec sa (ph2) ; DPD activé     │
│ UTM : dimensionner sur débit UTM, pas firewall       │
│ HA : display hrp state verbose ; tester 2×/an        │
│ SAUVEGARDE : save + export ; 3 générations           │
│ DEBUG : debugging ... puis TOUJOURS undo debugging all│
│ URGENCE : console + sauvegarde + rollback écrit      │
└──────────────────────────────────────────────────────┘
```

## 187. Les 15 commandes du quotidien

| # | Commande | Usage |
|---|----------|-------|
| 1 | `display version` | Version logicielle |
| 2 | `display current-configuration` | Config complète |
| 3 | `display zone` | Zones ↔ interfaces |
| 4 | `display security-policy rule all` | Politiques (ordre réel) |
| 5 | `display nat-policy rule all` | NAT (ordre réel) |
| 6 | `display firewall session table` | Sessions live |
| 7 | `display ip routing-table` | Routage |
| 8 | `display cpu` / `display memory` | Ressources |
| 9 | `display ike sa` / `display ipsec sa` | VPN |
| 10 | `display hrp state verbose` | HA |
| 11 | `display license` | Licences |
| 12 | `display logbuffer` | Logs récents |
| 13 | `display interface brief` | État interfaces |
| 14 | `display alarm` | Alarmes matérielles |
| 15 | `save` | Sauvegarder (le plus important) |

## 188. Glossaire

| Terme | Définition |
|-------|------------|
| **ALG** | Application Level Gateway — module qui comprend un protocole (FTP, SIP) pour le faire passer à travers le NAT. |
| **ASPF** | Application Specific Packet Filter — inspection applicative stateful (ex. FTP multi-canaux). |
| **DNAT** | Destination NAT — traduit l'adresse de destination (publier un serveur). |
| **DPD** | Dead Peer Detection — détecte qu'un pair VPN ne répond plus pour reconstruire le tunnel. |
| **Easy IP** | NAT de source utilisant l'IP de l'interface de sortie. |
| **EICAR** | Fichier test standard inoffensif détecté par tous les antivirus. |
| **ESP** | Encapsulating Security Payload — protocole 50, transporte les données chiffrées IPSec. |
| **HRP** | Huawei Redundancy Protocol — synchronisation d'état entre 2 USG en HA. |
| **IKE** | Internet Key Exchange — négocie les clés IPSec (phase 1 : IKE SA, phase 2 : IPSec SA). |
| **IPS** | Intrusion Prevention System — bloque les exploits connus par signatures. |
| **NAT-T** | NAT Traversal — encapsule ESP dans UDP 4500 pour traverser un NAT. |
| **PFS** | Perfect Forward Secrecy — renouvelle les clés à chaque phase 2 (compromis sécu/perf). |
| **PSK** | Pre-Shared Key — secret partagé pour l'authentification IKE. |
| **Server-map** | Table des mappings statiques (DNAT) de l'USG. |
| **SNAT** | Source NAT — traduit l'adresse source (accès Internet du LAN). |
| **UTM** | Unified Threat Management — AV+IPS+URL+applicatif+anti-spam dans un seul boîtier. |
| **VGMP** | VRRP Group Management Protocol — bascule groupée des interfaces en HA. |
| **VRP** | Versatile Routing Platform — système d'exploitation Huawei (CLI `system-view`). |
| **Vsys** | Virtual System — découpage d'un USG en firewalls virtuels indépendants. |

## 189. Quiz : 10 questions pour valider (réponses en fin de section)

**Q1.** Quelles sont les 4 zones prédéfinies de l'USG6000 et leurs priorités par défaut ?
**Q2.** Une règle `deny` placée APRÈS un `permit any any` pour le même flux sera-t-elle appliquée ? Pourquoi ?
**Q3.** Pourquoi faut-il une règle d'exemption NAT (`no-nat`) pour le trafic VPN, et où la placer ?
**Q4.** `display ike sa` montre une SA en état RD mais `display ipsec sa` est vide : où est le problème ?
**Q5.** Citez 3 raisons pour lesquelles un débit peut s'effondrer après activation de l'UTM.
**Q6.** En HA, que synchronise HRP et que NE synchronise-t-il PAS ?
**Q7.** Un utilisateur se plaint qu'un site légitime est bloqué par l'IPS après une MAJ de signatures : quelle est la bonne réaction ?
**Q8.** Pourquoi l'inspection SSL/TLS exige-t-elle des précautions juridiques et techniques ?
**Q9.** Quelle est la première chose à vérifier quand « l'interface est UP mais rien ne passe » ?
**Q10.** Citez les 5 étapes d'un upgrade firmware sans stress.

<details>
<summary><b>Réponses</b> (cliquer pour déplier)</summary>

