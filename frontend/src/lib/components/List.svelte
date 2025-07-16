<script lang="ts">
	import { todos } from '$lib/stores/todos';
	import { Status, statusLabels } from '$lib/enums/status';
	import ToDoCard from './ToDoCard.svelte';
	import type { Todo } from '$lib/types';
	import Icon from '@iconify/svelte';
	import { Icons } from '$lib/icons';

	export let status: Status;

	$: filteredTodos = $todos.filter((todo: Todo) => todo.status === status);
</script>

<div class="todo-list">
	<h2>
		{#if status === Status.Pending}
			<Icon icon={Icons.pending} />
		{:else if status === Status.InProgress}
			<Icon icon={Icons.inprogress} />
		{:else if status === Status.Completed}
			<Icon icon={Icons.completed} />
		{/if}
		{statusLabels[status]}
	</h2>

	<ul>
		{#each filteredTodos as todo (todo.id)}
			<ToDoCard {todo} />
		{/each}
	</ul>
</div>

<style lang="scss">
	.todo-list {
		border: 1px solid #ccc;
		border-radius: 5px;
		width: auto;
		flex-grow: 1;
	}

	h2 {
		display: inline-flex;
		align-items: center;
		justify-content: center;
		gap: 0.3em;
		height: 1rem;
		margin-bottom: 20px;
		width: 100%;
	}

	.todo-list {
		padding: 10px;
	}
</style>
