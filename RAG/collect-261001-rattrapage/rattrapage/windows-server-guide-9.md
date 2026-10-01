---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-9
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-26-09"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1427, 1605]
sha256: c7c7ed1a7034f93b78ff442a1188340db41b6ac92d5c95de5bcf79db0974ec61
---

# Windows Server en entreprise — Guide technique ultra-complet

```powershell
# Dupliquer un modèle (ex. "Utilisateur" -> "Utilisateur-Entreprise" avec auto-enrollment)
# Via la console : certsrv.msc → Modèles de certificats → Gérer → Dupliquer "Utilisateur"
# Onglet Sécurité : ajouter le groupe + cocher Inscrire + Inscription automatique
# Onglet Général : cocher "Publier le certificat dans Active Directory" si besoin

# Publier le modèle sur l'AC :
Add-CATemplate -Name "Utilisateur-Entreprise" -Force
Get-CATemplate | Format-Table Name -AutoSize
```

**GPO d'auto-enrollment :**

```
Configuration ordinateur → Stratégies → Paramètres Windows → Paramètres de sécurité →
Stratégies de clés publiques → Client des services de certificats - Inscription automatique
  • Modèle de configuration : Activé
  • Cocher : Renouveler les certificats expirés, mettre à jour les certificats...
```

**Cas d'usage internes typiques :**
| Besoin | Modèle | Auto-enrollment |
|---|---|---|
| LDAPS (AD chiffré) | Kerberos Authentication / Domain Controller | Oui, sur les DC |
| 802.1X Wi-Fi/filaire | Ordinateur + Utilisateur | Oui |
| RDP chiffré | Ordinateur | Oui |
| Serveurs web internes (IIS, WAC) | Serveur Web | Manuel ou auto |
| Signature e-mail (S/MIME) | Utilisateur | Oui |

---

## 50. Renouvellement et révocation (CRL, OCSP)

**Renouvellement :** avec l'auto-enrollment (§49), les certificats se renouvellent seuls avant expiration. Sans auto-enrollment : demande manuelle via `certlm.msc` / `certmgr.msc` ou le portail `http://srv-ca/certsrv`.

**Révocation :**

```powershell
# Révoquer un certificat (ex. départ d'un salarié, clé compromise)
certutil -revoke <NuméroDeSérie> 1   # 1 = keyCompromise ; 0 = unspecified
# Publier immédiatement la CRL
certutil -crl

# Vérifier qu'un certificat est révoqué
certutil -verify -urlfetch C:\temp\cert.cer
```

**Codes de révocation :** 0 = non spécifié, 1 = clé compromise, 3 = affiliation changée, 5 = cessation d'activité.

**OCSP (Online Responder)** : pour les gros parcs, installer le répondeur en ligne (service de rôle `ADCS-Online-Cert`) afin d'éviter le téléchargement de CRL volumineuses.

> **Point de vigilance :** si la CRL expire ou est injoignable, les clients **rejettent** les certificats (hard-fail par défaut sur beaucoup d'applis). Surveiller la date d'expiration de la CRL comme un certificat.

---

## 51. Dépannage PKI

| Symptôme | Cause probable | Action |
|---|---|---|
| Auto-enrollment ne distribue rien | Modèle non publié / droits insuffisants | `Add-CATemplate` ; onglet Sécurité du modèle : groupe = Inscrire + Inscription auto |
| "Le serveur RPC n'est pas disponible" | DCOM bloqué | Pare-feu : autoriser DCOM/RPC entre client et AC |
| Certificat rejeté : révocation inconnue | CRL injoignable/expirée | `certutil -urlfetch -verify` ; republier la CRL ; vérifier le HTTP |
| La demande reste "en attente" | Approbation manuelle requise sur le modèle | Console AC → Demandes en attente → Délivrer ; ou passer le modèle en auto |
| `certreq` échoue | Modèle mal orthographié | `certutil -template` pour lister les noms exacts |
| LDAPS ne fonctionne pas | Certificat DC sans le bon modèle | Modèle "Domain Controller Authentication" + redémarrer le DC |

```powershell
# Diagnostic express PKI
certutil -ping -config "SRV-CA-01\ENTREPRISE-Issuing-CA-01"   # l'AC répond-elle ?
certutil -cainfo                                              # infos de l'AC locale
Get-ChildItem Cert:\LocalMachine\My | Format-Table Subject, NotAfter, Issuer  # certificats machine
```

---

---

## 52. Sauvegarde : Windows Server Backup

Windows Server Backup (WSB) = sauvegarde locale correcte pour petits parcs ou appoint. Pas de déduplication avancée, pas de console centrale native.

```powershell
Install-WindowsFeature -Name Windows-Server-Backup -IncludeManagementTools

# Planifier une sauvegarde complète quotidienne à 22h vers un disque dédié
wbadmin enable backup -addtarget:E: -schedule:22:00 -allCritical -systemState -vssFull -quiet

# Sauvegarde bare metal (récupération complète sur matériel différent)
wbadmin start backup -backupTarget:E: -allCritical -systemState -bareMetalRecovery -quiet

# Sauvegarde immédiate d'un volume
wbadmin start backup -backupTarget:\\SRV-NAS-01\sav$ -include:D: -quiet

# Lister les sauvegardes disponibles
wbadmin get versions
wbadmin get disks
```

**Bonnes pratiques WSB :**
- Disque **dédié** (jamais le volume sauvegardé), de préférence **hors site** en rotation.
- `-vssFull` tronque les journaux applicatifs (SQL/Exchange) ; `-vssCopy` ne tronque pas.
- Tester la restauration (voir §55). Toujours.

---

## 53. Planifications, bare metal, restauration

**Planification type :**

| Quoi | Fréquence | Destination |
|---|---|---|
| État système + volumes critiques | Quotidien 22h | Disque USB en rotation (3 disques) |
| Bare Metal Recovery | Hebdo (dimanche) | NAS + disque externe |
| VM Hyper-V (hôte) | Quotidien | Stockage dédié |

```powershell
# Restaurer un fichier précis
wbadmin start recovery -version:09/26/2026-22:00 -itemType:File `
  -items:D:\Partages\Compta\budget.xlsx -recoveryTarget:D:\Restauration -quiet

# Restaurer un volume complet
wbadmin start recovery -version:09/26/2026-22:00 -itemType:Volume `
  -items:D: -recoveryTarget:D: -quiet

# État système (restauration AD sur un DC : voir mode DSRM)
wbadmin start systemstaterecovery -version:09/26/2026-22:00 -quiet

# Bare metal : démarrer sur le support d'installation → Réparer → Restauration d'image système
```

> **Restauration d'un contrôleur de domaine :** démarrer en **mode DSRM** (F8), restaurer l'état système, puis choisir restauration **autoritative** (`ntdsutil`) ou non selon le scénario. En pratique : si un autre DC sain existe, préférer **reconstruire** le DC plutôt que restaurer.

---

## 54. Stratégie 3-2-1 et rotation des supports

**3-2-1 :** 3 copies des données, sur 2 supports différents, dont 1 hors site.

| Support | Rôle | Rotation |
|---|---|---|
| Disque interne/dédié | Sauvegarde quotidienne rapide | Permanent |
| Disques USB chiffrés (×3) | Rotation hebdomadaire | LUN → MAR → MER, 1 hors site |
| NAS / cloud | Copie distante | Quotidien/hebdo |

```powershell
# Chiffrer le disque de sauvegarde (BitLocker) — obligatoire si le disque sort du site
Enable-BitLocker -MountPoint "E:" -PasswordProtector   # mot de passe FICTIF à la demande
# Ou via manage-bde :
manage-bde -on E: -pw
```

**Règle d'or :** une sauvegarde dont on n'a jamais testé la restauration est une **hypothèse**, pas une sauvegarde.

---

## 55. Restauration : scénarios pas à pas

**Scénario A — fichier supprimé par un utilisateur :**
1. Clic droit sur le dossier parent → **Propriétés → Versions précédentes** (clichés instantanés VSS à activer, voir ci-dessous).
2. Ou `wbadmin start recovery` (§53).

```powershell
# Activer les clichés instantanés (VSS) sur un volume de partages : 2×/jour
vssadmin add shadowstorage /for=D: /on=D: /maxsize=10%
# Planifier via tâche planifiée : wmic shadowcopy call create Volume='D:\'
schtasks /create /tn "VSS D midi" /tr "wmic shadowcopy call create Volume='D:\'" /sc daily /st 12:00 /ru SYSTEM
schtasks /create /tn "VSS D soir" /tr "wmic shadowcopy call create Volume='D:\'" /sc daily /st 18:00 /ru SYSTEM
```

**Scénario B — serveur HS, matériel OK :** restauration bare metal depuis le support d'installation.

**Scénario C — ransomware :** ne PAS restaurer sur l'infrastructure compromise. Isoler, reconstruire sur sain, restaurer depuis la copie **hors site** (la seule dont on est sûr qu'elle n'est pas chiffrée), changer tous les mots de passe.

---

## 56. Veeam et alternatives : quand passer au niveau supérieur

WSB suffit pour : 1–3 serveurs, budget nul, besoins simples.

