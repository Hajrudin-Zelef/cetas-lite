---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439-2
title: "fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439.md
source_anchor: ""
source_lines: [51, 66]
sha256: 86bf9dd3d3d3862fc860b351521735d7466863d45ce9adaf463416b25ef68b48
---

# fr-how-to-install-ssl-certificate-unifi-cloud-key-c9a83439

Après avoir installé le certificat, vérifiez votre configuration pour détecter d’éventuelles erreurs ou vulnérabilités, comme un certificat intermédiaire manquant ou une chaîne non fiable. Utilisez notre SSL Checker pour une analyse et un rapport instantanés. Comme le contrôleur UniFi écoute généralement sur le port 8443, pointez le vérificateur vers votre adresse complète, par exemple unifi.yoursite.com:8443. Vous pouvez également confirmer le résultat en ligne de commande :
echo | openssl s_client -connect unifi.yoursite.com:8443 -servername unifi.yoursite.com 2>/dev/null | openssl x509 -noout -issuer -datesCopied!
Cela affiche l’émetteur et les dates de validité du certificat, ce qui vous permet de confirmer que le Cloud Key sert bien votre certificat et non l’ancien certificat auto-signé.
Questions fréquentes
Où se trouve le keystore SSL sur un UniFi Cloud Key ?
Le keystore UniFi est le keystore Java situé à /usr/lib/unifi/data/keystore. Il contient le certificat que le contrôleur sert via HTTPS, stocké sous l’alias unifi et protégé par le mot de passe par défaut aircontrolenterprise.
Quel est le mot de passe par défaut du keystore sur UniFi Cloud Key ?
Le mot de passe par défaut du keystore est aircontrolenterprise. Conservez-le tel quel. Le service UniFi est codé en dur pour ouvrir le keystore avec ce mot de passe ; si vous le modifiez, le contrôleur ne pourra plus lire votre certificat. Utilisez-le comme mot de passe du magasin et de la clé de destination lors de l’importation avec keytool.
Pourquoi mon certificat n’apparaît-il pas après son importation ?
Les deux causes les plus fréquentes sont un mot de passe de keystore incorrect et l’absence de redémarrage du service. Vérifiez que vous avez importé dans le keystore situé à /usr/lib/unifi/data/keystore avec le mot de passe aircontrolenterprise et l’alias unifi, puis redémarrez le contrôleur avec service unifi restart. Le Cloud Key ne lit le keystore qu’au démarrage, le nouveau certificat n’est donc chargé qu’après un redémarrage.
Quel alias le certificat doit-il utiliser ?
L’alias doit être unifi. Définissez-le avec l’option -name unifi lorsque vous créez le fichier PKCS#12 dans openssl, et utilisez -alias unifi lors de l’importation avec keytool. Si vous importez sous un alias différent, le contrôleur continuera de servir l’ancien certificat auto-signé.
Comment redémarrer le service UniFi sur le Cloud Key ?
Exécutez service unifi restart via SSH. Sur les firmwares plus récents qui utilisent systemd, exécutez plutôt systemctl restart unifi. Vous pouvez aussi redémarrer l’appareil depuis l’interface web du Cloud Key, mais redémarrer uniquement le service UniFi est plus rapide et suffisant pour charger un nouveau certificat.
Economisez 10% sur les certificats SSL en commandant aujourd’hui!
Émission rapide, cryptage puissant, confiance de 99,99 % du navigateur, assistance dédiée et garantie de remboursement de 25 jours. Code de coupon: SAVE10
