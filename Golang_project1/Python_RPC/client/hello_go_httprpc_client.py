import requests
request = {
    "method": "HelloService.Hello",
    "params": ["python client httprpc"],
    "id": 0
}

rsp = requests.post("http://localhost:1234/jsonrpc", json=request)

print(rsp.text)