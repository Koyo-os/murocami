# Murocami

## Its modern CI tool, for selfhosting

### How use it?

```bash
git clone https://github.com/Koyo-os/murocami.git
cd murocami
```
### You can cange config values by config.yaml

```yaml
port : 8080
host : "localhost"
temp_dir_name : "murocami-temp"
input_point : "cmd/main.go"
output_point : "bin/app"
```

### Run murocami server

```bash
make run
```

## Use ngrok to get public url