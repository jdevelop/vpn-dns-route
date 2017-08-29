# VPN route helper script

Often you don't want to pass all the traffic through some VPN server (OpenVPN anonymizers etc). 

Since many of the domains may have multiple IP addresses, rotated in round-robin manner - the task of adding the routes manually could be tedious and boring.

So I created this simple script, that resolves `A`-records for the given host names using the DNS server provided (`8.8.8.8` for example, if the ISP blocks certain DNS names) and adds the routes for them to the given interface, using [netlink](https://github.com/docker/libcontainer/netlink) library from Docker.

The usage is simple:

```
sudo ./vpn-dns-route -iface tun0 mail.ru apiok.ru
```

The resulting route table:

```
10.8.1.217      0.0.0.0         255.255.255.255 UH        0 0          0 tun0
94.100.180.200  0.0.0.0         255.255.255.255 UH        0 0          0 tun0
94.100.180.201  0.0.0.0         255.255.255.255 UH        0 0          0 tun0
217.20.144.164  0.0.0.0         255.255.255.255 UH        0 0          0 tun0
217.20.149.108  0.0.0.0         255.255.255.255 UH        0 0          0 tun0
217.69.139.200  0.0.0.0         255.255.255.255 UH        0 0          0 tun0
217.69.139.201  0.0.0.0         255.255.255.255 UH        0 0          0 tun0
```

Lookup of the A records with `dig` :

```
dig apiok.ru @8.8.8.8 A                        

;; ANSWER SECTION:
apiok.ru.               242     IN      A       217.20.149.108
apiok.ru.               242     IN      A       217.20.144.164
```

and for **mail.ru**

```
;; ANSWER SECTION:
mail.ru.		36	IN	A	94.100.180.200
mail.ru.		36	IN	A	94.100.180.201
mail.ru.		36	IN	A	217.69.139.200
mail.ru.		36	IN	A	217.69.139.201
```

So far - looks good.

Enjoy!