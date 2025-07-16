<script lang="ts">
	import { Status, statusLabels } from '$lib/enums/status';
	import { todos } from '$lib/stores/todos';
	import type { Todo } from '$lib/types';
	import List from '$lib/components/List.svelte';
	import ButtonPrimary from '$lib/components/buttons/ButtonPrimary.svelte';

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

<div>
	<h1>To-Do App</h1>

	<div class="todo-form">
		<input placeholder="Nueva tarea" bind:value={newTodo} on:keydown={handleKeydown} />

		<select bind:value={todoStatus}>
			{#each Object.values(Status) as status}
				<option value={status}>{statusLabels[status]}</option>
			{/each}
		</select>

		<ButtonPrimary onClick={addTodo}>Añadir</ButtonPrimary>
	</div>

	<div class="status-lists">
		{#each Object.values(Status) as status}
			<List {status} />
		{/each}
	</div>
</div>

<style lang="scss">
	@media (max-width: 768px) {
		.status-lists {
			width: 100%;
		}
	}

	h1 {
		text-align: center;
		font-size: 2rem;
		padding: 5px;
		font-weight: bold;
	}

	.todo-form {
		box-sizing: border-box;
		display: flex;
		justify-content: center;
		gap: 10px;
		margin-bottom: 20px;
		width: 75vw;
		border: 1px solid black;
		border-radius: 5px;
		margin-inline: auto;
		padding: 10px;

		input {
			flex-grow: 1;
			padding: 5px;
			border: 1px solid #ccc;
			border-radius: 5px;
		}
	}

	.status-lists {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
		gap: 15px;
		width: 75vw;
		margin-inline: auto;
	}
</style>
