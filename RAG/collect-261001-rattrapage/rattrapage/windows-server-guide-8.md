---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-8
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attribution", "distribution"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1254, 1426]
sha256: f8cc2589b6de1ac28a21879699c4cbb02caa8d28503fdd83add53f80ad839c19
---

# Windows Server en entreprise — Guide technique ultra-complet

```powershell
# Diagnostic express d'un serveur d'impression
Get-Service Spooler | Select-Object Name, Status
Get-Printer -ComputerName SRV-PRINT-01 | Where-Object PrinterStatus -ne "Normal" |
  Format-Table Name, PrinterStatus, PortName -AutoSize
# Tester le port d'un copieur
Test-NetConnection -ComputerName 192.168.20.50 -Port 9100
```

---

## 43. RDS : les 4 rôles (Broker, Session Host, Passerelle, Accès Web)

| Rôle | Fonction |
|---|---|
| **RD Connection Broker** | Aiguille les connexions, gère les collections, reconnexion de session |
| **RD Session Host** | Héberge les sessions utilisateurs (le "serveur TSE") |
| **RD Gateway** | Proxy HTTPS (443) pour accès externes sans VPN |
| **RD Web Access** | Portail web listant les applis/bureaux publiés (RemoteApp) |

Topologies :
- **PME** : tout sur 1–2 serveurs (Broker + Web + Passerelle sur l'un, Session Hosts sur l'autre).
- **Entreprise** : Broker en HA (SQL), ferme de Session Hosts, Passerelle en DMZ.

```powershell
# Installer les 4 rôles (déploiement rapide, 1 serveur)
Install-WindowsFeature -Name RDS-RD-Server -IncludeManagementTools
# Puis : Server Manager → Services Bureau à distance → déploiement basé sur les scénarios
```

---

## 44. Déploiement RDS pas à pas

**Via Server Manager (recommandé pour la topologie) :**

1. Server Manager → **Services Bureau à distance** → **Vue d'ensemble** → déploiement standard.
2. Ajouter les serveurs aux 4 rôles.
3. Créer une **collection de sessions** : nom, Session Hosts membres, groupes d'utilisateurs autorisés.
4. Publier des **RemoteApp** (ex. l'ERP) ou le **bureau complet**.
5. Configurer le **certificat** (public si passerelle exposée, voir §48 PKI en interne).
6. Tester avec un compte utilisateur standard.

**En PowerShell :**

```powershell
# Créer la collection + publier une RemoteApp
New-RDSessionCollection -CollectionName "Bureautique" -SessionHost "SRV-RDS-01.entreprise.lan","SRV-RDS-02.entreprise.lan" `
  -ConnectionBroker "SRV-RDS-CB-01.entreprise.lan"
New-RDRemoteApp -CollectionName "Bureautique" -DisplayName "ERP Gestion" `
  -FilePath "C:\Program Files\ERP\erp.exe" -Alias "ERP"
Set-RDSessionCollectionConfiguration -CollectionName "Bureautique" -UserGroup "ENTREPRISE\GRP-RDS-Users"

# Équilibrage : pondération des hôtes
Set-RDSessionHost -SessionHost "SRV-RDS-01.entreprise.lan" -NewConnectionAllowed Yes
```

> Les hôtes de session RDS sont le **cas n°1 du Desktop Experience** (§10) : les utilisateurs ont besoin d'une GUI complète.

---

## 45. Licences CAL RDS : par utilisateur vs par appareil

En plus des CAL Windows Server (§58), RDS exige des **CAL RDS** :

| | Par utilisateur | Par appareil |
|---|---|---|
| Compte | Chaque **utilisateur** nommé | Chaque **appareil** |
| Idéal | Utilisateurs multi-appareils (PC + tablette + maison) | Postes partagés (3×8, ateliers, bornes) |
| Suivi | Attribution permanente, révocable après 90 j | Attribution à l'appareil |

```powershell
# Installer le rôle de gestion des licences
Install-WindowsFeature -Name RDS-Licensing -IncludeManagementTools

# Activer le serveur de licences + installer les CAL (via l'assistant ou)
# Configurer le mode de licence sur la collection :
Set-RDLicenseConfiguration -LicenseServer "SRV-RDS-LIC-01.entreprise.lan" -Mode PerUser -Force

# Vérifier l'état des licences
Get-RDLicenseConfiguration
```

> **Délai de grâce : 120 jours.** Passé ce délai sans serveur de licences configuré, les connexions sont refusées. C'est la panne RDS n°1.

---

## 46. Dépannage RDS

| Symptôme | Cause probable | Action |
|---|---|---|
| "Aucun serveur de licences disponible" | Délai de grâce expiré | Configurer le serveur de licences (§45) |
| Écran noir à la connexion | Session bloquée / GPO | Fermer la session depuis le Gestionnaire : `logoff <id> /server:SRV-RDS-01` |
| Lenteurs générales | Surcharge hôte | `Get-RDUserSession` : compter les sessions ; vérifier RAM/CPU |
| RemoteApp ne se lance pas | Chemin ou droits | Vérifier le FilePath publié + groupe autorisé |
| Passerelle inaccessible de l'extérieur | Certificat / 443 | Certificat valide (nom externe), NAT 443 → passerelle |
| Imprimantes locales non remontées | Redirection désactivée | GPO : autoriser la redirection d'imprimantes ; pilote Easy Print |

```powershell
# Sessions actives et déconnexion
Get-RDUserSession -ConnectionBroker "SRV-RDS-CB-01.entreprise.lan"
Invoke-RDUserLogoff -HostServer "SRV-RDS-01.entreprise.lan" -UnifiedSessionID 3 -Force

# Voir qui consomme quoi sur un hôte
quser /server:SRV-RDS-01
Get-Process -ComputerName SRV-RDS-01 | Sort-Object WS -Descending | Select-Object -First 10 Name, WS
```

---

## 47. AD CS / PKI : concepts (AC racine, subordonnée, modèles)

**Hiérarchie type en entreprise :**
- **AC racine autonome (offline)** : allumée 1×/an pour signer la subordonnée, sinon éteinte et coffrée. Durée de vie 15–20 ans.
- **AC subordonnée d'entreprise (online)** : intégrée à l'AD, émet les certificats au quotidien. Durée de vie 5–10 ans.

**Briques :**
- **Modèle de certificat** : définit qui peut demander quoi (ex. "Utilisateur", "Ordinateur", "Serveur Web").
- **Auto-enrollment** : distribution automatique via GPO, sans intervention.
- **CRL** (liste de révocation) : publiée sur HTTP/LDAP, consultée par les clients.
- **OCSP** : alternative en ligne à la CRL.

```powershell
# Installer AD CS (sur le futur serveur d'AC subordonnée membre du domaine)
Install-WindowsFeature -Name AD-Certificate -IncludeManagementTools
# Services de rôle : Certification Authority, (optionnel) Online Responder, Web Enrollment
Install-WindowsFeature -Name ADCS-Cert-Authority, ADCS-Online-Cert, ADCS-Web-Enrollment
```

---

## 48. Installation d'une AC d'entreprise pas à pas

**Étape A — AC racine autonome (hors domaine, VM isolée) :**

```powershell
# Sur la future racine (WORKGROUP, jamais jointe au domaine)
Install-WindowsFeature -Name AD-Certificate -IncludeManagementTools
Install-AdcsCertificationAuthority -CAType StandaloneRootCA `
  -CryptoProviderName "RSA#Microsoft Software Key Storage Provider" `
  -KeyLength 4096 -HashAlgorithmName SHA512 -ValidityPeriod Years -ValidityPeriodUnits 20 `
  -CACommonName "ENTREPRISE-Root-CA" -Force
```

**Étape B — AC subordonnée d'entreprise (membre du domaine) :**

```powershell
Install-WindowsFeature -Name AD-Certificate -IncludeManagementTools
# 1. Générer la demande :
Install-AdcsCertificationAuthority -CAType EnterpriseSubordinateCA -CACommonName "ENTREPRISE-Issuing-CA-01" `
  -KeyLength 4096 -HashAlgorithmName SHA384 -ValidityPeriod Years -ValidityPeriodUnits 10 `
  -OutputCertRequestFile "C:\req-issuing.req" -Force
# 2. Signer la demande sur la RACINE (console certsrv ou certreq)
# 3. Installer le certificat signé :
Install-AdcsCertificationAuthority -CAType EnterpriseSubordinateCA -CertificateFile "C:\issuing-signé.p7b" -Force
# 4. Éteindre et archiver la racine (export clé privée sur support chiffré, coffre-fort)
```

**Publier la CRL et AIA en HTTP** (indispensable : les clients non-joints au domaine doivent y accéder) :

```powershell
# Sur l'AC : ajouter un point de distribution HTTP
certutil -setreg CA\CRLPublicationURLs "1:C:\Windows\System32\CertSrv\CertEnroll\%%3%%8%%9.crl\n2:http://pki.entreprise.lan/crld/%%3%%8%%9.crl"
certutil -setreg CA\CACertPublicationURLs "1:C:\Windows\System32\CertSrv\CertEnroll\%%1_%%3%%4.crt\n2:http://pki.entreprise.lan/crld/%%1_%%3%%4.crt"
Restart-Service certsvc
# Publier une CRL fraîche
certutil -crl
```

---

## 49. Modèles de certificats et auto-enrollment via GPO

