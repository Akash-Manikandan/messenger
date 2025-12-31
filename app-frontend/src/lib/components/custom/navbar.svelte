<script lang="ts">
	import ThemeSwitcher from "./theme-switcher.svelte";
</script>

<nav class="font-sans antialiased flex items-center justify-between p-4">
    <h1 class="text-xl font-bold">Messenger</h1>
    <ThemeSwitcher />
</nav>

<style>
	.font-sans {
		font-family: 'DM Sans', sans-serif;
	}
</style>