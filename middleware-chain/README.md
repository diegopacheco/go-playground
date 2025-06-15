### Run

```bash
./run.sh
```

### Result

```
❯ curl -v http://localhost:8080/
* Host localhost:8080 was resolved.
* IPv6: ::1
* IPv4: 127.0.0.1
*   Trying [::1]:8080...
* Connected to localhost (::1) port 8080
* using HTTP/1.x
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/8.11.1
> Accept: */*
> 
* Request completely sent off
< HTTP/1.1 403 Forbidden
< Content-Type: text/plain; charset=utf-8
< X-Content-Type-Options: nosniff
< Date: Sat, 14 Jun 2025 19:34:06 GMT
< Content-Length: 10
< 
Forbidden
* Connection #0 to host localhost left intact
```
