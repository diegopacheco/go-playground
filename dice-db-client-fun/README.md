### Run DiceDB

```bash
docker run -p 7379:7379 dicedb/dicedb:latest
```
```
Unable to find image 'dicedb/dicedb:latest' locally
latest: Pulling from dicedb/dicedb
f18232174bc9: Pull complete 
3622eeabca00: Pull complete 
0be06685cc34: Pull complete 
7678b63cf684: Pull complete 
Digest: sha256:a1f529ca49b4315c42b050e5f34bf45415de6c5d2c33682c4e7719f0259c4b50
Status: Downloaded newer image for dicedb/dicedb:latest

	██████╗ ██╗ ██████╗███████╗██████╗ ██████╗ 
	██╔══██╗██║██╔════╝██╔════╝██╔══██╗██╔══██╗
	██║  ██║██║██║     █████╗  ██║  ██║██████╔╝
	██║  ██║██║██║     ██╔══╝  ██║  ██║██╔══██╗
	██████╔╝██║╚██████╗███████╗██████╔╝██████╔╝
	╚═════╝ ╚═╝ ╚═════╝╚══════╝╚═════╝ ╚═════╝
			
2025-05-19T04:46:53Z INF starting DiceDB version=v1.0.10
2025-05-19T04:46:53Z INF running with total_commands=40
2025-05-19T04:46:53Z INF running with engine=ironhawk
2025-05-19T04:46:53Z INF running with port=7379
2025-05-19T04:46:53Z INF running on cores=12
2025-05-19T04:46:53Z INF running with shards=12
```

### Run

```bash
./run.sh
```
```
❯ ./run.sh
successfully set key k1=v1
successfully got key k1=v1
```
