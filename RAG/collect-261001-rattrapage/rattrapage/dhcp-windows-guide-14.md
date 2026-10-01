---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-14
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-20", "2026-09-27"]
keywords: ["agent", "ethernet", "incident"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [2256, 2430]
sha256: f1192089bdfd5dd4e0510a75c9b544e1e17a2b25b5fbeb796d71890e48b3aa70
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

- [ ] Serveur dédié (ou cohabitation documentée : DC+DNS+DHCP acceptable en PME, jamais avec un serveur applicatif exposé)
- [ ] IP statique, hors plage DHCP, documentée
- [ ] Autorisation AD vérifiée (`Get-DhcpServerInDC`)
- [ ] Compte DNS dédié, mot de passe en coffre (section 71)
- [ ] DHCP snooping actif sur tous les VLANs utilisateurs
- [ ] Sauvegardes chiffrées/protégées, testées (restauration annuelle en labo)
- [ ] Journaux d'audit activés et centralisés (SIEM / syslog)
- [ ] Secret de basculement robuste, rotation annuelle

## 97. Détecter un rogue DHCP en continu (supervision)

```powershell
# Script de détection : interroge le réseau et compare les serveurs DHCP qui répondent
# Méthode simple : depuis une sonde, analyser les OFFER reçus
# (nécessite un outil type dhcpdump ou un script Python avec scapy — exemple de principe :)

# Alternative sans agent : surveiller les logs d'audit pour des OFFER multiples
# + sonde Nmap périodique depuis un poste de supervision :
# nmap --script broadcast-dhcp-discover -e Ethernet
# Comparer les "Server Identifier" répondants avec la liste autorisée (192.0.2.10, 192.0.2.11)

# Alerte Zabbix (UserParameter exemple) : compte les serveurs DHCP vus
# UserParameter=dhcp.rogue[*],powershell -File C:\Admin\Detect-RogueDhcp.ps1
```

> 💡 En pratique : une **sonde Nmap planifiée** (`broadcast-dhcp-discover`) par VLAN + alerte si un `Server Identifier` inconnu répond. Simple, efficace, sans agent.

## 98. Réagir à un incident rogue DHCP (procédure)

1. **Confirmer** : `ipconfig /all` sur un poste affecté → noter l'IP du serveur DHCP illégitime.
2. **Localiser** : `arp -a` → MAC → table MAC du switch → port.
3. **Isoler** : shutdown le port du switch (pas juste débrancher — garder la preuve).
4. **Assainir** : `ipconfig /release` + `/renew` sur les postes affectés (ou attendre le renouvellement T1).
5. **Chercher la cause** : équipement personnel ? box 4G ? boucle réseau ? malveillance ?
6. **Durcir** : vérifier le DHCP snooping sur le VLAN concerné, revoir la sensibilisation.
7. **Tracer** : consigner dans le registre des incidents (qui, quoi, quand, port, MAC).

## 99. DHCP et conformité : traçabilité des attributions

En cas d'incident de sécurité, on doit pouvoir dire **qui avait quelle IP à quel moment** :

```powershell
# Rechercher dans les logs d'audit archivés (exemple : IP suspecte le 2026-09-20)
$logs = Get-ChildItem "D:\Backup\DHCP\Logs\DhcpSrvLog-*.log"
foreach ($log in $logs) {
    Import-Csv $log.FullName -Header ID,Date,Time,Description,IP,HostName,MAC |
        Where-Object { $_.IP -eq "192.0.2.87" -and $_.ID -in 10,11 } |
        Select-Object Date, Time, HostName, MAC, Description
}
```

> ⚠️ **Conservation** : les logs d'audit DHCP doivent être archivés (au moins 1 an en entreprise, selon ta politique). Les fichiers natifs tournent sur 7 jours → les copier vers un archivage (script section 73 à étendre aux logs).

## 100. Bonnes pratiques — Sécurité (récap)

1. ✅ Autorisation AD systématique.
2. ✅ DHCP snooping sur **tous** les VLANs utilisateurs + rate limiting.
3. ✅ 802.1X à terme (feuille de route).
4. ✅ Sonde anti-rogue par VLAN.
5. ✅ Durcissement du serveur (section 96).
6. ✅ Logs archivés 1 an minimum.
7. ❌ Ne jamais brancher d'équipement réseau non géré (box, routeur domestique) sur le LAN.
8. ❌ Ne pas se reposer uniquement sur l'autorisation AD.

---

# Bloc I — Bonnes pratiques d'exploitation

## 101. Dimensionnement : combien d'adresses prévoir ?

Règle : **nombre d'équipements simultanés × 1,3** (30 % de marge), en tenant compte de la durée du bail :

| Profil | Équipements | Bail | Adresses à prévoir |
|--------|-------------|------|-------------------|
| Bureau 100 postes fixes | 100 | 8 j | 130 (/24 suffit) |
| Open space + portables 200 | 200 | 1 j | 260 → /23 ou 2×/24 |
| Wi-Fi visiteurs (événementiel) | 500 simultanés | 4 h | 650 → /23 |
| Site industriel / IoT | 300 capteurs | 30 j | 390 → /23 |

```powershell
# Calcul rapide : taille de plage nécessaire
$equipements = 200
$marge = 1.3
$adresses = [math]::Ceiling($equipements * $marge)
Write-Output "Prévoir $adresses adresses → $(if ($adresses -le 254) { '/24 suffit' } else { 'prévoir /23 ou plus' })"
```

> 💡 Le bail court **compense** partiellement le manque d'adresses (rotation rapide), mais ne remplace pas un dimensionnement correct.

## 102. Durées de bail : synthèse décisionnelle

| | Bail long (8–30 j) | Bail court (1–4 h) |
|---|---|---|
| Trafic DHCP | Faible | Plus élevé (renouvellements fréquents) |
| Récupération d'adresses | Lente | Rapide |
| Sensibilité aux pannes DHCP | Faible (les clients gardent leur IP) | Forte (renouvellement fréquent) |
| Idéal pour | Postes fixes, IoT, MFP | Wi-Fi invités, événementiel |

**Règle d'or** : plus le bail est court, plus la **redondance DHCP** (failover) devient critique.

## 103. Redondance : l'architecture minimale viable

```text
Minimum syndical en production :
  2 serveurs DHCP en Load Balance 50/50 (ou Hot Standby)
  + IP helpers vers les 2 sur chaque VLAN
  + Sauvegarde quotidienne testée
  + Supervision (taux d'utilisation + état du failover)

Niveau "robuste" :
  + DHCP snooping partout
  + DNS dynamique avec compte dédié
  + Documentation du plan d'adressage versionnée
  + Procédure de migration/restauration testée en labo 1×/an
```

## 104. Documentation : le plan d'adressage versionné

Modèle de document (à versionner en Git ou SharePoint) :

```markdown
# Plan d'adressage — Site principal (v2.3 — 2026-09-27)

## VLANs
| VLAN | Nom | Sous-réseau | Passerelle | DHCP (étendue) | Bail |
|------|-----|-------------|------------|----------------|------|
| 10 | LAN-Bureaux | 192.0.2.0/24 | 192.0.2.1 | .50-.200 | 8 j |

## IP statiques (hors DHCP)
| IP | Équipement | Rôle |
|----|------------|------|
| 192.0.2.1 | Routeur-RDC | Passerelle VLAN 10 |

## Réservations DHCP (MFP)
| IP | MAC | Nom | Localisation |
|----|-----|-----|--------------|
| 192.0.6.10 | 00-1B-A9-3F-2C-7D | MFP-Compta-RDC | Compta RDC |
```

Génération automatique depuis le DHCP :

```powershell
# Exporter le plan d'adressage réel en CSV (référence pour la doc)
Get-DhcpServerv4Scope | ForEach-Object {
    $s = $_
    $stats = Get-DhcpServerv4ScopeStatistics -ScopeId $s.ScopeId
    [PSCustomObject]@{
        Nom         = $s.Name
        SousReseau  = $s.ScopeId.IPAddressToString
        Plage       = "$($s.StartRange) - $($s.EndRange)"
        Masque      = $s.SubnetMask.IPAddressToString
        Bail        = $s.LeaseDuration.ToString()
        État        = $s.State
        Utilisation = "$([math]::Round($stats.PercentageInUse,1)) %"
    }
} | Export-Csv "C:\Admin\plan-adressage.csv" -NoTypeInformation -Encoding UTF8

# Exporter les réservations (inventaire MFP)
Get-DhcpServerv4Scope | ForEach-Object {
    Get-DhcpServerv4Reservation -ScopeId $_.ScopeId
} | Select-Object ScopeId, IPAddress, Name, ClientId, Description |
    Export-Csv "C:\Admin\reservations.csv" -NoTypeInformation -Encoding UTF8
```

## 105. Maintenance préventive : calendrier type

| Fréquence | Action |
|-----------|--------|
| Quotidien | Sauvegarde auto (script section 73), supervision des alertes |
| Hebdomadaire | Nettoyage des baux expirés (section 28), revue du taux d'utilisation |
| Mensuel | Revue des réservations (MFP remplacés ?), audit des filtres MAC |
| Trimestriel | Test de bascule du failover (couper le primaire 5 min en heure creuse) |
| Semestriel | Revue du plan d'adressage, rotation du secret de basculement |
| Annuel | Test de restauration en labo, revue des durées de bail, audit de sécurité |

