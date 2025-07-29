<script lang="ts">
	export let className: string = '';
	export let onClick: () => void = () => {};
	export let type: 'button' | 'submit' = 'button';
</script>

<button class={className} on:click={onClick} {type}>
	<slot />
</button>
