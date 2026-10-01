---
id: collect-261001-meraki/meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-seating-request-notifier-4130cdfa
title: "codeexchange-github-repo-gve-sw-gve-devnet-meraki-seating-request-notifier-4130cdfa"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "parameters"]
source: docs/RAG/collect-261001-meraki/codeexchange-github-repo-gve-sw-gve-devnet-meraki-seating-request-notifier-4130cdfa.md
source_anchor: ""
source_lines: [1, 53]
sha256: 46555abd631ac7500e8f6d5fd4918f5600e0e91edc93a9673082f3a53c31c2a1
---

# codeexchange-github-repo-gve-sw-gve-devnet-meraki-seating-request-notifier-4130cdfa

The purpose of this sample code is to notify restaurant staff members about guests in the entry area that are waiting to be seated.
Thereby, guests are detected via two Meraki MV cameras and a Meraki MT30 button sensor. The detection data is delivered via MQTT. A notification can be displayed on one or more tablets, mobile phones or screens within the restaurant.
Multiple scenarios and statuses are supported: no guests present, new guests arrived, guests have been waiting for quite some time and guests actively asked for seating.
MQTT is a Client-Server publish/subscribe messaging transport protocol. This sample code requires the setup of a locally installed or use of an online MQTT broker that gathers the data from all cameras and sensors and publish it to our sample script.
Popular MQTT brokers are for example: Mosquitto or HiveMQ.
Note: Some online MQTT brokers can involve major delays in data delivery. Delays can strongly affect the quality of this demo.
After the MQTT broker is successfully set up, some configurations in the Meraki Dashboard are required. Follow the MV Sense MQTT Instructions to configure MQTT for both cameras and MT MQTT Setup Guide to configure MQTT for the MT sensor.
Note: For demo purposes, use None as value for the field Security in both cases. Please be aware that it is recommended to use TLS in production setups. Further adaptions of this code are required for the latter.
This sample code allows to optionally narrow the camera and detection view via Meraki MV Privacy Windows and Zones. Follow the Instructions for Privacy Windows or Instructions for MV Zones to set these up.
cd [add name of virtual environment here] 
git clone [add github link here]
Access the downloaded folder:
cd gve_devnet_meraki_seating_request_notifier
Install all dependencies:
pip3 install -r requirements.txt
Define the MQTT broker to use for this script. Therefore, open the app.py file and adapt the following lines 25 and 26:
app.config['MQTT_BROKER_URL'] = "[Fill in ip of locally installed MQTT broker or URL for online MQTT broker]"
app.config['MQTT_BROKER_PORT'] = [Fill in MQTT port]
Run the script by using the command:
python3 app.py
Assuming you kept the default parameters for starting the Flask application, the address to navigate would be: http://localhost:5000/settings
Fill in all settings form fields and save the changes.
Form field descriptions:
MV Camera 1/2 - Serial Number: Serial number of Meraki MV Camera
MV Camera 1/2 - Zone: Use 0 for full screen with or without privacy windows or zone ID for Camera MV zone. Zone ID available under Cameras > [Choose Camera] > Settings Tab > **Sense Tab ** > /merakimv/xxxx-xxxx-xxxx/{zone ID}.
MT30 Button - Mac Address: Mac address of Meraki MT30 sensor in format XX:XX:XX:XX:XX:XX
MT30 Button - Local ID: Local ID of Meraki button sensor - available under Sensors > MONITOR > Sensors > MQTT Broker > MQTT Topics > meraki/v1/mt/{local ID}/....
Reviewing Interval (ms): Milliseconds interval in which received MQTT MV messages are reviewed based on the timestamp of the message.
Notification Interval (ms): Time span between the first and second/escalated notification in milliseconds.
Navigate to http://localhost:5000/ to access the notification dashboard.
Enter the camera view of both cameras and optionally press the button to trigger a notification.
The dashboard supports the following scenarios and statuses:
Provided under Cisco Sample Code License, for details see LICENSE
Our code of conduct is available here
See our contributing guidelines here
Please note: This script is meant for demo purposes only. All tools/ scripts in this repo are released for use "AS IS" without any warranties of any kind, including, but not limited to their installation, use, or performance. Any use of these scripts and tools is at your own risk. There is no guarantee that they have been through thorough testing in a comparable environment and we are not responsible for any damage or data loss incurred with their use.
You are responsible for reviewing and testing any scripts you run thoroughly before use in any non-testing environment.
Owner
Contributors
Categories
NetworkingIoTCollaboration
Products
Meraki
Programming Languages
CSS
License
Code Exchange Community
Get help, share code, and collaborate with other developers in the Code Exchange community.View Community
Lorsque vous consultez un site Web, celui-ci peut enregistrer ou récupérer des informations présentes dans votre navigateur, le plus souvent sous la forme de témoins. Ces informations peuvent vous concerner, ainsi que vos préférences ou votre périphérique, et servent principalement à faire fonctionner le site selon vos attentes. Généralement, ces informations ne permettent pas de vous identifier directement, mais elles sont utilisées afin de vous offrir une expérience Web personnalisée. Parce que nous respectons votre droit à la vie privée, vous pouvez choisir de ne pas autoriser certains types de témoins. Cliquez sur les différents en-têtes de catégorie pour en apprendre davantage et modifier nos paramètres par défaut. Toutefois, bloquer certains types de témoins peut avoir une incidence sur votre expérience du site et les services que nous proposons.
Ces cookies sont indispensables au bon fonctionnement du site web et ne peuvent pas être désactivés au niveau de nos systèmes. Ils ne résultent généralement que d'actions que vous avez effectuées et qui correspondent à une demande de service, comme lorsque vous définissez vos préférences en matière de confidentialité, que vous vous connectez ou que vous remplissez des formulaires. Vous pouvez configurer votre navigateur pour qu'il bloque ces cookies ou vous avertisse de leur présence, mais certaines parties du site ne seront alors plus opérationnelles. Ces cookies n'enregistrent aucune information d'identification personnelle.
Ces cookies fournissent des indicateurs relatifs aux performances et à la facilité d'utilisation de notre site. Ils visent principalement à récolter des informations sur la façon dont vous interagissez sur le site, notamment le temps de chargement des pages, les temps de réponse, les messages d'erreur et le déroulement des interactions pour chaque visiteur. Nous pouvons ainsi examiner et analyser le comportement des visiteurs pour améliorer l'ergonomie et les fonctionnalités de notre site. Ces cookies nous permettent également de comptabiliser les visites et les sources de trafic afin de mesurer et d'améliorer les performances du site. Ils nous aident à savoir quelles pages sont les plus et les moins populaires, et à voir la manière dont les visiteurs naviguent sur le site. Si vous n'autorisez pas ces cookies, nous ne saurons pas si vous avez déjà visité le site et ne pourrons pas en mesurer les performances.
Ces cookies peuvent être intégrés à notre site par nos partenaires publicitaires. Ils peuvent être utilisés par ces entreprises pour établir un profil de vos centres d'intérêt et vous présenter des publicités pertinentes sur d'autres sites. Si ces cookies n'enregistrent pas directement vos informations personnelles, ils identifient de manière unique votre navigateur et votre terminal d'accès à Internet. Si vous n'autorisez pas ces cookies, des publicités moins ciblées s'afficheront.
Ces cookies permettent d’améliorer et de personnaliser les fonctionnalités du site Web. Ils peuvent être activés par nos équipes, ou par des tiers dont les services sont utilisés sur les pages de notre site Web. Si vous n'acceptez pas ces cookies, une partie ou la totalité de ces services risquent de ne pas fonctionner correctement.
