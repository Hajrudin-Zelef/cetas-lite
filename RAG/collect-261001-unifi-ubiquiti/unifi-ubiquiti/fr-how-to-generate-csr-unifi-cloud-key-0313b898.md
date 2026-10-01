---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/fr-how-to-generate-csr-unifi-cloud-key-0313b898
title: "fr-how-to-generate-csr-unifi-cloud-key-0313b898"
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["California", "United States"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/fr-how-to-generate-csr-unifi-cloud-key-0313b898.md
source_anchor: ""
source_lines: [1, 49]
sha256: 6126926659f411192766c2e3db15f73b5cfaaa3cd138bbb218d06f21ae2d798b
---

# fr-how-to-generate-csr-unifi-cloud-key-0313b898

Ce tutoriel explique comment générer une CSR (Certificate Signing Request) sur un Ubiquiti UniFi Cloud Key. Le Cloud Key ne dispose pas d’un générateur de CSR intégré dans son interface web, vous devrez donc créer la CSR via SSH en utilisant OpenSSL. Ces étapes s’appliquent aux Cloud Key Gen1 et Gen2/Gen2+, ainsi qu’aux installations autohébergées de l’application UniFi Network fonctionnant sous Linux.
Générer une CSR sur UniFi Cloud Key
Si vous avez déjà généré votre CSR, passez directement aux instructions d’installation SSL pour UniFi Cloud Key.
Vous pouvez également utiliser notre Générateur de CSR pour créer la CSR automatiquement plutôt que d’exécuter des commandes OpenSSL sur l’appareil.
Prérequis
- Un accès SSH à votre Cloud Key (activez SSH dans les paramètres UniFi OS, ou connectez-vous directement via l’adresse IP de l’appareil sur le port 22).
- Une connaissance de base des commandes de terminal.
- OpenSSL installé (préinstallé sur le firmware du Cloud Key).
Étape 1 : Connectez-vous au Cloud Key via SSH
Ouvrez un terminal (ou PuTTY sous Windows) et connectez-vous à votre Cloud Key :
ssh root@192.168.1.xCopied!
Remplacez 192.168.1.x par l’adresse IP de votre Cloud Key. Sur les appareils UniFi OS (Gen2/Gen2+), connectez-vous avec les identifiants définis lors de la configuration initiale.
Étape 2 : Sauvegardez les fichiers de certificat existants (facultatif mais recommandé)
Avant de générer de nouveaux fichiers, créez une sauvegarde de votre répertoire de certificats actuel :
cp -r /etc/ssl/private ~/cloudkey-cert-backupCopied!
Cela préserve votre clé et votre certificat existants afin de pouvoir les restaurer si nécessaire. Cette sauvegarde compte ici : sur un Cloud Key Gen1, le répertoire contient déjà le certificat et la clé que l’appareil utilise (cloudkey.crt et cloudkey.key), et la commande de l’étape 3 écrase la nouvelle clé sur cloudkey.key. Tant que vous n’avez pas installé le certificat émis, le certificat présent sur l’appareil et la clé qui l’accompagne ne correspondent plus, et restaurer cette sauvegarde est le seul moyen de remettre l’appareil dans son état initial.
Remarque pour Gen2 et Gen2+ : sur les appareils fonctionnant sous UniFi OS, le certificat et la clé actifs sont unifi-core.crt et unifi-core.key dans /data/unifi-core/config/, et non dans /etc/ssl/private. Sauvegardez plutôt ce répertoire, et générez la demande dans un chemin de votre choix.
Étape 3 : Générez la CSR et la clé privée
Exécutez la commande OpenSSL suivante pour créer une clé privée RSA de 2048 bits et une CSR en une seule étape. Remplacez chaque valeur d’exemple par vos propres informations :
openssl req -new -newkey rsa:2048 -nodes -sha256 
  -subj "/C=US/ST=California/L=San Jose/O=Your Company LLC/CN=unifi.yourdomain.com" 
  -addext "subjectAltName=DNS:unifi.yourdomain.com" 
  -keyout /etc/ssl/private/cloudkey.key 
  -out /etc/ssl/private/cloudkey.csrCopied!
Remarque : l’option -addext nécessite OpenSSL 1.1.1 ou une version ultérieure (fournie sur les Cloud Key Gen2/Gen2+ et sur les appareils Gen1 exécutant le firmware 1.1.6+). Si votre appareil dispose d’une version plus ancienne d’OpenSSL, vous pouvez omettre -addext et ajouter le SAN via un fichier de configuration OpenSSL, ou utiliser le Générateur de CSR. Sur OpenSSL 3.x, vous pouvez utiliser -noenc au lieu de -nodes ; les deux produisent une clé privée non chiffrée.
Référence des champs de la CSR
Chaque champ de la chaîne -subj correspond à un attribut du certificat. Renseignez-les comme suit :
- C : code pays à deux lettres (ISO 3166), par exemple US.
- ST : état ou province (nom complet, non abrégé).
- L : ville ou localité.
- O : le nom légal de votre organisation. Pour les certificats à validation de domaine (DV), ce champ est facultatif.
- CN : le nom de domaine complet (FQDN) que vous utiliserez pour accéder au contrôleur Cloud Key, par exemple unifi.yourdomain.com.
La ligne -addext "subjectAltName=..." définit le champ Subject Alternative Name (SAN). Les navigateurs modernes exigent que le nom d’hôte apparaisse dans le SAN, et non uniquement dans le CN. Si vous devez couvrir plusieurs noms d’hôte, séparez-les par des virgules :
-addext "subjectAltName=DNS:unifi.yourdomain.com,DNS:cloudkey.yourdomain.com"Copied!
Étape 4 : Vérifiez la CSR
Avant de soumettre la CSR, vérifiez que les informations du sujet et le SAN sont corrects :
openssl req -in /etc/ssl/private/cloudkey.csr -noout -textCopied!
Recherchez la ligne Subject: (elle doit correspondre à vos valeurs -subj) et la section X509v3 Subject Alternative Name (elle doit lister votre FQDN). Vous pouvez également coller le contenu de la CSR dans le Décodeur de CSR de SSL Dragon pour une vérification visuelle.
Étape 5 : Soumettez la CSR à votre autorité de certification
Ouvrez le fichier CSR dans un éditeur de texte ou affichez-le dans le terminal :
cat /etc/ssl/private/cloudkey.csrCopied!
Copiez l’intégralité du résultat, y compris les lignes -----BEGIN CERTIFICATE REQUEST----- et -----END CERTIFICATE REQUEST-----. Collez-la dans le champ CSR lors du processus de commande du certificat SSL auprès de votre CA.
Conservez le fichier de clé privée (/etc/ssl/private/cloudkey.key) sur le Cloud Key. Vous en aurez besoin plus tard lorsque vous installerez le certificat SSL dans le magasin de clés UniFi.
Étape 6 : Sécurisez la clé privée
Restreignez le fichier de clé afin que seul root puisse le lire :
chmod 600 /etc/ssl/private/cloudkey.keyCopied!
Ne partagez et ne transférez jamais la clé privée via un canal non chiffré. Si vous pensez qu’elle a été compromise, générez une nouvelle CSR et faites réémettre le certificat.
Economisez 10% sur les certificats SSL en commandant aujourd’hui!
Émission rapide, cryptage puissant, confiance de 99,99 % du navigateur, assistance dédiée et garantie de remboursement de 25 jours. Code de coupon: SAVE10
