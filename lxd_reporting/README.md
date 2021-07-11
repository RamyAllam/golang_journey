## LXD Reporting
#### This tool reports the containers inventory as a table or CSV to the stdout or a file

### Usage
```
Usage of ./lxd_reporting.bin:
    --connector string   LXD connection method: ['tcp', 'socket'] (default "tcp")
    --crt string         LXD connection certificate (default "./certs/client.crt")
    --key string         LXD connection key (default "./certs/client.key")
    --socket string      LXD connection socket path (default "/var/snap/lxd/common/lxd/unix.socket")
    --url string         LXD connection URL. Ex. https://192.168.1.230:8443
    --format string      Output format ['table', 'csv'] (default "table")
    --output string      Output destination: ['output.txt', 'output.csv'] (default "os.Stdout")
```

### Examples
```
./lxd_reporting.bin --connector=socket --format csv
./lxd_reporting.bin --connector=socket --format csv --output containers.csv
./lxd_reporting.bin --connector=socket --format table
./lxd_reporting.bin --connector=socket --format table --output containers.txt
./lxd_reporting.bin --url=https://192.168.1.230:8443
./lxd_reporting.bin --url=https://192.168.1.230:8443 --crt="./certs/client.crt" --key="./certs/client.key"
```