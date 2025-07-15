<script lang="ts">
	import { todos } from '$lib/stores/todos';
	import type { Todo } from '$lib/types';
	import { Status, statusLabels } from '$lib/enums/status';

	export let status: Status;

	$: filteredTodos = $todos.filter((todo) => todo.status === status);

	function toggleTodo(id: string) {
		todos.update((current: Todo[]) =>
			current.map((todo) =>
				todo.id === id
					? { ...todo, done: !todo.done, status: !todo.done ? Status.Completed : Status.Pending }
					: todo
			)
		);
	}

	function updateTodoStatus(id: string, newStatus: Status) {
		todos.update((current: Todo[]) =>
			current.map((todo) => (todo.id === id ? { ...todo, status: newStatus } : todo))
		);
	}

	function deleteTodo(id: string) {
		todos.update((current: Todo[]) => current.filter((todo) => todo.id !== id));
	}
</script>

<div class="todo-list">
	<h2>{statusLabels[status]}</h2>

	<ul>
		{#each filteredTodos as todo (todo.id)}
			<li>
				<input type="checkbox" checked={todo.done} on:change={() => toggleTodo(todo.id)} />
				<span class:done={todo.done}>{todo.text}</span>
				<select
					bind:value={todo.status}
					on:change={(e) => {
						const target = e.target as HTMLSelectElement | null;
						if (target) {
							updateTodoStatus(todo.id, target.value as Status);
						}
					}}
				>
					{#each Object.values(Status) as status}
						<option value={status}>{statusLabels[status]}</option>
					{/each}
				</select>
				<button on:click={() => deleteTodo(todo.id)}>🗑️</button>
			</li>
		{/each}
	</ul>
</div>

<style>
	div {
		border: 1px solid #ccc;
		border-radius: 5px;
	}

	h2 {
		text-align: center;
	}

	.done {
		text-decoration: line-through;
	}

	.todo-list {
		padding: 10px;
	}
</style>
