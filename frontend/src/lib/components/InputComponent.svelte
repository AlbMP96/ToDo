<script lang="ts">
	import Icon from '@iconify/svelte';

	export let type: string = 'text';
	export let inputId: string = '';
	export let labelText: string = '';
	export let icon: string = '';
	export let value: string = '';
	export let required: boolean = false;
	let showPassword: boolean = false;
</script>

<div>
	<label class="input-component" for={inputId}>
		{labelText}
		<input type={showPassword ? 'text' : type} id={inputId} bind:value {required} />
		{#if icon.trim() !== ''}
			<div class="icon-wrapper">
				<Icon {icon} width="20" height="20" />
			</div>
		{/if}
	</label>
	{#if type === 'password'}
		<label for="show-pass">
			Show Password
			<input type="checkbox" bind:checked={showPassword} name="show-pass" id="show-pass" />
		</label>
	{/if}
</div>

<style lang="scss">
	.input-component {
		display: flex;
		flex-direction: column;
		position: relative;
	}

	.icon-wrapper {
		position: absolute;
		padding-left: 10px;
		top: 50%;
	}

	input {
		padding: 0.5rem;
		padding-left: 35px;
		border: 1px solid #ccc;
		border-radius: 5px;
	}
</style>
