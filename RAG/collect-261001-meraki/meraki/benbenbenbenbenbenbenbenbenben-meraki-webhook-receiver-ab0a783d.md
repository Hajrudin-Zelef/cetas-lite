---
id: collect-261001-meraki/meraki/benbenbenbenbenbenbenbenbenben-meraki-webhook-receiver-ab0a783d
title: "Door Sensor Open"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/benbenbenbenbenbenbenbenbenben-meraki-webhook-receiver-ab0a783d.md
source_anchor: ""
source_lines: [1, 44]
sha256: e3745a1857885c3ed374bc830abd4f49c067b8e9c1d0e62c4ffa09a7cc8a4261
---

# Door Sensor Open

Receive Webhooks from Meraki Dashboard and post to Webex Room.
- Webex Bot and Token
- Python
- Ngrok (Or similar if running locally)
- Meraki Dashboard
Skip if you already have a Bot and Token.
- Login to Webex Developer site https://developer.webex.com/my-apps
- Create a Bot (Create a new App > Create a new Bot)
  - Give your bot a name, username and description
  - Copy the generated Access Token (store somewhere safe for later)
- Add your Bot to a Webex Teams Space
  - In Webex Teams click the + button and create a space
  - Give your space a name and add your Bot (i.e. yourbot@webex.bot)
  - Click Create
  - Say “Hello” if you want :)
Exposes your development environment to the Internet using a public URL so we can receive Webhooks.
- Signup for Free Ngrok Account https://ngrok.com/ and download Ngrok
- Follow instructions in Ngrok to link authtoken
- Start Ngrok
$ ./ngrok http 5000
- Copy the Ngrok Forwarding URL for later use (ie. https://zzzzzzzz.ngrok.io)
Configure the webhook server in your Meraki Dashboard. Network-wide > Alerts
- Clone Github Repository
git clone https://github.com/benbenbenbenbenbenbenbenbenben/meraki-webhook-receiver.git
cd meraki-webhook-receiver
- Now "activate" the python virtual environment
python3 -m venv venv
source venv/bin/activate
- Install project requirements
pip install Flask requests
- Export Env Variables
export BOT_TOKEN=
export MERAKI_SECRET=
- Run The App. Initially this will return a list of webex rooms the bot has been added to. Copy the room Id for your room and export.
python app.pyexport WEBEX_ROOM=python app.py
Send a test webhook from Network-wide > Alerts Page or using the new Environmental > Alert Profiles test buttons.
Edit meraki_events.py to have your own custom messages.
# Door Sensor Open
if webhook['deviceModel'] == 'MT20' and webhook["alertData"]["triggerData"][0]["trigger"]["sensorValue"] == 1.0:
    message = f'**Door Opened** 🚪🏃\
                \n- **Network:** {webhook["networkName"]}\
                \n- **Sensor:** {webhook["deviceName"]}\
                \n- **Time:** {time_format(ts)}'
    return message
