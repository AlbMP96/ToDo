<script lang="ts">
	import { login } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import ButtonPrimary from '$lib/components/buttons/ButtonPrimary.svelte';
	import { Icons } from '$lib/icons';
	import InputComponent from '$lib/components/InputComponent.svelte';
	import FormBase from './FormBase.svelte';

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

<FormBase onSubmit={sendLogin}>
	<h2>Login</h2>
	<InputComponent type="email" bind:value={email} icon={Icons.user} labelText="Email" required />

	<InputComponent
		type="password"
		bind:value={password}
		icon={Icons.key}
		labelText="Password"
		required
	/>
	<ButtonPrimary type="submit">Login</ButtonPrimary>
	<div>
		<span
			>¿Aún no tienes cuenta? <a data-sveltekit-reload href="/register">Registrate ahora</a></span
		>
	</div>
</FormBase>
