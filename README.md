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
scp_for_cd : false
send_notify : false
save_history : false
file_history : history.json
notify_chat_id : -12345567
```

### Run murocami server

```bash
make run
```

## Use ngrok to get public url

## You can test it by

```bash
curl -X POST http://localhost:8080/webhook \
-H "Content-Type: application/json" \
-d '{
  "ref": "refs/heads/main",
  "repository": {
    "name": "simple",
    "html_url": "https://github.com/osamikoyo/Simple",
    "clone_url": "https://github.com/osamikoyo/Simple.git"
  },
  "commits": [
    {
      "id": "abc123",
      "message": "Test commit"
    }
  ]
}'
```