<script lang="ts">
	import { login } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import ButtonPrimary from '$lib/components/buttons/ButtonPrimary.svelte';
	import { Icons } from '$lib/icons';
	import InputComponent from '$lib/components/InputComponent.svelte';

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

<div class="login-page">
	<form on:submit|preventDefault={sendLogin}>
		<InputComponent
			type="email"
			inputId="email"
			bind:value={email}
			icon={Icons.user}
			labelText="Email:"
			required
		/>

		<InputComponent
			type="password"
			inputId="password"
			bind:value={password}
			icon={Icons.key}
			labelText="Password:"
			required
		/>
		<ButtonPrimary type="submit">Login</ButtonPrimary>
	</form>
</div>

<style lang="scss">
	.login-page {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		height: 100vh;

		form {
			background-color: oklch(86.9% 0.005 56.366);
			display: flex;
			flex-direction: column;
			gap: 1rem;
			margin: auto;
			padding: 2rem;
			border-radius: 10px;
			box-shadow: 0 0 10px oklch(70.9% 0.01 56.259);
		}
	}
</style>
