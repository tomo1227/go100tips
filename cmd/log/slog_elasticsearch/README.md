# how to use

## local

```shell
docker network inspect go100tips_app
```

Containersの中に以下3つのコンテナがあることを確認する。

```json
"Containers": {
    "***": {
        "Name": "kibana",
        "EndpointID": "***",
        "MacAddress": "02:90:1b:14:3e:45",
        "IPv4Address": "172.18.0.2/16",
        "IPv6Address": ""
    },
    "***": {
        "Name": "go100tips",
        "EndpointID": "***",
        "MacAddress": "1e:cc:1f:dc:c7:eb",
        "IPv4Address": "172.18.0.4/16",
        "IPv6Address": ""
    },
    "***": {
        "Name": "elasticsearch",
        "EndpointID": "260085dd7d642e25755b3e5641b4412423e75a6ed102d86bbc871c4c4e5c8bc2",
        "MacAddress": "36:2e:eb:91:d9:31",
        "IPv4Address": "172.18.0.3/16",
        "IPv6Address": ""
    }
},
```

ローカルから以下のURLにアクセスできることを確認する。

* Kibana : http://localhost:5601
* ElasticSearch : http://localhost:9200

## コンテナ

コンテナから別のコンテナにアクセスする場合は以下。

```shell
curl http://kibana:5601
curl http://elasticsearch:9200
```

コンテナからローカルにアクセスする場合は以下。

```shell
curl http://host.docker.internal:5601
curl http://host.docker.internal:9200
```
