<script lang="ts">
	import { page } from '$app/state';

	import LoginForm from '$lib/components/auth/LoginForm.svelte';
	import RegisterForm from '$lib/components/auth/RegisterForm.svelte';

	$: formType = page.url.pathname;
</script>

{#if formType === '/login'}
	<LoginForm />
{:else if formType === '/register'}
	<RegisterForm />
{:else}
	<p>Ruta no válida</p>
{/if}
