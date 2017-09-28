import websocket
import thread
import time

from threading import Thread
from websocket_server import WebsocketServer
from json import loads, dumps

class Server:
    # Called for every client connecting (after handshake)
    def new_client(self, client, server):
    	print("New client connected and was given id %d" % client['id'])

    # Called for every client disconnecting
    def client_left(self, client, server):
    	print("Client(%d) disconnected" % client['id'])


    # Called when a client sends a message
    def message_received(self, client, server, message):
    	if len(message) > 200:
    		message = message[:200]+'..'
    	print("Client(%d) said: %s" % (client['id'], message))

        # Format message
        messageDict = loads(message)
        actionName = messageDict['name']


        #TODO relay request on queue and wait for response
        def sendMessage():
            print("Sending Message")
            toSend = None
            if actionName == "fsm start":
                fsmName = messageDict['Data']['name']
                fsmID = messageDict['Data']['id']
                toSend  = {"name": actionName, "Data": {
                    "name": fsmName,
                    "id": fsmID,
                    "state": "started"
                    }
                }

            if actionName == "fsm stop":
                fsmName = messageDict['Data']['name']
                fsmID = messageDict['Data']['id']
                toSend  = {"name": actionName, "Data": {
                    "name": fsmName,
                    "id": fsmID,
                    "state": "stopped"
                    }
                }

            if actionName == "basic start":
              print actionName
              toSend  = {
                  "name": "basic start",
                  "data":
                  [
                      {"name": "Firewall", "id": "1", "state": "started"},
                      {"name": "VPN", "id": "2", "state": "started"}
                  ],
                  }

            if actionName == "anon start":
              print actionName
              toSend  = {
                  "name": "anon start",
                  "data":
                  [
                    {"name": "Firewall", "id": "1", "state": "started"},
                    {"name": "VPN", "id": "2", "state": "started"},
                    {"name": "TOR", "id": "3", "state": "started"},
                    {"name": "HTTP Proxy", "id": "4", "state": "started"},
                    {"name": "IDS", "id": "5", "state": "started"}
                  ],
                  }
            try:
                toSendJson = dumps(toSend)
                print toSendJson
                server.send_message(client, toSendJson)
            except Exception, e:
                print str(e)

        sendMessage()


    def listenToFSMRequests(self):
        print "Listening to Requests...!"
        port=9191
        host="selfservice-backend"
        server = WebsocketServer(port, host=host)
        server.set_fn_new_client(self.new_client)
        server.set_fn_client_left(self.client_left)
        server.set_fn_message_received(self.message_received)
        server.run_forever()

class Client:
    def on_message(self, ws, message):
        print(message)

    def on_error(self, ws, error):
        print(error)

    def on_close(self, ws):
        print("### closed ###")

    def on_open(self, ws):

        ##TODO replace toSend with real reply within the SSM
        toSend  = {"name": "fsm add", "data":
        [
            {"name": "Firewall", "fsmId": "1", "state": "stopped"},
            {"name": "VPN", "fsmId": "2", "state": "started"},
            {"name": "TOR", "fsmId": "3", "state": "started"},
            {"name": "HTTP Proxy", "fsmId": "4", "state": "started"},
            {"name": "IDS", "fsmId": "5", "state": "stopped"}
        ],
        }
        toSendJson = dumps(toSend)
        ws.send(toSendJson)
        ws.close()

    def advertiseFSMs(self):
        websocket.enableTrace(True)
        selfservice_backend = "localhost"
        ws = websocket.WebSocketApp("ws://"+selfservice_backend+":4000/ws",
                                  on_message = self.on_message,
                                  on_error = self.on_error,
                                  on_close = self.on_close)
        ws.on_open = self.on_open
        ws.run_forever()


if __name__ == "__main__":

    client = Client()
    server = Server()

    Thread(target = client.advertiseFSMs()).start()
    Thread(target = server.listenToFSMRequests()).start()
