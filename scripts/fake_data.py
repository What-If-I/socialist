import random

from faker import Faker
import requests

fake = Faker('en_US')

server = "http://localhost:8080"
for _ in range(1000):
    p = fake.simple_profile()

    s = requests.Session()
    resp = s.post(server + "/login", data={"username": p['username'], "password": "123"})
    resp.raise_for_status()

    name, surname = p['name'].split(" ", 1)
    age = random.randint(15, 70)
    resp = s.post(server+"/update-user-info", data={
        "name": name,
        "surname": surname,
        "age": age,
        "gender": p['sex'],
        "interests": f"work as {fake.job()} and like {fake.color_name()} color",
        "city": p['address'],
    })
    resp.raise_for_status()
