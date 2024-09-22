import { useEffect, useState } from "react";

const backendUrl = import.meta.env.VITE_BACKEND_URL;

type Todo = {
  id: number;
  title: string;
  completed_at: string;
};

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  useEffect(() => {
    fetch(`${backendUrl}/todo`)
      .then((res) => res.json())
      .then((data) => setTodos(data));
  }, []);

  return (
    <>
      <h1>TODO</h1>
      <ul>
        {todos.map((todo) => (
          <li key={todo.id}>
            {todo.id} {todo.title} {todo.completed_at}
          </li>
        ))}
      </ul>
    </>
  );
}

export default App;
