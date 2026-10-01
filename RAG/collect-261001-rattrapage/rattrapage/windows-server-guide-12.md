---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-12
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["backlog", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1931, 2052]
sha256: d869a8cb9474e3d7039846381ac5eaa802a0d840a30dad42df90ef4dd581023e
---

# Windows Server en entreprise — Guide technique ultra-complet

### Cas n°10 — Restaurer un fichier écrasé par erreur il y a 2 jours
1. Versions précédentes (VSS §55) : clic droit → restaurer la version d'il y a 2 jours.
2. Si VSS non activé : `wbadmin start recovery` depuis la sauvegarde de la veille.
3. Vérifier l'intégrité avec l'utilisateur avant de clore.
*Leçon : activer VSS 2×/jour sur tous les volumes de partages — ça sauve plus de temps que n'importe quel outil.*

---

## 67. Cas pratiques commentés (11 à 15)

### Cas n°11 — Sécuriser le RDP exposé pour 3 administrateurs nomades
Ne **jamais** exposer le 3389 sur Internet. Solution : passerelle RDS (§43) ou VPN (WireGuard, voir guide Debian/Ubuntu) + RDP interne + NLA + comptes nominatifs + LAPS (§63) + audit 4625. Alternative moderne : Windows Admin Center publié via reverse proxy avec MFA.

### Cas n°12 — Migrer DHCP vers un nouveau serveur sans perdre les baux
1. `Export-DhcpServer -File C:\dhcp.xml -Leases` sur l'ancien.
2. `Import-DhcpServer -File C:\dhcp.xml -Leases -BackupPath C:\dhcp-backup` sur le nouveau.
3. Autoriser le nouveau dans l'AD, désautoriser l'ancien, vérifier les étendues et les réservations (copieurs ! §42).

### Cas n°13 — Un disque du pool Storage Spaces passe en "retired"
1. `Get-PhysicalDisk | Where-Object HealthStatus -ne "Healthy"` → identifier.
2. Ajouter le disque de remplacement au pool, `Repair-VirtualDisk`.
3. Retirer l'ancien (`Remove-PhysicalDisk`) après resync complète.
4. **Ne jamais** retirer un disque avant la fin de la réparation.

### Cas n°14 — Déployer LAPS sur 80 serveurs existants
1. Étendre le schéma (§63), délégations sur les OU.
2. GPO Windows LAPS (2025 : natif) liée aux OU Serveurs.
3. Vérifier : `Get-LapsADPassword -Identity SRV-X-01` sur un échantillon.
4. Révoquer les anciens mots de passe locaux partagés.

### Cas n°15 — Préparer l'audit de conformité licences
1. Inventaire (§58) : serveurs, éditions, VM par hôte.
2. Calculer : licences Standard/Datacenter requises vs achetées.
3. Vérifier les CAL (utilisateurs/appareils) vs effectif.
4. Documenter KMS/ADBA et les clés dans le coffre.
*Leçon : l'audit se prépare en continu, pas la veille du contrôle.*

---

## 68. Cas pratiques commentés (16 à 18)

### Cas n°16 — Panne du PDC : l'heure dérive sur tout le domaine
**Symptômes :** échecs Kerberos aléatoires ("l'horloge n'est pas synchronisée"). **Cause :** le PDC émulateur pointait vers une source NTP morte. **Action :** `w32tm /config /manualpeerlist:"0.fr.pool.ntp.org,1.fr.pool.ntp.org" /syncfromflags:manual /reliable:yes /update` sur le PDC, `w32tm /resync` sur les membres. **Préventif :** superviser le décalage NTP (Zabbix : `system.localtime` vs référence).

### Cas n°17 — Un service métier ne démarre plus après patch Tuesday
1. Identifier le patch (historique : `Get-HotFix | Sort-Object InstalledOn -Descending`).
2. Si régression avérée : `wusa /uninstall /kb:XXXXXXX` (avec fenêtre de maintenance).
3. Bloquer le KB sur WSUS (refuser l'approbation, §27) en attendant le correctif.
4. Tester les patchs sur le groupe `Serveurs-Test` avant la prod (toujours).

### Cas n°18 — Dimensionner un hôte Hyper-V pour 12 VM
Méthode : inventorier les besoins (vCPU/RAM/disque par VM), appliquer les ratios (§23 : ~8:1 vCPU/pCPU max, RAM = somme + 20 % pour l'hôte), prévoir la croissance 3 ans (+30 %). Exemple : 12 VM × (2 vCPU, 4 Go) → 24 vCPU → 8 cœurs physiques (ratio 3:1, confortable) ; RAM : 48 Go + 8 Go hôte + marge → 64 Go ; stockage : 12 × 100 Go + marge → 2 To utiles en RAID/miroir. Licence : Datacenter (12 VM > 6, §2).

---

## 69. Erreurs classiques (1 à 5)

1. **Cloner une VM sans sysprep** → doublons de SID/SusClientId : WSUS voit 1 client, l'AD se mélange. Toujours sysprep ou déployer depuis un template propre.
2. **Mettre des permissions NTFS à des utilisateurs nommés** au lieu de groupes → ingérable au premier départ. Groupes AD systématiques (§31).
3. **Oublier le délai de grâce RDS de 120 jours** → plus personne ne se connecte un lundi matin (§45).
4. **Un checkpoint Hyper-V vieux de 6 mois** considéré comme "sauvegarde" → chaîne AVHDX corrompue un jour = VM perdue (§20).
5. **IP en DHCP sur un serveur, un copieur, une imprimante** → l'IP change, tout casse. Statique ou réservation, sans exception (§6, §42).

---

## 70. Erreurs classiques (6 à 10)

6. **Sauvegarder sans jamais tester la restauration** → le jour J, la sauvegarde est vide/corrompue/illisible (§54).
7. **Ouvrir le RDP (3389) sur Internet** → brute-force en quelques heures. Passerelle/VPN uniquement (§60).
8. **Désactiver le pare-feu Windows "pour que ça marche"** → ouvrir le port précis au lieu de tout ouvrir (§61).
9. **Installer des rôles applicatifs sur un contrôleur de domaine** (fichiers, impression, IIS...) → DC = AD/DNS/DHCP, rien d'autre.
10. **Ne pas documenter** (mots de passe dans un tableur, pas d'inventaire) → le jour où l'admin part, c'est le black-out. GLPI + coffre à mots de passe (§6).

---

## 71. Erreurs classiques (11 à 15)

11. **Même mot de passe admin local partout** → LAPS (§63), sans débat.
12. **WSUS sans maintenance** → base de 15 Go, console inutilisable, clients non patchés (§28).
13. **Approuver toutes les mises à jour pour tout le monde d'un coup** → un patch foireux = tout le parc HS. Anneaux Test → Pilotes → Prod (§27).
14. **Désactiver SMBv1... sur le papier seulement** : vérifier avec `Get-SmbServerConfiguration`, car certains vieux copieurs/scanners exigent encore SMBv1 → isoler ces équipements sur un VLAN dédié plutôt que réactiver SMBv1 partout.
15. **Oublier de sauvegarder la configuration** (GPO, DHCP, DNS, certificats AC) → on sauvegarde les données mais pas la conf, et la reconstruction prend 3 jours. Exporter régulièrement (`Backup-GPO -All`, `Export-DhcpServer`, sauvegarde de l'AC via `certutil -backup`).

---

## 72. Erreurs classiques (16 à 18)

16. **Stocker les VM sur C:** → le système sature, l'hôte plante. Volume dédié (§23).
17. **Ne pas séparer les réseaux du cluster** (heartbeat + VM + CSV sur le même lien) → bascules fantômes en pleine journée (§22).
18. **Faire confiance à la corbeille AD par défaut** : activer la **corbeille AD** (`Enable-ADOptionalFeature -Identity "CN=Recycle Bin Feature,CN=Optional Features,CN=Directory Service,CN=Windows NT,CN=Services,CN=Configuration,DC=entreprise,DC=lan" -Scope ForestOrConfigurationSet -Target "entreprise.lan"`) pour restaurer un objet supprimé sans sueur froide.

---

---

## 73. Checklist d'exploitation quotidienne

- [ ] Sauvegardes de la nuit : statut OK (console Veeam/WSB, e-mail de rapport)
- [ ] Espace disque : aucun volume > 85 % (Zabbix ou script)
- [ ] Journaux critiques : erreurs System/Application dernières 24 h
- [ ] Réplication : Hyper-V Replica / DFS-R backlog nominal
- [ ] Cluster : tous les nœuds `Up`, ressources `Online`
- [ ] Antivirus : signatures à jour, aucune menace non traitée
- [ ] Files d'impression : aucune file en erreur massive (spooler OK)

```powershell
# Rapport quotidien express (à mettre en tâche planifiée + envoi e-mail)
$rapport = @()
$rapport += "=== Disques > 80% ==="
$rapport += Get-CimInstance Win32_LogicalDisk -Filter "DriveType=3" |
  Where-Object { ($_.FreeSpace / $_.Size) -lt 0.2 } |
  ForEach-Object { "$($_.DeviceID) libre : $([math]::Round($_.FreeSpace/1GB,1)) Go" }
$rapport += "=== Erreurs System 24h ==="
$rapport += (Get-WinEvent -FilterHashtable @{LogName='System'; Level=1,2; StartTime=(Get-Date).AddDays(-1)} -MaxEvents 20 |
  ForEach-Object { "$($_.TimeCreated) [$($_.Id)] $($_.ProviderName)" })
$rapport | Out-File "D:\Rapports\quotidien-$(Get-Date -Format yyyy-MM-dd).txt"
```

---

## 74. Checklist hebdomadaire

