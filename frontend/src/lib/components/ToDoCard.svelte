<script lang="ts">
	import { todos } from '$lib/stores/todos';
	import type { Todo } from '$lib/types';
	import { Status, statusLabels } from '$lib/enums/status';
	import Icon from '@iconify/svelte';

	export let todo: Todo;

	function updateTodoStatus(id: string, newStatus: Status) {
		todos.update((current: Todo[]) =>
			current.map((todo) => (todo.id === id ? { ...todo, status: newStatus } : todo))
		);
	}

	function deleteTodo(id: string) {
		todos.update((current: Todo[]) => current.filter((todo) => todo.id !== id));
	}
</script>

<li>
	<span>{todo.text}</span>
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
	<button on:click={() => deleteTodo(todo.id)}><Icon icon="fluent:delete-32-light" /></button>
</li>

<style lang="scss">
	li {
		box-shadow: 0 0 10px oklch(70.9% 0.01 56.259);
		padding: 5px;
		margin: 5px;
		border-radius: 5px;
		display: flex;
		justify-content: space-around;
	}
</style>
