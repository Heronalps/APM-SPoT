import json
from pprint import pprint

with open('xray.json') as data_file:
    data = json.load(data_file)

target = open('xray.dat', 'w')
#print(len(data))
for i in range(0, len(data)):
    '''
    if len(data[i]) != 0:
        pprint(data[i])

    '''
    if len(data[i]) != 0 and "Duration" in data[i][0]:
        target.write(str(data[i][0]["Duration"]))
        target.write("\n")
