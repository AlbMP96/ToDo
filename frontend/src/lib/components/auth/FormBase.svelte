<script lang="ts">
	export let onSubmit: () => void;
</script>

<div class="login-page">
	<form on:submit|preventDefault={onSubmit}>
		<slot />
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
