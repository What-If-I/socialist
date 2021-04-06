import random

from faker import Faker
import requests

fake = Faker('en_US')

with open("generate_profile.sql", "w") as fh:
    for i in range(1, 1_000_001):
        p = fake.simple_profile()
        name, surname = p['name'].split(" ", 1)
        age = random.randint(15, 70)
        interest = f"work as {fake.job()} and like {fake.color_name()} color".replace("'", "")
        username = f"{p['username']}_{i}"
        fh.write(
            f"INSERT INTO user (id, username, password_hash) VALUES ({i}, '{username}', '$2a$04$h8hE36SAs5j0bCr4E2bjLONxzgv1Sb3HdHAcDlDDxIz/kOx8F/K6y');\n")
        fh.write(f"INSERT INTO profile (user_id, name, surname, age, gender, interests, city) VALUES ({i}, '{name}', '{surname}', {age}, '{p['sex']}', '{interest}', '{p['address']}');\n")
        if i % 10000 == 0:
            print(f"done {i}")


# server = "http://142.93.169.163:8080"



# s = requests.Session()
# resp = s.post(server + "/login", data={"username": p['username'], "password": "123"})
# resp.raise_for_status()
#

# resp = s.post(server+"/update-user-info", data={
#     "name": name,
#     "surname": surname,
#     "age": age,
#     "gender": p['sex'],
#     "interests": f"work as {fake.job()} and like {fake.color_name()} color",
#     "city": p['address'],
# })
# resp.raise_for_status()

