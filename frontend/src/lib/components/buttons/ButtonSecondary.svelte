<script lang="ts">
	import Button from './Button.svelte';
	export let onClick: () => void = () => {};
	export let type: 'button' | 'submit' = 'button';
</script>

<Button className="button-secondary" {onClick} {type}>
	<slot />
</Button>
