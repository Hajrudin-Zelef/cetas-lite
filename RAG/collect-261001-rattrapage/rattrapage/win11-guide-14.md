---
id: collect-261001-rattrapage/rattrapage/win11-guide-14
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [2109, 2304]
sha256: e999b089f54a20879f3f89348f66b1c2162c059d2ea6e838b6ab6a26973a2059
---

# Désactiver LLMNR + NetBIOS (via registre, déployé par GPO)
$tcp = "HKLM:\SOFTWARE\Policies\Microsoft\Windows NT\DNSClient"
New-Item -Path $tcp -Force | Out-Null
Set-ItemProperty -Path $tcp -Name "EnableMulticast" -Value 0 -Type DWord  # LLMNR off
```

### 33.3 Suivi du durcissement : Microsoft Secure Score

- **Microsoft Secure Score** (portail Defender) : note + actions recommandées pour l'organisation.
- **Exposure Score** : exposition des appareils.
- À suivre mensuellement en comité sécurité.

---

## 34. LAPS Windows (nouveau) : principe, GPO, récupération

### 34.1 Windows LAPS vs l'ancien LAPS (différences clés)

| Point | Ancien LAPS (outil) | Windows LAPS (intégré) |
|---|---|---|
| Installation | MSI à déployer | **Natif** depuis avril 2023 (MàJ cumulative) |
| Stockage | AD uniquement | **AD ou Entra ID** |
| Chiffrement du mot de passe | Non | Oui (optionnel, AD) |
| Historique des mots de passe | Non | Oui |
| Gestion DSRM | Non | Oui |

> ✅ **Windows LAPS est le standard** : l'ancien LAPS est déprécié. Si l'ancien est encore déployé, migrez (les stratégies sont distinctes, pas de conflit si on migre proprement par OU).

### 34.2 Configurer Windows LAPS (GPO)

`Configuration ordinateur\Modèles d'administration\Système\LAPS` :

| Stratégie | Réglage recommandé |
|---|---|
| Configurer le répertoire de sauvegarde | **Active Directory** (ou Entra ID si cloud) |
| Nom du compte administrateur à gérer | `Administrateur` (ou compte dédié `adm-local`) |
| Longueur du mot de passe | 20+ caractères |
| Âge maximal du mot de passe | 30 jours |
| Activer le chiffrement du mot de passe | Oui (nécessite le schéma AD à jour 2023+) |

```powershell
# Équivalent : forcer la rotation immédiate (test)
Reset-LapsPassword  # (module LAPS PowerShell si installé)

# Vérifier l'état LAPS sur un poste
Get-LapsAADPassword -DeviceId (Get-DeviceId) -AsPlainText  # Entra : via Graph/portail
```

### 34.3 Récupérer le mot de passe (AD)

```powershell
# Depuis un poste avec les outils RSAT / module LAPS :
Get-LapsADPassword -Identity "PC-COMPTA-042" -AsPlainText

# Délégation : donner au helpdesk le droit "Lire le mot de passe LAPS"
# (autorisation étendue ms-LAPS-Password sur l'OU des postes)
```

### 34.4 Procédure helpdesk « mot de passe admin local »

```
1. Vérifier l'identité du demandeur + le motif (ticket obligatoire)
2. Get-LapsADPassword -Identity <poste> -AsPlainText
3. Communiquer le mot de passe (jamais par e-mail), usage unique
4. Forcer la rotation après usage : le mot de passe expire automatiquement
   au prochain cycle (ou forcer via stratégie "expiration après utilisation")
```

> 🔐 Le compte admin local géré par LAPS **ne doit jamais** servir au quotidien : uniquement dépannage hors réseau / hors domaine. Toute utilisation = ticket + rotation.

---

## 35. Pare-feu Windows Defender : profils, règles, GPO, journalisation

### 35.1 Les 3 profils

| Profil | Quand | Politique par défaut recommandée |
|---|---|---|
| Domaine | Connecté au domaine AD | Entrant : bloquer (sauf exceptions) |
| Privé | Réseau de confiance (domicile) | Entrant : bloquer |
| Public | Wi-Fi public, inconnu | Entrant : bloquer strict |

### 35.2 Gérer par GPO (recommandé) plutôt qu'en local

`Configuration ordinateur\Stratégies\Paramètres Windows\Paramètres de sécurité\Pare-feu Windows Defender avec fonctions avancées de sécurité`.

Bonnes pratiques :

- Définir les règles **une fois** en GPO, interdire la fusion avec les règles locales sur les postes sensibles (`Appliquer les règles locales : Non`).
- Nommer les règles : `[APP]-[Sens]-[Port]` (ex. `Zabbix-Agent-In-10050`).
- Documenter chaque exception dans un registre des flux.

```powershell
# Créer une règle en PowerShell (pour script / image de référence)
New-NetFirewallRule -DisplayName "Zabbix-Agent-In-10050" `
    -Direction Inbound -Protocol TCP -LocalPort 10050 `
    -RemoteAddress "10.0.50.0/24" -Action Allow -Profile Domain

# Lister les règles actives
Get-NetFirewallRule -Enabled True -Direction Inbound |
    Select-Object DisplayName, Profile, Action | Sort-Object DisplayName

# Règle de test temporaire (à supprimer après)
New-NetFirewallRule -DisplayName "TEMP-Debug" -Direction Inbound -Action Allow -Profile Any
```

### 35.3 Journalisation (indispensable en dépannage réseau)

```powershell
# Activer le log des paquets rejetés (par profil)
Set-NetFirewallProfile -Profile Domain, Private, Public -LogBlocked True -LogMaxSizeKilobytes 16384
# Fichier : C:\Windows\System32\LogFiles\Firewall\pfirewall.log

# Lire les rejets récents
Get-Content C:\Windows\System32\LogFiles\Firewall\pfirewall.log -Tail 20
```

> 💡 **Réflexe dépannage** : « l'application ne répond plus sur le réseau » → 1) `Test-NetConnection`, 2) `pfirewall.log`, 3) règle GPO. Dans 50 % des cas, c'est le pare-feu.

---

## 36. Réseau : Wi-Fi d'entreprise 802.1X (WPA2/WPA3-Enterprise)

### 36.1 Architecture

```
Client Windows 11
   → Point d'accès (authenticator)
   → Serveur RADIUS (NPS / ISE / ClearPass)
   → Active Directory / Entra (annuaire)
   Méthode EAP : PEAP-MSCHAPv2 (identifiants) ou EAP-TLS (certificats, recommandé)
```

| Méthode EAP | Sécurité | Contrainte |
|---|---|---|
| PEAP + MSCHAPv2 | Correcte | Mots de passe (phishing possible) |
| **EAP-TLS** | **Forte** | PKI : certificat machine et/ou utilisateur à déployer |

**Recommandation** : EAP-TLS avec certificats **machine** (auto-enrollment AD CS ou Intune/SCEP) : l'utilisateur n'a rien à saisir, le poste s'authentifie avant l'ouverture de session (utile pour les GPO au démarrage).

### 36.2 Déployer le profil Wi-Fi par GPO

`Configuration ordinateur\Stratégies\Paramètres Windows\Paramètres de sécurité\Stratégies de réseau sans fil` → créer un profil :

- SSID : `ENTREPRISE` (ne pas diffuser le SSID invité ici)
- Sécurité : WPA2-Entreprise (ou WPA3-Entreprise si le parc AP le supporte)
- EAP : Microsoft Smart Card ou EAP-TLS
- Cocher « Authentification ordinateur » pour le pré-logon

### 36.3 Déployer par Intune

Intune → Appareils → Profils de configuration → **Wi-Fi** :

- Type : Entreprise
- EAP : EAP-TLS, certificat client SCEP
- Serveurs RADIUS de confiance (noms des certificats du NPS)

### 36.4 Commandes utiles côté client

```powershell
# Profils Wi-Fi mémorisés
netsh wlan show profiles

# Détail d'un profil (dont la méthode EAP)
netsh wlan show profile name="ENTREPRISE" key=clear

# Exporter / importer un profil (déploiement manuel / dépannage)
netsh wlan export profile name="ENTREPRISE" folder="C:\Admin" key=clear
netsh wlan add profile filename="C:\Admin\Wi-Fi-ENTREPRISE.xml"

# Voir l'état de la connexion et le chiffrement négocié
netsh wlan show interfaces
```

---

## 37. Réseau : VPN — Always On VPN (intro), profils, dépannage

### 37.1 Always On VPN : principe

Successeur de DirectAccess (déprécié) : VPN **IKEv2/SSTP** natif Windows, déployé par **profil XML via Intune/GPO**, avec deux tunnels :

| Tunnel | Usage | Authentification |
|---|---|---|
| **Device Tunnel** | Connectivité machine avant logon (GPO, Intune) | Certificat machine |
| **User Tunnel** | Accès utilisateur | Certificat utilisateur ou EAP |

### 37.2 Infrastructure minimale

- Serveur **RRAS** (Windows Server) avec IKEv2, ou solution tierce compatible ;
- **PKI** (AD CS) : certificats machine + serveur (SAN = nom public) ;
- **NPS** pour l'authentification RADIUS (optionnel selon l'architecture) ;
- Publication : le serveur VPN doit être joignable depuis Internet (443/500/4500).

### 37.3 Profil XML (extrait simplifié)

