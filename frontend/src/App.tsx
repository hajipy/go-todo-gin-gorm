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

  const [newTodoTitle, setNewTodoTitle] = useState("");

  return (
    <>
      <h1>TODO</h1>
      <div>
        {todos.map((todo) => (
          <div key={todo.id}>
            <input
              type="checkbox"
              name="is_completed"
              checked={todo.completed_at !== null}
              onChange={(e) => {
                fetch(`${backendUrl}/todo/${todo.id}`, {
                  method: "PATCH",
                  headers: {
                    "Content-Type": "application/json",
                  },
                  body: JSON.stringify({ is_completed: e.target.checked }),
                })
                  .then((response) => response.json())
                  .then((updateTodo: Todo) =>
                    setTodos(
                      todos.map((todo) =>
                        todo.id === updateTodo.id ? updateTodo : todo,
                      ),
                    ),
                  );
              }}
            />
            <span>{todo.title}</span>
            &nbsp;
            <span>{todo.completed_at}</span>
            <button
              onClick={() => {
                fetch(`${backendUrl}/todo/${todo.id}`, {
                  method: "DELETE",
                }).then(() => setTodos(todos.filter((t) => t.id !== todo.id)));
              }}
            >
              Delete
            </button>
          </div>
        ))}
      </div>

      <form
        onSubmit={(e) => {
          e.preventDefault();
          fetch(`${backendUrl}/todo`, {
            method: "POST",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify({ title: newTodoTitle }),
          })
            .then((response) => response.json())
            .then((newTodo: Todo) => {
              setTodos([...todos, newTodo]);
              setNewTodoTitle("");
            });
        }}
      >
        <input
          type="text"
          name="title"
          value={newTodoTitle}
          onChange={(e) => setNewTodoTitle(e.target.value)}
        />
        <button type="submit">Add</button>
      </form>
    </>
  );
}

export default App;
