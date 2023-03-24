from flask import Flask, send_from_directory
from flask import request

app = Flask(__name__)

@app.route('/')
def index():
    return send_from_directory('dist/', 'index.html')

@app.route('/<path:path>')
def static_proxy(path):
    return send_from_directory('dist/', path)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8900)
