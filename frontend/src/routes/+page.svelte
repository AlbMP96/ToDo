<script lang="ts">
	import { Status, statusLabels } from '$lib/enums/status';
	import { todos } from '$lib/stores/todos';
	import type { Todo } from '$lib/types';
	import List from '$lib/components/List.svelte';

	let newTodo = '';
	let todoStatus: Status = Status.Pending;

	function addTodo() {
		if (newTodo.trim() === '') return;
		todos.update((current: Todo[]) => [
			...current,
			{
				id: crypto.randomUUID(),
				text: newTodo.trim(),
				done: todoStatus == Status.Completed ? true : false,
				status: todoStatus
			}
		]);
		newTodo = '';
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Enter') addTodo();
	}
</script>

<h1>To-Do App</h1>

<div class="todo-form">
	<input placeholder="Nueva tarea" bind:value={newTodo} on:keydown={handleKeydown} />

	<select bind:value={todoStatus}>
		{#each Object.values(Status) as status}
			<option value={status}>{statusLabels[status]}</option>
		{/each}
	</select>

	<button on:click={addTodo}>Añadir</button>
</div>

<div class="status-lists">
	{#each Object.values(Status) as status}
		<List {status} />
	{/each}
</div>

<style>
	h1 {
		text-align: center;
		font-size: 2rem;
		padding: 5px;
		font-weight: bold;
	}

	.todo-form {
		display: flex;
		justify-content: center;
		gap: 10px;
		margin-bottom: 20px;
	}

	.status-lists {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
		gap: 15px;
	}
</style>
