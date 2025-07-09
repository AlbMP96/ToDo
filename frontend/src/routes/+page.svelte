<script lang="ts">
	import { Status, statusLabels } from '$lib/enums/status';
	import { todos } from '$lib/stores/todos';
	import type { Todo } from '$lib/types';

	let newTodo = '';
	let todoStatus: Status = Status.Pending;

	function addTodo() {
		if (newTodo.trim() === '') return;
		todos.update((current: Todo[]) => [
			...current,
			{ id: crypto.randomUUID(), text: newTodo.trim(), done: false, status: todoStatus }
		]);
		newTodo = '';
	}

	function toggleTodo(id: string) {
		todos.update((current: Todo[]) =>
			current.map((todo) => (todo.id === id ? { ...todo, done: !todo.done } : todo))
		);
	}

	function deleteTodo(id: string) {
		todos.update((current: Todo[]) => current.filter((todo) => todo.id !== id));
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') addTodo();
	}
</script>

<h1>To-Do App</h1>

<input placeholder="Nueva tarea" bind:value={newTodo} on:keydown={handleKeydown} />

<select bind:value={todoStatus}>
	{#each Object.values(Status) as status}
		<option value={status}>{statusLabels[status]}</option>
	{/each}
</select>

<button on:click={addTodo}>Añadir</button>

<ul>
	{#each $todos as todo (todo.id)}
		<li>
			<input type="checkbox" checked={todo.done} on:change={() => toggleTodo(todo.id)} />
			<span class:done={todo.done}>
				{todo.text}
			</span>
			<select bind:value={todo.status}>
				{#each Object.values(Status) as status}
					<option value={status}>{statusLabels[status]}</option>
				{/each}
			</select>
			<button on:click={() => deleteTodo(todo.id)}>🗑️</button>
		</li>
	{/each}
</ul>

<style>
	.done {
		text-decoration: line-through;
	}
</style>
