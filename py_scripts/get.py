import json


d = dict()
d["name"] = "John"
d["id"] = 7
d["score"] = 100

j = json.dumps(d,separators=(',', ':'))

print(j)