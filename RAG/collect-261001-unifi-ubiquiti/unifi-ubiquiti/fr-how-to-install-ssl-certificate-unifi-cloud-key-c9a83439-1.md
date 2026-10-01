---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439-1
title: "fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439.md
source_anchor: ""
source_lines: [1, 50]
sha256: b3c803854d9be8522e8383bfa1bd1ed8af88c5e1511753b5df75ec6664183c37
---

# fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439

Ce tutoriel vous donne des instructions étape par étape sur comment installer un certificat SSL sur un UniFi Cloud Key. Le Cloud Key exécute l’application UniFi Network au-dessus d’un keystore Java, l’installation se fait donc via SSH : vous créez un ensemble PKCS#12 à partir de votre certificat, de votre clé privée et de la chaîne CA, puis vous l’importez dans le keystore du Cloud Key à l’aide de l’utilitaire Java keytool et vous redémarrez le service UniFi.
Générer un code CSR sur UniFi Cloud Key
La demande de signature de certificat, ou CSR, est un bloc de texte encodé qui contient vos coordonnées, telles que votre domaine et votre organisation. Chaque demandeur doit générer une CSR et l’envoyer à l’autorité de certification (CA) pour validation avant qu’un certificat puisse être émis. Pour un Cloud Key, le nom commun dans la CSR doit être le nom d’hôte que vous utilisez pour accéder au contrôleur, par exemple unifi.yoursite.com.
Vous avez deux options :
- Utilisez notre générateur de CSR pour créer la CSR automatiquement.
- Suivez notre tutoriel étape par étape sur comment générer une CSR sur UniFi Cloud Key.
Vous pouvez ouvrir votre fichier CSR avec n’importe quel éditeur de texte, comme le Bloc-notes, et le coller lors de la commande de certificat auprès de votre fournisseur SSL. Conservez la clé privée correspondante en lieu sûr : vous en aurez à nouveau besoin à l’étape 3 pour créer l’ensemble PKCS#12. Une fois que la CA valide votre demande et vous envoie les fichiers SSL par e-mail, poursuivez avec l’installation ci-dessous.
Installer un certificat SSL sur UniFi Cloud Key
Le Cloud Key sert le HTTPS à partir d’un keystore Java situé à /usr/lib/unifi/data/keystore. Par défaut, ce keystore contient un certificat auto-signé sous l’alias unifi, protégé par le mot de passe aircontrolenterprise. Pour installer votre propre certificat, vous remplacez cette entrée : vous regroupez vos fichiers dans une archive PKCS#12, vous l’importez sous le même alias, puis vous redémarrez le service UniFi. Vous travaillerez via SSH, connectez-vous donc d’abord au Cloud Key (l’utilisateur SSH par défaut est root sur Gen1, ou le compte administrateur que vous avez défini lors de l’adoption sur Gen2).
Étape 1 : Préparez vos fichiers SSL
Vérifiez votre boîte de réception et téléchargez l’archive ZIP contenant votre certificat. Extrayez-la. Selon votre fournisseur SSL, vous devriez avoir tout ou partie de ces fichiers :
- Le certificat SSL principal (votre certificat de serveur).
- Le certificat SSL intermédiaire.
- Le certificat SSL racine.
- Le CA bundle, un fichier unique qui contient déjà les certificats racine et intermédiaire.
Vous avez également besoin de la clé privée que vous avez générée avec la CSR. Si vous avez utilisé notre générateur de CSR, la clé était proposée au téléchargement à ce moment-là.
Étape 2 : Copiez vos fichiers SSL sur le Cloud Key
Téléversez vos fichiers sur le Cloud Key via SCP ou SFTP et placez-les dans /etc/ssl/private/. Utilisez ces noms afin qu’ils correspondent aux commandes des étapes suivantes :
- Certificat principal sous /etc/ssl/private/cloudkey.crt
- Clé privée sous /etc/ssl/private/cloudkey.key
- CA bundle (racine et intermédiaire) sous /etc/ssl/private/yourcaname.crt
Remarque : chaque fichier PEM doit se terminer par un saut de ligne après la ligne de clôture. La dernière ligne du certificat est exactement cinq tirets, les mots END CERTIFICATE, puis cinq tirets (—–END CERTIFICATE—–) ; assurez-vous qu’il y a un saut de ligne après. L’absence de saut de ligne final est une cause fréquente d’erreurs d’importation.
Étape 3 : Regroupez votre certificat dans un fichier PKCS#12
Combinez votre clé privée, votre certificat principal et le CA bundle en un seul fichier PKCS#12. L’alias doit être unifi, car c’est l’alias attendu par le keystore UniFi. Remplacez enteryourpassword par un mot de passe de votre choix et souvenez-vous-en : vous transmettrez cette même valeur à keytool à l’étape 4 comme mot de passe source.
openssl pkcs12 -export -in /etc/ssl/private/cloudkey.crt -inkey /etc/ssl/private/cloudkey.key -out /etc/ssl/private/cloudkey.p12 -name unifi -CAfile /etc/ssl/private/yourcaname.crt -caname root -password pass:enteryourpasswordCopied!
Cela crée le fichier /etc/ssl/private/cloudkey.p12, contenant la clé, votre certificat et la chaîne, le tout sous l’alias unifi.
Étape 4 : Importez le fichier PKCS#12 dans le keystore du Cloud Key
Importez le fichier PKCS#12 dans le keystore existant situé à /usr/lib/unifi/data/keystore. Le keystore de destination utilise déjà le mot de passe aircontrolenterprise, donc les mots de passe du magasin et de la clé de destination doivent rester aircontrolenterprise. Le mot de passe source est celui que vous avez défini sur le fichier PKCS#12 à l’étape 3.
keytool -importkeystore -deststorepass aircontrolenterprise -destkeypass aircontrolenterprise -destkeystore /usr/lib/unifi/data/keystore -srckeystore /etc/ssl/private/cloudkey.p12 -srcstoretype PKCS12 -srcstorepass enteryourpassword -alias unifiCopied!
Si keytool signale que l’alias unifi existe déjà, il fait référence à l’entrée auto-signée par défaut. Confirmez l’écrasement lorsque vous y êtes invité, ou supprimez d’abord l’ancienne entrée avec la commande ci-dessous puis relancez l’importation :
keytool -delete -alias unifi -keystore /usr/lib/unifi/data/keystore -storepass aircontrolenterpriseCopied!
Important : ne changez pas le mot de passe du keystore pour une valeur personnalisée. Le service UniFi est codé en dur pour ouvrir le keystore avec aircontrolenterprise ; si vous le modifiez, le contrôleur ne pourra plus lire le certificat et le HTTPS échouera.
Étape 5 : Définissez les permissions et supprimez les fichiers temporaires
Renforcez la propriété et les permissions des fichiers dans /etc/ssl/private/ pour que la clé privée ne soit pas lisible par tout le monde, puis supprimez les fichiers intermédiaires dont vous n’avez plus besoin. Exécutez ces commandes une par une :
chown root:ssl-cert /etc/ssl/private/*
chmod 640 /etc/ssl/private/*
rm /etc/ssl/private/cloudkey.p12Copied!
Le certificat se trouve désormais dans le keystore, donc le fichier PKCS#12 n’est plus nécessaire. Si vous avez également téléversé la CSR ou le fichier CA uniquement pour ce processus, vous pouvez aussi les supprimer :
rm /etc/ssl/private/cloudkey.csr
rm /etc/ssl/private/yourcaname.crtCopied!
Astuce : avant de redémarrer, vous pouvez sauvegarder le keystore afin de pouvoir revenir en arrière si nécessaire. Depuis /usr/lib/unifi/data/, archivez-le avec tar -cvf keystore-backup.tar keystore.
Étape 6 : Faites correspondre le nom d’hôte du contrôleur à votre certificat
Pour que les navigateurs fassent confiance au certificat sans avertissement de non-correspondance de nom, définissez le nom d’hôte ou l’IP du contrôleur sur le nom figurant dans votre certificat. Dans l’application UniFi Network, allez dans Paramètres > Système (versions plus anciennes : Paramètres > Contrôleur) et définissez le nom d’hôte/IP du contrôleur sur le nom commun de votre certificat, par exemple unifi.yoursite.com. Assurez-vous que ce nom d’hôte pointe bien vers le Cloud Key dans votre DNS.
Étape 7 : Redémarrez le service UniFi
Le nouveau certificat n’est lu qu’au démarrage du service UniFi, redémarrez-le donc pour charger votre certificat :
service unifi restartCopied!
Sur les firmwares plus récents du Cloud Key qui utilisent systemd, la commande équivalente est :
systemctl restart unifiCopied!
Patientez une minute que le service redémarre, puis rechargez le contrôleur dans votre navigateur. Félicitations, votre certificat SSL est maintenant installé sur votre UniFi Cloud Key.
Testez votre installation SSL
