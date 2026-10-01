---
id: collect-261001-rattrapage/rattrapage/dns-windows-guide-7
title: "DNS sous Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/dns_windows_guide.md
source_anchor: ""
source_lines: [828, 958]
sha256: 5e42e39ee295690c77a3895c7ae3bd8afabad1e7f7588ae78ffa180467db143c
---

# DNS sous Windows Server en entreprise — Guide technique ultra-complet

- **No-refresh interval** (défaut 7 jours) : période pendant laquelle un client *ne peut pas* rafraîchir son timestamp → évite une réplication AD à chaque renouvellement DHCP.
- **Refresh interval** (défaut 7 jours) : période pendant laquelle le client *peut* rafraîchir ; passé ce délai sans refresh, l'enregistrement est considéré comme périmé.
- **Scavenging period** (défaut 7 jours, niveau serveur) : fréquence du passage du ramasse-miettes.

**Règle de cohérence :** `no-refresh + refresh` doit être **supérieur** à la durée de bail DHCP typique, sinon des machines éteintes le week-end (bail 8 jours > cycle) se font scavenger à tort. Avec des baux DHCP de 8 jours, les défauts 7+7 jours sont corrects.

## 54. Activer le vieillissement sur une zone

Le vieillissement s'active en **deux endroits** : sur chaque zone, ET au niveau serveur (§55). L'un sans l'autre = rien ne se passe.

```powershell
# 1. Activer l'aging sur la zone (avec les intervalles)
Set-DnsServerZoneAging -Name "contoso.local" -Aging $true -RefreshInterval 7.00:00:00 -NoRefreshInterval 7.00:00:00 -PassThru
Set-DnsServerZoneAging -Name "1.10.in-addr.arpa" -Aging $true -RefreshInterval 7.00:00:00 -NoRefreshInterval 7.00:00:00

# 2. Vérifier
Get-DnsServerZoneAging -Name "contoso.local" | Format-List *

# 3. Appliquer l'horodatage aux enregistrements EXISTANTS (sinon ils restent "statiques" pour toujours)
dnscmd SRV-DNS-01 /AgeAllRecords contoso.local
```

⚠️ `/AgeAllRecords` : à n'exécuter qu'**une fois**, après avoir vérifié que les enregistrements existants sont légitimes. Il rend éligibles au scavenging des enregistrements qui ne l'étaient pas.

GUI : Propriétés de la zone → onglet **Général** → bouton **Vieillissement** → cocher « Éliminer les enregistrements de ressources périmés ».

## 55. Activer le scavenging au niveau serveur

```powershell
# Activer le scavenging serveur (cycle de 7 jours, valeurs par défaut 7j/7j)
Set-DnsServerScavenging -ScavengingState $true -ScavengingInterval 7.00:00:00 `
    -RefreshInterval 7.00:00:00 -NoRefreshInterval 7.00:00:00 -PassThru

# Vérifier
Get-DnsServerScavenging | Format-List *

# Forcer un passage manuel (avec -Force, sans confirmation)
Start-DnsServerScavenging -Force
```

GUI : Propriétés du serveur → onglet **Avancé** → cocher « Activer le nettoyage automatique des enregistrements périmés » + « Période de nettoyage ».

⚠️ **N'activez le scavenging serveur que sur UN SEUL serveur DNS** (le « scavengeur » désigné, typiquement le premier DC). Si plusieurs serveurs scavengent, ils se marchent dessus et les logs deviennent illisibles. Documentez lequel dans votre dossier d'exploitation.

## 56. Bonnes pratiques et pièges du scavenging

**Bonnes pratiques :**
1. Activez l'aging sur **toutes** les zones dynamiques (directes + inversées), pas seulement `contoso.local`.
2. `no-refresh + refresh` > durée de bail DHCP (ex : baux 8 j → 7+7 j OK ; baux 30 j → passez à 14+14 j).
3. Scavenging serveur sur **un seul** serveur ; cycle ≥ 7 jours.
4. Avant la première activation : inventaire des enregistrements statiques critiques (imprimantes, équipements réseau en IP fixe enregistrés à la main) → vérifiez qu'ils sont bien **statiques** (timestamp à 0) ou protégés.
5. Surveillez l'Event ID **2501** (nombre d'enregistrements scavengés) après chaque cycle les premiers mois.

**Pièges mortels :**
- Scavenging activé avec des baux DHCP longs (ex : 30 jours) et des intervalles 7+7 : des machines en congés se font supprimer → cas n°5 (§104).
- `/AgeAllRecords` lancé sans inventaire : des enregistrements créés à la main deviennent éligibles et disparaissent 14 jours plus tard.
- Enregistrements d'équipements réseau (imprimantes, onduleurs) créés par DHCP puis passés en IP fixe sans passer en statique : le DHCP ne les rafraîchit plus → scavengés.

---

# Chapitre 7 — Redirecteurs et indications de racine

## 57. Redirecteurs : principe

Un **redirecteur** (forwarder) est un serveur DNS vers lequel votre serveur **transfère** les requêtes récursives qu'il ne peut pas résoudre lui-même, au lieu d'interroger la racine. Typiquement : les DNS du FAI, `1.1.1.1`, `9.9.9.9`, ou le DNS central du groupe.

```
Client → SRV-DNS-01 (ne sait pas) → Redirecteur 1.1.1.1 → Internet → réponse → cache → client
```

Avantages : **cache partagé** (le redirecteur a souvent déjà la réponse → plus rapide), **un seul point de sortie** à autoriser au pare-feu, possibilité de filtrage (DNS filtrant type Quad9).

## 58. Indications de racine : principe

Les **indications de racine** (root hints) sont la liste des 13 serveurs racine (`a.root-servers.net` → `m.root-servers.net`, fichier `cache.dns`). Sans redirecteur, votre serveur fait lui-même la résolution **itérative** complète depuis la racine.

```
Client → SRV-DNS-01 → . (racine) → serveur du TLD → serveur autoritaire → réponse
```

Avantages : **aucune dépendance tierce** (pas de FAI, pas de Cloudflare), pas de logging de vos requêtes chez un tiers, conforme à une posture « zéro confiance externe ». Inconvénient : première résolution plus lente, tout le trafic DNS sortant (UDP/TCP 53 vers Internet) doit être autorisé.

## 59. Redirecteurs vs racine : quand utiliser quoi

| Critère | Redirecteurs | Indications de racine |
|---|---|---|
| Performance (cache chaud) | ✅ Meilleure | Moyenne |
| Indépendance / confidentialité | ❌ Dépend d'un tiers | ✅ Aucun tiers |
| Pare-feu sortant | Simple (quelques IP) | Large (tout l'Internet en 53) |
| Filtrage (malware, parental) | ✅ Possible (Quad9, OpenDNS) | ❌ Non |
| Robustesse si le FAI tombe | ❌ Si redirecteur = FAI | ✅ |
| Recommandé quand | Poste client standard, besoin de filtrage | Serveur exposé, contrainte de souveraineté |

**Recommandation entreprise :** redirecteurs vers **2 résolveurs fiables et rapides** (ex : `1.1.1.1` + `9.9.9.9`, ou les DNS de votre opérateur si SLA), avec **repli sur la racine désactivé ou activé selon votre politique** (case « Utiliser les indications de racine si aucun redirecteur n'est disponible » : à cocher = résilience, à décocher = contrôle strict du chemin).

## 60. Configurer les redirecteurs

```powershell
# Définir les redirecteurs (remplace la liste) et désactiver le repli racine
Set-DnsServerForwarder -IPAddress @("1.1.1.1","9.9.9.9") -UseRootHint $false -Timeout 3 -PassThru

# Ajouter un redirecteur sans écraser
Add-DnsServerForwarder -IPAddress "8.8.8.8" -PassThru

# Vérifier + tester chaque redirecteur
Get-DnsServerForwarder | Select-Object -ExpandProperty IPAddress
Test-DnsServer -IPAddress "1.1.1.1" -Context "Forwarder"
```

GUI : Propriétés du serveur → onglet **Redirecteurs** → Modifier. Le bouton **« Utiliser les indications de racine si aucun redirecteur n'est disponible »** correspond à `-UseRootHint`.

⚠️ **Timeout** : 3 secondes par défaut. Avec des redirecteurs lents, les clients subissent des latences en cascade → cas n°8 (§107).

## 61. Redirecteurs conditionnels

Un redirecteur conditionnel ne s'applique qu'à **une zone précise** : « pour `partenaire.local`, interroge toujours ces serveurs-là ».

```powershell
# Tout ce qui concerne partenaire.local va vers leurs DNS
Add-DnsServerConditionalForwarderZone -Name "partenaire.local" -MasterServers @("192.168.50.10","192.168.50.11") -PassThru

# Lister / modifier / supprimer
Get-DnsServerZone -Name "partenaire.local"
Set-DnsServerConditionalForwarderZone -Name "partenaire.local" -MasterServers @("192.168.50.12")
Remove-DnsServerZone -Name "partenaire.local" -Force
```

Cas d'usage rois : interconnexions VPN avec un partenaire, forêt AD distante sans approbation de forêt, zone d'un hébergeur cloud privé. **Priorité** : le redirecteur conditionnel est consulté *avant* les redirecteurs généraux.

## 62. Chaînage et boucles : à éviter

