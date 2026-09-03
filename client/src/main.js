import '@xterm/xterm/css/xterm.css'
import './style.css'
import { Terminal } from '@xterm/xterm'
import { FitAddon } from '@xterm/addon-fit'

const term = new Terminal({
  cursorBlink: true,
  scrollback: 1000,
  fontFamily: '"Share Tech Mono", "Courier New", monospace',
  fontSize: 20,
  lineHeight: 1.15,
  theme: {
    background: '#041208',
    foreground: '#3cff6a',
    cursor: '#7dff9a',
    cursorAccent: '#041208',
    selectionBackground: 'rgba(60, 255, 106, 0.28)',
    black: '#041208',
    green: '#3cff6a',
    brightGreen: '#9affb4',
  },
})

const fitAddon = new FitAddon()
term.loadAddon(fitAddon)

const screen = document.getElementById('terminal')
term.open(screen)

function fitTerm() {
  fitAddon.fit()
  const cols = document.querySelector('.monitor-chin .cols')
  if (cols) {
    cols.textContent = `${term.cols} COL`
  }
}

fitTerm()
document.fonts?.ready.then(fitTerm)

const resizeObserver = new ResizeObserver(() => {
  fitTerm()
})
resizeObserver.observe(screen)
window.addEventListener('resize', fitTerm)

term.write('Connecting to Go application backend...\r\n')

const socket = new WebSocket(`wss://ginrummy.live/ws`)

socket.onmessage = (event) => {
  term.write(event.data)
}

term.onData((data) => {
  if (socket.readyState === WebSocket.OPEN) {
    socket.send(data)
  }
})

socket.onclose = () => {
  term.write('\r\n[Connection to Go backend closed]\r\n')
}
