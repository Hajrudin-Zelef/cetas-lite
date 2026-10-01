---
id: collect-261001-rattrapage/rattrapage/dhcp-windows-guide-15
title: "Guide technique ultra-complet : DHCP sous Windows Server en entreprise"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/dhcp_windows_guide.md
source_anchor: ""
source_lines: [2431, 2550]
sha256: bcdbbcb9044e296b0737d2e2801fb55d725e547f2f09dd975ddce49e8ca5d381
---

# Guide technique ultra-complet : DHCP sous Windows Server en entreprise

```powershell
# Test de bascule trimestriel (procédure) :
# 1. Sur le primaire : Stop-Service DHCPServer
# 2. Depuis un client de chaque VLAN : ipconfig /release ; ipconfig /renew
# 3. Vérifier que le secondaire répond (ipconfig /all -> Serveur DHCP = .11)
# 4. Redémarrer le primaire : Start-Service DHCPServer
# 5. Vérifier l'état du failover : State = Normal
```

## 106. Supervision : quoi monitorer (Zabbix / autre)

| Sonde | Seuil | Moyen |
|-------|-------|-------|
| Service DHCPServer démarré | — | Agent Zabbix / WinRM |
| Taux d'utilisation par étendue | Alerte 80 %, critique 90 % | Script PowerShell (section 24) |
| État du failover = Normal | Alerte si ≠ Normal | `Get-DhcpServerv4Failover` |
| Événement 1020 (épuisement) | Alerte immédiate | Log Windows |
| Événement 1046 (non autorisé) | Alerte immédiate | Log Windows |
| Événement 1056 (échec DNS) | Alerte | Log Windows |
| Sonde anti-rogue par VLAN | Alerte si serveur inconnu | Nmap planifié (section 97) |
| Fraîcheur de la sauvegarde | Alerte si > 26 h | Vérification du fichier |

```powershell
# Script de supervision : retourne les étendues en alerte (format simple pour Zabbix)
$alertes = Get-DhcpServerv4ScopeStatistics | Where-Object { $_.PercentageInUse -gt 80 } |
    Select-Object @{N="Scope";E={$_.ScopeId.IPAddressToString}},
                  @{N="Pct";E={[math]::Round($_.PercentageInUse,1)}}
if ($alertes) { $alertes | ConvertTo-Json } else { Write-Output "OK" }
```

## 107. Gestion des changements : procédure type

Toute modification DHCP en production suit ce cycle :

1. **Demande** : qui, quoi, pourquoi, quel VLAN/étendue.
2. **Analyse d'impact** : combien de clients affectés ? besoin d'heure creuse ?
3. **Sauvegarde** : `Export-DhcpServer` avant toute modif (section 73).
4. **Test** : si possible sur une étendue de labo ou une réservation test.
5. **Application** : en heure creuse pour les changements d'options/passerelle.
6. **Vérification** : `ipconfig /renew` sur un échantillon, logs, supervision.
7. **Documentation** : mise à jour du plan d'adressage (section 104).
8. **Retour arrière** : `Import-DhcpServer` si problème (procédure connue d'avance).

## 108. Pense-bête d'exploitation (commandes du quotidien)

```powershell
# --- Quotidien ---
Get-DhcpServerv4ScopeStatistics | # taux d'utilisation
    Sort-Object PercentageInUse -Descending | Select-Object -First 5
Get-DhcpServerv4Failover | Select-Object Name, State  # état du basculement

# --- Chercher un équipement ---
# Par IP :
Get-DhcpServerv4Lease -ScopeId 192.0.2.0 | Where-Object IPAddress -eq 192.0.2.87
# Par nom :
Get-DhcpServerv4Scope | % { Get-DhcpServerv4Lease -ScopeId $_.ScopeId } |
    Where-Object HostName -like "*MFP*"

# --- Forcer le renouvellement d'un parc (via GPO / script de démarrage) ---
# ipconfig /renew

# --- Sauvegarde manuelle avant une modif ---
Export-DhcpServer -File "D:\Backup\DHCP\avant-modif-$(Get-Date -Format yyyyMMdd-HHmm).xml" -Leases -Force

# --- Redémarrer le service (rarement nécessaire) ---
Restart-Service -Name DHCPServer
```

---

# Annexes

## A1. Checklists

### Checklist — Mise en service d'un serveur DHCP

- [ ] Windows Server 2019/2022/2025 installé, à jour, IP statique configurée
- [ ] Serveur membre du domaine (ou workgroup assumé et documenté)
- [ ] Rôle DHCP installé (`Install-WindowsFeature DHCP -IncludeManagementTools`)
- [ ] Post-configuration effectuée (groupes de sécurité créés)
- [ ] Serveur **autorisé dans AD** (`Get-DhcpServerInDC` le liste)
- [ ] Plan d'adressage écrit et validé (section 13)
- [ ] Étendues créées (une par VLAN), actives
- [ ] Exclusions configurées pour les IP statiques
- [ ] Options : 003/006/015 au minimum, niveaux corrects (section 38)
- [ ] Réservations MFP/imprimantes créées et testées
- [ ] DNS dynamique activé + compte dédié (sections 69-71)
- [ ] Basculement configuré avec le 2e serveur, état Normal
- [ ] IP helpers configurés vers les 2 serveurs sur chaque VLAN
- [ ] DHCP snooping actif sur les switchs (tous les VLANs utilisateurs)
- [ ] Sauvegarde quotidienne planifiée et testée
- [ ] Supervision en place (section 106)
- [ ] Documentation à jour (plan d'adressage versionné)

### Checklist — Audit annuel du DHCP

- [ ] Restauration testée en labo depuis la sauvegarde
- [ ] Test de bascule du failover effectué
- [ ] Rotation du secret partagé de basculement
- [ ] Revue des durées de bail (toujours adaptées ?)
- [ ] Nettoyage des réservations obsolètes (équipements remplacés)
- [ ] Revue des filtres MAC (Allow/Deny toujours pertinents ?)
- [ ] Vérification du DHCP snooping (nouveaux VLANs couverts ?)
- [ ] Rotation du mot de passe du compte DNS dédié
- [ ] Archivage des logs d'audit vérifié
- [ ] Plan d'adressage à jour et versionné

### Checklist — Intervention "plus de DHCP sur un VLAN"

- [ ] Le problème est-il global ou localisé ? (un VLAN ? un client ? tout ?)
- [ ] Service DHCP démarré sur les 2 serveurs ?
- [ ] Étendue active et non pleine ?
- [ ] IP helper présent sur le VLAN ?
- [ ] Pare-feu UDP 67/68 OK entre relais et serveurs ?
- [ ] Failover en état Normal ?
- [ ] Rogue DHCP ? (`ipconfig /all` → Serveur DHCP)
- [ ] Logs : événements 1020 / 1046 / 1034 ?

## A2. Pense-bête de poche

