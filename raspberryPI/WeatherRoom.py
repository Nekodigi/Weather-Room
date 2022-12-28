import smbus2
import bme280
import requests
import time

API_URL = "PUT YOUR API URL HERE"

port = 1
address = 0x76
bus = smbus2.SMBus(port)

calibration_params = bme280.load_calibration_params(bus, address)

# the sample method will take a single reading and return a
# compensated_reading object

def GetJson():
    data = bme280.sample(bus, address, calibration_params)
    #print(f"{data.temperature} {data.humidity} {data.pressure}")

    return {
        "temperature": data.temperature,
        "humidity": data.humidity,
        "atmosphere": data.pressure,
        "co2": 0
    }
    
count = 0
import sys
print(sys.version)
while True:
    count+=1
    json = GetJson()
    response = requests.post(API_URL+'/indicator', json=json)
    print(f"Indocator update: {count} {response.status_code} {response.text}")   
    time.sleep(60)