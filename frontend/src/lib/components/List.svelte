<script lang="ts">
	import { todos } from '$lib/stores/todos';
	import { Status, statusLabels } from '$lib/enums/status';
	import ToDoCard from './ToDoCard.svelte';
	import type { Todo } from '$lib/types';

	export let status: Status;

	$: filteredTodos = $todos.filter((todo: Todo) => todo.status === status);
</script>

<div class="todo-list">
	<h2>{statusLabels[status]}</h2>

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
		width: 20%;
	}

	@media (max-width: 600px) {
		.todo-list {
			width: 100%;
		}
	}

	h2 {
		text-align: center;
	}

	.todo-list {
		padding: 10px;
	}
</style>
