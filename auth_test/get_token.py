import requests
from dotenv import load_dotenv
import os

load_dotenv()

API_KEY = os.getenv('EMAIL_API_KEY')
EMAIL = 'test@pigeons.house'
PASSWORD = 'hogehoge'

def print_token(email: str, password: str):
    uri = f"https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key={API_KEY}"
    data = {"email": email, "password": password, "returnSecureToken": True}

    result = requests.post(url=uri, data=data)

    user = result.json()

    print(email)
    print(user.get('idToken'))

print_token(EMAIL, PASSWORD)
