import { useEffect, useState } from "react";

import viteLogo from "/vite.svg";
import reactLogo from "./assets/react.svg";
import "./App.css";

const backendUrl = import.meta.env.VITE_BACKEND_URL;

type Todo = {
  id: number;
  title: string;
  completed_at: string;
};

function App() {
  const [count, setCount] = useState(0);
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetch(`${backendUrl}/todo`)
      .then((res) => res.json())
      .then((data) => setTodos(data));
  }, []);

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
      <div>
        <h2>TODO</h2>
        <ul>
          {todos.map((todo) => (
            <li key={todo.id}>
              {todo.id} {todo.title} {todo.completed_at}
            </li>
          ))}
        </ul>
      </div>
    </>
  );
}

export default App;
