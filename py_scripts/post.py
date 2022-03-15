import sys
import json

out = dict()

n = len(sys.argv)
if not n>1:
    out["error"]= "no input received"

inp = sys.argv[1]

data = json.loads(inp)

j = json.dumps(data,separators=(',', ':'))

print(j)