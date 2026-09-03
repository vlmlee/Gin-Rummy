import './style.css'
import heroImg from './assets/hero.png'
import javascriptLogo from './assets/javascript.svg'
import viteLogo from './assets/vite.svg'
import { setupCounter } from './counter.js'
import { Terminal } from 'xterm'

const term = new Terminal({
  cursorBlink: true,
  theme: { background: '#1e1e1e' }
})

term.open(document.getElementById('terminal'))
term.write('Connecting to Go application backend...\r\n');

const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = (event) => {
    term.write(event.data);
};

term.onData((data) => {
    if (socket.readyState === WebSocket.OPEN) {
        socket.send(data);
    }
});

socket.onclose = () => {
    term.write('\r\n[Connection to Go backend closed]\r\n');
};