This is an updated version from:
https://blog.logrocket.com/making-http-requests-in-go/

It supports a couple of additioanl flags such as apikey for Nomics API and slack webhook URL

```
% ./coinsapp -h
Usage of ./coinsapp:
  -apikey string
        Nomics API key
  -crypto string
        The crypto currency name (default "BTC")
  -fiat string
        The name of the fiat currency (default "USD")
  -slack string
        Slack webhook URL
```
