<script lang="ts">
	import { login } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import ButtonPrimary from '$lib/components/buttons/ButtonPrimary.svelte';

	let email: string = '';
	let password: string = '';
	let showPassword: boolean = false;

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
		<label class="email-field" for="email">
			Email:
			<input type="email" id="email" bind:value={email} required />
		</label>
		<div class="password-field">
			<label for="password">
				Password:
				<input
					type={showPassword ? 'text' : 'password'}
					id="password"
					bind:value={password}
					required
				/>
			</label>
			<label for="show-pass">
				Show Password
				<input type="checkbox" bind:checked={showPassword} name="show-pass" id="show-pass" />
			</label>
		</div>
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

		.email-field,
		.password-field > label:first-child {
			display: flex;
			flex-direction: column;
		}

		.password-field {
			display: flex;
			flex-direction: column;
			gap: 0.5rem;
		}

		input {
			padding: 0.5rem;
			border: 1px solid #ccc;
			border-radius: 4px;
		}
	}
</style>
