---
id: collect-261001-cisco/cisco/reseau-avec-gns3-comment-ajouter-un-switch-cisco-2
title: "reseau-avec-gns3-comment-ajouter-un-switch-cisco"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/reseau-avec-gns3-comment-ajouter-un-switch-cisco.md
source_anchor: ""
source_lines: [138, 142]
sha256: a3437af4674d1ad50209ea02103490abc57c42e3782611a381f16ce4b612d1c6
---

# reseau-avec-gns3-comment-ajouter-un-switch-cisco

Effectivement, c’est toujours le même soucis, trouver les images cisco. En effet, dans le tuto on passe d’un .gns3a de 4Ko à un fichier à importer de 92Mo mais rien n’indique, sauf erreur de ma part, comment l’obtenir. Le bouton download nous envoie sur le site de Cisco, mais même avec un compte, celui ne propose pas de téléchargeme,t, et donc l’importation reste impossible…

Bonjour, existe-il une solution pour obtenir l’image vios_l2 ? Ce n’est pas mentionné dans l’article

Bonjour, il faut acheter CML sur le site Cisco. Ensuite télécharger le .zip qui contient tous les fichiers utiles à la créations d’environnement virtuels Cisco. Le IOSvl2 est dedans. C’est la façon légale de se le procurer si on ne veut pas télécharger des qcow2 bizarre n’importe où
