import { createHteaApi } from "../plugin/htea_api";
import { TodoList } from "../components/todos/TodoList";
import { Todo } from "../components/todos/Todo";

export const todosApi = createHteaApi()
  .get("/api/todos", ({ render }) =>
    render(TodoList, { todos: [{ text: "todo1" }, { text: "todo2" }] })
  )
  .post("/api/todos", ({ render, body }) => render(Todo, { text: body.title }));
