---
id: collect-261001-rattrapage/rattrapage/win11-guide-5
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "attention", "ethernet", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [388, 573]
sha256: 77439f8c5ece09b371a5167a1e80841bdbe0e3b28f85f647bf13cf7bbddb69ee
---

# Windows 11 en entreprise — Guide technique ultra-complet

1. Media Creation Tool → « Créer un support d'installation ».
2. Ou en ligne de commande avec `diskpart` + copie des fichiers (clé en FAT32 pour boot UEFI ; si `install.wim` > 4 Go, splitter) :

```powershell
# Préparer la clé (ATTENTION : remplacez N par le bon numéro de disque !)
diskpart
list disk
select disk N
clean
convert gpt
create partition primary
format fs=fat32 quick label="WIN11"
assign
exit
```

```powershell
# Si install.wim dépasse 4 Go (limite FAT32), le scinder :
dism /Split-Image /ImageFile:C:\ISO\sources\install.wim /SWMFile:E:\sources\install.swm /FileSize:3800
```

Alternative moderne : **Rufus** (utile pour forcer certains paramètres en lab, mais en entreprise préférez la méthode Microsoft + unattend.xml).

### 6.2 Étapes d'installation propre

1. Boot UEFI sur la clé (désactiver temporairement Secure Boot n'est **pas** nécessaire avec un média officiel signé).
2. Langue → « Installer maintenant ».
3. Clé de produit : saisir ou « Je n'ai pas de clé » (l'activation se fera après, via KMS/ADBA/Entra).
4. Choisir l'édition correspondant à votre licence.
5. Type : **Personnalisé** → supprimer/recréer les partitions (sauf cas de double boot).
6. Laisser l'installation se dérouler, puis OOBE.

### 6.3 Pilotes : ordre d'installation recommandé

1. Chipset (Intel Chipset Device Software / AMD Chipset Drivers)
2. Stockage (NVMe / RAID / RST)
3. Réseau (Ethernet puis Wi-Fi)
4. Graphiques (Intel/AMD/NVIDIA — version **DCH**)
5. Audio
6. Périphériques (lecteur d'empreinte, etc.)
7. Utilitaires constructeur (Dell Command Update, Lenovo System Update…)

```powershell
# Exporter les pilotes d'un poste de référence (pour MDT, section 11)
Export-WindowsDriver -Online -Destination C:\Drivers\Export-Ref

# Installer un pilote en ligne de commande
pnputil /add-driver C:\Drivers\*.inf /subdirs /install
```

---

## 7. Mise à niveau depuis Windows 10 : chemins, outils, pièges

### 7.1 Chemins de mise à niveau supportés

| Source | Cible | Méthode |
|---|---|---|
| Windows 10 22H2 Pro/Entreprise | Windows 11 24H2 même édition | Windows Update, ISO (setup.exe), WUfB, MDT (refresh), Intune (feature update) |
| Windows 10 Famille | Windows 11 Pro | Changement de clé + MàJ |

> ⚠️ **Pas de mise à niveau** 32 bits → 64 bits, ni d'une langue vers une autre sans réinstallation. Et Windows 10 doit être en 22H2 avec les dernières MàJ cumulatives.

### 7.2 Lancement via setup.exe (ISO montée)

```powershell
# Mise à niveau silencieuse conservant fichiers et applications
E:\setup.exe /auto upgrade /quiet /noreboot /showoobe none /compat ignorewarning

# Options utiles :
# /auto upgrade            : tout conserver
# /quiet                  : sans interface
# /noreboot               : ne pas redémarrer (planifier le reboot)
# /dynamicupdate enable   : télécharger les dernières MàJ pendant l'install
# /compat ignorewarning   : ignorer les avertissements (pas les erreurs bloquantes)
```

### 7.3 Pièges classiques de la mise à niveau

| Piège | Symptôme | Parade |
|---|---|---|
| Espace disque insuffisant | Échec à ~30 %, code 0x80070070 | 20 Go libres minimum ; nettoyer avec cleanmgr |
| Pilote incompatible | Rollback, 0xC1900101 | Mettre à jour BIOS/pilotes avant ; débrancher les périphériques USB non essentiels |
| Antivirus tiers | Blocage setup | Désinstaller (pas seulement désactiver) avant |
| Chiffrement tiers | Échec | Suspendre/déchiffrer avant |
| Applications incompatibles | Avertissement compat | Tester sur lot pilote ; utiliser Upgrade Readiness |
| TPM désactivé | « Ce PC ne peut pas exécuter Windows 11 » | Activer Intel PTT / AMD fTPM dans l'UEFI |

### 7.4 Journaux de mise à niveau (où chercher quand ça échoue)

| Fichier | Rôle |
|---|---|
| `C:\$Windows.~BT\Sources\Panther\setupact.log` | Journal principal du setup |
| `C:\$Windows.~BT\Sources\Panther\setuperr.log` | Erreurs uniquement |
| `C:\$Windows.~BT\Sources\Rollback\setupact.log` | En cas de rollback |
| `C:\Windows\Logs\MoSetup\BlueBox.log` | Lancement via Windows Update |
| `setupdiag.exe` (outil Microsoft) | Diagnostic automatisé : `SetupDiag.exe /Output:C:\Logs\result.xml` |

```powershell
# Télécharger et lancer SetupDiag (nécessite .NET)
# https://aka.ms/SetupDiag
.\SetupDiag.exe /Output:C:\Admin\SetupDiag_Results.log
```

---

## 8. Installation automatisée : unattend.xml, clés de produit, setupconfig

### 8.1 Principe des passes d'installation

Le fichier `unattend.xml` pilote le setup via des « passes » : `windowsPE`, `offlineServicing`, `specialize`, `oobeSystem`, etc. L'outil graphique officiel est le **Windows System Image Manager (WSIM)**, inclus dans l'ADK.

### 8.2 Exemple unattend.xml minimal (lab — à adapter)

```xml
<?xml version="1.0" encoding="utf-8"?>
<unattend xmlns="urn:schemas-microsoft-com:unattend">
  <settings pass="windowsPE">
    <component name="Microsoft-Windows-International-Core-WinPE"
               processorArchitecture="amd64"
               publicKeyToken="31bf3856ad364e35" language="neutral" versionScope="nonSxS">
      <SetupUILanguage><UILanguage>fr-FR</UILanguage></SetupUILanguage>
      <InputLocale>040c:0000040c</InputLocale>
      <SystemLocale>fr-FR</SystemLocale>
      <UILanguage>fr-FR</UILanguage>
      <UserLocale>fr-FR</UserLocale>
    </component>
    <component name="Microsoft-Windows-Setup"
               processorArchitecture="amd64"
               publicKeyToken="31bf3856ad364e35" language="neutral" versionScope="nonSxS">
      <UserData>
        <ProductKey>
          <!-- Clé générique d'installation (GVLK) — N'ACTIVE PAS, sert au setup -->
          <Key>XXXXX-XXXXX-XXXXX-XXXXX-XXXXX</Key>
          <WillShowUI>Never</WillShowUI>
        </ProductKey>
        <AcceptEula>true</AcceptEula>
      </UserData>
    </component>
  </settings>
  <settings pass="oobeSystem">
    <component name="Microsoft-Windows-Shell-Setup"
               processorArchitecture="amd64"
               publicKeyToken="31bf3856ad364e35" language="neutral" versionScope="nonSxS">
      <OOBE>
        <HideEULAPage>true</HideEULAPage>
        <HideOEMRegistrationScreen>true</HideOEMRegistrationScreen>
        <HideOnlineAccountScreens>true</HideOnlineAccountScreens>
        <HideWirelessSetupInOOBE>true</HideWirelessSetupInOOBE>
        <ProtectYourPC>3</ProtectYourPC>
      </OOBE>
      <UserAccounts>
        <LocalAccounts>
          <LocalAccount wcm:action="add">
            <Name>tech</Name>
            <Group>Administrators</Group>
            <Password><Value>U2VjcmV0UGFzc3dvcmQxMjM=</Value><PlainText>false</PlainText></Password>
          </LocalAccount>
        </LocalAccounts>
      </UserAccounts>
    </component>
  </settings>
</unattend>
```

> ⚠️ Le mot de passe ci-dessus est un exemple encodé en base64 (`U2VjcmV0UGFzc3dvcmQxMjM=` = « SecretPassword123 » fictif). **Ne commitez jamais un unattend.xml contenant un vrai mot de passe** dans un dépôt : préférez créer le compte admin via une task sequence MDT ou une stratégie LAPS (section 34).

### 8.3 Clés de produit : les 4 modèles

| Modèle | Usage | Remarque |
|---|---|---|
| OEM | PC livrés avec Windows | Liée au firmware (OA3), réinstallation auto |
| Retail | Boîte / Microsoft Store | Transférable |
| MAK (Multiple Activation Key) | Volume, activation en ligne | Nombre d'activations limité, à réserver aux postes isolés |
| KMS / GVLK | Volume, activation locale | **Standard entreprise** : serveur KMS ou ADBA |

```powershell
# Installer une clé générique de volume (GVLK) — exemple fictif de format
slmgr.vbs /ipk XXXXX-XXXXX-XXXXX-XXXXX-XXXXX
# Pointer vers le KMS interne
slmgr.vbs /skms kms.entreprise.local:1688
slmgr.vbs /ato
# Vérifier
slmgr.vbs /dlv
```

