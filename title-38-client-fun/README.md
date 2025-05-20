### Run Tile38

```bash
docker run -p 9851:9851 tile38/
```
```
Unable to find image 'tile38/tile38:latest' locally
latest: Pulling from tile38/tile38
0a9a5dfd008f: Pull complete 
0fd51b3e746d: Pull complete 
27019f25c32b: Pull complete 
ce7dc10cf10a: Pull complete 
7062ad67c29e: Pull complete 
d90581ea4de8: Pull complete 
Digest: sha256:48849dfa95129ab5172a2d5c82ae6ec3561a882616062e83a54a61ea1e5c6f8b
Status: Downloaded newer image for tile38/tile38:latest

   _____ _ _     ___ ___
  |_   _|_| |___|_  | . |  Tile38 1.34.4 (fbf767a6) 64 bit (amd64/linux)
    | | | | | -_|_  | . |  Port: 9851, PID: 1
    |_| |_|_|___|___|___|  tile38.com

2025/05/19 07:36:24 [INFO] Server started, Tile38 version 1.34.4, git fbf767a6
2025/05/19 07:36:24 [INFO] AOF loaded 0 commands: 0.00s, 0/s, 0 bytes/s
2025/05/19 07:36:24 [INFO] Ready to accept connections at [::]:9851
```

### Run

```bash
./run.sh
```
```
❯ ./run.sh
2025/05/19 00:37:44 [PING]: {"ok":true,"ping":"pong","elapsed":"342ns"}
2025/05/19 00:37:44 [SET fleet truck1 POINT 33.5123 -112.2693]: {"ok":true,"elapsed":"24.502µs"}
2025/05/19 00:37:44 [SET fleet truck2 EX 20 FIELD speed 20 POINT 33.4626 -112.1695]: {"ok":true,"elapsed":"9.83µs"}
2025/05/19 00:37:44 [NEARBY fleet WHERE speed 0 100 MATCH truck* POINTS POINT 33.462 -112.268 6000]: {"ok":true,"points":[{"id":"truck1","point":{"lat":33.5123,"lon":-112.2693}}],"count":1,"cursor":0,"elapsed":"70.157µs"}
truck1 {33.5123 -112.2693}
```
