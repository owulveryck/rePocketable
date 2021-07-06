WIP

| KEY                             | TYPE        | DEFAULT    | REQUIRED    | DESCRIPTION    |
|---------------------------------|-------------|------------|-------------|----------------|
| DOWNLOADER_LIVENESS_CHECK       | Duration    | 5m         | true        |                |
| DOWNLOADER_PROBE_TIMEOUT        | Duration    | 60m        | true        |                |
| DOWNLOADER_HTTP_TIMEOUT         | Duration    | 10s        | true        |                |
| DOWNLOADER_TRANSPORT_TIMEOUT    | Duration    | 5s         | true        |                |

| KEY                        | TYPE        | DEFAULT                         | REQUIRED    | DESCRIPTION                                                        |
|----------------------------|-------------|---------------------------------|-------------|--------------------------------------------------------------------|
| POCKET_UPDATE_FREQUENCY    | Duration    | 1h                              | true        | How often to query getPocket                                       |
| POCKET_HEALTH_CHECK        | Duration    | 30s                             | true        |                                                                    |
| POCKET_POCKET_URL          | String      | https://getpocket.com/v3/get    | true        |                                                                    |
| POCKET_CONSUMER_KEY        | String      |                                 | true        | See https://getpocket.com/developer/apps/ to get a consumer key    |
| POCKET_USERNAME            | String      |                                 |             | The pocket username (will try to fetch it if not found)            |
| POCKET_TOKEN               | String      |                                 |             | The access token, will try to fetch it if not found or invalid     |
