---
id: collect-261001-meraki/meraki/docs-documentation-wi-fi-meraki-configuration-9903c4bc
title: "docs-documentation-wi-fi-meraki-configuration-9903c4bc"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/docs-documentation-wi-fi-meraki-configuration-9903c4bc.md
source_anchor: ""
source_lines: [1, 48]
sha256: c979d02ea54e4cda8ddd5f38f1c7c26b2ca19160b89a46c4f48ac0c3d170b981
---

# docs-documentation-wi-fi-meraki-configuration-9903c4bc

Voici comment configurer vos point d’accès Meraki pour utiliser le système de connexion d’Eye-Intelligence.
Pour commencer, allez dans Wi-Fi / Contrôle d’accès
Sélectionnez Ouvert (sans chiffrement) pour Exigences de connexion, et Mon server RADIUS pour Portail captif
Dans les paramètres de Radius, assurez-vous que le paramètre Compatibilité RADIUS est activée
Utilisez radius.eyeinwireless.com pour les Hôtes RADIUS et utilisez les ports (standards) 1812 pour la page d’introduction et 1813 pour la gestion des comptes.
Contactez-nous pour obtenir la Clé secrète…
Note : si l’option pour la comptabilité RADIUS n’est pas disponible, vous devez contacter le support Meraki pour l’activer.
Nous vous recommandons de Refuser l’accès en Politique de basculement et de conserver le paramètre Ordre de priorité strict pour la Politique d’équilibrage de la charge
Activez le Jardin clos, et ajoutez les URL suivantes :
 api.eyeinwireless.com,
azure.eyeinwireless.com
eyeinlivestorage.blob.core.windows.net
eye-in.com
Connexion sociales
Pour autoriser la connexion avec l’un des réseaux sociaux, ajoutez les URLs suivantes :
facebook.com
*.facebook.com
*.fbcdn.net
facebook.net
*.facebook.net
*.akamaihd.net
twitter.com
api.twitter.com
*.twimg.com
*.akamaihd.net
google.com
www.google.com
clients1.google.com
accounts.youtube.com
accounts.google.*
ssl.gstatic.com
ssl.google-analytics.com
*.googleusercontent.com
*.akamaihd.net
 instagram.com
*.instagram.com
api.instagram.com
d36xtkk24g8jdx.cloudfront.net
www.facebook.com
connect.facebook.net
*.akamaihd.net
Note : vous devrez peut-être contacter le support Meraki pour activer la gestion du Jardin Clos par URL et non seulement par adresse IP. Si un de ces domaine change d’IP, vous n’aurez pas à modifier la configuration de votre réseau.
Utilisez le paramètre Bloquer tout accès jusqu’à ce que la connexion soit achevée pour la Force du portail captif
Configuration du portail captif
Allez dans Wi-Fi / Portail captif
Sélectionnez URL d’accueil personnalisée et entrez l’URL suivante :
https://azure.eyeinwireless.com/?step=prelogin&dealerid=[site]&manufacturer=meraki
Remplacez [site] par l’un des codes de vos sites fournis par Eye-In
