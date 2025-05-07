import { useEffect, useRef, useState } from 'react';

function App() {
  const wsRef = useRef<WebSocket | null>(null);
  const [messages, setMessages] = useState<string[]>([]);
  const [status, setStatus] = useState('Connecting...');

  useEffect(() => {
    const ws = new WebSocket('ws://localhost:8080/ws');
    wsRef.current = ws;

    ws.onopen = () => {
      setStatus('Connected to backend');
      ws.send('hello from client');
    };

    ws.onmessage = (event) => {
      setMessages((prev) => [...prev, `Received: ${event.data}`]);
    };

    ws.onerror = () => {
      setStatus('WebSocket error');
    };

    ws.onclose = () => {
      setStatus('Connection closed');
    };

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div style={{ padding: '2rem', fontFamily: 'monospace' }}>
      <h1>🖥️ Deskrun</h1>
      <p>Status: {status}</p>
      <div>
        {messages.map((msg, i) => (
          <p key={i}>{msg}</p>
        ))}
      </div>
    </div>
  );
}

export default App;
