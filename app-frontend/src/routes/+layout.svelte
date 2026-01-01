<script lang="ts">
	import './layout.css';
	import favicon from '$lib/assets/favicon.svg';
	import { ModeWatcher } from 'mode-watcher';
	import Navbar from '$lib/components/custom/navbar.svelte';
	import * as Tooltip from '$lib/components/ui/tooltip';

	let { children, data } = $props();
	let user = $derived(data.user);
</script>

<svelte:head
	><link rel="icon" href={favicon} />
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link
		href="https://fonts.googleapis.com/css2?family=DM+Sans:ital,opsz,wght@0,9..40,100..1000;1,9..40,100..1000&family=Dosis:wght@200..800&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<Navbar {user} />
<ModeWatcher />

<main class="font-sans antialiased">
	<Tooltip.Provider>
		{@render children()}
	</Tooltip.Provider>
</main>

<style>
	.font-sans {
		font-family: 'DM Sans', sans-serif;
	}
</style>
