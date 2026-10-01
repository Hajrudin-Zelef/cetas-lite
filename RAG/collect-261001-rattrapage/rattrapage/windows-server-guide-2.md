---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-2
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["arr", "ethernet"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [166, 359]
sha256: 91a116597a2706743d10a97b498d47cacca04017af52ab8b8731c102c5fa319f
---

# Windows Server en entreprise — Guide technique ultra-complet

1. Démarrer sur l'ISO (ou clé USB bootable via Rufus / Media Creation).
2. Langue, heure/devise, clavier → **Suivant** → **Installer maintenant**.
3. Choisir l'édition **avec Expérience de bureau** (Desktop Experience) si tu veux une GUI.
4. Type d'installation : **Personnalisé** (jamais de mise à niveau sur un serveur de prod).
5. Partitionnement : supprimer/recréer si disque neuf. Laisser 100+ Go pour le système.
6. Mot de passe Administrateur local (fictif ici) : `Adm1n-L0cal-Exemple!` → **Terminer**.
7. Premier boot : `Ctrl+Alt+Suppr`, connexion Administrateur.
8. Server Manager se lance → passer à la checklist §6.

```powershell
# En une fois après le premier boot (à adapter) :
Rename-Computer -NewName "SRV-FICHIER-01" -Restart
```

> Astuce : note le nommage normalisé dès l'installation : `SRV-<ROLE>-<NN>` (ex. `SRV-HV-01`, `SRV-AD-01`, `SRV-PRINT-01`). Un nommage propre évite 50 % des confusions d'exploitation.

---

## 5. Installation Server Core pas à pas

Même procédure qu'au §4, mais choisir l'édition **sans** mention "Expérience de bureau" (ex. `Windows Server 2025 Standard` tout court).

Au premier démarrage : une fenêtre `cmd` + `sconfig` (voir §8). Pas de menu Démarrer, pas d'Explorateur.

```powershell
# Depuis la console Core, ouvrir PowerShell :
powershell

# Vérifier qu'on est bien en Core :
Get-ComputerInfo | Select-Object WindowsInstallationType
# Résultat attendu : Server Core
```

**Pièges fréquents à l'installation :**
- Oublier le mot de passe Administrateur → pas de récupération simple, réinstaller.
- Choisir la mauvaise édition (Core vs Desktop) → réinstallation obligatoire (pas de bascule GUI ↔ Core depuis 2016).
- Pilotes RAID/NVMe manquants → charger via **Charger un pilote** pendant l'installation.

---

## 6. Post-installation : la checklist des 20 minutes

À faire sur **chaque** serveur, dans l'ordre :

```powershell
# 1. Nom + domaine/groupe de travail
Rename-Computer -NewName "SRV-ROLE-01"

# 2. IP statique (jamais de DHCP sur un serveur d'infra)
New-NetIPAddress -InterfaceAlias "Ethernet" -IPAddress "192.168.10.11" `
  -PrefixLength 24 -DefaultGateway "192.168.10.1"
Set-DnsClientServerAddress -InterfaceAlias "Ethernet" -ServerAddresses "192.168.10.10","192.168.10.12"

# 3. Fuseau horaire + NTP
Set-TimeZone -Id "Romance Standard Time"   # France métropolitaine
w32tm /config /manualpeerlist:"192.168.10.10" /syncfromflags:manual /reliable:yes /update

# 4. Mises à jour
Install-Module PSWindowsUpdate -Force   # une fois
Get-WindowsUpdate -Install -AcceptAll -AutoReboot

# 5. Activer l'administration à distance
Enable-PSRemoting -Force
Set-NetFirewallRule -DisplayGroup "Gestion à distance de Windows" -Enabled True

# 6. Renommer + joindre au domaine (redémarrage requis)
Add-Computer -DomainName "entreprise.lan" -Credential (Get-Credential) -Restart
```

Checklist manuelle :
- [ ] IP statique + DNS corrects, passerelle OK (`Test-NetConnection 8.8.8.8`)
- [ ] Nom conforme à la convention
- [ ] Fuseau horaire + heure synchronisée (`w32tm /query /status`)
- [ ] Mises à jour installées
- [ ] Pare-feu actif avec profil Domaine
- [ ] RDP activé si besoin (voir §9)
- [ ] Antivirus / Defender à jour
- [ ] Sauvegarde planifiée (voir §52)
- [ ] Serveur documenté (CMDB/GLPI : rôle, IP, garanties)

---

## 7. Server Core vs Desktop Experience : différences concrètes

| Point | Server Core | Desktop Experience |
|---|---|---|
| Interface | Console texte + sconfig | GUI complète |
| Empreinte disque | ~6–9 Go | ~12–20 Go |
| Surface d'attaque | Réduite (~40 % de correctifs en moins) | Complète |
| Redémarrages (patchs) | Moins fréquents | Plus fréquents |
| Outils d'admin | PowerShell, sconfig, RSAT, WAC | + Server Manager, MMC locales |
| Cas d'usage | Hyper-V, AD, DNS, DHCP, fichiers | RDS Session Host, applis à GUI, transition |

**Depuis Windows Server 2016 : pas de bascule** entre Core et Desktop sans réinstallation. Le choix est définitif à l'installation.

---

## 8. sconfig : le couteau suisse de Server Core

Tape `sconfig` dans la console. Menu numéroté :

| N° | Action |
|---|---|
| 1 | Joindre au domaine / groupe de travail |
| 2 | Changer le nom de l'ordinateur |
| 3 | Ajouter un compte administrateur local |
| 4 | Configurer l'administration à distance (WinRM) |
| 5 | Paramètres Windows Update (Auto/Manuel/Désactivé) |
| 6 | Télécharger et installer les mises à jour |
| 7 | Activer/désactiver le Bureau à distance |
| 8 | Configurer la carte réseau (IP statique/DHCP, DNS) |
| 9 | Date et heure |
| 10 | Télémétrie |
| 11 | Paramètres Windows Update avancés |
| 13 | Redémarrer / arrêter |

Équivalents PowerShell directs (plus rapides en série) :

```powershell
# Tout ce que fait sconfig, en PowerShell :
sconfig                                  # lancer le menu
# ou en direct :
Set-NetIPInterface -InterfaceAlias "Ethernet" -Dhcp Disabled   # passer en statique via New-NetIPAddress
Enable-PSRemoting -Force                 # = option 4
cscript C:\Windows\System32\slmgr.vbs    # activation (option 11/12 selon version)
shutdown /r /t 0                         # = option 13
```

---

## 9. Administration à distance : RSAT, Windows Admin Center, WinRM

**Option A — RSAT (Remote Server Administration Tools)**, sur un poste Windows 10/11 :

```powershell
# Lister et installer RSAT (Windows 10 1809+ / 11 : fonctionnalités à la demande)
Get-WindowsCapability -Online | Where-Object Name -like 'Rsat*' | Select-Object Name, State
Add-WindowsCapability -Online -Name Rsat.ActiveDirectory.DS-LDS.Tools~~~~0.0.1.0
Add-WindowsCapability -Online -Name Rsat.DHCP.Tools~~~~0.0.1.0
Add-WindowsCapability -Online -Name Rsat.DNS.Tools~~~~0.0.1.0
Add-WindowsCapability -Online -Name Rsat.FileServices.Tools~~~~0.0.1.0
Add-WindowsCapability -Online -Name Rsat.Hyper-V.Tools~~~~0.0.1.0
```

**Option B — Windows Admin Center (WAC)** : portail web gratuit, idéal pour Server Core.

```powershell
# Installation silencieuse de WAC sur une machine d'administration
msiexec /i WindowsAdminCenter.msi /qn /L*v log.txt SME_PORT=443 SSL_CERTIFICATE_OPTION=generate
```

**Option C — WinRM / PowerShell Remoting** (la base de tout) :

```powershell
# Sur le serveur cible (une fois) :
Enable-PSRemoting -Force

# Depuis ta station :
Enter-PSSession -ComputerName SRV-HV-01 -Credential (Get-Credential)
Invoke-Command -ComputerName SRV-HV-01,SRV-HV-02 -ScriptBlock { Get-Service spooler }
```

> En 2025, privilégie **Windows Admin Center + PowerShell Remoting**. Le RDP permanent sur chaque serveur est un anti-pattern (sessions oubliées, consommation mémoire).

---

## 10. Quand choisir Core, quand choisir Desktop

**Choisir Server Core :**
- Hôtes Hyper-V (recommandation Microsoft officielle)
- Contrôleurs de domaine, DNS, DHCP
- Serveurs de fichiers (avec admin via WAC/RSAT)
- WSUS, serveurs d'impression

**Choisir Desktop Experience :**
- Hôtes de session RDS (les utilisateurs ont besoin d'une GUI complète)
- Applications métier qui exigent une interface graphique
- Serveur de transition / équipe junior en montée en compétence
- Outils fournisseurs sans équivalent ligne de commande (ex. certaines consoles de sauvegarde)

**Règle pratique :** par défaut Core ; Desktop uniquement sur justification écrite.

---

## 11. Rôles et fonctionnalités : concepts

- **Rôle** : une fonction serveur majeure (AD DS, Hyper-V, DNS, Serveur de fichiers, RDS...). Un rôle peut dépendre d'autres rôles.
- **Fonctionnalité** : un composant auxiliaire (.NET Framework, Windows Server Backup, SNMP, Telnet Client...).
- **Service de rôle** : sous-composant d'un rôle (ex. pour RDS : Broker, Session Host, Passerelle, Accès Web).

```powershell
# Lister tout ce qui est disponible
Get-WindowsFeature | Where-Object InstallState -eq 'Available' | Select-Object Name, DisplayName

