<script lang="ts">
	import type { UserResponse } from '$lib/api';
	import * as Avatar from '$lib/components/ui/avatar';
	import ThemeSwitcher from './theme-switcher.svelte';

	let { user }: { user: UserResponse } = $props();
</script>

<nav class="flex items-center justify-between p-4 font-sans antialiased">
	<h1 class="text-xl font-bold">Messenger</h1>
	<div class="flex items-center gap-4">
		<ThemeSwitcher />
		<div class="rounded-full border p-1">
			<Avatar.Root>
				<Avatar.Image src={user.avatar} alt={user.username} width={36} height={36} />
				<Avatar.Fallback>{user.username.slice(0, 2).toUpperCase()}</Avatar.Fallback>
			</Avatar.Root>
		</div>
	</div>
</nav>

<style>
	.font-sans {
		font-family: 'DM Sans', sans-serif;
	}
</style>
