---
id: collect-261001-rattrapage/rattrapage/win11-guide-15
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [2305, 2525]
sha256: f5263f6b5d89a6122f493e3a63b7600a90dddc5647105e2bc794e48a952953db
---

# Windows 11 en entreprise — Guide technique ultra-complet

```xml
<VPNProfile>
  <NativeProfile>
    <Servers>vpn.entreprise.example</Servers>
    <NativeProtocolType>IKEv2</NativeProtocolType>
    <Authentication>
      <MachineMethod>Certificate</MachineMethod>
    </Authentication>
    <RoutingPolicyType>SplitTunnel</RoutingPolicyType>
  </NativeProfile>
  <AlwaysOn>true</AlwaysOn>
  <DnsSuffix>entreprise.local</DnsSuffix>
  <TrustedNetworkDetection>entreprise.local</TrustedNetworkDetection>
</VPNProfile>
```

```powershell
# Appliquer un profil VPN en PowerShell (test / dépannage)
Add-VpnConnection -Name "ENTREPRISE-AO" -ServerAddress "vpn.entreprise.example" `
    -TunnelType Ikev2 -AuthenticationMethod MachineCertificate `
    -SplitTunneling -DnsSuffix "entreprise.local" -AllUserConnection

# État
Get-VpnConnection -Name "ENTREPRISE-AO" | Select-Object Name, ConnectionStatus, TunnelType
```

> ⚠️ **Split tunnel vs full tunnel** : le split tunnel (seul le trafic d'entreprise passe dans le VPN) économise la bande passante mais expose la navigation Internet du poste. Choix à valider avec la politique de sécurité (souvent : full tunnel pour les postes sensibles).

---

## 38. Réseau : dépannage — pile TCP/IP, DNS, proxy, certificats

### 38.1 La séquence de dépannage réseau (dans l'ordre)

```powershell
# 1. L'interface a-t-elle une adresse ?
Get-NetIPAddress -AddressFamily IPv4 | Where-Object { $_.InterfaceAlias -notlike "*Loopback*" }

# 2. La passerelle répond-elle ?
Test-NetConnection 10.0.0.254 -InformationLevel Detailed  # (adresse fictive d'exemple)

# 3. Le DNS résout-il ?
Resolve-DnsName intranet.entreprise.local -Server 10.0.0.10
nslookup intranet.entreprise.local

# 4. Le port du service est-il ouvert ?
Test-NetConnection srv-fichiers.entreprise.local -Port 445

# 5. Traceroute
tracert -d 8.8.8.8
```

### 38.2 Réinitialiser la pile réseau (quand tout est « bizarre »)

```powershell
# Dans l'ordre, puis redémarrer :
netsh winsock reset
netsh int ip reset
ipconfig /release
ipconfig /renew
ipconfig /flushdns
# Redémarrage obligatoire après winsock reset.
```

### 38.3 Proxy (cause fréquente en entreprise)

```powershell
# Voir la config proxy (WinHTTP = services, WinINET = utilisateur)
netsh winhttp show proxy
Get-ItemProperty "HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings" |
    Select-Object ProxyEnable, ProxyServer, AutoConfigURL

# Forcer l'alignement WinHTTP sur la config IE (utile pour Windows Update / Intune)
netsh winhttp import proxy source=ie
```

### 38.4 Certificats (802.1X, VPN, Wi-Fi, sites internes)

```powershell
# Certificats machine / utilisateur
Get-ChildItem Cert:\LocalMachine\My | Select-Object Subject, NotAfter, Thumbprint
Get-ChildItem Cert:\CurrentUser\My  | Select-Object Subject, NotAfter, Thumbprint

# Tester la chaîne de confiance d'un certificat
# (PowerShell 7+ : Test-Certificate ; sinon : certutil)
certutil -verify -urlfetch C:\Admin\cert.cer
```

---

## 39. WinRE : environnement de récupération, outils, personnalisation

### 39.1 Accéder à WinRE

| Méthode | Comment |
|---|---|
| Échecs de démarrage répétés (2-3) | Automatique |
| Depuis Windows | Paramètres → Récupération → Démarrage avancé → Redémarrer |
| Ligne de commande | `shutdown /r /o /f /t 0` |
| Écran d'ouverture de session | Maj + Redémarrer |

### 39.2 Outils disponibles dans WinRE

- **Invite de commandes** : diskpart, bcdboot, bootrec, manage-bde, chkdsk, sfc, dism (offline).
- **Désinstaller les mises à jour** : rollback d'une KB fautive sans démarrer Windows.
- **Restauration du système** : points de restauration.
- **Réparation du démarrage** : automatique (efficacité variable, voir cas pratique 1).
- **Réinitialiser ce PC**.

```powershell
# Vérifier que WinRE est activé et où se trouve son image
reagentc /info
# Sortie : Windows RE status: Enabled, emplacement \\?\GLOBALROOT\device\harddisk0\partition4\Recovery\WindowsRE

# Activer / désactiver
reagentc /enable
reagentc /disable
```

### 39.3 Commandes de réparation depuis WinRE (antisèche)

```powershell
# Reconstruire le BCD (UEFI)
diskpart
list vol
# identifier la partition EFI (FAT32, ~100 Mo) → lettre S:, et Windows → C:
exit
bcdboot C:\Windows /s S: /f UEFI

# Vérifier le disque
chkdsk C: /f /r

# Réparer les fichiers système hors ligne
sfc /scannow /offbootdir=C:\ /offwindir=C:\Windows

# Réparer l'image hors ligne (source : ISO montée ou WIM)
dism /Image:C:\ /Cleanup-Image /RestoreHealth /Source:E:\sources\install.wim:1
```

### 39.4 Personnaliser WinRE (avancé)

On peut injecter des pilotes/outils dans `winre.wim` (ex. : pilote RAID pour voir les disques) :

```powershell
# Monter, injecter, démonter (à faire sur une machine de lab)
dism /Mount-Wim /WimFile:C:\WinRE\winre.wim /Index:1 /MountDir:C:\Mount
dism /Image:C:\Mount /Add-Driver /Driver:C:\Drivers\RAID /Recurse
dism /Unmount-Wim /MountDir:C:\Mount /Commit
reagentc /setreimage /path C:\WinRE /target C:\Windows
```

---

## 40. Points de restauration, sauvegardes système, récupération

### 40.1 Points de restauration

```powershell
# Activer la protection système sur C:
Enable-ComputerRestore -Drive "C:\"

# Créer un point manuel (avant une intervention risquée !)
Checkpoint-Computer -Description "Avant MàJ pilote GPU" -RestorePointType "MODIFY_SETTINGS"

# Lister les points
Get-ComputerRestorePoint | Select-Object SequenceNumber, Description, CreationTime

# Restaurer (nécessite un redémarrage)
Restore-Computer -RestorePoint 12
```

> ⚠️ Les points de restauration ne protègent **ni les fichiers utilisateurs** ni les volumes chiffrés en cours de chiffrement. C'est un filet anti-bêtise système, pas une sauvegarde.

### 40.2 Sauvegarde image système (héritage, toujours présent)

```powershell
# Sauvegarde complète vers un disque externe (outil historique wbadmin)
wbadmin start backup -backupTarget:E: -include:C: -allCritical -quiet
```

### 40.3 Stratégie réaliste « récupération poste » en entreprise

| Scénario | Solution |
|---|---|
| Fichiers utilisateur | OneDrive KFM (continu) |
| Régression système (pilote, KB) | Point de restauration + désinstallation KB |
| Poste HS / volé / réaffecté | **Re-déploiement** (MDT/Autopilot) + OneDrive, pas de « réparation héroïque » |
| Besoin de preuve / forensique | Image disque bit-à-bit avant toute manipulation |

> 💡 **Doctrine** : en entreprise, un poste est **jetable**, les données ne le sont pas. Investissez dans OneDrive/KFM et le redéploiement rapide plutôt que dans des sauvegardes images de postes.

---

## 41. Dépannage : méthodologie et boîte à outils de l'admin

### 41.1 Méthode en 6 étapes

```
1. CADRER      : qui, quoi, quand, depuis quand, quoi de changé ?
2. REPRODUIRE  : le problème est-il reproductible ? Sur un autre poste / autre utilisateur ?
3. ISOLER      : dichotomie — réseau ? profil ? GPO ? pilote ? MàJ récente ?
4. JOURNAUX    : Event Viewer, journaux applicatifs, logs (toujours AVANT de toucher)
5. CORRIGER    : une seule modification à la fois, documentée
6. VÉRIFIER    : le problème a disparu ET rien d'autre n'a cassé
```

### 41.2 Boîte à outils (à avoir sur sa clé / son poste)

| Outil | Usage |
|---|---|
| Sysinternals Suite (Process Explorer, Autoruns, TCPView…) | Diagnostic avancé |
| SetupDiag | Échecs de mise à niveau |
| MDT / WinPE bootable | Réparation hors ligne |
| `winget` + scripts | Réinstallation rapide d'apps |
| RSAT | AD, GPO, DNS à distance |
| Wireshark | Captures réseau (avec accord) |
| CrystalDiskInfo | Santé SMART des disques |
| Carnet de bord | Noter chaque intervention (ticket GLPI !) |

### 41.3 Les journaux à connaître par cœur

