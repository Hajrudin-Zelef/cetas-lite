---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-16
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr", "attribution"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [2551, 2641]
sha256: 1a4ce328b8087cca8499b8ec01c5ccd7eeb3596f5acce84874107cf0afe5db19
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

```text
┌─────────────────────────────────────────────────────────────┐
│              PENSE-BÊTE DHCP — Windows Server                │
├─────────────────────────────────────────────────────────────┤
│ DORA : Discover → Offer → Request → Ack (UDP 67/68)         │
│ Renouvellement : T1 = 50 % du bail (unicast)                │
│                  T2 = 87,5 % (broadcast)                    │
├─────────────────────────────────────────────────────────────┤
│ Options : 003 passerelle │ 006 DNS │ 015 domaine            │
│           044/046 WINS   │ 066/067 PXE │ 042 NTP            │
├─────────────────────────────────────────────────────────────┤
│ Précédence : Serveur < Étendue < Réservation < Stratégie    │
├─────────────────────────────────────────────────────────────┤
│ Baux : 8 j fixe │ 1 j Wi-Fi │ 4 h invités │ 30 j IoT/MFP    │
├─────────────────────────────────────────────────────────────┤
│ Alertes : étendue > 80 % │ failover ≠ Normal │ event 1020   │
├─────────────────────────────────────────────────────────────┤
│ Commandes :                                                  │
│  Get-DhcpServerv4Scope                  (étendues)          │
│  Get-DhcpServerv4ScopeStatistics        (utilisation)       │
│  Get-DhcpServerv4Lease -ScopeId <id>    (baux)              │
│  Get-DhcpServerv4Reservation            (réservations)      │
│  Get-DhcpServerv4Failover               (basculement)       │
│  Export-DhcpServer -File <xml> -Leases  (sauvegarde)        │
├─────────────────────────────────────────────────────────────┤
│ Client : ipconfig /release │ /renew │ /all │ /setclassid    │
│ Rogue ? → ipconfig /all (Serveur DHCP) → arp -a → switch    │
├─────────────────────────────────────────────────────────────┤
│ APIPA 169.254.x.x = aucun OFFER reçu → voir arbre §85       │
└─────────────────────────────────────────────────────────────┘
```

## A3. Glossaire complet

| Terme | Définition |
|-------|------------|
| ACK (DHCPACK) | Message de confirmation finale du serveur |
| Agent de relais | Équipement/service qui relaie les broadcasts DHCP entre sous-réseaux (champ giaddr) |
| APIPA | Adressage automatique 169.254.0.0/16 quand aucun DHCP ne répond |
| Autorisation AD | Enregistrement du serveur DHCP dans Active Directory (obligatoire en domaine) |
| Bail (lease) | Attribution temporaire d'une IP à un client |
| Basculement (failover) | Réplication temps réel entre 2 serveurs DHCP |
| Classe fournisseur | Identifiant déclaré par le client (option 60, ex. : PXEClient) |
| Classe utilisateur | Étiquette configurée côté client (option 77) pour les stratégies |
| DHCID | Enregistrement DNS lié à la protection de nom |
| DHCP snooping | Fonction switch qui filtre les messages DHCP par port (trusted/untrusted) |
| DISCOVER | 1er message du client (broadcast : "qui peut me donner une IP ?") |
| DNS dynamique | Mise à jour automatique des zones DNS par le DHCP |
| DORA | Discover, Offer, Request, Acknowledge — le cycle DHCP |
| Étendue (scope) | Plage d'adresses distribuée sur un sous-réseau |
| Exclusion | Plage jamais attribuée par le DHCP |
| giaddr | Champ du paquet DHCP indiquant l'agent de relais (donc le sous-réseau d'origine) |
| Hot Standby | Mode de basculement actif/passif |
| IP helper-address | Commande routeur/switch désignant les serveurs DHCP (agent de relais) |
| Load Balance | Mode de basculement actif/actif avec répartition de charge |
| Multicast scope | Étendue d'adresses multicast (224.0.0.0/4) |
| NAK (DHCPNAK) | Refus du serveur (le client doit recommencer un DORA) |
| OFFER (DHCPOFFER) | Proposition d'adresse par le serveur |
| Option DHCP | Paramètre réseau transmis avec le bail (numéroté, RFC 2132) |
| Protection de nom | Mécanisme anti-écrasement des enregistrements DNS (DHCID) |
| Réconciliation | Réparation des incohérences entre la base des baux et le registre |
| Réservation | Attribution permanente d'une IP à une adresse MAC |
| REQUEST | 3e message : le client accepte (ou renouvelle) |
| Rogue DHCP | Serveur DHCP illégitime |
| Scavenging | Nettoyage automatique des enregistrements DNS périmés |
| Starvation | Attaque par épuisement de l'étendue (MAC forgées) |
| Stratégie | Règle d'attribution fine (MAC, classe, plage) au sein d'une étendue |
| Superscope | Regroupement de plusieurs étendues sur un même segment |
| T1 / T2 | Timers de renouvellement (50 % / 87,5 % du bail) |
| WPAD | Web Proxy Auto-Discovery (option 252) |

## A4. Quiz — 10 questions + réponses

**Q1. Quelles sont les 4 étapes du cycle DORA et qui émet chaque message ?**
> R : 1. DISCOVER (client, broadcast) — 2. OFFER (serveur) — 3. REQUEST (client, broadcast) — 4. ACK (serveur). Le REQUEST est en broadcast pour prévenir les autres serveurs ayant fait une offre.

**Q2. Pourquoi un serveur DHCP Windows non autorisé dans AD ne distribue-t-il pas d'adresses ?**
> R : Protection anti-rogue intégrée : au démarrage, le service vérifie son autorisation dans la partition Configuration de l'AD. Non autorisé → il reste démarré mais ignore les DISCOVER (événement 1046).

**Q3. Un client en Wi-Fi invités obtient une IP en 169.254.x.x. Citez 3 causes possibles dans l'ordre de vérification.**
> R : 1. Étendue pleine (bail trop long pour la rotation des visiteurs) — 2. Agent de relais (IP helper) manquant ou mal configuré sur le VLAN invités — 3. Service DHCP arrêté / étendue désactivée. (Voir arbre de décision section 85.)

**Q4. Quelle est la précédence des options DHCP entre les niveaux serveur, étendue, réservation et stratégie ?**
> R : Serveur < Étendue < Réservation < Stratégie. Le niveau le plus spécifique l'emporte toujours.

**Q5. Quelle différence entre Hot Standby et Load Balance pour le basculement DHCP ?**
> R : Hot Standby = actif/passif (le secondaire ne sert qu'en cas de panne du primaire, avec un % d'adresses réservées) ; Load Balance = actif/actif (les deux servent en permanence, charge répartie selon le ratio configuré, 50/50 par défaut).

**Q6. Pourquoi faut-il déclarer les DEUX serveurs DHCP dans l'ip helper-address ?**
> R : Si seul le primaire est déclaré et qu'il tombe, le relais n'envoie les DISCOVER qu'à un serveur mort → le VLAN n'a plus de DHCP malgré le basculement configuré côté serveurs.

