// import ChatMessage from '@/core/classes/ChatMessage';

// import Classes from './Classes';

const apiHost = import.meta.env.VITE_APP_API_HOST ?? '';
type OnMessageF = (...args: any[]) => void;

export default class WS {
  private client?: WebSocket;
  private endpoint = '';
  private query = '';
  private period = 20000;

  private onMessageF: OnMessageF = () => {};

  static Create(onMessage: OnMessageF, endpoint = '', query = ''): WS {
    const ws = new WS(endpoint, query);
    ws.connect();
    ws.ping;
    ws.setOnMessage(onMessage);
    return ws;
  }

  constructor(endpoint = '', query = '') {
    PF.H.Classes.BuildClass(this, { endpoint, query });
  }

  send(data: string | ArrayBufferLike | Blob | ArrayBufferView) {
    console.log(this.client);
    if (!this.client || this.client.readyState !== 1) {
      return;
    }
    this.client.send(data);
  }

  ping() {
    setInterval(this.sendPing.bind(this), this.period);
  }

  private sendPing() {
    if (!this.client || this.client.readyState !== 1) {
      return;
    }
    this.send(JSON.stringify('ping'));
  }

  private parseOnMessageF(f: MessageEvent) {
    const m = JSON.parse(f.data);
    if (m.type !== 'message') {
      return;
    }
    this.onMessageF(m);
  }

  setOnMessage(onMessageF: OnMessageF): void {
    this.onMessageF = onMessageF;
    if (!this.client) {
      return;
    }
    this.client.onmessage = this.parseOnMessageF;
  }

  private buildUrl(): string {
    const path = `ws/${this.endpoint}/${this.query}`;
    console.log(path);
    return new URL(path, apiHost.replace('http', 'ws')).toString();
  }

  connect() {
    this.client = new WebSocket(this.buildUrl());
    this.client.onclose = this.reconnect;
    this.client.onerror = this.reconnect;
  }

  private reconnect() {
    setTimeout(this.connect, 3000);
  }
}
