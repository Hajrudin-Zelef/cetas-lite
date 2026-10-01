---
id: collect-261001-rattrapage/rattrapage/win11-guide-18
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [2939, 3138]
sha256: 5cbbea561448f0427d0a1d66c795d24eb345a82906bc972d5de51988918b5310
---

# 7. Forcer + journal
gpupdate /force
Get-WinEvent -FilterHashtable @{LogName='Microsoft-Windows-GroupPolicy/Operational'; Level=2; StartTime=(Get-Date).AddDays(-1)} |
    Select-Object -First 5 TimeCreated, Id, Message
```

**Causes top 5** : filtrage de sécurité mal configuré (droits manquants) · héritage bloqué · GPO liée à la mauvaise OU · réplication SYSVOL incomplète · boucle de traitement (loopback) non configurée pour les GPO utilisateur sur postes partagés.

---

## 52. Cas pratique 11 : Wi-Fi 802.1X — échec d'authentification

**Symptômes** : « Impossible de se connecter à ENTREPRISE », ou connexion puis déconnexion en boucle.

**Diagnostic** :

```powershell
# 1. Le profil est-il correct ?
netsh wlan show profile name="ENTREPRISE"

# 2. Journaux WLAN (très verbeux, très utiles)
Get-WinEvent -FilterHashtable @{LogName='Microsoft-Windows-WLAN-AutoConfig/Operational'; Level=2; StartTime=(Get-Date).AddDays(-1)} |
    Select-Object TimeCreated, Id, Message | Format-List

# 3. Le certificat machine est-il présent et valide ? (EAP-TLS)
Get-ChildItem Cert:\LocalMachine\My | Where-Object { $_.NotAfter -gt (Get-Date) } |
    Select-Object Subject, NotAfter

# 4. L'heure est-elle correcte ? (EAP-TLS échoue si l'horloge dérive)
w32tm /query /status
```

**Résolutions** :

| Cause | Action |
|---|---|
| Certificat expiré / absent | Renouveler via auto-enrollment (`certutil -pulse`), vérifier le modèle AD CS |
| Serveur RADIUS non approuvé | Ajouter le nom du certificat NPS dans le profil (serveurs de confiance) |
| Identifiants PEAP erronés | Effacer les identifiants mémorisés : `netsh wlan delete profile name="ENTREPRISE"` puis re-déployer |
| Heure désynchronisée | Forcer la synchro NTP |
| Pilote Wi-Fi | Mettre à jour (constructeur, pas Windows Update) |

---

## 53. Cas pratique 12 : VPN qui ne se connecte pas

**Symptômes** : erreur 800/809/13801/691 sur la connexion VPN.

| Code | Signification | Piste |
|---|---|---|
| 800 | Serveur injoignable | Réseau, pare-feu, nom DNS |
| 809 | Bloqué par NAT/pare-feu | Ports UDP 500/4500, activer le NAT-T |
| 13801 | Échec IKEv2 (certificat) | Certificat serveur (SAN), chaîne de confiance |
| 691 | Identifiants refusés | Login/mot de passe, compte verrouillé |
| 942 | Certificat machine absent | Auto-enrollment, modèle de certificat |

```powershell
# Diagnostic
Get-VpnConnection | Select-Object Name, ConnectionStatus, TunnelType, ServerAddress
# Journaux RRAS côté client :
Get-WinEvent -FilterHashtable @{LogName='Application'; ProviderName='RasClient'; StartTime=(Get-Date).AddDays(-1)} |
    Select-Object TimeCreated, Id, Message | Format-List

# Tester la connectivité vers le serveur VPN
Test-NetConnection vpn.entreprise.example -Port 443
Resolve-DnsName vpn.entreprise.example
```

> 💡 En télétravail, 80 % des tickets VPN sont : Wi-Fi domestique instable, heure désynchronisée, ou certificat expiré. Vérifiez ces trois points avant d'ouvrir la console RRAS.

---

## 54. Cas pratique 13 : applications qui ne se lancent plus / Store corrompu

**Symptômes** : les applications du Store (ou toutes les UWP) plantent au lancement, le Store ne s'ouvre pas.

**Résolution** (dans l'ordre, du moins au plus invasif) :

```powershell
# 1. Réparer le Store
wsreset.exe   # vide le cache du Store (fenêtre noire, attendre la fin)

# 2. Réinscrire les applications pour l'utilisateur courant
Get-AppXPackage -AllUsers | ForEach-Object {
    Add-AppxPackage -DisableDevelopmentMode -Register "$($_.InstallLocation)\AppXManifest.xml"
}

# 3. Réparer une application précise (ex. : Calculatrice)
Get-AppxPackage *calculator* | Reset-AppxPackage

# 4. Vérifier l'intégrité système
dism /Online /Cleanup-Image /RestoreHealth
sfc /scannow

# 5. Si un profil est en cause : tester avec un nouvel utilisateur local
# (si ça marche avec le nouvel utilisateur → profil corrompu, voir cas n°8)
```

---

## 55. Cas pratique 14 : espace disque saturé après mise à niveau

**Symptômes** : C: plein à 95 %+ après une feature update (`Windows.old` = 20-30 Go).

**Résolution** :

```powershell
# 1. Voir ce qui prend de la place
Get-PSDrive C | Select-Object Used, Free
# Dossiers suspects :
Get-ChildItem C:\ -Directory | ForEach-Object {
    $size = (Get-ChildItem $_.FullName -Recurse -Force -ErrorAction SilentlyContinue |
             Measure-Object Length -Sum).Sum / 1GB
    [pscustomobject]@{ Dossier=$_.Name; Go=[math]::Round($size,1) }
} | Sort-Object Go -Descending | Select-Object -First 10

# 2. Nettoyage système (lance l'assistant)
cleanmgr.exe /sageset:1
cleanmgr.exe /sagerun:1

# 3. Supprimer Windows.old APRÈS validation (10 jours de rollback par défaut !)
# Paramètres → Système → Stockage → Fichiers temporaires → cocher "Version précédente de Windows"
# En ligne de commande :
dism /Online /Cleanup-Image /StartComponentCleanup /ResetBase
# /ResetBase : irréversible, supprime la possibilité de désinstaller les MàJ — à n'utiliser
# qu'après validation complète du poste

# 4. Recommandation disque (Storage Sense) : activer par GPO
$ss = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\StorageSense"
New-Item -Path $ss -Force | Out-Null
Set-ItemProperty -Path $ss -Name "AllowStorageSenseGlobal" -Value 1 -Type DWord
```

> ⚠️ Ne supprimez **jamais** `Windows.old` manuellement à la main dans l'Explorateur (droits système) : utilisez l'outil de nettoyage de disque. Et attendez la fin de la période de rollback (10 jours) avant `/ResetBase`.

---

## 56. Cas pratique 15 : Windows Hello / biométrie ne fonctionne plus

**Symptômes** : le PIN ou l'empreinte ne fonctionne plus après une MàJ ou un changement de domaine.

**Diagnostic** :

```powershell
# 1. Le TPM est-il sain ? (Windows Hello repose dessus)
Get-Tpm | Select-Object TpmReady, TpmEnabled, LockedOut

# 2. État de l'inscription Hello
# Paramètres → Comptes → Options de connexion → tout est grisé ?
# → souvent lié à une stratégie (désactivé par GPO) ou au PIN "oublié" par le TPM
```

**Résolution** :

```
1. Supprimer puis recréer le PIN : Paramètres → Comptes → Options de connexion
   → Code PIN → Supprimer → Redémarrer → Reconfigurer
2. Si le TPM a été effacé : la clé Hello est perdue → recréation obligatoire
   (c'est normal, c'est le principe de sécurité)
3. Pilote du lecteur d'empreinte : réinstaller depuis le constructeur
4. En entreprise : vérifier la GPO "Utiliser Windows Hello Entreprise"
   (nécessite une infrastructure : certificats ou clés, selon le modèle)
```

```powershell
# Réinitialiser le conteneur NGC (base des PIN Hello) — méthode radicale :
# 1. Prendre possession de C:\Windows\ServiceProfiles\LocalService\AppData\Local\Microsoft\NGC
# 2. Supprimer son contenu, redémarrer, recréer le PIN
takeown /f "C:\Windows\ServiceProfiles\LocalService\AppData\Local\Microsoft\NGC" /r /d y
icacls "C:\Windows\ServiceProfiles\LocalService\AppData\Local\Microsoft\NGC" /grant administrateurs:F /t
Remove-Item "C:\Windows\ServiceProfiles\LocalService\AppData\Local\Microsoft\NGC\*" -Recurse -Force
```

---

## 57. Cas pratique 16 : imprimante réseau / partage SMB inaccessible

**Symptômes** : « Windows ne peut pas accéder à \\SRV\partage », erreur 0x80070035, imprimante hors ligne.

**Diagnostic** :

```powershell
# 1. Résolution de noms
Resolve-DnsName srv-fichiers.entreprise.local
nbtstat -a SRV-FICHIERS   # NetBIOS (si encore utilisé)

# 2. Connectivité SMB
Test-NetConnection srv-fichiers.entreprise.local -Port 445

# 3. Version SMB négociée (doit être 3.x)
Get-SmbConnection | Select-Object ServerName, Dialect, NumOpens

# 4. Identifiants : avec quel compte suis-je authentifié ?
klist   # tickets Kerberos
# Tester avec des identifiants explicites :
net use \\srv-fichiers\partage /user:ENTREPRISE\jdupont *
```

**Causes fréquentes** :

