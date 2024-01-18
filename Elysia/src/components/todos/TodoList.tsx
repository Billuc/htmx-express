const TodoList = createComponent("todo-list", {
  style: /* css */ `
        .todo-list {
            display: flex;
            flex-direction: column;
        }
    `,
  render(props: { todos: { text: string }[] }) {
    const todoItems = props.todos.map((todo) => <div>{todo.text}</div>);

    return <div class="todo-list">{todoItems}</div>;
  },
});
