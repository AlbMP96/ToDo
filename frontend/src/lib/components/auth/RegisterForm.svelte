<script lang="ts">
	import FormBase from './FormBase.svelte';
	import { create, login } from '$lib/api/auth';
	import { goto } from '$app/navigation';
	import ButtonPrimary from '$lib/components/buttons/ButtonPrimary.svelte';
	import { Icons } from '$lib/icons';
	import InputComponent from '$lib/components/InputComponent.svelte';

	let email: string = '';
	let password: string = '';
	const regexEmail: RegExp = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
	const regexPassword: RegExp = /^[a-zA-Z0-9!@#$%^&*]{8,}$/;

	async function sendRegister() {
		if (regexEmail.test(email) && regexPassword.test(password)) {
			try {
				await create(email, password);
				await login(email, password);
				goto('/');
			} catch (error) {
				console.error('Login failed:', error);
			}
		} else {
			console.log('Campos no válidos');
		}
	}
</script>

<FormBase onSubmit={sendRegister}>
	<h2>Sign Up</h2>
	<InputComponent type="email" bind:value={email} icon={Icons.user} labelText="Email" required />

	<InputComponent
		type="password"
		bind:value={password}
		icon={Icons.key}
		labelText="Password"
		required
	/>

	<ButtonPrimary type="submit">Sign Up</ButtonPrimary>
	<div>
		<span>¿Ya tienes cuenta? <a data-sveltekit-reload href="/login">Inicia sesión</a></span>
	</div>
</FormBase>
