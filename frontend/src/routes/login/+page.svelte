<script lang="ts">
	import { login } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import ButtonPrimary from '$lib/components/buttons/ButtonPrimary.svelte';

	let email: string = '';
	let password: string = '';

	async function sendLogin() {
		try {
			await login(email, password);
			goto('/');
		} catch (error) {
			console.error('Login failed:', error);
		}
	}
</script>

<div>
	<form on:submit|preventDefault={sendLogin}>
		<label>
			Email:
			<input type="email" bind:value={email} required />
		</label>
		<label>
			Password:
			<input type="password" bind:value={password} required />
		</label>
		<ButtonPrimary type="submit">Login</ButtonPrimary>
	</form>
</div>
