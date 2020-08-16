![shazam-webhook](https://raw.githubusercontent.com/bst27/shazam-webhook/master/website/banner.png)

shazam-webhook
======
This application allows to monitor Shazam charts available at shazam.com and
sending webhooks holding artist and title when new tracks are added. E.g. if
you want to **add Shazam charts to a Spotify playlist** you can do this using
this application and a 3rd party service to receive webhooks like 
[IFTTT](https://ifttt.com/) or [Zapier](https://zapier.com/).

# Usage
At first you have to customize the **app-config.json**: Add some Shazam charts
you want to monitor. Also add tracklists to receive tracks from monitored
Shazam charts and define some webhooks.

Below is an example configuration. By disabling initial hooks no webhooks
are triggered when tracks are added to a tracklist during the first iteration. The polling 
interval is set to fetch Shazam charts every 10800 seconds (= 3 hours). Two
Shazam charts are defined to be monitored: Berlin (81310) and Paris (80321).
A single tracklist is defined. By watching two charts the tracks of both 
Shazam charts will be added to the tracklist. When a track is added the defined webhook
will receive artist and title as a POST request. 
```
{
    "SendInitialHooks": false,
    "PollingInterval": 10800,
    "MonitoredShazamCharts": [
        "81310",
        "80321"
    ],
    "Tracklists": [ 
        {
            "WatchedShazamCharts": [
                "81310",
                "80321"
            ],
            "WebhookTargets": [
                {
                    "URL": "http://localhost:8080/72A4E"
                }
            ]
        }
    ]
}
```
You can find a list of available Shazam city IDs to be monitored at [bst27/shazam](https://github.com/bst27/shazam/blob/24cdf23c38ce845e2b0333368bba697888bb0412/cities.go).

After customizing the configuration you can build and run the application:
```
go build; ./shazam-webhook
```

# TODO
This section holds some ideas:

## Config file path
Allow to set path to config file via command line argument.

## Config Creator
Create a command line tool to generate an **app-config.json** file. This file will hold the app configuration.

## Tests
As always there could be more tests.