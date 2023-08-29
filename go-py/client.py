import requests

response = requests.get("http://localhost:8080/api/myendpoint")
data = response.json()
print(data["text"])  # Output: Hello from Go!
