<script lang="ts">
	import Button from './Button.svelte';
	export let onClick: () => void;
</script>

<Button className="button-secondary" {onClick}>
	<slot />
</Button>
