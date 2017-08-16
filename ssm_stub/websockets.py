import websocket, json
import thread
import time

def on_message(ws, message):
    print(message)

def on_error(ws, error):
    print(error)

def on_close(ws):
    print("### closed ###")

def on_open(ws):
    # def run(*args):
    #     for i in range(3):
    #         time.sleep(1)
    toSend  = {"name": "fsm update", "data": {"fsmId": "1234", "state": "started"}}
    #
    #     time.sleep(1)
    # toSend  = {"name": "fsm update", "data":
    # [
    #     {"name": "test", "fsmId": "2", "state": "stopped"},
    #     {"name": "test", "fsmId": "3", "state": "stopped"}
    # ],
    # }
    toSendJson = json.dumps(toSend)
    ws.send(toSendJson)
    ws.close()
    print("thread terminating...")
    # thread.start_new_thread(run, ())

if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.WebSocketApp("ws://localhost:4000/ws",
                              on_message = on_message,
                              on_error = on_error,
                              on_close = on_close)
    ws.on_open = on_open
    ws.run_forever()
