<script lang="ts">
	export let className: string = '';
	export let onClick: () => void;
</script>

<button class={className} on:click={onClick}>
	<slot />
</button>
