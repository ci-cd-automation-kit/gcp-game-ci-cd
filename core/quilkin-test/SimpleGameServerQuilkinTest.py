import socket
import time
import base64

#"35.240.188.100"
ip = "35.223.66.108"
port = 7136

message = b"abcEXIT"
# message = "abcHello".encode()
#message_bytes = message.encode('ascii')
base64_bytes = base64.b64encode(message)
base64_message = base64_bytes.decode('ascii')

# print("decode OGdqM3YyaQ== : ", base64.b64decode("OGdqM3YyaQ=="), "\n")
# print("decode YWJj : ", base64.b64decode("YWJj"), "\n\n")

# print(len(base64_bytes))
print(base64_message)
print(base64_bytes)
# Create socket for server
s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)
print("Do Ctrl+c to exit the program !!")

# Let's send data through UDP protocol
#while True:
#    send_data = input("Type some text to send =>");
s.sendto(message, (ip, port))
print("\n\n 1. Client Sent : ", base64_message, "\n\n")
time.sleep(5)
data, address = s.recvfrom(4096)
print("\n\n 2. Client received : ", data.decode('utf-8'), "\n\n")
# close the socket
s.close()
